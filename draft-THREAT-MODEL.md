<!--
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
-->

# Apache CloudStack Go SDK Security Threat Model — delta (draft)

> **Delta document.** This document is a delta over the canonical
> `apache/cloudstack` threat model. It inherits §3 (out-of-scope),
> §4 B1 (API trust boundary), and §7 (adversary model) from the
> main model. It restates only what `apache/cloudstack-go` uniquely
> introduces. Read the main model first.

## §1 Header

- **Project:** Apache CloudStack Go SDK (`apache/cloudstack-go`) — Go
  client library that calls the `apache/cloudstack` JSON API.
- **Commit:** `358fe85` (HEAD of `main` at draft time).
- **Date:** 2026-05-29.
- **Authors:** ASF Security team draft, awaiting CloudStack PMC review.
- **Status:** Draft delta over `cloudstack-threat-model-draft.md`.
- **Version binding:** as of the commit above. The SDK's release line is
  versioned independently of the management-server release line
  *(documented: `README.md` — "SDK releases tagged based on the ACS
  version")*.
- **Reporting:** same as the main model — `security@apache.org`.
- **Provenance legend:** as in the main model.
- **Draft confidence:** 9 documented / 0 maintainer / 11 inferred.

**About the project.** `cloudstack-go` is a generated-then-curated Go
library that wraps every CloudStack JSON API command in a typed
parameter struct and a typed response *(documented: `README.md`,
`generate/generate.go`)*. Its main moving parts: an HMAC-SHA1 signer
+ HTTP client (`cloudstack/cloudstack.go`) and one `<Service>.go` file
per API service group (Account, Address, VirtualMachine, …) generated
from `generate/listApis.json` against the latest stable CloudStack
release (currently v4.18.x baseline) *(documented: `README.md`)*.

## §2 Scope and intended use

**Primary intended use.** *(documented — README)* In-process import by
a Go program that wants to drive a CloudStack management server. The
embedding application instantiates a client with `NewClient(endpoint,
apiKey, secretKey, verifySSL)` or `NewAsyncClient(…)`, then issues
typed API calls.

**Deployment shape.** Always in-process; the SDK is a library, never
a daemon.

**Caller expectations.** The caller is trusted to:

- choose the management-server endpoint URL,
- supply a valid `apiKey` + `secretKey` pair,
- pass `verifySSL = true` *unless* connecting to a known dev cluster
  using a self-signed CloudStack Root CA cert,
- not source any of the above from end-user input.

**Component-family table.**

| Family | Representative entry | Touches outside the process? | In this delta? |
| --- | --- | --- | --- |
| Client + signer (`cloudstack/cloudstack.go`) | `NewAsyncClient`, signer in `cs.sign(params)` | **yes — network + creds** | yes |
| Per-service generated wrappers (`cloudstack/<Service>.go`, ~600 services) | `cs.VirtualMachine.DeployVirtualMachine(…)` | inherited from client | yes |
| Code generator (`generate/`) | reads `generate/listApis.json` and emits `cloudstack/*.go` | host filesystem at generation time | **out of model** *(§3)* — generator output is in-model, generator itself is build-time tooling |
| `examples/` | `examples/AddHost.go`, `examples/CreateDeleteDomain.go`, `mock_test.go` | network | **out of model** *(§3)* |
| `test/`, `ci/` | tests + CI scaffolding | varies | **out of model** *(§3)* |
| `cloudstack/*_mock.go` | generated mocks | none | **out of model** *(§3)* |

## §3 Out of scope (explicit non-goals)

The main model's §3 applies in full. **Additional** out-of-scope items
specific to the SDK:

1. **Storage of `apiKey` / `secretKey` on disk.** Where the caller
   chooses to persist credentials is out of model. *(inferred — Q1)*
2. **Server-side correctness of the management server's response.** If
   the management server returns wrong data, the SDK forwards it; the
   SDK is not a re-verification layer. *(inferred — Q2)*
3. **TLS transport configuration.** When the caller passes `verifySSL =
   false` (the fourth `NewAsyncClient` argument), the SDK explicitly
   sets `InsecureSkipVerify: true` *(documented:
   `cloudstack/cloudstack.go` line ~216
   `TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifyssl}`)*.
   That is a documented caller-supplied weakening; reports of "the SDK
   permits TLS verification to be disabled" are `OUT-OF-MODEL:
   trusted-input` because the caller chose it.
4. **`examples/`, `cloudstack/*_mock.go`, `test/`, `ci/`.**
5. **The four sibling repos** — `apache/cloudstack`,
   `apache/cloudstack-cloudmonkey`, `apache/cloudstack-terraform-provider`,
   `apache/cloudstack-kubernetes-provider`. Covered by their own models.

## §4 Trust boundaries and data flow

The only boundary the SDK contributes is **the embedded application's
call into `NewAsyncClient` (or `NewClient`) and subsequent typed API
calls**. Bytes the caller passes — endpoint, `apiKey`, `secretKey`,
verifySSL, per-method parameters — are caller-controlled config and not
attacker-controlled per the main-model §7 adversary set.

The bytes that come *back* from the management server (JSON responses)
are trusted control-plane content per the main model's B1 transition.

**Per-call request shape** *(documented: `cloudstack/cloudstack.go`
lines ~547–575)*:

- the SDK forces `signatureversion=3`, which mandates an `expires`
  parameter (good — see main-model §5a "api.signature.version" row);
- the SDK computes HMAC-SHA1 of the lowercase-sorted parameter string
  under `secretKey` (`crypto/hmac` + `crypto/sha1`);
- the signature is base64-encoded and placed in either the POST body
  or the URL (when `HTTPGETOnly` is set);
- when `HTTPGETOnly = true`, the SDK constructs the URL as
  `baseURL + "?" + paramString + "&signature=" + url.QueryEscape(sig)`
  — note the signature lands in the *URL* (and thus the access log /
  reverse proxy log) for GET requests *(inferred — Q3)*.

## §5 Assumptions about the environment

- **Go version**: per `go.mod` (`go 1.x`) — confirm with maintainer
  *(inferred — Q4)*.
- **TLS**: provided by `crypto/tls` from the Go stdlib; verification
  default is on, weakened only by the explicit `verifyssl=false` caller
  choice.
- **Transport**: `net/http` with stdlib defaults; no connection pool
  sharing across `cloudstack.Client` instances *(inferred — Q5)*.
- **Concurrency**: `cloudstack.Client` is intended to be safe for
  concurrent use *(inferred — Q5)*.
- **What the SDK does not do to its host**: no global state mutation,
  no signal handlers, no env-var consumption (credentials are passed
  in via the constructor, not read from env). *(inferred — Q6)*

## §5a Build-time and configuration variants

| Knob | Default | Stance | Effect |
| --- | --- | --- | --- |
| `verifyssl` arg to `NewClient` / `NewAsyncClient` | passed-in *(no default)* | calling with `false` is documented for dev / self-signed-CA setups | when `false`, `InsecureSkipVerify: true` is set in the TLS config |
| `cs.HTTPGETOnly` | `false` *(inferred — Q3)* | the Kubernetes / Terraform satellites set this `true` to force GET, which inlines signatures in the URL | when `true`, requests are GET with signature in URL |
| async timeout via `cs.AsyncTimeout(…)` | `300s` default *(inferred — Q7)* | bounds how long async jobs are polled | timeout exhaustion returns an error, not a hang |

## §6 Assumptions about inputs

| Entry point | Parameter | Attacker-controllable in the model? | Caller must enforce |
| --- | --- | --- | --- |
| `NewClient` / `NewAsyncClient` | `endpoint`, `apiKey`, `secretKey`, `verifyssl` | **no** — operator config per §3 item 1 | sanity-check the endpoint URL; do not source from end-user input |
| Per-service typed call (e.g. `VirtualMachine.NewDeployVirtualMachineParams`) | typed params | **no** — typed params come from the embedding app's own logic | the embedding app applies its own user authn/authz before constructing params |
| HTTP response | JSON body | trusted as far as the management server is trusted | bytes are JSON-decoded into typed responses |

## §7 Adversary model

Main-model §7 applies. **Adjustments specific to the SDK**:

- The "unauthenticated network peer reaching `:8080`/`:8443`" actor is
  *upstream* of the SDK and not in scope for SDK code.
- An additional adversary worth naming: **a passive observer of HTTP
  access logs / reverse-proxy logs** when `HTTPGETOnly = true`. The
  signature lands in the URL and thus in any log that records URLs.
  Whether this is a §10 caller responsibility or a §11 misuse is the
  Q3 ruling. *(inferred — Q3)*

## §8 Security properties the SDK provides

### S1 — HMAC-SHA1 signature with `signatureversion=3`

- **Property.** Every API call carries a v3 signature with an `expires`
  parameter, computed via `crypto/hmac` + `crypto/sha1` over the
  canonical parameter string. *(documented: `cloudstack/cloudstack.go`
  lines ~547–575)*
- **Conditions.** Caller supplied a valid `secretKey`.
- **Violation symptom.** A request leaves the SDK without a signature,
  or with a signature derived from input outside the canonical
  parameter string.
- **Severity.** Security-critical, `VALID` per main-model §13.

### S2 — TLS verification on by default

- **Property.** The HTTP client uses stdlib TLS defaults; verification
  is on unless the caller passes `verifyssl = false`.
- **Conditions.** Caller did not pass `verifyssl = false`.
- **Violation symptom.** TLS verification skipped despite `verifyssl =
  true`.
- **Severity.** Security-critical, `VALID`.

### S3 — Constant-time secret-handling at the signer

- **Property.** *(inferred — Q8)* The SDK does not compare secrets, so
  the question is whether the signer leaks the secret via timing /
  branch-predictable paths. The Go `crypto/hmac` is the only consumer.
- **Conditions.** as above.
- **Violation symptom.** Side-channel recovery of `secretKey`.
- **Severity.** Security-critical if reachable; otherwise `OUT-OF-MODEL`
  per main-model §7 side-channel exclusion.

## §9 Security properties the SDK does *not* provide

- **No protection of `apiKey` / `secretKey` at rest.** The caller decides
  where to store them. *(inferred — Q1)*
- **No defence when `verifyssl = false`.** Explicit caller-supplied
  TLS-verification disablement.
- **No retry / rate-limit / circuit-breaker by default.** The async
  client polls; that is the only built-in scheduling.
- **No re-validation of management-server response correctness.**
- **No log-redaction of signatures or credentials by the SDK.** If the
  embedding application logs request URLs (which would include the
  signature when `HTTPGETOnly = true`), the signature ends up in the
  log. *(inferred — Q3)*

### False-friend properties

- **`HTTPGETOnly = true` is a transport choice, not "more secure".** It
  forces signatures into the URL where they will appear in any URL log.
- **"HMAC-SHA1" is not a SHA1-broken construction.** Collision attacks
  on SHA1 do not extend to HMAC-SHA1. Reports flagging "SHA1 is
  deprecated" against the signer are `KNOWN-NON-FINDING` per the main
  model's §11a.

## §10 Downstream responsibilities

The embedding Go application MUST:

1. Pass `verifyssl = true` unless connecting to a known dev cluster.
2. Source `apiKey` / `secretKey` from a secret store, not from
   user-controlled input.
3. Not log full request URLs when `HTTPGETOnly = true`.
4. Apply its own timeouts and retries on top of the SDK as appropriate.
5. Use the SDK version matched to the management-server version
   *(documented: README — "SDK releases tagged based on the ACS
   version")*.

## §11 Known misuse patterns

- Calling `NewAsyncClient(endpoint, key, secret, false)` in production
  (TLS verification disabled).
- Embedding credentials in source / build artifacts.
- Using `HTTPGETOnly = true` *and* shipping URL logs off-host.
- Using a mismatched SDK / server major version and assuming the
  request shape will still verify.

## §11a Known non-findings (recurring false positives)

- **"HMAC-SHA1 — SHA1 is deprecated."** → `KNOWN-NON-FINDING` per main
  model §11a.
- **"`InsecureSkipVerify: true` is hardcoded."** It is gated by the
  caller-supplied `verifyssl` argument. → `OUT-OF-MODEL: trusted-input`.
- **"Signature appears in URL."** Only when caller sets `HTTPGETOnly =
  true`. → `OUT-OF-MODEL: non-default-build` if `false` is the default
  per Q3. Else escalate.
- **"Examples in `examples/` have weak input handling."** Out of
  scope. → `OUT-OF-MODEL: unsupported-component`.

## §12 Conditions that would change this delta

- Change in default signing algorithm in the main model (SHA1 → SHA256).
- Change in `HTTPGETOnly` default or new option that affects signature
  placement.
- Addition of credential storage to the SDK.

## §13 Triage dispositions

Use the same table as the main model.

## §14 Open questions for the maintainers

**Q1.** Out-of-scope: where the caller stores `apiKey` / `secretKey`
on disk. Confirm.

**Q2.** Out-of-scope: revalidating management-server response
correctness in the SDK. Confirm.

**Q3.** `HTTPGETOnly` default and signature-in-URL leakage — is `false`
the default and is "do not log URLs when `HTTPGETOnly = true`" a
documented caller responsibility? *(maps to §5a, §6, §10, §11a)*

**Q4.** Confirm the Go MSRV.

**Q5.** Concurrency: is `cloudstack.Client` documented as safe for
concurrent use across goroutines? Is the underlying `http.Client`
shared or per-call?

**Q6.** Confirm the §5 negative side-effect inventory: no env-var
consumption, no signal handlers, no global mutation.

**Q7.** Confirm `AsyncTimeout` default.

**Q8.** Constant-time / side-channel posture at the signer — confirm
the SDK delegates entirely to Go stdlib `crypto/hmac` and makes no
additional claim.

**Q9.** Meta — should this delta document live in `apache/cloudstack-go`
at `docs/threat-model.md`, or only on the website?

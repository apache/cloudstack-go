# cloudstack-go

[![Go Reference](https://pkg.go.dev/badge/github.com/apache/cloudstack-go/v2/cloudstack.svg)](https://pkg.go.dev/github.com/apache/cloudstack-go/v2/cloudstack)

A CloudStack API client enabling Go programs to interact with CloudStack in a simple and uniform way

## Status

This package covers the complete CloudStack API and is well tested. There may be some untested corner cases as there are over 600 API commands, but over all, it's safe to use this package.

To be able to find the API command you want, they are grouped by 'services' which match the grouping you can see/find on the [CloudStack API docs](https://cloudstack.apache.org/api/apidocs-4.18/) website.

## Usage

The cloudstack package is always generated against the latest stable CloudStack release (currently v4.15.x). As long as the API changes were minimum across subsequent CloudStack releases it was possible to have the generator code handle API changes such that they were backward compatible.
However, over time, with frequent API changes observed across releases of Apache CloudStack(ACS), we will have the SDK releases tagged based on the ACS version.

Please see the package documentation on [go.dev](https://pkg.go.dev/github.com/apache/cloudstack-go/v2/cloudstack).

## Example

```go
// Create a new API client
cs := cloudstack.NewAsyncClient("https://cloudstack.company.com", "your-api-key", "your-api-secret", false)

// Create a new parameter struct
p := cs.VirtualMachine.NewDeployVirtualMachineParams("service-offering-id", "template-id", "zone-id")

// Set the name
name := "server-1"
p.SetName(name)

// Set the display name
p.SetDisplayname("Test server 1")

// Set any other options required by your setup in the same way

// Create the new instance
r, err := cs.VirtualMachine.DeployVirtualMachine(p)
if err != nil {
	log.Fatalf("Error creating the new instance %s: %s", name, err)
}

fmt.Printf("UUID or the newly created machine: %s", r.Id)
```

## Features

Apart from the API commands CloudStack itself offers, there are a few additional features/functions that are helpful. For starters, there are two clients, a normal one (created with `NewClient(...)`) and an async client (created with `NewAsyncClient(...)`). The async client has a builtin waiting/polling feature that waits for a configured amount of time (defaults to 300 seconds) on running async jobs. This is very helpful if you do not want to continue with your program execution until the async job is done.

There is also a function that can be called manually (`GetAsyncJobResult(...)`) that does the same, but then as a separate call after the async job has started.

Another nice feature is the fact that for every API command you can create the needed parameter struct using a `New...Params` function, like for example `NewListTemplatesParams`. The advantage of using this functions to create a new parameter struct, is that these functions know what the required parameters are for every API command, and they require you to supply these when creating the new struct. Every additional parameter can be set after creating the struct by using the appropriate setters, e.g., `SetName()`.

Last but not the least, there are a lot of helper functions that will try to automatically find a UUID for you for various resources (disk, template, virtualmachine, network...). This makes it much easier and faster to work with the API commands and in most cases you can just use then if you know the name instead of the UUID.

## Developer Guide

The SDK relies on the  `generate.go` script to auto generate the code for all the supported APIs listed in the `listApis.json` file.
The `listAPIs.json` file holds the output of `listApis` command for a specific release of CloudStack.

```
# Run it via the Makefile
make all

```

## Getting Help

_Please try to see if the [module documentation](https://pkg.go.dev/github.com/apache/cloudstack-go/v2/cloudstack) can provide some answers first!_

* If you have an issue: report it on the [issue tracker](https://github.com/apache/cloudstack-go/issues)

## History

Sander van Harmelen (<sander@vanharmelen.nl>) was the original author of this repo
which was donated to the Apache CloudStack project under an [IP
clearance](https://github.com/apache/cloudstack/issues/5159) process.

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at <http://www.apache.org/licenses/LICENSE-2.0>

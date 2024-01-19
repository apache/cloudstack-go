Our goal is to add the follwoing api call to cloudstack-go:
https://cloudstack.apache.org/api/apidocs-4.18/apis/createConsoleEndpoint.html

Request parameters
virtualmachineid	(mandatory) ID of the VM
token				(optional) extra security token, valid when the extra validation is enabled

Response tags
details		details in case of an error
success		true if the console endpoint is generated properly
url			the console url
websocket	the console websocket options

If I understand it correctly the correct file to add this api call is this file:
cloudstack-go/blob/main/cloudstack/VirtualMachineService.go
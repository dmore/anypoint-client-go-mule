# Go API client for connected_app

Description of the Connected App API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./connected_app"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://anypoint.mulesoft.com/accounts/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ConnectedApplicationsConnAppIdDelete**](docs/DefaultApi.md#connectedapplicationsconnappiddelete) | **Delete** /connectedApplications/{connAppId} | 
*DefaultApi* | [**ConnectedApplicationsConnAppIdGet**](docs/DefaultApi.md#connectedapplicationsconnappidget) | **Get** /connectedApplications/{connAppId} | 
*DefaultApi* | [**ConnectedApplicationsConnAppIdPatch**](docs/DefaultApi.md#connectedapplicationsconnappidpatch) | **Patch** /connectedApplications/{connAppId} | 
*DefaultApi* | [**ConnectedApplicationsConnAppIdScopesGet**](docs/DefaultApi.md#connectedapplicationsconnappidscopesget) | **Get** /connectedApplications/{connAppId}/scopes | 
*DefaultApi* | [**ConnectedApplicationsConnAppIdScopesPut**](docs/DefaultApi.md#connectedapplicationsconnappidscopesput) | **Put** /connectedApplications/{connAppId}/scopes | 
*DefaultApi* | [**ConnectedApplicationsGet**](docs/DefaultApi.md#connectedapplicationsget) | **Get** /connectedApplications | 
*DefaultApi* | [**ConnectedApplicationsPost**](docs/DefaultApi.md#connectedapplicationspost) | **Post** /connectedApplications | 


## Documentation For Models

 - [ConnectedAppCore](docs/ConnectedAppCore.md)
 - [ConnectedAppPatchExt](docs/ConnectedAppPatchExt.md)
 - [ConnectedAppPatchExtAllOf](docs/ConnectedAppPatchExtAllOf.md)
 - [ConnectedAppRespExt](docs/ConnectedAppRespExt.md)
 - [ConnectedAppRespExtAllOf](docs/ConnectedAppRespExtAllOf.md)
 - [ConnectedAppScopesPutBody](docs/ConnectedAppScopesPutBody.md)
 - [ContextParams](docs/ContextParams.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [ScopeCore](docs/ScopeCore.md)


## Documentation For Authorization



### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




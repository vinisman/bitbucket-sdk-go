# \MirroringUpstreamAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Accept**](MirroringUpstreamAPI.md#Accept) | **Post** /mirroring/latest/requests/{mirroringRequestId}/accept | Accept a mirroring request
[**AnalyticsSettings**](MirroringUpstreamAPI.md#AnalyticsSettings) | **Get** /mirroring/latest/analyticsSettings | Get analytics settings from upstream
[**DeleteMirroringRequest**](MirroringUpstreamAPI.md#DeleteMirroringRequest) | **Delete** /mirroring/latest/requests/{mirroringRequestId} | Delete a mirroring request
[**DeletePreferredMirrorId**](MirroringUpstreamAPI.md#DeletePreferredMirrorId) | **Delete** /mirroring/latest/account/settings/preferred-mirror | Remove preferred mirror
[**GetAllContentHashes**](MirroringUpstreamAPI.md#GetAllContentHashes) | **Get** /mirroring/latest/repos | Get content hashes for repositories
[**GetAllReposForProject**](MirroringUpstreamAPI.md#GetAllReposForProject) | **Get** /mirroring/latest/projects/{projectId}/repos | Get hashes for repositories in project
[**GetContentHashById**](MirroringUpstreamAPI.md#GetContentHashById) | **Get** /mirroring/latest/repos/{repoId} | Get content hash for a repository
[**GetMirror**](MirroringUpstreamAPI.md#GetMirror) | **Get** /mirroring/latest/mirrorServers/{mirrorId} | Get mirror by ID
[**GetMirroringRequest**](MirroringUpstreamAPI.md#GetMirroringRequest) | **Get** /mirroring/latest/requests/{mirroringRequestId} | Get a mirroring request
[**GetPreferredMirrorId**](MirroringUpstreamAPI.md#GetPreferredMirrorId) | **Get** /mirroring/latest/account/settings/preferred-mirror | Get preferred mirror
[**GetProjectById**](MirroringUpstreamAPI.md#GetProjectById) | **Get** /mirroring/latest/projects/{projectId} | Get project
[**GetRepositoryMirrors**](MirroringUpstreamAPI.md#GetRepositoryMirrors) | **Get** /mirroring/latest/repos/{repoId}/mirrors | Get mirrors for repository
[**ListMirrors**](MirroringUpstreamAPI.md#ListMirrors) | **Get** /mirroring/latest/mirrorServers | Get all mirrors
[**ListRequests**](MirroringUpstreamAPI.md#ListRequests) | **Get** /mirroring/latest/requests | Get mirroring requests
[**MAuthenticate**](MirroringUpstreamAPI.md#MAuthenticate) | **Post** /mirroring/latest/authenticate | Authenticate on behalf of a user
[**PublishEvent**](MirroringUpstreamAPI.md#PublishEvent) | **Post** /mirroring/latest/mirrorServers/{mirrorId}/events | Publish RepositoryMirrorEvent
[**Register**](MirroringUpstreamAPI.md#Register) | **Post** /mirroring/latest/requests | Create a mirroring request
[**Reject**](MirroringUpstreamAPI.md#Reject) | **Post** /mirroring/latest/requests/{mirroringRequestId}/reject | Reject a mirroring request
[**Remove**](MirroringUpstreamAPI.md#Remove) | **Delete** /mirroring/latest/mirrorServers/{mirrorId} | Delete mirror by ID
[**SetPreferredMirrorId**](MirroringUpstreamAPI.md#SetPreferredMirrorId) | **Post** /mirroring/latest/account/settings/preferred-mirror | Set preferred mirror
[**Upgrade**](MirroringUpstreamAPI.md#Upgrade) | **Put** /mirroring/latest/mirrorServers/{mirrorId} | Upgrade mirror server



## Accept

> RestMirrorServer Accept(ctx, mirroringRequestId).Execute()

Accept a mirroring request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	mirroringRequestId := "mirroringRequestId_example" // string | the ID of the request to accept

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.Accept(context.Background(), mirroringRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.Accept``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Accept`: RestMirrorServer
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.Accept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mirroringRequestId** | **string** | the ID of the request to accept | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestMirrorServer**](RestMirrorServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsSettings

> RestAnalyticsSettings AnalyticsSettings(ctx).Execute()

Get analytics settings from upstream



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.AnalyticsSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.AnalyticsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsSettings`: RestAnalyticsSettings
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.AnalyticsSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsSettingsRequest struct via the builder pattern


### Return type

[**RestAnalyticsSettings**](RestAnalyticsSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMirroringRequest

> DeleteMirroringRequest(ctx, mirroringRequestId).Execute()

Delete a mirroring request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	mirroringRequestId := "mirroringRequestId_example" // string | the ID of the mirroring request to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MirroringUpstreamAPI.DeleteMirroringRequest(context.Background(), mirroringRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.DeleteMirroringRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mirroringRequestId** | **string** | the ID of the mirroring request to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMirroringRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePreferredMirrorId

> DeletePreferredMirrorId(ctx).Execute()

Remove preferred mirror



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MirroringUpstreamAPI.DeletePreferredMirrorId(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.DeletePreferredMirrorId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePreferredMirrorIdRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllContentHashes

> EnrichedRepository GetAllContentHashes(ctx).IncludeDefaultBranch(includeDefaultBranch).Execute()

Get content hashes for repositories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	includeDefaultBranch := "includeDefaultBranch_example" // string | includes defaultBranchId for each repository in the response, if <code>true</code>. Default value is <code>false</code>. (optional) (default to "false")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.GetAllContentHashes(context.Background()).IncludeDefaultBranch(includeDefaultBranch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.GetAllContentHashes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllContentHashes`: EnrichedRepository
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.GetAllContentHashes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllContentHashesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDefaultBranch** | **string** | includes defaultBranchId for each repository in the response, if &lt;code&gt;true&lt;/code&gt;. Default value is &lt;code&gt;false&lt;/code&gt;. | [default to &quot;false&quot;]

### Return type

[**EnrichedRepository**](EnrichedRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllReposForProject

> GetAllReposForProject200Response GetAllReposForProject(ctx, projectId).IncludeDefaultBranch(includeDefaultBranch).Start(start).Limit(limit).Execute()

Get hashes for repositories in project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	projectId := "projectId_example" // string | the id of the requested project
	includeDefaultBranch := "includeDefaultBranch_example" // string | includes defaultBranchId in the response, if <code>true</code>. Default value is <code>false</code> (optional) (default to "false")
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.GetAllReposForProject(context.Background(), projectId).IncludeDefaultBranch(includeDefaultBranch).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.GetAllReposForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllReposForProject`: GetAllReposForProject200Response
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.GetAllReposForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the id of the requested project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReposForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDefaultBranch** | **string** | includes defaultBranchId in the response, if &lt;code&gt;true&lt;/code&gt;. Default value is &lt;code&gt;false&lt;/code&gt; | [default to &quot;false&quot;]
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetAllReposForProject200Response**](GetAllReposForProject200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentHashById

> EnrichedRepository GetContentHashById(ctx, repoId).IncludeDefaultBranch(includeDefaultBranch).Execute()

Get content hash for a repository



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	repoId := "repoId_example" // string | the ID of the requested repository
	includeDefaultBranch := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.GetContentHashById(context.Background(), repoId).IncludeDefaultBranch(includeDefaultBranch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.GetContentHashById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentHashById`: EnrichedRepository
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.GetContentHashById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoId** | **string** | the ID of the requested repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentHashByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDefaultBranch** | **bool** |  | [default to false]

### Return type

[**EnrichedRepository**](EnrichedRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMirror

> RestMirrorServer GetMirror(ctx, mirrorId).Execute()

Get mirror by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	mirrorId := "mirrorId_example" // string | the mirror ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.GetMirror(context.Background(), mirrorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.GetMirror``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMirror`: RestMirrorServer
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.GetMirror`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mirrorId** | **string** | the mirror ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMirrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestMirrorServer**](RestMirrorServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMirroringRequest

> RestMirroringRequest GetMirroringRequest(ctx, mirroringRequestId).Execute()

Get a mirroring request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	mirroringRequestId := "mirroringRequestId_example" // string | the ID of the mirroring request to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.GetMirroringRequest(context.Background(), mirroringRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.GetMirroringRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMirroringRequest`: RestMirroringRequest
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.GetMirroringRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mirroringRequestId** | **string** | the ID of the mirroring request to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMirroringRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestMirroringRequest**](RestMirroringRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreferredMirrorId

> RestMirrorServer GetPreferredMirrorId(ctx).Execute()

Get preferred mirror



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.GetPreferredMirrorId(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.GetPreferredMirrorId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPreferredMirrorId`: RestMirrorServer
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.GetPreferredMirrorId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreferredMirrorIdRequest struct via the builder pattern


### Return type

[**RestMirrorServer**](RestMirrorServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectById

> RestProject GetProjectById(ctx, projectId).Execute()

Get project



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	projectId := "projectId_example" // string | the ID of the requested project

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.GetProjectById(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.GetProjectById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectById`: RestProject
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.GetProjectById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the ID of the requested project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestProject**](RestProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryMirrors

> RestMirroredRepositoryDescriptor GetRepositoryMirrors(ctx, repoId).PreAuthorized(preAuthorized).Execute()

Get mirrors for repository



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	repoId := "repoId_example" // string | the ID of the requested repository
	preAuthorized := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.GetRepositoryMirrors(context.Background(), repoId).PreAuthorized(preAuthorized).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.GetRepositoryMirrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryMirrors`: RestMirroredRepositoryDescriptor
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.GetRepositoryMirrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoId** | **string** | the ID of the requested repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryMirrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preAuthorized** | **bool** |  | 

### Return type

[**RestMirroredRepositoryDescriptor**](RestMirroredRepositoryDescriptor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMirrors

> ListMirrors200Response ListMirrors(ctx).Start(start).Limit(limit).Execute()

Get all mirrors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.ListMirrors(context.Background()).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.ListMirrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMirrors`: ListMirrors200Response
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.ListMirrors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMirrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**ListMirrors200Response**](ListMirrors200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequests

> ListRequests200Response ListRequests(ctx).State(state).Start(start).Limit(limit).Execute()

Get mirroring requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	state := "state_example" // string | (optional) the request state to filter on (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.ListRequests(context.Background()).State(state).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.ListRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRequests`: ListRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.ListRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | (optional) the request state to filter on | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**ListRequests200Response**](ListRequests200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MAuthenticate

> RestApplicationUserWithPermissions MAuthenticate(ctx).RestAuthenticationRequest(restAuthenticationRequest).Execute()

Authenticate on behalf of a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restAuthenticationRequest := *openapiclient.NewRestAuthenticationRequest(openapiclient.Credentials{RestBearerTokenCredentials: openapiclient.NewRestBearerTokenCredentials("<token>")}) // RestAuthenticationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.MAuthenticate(context.Background()).RestAuthenticationRequest(restAuthenticationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.MAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MAuthenticate`: RestApplicationUserWithPermissions
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.MAuthenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restAuthenticationRequest** | [**RestAuthenticationRequest**](RestAuthenticationRequest.md) |  | 

### Return type

[**RestApplicationUserWithPermissions**](RestApplicationUserWithPermissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishEvent

> PublishEvent(ctx, mirrorId).RestRepositoryMirrorEvent(restRepositoryMirrorEvent).Execute()

Publish RepositoryMirrorEvent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	mirrorId := "mirrorId_example" // string | the server id of the mirror that raised this event
	restRepositoryMirrorEvent := *openapiclient.NewRestRepositoryMirrorEvent() // RestRepositoryMirrorEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MirroringUpstreamAPI.PublishEvent(context.Background(), mirrorId).RestRepositoryMirrorEvent(restRepositoryMirrorEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.PublishEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mirrorId** | **string** | the server id of the mirror that raised this event | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restRepositoryMirrorEvent** | [**RestRepositoryMirrorEvent**](RestRepositoryMirrorEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> RestMirroringRequest Register(ctx).RestMirroringRequest(restMirroringRequest).Execute()

Create a mirroring request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restMirroringRequest := *openapiclient.NewRestMirroringRequest() // RestMirroringRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.Register(context.Background()).RestMirroringRequest(restMirroringRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.Register``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Register`: RestMirroringRequest
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.Register`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restMirroringRequest** | [**RestMirroringRequest**](RestMirroringRequest.md) |  | 

### Return type

[**RestMirroringRequest**](RestMirroringRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reject

> RestMirrorServer Reject(ctx, mirroringRequestId).Execute()

Reject a mirroring request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	mirroringRequestId := "mirroringRequestId_example" // string | the ID of the request to reject

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.Reject(context.Background(), mirroringRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.Reject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Reject`: RestMirrorServer
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.Reject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mirroringRequestId** | **string** | the ID of the request to reject | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestMirrorServer**](RestMirrorServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Remove

> Remove(ctx, mirrorId).Execute()

Delete mirror by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	mirrorId := "mirrorId_example" // string | the ID of the mirror to remove

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MirroringUpstreamAPI.Remove(context.Background(), mirrorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.Remove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mirrorId** | **string** | the ID of the mirror to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPreferredMirrorId

> SetPreferredMirrorId(ctx).Body(body).Execute()

Set preferred mirror



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	body := "body_example" // string | the mirror ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MirroringUpstreamAPI.SetPreferredMirrorId(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.SetPreferredMirrorId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPreferredMirrorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | the mirror ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upgrade

> RestMirrorServer Upgrade(ctx, mirrorId).RestMirrorUpgradeRequest(restMirrorUpgradeRequest).Execute()

Upgrade mirror server



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	mirrorId := "mirrorId_example" // string | the ID of the mirror to upgrade
	restMirrorUpgradeRequest := *openapiclient.NewRestMirrorUpgradeRequest() // RestMirrorUpgradeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringUpstreamAPI.Upgrade(context.Background(), mirrorId).RestMirrorUpgradeRequest(restMirrorUpgradeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringUpstreamAPI.Upgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Upgrade`: RestMirrorServer
	fmt.Fprintf(os.Stdout, "Response from `MirroringUpstreamAPI.Upgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mirrorId** | **string** | the ID of the mirror to upgrade | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restMirrorUpgradeRequest** | [**RestMirrorUpgradeRequest**](RestMirrorUpgradeRequest.md) |  | 

### Return type

[**RestMirrorServer**](RestMirrorServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


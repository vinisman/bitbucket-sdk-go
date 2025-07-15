# \MirroringMirrorAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndRollingUpgrade**](MirroringMirrorAPI.md#EndRollingUpgrade) | **Post** /mirroring/latest/zdu/end | End ZDU upgrade on mirror farm
[**GetDelayedSyncRepositories**](MirroringMirrorAPI.md#GetDelayedSyncRepositories) | **Get** /mirroring/latest/mirrorRepos/delayed-sync | Get delayed sync repositories
[**GetFarmNodes**](MirroringMirrorAPI.md#GetFarmNodes) | **Get** /mirroring/latest/farmNodes | Get farm nodes
[**GetMirrorMode**](MirroringMirrorAPI.md#GetMirrorMode) | **Get** /mirroring/latest/syncSettings/mode | Get mirror mode
[**GetMirrorSettings**](MirroringMirrorAPI.md#GetMirrorSettings) | **Get** /mirroring/latest/syncSettings | Get upstream settings
[**GetMirroredProjects**](MirroringMirrorAPI.md#GetMirroredProjects) | **Get** /mirroring/latest/syncSettings/projects | Get mirrored project IDs
[**GetMirroredRepository**](MirroringMirrorAPI.md#GetMirroredRepository) | **Get** /mirroring/latest/mirrorRepos/{externalRepositoryId} | Get clone URLs
[**GetRefChangesQueue**](MirroringMirrorAPI.md#GetRefChangesQueue) | **Get** /mirroring/latest/supportInfo/refChangesQueue | Get items in ref changes queue
[**GetRefChangesQueueCount**](MirroringMirrorAPI.md#GetRefChangesQueueCount) | **Get** /mirroring/latest/supportInfo/refChangesQueue/count | Get total number of items in ref changes queue
[**GetRepoSyncStatus**](MirroringMirrorAPI.md#GetRepoSyncStatus) | **Get** /mirroring/latest/supportInfo/repoSyncStatus | Get sync status of repositories
[**GetRepoSyncStatus1**](MirroringMirrorAPI.md#GetRepoSyncStatus1) | **Get** /mirroring/latest/supportInfo/projects/{projectKey}/repos/{repositorySlug}/repoSyncStatus | Gets information about the mirrored repository
[**GetRepositoryLockOwner**](MirroringMirrorAPI.md#GetRepositoryLockOwner) | **Get** /mirroring/latest/supportInfo/projects/{projectKey}/repos/{repositorySlug}/repo-lock-owner | Get the repository lock owner for the syncing process
[**GetRepositoryLockOwners**](MirroringMirrorAPI.md#GetRepositoryLockOwners) | **Get** /mirroring/latest/supportInfo/repo-lock-owners | Get all the repository lock owners for the syncing process
[**GetSynchronizationProgress**](MirroringMirrorAPI.md#GetSynchronizationProgress) | **Get** /mirroring/latest/progress | Get synchronization progress state
[**GetUpstreamServer**](MirroringMirrorAPI.md#GetUpstreamServer) | **Get** /mirroring/latest/upstreamServer | Get upstream server
[**SetMirrorMode**](MirroringMirrorAPI.md#SetMirrorMode) | **Put** /mirroring/latest/syncSettings/mode | Update mirror mode
[**SetMirrorSettings**](MirroringMirrorAPI.md#SetMirrorSettings) | **Put** /mirroring/latest/syncSettings | Update upstream settings
[**StartMirroringProject**](MirroringMirrorAPI.md#StartMirroringProject) | **Post** /mirroring/latest/syncSettings/projects/{projectId} | Add project to be mirrored
[**StartMirroringProjects**](MirroringMirrorAPI.md#StartMirroringProjects) | **Post** /mirroring/latest/syncSettings/projects | Add multiple projects to be mirrored
[**StartRollingUpgrade**](MirroringMirrorAPI.md#StartRollingUpgrade) | **Post** /mirroring/latest/zdu/start | Start ZDU upgrade on mirror farm
[**StopMirroringProject**](MirroringMirrorAPI.md#StopMirroringProject) | **Delete** /mirroring/latest/syncSettings/projects/{projectId} | Stop mirroring project



## EndRollingUpgrade

> RestRollingUpgradeState EndRollingUpgrade(ctx).Execute()

End ZDU upgrade on mirror farm



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
	resp, r, err := apiClient.MirroringMirrorAPI.EndRollingUpgrade(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.EndRollingUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndRollingUpgrade`: RestRollingUpgradeState
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.EndRollingUpgrade`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEndRollingUpgradeRequest struct via the builder pattern


### Return type

[**RestRollingUpgradeState**](RestRollingUpgradeState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDelayedSyncRepositories

> []RestDelayedSyncRepository GetDelayedSyncRepositories(ctx).DelayThreshold(delayThreshold).Limit(limit).Execute()

Get delayed sync repositories



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
	delayThreshold := "delayThreshold_example" // string | Returns only those repositories that are delayed for the given duration. The minimum allowed value is the configured value for the property <code>plugin.mirroring.synchronization.interval</code> (optional)
	limit := "limit_example" // string | Limit the number of delayed sync repositories returned, the maximum allowed value is 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringMirrorAPI.GetDelayedSyncRepositories(context.Background()).DelayThreshold(delayThreshold).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetDelayedSyncRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDelayedSyncRepositories`: []RestDelayedSyncRepository
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetDelayedSyncRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDelayedSyncRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delayThreshold** | **string** | Returns only those repositories that are delayed for the given duration. The minimum allowed value is the configured value for the property &lt;code&gt;plugin.mirroring.synchronization.interval&lt;/code&gt; | 
 **limit** | **string** | Limit the number of delayed sync repositories returned, the maximum allowed value is 100 | 

### Return type

[**[]RestDelayedSyncRepository**](RestDelayedSyncRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFarmNodes

> []RestClusterNode GetFarmNodes(ctx).Execute()

Get farm nodes



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
	resp, r, err := apiClient.MirroringMirrorAPI.GetFarmNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetFarmNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFarmNodes`: []RestClusterNode
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetFarmNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFarmNodesRequest struct via the builder pattern


### Return type

[**[]RestClusterNode**](RestClusterNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMirrorMode

> GetMirrorMode(ctx).Execute()

Get mirror mode



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
	r, err := apiClient.MirroringMirrorAPI.GetMirrorMode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetMirrorMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMirrorModeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMirrorSettings

> RestUpstreamSettings GetMirrorSettings(ctx).Execute()

Get upstream settings



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
	resp, r, err := apiClient.MirroringMirrorAPI.GetMirrorSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetMirrorSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMirrorSettings`: RestUpstreamSettings
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetMirrorSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMirrorSettingsRequest struct via the builder pattern


### Return type

[**RestUpstreamSettings**](RestUpstreamSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMirroredProjects

> GetMirroredProjects(ctx).Execute()

Get mirrored project IDs



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
	r, err := apiClient.MirroringMirrorAPI.GetMirroredProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetMirroredProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMirroredProjectsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMirroredRepository

> RestMirroredRepository GetMirroredRepository(ctx, externalRepositoryId).Execute()

Get clone URLs



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
	externalRepositoryId := "externalRepositoryId_example" // string | the repository ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringMirrorAPI.GetMirroredRepository(context.Background(), externalRepositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetMirroredRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMirroredRepository`: RestMirroredRepository
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetMirroredRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalRepositoryId** | **string** | the repository ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMirroredRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestMirroredRepository**](RestMirroredRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRefChangesQueue

> RestRefSyncQueue GetRefChangesQueue(ctx).Execute()

Get items in ref changes queue



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
	resp, r, err := apiClient.MirroringMirrorAPI.GetRefChangesQueue(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetRefChangesQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRefChangesQueue`: RestRefSyncQueue
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetRefChangesQueue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefChangesQueueRequest struct via the builder pattern


### Return type

[**RestRefSyncQueue**](RestRefSyncQueue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRefChangesQueueCount

> GetRefChangesQueueCount(ctx).Execute()

Get total number of items in ref changes queue



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
	r, err := apiClient.MirroringMirrorAPI.GetRefChangesQueueCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetRefChangesQueueCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefChangesQueueCountRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepoSyncStatus

> GetRepoSyncStatus200Response GetRepoSyncStatus(ctx).Start(start).Limit(limit).Execute()

Get sync status of repositories



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
	resp, r, err := apiClient.MirroringMirrorAPI.GetRepoSyncStatus(context.Background()).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetRepoSyncStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepoSyncStatus`: GetRepoSyncStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetRepoSyncStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepoSyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetRepoSyncStatus200Response**](GetRepoSyncStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepoSyncStatus1

> RestMirrorRepositorySynchronizationStatus GetRepoSyncStatus1(ctx, projectKey, repositorySlug).Execute()

Gets information about the mirrored repository



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
	projectKey := "projectKey_example" // string | The project key
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringMirrorAPI.GetRepoSyncStatus1(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetRepoSyncStatus1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepoSyncStatus1`: RestMirrorRepositorySynchronizationStatus
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetRepoSyncStatus1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepoSyncStatus1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestMirrorRepositorySynchronizationStatus**](RestMirrorRepositorySynchronizationStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryLockOwner

> RestRepositoryLockOwner GetRepositoryLockOwner(ctx, projectKey, repositorySlug).Execute()

Get the repository lock owner for the syncing process



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
	projectKey := "projectKey_example" // string | The project key
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringMirrorAPI.GetRepositoryLockOwner(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetRepositoryLockOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryLockOwner`: RestRepositoryLockOwner
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetRepositoryLockOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryLockOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestRepositoryLockOwner**](RestRepositoryLockOwner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryLockOwners

> []RestRepositoryLockOwner GetRepositoryLockOwners(ctx).Execute()

Get all the repository lock owners for the syncing process



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
	resp, r, err := apiClient.MirroringMirrorAPI.GetRepositoryLockOwners(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetRepositoryLockOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryLockOwners`: []RestRepositoryLockOwner
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetRepositoryLockOwners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryLockOwnersRequest struct via the builder pattern


### Return type

[**[]RestRepositoryLockOwner**](RestRepositoryLockOwner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSynchronizationProgress

> RestSyncProgress GetSynchronizationProgress(ctx).Execute()

Get synchronization progress state



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
	resp, r, err := apiClient.MirroringMirrorAPI.GetSynchronizationProgress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetSynchronizationProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSynchronizationProgress`: RestSyncProgress
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetSynchronizationProgress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSynchronizationProgressRequest struct via the builder pattern


### Return type

[**RestSyncProgress**](RestSyncProgress.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpstreamServer

> RestUpstreamServer GetUpstreamServer(ctx).Execute()

Get upstream server



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
	resp, r, err := apiClient.MirroringMirrorAPI.GetUpstreamServer(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.GetUpstreamServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpstreamServer`: RestUpstreamServer
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.GetUpstreamServer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpstreamServerRequest struct via the builder pattern


### Return type

[**RestUpstreamServer**](RestUpstreamServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMirrorMode

> SetMirrorMode(ctx).Body(body).Execute()

Update mirror mode



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MirroringMirrorAPI.SetMirrorMode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.SetMirrorMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMirrorModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMirrorSettings

> RestUpstreamSettings SetMirrorSettings(ctx).RestUpstreamSettings(restUpstreamSettings).Execute()

Update upstream settings



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
	restUpstreamSettings := *openapiclient.NewRestUpstreamSettings() // RestUpstreamSettings | the mirror settings to update to (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MirroringMirrorAPI.SetMirrorSettings(context.Background()).RestUpstreamSettings(restUpstreamSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.SetMirrorSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMirrorSettings`: RestUpstreamSettings
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.SetMirrorSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMirrorSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restUpstreamSettings** | [**RestUpstreamSettings**](RestUpstreamSettings.md) | the mirror settings to update to | 

### Return type

[**RestUpstreamSettings**](RestUpstreamSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartMirroringProject

> StartMirroringProject(ctx, projectId).Execute()

Add project to be mirrored



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
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MirroringMirrorAPI.StartMirroringProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.StartMirroringProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartMirroringProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartMirroringProjects

> StartMirroringProjects(ctx).RequestBody(requestBody).Execute()

Add multiple projects to be mirrored



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
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MirroringMirrorAPI.StartMirroringProjects(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.StartMirroringProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartMirroringProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartRollingUpgrade

> RestRollingUpgradeState StartRollingUpgrade(ctx).Execute()

Start ZDU upgrade on mirror farm



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
	resp, r, err := apiClient.MirroringMirrorAPI.StartRollingUpgrade(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.StartRollingUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartRollingUpgrade`: RestRollingUpgradeState
	fmt.Fprintf(os.Stdout, "Response from `MirroringMirrorAPI.StartRollingUpgrade`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartRollingUpgradeRequest struct via the builder pattern


### Return type

[**RestRollingUpgradeState**](RestRollingUpgradeState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopMirroringProject

> StopMirroringProject(ctx, projectId).Execute()

Stop mirroring project



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
	projectId := "projectId_example" // string | the project ID to stop mirroring

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MirroringMirrorAPI.StopMirroringProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MirroringMirrorAPI.StopMirroringProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID to stop mirroring | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopMirroringProjectRequest struct via the builder pattern


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


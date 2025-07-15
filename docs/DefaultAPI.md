# \DefaultAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPullRequestCount**](DefaultAPI.md#GetPullRequestCount) | **Get** /api/latest/inbox/pull-requests/count | Get total number of pull requests in inbox
[**GetPullRequests2**](DefaultAPI.md#GetPullRequests2) | **Get** /api/latest/inbox/pull-requests | Get pull requests in inbox
[**ReindexRepositories**](DefaultAPI.md#ReindexRepositories) | **Post** /indexing/latest/reindex | Re-indexes the search index of the provided list of repositories
[**RestartIndexingThreadWorker**](DefaultAPI.md#RestartIndexingThreadWorker) | **Post** /indexing/latest/restart | Restarts the search indexing worker thread



## GetPullRequestCount

> GetPullRequestCount(ctx).Execute()

Get total number of pull requests in inbox



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
	r, err := apiClient.DefaultAPI.GetPullRequestCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPullRequestCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestCountRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequests2

> GetPullRequests2(ctx).Role(role).Limit(limit).Start(start).Execute()

Get pull requests in inbox



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
	role := "role_example" // string |  (optional) (default to "reviewer")
	limit := int32(56) // int32 |  (optional) (default to 25)
	start := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetPullRequests2(context.Background()).Role(role).Limit(limit).Start(start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPullRequests2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequests2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string** |  | [default to &quot;reviewer&quot;]
 **limit** | **int32** |  | [default to 25]
 **start** | **int32** |  | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReindexRepositories

> ReindexRepositories(ctx).RestRepositorySelector(restRepositorySelector).Execute()

Re-indexes the search index of the provided list of repositories



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
	restRepositorySelector := []openapiclient.RestRepositorySelector{*openapiclient.NewRestRepositorySelector()} // []RestRepositorySelector |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ReindexRepositories(context.Background()).RestRepositorySelector(restRepositorySelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ReindexRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReindexRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restRepositorySelector** | [**[]RestRepositorySelector**](RestRepositorySelector.md) |  | 

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


## RestartIndexingThreadWorker

> RestartIndexingThreadWorker(ctx).RestIndexingWorkerRestartRequest(restIndexingWorkerRestartRequest).Execute()

Restarts the search indexing worker thread



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
	restIndexingWorkerRestartRequest := *openapiclient.NewRestIndexingWorkerRestartRequest() // RestIndexingWorkerRestartRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RestartIndexingThreadWorker(context.Background()).RestIndexingWorkerRestartRequest(restIndexingWorkerRestartRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RestartIndexingThreadWorker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestartIndexingThreadWorkerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restIndexingWorkerRestartRequest** | [**RestIndexingWorkerRestartRequest**](RestIndexingWorkerRestartRequest.md) |  | 

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


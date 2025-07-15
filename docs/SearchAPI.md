# \SearchAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrokenIndexStatusRepos**](SearchAPI.md#GetBrokenIndexStatusRepos) | **Get** /indexing/latest/support-info/broken-index-status-repos | Retrieve a paged list of repositories which have exceeded the configured maximum indexing retries.
[**GetDetails**](SearchAPI.md#GetDetails) | **Get** /indexing/latest/projects/{projectKey}/repos/{repositorySlug} | Get repository search indexing details.
[**GetIndexingThreadSnapshot**](SearchAPI.md#GetIndexingThreadSnapshot) | **Get** /indexing/latest/support-info/indexing-thread-snapshot | Retrieve a snapshot of the indexing thread details.
[**GetQueueDetails**](SearchAPI.md#GetQueueDetails) | **Get** /indexing/latest/projects/{projectKey}/repos/{repositorySlug}/indexing-queue-details | Retrieve detailed queue information for a repository
[**IndexingQueuedStatus**](SearchAPI.md#IndexingQueuedStatus) | **Get** /indexing/latest/projects/{projectKey}/repos/{repositorySlug}/indexing-queued-status | Checks if a repository has been queued for indexing.



## GetBrokenIndexStatusRepos

> GetBrokenIndexStatusRepos200Response GetBrokenIndexStatusRepos(ctx).Start(start).Limit(limit).Execute()

Retrieve a paged list of repositories which have exceeded the configured maximum indexing retries.



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
	resp, r, err := apiClient.SearchAPI.GetBrokenIndexStatusRepos(context.Background()).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetBrokenIndexStatusRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrokenIndexStatusRepos`: GetBrokenIndexStatusRepos200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetBrokenIndexStatusRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrokenIndexStatusReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetBrokenIndexStatusRepos200Response**](GetBrokenIndexStatusRepos200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetails

> RestRepositoryIndexingDetails GetDetails(ctx, projectKey, repositorySlug).Execute()

Get repository search indexing details.



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
	resp, r, err := apiClient.SearchAPI.GetDetails(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetails`: RestRepositoryIndexingDetails
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestRepositoryIndexingDetails**](RestRepositoryIndexingDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndexingThreadSnapshot

> []RestIndexingThreadDetails GetIndexingThreadSnapshot(ctx).Execute()

Retrieve a snapshot of the indexing thread details.



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
	resp, r, err := apiClient.SearchAPI.GetIndexingThreadSnapshot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetIndexingThreadSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexingThreadSnapshot`: []RestIndexingThreadDetails
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetIndexingThreadSnapshot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexingThreadSnapshotRequest struct via the builder pattern


### Return type

[**[]RestIndexingThreadDetails**](RestIndexingThreadDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueDetails

> RestRepositoryIndexingQueueDetails GetQueueDetails(ctx, projectKey, repositorySlug).Execute()

Retrieve detailed queue information for a repository



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
	resp, r, err := apiClient.SearchAPI.GetQueueDetails(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetQueueDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueueDetails`: RestRepositoryIndexingQueueDetails
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetQueueDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestRepositoryIndexingQueueDetails**](RestRepositoryIndexingQueueDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndexingQueuedStatus

> RestIndexingIsRepositoryQueued IndexingQueuedStatus(ctx, projectKey, repositorySlug).Execute()

Checks if a repository has been queued for indexing.



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
	resp, r, err := apiClient.SearchAPI.IndexingQueuedStatus(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.IndexingQueuedStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IndexingQueuedStatus`: RestIndexingIsRepositoryQueued
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.IndexingQueuedStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiIndexingQueuedStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestIndexingIsRepositoryQueued**](RestIndexingIsRepositoryQueued.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


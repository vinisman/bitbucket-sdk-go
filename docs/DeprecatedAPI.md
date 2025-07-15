# \DeprecatedAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBuildStatus**](DeprecatedAPI.md#AddBuildStatus) | **Post** /build-status/latest/commits/{commitId} | Create build status for commit
[**AddGroupToUser**](DeprecatedAPI.md#AddGroupToUser) | **Post** /api/latest/admin/users/add-group | Add user to group
[**AddUserToGroup**](DeprecatedAPI.md#AddUserToGroup) | **Post** /api/latest/admin/groups/add-user | Add user to group
[**Approve**](DeprecatedAPI.md#Approve) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/approve | Approve pull request
[**GetBuildStatus**](DeprecatedAPI.md#GetBuildStatus) | **Get** /build-status/latest/commits/{commitId} | Get build statuses for commit
[**GetBuildStatusStats**](DeprecatedAPI.md#GetBuildStatusStats) | **Get** /build-status/latest/commits/stats/{commitId} | Get build status statistics for commit
[**GetDefaultBranch1**](DeprecatedAPI.md#GetDefaultBranch1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/branches/default | Get default branch
[**GetMultipleBuildStatusStats**](DeprecatedAPI.md#GetMultipleBuildStatusStats) | **Post** /build-status/latest/commits/stats | Get build status statistics for multiple commits
[**RemoveUserFromGroup**](DeprecatedAPI.md#RemoveUserFromGroup) | **Post** /api/latest/admin/groups/remove-user | Remove user from group
[**SetDefaultBranch1**](DeprecatedAPI.md#SetDefaultBranch1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/branches/default | Update default branch
[**UnassignParticipantRole1**](DeprecatedAPI.md#UnassignParticipantRole1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants | Unassign pull request participant
[**WithdrawApproval**](DeprecatedAPI.md#WithdrawApproval) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/approve | Unapprove pull request



## AddBuildStatus

> AddBuildStatus(ctx, commitId).RestBuildStatus(restBuildStatus).Execute()

Create build status for commit



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
	commitId := "e00cf62997a027bbf785614a93e2e55bb331d268" // string | full SHA1 of the commit
	restBuildStatus := *openapiclient.NewRestBuildStatus() // RestBuildStatus | build status to associate with the commit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeprecatedAPI.AddBuildStatus(context.Background(), commitId).RestBuildStatus(restBuildStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.AddBuildStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitId** | **string** | full SHA1 of the commit | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBuildStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restBuildStatus** | [**RestBuildStatus**](RestBuildStatus.md) | build status to associate with the commit | 

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


## AddGroupToUser

> AddGroupToUser(ctx).GroupPickerContext(groupPickerContext).Execute()

Add user to group



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
	groupPickerContext := *openapiclient.NewGroupPickerContext() // GroupPickerContext |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeprecatedAPI.AddGroupToUser(context.Background()).GroupPickerContext(groupPickerContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.AddGroupToUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupPickerContext** | [**GroupPickerContext**](GroupPickerContext.md) |  | 

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


## AddUserToGroup

> AddUserToGroup(ctx).UserPickerContext(userPickerContext).Execute()

Add user to group



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
	userPickerContext := *openapiclient.NewUserPickerContext() // UserPickerContext |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeprecatedAPI.AddUserToGroup(context.Background()).UserPickerContext(userPickerContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.AddUserToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPickerContext** | [**UserPickerContext**](UserPickerContext.md) |  | 

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


## Approve

> RestPullRequestParticipant Approve(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Approve pull request



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
	projectKey := "projectKey_example" // string | The project key.
	pullRequestId := "pullRequestId_example" // string | The ID of the pull request within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeprecatedAPI.Approve(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.Approve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Approve`: RestPullRequestParticipant
	fmt.Fprintf(os.Stdout, "Response from `DeprecatedAPI.Approve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The ID of the pull request within the repository | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestPullRequestParticipant**](RestPullRequestParticipant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildStatus

> GetBuildStatus200Response GetBuildStatus(ctx, commitId).OrderBy(orderBy).Start(start).Limit(limit).Execute()

Get build statuses for commit



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
	commitId := "commitId_example" // string | Full SHA1 of the commit (ex: <code>e00cf62997a027bbf785614a93e2e55bb331d268</code>)
	orderBy := "newest, oldest, or status" // string | How the results should be ordered. Options are NEWEST, OLDEST, STATUS (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeprecatedAPI.GetBuildStatus(context.Background(), commitId).OrderBy(orderBy).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.GetBuildStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildStatus`: GetBuildStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `DeprecatedAPI.GetBuildStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitId** | **string** | Full SHA1 of the commit (ex: &lt;code&gt;e00cf62997a027bbf785614a93e2e55bb331d268&lt;/code&gt;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderBy** | **string** | How the results should be ordered. Options are NEWEST, OLDEST, STATUS | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetBuildStatus200Response**](GetBuildStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildStatusStats

> RestBuildStats GetBuildStatusStats(ctx, commitId).IncludeUnique(includeUnique).Execute()

Get build status statistics for commit



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
	commitId := "e00cf62997a027bbf785614a93e2e55bb331d268" // string | full SHA1 of the commit
	includeUnique := true // bool | include a unique build result if there is either only one failed build, only one in-progress build or only one successful build (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeprecatedAPI.GetBuildStatusStats(context.Background(), commitId).IncludeUnique(includeUnique).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.GetBuildStatusStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildStatusStats`: RestBuildStats
	fmt.Fprintf(os.Stdout, "Response from `DeprecatedAPI.GetBuildStatusStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitId** | **string** | full SHA1 of the commit | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildStatusStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeUnique** | **bool** | include a unique build result if there is either only one failed build, only one in-progress build or only one successful build | 

### Return type

[**RestBuildStats**](RestBuildStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultBranch1

> RestBranch GetDefaultBranch1(ctx, projectKey, repositorySlug).Execute()

Get default branch



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
	projectKey := "projectKey_example" // string | The project key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeprecatedAPI.GetDefaultBranch1(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.GetDefaultBranch1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultBranch1`: RestBranch
	fmt.Fprintf(os.Stdout, "Response from `DeprecatedAPI.GetDefaultBranch1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultBranch1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestBranch**](RestBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleBuildStatusStats

> interface{} GetMultipleBuildStatusStats(ctx).RequestBody(requestBody).Execute()

Get build status statistics for multiple commits



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
	requestBody := []string{"db0dd118fa6ccdf1d593fef00df839dd27702df7"} // []string | full SHA1 of each commit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeprecatedAPI.GetMultipleBuildStatusStats(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.GetMultipleBuildStatusStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipleBuildStatusStats`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeprecatedAPI.GetMultipleBuildStatusStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleBuildStatusStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | full SHA1 of each commit | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromGroup

> RemoveUserFromGroup(ctx).UserPickerContext(userPickerContext).Execute()

Remove user from group



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
	userPickerContext := *openapiclient.NewUserPickerContext() // UserPickerContext |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeprecatedAPI.RemoveUserFromGroup(context.Background()).UserPickerContext(userPickerContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.RemoveUserFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPickerContext** | [**UserPickerContext**](UserPickerContext.md) |  | 

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


## SetDefaultBranch1

> SetDefaultBranch1(ctx, projectKey, repositorySlug).RestBranch(restBranch).Execute()

Update default branch



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
	projectKey := "projectKey_example" // string | The project key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restBranch := *openapiclient.NewRestBranch() // RestBranch | The branch to set as default (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeprecatedAPI.SetDefaultBranch1(context.Background(), projectKey, repositorySlug).RestBranch(restBranch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.SetDefaultBranch1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultBranch1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restBranch** | [**RestBranch**](RestBranch.md) | The branch to set as default | 

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


## UnassignParticipantRole1

> UnassignParticipantRole1(ctx, projectKey, pullRequestId, repositorySlug).Username(username).Execute()

Unassign pull request participant



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
	projectKey := "projectKey_example" // string | The project key.
	pullRequestId := "pullRequestId_example" // string | The ID of the pull request within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	username := "username_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeprecatedAPI.UnassignParticipantRole1(context.Background(), projectKey, pullRequestId, repositorySlug).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.UnassignParticipantRole1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The ID of the pull request within the repository | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignParticipantRole1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **username** | **string** |  | 

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


## WithdrawApproval

> RestPullRequestParticipant WithdrawApproval(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Unapprove pull request



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
	projectKey := "projectKey_example" // string | The project key.
	pullRequestId := "pullRequestId_example" // string | The ID of the pull request within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeprecatedAPI.WithdrawApproval(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedAPI.WithdrawApproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WithdrawApproval`: RestPullRequestParticipant
	fmt.Fprintf(os.Stdout, "Response from `DeprecatedAPI.WithdrawApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The ID of the pull request within the repository | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestPullRequestParticipant**](RestPullRequestParticipant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


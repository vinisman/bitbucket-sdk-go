# \PullRequestsAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplySuggestion**](PullRequestsAPI.md#ApplySuggestion) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId}/apply-suggestion | Apply pull request suggestion
[**Approve**](PullRequestsAPI.md#Approve) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/approve | Approve pull request
[**AssignParticipantRole**](PullRequestsAPI.md#AssignParticipantRole) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants | Assign pull request participant role
[**CanMerge**](PullRequestsAPI.md#CanMerge) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/merge | Test if pull request can be merged
[**CanRebase**](PullRequestsAPI.md#CanRebase) | **Get** /git/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/rebase | Check PR rebase precondition
[**CancelAutoMerge**](PullRequestsAPI.md#CancelAutoMerge) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/auto-merge | Cancel auto-merge for pull request
[**Create**](PullRequestsAPI.md#Create) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests | Create pull request
[**Create1**](PullRequestsAPI.md#Create1) | **Post** /api/latest/projects/{projectKey}/settings/reviewer-groups | Create reviewer group
[**Create2**](PullRequestsAPI.md#Create2) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/reviewer-groups | Create reviewer group
[**CreateComment1**](PullRequestsAPI.md#CreateComment1) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/blocker-comments | Add new blocker comment
[**CreateComment2**](PullRequestsAPI.md#CreateComment2) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments | Add pull request comment
[**CreatePullRequestCondition**](PullRequestsAPI.md#CreatePullRequestCondition) | **Post** /default-reviewers/latest/projects/{projectKey}/condition | Create default reviewer condition
[**CreatePullRequestCondition1**](PullRequestsAPI.md#CreatePullRequestCondition1) | **Post** /default-reviewers/latest/projects/{projectKey}/repos/{repositorySlug}/condition | Create default reviewer condition
[**Decline**](PullRequestsAPI.md#Decline) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/decline | Decline pull request
[**Delete3**](PullRequestsAPI.md#Delete3) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | Delete pull request
[**Delete6**](PullRequestsAPI.md#Delete6) | **Delete** /api/latest/projects/{projectKey}/settings/reviewer-groups/{id} | Delete reviewer group
[**Delete7**](PullRequestsAPI.md#Delete7) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/reviewer-groups/{id} | Delete reviewer group
[**DeleteComment1**](PullRequestsAPI.md#DeleteComment1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/blocker-comments/{commentId} | Delete pull request comment
[**DeleteComment2**](PullRequestsAPI.md#DeleteComment2) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId} | Delete a pull request comment
[**DeletePullRequestCondition**](PullRequestsAPI.md#DeletePullRequestCondition) | **Delete** /default-reviewers/latest/projects/{projectKey}/condition/{id} | Delete default reviewer condition
[**DeletePullRequestCondition1**](PullRequestsAPI.md#DeletePullRequestCondition1) | **Delete** /default-reviewers/latest/projects/{projectKey}/repos/{repositorySlug}/condition/{id} | Delete default reviewer condition
[**DiscardReview**](PullRequestsAPI.md#DiscardReview) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/review | Discard pull request review
[**FinishReview**](PullRequestsAPI.md#FinishReview) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/review | Complete pull request review
[**Get3**](PullRequestsAPI.md#Get3) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | Get pull request
[**GetActivities**](PullRequestsAPI.md#GetActivities) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/activities | Get pull request activity
[**GetAutoMergeRequest**](PullRequestsAPI.md#GetAutoMergeRequest) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/auto-merge | Get auto-merge request for pull request
[**GetComment1**](PullRequestsAPI.md#GetComment1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/blocker-comments/{commentId} | Get pull request comment
[**GetComment2**](PullRequestsAPI.md#GetComment2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId} | Get a pull request comment
[**GetComments1**](PullRequestsAPI.md#GetComments1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/blocker-comments | Search pull request comments
[**GetComments2**](PullRequestsAPI.md#GetComments2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments | Get pull request comments for path
[**GetCommitMessageSuggestion**](PullRequestsAPI.md#GetCommitMessageSuggestion) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/commit-message-suggestion | Get commit message suggestion
[**GetCommits1**](PullRequestsAPI.md#GetCommits1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/commits | Get pull request commits
[**GetDiffStatsSummary2**](PullRequestsAPI.md#GetDiffStatsSummary2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/diff-stats-summary/{path} | Get diff stats summary for pull request
[**GetMergeBase1**](PullRequestsAPI.md#GetMergeBase1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/merge-base | Get the common ancestor between the latest commits of the source and target branches of the pull request
[**GetMergeConfig**](PullRequestsAPI.md#GetMergeConfig) | **Get** /api/latest/admin/pull-requests/{scmId} | Get merge strategies
[**GetPage**](PullRequestsAPI.md#GetPage) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests | Get pull requests for repository
[**GetPullRequestConditions**](PullRequestsAPI.md#GetPullRequestConditions) | **Get** /default-reviewers/latest/projects/{projectKey}/conditions | Get default reviewer conditions
[**GetPullRequestConditions1**](PullRequestsAPI.md#GetPullRequestConditions1) | **Get** /default-reviewers/latest/projects/{projectKey}/repos/{repositorySlug}/conditions | Get default reviewer conditions
[**GetPullRequests**](PullRequestsAPI.md#GetPullRequests) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/pull-requests | Get repository pull requests containing commit
[**GetReview**](PullRequestsAPI.md#GetReview) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/review | Get pull request comment thread
[**GetReviewerGroup**](PullRequestsAPI.md#GetReviewerGroup) | **Get** /api/latest/projects/{projectKey}/settings/reviewer-groups/{id} | Get reviewer group
[**GetReviewerGroup1**](PullRequestsAPI.md#GetReviewerGroup1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/reviewer-groups/{id} | Get reviewer group
[**GetReviewerGroups**](PullRequestsAPI.md#GetReviewerGroups) | **Get** /api/latest/projects/{projectKey}/settings/reviewer-groups | Get all reviewer groups
[**GetReviewerGroups1**](PullRequestsAPI.md#GetReviewerGroups1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/reviewer-groups | Get all reviewer groups
[**GetReviewers**](PullRequestsAPI.md#GetReviewers) | **Get** /default-reviewers/latest/projects/{projectKey}/repos/{repositorySlug}/reviewers | Get required reviewers for PR creation
[**GetUsers**](PullRequestsAPI.md#GetUsers) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/reviewer-groups/{id}/users | Get reviewer group users
[**ListParticipants**](PullRequestsAPI.md#ListParticipants) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants | Get pull request participants
[**Merge**](PullRequestsAPI.md#Merge) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/merge | Merge pull request
[**React1**](PullRequestsAPI.md#React1) | **Put** /comment-likes/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId}/reactions/{emoticon} | React to a PR comment
[**Rebase**](PullRequestsAPI.md#Rebase) | **Post** /git/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/rebase | Rebase pull request
[**Reopen**](PullRequestsAPI.md#Reopen) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/reopen | Re-open pull request
[**Search**](PullRequestsAPI.md#Search) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/participants | Search pull request participants
[**SetMergeConfig**](PullRequestsAPI.md#SetMergeConfig) | **Post** /api/latest/admin/pull-requests/{scmId} | Update merge strategies
[**StreamChanges1**](PullRequestsAPI.md#StreamChanges1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/changes | Gets pull request changes
[**StreamDiff2**](PullRequestsAPI.md#StreamDiff2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/diff/{path} | Stream a diff within a pull request
[**StreamPatch1**](PullRequestsAPI.md#StreamPatch1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}.patch | Stream pull request as patch
[**StreamRawDiff2**](PullRequestsAPI.md#StreamRawDiff2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}.diff | Stream raw pull request diff
[**TryAutoMerge**](PullRequestsAPI.md#TryAutoMerge) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/auto-merge | Auto-merge pull request
[**UnReact1**](PullRequestsAPI.md#UnReact1) | **Delete** /comment-likes/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId}/reactions/{emoticon} | Remove a reaction from a PR comment
[**UnassignParticipantRole**](PullRequestsAPI.md#UnassignParticipantRole) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug} | Unassign pull request participant
[**UnassignParticipantRole1**](PullRequestsAPI.md#UnassignParticipantRole1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants | Unassign pull request participant
[**Unwatch1**](PullRequestsAPI.md#Unwatch1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/watch | Stop watching pull request
[**Update**](PullRequestsAPI.md#Update) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId} | Update pull request metadata
[**Update1**](PullRequestsAPI.md#Update1) | **Put** /api/latest/projects/{projectKey}/settings/reviewer-groups/{id} | Update reviewer group attributes
[**Update2**](PullRequestsAPI.md#Update2) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/reviewer-groups/{id} | Update reviewer group attributes
[**UpdateComment1**](PullRequestsAPI.md#UpdateComment1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/blocker-comments/{commentId} | Update pull request comment
[**UpdateComment2**](PullRequestsAPI.md#UpdateComment2) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/comments/{commentId} | Update pull request comment
[**UpdatePullRequestCondition**](PullRequestsAPI.md#UpdatePullRequestCondition) | **Put** /default-reviewers/latest/projects/{projectKey}/condition/{id} | Update default reviewer condition
[**UpdatePullRequestCondition1**](PullRequestsAPI.md#UpdatePullRequestCondition1) | **Put** /default-reviewers/latest/projects/{projectKey}/repos/{repositorySlug}/condition/{id} | Update default reviewer condition
[**UpdateStatus**](PullRequestsAPI.md#UpdateStatus) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/participants/{userSlug} | Change pull request status
[**Watch1**](PullRequestsAPI.md#Watch1) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/watch | Watch pull request
[**WithdrawApproval**](PullRequestsAPI.md#WithdrawApproval) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/approve | Unapprove pull request



## ApplySuggestion

> ApplySuggestion(ctx, projectKey, commentId, pullRequestId, repositorySlug).RestApplySuggestionRequest(restApplySuggestionRequest).Execute()

Apply pull request suggestion



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
	commentId := "commentId_example" // string | The ID of the comment to retrieve.
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restApplySuggestionRequest := *openapiclient.NewRestApplySuggestionRequest() // RestApplySuggestionRequest | A request containing other parameters required to apply a suggestion - The given versions/hashes must match the server's version/hashes or the suggestion application will fail (in order to avoid applying the suggestion to the wrong place (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.ApplySuggestion(context.Background(), projectKey, commentId, pullRequestId, repositorySlug).RestApplySuggestionRequest(restApplySuggestionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.ApplySuggestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The ID of the comment to retrieve. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplySuggestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **restApplySuggestionRequest** | [**RestApplySuggestionRequest**](RestApplySuggestionRequest.md) | A request containing other parameters required to apply a suggestion - The given versions/hashes must match the server&#39;s version/hashes or the suggestion application will fail (in order to avoid applying the suggestion to the wrong place | 

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
	resp, r, err := apiClient.PullRequestsAPI.Approve(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Approve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Approve`: RestPullRequestParticipant
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Approve`: %v\n", resp)
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


## AssignParticipantRole

> RestPullRequestParticipant AssignParticipantRole(ctx, projectKey, pullRequestId, repositorySlug).RestPullRequestAssignParticipantRoleRequest(restPullRequestAssignParticipantRoleRequest).Execute()

Assign pull request participant role



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
	restPullRequestAssignParticipantRoleRequest := *openapiclient.NewRestPullRequestAssignParticipantRoleRequest() // RestPullRequestAssignParticipantRoleRequest | The participant to be added to the pull request, includes the user and their role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.AssignParticipantRole(context.Background(), projectKey, pullRequestId, repositorySlug).RestPullRequestAssignParticipantRoleRequest(restPullRequestAssignParticipantRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.AssignParticipantRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignParticipantRole`: RestPullRequestParticipant
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.AssignParticipantRole`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAssignParticipantRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restPullRequestAssignParticipantRoleRequest** | [**RestPullRequestAssignParticipantRoleRequest**](RestPullRequestAssignParticipantRoleRequest.md) | The participant to be added to the pull request, includes the user and their role | 

### Return type

[**RestPullRequestParticipant**](RestPullRequestParticipant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CanMerge

> RestPullRequestMergeability CanMerge(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Test if pull request can be merged



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
	resp, r, err := apiClient.PullRequestsAPI.CanMerge(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.CanMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CanMerge`: RestPullRequestMergeability
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.CanMerge`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCanMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestPullRequestMergeability**](RestPullRequestMergeability.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CanRebase

> RestPullRequestRebaseability CanRebase(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Check PR rebase precondition



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
	pullRequestId := "pullRequestId_example" // string | The ID of the pull request within the repository.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.CanRebase(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.CanRebase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CanRebase`: RestPullRequestRebaseability
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.CanRebase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The ID of the pull request within the repository. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCanRebaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestPullRequestRebaseability**](RestPullRequestRebaseability.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelAutoMerge

> CancelAutoMerge(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Cancel auto-merge for pull request



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
	r, err := apiClient.PullRequestsAPI.CancelAutoMerge(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.CancelAutoMerge``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCancelAutoMergeRequest struct via the builder pattern


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


## Create

> RestPullRequest Create(ctx, projectKey, repositorySlug).RestPullRequest(restPullRequest).Execute()

Create pull request



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
	restPullRequest := *openapiclient.NewRestPullRequest() // RestPullRequest | The pull request data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Create(context.Background(), projectKey, repositorySlug).RestPullRequest(restPullRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: RestPullRequest
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restPullRequest** | [**RestPullRequest**](RestPullRequest.md) | The pull request data | 

### Return type

[**RestPullRequest**](RestPullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create1

> RestReviewerGroup Create1(ctx, projectKey).RestReviewerGroup(restReviewerGroup).Execute()

Create reviewer group



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
	restReviewerGroup := *openapiclient.NewRestReviewerGroup() // RestReviewerGroup | The reviewer group to be create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Create1(context.Background(), projectKey).RestReviewerGroup(restReviewerGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Create1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create1`: RestReviewerGroup
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Create1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restReviewerGroup** | [**RestReviewerGroup**](RestReviewerGroup.md) | The reviewer group to be create | 

### Return type

[**RestReviewerGroup**](RestReviewerGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create2

> RestReviewerGroup Create2(ctx, projectKey, repositorySlug).RestReviewerGroup(restReviewerGroup).Execute()

Create reviewer group



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
	restReviewerGroup := *openapiclient.NewRestReviewerGroup() // RestReviewerGroup | The request containing the details of the reviewer group. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Create2(context.Background(), projectKey, repositorySlug).RestReviewerGroup(restReviewerGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Create2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create2`: RestReviewerGroup
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Create2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restReviewerGroup** | [**RestReviewerGroup**](RestReviewerGroup.md) | The request containing the details of the reviewer group. | 

### Return type

[**RestReviewerGroup**](RestReviewerGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComment1

> RestComment CreateComment1(ctx, projectKey, pullRequestId, repositorySlug).RestComment(restComment).Execute()

Add new blocker comment



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restComment := *openapiclient.NewRestComment() // RestComment | The comment to add. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.CreateComment1(context.Background(), projectKey, pullRequestId, repositorySlug).RestComment(restComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.CreateComment1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComment1`: RestComment
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.CreateComment1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateComment1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restComment** | [**RestComment**](RestComment.md) | The comment to add. | 

### Return type

[**RestComment**](RestComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComment2

> RestComment CreateComment2(ctx, projectKey, pullRequestId, repositorySlug).RestComment(restComment).Execute()

Add pull request comment



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restComment := *openapiclient.NewRestComment() // RestComment | The comment to add (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.CreateComment2(context.Background(), projectKey, pullRequestId, repositorySlug).RestComment(restComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.CreateComment2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComment2`: RestComment
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.CreateComment2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateComment2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restComment** | [**RestComment**](RestComment.md) | The comment to add | 

### Return type

[**RestComment**](RestComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePullRequestCondition

> RestPullRequestCondition CreatePullRequestCondition(ctx, projectKey).RestDefaultReviewersRequest(restDefaultReviewersRequest).Execute()

Create default reviewer condition



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
	restDefaultReviewersRequest := *openapiclient.NewRestDefaultReviewersRequest() // RestDefaultReviewersRequest | The details needed to create a default reviewer pull request condition. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.CreatePullRequestCondition(context.Background(), projectKey).RestDefaultReviewersRequest(restDefaultReviewersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.CreatePullRequestCondition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePullRequestCondition`: RestPullRequestCondition
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.CreatePullRequestCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePullRequestConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restDefaultReviewersRequest** | [**RestDefaultReviewersRequest**](RestDefaultReviewersRequest.md) | The details needed to create a default reviewer pull request condition. | 

### Return type

[**RestPullRequestCondition**](RestPullRequestCondition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePullRequestCondition1

> RestPullRequestCondition CreatePullRequestCondition1(ctx, projectKey, repositorySlug).RestDefaultReviewersRequest(restDefaultReviewersRequest).Execute()

Create default reviewer condition



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
	restDefaultReviewersRequest := *openapiclient.NewRestDefaultReviewersRequest() // RestDefaultReviewersRequest | The details needed to create a default reviewer pull request condition. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.CreatePullRequestCondition1(context.Background(), projectKey, repositorySlug).RestDefaultReviewersRequest(restDefaultReviewersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.CreatePullRequestCondition1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePullRequestCondition1`: RestPullRequestCondition
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.CreatePullRequestCondition1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePullRequestCondition1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restDefaultReviewersRequest** | [**RestDefaultReviewersRequest**](RestDefaultReviewersRequest.md) | The details needed to create a default reviewer pull request condition. | 

### Return type

[**RestPullRequestCondition**](RestPullRequestCondition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Decline

> RestPullRequest Decline(ctx, projectKey, pullRequestId, repositorySlug).Version(version).RestPullRequestDeclineRequest(restPullRequestDeclineRequest).Execute()

Decline pull request



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
	pullRequestId := "pullRequestId_example" // string | The pullrequest ID provided by the path
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	version := "version_example" // string | The current version of the pull request. If the server's version isn't the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the 'version' attribute in the returned JSON structure. (optional)
	restPullRequestDeclineRequest := *openapiclient.NewRestPullRequestDeclineRequest() // RestPullRequestDeclineRequest | Optional body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Decline(context.Background(), projectKey, pullRequestId, repositorySlug).Version(version).RestPullRequestDeclineRequest(restPullRequestDeclineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Decline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Decline`: RestPullRequest
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Decline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pullrequest ID provided by the path | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **version** | **string** | The current version of the pull request. If the server&#39;s version isn&#39;t the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the &#39;version&#39; attribute in the returned JSON structure. | 
 **restPullRequestDeclineRequest** | [**RestPullRequestDeclineRequest**](RestPullRequestDeclineRequest.md) | Optional body | 

### Return type

[**RestPullRequest**](RestPullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete3

> Delete3(ctx, projectKey, pullRequestId, repositorySlug).RestPullRequestDeleteRequest(restPullRequestDeleteRequest).Execute()

Delete pull request



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
	restPullRequestDeleteRequest := *openapiclient.NewRestPullRequestDeleteRequest() // RestPullRequestDeleteRequest | A body containing the version of the pull request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.Delete3(context.Background(), projectKey, pullRequestId, repositorySlug).RestPullRequestDeleteRequest(restPullRequestDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Delete3``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDelete3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restPullRequestDeleteRequest** | [**RestPullRequestDeleteRequest**](RestPullRequestDeleteRequest.md) | A body containing the version of the pull request | 

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


## Delete6

> Delete6(ctx, projectKey, id).Execute()

Delete reviewer group



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
	id := "id_example" // string | The ID of the reviewer group to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.Delete6(context.Background(), projectKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Delete6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the reviewer group to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete6Request struct via the builder pattern


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


## Delete7

> Delete7(ctx, projectKey, id, repositorySlug).Execute()

Delete reviewer group



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
	id := "id_example" // string | The ID of the reviewer group to be deleted
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.Delete7(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Delete7``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the reviewer group to be deleted | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete7Request struct via the builder pattern


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


## DeleteComment1

> DeleteComment1(ctx, projectKey, commentId, pullRequestId, repositorySlug).Version(version).Execute()

Delete pull request comment



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
	commentId := "commentId_example" // string | The ID of the comment to retrieve.
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	version := "version_example" // string | The expected version of the comment. This must match the server's version of the comment or the delete will fail. To determine the current version of the comment, the comment should be fetched from the server prior to the delete. Look for the 'version' attribute in the returned JSON structure. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.DeleteComment1(context.Background(), projectKey, commentId, pullRequestId, repositorySlug).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.DeleteComment1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The ID of the comment to retrieve. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComment1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **version** | **string** | The expected version of the comment. This must match the server&#39;s version of the comment or the delete will fail. To determine the current version of the comment, the comment should be fetched from the server prior to the delete. Look for the &#39;version&#39; attribute in the returned JSON structure. | 

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


## DeleteComment2

> DeleteComment2(ctx, projectKey, commentId, pullRequestId, repositorySlug).Version(version).Execute()

Delete a pull request comment



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
	commentId := "commentId_example" // string | The ID of the comment to retrieve.
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	version := "version_example" // string | The expected version of the comment. This must match the server's version of the comment or the delete will fail. To determine the current version of the comment, the comment should be fetched from the server prior to the delete. Look for the 'version' attribute in the returned JSON structure. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.DeleteComment2(context.Background(), projectKey, commentId, pullRequestId, repositorySlug).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.DeleteComment2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The ID of the comment to retrieve. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComment2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **version** | **string** | The expected version of the comment. This must match the server&#39;s version of the comment or the delete will fail. To determine the current version of the comment, the comment should be fetched from the server prior to the delete. Look for the &#39;version&#39; attribute in the returned JSON structure. | 

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


## DeletePullRequestCondition

> DeletePullRequestCondition(ctx, projectKey, id).Execute()

Delete default reviewer condition



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
	id := "id_example" // string | The ID of the pull request condition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.DeletePullRequestCondition(context.Background(), projectKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.DeletePullRequestCondition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the pull request condition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePullRequestConditionRequest struct via the builder pattern


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


## DeletePullRequestCondition1

> DeletePullRequestCondition1(ctx, projectKey, id, repositorySlug).Execute()

Delete default reviewer condition



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
	id := int32(56) // int32 | 
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.DeletePullRequestCondition1(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.DeletePullRequestCondition1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **int32** |  | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePullRequestCondition1Request struct via the builder pattern


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


## DiscardReview

> DiscardReview(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Discard pull request review



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.DiscardReview(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.DiscardReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscardReviewRequest struct via the builder pattern


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


## FinishReview

> FinishReview(ctx, projectKey, pullRequestId, repositorySlug).Version(version).RestPullRequestFinishReviewRequest(restPullRequestFinishReviewRequest).Execute()

Complete pull request review



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	version := "version_example" // string | The current version of the pull request. If the server's version isn't the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the 'version' attribute in the returned JSON structure. Note: This parameter is deprecated. Use last reviewed commit in request body instead (optional)
	restPullRequestFinishReviewRequest := *openapiclient.NewRestPullRequestFinishReviewRequest() // RestPullRequestFinishReviewRequest | The REST request which contains comment text, last reviewed commit and participant status. If last reviewed commit is provided, it will be used to update the participant status. The operation will fail if the latest commit of the pull request does not match the provided last reviewed commit. If last reviewed commit is not provided, the latest commit of the pull request will be used for the update by default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.FinishReview(context.Background(), projectKey, pullRequestId, repositorySlug).Version(version).RestPullRequestFinishReviewRequest(restPullRequestFinishReviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.FinishReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinishReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **version** | **string** | The current version of the pull request. If the server&#39;s version isn&#39;t the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the &#39;version&#39; attribute in the returned JSON structure. Note: This parameter is deprecated. Use last reviewed commit in request body instead | 
 **restPullRequestFinishReviewRequest** | [**RestPullRequestFinishReviewRequest**](RestPullRequestFinishReviewRequest.md) | The REST request which contains comment text, last reviewed commit and participant status. If last reviewed commit is provided, it will be used to update the participant status. The operation will fail if the latest commit of the pull request does not match the provided last reviewed commit. If last reviewed commit is not provided, the latest commit of the pull request will be used for the update by default. | 

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


## Get3

> RestPullRequest Get3(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Get pull request



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
	resp, r, err := apiClient.PullRequestsAPI.Get3(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Get3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get3`: RestPullRequest
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Get3`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGet3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestPullRequest**](RestPullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivities

> GetActivities200Response GetActivities(ctx, projectKey, pullRequestId, repositorySlug).FromType(fromType).FromId(fromId).Start(start).Limit(limit).Execute()

Get pull request activity



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
	fromType := "fromType_example" // string | (required if <strong>fromId</strong> is present) the type of the activity item specified by <strong>fromId</strong> (either <strong>COMMENT</strong> or <strong>ACTIVITY</strong>) (optional)
	fromId := "fromId_example" // string | (optional) the ID of the activity item to use as the first item in the returned page (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetActivities(context.Background(), projectKey, pullRequestId, repositorySlug).FromType(fromType).FromId(fromId).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivities`: GetActivities200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetActivities`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromType** | **string** | (required if &lt;strong&gt;fromId&lt;/strong&gt; is present) the type of the activity item specified by &lt;strong&gt;fromId&lt;/strong&gt; (either &lt;strong&gt;COMMENT&lt;/strong&gt; or &lt;strong&gt;ACTIVITY&lt;/strong&gt;) | 
 **fromId** | **string** | (optional) the ID of the activity item to use as the first item in the returned page | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetActivities200Response**](GetActivities200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoMergeRequest

> RestAutoMergeRequest GetAutoMergeRequest(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Get auto-merge request for pull request



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
	resp, r, err := apiClient.PullRequestsAPI.GetAutoMergeRequest(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetAutoMergeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoMergeRequest`: RestAutoMergeRequest
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetAutoMergeRequest`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAutoMergeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestAutoMergeRequest**](RestAutoMergeRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComment1

> RestComment GetComment1(ctx, projectKey, commentId, pullRequestId, repositorySlug).Execute()

Get pull request comment



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
	commentId := "commentId_example" // string | The ID of the comment to retrieve
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetComment1(context.Background(), projectKey, commentId, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetComment1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComment1`: RestComment
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetComment1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The ID of the comment to retrieve | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComment1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**RestComment**](RestComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComment2

> RestComment GetComment2(ctx, projectKey, commentId, pullRequestId, repositorySlug).Execute()

Get a pull request comment



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
	commentId := "commentId_example" // string | The ID of the comment to retrieve.
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetComment2(context.Background(), projectKey, commentId, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetComment2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComment2`: RestComment
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetComment2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The ID of the comment to retrieve. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComment2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**RestComment**](RestComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComments1

> GetComments200Response GetComments1(ctx, projectKey, pullRequestId, repositorySlug).Count(count).State(state).States(states).Start(start).Limit(limit).Execute()

Search pull request comments



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	count := "count_example" // string | If true only the count of the comments by state will be returned (and not the body of the comments). (optional)
	state := []string{"Inner_example"} // []string |  (optional)
	states := "states_example" // string | (optional). If supplied, only comments with a state in the given list will be returned. The state can be OPEN or RESOLVED. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetComments1(context.Background(), projectKey, pullRequestId, repositorySlug).Count(count).State(state).States(states).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetComments1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComments1`: GetComments200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetComments1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComments1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **string** | If true only the count of the comments by state will be returned (and not the body of the comments). | 
 **state** | **[]string** |  | 
 **states** | **string** | (optional). If supplied, only comments with a state in the given list will be returned. The state can be OPEN or RESOLVED. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetComments200Response**](GetComments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComments2

> GetComments200Response GetComments2(ctx, projectKey, pullRequestId, repositorySlug).Path(path).FromHash(fromHash).AnchorState(anchorState).DiffType(diffType).ToHash(toHash).State(state).DiffTypes(diffTypes).States(states).Start(start).Limit(limit).Execute()

Get pull request comments for path



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
	path := "path_example" // string | The path to stream comments for a given path
	projectKey := "projectKey_example" // string | The project key.
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	fromHash := "fromHash_example" // string | The from commit hash to stream comments for a RANGE or COMMIT arbitrary change scope (optional)
	anchorState := "anchorState_example" // string | ACTIVE to stream the active comments; ORPHANED to stream the orphaned comments; ALL to stream both the active and the orphaned comments; (optional)
	diffType := []string{"Inner_example"} // []string |  (optional)
	toHash := "toHash_example" // string | The to commit hash to stream comments for a RANGE or COMMIT arbitrary change scope (optional)
	state := []string{"Inner_example"} // []string |  (optional)
	diffTypes := "diffTypes_example" // string | EFFECTIVE to stream the comments related to the effective diff of the pull request; RANGE to stream comments related to a commit range between two arbitrary commits (requires 'fromHash' and 'toHash'); COMMIT to stream comments related to a commit between two arbitrary commits (requires 'fromHash' and 'toHash') (optional)
	states := "states_example" // string | (optional). If supplied, only comments with a state in the given list will be returned. The state can be OPEN or RESOLVED. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetComments2(context.Background(), projectKey, pullRequestId, repositorySlug).Path(path).FromHash(fromHash).AnchorState(anchorState).DiffType(diffType).ToHash(toHash).State(state).DiffTypes(diffTypes).States(states).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetComments2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComments2`: GetComments200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetComments2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComments2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | The path to stream comments for a given path | 



 **fromHash** | **string** | The from commit hash to stream comments for a RANGE or COMMIT arbitrary change scope | 
 **anchorState** | **string** | ACTIVE to stream the active comments; ORPHANED to stream the orphaned comments; ALL to stream both the active and the orphaned comments; | 
 **diffType** | **[]string** |  | 
 **toHash** | **string** | The to commit hash to stream comments for a RANGE or COMMIT arbitrary change scope | 
 **state** | **[]string** |  | 
 **diffTypes** | **string** | EFFECTIVE to stream the comments related to the effective diff of the pull request; RANGE to stream comments related to a commit range between two arbitrary commits (requires &#39;fromHash&#39; and &#39;toHash&#39;); COMMIT to stream comments related to a commit between two arbitrary commits (requires &#39;fromHash&#39; and &#39;toHash&#39;) | 
 **states** | **string** | (optional). If supplied, only comments with a state in the given list will be returned. The state can be OPEN or RESOLVED. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetComments200Response**](GetComments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommitMessageSuggestion

> RestCommitMessageSuggestion GetCommitMessageSuggestion(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Get commit message suggestion



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
	pullRequestId := "pullRequestId_example" // string | The ID of the pull request to generate the suggestion for
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetCommitMessageSuggestion(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetCommitMessageSuggestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommitMessageSuggestion`: RestCommitMessageSuggestion
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetCommitMessageSuggestion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The ID of the pull request to generate the suggestion for | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitMessageSuggestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestCommitMessageSuggestion**](RestCommitMessageSuggestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommits1

> GetCommits200Response GetCommits1(ctx, projectKey, pullRequestId, repositorySlug).AvatarScheme(avatarScheme).WithCounts(withCounts).AvatarSize(avatarSize).Start(start).Limit(limit).Execute()

Get pull request commits



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
	pullRequestId := "pullRequestId_example" // string | ID of the pullrequest, part of the path
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	avatarScheme := "avatarScheme_example" // string | The desired scheme for the avatar URL. If the parameter is not present URLs will use the same scheme as this request (optional)
	withCounts := "withCounts_example" // string | If set to true, the service will add \"authorCount\" and \"totalCount\" at the end of the page. \"authorCount\" is the number of different authors and \"totalCount\" is the total number of commits. (optional)
	avatarSize := "avatarSize_example" // string | If present the service adds avatar URLs for commit authors. Should be an integer specifying the desired size in pixels. If the parameter is not present, avatar URLs will not be setCOMMIT to stream comments related to a commit between two arbitrary commits (requires 'fromHash' and 'toHash') (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetCommits1(context.Background(), projectKey, pullRequestId, repositorySlug).AvatarScheme(avatarScheme).WithCounts(withCounts).AvatarSize(avatarSize).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetCommits1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommits1`: GetCommits200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetCommits1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | ID of the pullrequest, part of the path | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommits1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **avatarScheme** | **string** | The desired scheme for the avatar URL. If the parameter is not present URLs will use the same scheme as this request | 
 **withCounts** | **string** | If set to true, the service will add \&quot;authorCount\&quot; and \&quot;totalCount\&quot; at the end of the page. \&quot;authorCount\&quot; is the number of different authors and \&quot;totalCount\&quot; is the total number of commits. | 
 **avatarSize** | **string** | If present the service adds avatar URLs for commit authors. Should be an integer specifying the desired size in pixels. If the parameter is not present, avatar URLs will not be setCOMMIT to stream comments related to a commit between two arbitrary commits (requires &#39;fromHash&#39; and &#39;toHash&#39;) | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetCommits200Response**](GetCommits200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiffStatsSummary2

> interface{} GetDiffStatsSummary2(ctx, path, projectKey, pullRequestId, repositorySlug).SinceId(sinceId).SrcPath(srcPath).UntilId(untilId).Whitespace(whitespace).Execute()

Get diff stats summary for pull request



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
	path := "path_example" // string | Optional path to the file which should be diffed
	projectKey := "projectKey_example" // string | The project key.
	pullRequestId := "pullRequestId_example" // string | The ID of the pull request within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	sinceId := "sinceId_example" // string | The since commit hash to stream a diff between two arbitrary hashes (optional)
	srcPath := "srcPath_example" // string | The previous path to the file, if the file has been copied, moved or renamed (optional)
	untilId := "untilId_example" // string | The until commit hash to stream a diff between two arbitrary hashes (optional)
	whitespace := "whitespace_example" // string | Optional whitespace flag which can be set to <code>ignore-all</code> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetDiffStatsSummary2(context.Background(), path, projectKey, pullRequestId, repositorySlug).SinceId(sinceId).SrcPath(srcPath).UntilId(untilId).Whitespace(whitespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetDiffStatsSummary2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiffStatsSummary2`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetDiffStatsSummary2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | Optional path to the file which should be diffed | 
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The ID of the pull request within the repository | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiffStatsSummary2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sinceId** | **string** | The since commit hash to stream a diff between two arbitrary hashes | 
 **srcPath** | **string** | The previous path to the file, if the file has been copied, moved or renamed | 
 **untilId** | **string** | The until commit hash to stream a diff between two arbitrary hashes | 
 **whitespace** | **string** | Optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergeBase1

> RestCommit GetMergeBase1(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Get the common ancestor between the latest commits of the source and target branches of the pull request



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetMergeBase1(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetMergeBase1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMergeBase1`: RestCommit
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetMergeBase1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergeBase1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestCommit**](RestCommit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergeConfig

> RestPullRequestMergeConfig GetMergeConfig(ctx, scmId).Execute()

Get merge strategies



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
	scmId := "scmId_example" // string | the id of the scm to get strategies for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetMergeConfig(context.Background(), scmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetMergeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMergeConfig`: RestPullRequestMergeConfig
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetMergeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scmId** | **string** | the id of the scm to get strategies for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestPullRequestMergeConfig**](RestPullRequestMergeConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPage

> GetPullRequests1200Response GetPage(ctx, projectKey, repositorySlug).WithAttributes(withAttributes).At(at).WithProperties(withProperties).Draft(draft).FilterText(filterText).State(state).Order(order).Direction(direction).Start(start).Limit(limit).Execute()

Get pull requests for repository



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
	withAttributes := "withAttributes_example" // string | (optional) defaults to true, whether to return additional pull request attributes (optional)
	at := "at_example" // string | (optional) a <i>fully-qualified</i> branch ID to find pull requests to or from, such as refs/heads/master (optional)
	withProperties := "withProperties_example" // string | (optional) defaults to true, whether to return additional pull request properties (optional)
	draft := "draft_example" // string | (optional) If specified, only pull requests matching the supplied draft status will be returned. (optional)
	filterText := "filterText_example" // string | (optional) If specified, only pull requests where the title or description contains the supplied string will be returned. (optional)
	state := "state_example" // string | (optional, defaults to <strong>OPEN</strong>). Supply <strong>ALL</strong> to return pull request in any state. If a state is supplied only pull requests in the specified state will be returned. Either <strong>OPEN</strong>, <strong>DECLINED</strong> or <strong>MERGED</strong>. (optional)
	order := "order_example" // string | (optional, defaults to <strong>NEWEST</strong>) the order to return pull requests in, either <strong>OLDEST</strong> (as in: \"oldest first\") or <strong>NEWEST</strong>. (optional)
	direction := "direction_example" // string | (optional, defaults to <strong>INCOMING</strong>) the direction relative to the specified repository. Either <strong>INCOMING</strong> or <strong>OUTGOING</strong>. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetPage(context.Background(), projectKey, repositorySlug).WithAttributes(withAttributes).At(at).WithProperties(withProperties).Draft(draft).FilterText(filterText).State(state).Order(order).Direction(direction).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPage`: GetPullRequests1200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withAttributes** | **string** | (optional) defaults to true, whether to return additional pull request attributes | 
 **at** | **string** | (optional) a &lt;i&gt;fully-qualified&lt;/i&gt; branch ID to find pull requests to or from, such as refs/heads/master | 
 **withProperties** | **string** | (optional) defaults to true, whether to return additional pull request properties | 
 **draft** | **string** | (optional) If specified, only pull requests matching the supplied draft status will be returned. | 
 **filterText** | **string** | (optional) If specified, only pull requests where the title or description contains the supplied string will be returned. | 
 **state** | **string** | (optional, defaults to &lt;strong&gt;OPEN&lt;/strong&gt;). Supply &lt;strong&gt;ALL&lt;/strong&gt; to return pull request in any state. If a state is supplied only pull requests in the specified state will be returned. Either &lt;strong&gt;OPEN&lt;/strong&gt;, &lt;strong&gt;DECLINED&lt;/strong&gt; or &lt;strong&gt;MERGED&lt;/strong&gt;. | 
 **order** | **string** | (optional, defaults to &lt;strong&gt;NEWEST&lt;/strong&gt;) the order to return pull requests in, either &lt;strong&gt;OLDEST&lt;/strong&gt; (as in: \&quot;oldest first\&quot;) or &lt;strong&gt;NEWEST&lt;/strong&gt;. | 
 **direction** | **string** | (optional, defaults to &lt;strong&gt;INCOMING&lt;/strong&gt;) the direction relative to the specified repository. Either &lt;strong&gt;INCOMING&lt;/strong&gt; or &lt;strong&gt;OUTGOING&lt;/strong&gt;. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetPullRequests1200Response**](GetPullRequests1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequestConditions

> []RestPullRequestCondition GetPullRequestConditions(ctx, projectKey).Execute()

Get default reviewer conditions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetPullRequestConditions(context.Background(), projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetPullRequestConditions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPullRequestConditions`: []RestPullRequestCondition
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetPullRequestConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RestPullRequestCondition**](RestPullRequestCondition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequestConditions1

> []RestPullRequestCondition GetPullRequestConditions1(ctx, projectKey, repositorySlug).Execute()

Get default reviewer conditions



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
	resp, r, err := apiClient.PullRequestsAPI.GetPullRequestConditions1(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetPullRequestConditions1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPullRequestConditions1`: []RestPullRequestCondition
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetPullRequestConditions1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestConditions1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RestPullRequestCondition**](RestPullRequestCondition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequests

> GetPullRequests1200Response GetPullRequests(ctx, projectKey, commitId, repositorySlug).Start(start).Limit(limit).Execute()

Get repository pull requests containing commit



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
	commitId := "commitId_example" // string | the commit ID
	repositorySlug := "repositorySlug_example" // string | The repository slug
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetPullRequests(context.Background(), projectKey, commitId, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetPullRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPullRequests`: GetPullRequests1200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetPullRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | the commit ID | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetPullRequests1200Response**](GetPullRequests1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReview

> GetComments200Response GetReview(ctx, projectKey, pullRequestId, repositorySlug).Start(start).Limit(limit).Execute()

Get pull request comment thread



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetReview(context.Background(), projectKey, pullRequestId, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReview`: GetComments200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetComments200Response**](GetComments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReviewerGroup

> RestReviewerGroup GetReviewerGroup(ctx, projectKey, id).Execute()

Get reviewer group



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
	id := "id_example" // string | The ID of the reviewer group to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetReviewerGroup(context.Background(), projectKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetReviewerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReviewerGroup`: RestReviewerGroup
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetReviewerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the reviewer group to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReviewerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestReviewerGroup**](RestReviewerGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReviewerGroup1

> RestReviewerGroup GetReviewerGroup1(ctx, projectKey, id, repositorySlug).Execute()

Get reviewer group



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
	id := "id_example" // string | The ID of the reviewer group to be retrieved
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetReviewerGroup1(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetReviewerGroup1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReviewerGroup1`: RestReviewerGroup
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetReviewerGroup1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the reviewer group to be retrieved | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReviewerGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestReviewerGroup**](RestReviewerGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReviewerGroups

> GetReviewerGroups1200Response GetReviewerGroups(ctx, projectKey).Start(start).Limit(limit).Execute()

Get all reviewer groups



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
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetReviewerGroups(context.Background(), projectKey).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetReviewerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReviewerGroups`: GetReviewerGroups1200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetReviewerGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReviewerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetReviewerGroups1200Response**](GetReviewerGroups1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReviewerGroups1

> GetReviewerGroups1200Response GetReviewerGroups1(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Execute()

Get all reviewer groups



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
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetReviewerGroups1(context.Background(), projectKey, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetReviewerGroups1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReviewerGroups1`: GetReviewerGroups1200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetReviewerGroups1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReviewerGroups1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetReviewerGroups1200Response**](GetReviewerGroups1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReviewers

> []RestPullRequestCondition GetReviewers(ctx, projectKey, repositorySlug).TargetRepoId(targetRepoId).SourceRepoId(sourceRepoId).SourceRefId(sourceRefId).TargetRefId(targetRefId).Execute()

Get required reviewers for PR creation



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
	targetRepoId := "targetRepoId_example" // string | The ID of the repository in which the target ref exists (optional)
	sourceRepoId := "sourceRepoId_example" // string | The ID of the repository in which the source ref exists (optional)
	sourceRefId := "sourceRefId_example" // string | The ID of the source ref (optional)
	targetRefId := "targetRefId_example" // string | The ID of the target ref (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetReviewers(context.Background(), projectKey, repositorySlug).TargetRepoId(targetRepoId).SourceRepoId(sourceRepoId).SourceRefId(sourceRefId).TargetRefId(targetRefId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetReviewers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReviewers`: []RestPullRequestCondition
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **targetRepoId** | **string** | The ID of the repository in which the target ref exists | 
 **sourceRepoId** | **string** | The ID of the repository in which the source ref exists | 
 **sourceRefId** | **string** | The ID of the source ref | 
 **targetRefId** | **string** | The ID of the target ref | 

### Return type

[**[]RestPullRequestCondition**](RestPullRequestCondition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []RestApplicationUser GetUsers(ctx, projectKey, id, repositorySlug).Execute()

Get reviewer group users



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
	id := "id_example" // string | The ID of the reviewer group to be retrieved
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.GetUsers(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: []RestApplicationUser
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the reviewer group to be retrieved | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RestApplicationUser**](RestApplicationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParticipants

> ListParticipants200Response ListParticipants(ctx, projectKey, pullRequestId, repositorySlug).Start(start).Limit(limit).Execute()

Get pull request participants



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
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.ListParticipants(context.Background(), projectKey, pullRequestId, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.ListParticipants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListParticipants`: ListParticipants200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.ListParticipants`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**ListParticipants200Response**](ListParticipants200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Merge

> RestPullRequest Merge(ctx, projectKey, pullRequestId, repositorySlug).Version(version).RestPullRequestMergeRequest(restPullRequestMergeRequest).Execute()

Merge pull request



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
	version := "version_example" // string | The current version of the pull request. If the server's version isn't the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the 'version' attribute in the returned JSON structure. (optional)
	restPullRequestMergeRequest := *openapiclient.NewRestPullRequestMergeRequest() // RestPullRequestMergeRequest | The body holder (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Merge(context.Background(), projectKey, pullRequestId, repositorySlug).Version(version).RestPullRequestMergeRequest(restPullRequestMergeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Merge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Merge`: RestPullRequest
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Merge`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **version** | **string** | The current version of the pull request. If the server&#39;s version isn&#39;t the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the &#39;version&#39; attribute in the returned JSON structure. | 
 **restPullRequestMergeRequest** | [**RestPullRequestMergeRequest**](RestPullRequestMergeRequest.md) | The body holder | 

### Return type

[**RestPullRequest**](RestPullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## React1

> RestUserReaction React1(ctx, projectKey, commentId, pullRequestId, emoticon, repositorySlug).Execute()

React to a PR comment



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
	commentId := "commentId_example" // string | The comment id.
	pullRequestId := "pullRequestId_example" // string | The pull request id.
	emoticon := "emoticon_example" // string | The emoticon to add
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.React1(context.Background(), projectKey, commentId, pullRequestId, emoticon, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.React1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `React1`: RestUserReaction
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.React1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The comment id. | 
**pullRequestId** | **string** | The pull request id. | 
**emoticon** | **string** | The emoticon to add | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReact1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**RestUserReaction**](RestUserReaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Rebase

> RestPullRequestRebaseResult Rebase(ctx, projectKey, pullRequestId, repositorySlug).RestPullRequestRebaseRequest(restPullRequestRebaseRequest).Execute()

Rebase pull request



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
	pullRequestId := "pullRequestId_example" // string | The ID of the pull request within the repository.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restPullRequestRebaseRequest := *openapiclient.NewRestPullRequestRebaseRequest() // RestPullRequestRebaseRequest | The pull request rebase request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Rebase(context.Background(), projectKey, pullRequestId, repositorySlug).RestPullRequestRebaseRequest(restPullRequestRebaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Rebase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Rebase`: RestPullRequestRebaseResult
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Rebase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The ID of the pull request within the repository. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restPullRequestRebaseRequest** | [**RestPullRequestRebaseRequest**](RestPullRequestRebaseRequest.md) | The pull request rebase request. | 

### Return type

[**RestPullRequestRebaseResult**](RestPullRequestRebaseResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reopen

> RestPullRequest Reopen(ctx, projectKey, pullRequestId, repositorySlug).Version(version).RestPullRequestReopenRequest(restPullRequestReopenRequest).Execute()

Re-open pull request



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
	version := "version_example" // string | The current version of the pull request. If the server's version isn't the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the 'version' attribute in the returned JSON structure. (optional)
	restPullRequestReopenRequest := *openapiclient.NewRestPullRequestReopenRequest() // RestPullRequestReopenRequest | The body holder (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Reopen(context.Background(), projectKey, pullRequestId, repositorySlug).Version(version).RestPullRequestReopenRequest(restPullRequestReopenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Reopen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Reopen`: RestPullRequest
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Reopen`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReopenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **version** | **string** | The current version of the pull request. If the server&#39;s version isn&#39;t the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the &#39;version&#39; attribute in the returned JSON structure. | 
 **restPullRequestReopenRequest** | [**RestPullRequestReopenRequest**](RestPullRequestReopenRequest.md) | The body holder | 

### Return type

[**RestPullRequest**](RestPullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> GetUsersWithoutAnyPermission200Response Search(ctx, projectKey, repositorySlug).Filter(filter).Role(role).Direction(direction).Start(start).Limit(limit).Execute()

Search pull request participants



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
	filter := "filter_example" // string | (optional) Return only users, whose username, name or email address <i>contain</i> the filter value (optional)
	role := "role_example" // string | (optional) The role associated with the pull request participant. This must be one of AUTHOR, REVIEWER, or PARTICIPANT (optional)
	direction := "direction_example" // string | (optional), Defaults to <strong>INCOMING</strong>) the direction relative to the specified repository. Either <strong>INCOMING</strong> or <strong>OUTGOING</strong>. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Search(context.Background(), projectKey, repositorySlug).Filter(filter).Role(role).Direction(direction).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: GetUsersWithoutAnyPermission200Response
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Search`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | (optional) Return only users, whose username, name or email address &lt;i&gt;contain&lt;/i&gt; the filter value | 
 **role** | **string** | (optional) The role associated with the pull request participant. This must be one of AUTHOR, REVIEWER, or PARTICIPANT | 
 **direction** | **string** | (optional), Defaults to &lt;strong&gt;INCOMING&lt;/strong&gt;) the direction relative to the specified repository. Either &lt;strong&gt;INCOMING&lt;/strong&gt; or &lt;strong&gt;OUTGOING&lt;/strong&gt;. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetUsersWithoutAnyPermission200Response**](GetUsersWithoutAnyPermission200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMergeConfig

> RestPullRequestMergeConfig SetMergeConfig(ctx, scmId).RestPullRequestSettings(restPullRequestSettings).Execute()

Update merge strategies



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
	scmId := "scmId_example" // string | the id of the scm to get strategies for
	restPullRequestSettings := *openapiclient.NewRestPullRequestSettings() // RestPullRequestSettings | the settings (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.SetMergeConfig(context.Background(), scmId).RestPullRequestSettings(restPullRequestSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.SetMergeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMergeConfig`: RestPullRequestMergeConfig
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.SetMergeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scmId** | **string** | the id of the scm to get strategies for | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMergeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restPullRequestSettings** | [**RestPullRequestSettings**](RestPullRequestSettings.md) | the settings | 

### Return type

[**RestPullRequestMergeConfig**](RestPullRequestMergeConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamChanges1

> RestChange StreamChanges1(ctx, projectKey, pullRequestId, repositorySlug).SinceId(sinceId).ChangeScope(changeScope).UntilId(untilId).WithComments(withComments).Start(start).Limit(limit).Execute()

Gets pull request changes



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	sinceId := "sinceId_example" // string | The since commit hash to stream changes for a RANGE arbitrary change scope (optional)
	changeScope := "changeScope_example" // string | UNREVIEWED to stream the unreviewed changes for the current user (if they exist); RANGE to stream changes between two arbitrary commits (requires 'sinceId' and 'untilId'); otherwise ALL to stream all changes (the default) (optional)
	untilId := "untilId_example" // string | The until commit hash to stream changes for a RANGE arbitrary change scope (optional)
	withComments := "withComments_example" // string | true to apply comment counts in the changes (the default); otherwise, false to stream changes without comment counts (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.StreamChanges1(context.Background(), projectKey, pullRequestId, repositorySlug).SinceId(sinceId).ChangeScope(changeScope).UntilId(untilId).WithComments(withComments).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.StreamChanges1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamChanges1`: RestChange
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.StreamChanges1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamChanges1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sinceId** | **string** | The since commit hash to stream changes for a RANGE arbitrary change scope | 
 **changeScope** | **string** | UNREVIEWED to stream the unreviewed changes for the current user (if they exist); RANGE to stream changes between two arbitrary commits (requires &#39;sinceId&#39; and &#39;untilId&#39;); otherwise ALL to stream all changes (the default) | 
 **untilId** | **string** | The until commit hash to stream changes for a RANGE arbitrary change scope | 
 **withComments** | **string** | true to apply comment counts in the changes (the default); otherwise, false to stream changes without comment counts | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**RestChange**](RestChange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamDiff2

> RestDiff StreamDiff2(ctx, path, projectKey, pullRequestId, repositorySlug).AvatarScheme(avatarScheme).ContextLines(contextLines).SinceId(sinceId).SrcPath(srcPath).DiffType(diffType).UntilId(untilId).Whitespace(whitespace).WithComments(withComments).AvatarSize(avatarSize).Execute()

Stream a diff within a pull request



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
	path := "path_example" // string | The path to the file which should be diffed (optional)
	projectKey := "projectKey_example" // string | The project key.
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	avatarScheme := "avatarScheme_example" // string | The security scheme for avatar URLs. If the scheme is not present then it is inherited from the request. It can be set to \"https\" to force the use of secure URLs. Not applicable if streaming raw diff (optional)
	contextLines := "contextLines_example" // string | The number of context lines to include around added/removed lines in the diff (optional)
	sinceId := "sinceId_example" // string | The since commit hash to stream a diff between two arbitrary hashes (optional)
	srcPath := "srcPath_example" // string | The previous path to the file, if the file has been copied, moved or renamed (optional)
	diffType := "diffType_example" // string | The type of diff being requested. When withComments is true this works as a hint to the system to attach the correct set of comments to the diff. Not applicable if streaming raw diff (optional)
	untilId := "untilId_example" // string | The until commit hash to stream a diff between two arbitrary hashes (optional)
	whitespace := "whitespace_example" // string | Optional whitespace flag which can be set to <code>ignore-all</code> (optional)
	withComments := "withComments_example" // string | <code>true</code> to embed comments in the diff (the default); otherwise, <code>false</code> to stream the diff without comments. Not applicable if streaming raw diff (optional)
	avatarSize := "avatarSize_example" // string | If present the service adds avatar URLs for comment authors where the provided value specifies the desired avatar size in pixels. Not applicable if streaming raw diff (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.StreamDiff2(context.Background(), path, projectKey, pullRequestId, repositorySlug).AvatarScheme(avatarScheme).ContextLines(contextLines).SinceId(sinceId).SrcPath(srcPath).DiffType(diffType).UntilId(untilId).Whitespace(whitespace).WithComments(withComments).AvatarSize(avatarSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.StreamDiff2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamDiff2`: RestDiff
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.StreamDiff2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to the file which should be diffed (optional) | 
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamDiff2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **avatarScheme** | **string** | The security scheme for avatar URLs. If the scheme is not present then it is inherited from the request. It can be set to \&quot;https\&quot; to force the use of secure URLs. Not applicable if streaming raw diff | 
 **contextLines** | **string** | The number of context lines to include around added/removed lines in the diff | 
 **sinceId** | **string** | The since commit hash to stream a diff between two arbitrary hashes | 
 **srcPath** | **string** | The previous path to the file, if the file has been copied, moved or renamed | 
 **diffType** | **string** | The type of diff being requested. When withComments is true this works as a hint to the system to attach the correct set of comments to the diff. Not applicable if streaming raw diff | 
 **untilId** | **string** | The until commit hash to stream a diff between two arbitrary hashes | 
 **whitespace** | **string** | Optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 
 **withComments** | **string** | &lt;code&gt;true&lt;/code&gt; to embed comments in the diff (the default); otherwise, &lt;code&gt;false&lt;/code&gt; to stream the diff without comments. Not applicable if streaming raw diff | 
 **avatarSize** | **string** | If present the service adds avatar URLs for comment authors where the provided value specifies the desired avatar size in pixels. Not applicable if streaming raw diff | 

### Return type

[**RestDiff**](RestDiff.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamPatch1

> StreamPatch1(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Stream pull request as patch



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
	r, err := apiClient.PullRequestsAPI.StreamPatch1(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.StreamPatch1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStreamPatch1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamRawDiff2

> StreamRawDiff2(ctx, projectKey, pullRequestId, repositorySlug).ContextLines(contextLines).Whitespace(whitespace).Execute()

Stream raw pull request diff



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
	contextLines := "contextLines_example" // string | The number of context lines to include around added/removed lines in the diff (optional)
	whitespace := "whitespace_example" // string | optional whitespace flag which can be set to <code>ignore-all</code> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.StreamRawDiff2(context.Background(), projectKey, pullRequestId, repositorySlug).ContextLines(contextLines).Whitespace(whitespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.StreamRawDiff2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStreamRawDiff2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contextLines** | **string** | The number of context lines to include around added/removed lines in the diff | 
 **whitespace** | **string** | optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TryAutoMerge

> RestAutoMergeProcessingResult TryAutoMerge(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Auto-merge pull request



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
	resp, r, err := apiClient.PullRequestsAPI.TryAutoMerge(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.TryAutoMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TryAutoMerge`: RestAutoMergeProcessingResult
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.TryAutoMerge`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiTryAutoMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestAutoMergeProcessingResult**](RestAutoMergeProcessingResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnReact1

> UnReact1(ctx, projectKey, commentId, pullRequestId, emoticon, repositorySlug).Execute()

Remove a reaction from a PR comment



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
	commentId := "commentId_example" // string | The comment id.
	pullRequestId := "pullRequestId_example" // string | The pull request id.
	emoticon := "emoticon_example" // string | The emoticon to remove
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.UnReact1(context.Background(), projectKey, commentId, pullRequestId, emoticon, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.UnReact1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The comment id. | 
**pullRequestId** | **string** | The pull request id. | 
**emoticon** | **string** | The emoticon to remove | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnReact1Request struct via the builder pattern


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


## UnassignParticipantRole

> UnassignParticipantRole(ctx, projectKey, userSlug, pullRequestId, repositorySlug).Execute()

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
	userSlug := "userSlug_example" // string | The slug for the user being unassigned
	pullRequestId := "pullRequestId_example" // string | The ID of the pull request within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.UnassignParticipantRole(context.Background(), projectKey, userSlug, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.UnassignParticipantRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**userSlug** | **string** | The slug for the user being unassigned | 
**pullRequestId** | **string** | The ID of the pull request within the repository | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignParticipantRoleRequest struct via the builder pattern


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
	r, err := apiClient.PullRequestsAPI.UnassignParticipantRole1(context.Background(), projectKey, pullRequestId, repositorySlug).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.UnassignParticipantRole1``: %v\n", err)
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


## Unwatch1

> Unwatch1(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Stop watching pull request



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.Unwatch1(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Unwatch1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnwatch1Request struct via the builder pattern


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


## Update

> RestPullRequest Update(ctx, projectKey, pullRequestId, repositorySlug).RestPullRequest(restPullRequest).Execute()

Update pull request metadata



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
	restPullRequest := *openapiclient.NewRestPullRequest() // RestPullRequest | The updated pull request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Update(context.Background(), projectKey, pullRequestId, repositorySlug).RestPullRequest(restPullRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: RestPullRequest
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Update`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restPullRequest** | [**RestPullRequest**](RestPullRequest.md) | The updated pull request | 

### Return type

[**RestPullRequest**](RestPullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update1

> RestReviewerGroup Update1(ctx, projectKey, id).RestReviewerGroup(restReviewerGroup).Execute()

Update reviewer group attributes



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
	id := "id_example" // string | The ID of the reviewer group to be updated
	restReviewerGroup := *openapiclient.NewRestReviewerGroup() // RestReviewerGroup | The request containing the attributes of the reviewer group to be updated. Only the attributes to be updated need to be present in this object. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Update1(context.Background(), projectKey, id).RestReviewerGroup(restReviewerGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Update1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update1`: RestReviewerGroup
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Update1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the reviewer group to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restReviewerGroup** | [**RestReviewerGroup**](RestReviewerGroup.md) | The request containing the attributes of the reviewer group to be updated. Only the attributes to be updated need to be present in this object. | 

### Return type

[**RestReviewerGroup**](RestReviewerGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update2

> RestReviewerGroup Update2(ctx, projectKey, id, repositorySlug).RestReviewerGroup(restReviewerGroup).Execute()

Update reviewer group attributes



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
	id := "id_example" // string | The ID of the reviewer group to be updated
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restReviewerGroup := *openapiclient.NewRestReviewerGroup() // RestReviewerGroup | The request containing the attributes of the reviewer group to be updated. Only the attributes to be updated need to be present in this object. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.Update2(context.Background(), projectKey, id, repositorySlug).RestReviewerGroup(restReviewerGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Update2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update2`: RestReviewerGroup
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.Update2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the reviewer group to be updated | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdate2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restReviewerGroup** | [**RestReviewerGroup**](RestReviewerGroup.md) | The request containing the attributes of the reviewer group to be updated. Only the attributes to be updated need to be present in this object. | 

### Return type

[**RestReviewerGroup**](RestReviewerGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComment1

> RestComment UpdateComment1(ctx, projectKey, commentId, pullRequestId, repositorySlug).RestComment(restComment).Execute()

Update pull request comment



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
	commentId := "commentId_example" // string | The ID of the comment to retrieve.
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restComment := *openapiclient.NewRestComment() // RestComment | The comment to add. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.UpdateComment1(context.Background(), projectKey, commentId, pullRequestId, repositorySlug).RestComment(restComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.UpdateComment1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComment1`: RestComment
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.UpdateComment1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The ID of the comment to retrieve. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComment1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **restComment** | [**RestComment**](RestComment.md) | The comment to add. | 

### Return type

[**RestComment**](RestComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComment2

> RestComment UpdateComment2(ctx, projectKey, commentId, pullRequestId, repositorySlug).RestComment(restComment).Execute()

Update pull request comment



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
	commentId := "commentId_example" // string | The ID of the comment to retrieve.
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restComment := *openapiclient.NewRestComment() // RestComment | The updated comment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.UpdateComment2(context.Background(), projectKey, commentId, pullRequestId, repositorySlug).RestComment(restComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.UpdateComment2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComment2`: RestComment
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.UpdateComment2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The ID of the comment to retrieve. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComment2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **restComment** | [**RestComment**](RestComment.md) | The updated comment | 

### Return type

[**RestComment**](RestComment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePullRequestCondition

> RestPullRequestCondition UpdatePullRequestCondition(ctx, projectKey, id).RestDefaultReviewersRequest(restDefaultReviewersRequest).Execute()

Update default reviewer condition



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
	id := "id_example" // string | The ID of the pull request condition.
	restDefaultReviewersRequest := *openapiclient.NewRestDefaultReviewersRequest() // RestDefaultReviewersRequest | The new details for the default reviewer pull request condition. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.UpdatePullRequestCondition(context.Background(), projectKey, id).RestDefaultReviewersRequest(restDefaultReviewersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.UpdatePullRequestCondition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePullRequestCondition`: RestPullRequestCondition
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.UpdatePullRequestCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the pull request condition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePullRequestConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restDefaultReviewersRequest** | [**RestDefaultReviewersRequest**](RestDefaultReviewersRequest.md) | The new details for the default reviewer pull request condition. | 

### Return type

[**RestPullRequestCondition**](RestPullRequestCondition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePullRequestCondition1

> RestPullRequestCondition UpdatePullRequestCondition1(ctx, projectKey, id, repositorySlug).UpdatePullRequestCondition1Request(updatePullRequestCondition1Request).Execute()

Update default reviewer condition



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
	id := "id_example" // string | The ID of the pull request condition
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	updatePullRequestCondition1Request := *openapiclient.NewUpdatePullRequestCondition1Request() // UpdatePullRequestCondition1Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.UpdatePullRequestCondition1(context.Background(), projectKey, id, repositorySlug).UpdatePullRequestCondition1Request(updatePullRequestCondition1Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.UpdatePullRequestCondition1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePullRequestCondition1`: RestPullRequestCondition
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.UpdatePullRequestCondition1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The ID of the pull request condition | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePullRequestCondition1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updatePullRequestCondition1Request** | [**UpdatePullRequestCondition1Request**](UpdatePullRequestCondition1Request.md) |  | 

### Return type

[**RestPullRequestCondition**](RestPullRequestCondition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStatus

> RestPullRequestParticipant UpdateStatus(ctx, projectKey, userSlug, pullRequestId, repositorySlug).RestPullRequestAssignStatusRequest(restPullRequestAssignStatusRequest).Version(version).Execute()

Change pull request status



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
	userSlug := "userSlug_example" // string | The slug for the user changing their status
	pullRequestId := "pullRequestId_example" // string | The ID of the pull request within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restPullRequestAssignStatusRequest := *openapiclient.NewRestPullRequestAssignStatusRequest() // RestPullRequestAssignStatusRequest | The participant representing the status to set, includes the status of the participant and last reviewed commit. If last reviewed commit is provided, it will be used to update the participant status. The operation will fail if the latest commit of the pull request does not match the provided last reviewed commit. If last reviewed commit is not provided, the latest commit of the pull request will be used for the update by default.
	version := "version_example" // string | The current version of the pull request. If the server's version isn't the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the 'version' attribute in the returned JSON structure. Note: This parameter is deprecated. Use last reviewed commit in request body instead (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PullRequestsAPI.UpdateStatus(context.Background(), projectKey, userSlug, pullRequestId, repositorySlug).RestPullRequestAssignStatusRequest(restPullRequestAssignStatusRequest).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.UpdateStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStatus`: RestPullRequestParticipant
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.UpdateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**userSlug** | **string** | The slug for the user changing their status | 
**pullRequestId** | **string** | The ID of the pull request within the repository | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **restPullRequestAssignStatusRequest** | [**RestPullRequestAssignStatusRequest**](RestPullRequestAssignStatusRequest.md) | The participant representing the status to set, includes the status of the participant and last reviewed commit. If last reviewed commit is provided, it will be used to update the participant status. The operation will fail if the latest commit of the pull request does not match the provided last reviewed commit. If last reviewed commit is not provided, the latest commit of the pull request will be used for the update by default. | 
 **version** | **string** | The current version of the pull request. If the server&#39;s version isn&#39;t the same as the specified version the operation will fail. To determine the current version of the pull request it should be fetched from the server prior to this operation. Look for the &#39;version&#39; attribute in the returned JSON structure. Note: This parameter is deprecated. Use last reviewed commit in request body instead | 

### Return type

[**RestPullRequestParticipant**](RestPullRequestParticipant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Watch1

> Watch1(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Watch pull request



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
	pullRequestId := "pullRequestId_example" // string | The pull request ID.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PullRequestsAPI.Watch1(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.Watch1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**pullRequestId** | **string** | The pull request ID. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatch1Request struct via the builder pattern


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
	resp, r, err := apiClient.PullRequestsAPI.WithdrawApproval(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PullRequestsAPI.WithdrawApproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WithdrawApproval`: RestPullRequestParticipant
	fmt.Fprintf(os.Stdout, "Response from `PullRequestsAPI.WithdrawApproval`: %v\n", resp)
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


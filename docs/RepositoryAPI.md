# \RepositoryAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDefaultTask1**](RepositoryAPI.md#AddDefaultTask1) | **Post** /default-tasks/latest/projects/{projectKey}/repos/{repositorySlug}/tasks | Add a default task
[**AddLabel**](RepositoryAPI.md#AddLabel) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/labels | Add repository label
[**CreateBranch**](RepositoryAPI.md#CreateBranch) | **Post** /branch-utils/latest/projects/{projectKey}/repos/{repositorySlug}/branches | Create branch
[**CreateBranchForRepository**](RepositoryAPI.md#CreateBranchForRepository) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/branches | Create branch
[**CreateComment**](RepositoryAPI.md#CreateComment) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments | Add a new commit comment
[**CreateRestrictions1**](RepositoryAPI.md#CreateRestrictions1) | **Post** /branch-permissions/latest/projects/{projectKey}/repos/{repositorySlug}/restrictions | Create multiple ref restrictions
[**CreateTag**](RepositoryAPI.md#CreateTag) | **Post** /git/latest/projects/{projectKey}/repos/{repositorySlug}/tags | Create tag
[**CreateTagForRepository**](RepositoryAPI.md#CreateTagForRepository) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/tags | Create tag
[**CreateWebhook1**](RepositoryAPI.md#CreateWebhook1) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks | Create webhook
[**Delete5**](RepositoryAPI.md#Delete5) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/auto-merge | Delete pull request auto-merge settings
[**DeleteAllDefaultTasks1**](RepositoryAPI.md#DeleteAllDefaultTasks1) | **Delete** /default-tasks/latest/projects/{projectKey}/repos/{repositorySlug}/tasks | Deletes all default tasks for the repository
[**DeleteAttachment**](RepositoryAPI.md#DeleteAttachment) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/attachments/{attachmentId} | Delete an attachment
[**DeleteAttachmentMetadata**](RepositoryAPI.md#DeleteAttachmentMetadata) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/attachments/{attachmentId}/metadata | Delete attachment metadata
[**DeleteAutoDeclineSettings1**](RepositoryAPI.md#DeleteAutoDeclineSettings1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/auto-decline | Delete auto decline settings
[**DeleteBranch**](RepositoryAPI.md#DeleteBranch) | **Delete** /branch-utils/latest/projects/{projectKey}/repos/{repositorySlug}/branches | Delete branch
[**DeleteComment**](RepositoryAPI.md#DeleteComment) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId} | Delete a commit comment
[**DeleteDefaultTask1**](RepositoryAPI.md#DeleteDefaultTask1) | **Delete** /default-tasks/latest/projects/{projectKey}/repos/{repositorySlug}/tasks/{taskId} | Delete a specific default task
[**DeleteRepositoryHook**](RepositoryAPI.md#DeleteRepositoryHook) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey} | Delete repository hook
[**DeleteRestriction1**](RepositoryAPI.md#DeleteRestriction1) | **Delete** /branch-permissions/latest/projects/{projectKey}/repos/{repositorySlug}/restrictions/{id} | Delete a ref restriction
[**DeleteTag**](RepositoryAPI.md#DeleteTag) | **Delete** /git/latest/projects/{projectKey}/repos/{repositorySlug}/tags/{name} | Delete tag
[**DeleteWebhook1**](RepositoryAPI.md#DeleteWebhook1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Delete webhook
[**DisableHook1**](RepositoryAPI.md#DisableHook1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/enabled | Disable repository hook
[**EditFile**](RepositoryAPI.md#EditFile) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/browse/{path} | Edit file
[**EnableHook1**](RepositoryAPI.md#EnableHook1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/enabled | Enable repository hook
[**FindBranches**](RepositoryAPI.md#FindBranches) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/ref-change-activities/branches | Get branches with ref change activities for repository
[**FindByCommit**](RepositoryAPI.md#FindByCommit) | **Get** /branch-utils/latest/projects/{projectKey}/repos/{repositorySlug}/branches/info/{commitId} | Get branch
[**FindWebhooks1**](RepositoryAPI.md#FindWebhooks1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks | Find webhooks
[**Get5**](RepositoryAPI.md#Get5) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/auto-merge | Get pull request auto-merge settings
[**GetAllLabelsForRepository**](RepositoryAPI.md#GetAllLabelsForRepository) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/labels | Get repository labels
[**GetArchive**](RepositoryAPI.md#GetArchive) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/archive | Stream archive of repository
[**GetAttachment**](RepositoryAPI.md#GetAttachment) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/attachments/{attachmentId} | Get an attachment
[**GetAttachmentMetadata**](RepositoryAPI.md#GetAttachmentMetadata) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/attachments/{attachmentId}/metadata | Get attachment metadata
[**GetAutoDeclineSettings1**](RepositoryAPI.md#GetAutoDeclineSettings1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/auto-decline | Get auto decline settings
[**GetBranches**](RepositoryAPI.md#GetBranches) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/branches | Find branches
[**GetChanges**](RepositoryAPI.md#GetChanges) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/changes | Get changes in commit
[**GetChanges1**](RepositoryAPI.md#GetChanges1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/changes | Get changes made in commit
[**GetComment**](RepositoryAPI.md#GetComment) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId} | Get a commit comment
[**GetComments**](RepositoryAPI.md#GetComments) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments | Search for commit comments
[**GetCommit**](RepositoryAPI.md#GetCommit) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId} | Get commit by ID
[**GetCommits**](RepositoryAPI.md#GetCommits) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits | Get commits
[**GetConfigurations1**](RepositoryAPI.md#GetConfigurations1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/hook-scripts | Get hook scripts
[**GetContent**](RepositoryAPI.md#GetContent) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/browse | Get file content at revision
[**GetContent1**](RepositoryAPI.md#GetContent1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/browse/{path} | Get file content
[**GetDefaultBranch1**](RepositoryAPI.md#GetDefaultBranch1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/branches/default | Get default branch
[**GetDefaultTasks1**](RepositoryAPI.md#GetDefaultTasks1) | **Get** /default-tasks/latest/projects/{projectKey}/repos/{repositorySlug}/tasks | Get a page of default tasks
[**GetDiffStatsSummary**](RepositoryAPI.md#GetDiffStatsSummary) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/diff-stats-summary/{path} | Get diff stats summary between revisions
[**GetDiffStatsSummary1**](RepositoryAPI.md#GetDiffStatsSummary1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/compare/diff-stats-summary{path} | Retrieve the diff stats summary between commits
[**GetLatestInvocation1**](RepositoryAPI.md#GetLatestInvocation1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}/latest | Get last webhook invocation details
[**GetMergeBase**](RepositoryAPI.md#GetMergeBase) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/merge-base | Get the common ancestor between two commits
[**GetPullRequestSettings1**](RepositoryAPI.md#GetPullRequestSettings1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/pull-requests | Get pull request settings
[**GetRefChangeActivity**](RepositoryAPI.md#GetRefChangeActivity) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/ref-change-activities | Get ref change activity
[**GetRepositories1**](RepositoryAPI.md#GetRepositories1) | **Get** /api/latest/repos | Search for repositories
[**GetRepositoriesRecentlyAccessed**](RepositoryAPI.md#GetRepositoriesRecentlyAccessed) | **Get** /api/latest/profile/recent/repos | Get recently accessed repositories
[**GetRepositoryHook1**](RepositoryAPI.md#GetRepositoryHook1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey} | Get repository hook
[**GetRepositoryHooks1**](RepositoryAPI.md#GetRepositoryHooks1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/hooks | Get repository hooks
[**GetRestriction1**](RepositoryAPI.md#GetRestriction1) | **Get** /branch-permissions/latest/projects/{projectKey}/repos/{repositorySlug}/restrictions/{id} | Get a ref restriction
[**GetRestrictions1**](RepositoryAPI.md#GetRestrictions1) | **Get** /branch-permissions/latest/projects/{projectKey}/repos/{repositorySlug}/restrictions | Search for ref restrictions
[**GetSettings1**](RepositoryAPI.md#GetSettings1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/settings | Get repository hook settings
[**GetStatistics1**](RepositoryAPI.md#GetStatistics1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}/statistics | Get webhook statistics
[**GetStatisticsSummary1**](RepositoryAPI.md#GetStatisticsSummary1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId}/statistics/summary | Get webhook statistics summary
[**GetTag**](RepositoryAPI.md#GetTag) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/tags/{name} | Get tag
[**GetTags**](RepositoryAPI.md#GetTags) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/tags | Find tag
[**GetWebhook1**](RepositoryAPI.md#GetWebhook1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Get webhook
[**MGetStatus**](RepositoryAPI.md#MGetStatus) | **Get** /sync/latest/projects/{projectKey}/repos/{repositorySlug} | Get synchronization status
[**React**](RepositoryAPI.md#React) | **Put** /comment-likes/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId}/reactions/{emoticon} | React to a comment
[**RemoveConfiguration1**](RepositoryAPI.md#RemoveConfiguration1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/hook-scripts/{scriptId} | Remove a hook script
[**RemoveLabel**](RepositoryAPI.md#RemoveLabel) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/labels/{labelName} | Remove repository label
[**SaveAttachmentMetadata**](RepositoryAPI.md#SaveAttachmentMetadata) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/attachments/{attachmentId}/metadata | Save attachment metadata
[**SearchWebhooks**](RepositoryAPI.md#SearchWebhooks) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks/search | Search webhooks
[**Set1**](RepositoryAPI.md#Set1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/auto-merge | Create or update the pull request auto-merge settings
[**SetAutoDeclineSettings1**](RepositoryAPI.md#SetAutoDeclineSettings1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/auto-decline | Create auto decline settings
[**SetConfiguration1**](RepositoryAPI.md#SetConfiguration1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/hook-scripts/{scriptId} | Create/update a hook script
[**SetDefaultBranch1**](RepositoryAPI.md#SetDefaultBranch1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/branches/default | Update default branch
[**SetEnabled**](RepositoryAPI.md#SetEnabled) | **Post** /sync/latest/projects/{projectKey}/repos/{repositorySlug} | Disable synchronization
[**SetSettings1**](RepositoryAPI.md#SetSettings1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/hooks/{hookKey}/settings | Update repository hook settings
[**Stream**](RepositoryAPI.md#Stream) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/last-modified | Stream files
[**Stream1**](RepositoryAPI.md#Stream1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/last-modified/{path} | Stream files with last modified commit in path
[**StreamChanges**](RepositoryAPI.md#StreamChanges) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/compare/changes | Compare commits
[**StreamCommits**](RepositoryAPI.md#StreamCommits) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/compare/commits | Get accessible commits
[**StreamDiff**](RepositoryAPI.md#StreamDiff) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/diff/{path} | Get diff between revisions
[**StreamDiff1**](RepositoryAPI.md#StreamDiff1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/compare/diff{path} | Get diff between commits
[**StreamFiles**](RepositoryAPI.md#StreamFiles) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/files | Get files in directory
[**StreamFiles1**](RepositoryAPI.md#StreamFiles1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/files/{path} | Get files in directory
[**StreamPatch**](RepositoryAPI.md#StreamPatch) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/patch | Get patch content at revision
[**StreamRaw**](RepositoryAPI.md#StreamRaw) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/raw/{path} | Get raw content of a file at revision
[**StreamRawDiff**](RepositoryAPI.md#StreamRawDiff) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/diff | Get raw diff for path
[**StreamRawDiff1**](RepositoryAPI.md#StreamRawDiff1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/diff/{path} | Get raw diff for path
[**Synchronize**](RepositoryAPI.md#Synchronize) | **Post** /sync/latest/projects/{projectKey}/repos/{repositorySlug}/synchronize | Manual synchronization
[**TestWebhook1**](RepositoryAPI.md#TestWebhook1) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks/test | Test webhook
[**UnReact**](RepositoryAPI.md#UnReact) | **Delete** /comment-likes/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId}/reactions/{emoticon} | Remove a reaction from comment
[**Unwatch**](RepositoryAPI.md#Unwatch) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/watch | Stop watching commit
[**Unwatch2**](RepositoryAPI.md#Unwatch2) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/watch | Stop watching repository
[**UpdateComment**](RepositoryAPI.md#UpdateComment) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments/{commentId} | Update a commit comment
[**UpdateDefaultTask1**](RepositoryAPI.md#UpdateDefaultTask1) | **Put** /default-tasks/latest/projects/{projectKey}/repos/{repositorySlug}/tasks/{taskId} | Update a default task
[**UpdatePullRequestSettings1**](RepositoryAPI.md#UpdatePullRequestSettings1) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/settings/pull-requests | Update pull request settings
[**UpdateWebhook1**](RepositoryAPI.md#UpdateWebhook1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Update webhook
[**Watch**](RepositoryAPI.md#Watch) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/watch | Watch commit
[**Watch2**](RepositoryAPI.md#Watch2) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/watch | Watch repository



## AddDefaultTask1

> RestDefaultTask AddDefaultTask1(ctx, projectKey, repositorySlug).RestDefaultTaskRequest(restDefaultTaskRequest).Execute()

Add a default task



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
	restDefaultTaskRequest := *openapiclient.NewRestDefaultTaskRequest() // RestDefaultTaskRequest | The task to be added

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.AddDefaultTask1(context.Background(), projectKey, repositorySlug).RestDefaultTaskRequest(restDefaultTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.AddDefaultTask1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDefaultTask1`: RestDefaultTask
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.AddDefaultTask1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDefaultTask1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restDefaultTaskRequest** | [**RestDefaultTaskRequest**](RestDefaultTaskRequest.md) | The task to be added | 

### Return type

[**RestDefaultTask**](RestDefaultTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLabel

> RestLabel AddLabel(ctx, projectKey, repositorySlug).RestLabel(restLabel).Execute()

Add repository label



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
	restLabel := *openapiclient.NewRestLabel() // RestLabel | The label to apply (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.AddLabel(context.Background(), projectKey, repositorySlug).RestLabel(restLabel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.AddLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLabel`: RestLabel
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.AddLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restLabel** | [**RestLabel**](RestLabel.md) | The label to apply | 

### Return type

[**RestLabel**](RestLabel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBranch

> RestBranch CreateBranch(ctx, projectKey, repositorySlug).RestBranchCreateRequest(restBranchCreateRequest).Execute()

Create branch



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
	restBranchCreateRequest := *openapiclient.NewRestBranchCreateRequest() // RestBranchCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.CreateBranch(context.Background(), projectKey, repositorySlug).RestBranchCreateRequest(restBranchCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.CreateBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBranch`: RestBranch
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.CreateBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restBranchCreateRequest** | [**RestBranchCreateRequest**](RestBranchCreateRequest.md) |  | 

### Return type

[**RestBranch**](RestBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBranchForRepository

> RestBranch CreateBranchForRepository(ctx, projectKey, repositorySlug).RestCreateBranchRequest(restCreateBranchRequest).Execute()

Create branch



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
	restCreateBranchRequest := *openapiclient.NewRestCreateBranchRequest() // RestCreateBranchRequest | The request to create a branch containing a <strong>name</strong>, <strong>startPoint</strong>, and optionally a <strong>message</strong> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.CreateBranchForRepository(context.Background(), projectKey, repositorySlug).RestCreateBranchRequest(restCreateBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.CreateBranchForRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBranchForRepository`: RestBranch
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.CreateBranchForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBranchForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restCreateBranchRequest** | [**RestCreateBranchRequest**](RestCreateBranchRequest.md) | The request to create a branch containing a &lt;strong&gt;name&lt;/strong&gt;, &lt;strong&gt;startPoint&lt;/strong&gt;, and optionally a &lt;strong&gt;message&lt;/strong&gt; | 

### Return type

[**RestBranch**](RestBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComment

> RestComment CreateComment(ctx, projectKey, commitId, repositorySlug).Since(since).RestComment(restComment).Execute()

Add a new commit comment



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
	commitId := "commitId_example" // string | The <i>full ID</i> of the commit within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug
	since := "since_example" // string | For a merge commit, a parent can be provided to specify which diff the comments should be on. For a commit range, a sinceId can be provided to specify where the comments should be anchored from. (optional)
	restComment := *openapiclient.NewRestComment() // RestComment | the comment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.CreateComment(context.Background(), projectKey, commitId, repositorySlug).Since(since).RestComment(restComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.CreateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComment`: RestComment
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.CreateComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | The &lt;i&gt;full ID&lt;/i&gt; of the commit within the repository | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **since** | **string** | For a merge commit, a parent can be provided to specify which diff the comments should be on. For a commit range, a sinceId can be provided to specify where the comments should be anchored from. | 
 **restComment** | [**RestComment**](RestComment.md) | the comment | 

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


## CreateRestrictions1

> RestRefRestriction CreateRestrictions1(ctx, projectKey, repositorySlug).RestRestrictionRequest(restRestrictionRequest).Execute()

Create multiple ref restrictions



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
	restRestrictionRequest := []openapiclient.RestRestrictionRequest{*openapiclient.NewRestRestrictionRequest()} // []RestRestrictionRequest | The request containing a list of the details of the restrictions to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.CreateRestrictions1(context.Background(), projectKey, repositorySlug).RestRestrictionRequest(restRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.CreateRestrictions1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRestrictions1`: RestRefRestriction
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.CreateRestrictions1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRestrictions1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restRestrictionRequest** | [**[]RestRestrictionRequest**](RestRestrictionRequest.md) | The request containing a list of the details of the restrictions to create. | 

### Return type

[**RestRefRestriction**](RestRefRestriction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.atl.bitbucket.bulk+json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTag

> RestTag CreateTag(ctx, projectKey, repositorySlug).RestGitTagCreateRequest(restGitTagCreateRequest).Execute()

Create tag



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
	restGitTagCreateRequest := *openapiclient.NewRestGitTagCreateRequest() // RestGitTagCreateRequest | The create git tag request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.CreateTag(context.Background(), projectKey, repositorySlug).RestGitTagCreateRequest(restGitTagCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: RestTag
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.CreateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restGitTagCreateRequest** | [**RestGitTagCreateRequest**](RestGitTagCreateRequest.md) | The create git tag request. | 

### Return type

[**RestTag**](RestTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagForRepository

> RestTag CreateTagForRepository(ctx, projectKey, repositorySlug).RestCreateTagRequest(restCreateTagRequest).Execute()

Create tag



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
	restCreateTagRequest := *openapiclient.NewRestCreateTagRequest() // RestCreateTagRequest | The request to create a tag containing a <strong>name</strong>, <strong>startPoint</strong>, and optionally a <strong>message</strong> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.CreateTagForRepository(context.Background(), projectKey, repositorySlug).RestCreateTagRequest(restCreateTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.CreateTagForRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTagForRepository`: RestTag
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.CreateTagForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restCreateTagRequest** | [**RestCreateTagRequest**](RestCreateTagRequest.md) | The request to create a tag containing a &lt;strong&gt;name&lt;/strong&gt;, &lt;strong&gt;startPoint&lt;/strong&gt;, and optionally a &lt;strong&gt;message&lt;/strong&gt; | 

### Return type

[**RestTag**](RestTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook1

> RestWebhook CreateWebhook1(ctx, projectKey, repositorySlug).RestWebhook(restWebhook).Execute()

Create webhook



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
	restWebhook := *openapiclient.NewRestWebhook() // RestWebhook | The webhook to be created for this repository. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.CreateWebhook1(context.Background(), projectKey, repositorySlug).RestWebhook(restWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.CreateWebhook1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook1`: RestWebhook
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.CreateWebhook1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhook1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restWebhook** | [**RestWebhook**](RestWebhook.md) | The webhook to be created for this repository. | 

### Return type

[**RestWebhook**](RestWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete5

> Delete5(ctx, projectKey, repositorySlug).Execute()

Delete pull request auto-merge settings



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
	r, err := apiClient.RepositoryAPI.Delete5(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Delete5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete5Request struct via the builder pattern


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


## DeleteAllDefaultTasks1

> DeleteAllDefaultTasks1(ctx, projectKey, repositorySlug).Execute()

Deletes all default tasks for the repository



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
	r, err := apiClient.RepositoryAPI.DeleteAllDefaultTasks1(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteAllDefaultTasks1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAllDefaultTasks1Request struct via the builder pattern


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


## DeleteAttachment

> DeleteAttachment(ctx, projectKey, attachmentId, repositorySlug).Execute()

Delete an attachment



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
	attachmentId := "attachmentId_example" // string | the attachment ID
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.DeleteAttachment(context.Background(), projectKey, attachmentId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**attachmentId** | **string** | the attachment ID | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentRequest struct via the builder pattern


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


## DeleteAttachmentMetadata

> DeleteAttachmentMetadata(ctx, projectKey, attachmentId, repositorySlug).Execute()

Delete attachment metadata



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
	attachmentId := "attachmentId_example" // string | the attachment ID
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.DeleteAttachmentMetadata(context.Background(), projectKey, attachmentId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteAttachmentMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**attachmentId** | **string** | the attachment ID | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentMetadataRequest struct via the builder pattern


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


## DeleteAutoDeclineSettings1

> DeleteAutoDeclineSettings1(ctx, projectKey, repositorySlug).Execute()

Delete auto decline settings



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
	r, err := apiClient.RepositoryAPI.DeleteAutoDeclineSettings1(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteAutoDeclineSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoDeclineSettings1Request struct via the builder pattern


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


## DeleteBranch

> DeleteBranch(ctx, projectKey, repositorySlug).Execute()

Delete branch



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
	r, err := apiClient.RepositoryAPI.DeleteBranch(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteBranch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteBranchRequest struct via the builder pattern


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


## DeleteComment

> DeleteComment(ctx, projectKey, commentId, commitId, repositorySlug).Version(version).Execute()

Delete a commit comment



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
	commentId := "commentId_example" // string | the comment
	commitId := "commitId_example" // string | The <i>full ID</i> of the commit within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug
	version := "version_example" // string | The expected version of the comment. This must match the server's version of the comment or the delete will fail. To determine the current version of the comment, the comment should be fetched from the server prior to the delete. Look for the 'version' attribute in the returned JSON structure. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.DeleteComment(context.Background(), projectKey, commentId, commitId, repositorySlug).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commentId** | **string** | the comment | 
**commitId** | **string** | The &lt;i&gt;full ID&lt;/i&gt; of the commit within the repository | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentRequest struct via the builder pattern


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


## DeleteDefaultTask1

> DeleteDefaultTask1(ctx, projectKey, repositorySlug, taskId).Execute()

Delete a specific default task



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
	taskId := "taskId_example" // string | The ID of the default task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.DeleteDefaultTask1(context.Background(), projectKey, repositorySlug, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteDefaultTask1``: %v\n", err)
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
**taskId** | **string** | The ID of the default task | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultTask1Request struct via the builder pattern


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


## DeleteRepositoryHook

> DeleteRepositoryHook(ctx, projectKey, hookKey, repositorySlug).Execute()

Delete repository hook



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
	hookKey := "hookKey_example" // string | The hook key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.DeleteRepositoryHook(context.Background(), projectKey, hookKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteRepositoryHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryHookRequest struct via the builder pattern


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


## DeleteRestriction1

> DeleteRestriction1(ctx, projectKey, id, repositorySlug).Execute()

Delete a ref restriction



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
	id := "id_example" // string | The restriction id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.DeleteRestriction1(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteRestriction1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The restriction id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRestriction1Request struct via the builder pattern


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


## DeleteTag

> DeleteTag(ctx, projectKey, name, repositorySlug).Execute()

Delete tag



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
	name := "name_example" // string | The name of the tag to be deleted.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.DeleteTag(context.Background(), projectKey, name, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**name** | **string** | The name of the tag to be deleted. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## DeleteWebhook1

> DeleteWebhook1(ctx, projectKey, webhookId, repositorySlug).Execute()

Delete webhook



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
	webhookId := "webhookId_example" // string | The ID of the webhook to be deleted.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.DeleteWebhook1(context.Background(), projectKey, webhookId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteWebhook1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | The ID of the webhook to be deleted. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhook1Request struct via the builder pattern


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


## DisableHook1

> RestRepositoryHook DisableHook1(ctx, projectKey, hookKey, repositorySlug).Execute()

Disable repository hook



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
	hookKey := "hookKey_example" // string | The hook key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.DisableHook1(context.Background(), projectKey, hookKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DisableHook1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableHook1`: RestRepositoryHook
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.DisableHook1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableHook1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestRepositoryHook**](RestRepositoryHook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditFile

> RestCommit EditFile(ctx, path, projectKey, repositorySlug).Branch(branch).Content(content).Message(message).SourceBranch(sourceBranch).SourceCommitId(sourceCommitId).Execute()

Edit file



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
	path := "path_example" // string | The path of the file that is to be modified or created
	projectKey := "projectKey_example" // string | The project key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	branch := "branch_example" // string | The branch on which the <code>path</code> should be modified or created. (optional)
	content := "content_example" // string | The full content of the file at <code>path</code>. (optional)
	message := "message_example" // string | The message associated with this change, to be used as the commit message. Or null if the default message should be used. (optional)
	sourceBranch := "sourceBranch_example" // string | The starting point for <code>branch</code>. If provided and different from <code>branch</code>, <code>branch</code> will be created as a new branch, branching off from <code>sourceBranch</code>. (optional)
	sourceCommitId := "sourceCommitId_example" // string | The commit ID of the file before it was edited, used to identify if content has changed. Or null if this is a new file (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.EditFile(context.Background(), path, projectKey, repositorySlug).Branch(branch).Content(content).Message(message).SourceBranch(sourceBranch).SourceCommitId(sourceCommitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.EditFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditFile`: RestCommit
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.EditFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path of the file that is to be modified or created | 
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branch** | **string** | The branch on which the &lt;code&gt;path&lt;/code&gt; should be modified or created. | 
 **content** | **string** | The full content of the file at &lt;code&gt;path&lt;/code&gt;. | 
 **message** | **string** | The message associated with this change, to be used as the commit message. Or null if the default message should be used. | 
 **sourceBranch** | **string** | The starting point for &lt;code&gt;branch&lt;/code&gt;. If provided and different from &lt;code&gt;branch&lt;/code&gt;, &lt;code&gt;branch&lt;/code&gt; will be created as a new branch, branching off from &lt;code&gt;sourceBranch&lt;/code&gt;. | 
 **sourceCommitId** | **string** | The commit ID of the file before it was edited, used to identify if content has changed. Or null if this is a new file | 

### Return type

[**RestCommit**](RestCommit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableHook1

> RestRepositoryHook EnableHook1(ctx, projectKey, hookKey, repositorySlug).ContentLength(contentLength).Execute()

Enable repository hook



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
	hookKey := "hookKey_example" // string | The hook key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	contentLength := "contentLength_example" // string | The content length. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.EnableHook1(context.Background(), projectKey, hookKey, repositorySlug).ContentLength(contentLength).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.EnableHook1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableHook1`: RestRepositoryHook
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.EnableHook1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableHook1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentLength** | **string** | The content length. | 

### Return type

[**RestRepositoryHook**](RestRepositoryHook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindBranches

> FindByCommit200Response FindBranches(ctx, projectKey, repositorySlug).FilterText(filterText).Start(start).Limit(limit).Execute()

Get branches with ref change activities for repository



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
	filterText := "filterText_example" // string | (optional) Partial match for a ref ID to filter minimal refs for (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.FindBranches(context.Background(), projectKey, repositorySlug).FilterText(filterText).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.FindBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindBranches`: FindByCommit200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.FindBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterText** | **string** | (optional) Partial match for a ref ID to filter minimal refs for | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**FindByCommit200Response**](FindByCommit200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindByCommit

> FindByCommit200Response FindByCommit(ctx, projectKey, commitId, repositorySlug).Start(start).Limit(limit).Execute()

Get branch



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
	commitId := "commitId_example" // string | 
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.FindByCommit(context.Background(), projectKey, commitId, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.FindByCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindByCommit`: FindByCommit200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.FindByCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** |  | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindByCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**FindByCommit200Response**](FindByCommit200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWebhooks1

> FindWebhooks1(ctx, projectKey, repositorySlug).Event(event).Statistics(statistics).Execute()

Find webhooks



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
	event := "event_example" // string | List of <code>com.atlassian.webhooks.WebhookEvent</code> IDs to filter for (optional)
	statistics := true // bool | <code>true</code> if statistics should be provided for all found webhooks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.FindWebhooks1(context.Background(), projectKey, repositorySlug).Event(event).Statistics(statistics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.FindWebhooks1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFindWebhooks1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **event** | **string** | List of &lt;code&gt;com.atlassian.webhooks.WebhookEvent&lt;/code&gt; IDs to filter for | 
 **statistics** | **bool** | &lt;code&gt;true&lt;/code&gt; if statistics should be provided for all found webhooks | 

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


## Get5

> RestAutoMergeRestrictedSettings Get5(ctx, projectKey, repositorySlug).Execute()

Get pull request auto-merge settings



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
	resp, r, err := apiClient.RepositoryAPI.Get5(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Get5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get5`: RestAutoMergeRestrictedSettings
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.Get5`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestAutoMergeRestrictedSettings**](RestAutoMergeRestrictedSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllLabelsForRepository

> RestLabel GetAllLabelsForRepository(ctx, projectKey, repositorySlug).Execute()

Get repository labels



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
	resp, r, err := apiClient.RepositoryAPI.GetAllLabelsForRepository(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetAllLabelsForRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllLabelsForRepository`: RestLabel
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetAllLabelsForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLabelsForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestLabel**](RestLabel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchive

> GetArchive(ctx, projectKey, repositorySlug).Path(path).Filename(filename).At(at).Prefix(prefix).Format(format).Execute()

Stream archive of repository



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
	path := "path_example" // string | Paths to include in the streamed archive; may be repeated to include multiple paths (optional)
	filename := "filename_example" // string | A filename to include the \"Content-Disposition\" header (optional)
	at := "at_example" // string | The commit to stream an archive of; if not supplied, an archive of the default branch is streamed (optional)
	prefix := "prefix_example" // string | A prefix to apply to all entries in the streamed archive; if the supplied prefix does not end with a trailing /, one will be added automatically (optional)
	format := "format_example" // string | The format to stream the archive in; must be one of: zip, tar, tar.gz or tgz (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.GetArchive(context.Background(), projectKey, repositorySlug).Path(path).Filename(filename).At(at).Prefix(prefix).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetArchive``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **path** | **string** | Paths to include in the streamed archive; may be repeated to include multiple paths | 
 **filename** | **string** | A filename to include the \&quot;Content-Disposition\&quot; header | 
 **at** | **string** | The commit to stream an archive of; if not supplied, an archive of the default branch is streamed | 
 **prefix** | **string** | A prefix to apply to all entries in the streamed archive; if the supplied prefix does not end with a trailing /, one will be added automatically | 
 **format** | **string** | The format to stream the archive in; must be one of: zip, tar, tar.gz or tgz | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/x-tar, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachment

> GetAttachment(ctx, projectKey, attachmentId, repositorySlug).UserAgent(userAgent).Range_(range_).Execute()

Get an attachment



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
	attachmentId := "attachmentId_example" // string | the attachment ID
	repositorySlug := "repositorySlug_example" // string | The repository slug
	userAgent := "userAgent_example" // string |  (optional)
	range_ := "range__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.GetAttachment(context.Background(), projectKey, attachmentId, repositorySlug).UserAgent(userAgent).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**attachmentId** | **string** | the attachment ID | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userAgent** | **string** |  | 
 **range_** | **string** |  | 

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


## GetAttachmentMetadata

> RestAttachmentMetadata GetAttachmentMetadata(ctx, projectKey, attachmentId, repositorySlug).Execute()

Get attachment metadata



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
	attachmentId := "attachmentId_example" // string | the attachment ID
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetAttachmentMetadata(context.Background(), projectKey, attachmentId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetAttachmentMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentMetadata`: RestAttachmentMetadata
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetAttachmentMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**attachmentId** | **string** | the attachment ID | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestAttachmentMetadata**](RestAttachmentMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoDeclineSettings1

> RestAutoDeclineSettings GetAutoDeclineSettings1(ctx, projectKey, repositorySlug).Execute()

Get auto decline settings



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
	resp, r, err := apiClient.RepositoryAPI.GetAutoDeclineSettings1(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetAutoDeclineSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoDeclineSettings1`: RestAutoDeclineSettings
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetAutoDeclineSettings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoDeclineSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestAutoDeclineSettings**](RestAutoDeclineSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranches

> GetBranches200Response GetBranches(ctx, projectKey, repositorySlug).BoostMatches(boostMatches).Context(context).OrderBy(orderBy).Details(details).FilterText(filterText).Base(base).Start(start).Limit(limit).Execute()

Find branches



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
	boostMatches := true // bool | Controls whether exact and prefix matches will be boosted to the top (optional)
	context := "context_example" // string |  (optional)
	orderBy := "orderBy_example" // string | Ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) (optional)
	details := true // bool | Whether to retrieve plugin-provided metadata about each branch (optional)
	filterText := "filterText_example" // string | The text to match on (optional)
	base := "base_example" // string | Base branch or tag to compare each branch to (for the metadata providers that uses that information (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetBranches(context.Background(), projectKey, repositorySlug).BoostMatches(boostMatches).Context(context).OrderBy(orderBy).Details(details).FilterText(filterText).Base(base).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBranches`: GetBranches200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **boostMatches** | **bool** | Controls whether exact and prefix matches will be boosted to the top | 
 **context** | **string** |  | 
 **orderBy** | **string** | Ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) | 
 **details** | **bool** | Whether to retrieve plugin-provided metadata about each branch | 
 **filterText** | **string** | The text to match on | 
 **base** | **string** | Base branch or tag to compare each branch to (for the metadata providers that uses that information | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetBranches200Response**](GetBranches200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChanges

> GetChanges1200Response GetChanges(ctx, projectKey, commitId, repositorySlug).WithComments(withComments).Since(since).Start(start).Limit(limit).Execute()

Get changes in commit



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
	commitId := "commitId_example" // string | The commit to retrieve changes for
	repositorySlug := "repositorySlug_example" // string | The repository slug
	withComments := "withComments_example" // string | <code>true</code> to apply comment counts in the changes (the default); otherwise, <code>false</code> to stream changes without comment counts (optional)
	since := "since_example" // string | The commit to which <code>until</code> should be compared to produce a page of changes. If not specified the commit's first parent is assumed (if one exists) (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetChanges(context.Background(), projectKey, commitId, repositorySlug).WithComments(withComments).Since(since).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChanges`: GetChanges1200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | The commit to retrieve changes for | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **withComments** | **string** | &lt;code&gt;true&lt;/code&gt; to apply comment counts in the changes (the default); otherwise, &lt;code&gt;false&lt;/code&gt; to stream changes without comment counts | 
 **since** | **string** | The commit to which &lt;code&gt;until&lt;/code&gt; should be compared to produce a page of changes. If not specified the commit&#39;s first parent is assumed (if one exists) | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetChanges1200Response**](GetChanges1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChanges1

> GetChanges1200Response GetChanges1(ctx, projectKey, repositorySlug).Until(until).Since(since).Start(start).Limit(limit).Execute()

Get changes made in commit



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
	until := "until_example" // string | The commit to retrieve changes for (optional)
	since := "since_example" // string | The commit to which <code>until</code> should be compared to produce a page of changes. If not specified the commit's first parent is assumed (if one exists) (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetChanges1(context.Background(), projectKey, repositorySlug).Until(until).Since(since).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetChanges1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChanges1`: GetChanges1200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetChanges1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChanges1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **until** | **string** | The commit to retrieve changes for | 
 **since** | **string** | The commit to which &lt;code&gt;until&lt;/code&gt; should be compared to produce a page of changes. If not specified the commit&#39;s first parent is assumed (if one exists) | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetChanges1200Response**](GetChanges1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComment

> RestComment GetComment(ctx, projectKey, commentId, commitId, repositorySlug).Execute()

Get a commit comment



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
	commentId := "commentId_example" // string | The ID of the comment to retrieve
	commitId := "commitId_example" // string | The <i>full ID</i> of the commit within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetComment(context.Background(), projectKey, commentId, commitId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComment`: RestComment
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commentId** | **string** | The ID of the comment to retrieve | 
**commitId** | **string** | The &lt;i&gt;full ID&lt;/i&gt; of the commit within the repository | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentRequest struct via the builder pattern


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


## GetComments

> GetComments200Response GetComments(ctx, projectKey, commitId, repositorySlug).Path(path).Since(since).Start(start).Limit(limit).Execute()

Search for commit comments



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
	commitId := "commitId_example" // string | The <i>full ID</i> of the commit within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug
	path := "path_example" // string | The path to the file on which comments were made (optional)
	since := "since_example" // string | For a merge commit, a parent can be provided to specify which diff the comments are on. For a commit range, a sinceId can be provided to specify where the comments are anchored from. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetComments(context.Background(), projectKey, commitId, repositorySlug).Path(path).Since(since).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComments`: GetComments200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | The &lt;i&gt;full ID&lt;/i&gt; of the commit within the repository | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **string** | The path to the file on which comments were made | 
 **since** | **string** | For a merge commit, a parent can be provided to specify which diff the comments are on. For a commit range, a sinceId can be provided to specify where the comments are anchored from. | 
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


## GetCommit

> RestCommit GetCommit(ctx, projectKey, commitId, repositorySlug).Path(path).Execute()

Get commit by ID



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
	commitId := "commitId_example" // string | The commit ID to retrieve
	repositorySlug := "repositorySlug_example" // string | The repository slug
	path := "path_example" // string | An optional path to filter the commit by. If supplied the details returned <i>may not</i> be for the specified commit. Instead, starting from the specified commit, they will be the details for the first commit affecting the specified path. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetCommit(context.Background(), projectKey, commitId, repositorySlug).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetCommit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommit`: RestCommit
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | The commit ID to retrieve | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **string** | An optional path to filter the commit by. If supplied the details returned &lt;i&gt;may not&lt;/i&gt; be for the specified commit. Instead, starting from the specified commit, they will be the details for the first commit affecting the specified path. | 

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


## GetCommits

> GetCommits200Response GetCommits(ctx, projectKey, repositorySlug).AvatarScheme(avatarScheme).Path(path).WithCounts(withCounts).FollowRenames(followRenames).Until(until).AvatarSize(avatarSize).Since(since).Merges(merges).IgnoreMissing(ignoreMissing).Start(start).Limit(limit).Execute()

Get commits



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
	avatarScheme := "avatarScheme_example" // string | The desired scheme for the avatar URL. If the parameter is not present URLs will use the same scheme as this request (optional)
	path := "path_example" // string | An optional path to filter commits by (optional)
	withCounts := "withCounts_example" // string | Optionally include the total number of commits and total number of unique authors (optional)
	followRenames := "followRenames_example" // string | If <code>true</code>, the commit history of the specified file will be followed past renames. Only valid for a path to a single file. (optional)
	until := "until_example" // string | The commit ID (SHA1) or ref (inclusively) to retrieve commits before (optional)
	avatarSize := "avatarSize_example" // string | If present the service adds avatar URLs for commit authors. Should be an integer specifying the desired size in pixels. If the parameter is not present, avatar URLs will not be set (optional)
	since := "since_example" // string | The commit ID or ref (exclusively) to retrieve commits after (optional)
	merges := "merges_example" // string | If present, controls how merge commits should be filtered. Can be either <code>exclude</code>, to exclude merge commits, <code>include</code>, to include both merge commits and non-merge commits or <code>only</code>, to only return merge commits. (optional)
	ignoreMissing := "ignoreMissing_example" // string | <code>true</code> to ignore missing commits, <code>false</code> otherwise (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetCommits(context.Background(), projectKey, repositorySlug).AvatarScheme(avatarScheme).Path(path).WithCounts(withCounts).FollowRenames(followRenames).Until(until).AvatarSize(avatarSize).Since(since).Merges(merges).IgnoreMissing(ignoreMissing).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommits`: GetCommits200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **avatarScheme** | **string** | The desired scheme for the avatar URL. If the parameter is not present URLs will use the same scheme as this request | 
 **path** | **string** | An optional path to filter commits by | 
 **withCounts** | **string** | Optionally include the total number of commits and total number of unique authors | 
 **followRenames** | **string** | If &lt;code&gt;true&lt;/code&gt;, the commit history of the specified file will be followed past renames. Only valid for a path to a single file. | 
 **until** | **string** | The commit ID (SHA1) or ref (inclusively) to retrieve commits before | 
 **avatarSize** | **string** | If present the service adds avatar URLs for commit authors. Should be an integer specifying the desired size in pixels. If the parameter is not present, avatar URLs will not be set | 
 **since** | **string** | The commit ID or ref (exclusively) to retrieve commits after | 
 **merges** | **string** | If present, controls how merge commits should be filtered. Can be either &lt;code&gt;exclude&lt;/code&gt;, to exclude merge commits, &lt;code&gt;include&lt;/code&gt;, to include both merge commits and non-merge commits or &lt;code&gt;only&lt;/code&gt;, to only return merge commits. | 
 **ignoreMissing** | **string** | &lt;code&gt;true&lt;/code&gt; to ignore missing commits, &lt;code&gt;false&lt;/code&gt; otherwise | 
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


## GetConfigurations1

> GetConfigurations200Response GetConfigurations1(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Execute()

Get hook scripts



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
	resp, r, err := apiClient.RepositoryAPI.GetConfigurations1(context.Background(), projectKey, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetConfigurations1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurations1`: GetConfigurations200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetConfigurations1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurations1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetConfigurations200Response**](GetConfigurations200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContent

> GetContent(ctx, projectKey, repositorySlug).NoContent(noContent).At(at).Size(size).Blame(blame).Type_(type_).Execute()

Get file content at revision



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
	noContent := "noContent_example" // string | If blame&amp;noContent only the blame is retrieved instead of the contents (optional)
	at := "at_example" // string | The commit ID or ref to retrieve the content for (optional)
	size := "size_example" // string | If true only the size will be returned for the file path instead of the contents (optional)
	blame := "blame_example" // string | If present and not equal to 'false', the blame will be returned for the file as well (optional)
	type_ := "type__example" // string | If true only the type will be returned for the file path instead of the contents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.GetContent(context.Background(), projectKey, repositorySlug).NoContent(noContent).At(at).Size(size).Blame(blame).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetContent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **noContent** | **string** | If blame&amp;amp;noContent only the blame is retrieved instead of the contents | 
 **at** | **string** | The commit ID or ref to retrieve the content for | 
 **size** | **string** | If true only the size will be returned for the file path instead of the contents | 
 **blame** | **string** | If present and not equal to &#39;false&#39;, the blame will be returned for the file as well | 
 **type_** | **string** | If true only the type will be returned for the file path instead of the contents | 

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


## GetContent1

> GetContent1(ctx, path, projectKey, repositorySlug).NoContent(noContent).At(at).Size(size).Blame(blame).Type_(type_).Execute()

Get file content



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
	path := "path_example" // string | The file path to retrieve content from
	projectKey := "projectKey_example" // string | The project key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	noContent := "noContent_example" // string | If blame&amp;noContent only the blame is retrieved instead of the contents (optional)
	at := "at_example" // string | The commit ID or ref to retrieve the content for (optional)
	size := "size_example" // string | If true only the size will be returned for the file path instead of the contents (optional)
	blame := "blame_example" // string | If present and not equal to 'false', the blame will be returned for the file as well (optional)
	type_ := "type__example" // string | If true only the type will be returned for the file path instead of the contents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.GetContent1(context.Background(), path, projectKey, repositorySlug).NoContent(noContent).At(at).Size(size).Blame(blame).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetContent1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The file path to retrieve content from | 
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContent1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **noContent** | **string** | If blame&amp;amp;noContent only the blame is retrieved instead of the contents | 
 **at** | **string** | The commit ID or ref to retrieve the content for | 
 **size** | **string** | If true only the size will be returned for the file path instead of the contents | 
 **blame** | **string** | If present and not equal to &#39;false&#39;, the blame will be returned for the file as well | 
 **type_** | **string** | If true only the type will be returned for the file path instead of the contents | 

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
	resp, r, err := apiClient.RepositoryAPI.GetDefaultBranch1(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetDefaultBranch1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultBranch1`: RestBranch
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetDefaultBranch1`: %v\n", resp)
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


## GetDefaultTasks1

> GetDefaultTasks1200Response GetDefaultTasks1(ctx, projectKey, repositorySlug).Markup(markup).Start(start).Limit(limit).Execute()

Get a page of default tasks



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
	markup := "markup_example" // string | If present or `\"true\"`, includes a markup-rendered description (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetDefaultTasks1(context.Background(), projectKey, repositorySlug).Markup(markup).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetDefaultTasks1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultTasks1`: GetDefaultTasks1200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetDefaultTasks1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultTasks1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **markup** | **string** | If present or &#x60;\&quot;true\&quot;&#x60;, includes a markup-rendered description | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetDefaultTasks1200Response**](GetDefaultTasks1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiffStatsSummary

> interface{} GetDiffStatsSummary(ctx, path, projectKey, commitId, repositorySlug).SrcPath(srcPath).AutoSrcPath(autoSrcPath).Whitespace(whitespace).Since(since).Execute()

Get diff stats summary between revisions



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
	projectKey := "projectKey_example" // string | The project key
	commitId := "commitId_example" // string | The commit ID to diff to.
	repositorySlug := "repositorySlug_example" // string | The repository slug
	srcPath := "srcPath_example" // string | The source path for the file, if it was copied, moved or renamed (optional)
	autoSrcPath := "autoSrcPath_example" // string | <code>true</code> to automatically try to find the source path when it's not provided, <code>false</code> otherwise. Requires the path to be provided. (optional)
	whitespace := "whitespace_example" // string | Optional whitespace flag which can be set to ignore-all (optional)
	since := "since_example" // string | The base revision to diff from. If omitted the parent revision of the commit ID is used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetDiffStatsSummary(context.Background(), path, projectKey, commitId, repositorySlug).SrcPath(srcPath).AutoSrcPath(autoSrcPath).Whitespace(whitespace).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetDiffStatsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiffStatsSummary`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetDiffStatsSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to the file which should be diffed (optional) | 
**projectKey** | **string** | The project key | 
**commitId** | **string** | The commit ID to diff to. | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiffStatsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **srcPath** | **string** | The source path for the file, if it was copied, moved or renamed | 
 **autoSrcPath** | **string** | &lt;code&gt;true&lt;/code&gt; to automatically try to find the source path when it&#39;s not provided, &lt;code&gt;false&lt;/code&gt; otherwise. Requires the path to be provided. | 
 **whitespace** | **string** | Optional whitespace flag which can be set to ignore-all | 
 **since** | **string** | The base revision to diff from. If omitted the parent revision of the commit ID is used | 

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


## GetDiffStatsSummary1

> RestDiff GetDiffStatsSummary1(ctx, path, projectKey, repositorySlug).FromRepo(fromRepo).SrcPath(srcPath).From(from).To(to).Whitespace(whitespace).Execute()

Retrieve the diff stats summary between commits



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
	path := "path_example" // string | the path to the file to diff (optional)
	projectKey := "projectKey_example" // string | The project key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	fromRepo := "fromRepo_example" // string | an optional parameter specifying the source repository containing the source commit if that commit is not present in the current repository; the repository can be specified by either its ID <em>fromRepo=42</em> or by its project key plus its repo slug separated by a slash: <em>fromRepo=projectKey/repoSlug</em> (optional)
	srcPath := "srcPath_example" // string | source path (optional)
	from := "from_example" // string | the source commit (can be a partial/full commit ID or qualified/unqualified ref name) (optional)
	to := "to_example" // string | the target commit (can be a partial/full commit ID or qualified/unqualified ref name) (optional)
	whitespace := "whitespace_example" // string | an optional whitespace flag which can be set to <code>ignore-all</code> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetDiffStatsSummary1(context.Background(), path, projectKey, repositorySlug).FromRepo(fromRepo).SrcPath(srcPath).From(from).To(to).Whitespace(whitespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetDiffStatsSummary1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiffStatsSummary1`: RestDiff
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetDiffStatsSummary1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | the path to the file to diff (optional) | 
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiffStatsSummary1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromRepo** | **string** | an optional parameter specifying the source repository containing the source commit if that commit is not present in the current repository; the repository can be specified by either its ID &lt;em&gt;fromRepo&#x3D;42&lt;/em&gt; or by its project key plus its repo slug separated by a slash: &lt;em&gt;fromRepo&#x3D;projectKey/repoSlug&lt;/em&gt; | 
 **srcPath** | **string** | source path | 
 **from** | **string** | the source commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **to** | **string** | the target commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **whitespace** | **string** | an optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 

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


## GetLatestInvocation1

> RestDetailedInvocation GetLatestInvocation1(ctx, projectKey, webhookId, repositorySlug).Event(event).Outcome(outcome).Execute()

Get last webhook invocation details



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
	webhookId := "webhookId_example" // string | ID of the webhook
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	event := "event_example" // string | The string ID of a specific event to retrieve the last invocation for. (optional)
	outcome := "outcome_example" // string | The outcome to filter for. Can be SUCCESS, FAILURE, ERROR. None specified means that the all will be considered (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetLatestInvocation1(context.Background(), projectKey, webhookId, repositorySlug).Event(event).Outcome(outcome).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetLatestInvocation1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestInvocation1`: RestDetailedInvocation
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetLatestInvocation1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | ID of the webhook | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestInvocation1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **event** | **string** | The string ID of a specific event to retrieve the last invocation for. | 
 **outcome** | **string** | The outcome to filter for. Can be SUCCESS, FAILURE, ERROR. None specified means that the all will be considered | 

### Return type

[**RestDetailedInvocation**](RestDetailedInvocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMergeBase

> RestCommit GetMergeBase(ctx, projectKey, commitId, repositorySlug).OtherCommitId(otherCommitId).Execute()

Get the common ancestor between two commits



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
	commitId := "commitId_example" // string | The <i>full ID</i> of the commit within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug
	otherCommitId := "otherCommitId_example" // string | The other commit id to calculate the merge-base on (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetMergeBase(context.Background(), projectKey, commitId, repositorySlug).OtherCommitId(otherCommitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetMergeBase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMergeBase`: RestCommit
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetMergeBase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | The &lt;i&gt;full ID&lt;/i&gt; of the commit within the repository | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMergeBaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **otherCommitId** | **string** | The other commit id to calculate the merge-base on | 

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


## GetPullRequestSettings1

> RestRepositoryPullRequestSettings GetPullRequestSettings1(ctx, projectKey, repositorySlug).Execute()

Get pull request settings



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
	resp, r, err := apiClient.RepositoryAPI.GetPullRequestSettings1(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetPullRequestSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPullRequestSettings1`: RestRepositoryPullRequestSettings
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetPullRequestSettings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestRepositoryPullRequestSettings**](RestRepositoryPullRequestSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRefChangeActivity

> GetRefChangeActivity200Response GetRefChangeActivity(ctx, projectKey, repositorySlug).Ref(ref).Start(start).Limit(limit).Execute()

Get ref change activity



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
	ref := "ref_example" // string | (optional) exact match for a ref ID to filter ref change activity for (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetRefChangeActivity(context.Background(), projectKey, repositorySlug).Ref(ref).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRefChangeActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRefChangeActivity`: GetRefChangeActivity200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRefChangeActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefChangeActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ref** | **string** | (optional) exact match for a ref ID to filter ref change activity for | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetRefChangeActivity200Response**](GetRefChangeActivity200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositories1

> GetRepositoriesRecentlyAccessed200Response GetRepositories1(ctx).Archived(archived).Projectname(projectname).Projectkey(projectkey).Visibility(visibility).Name(name).Permission(permission).State(state).Start(start).Limit(limit).Execute()

Search for repositories



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
	archived := "archived_example" // string | (optional) if specified, this will limit the resulting repository list to ones whose are <tt>ACTIVE</tt>, <tt>ARCHIVED</tt> or <tt>ALL</tt> for both. The match performed is case-insensitive. This filter defaults to <tt>ACTIVE</tt> when not set. <em>Available since 8.0</em> (optional)
	projectname := "projectname_example" // string | (optional) if specified, this will limit the resulting repository list to ones whose project's name matches this parameter's value. The match performed is case-insensitive and any leading and/or trailing whitespace characters on the <code>projectname</code> parameter will be stripped. (optional)
	projectkey := "projectkey_example" // string | (optional) if specified, this will limit the resulting repository list to ones whose project's key matches this parameter's value. The match performed is case-insensitive and any leading  and/or trailing whitespace characters on the <code>projectKey</code> parameter will be stripped. <em>Available since 8.0</em> (optional)
	visibility := "visibility_example" // string | (optional) if specified, this will limit the resulting repository list based on the repositories visibility. Valid values are <em>public</em> or <em>private</em>. (optional)
	name := "name_example" // string | (optional) if specified, this will limit the resulting repository list to ones whose name matches this parameter's value. The match performed is case-insensitive and any leading and/or trailing whitespace characters on the <code>name</code> parameter will be stripped. (optional)
	permission := "permission_example" // string | (optional) if specified, it must be a valid repository permission level name and will limit the resulting repository list to ones that the requesting user has the specified permission level to. If not specified, the default implicit 'read' permission level will be assumed. The currently supported explicit permission values are <tt>REPO_READ</tt>, <tt>REPO_WRITE</tt> and <tt>REPO_ADMIN</tt>. (optional)
	state := "state_example" // string | (optional) if specified, it must be a valid repository state name and will limit the resulting repository list to ones that are in the specified state. The currently supported explicit state values are <tt>AVAILABLE</tt>, <tt>INITIALISING</tt> and <tt>INITIALISATION_FAILED</tt>. Filtering by <tt>OFFLINE</tt> repositories is not supported.<br><em>Available since 5.13</em> (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetRepositories1(context.Background()).Archived(archived).Projectname(projectname).Projectkey(projectkey).Visibility(visibility).Name(name).Permission(permission).State(state).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRepositories1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositories1`: GetRepositoriesRecentlyAccessed200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRepositories1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositories1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **string** | (optional) if specified, this will limit the resulting repository list to ones whose are &lt;tt&gt;ACTIVE&lt;/tt&gt;, &lt;tt&gt;ARCHIVED&lt;/tt&gt; or &lt;tt&gt;ALL&lt;/tt&gt; for both. The match performed is case-insensitive. This filter defaults to &lt;tt&gt;ACTIVE&lt;/tt&gt; when not set. &lt;em&gt;Available since 8.0&lt;/em&gt; | 
 **projectname** | **string** | (optional) if specified, this will limit the resulting repository list to ones whose project&#39;s name matches this parameter&#39;s value. The match performed is case-insensitive and any leading and/or trailing whitespace characters on the &lt;code&gt;projectname&lt;/code&gt; parameter will be stripped. | 
 **projectkey** | **string** | (optional) if specified, this will limit the resulting repository list to ones whose project&#39;s key matches this parameter&#39;s value. The match performed is case-insensitive and any leading  and/or trailing whitespace characters on the &lt;code&gt;projectKey&lt;/code&gt; parameter will be stripped. &lt;em&gt;Available since 8.0&lt;/em&gt; | 
 **visibility** | **string** | (optional) if specified, this will limit the resulting repository list based on the repositories visibility. Valid values are &lt;em&gt;public&lt;/em&gt; or &lt;em&gt;private&lt;/em&gt;. | 
 **name** | **string** | (optional) if specified, this will limit the resulting repository list to ones whose name matches this parameter&#39;s value. The match performed is case-insensitive and any leading and/or trailing whitespace characters on the &lt;code&gt;name&lt;/code&gt; parameter will be stripped. | 
 **permission** | **string** | (optional) if specified, it must be a valid repository permission level name and will limit the resulting repository list to ones that the requesting user has the specified permission level to. If not specified, the default implicit &#39;read&#39; permission level will be assumed. The currently supported explicit permission values are &lt;tt&gt;REPO_READ&lt;/tt&gt;, &lt;tt&gt;REPO_WRITE&lt;/tt&gt; and &lt;tt&gt;REPO_ADMIN&lt;/tt&gt;. | 
 **state** | **string** | (optional) if specified, it must be a valid repository state name and will limit the resulting repository list to ones that are in the specified state. The currently supported explicit state values are &lt;tt&gt;AVAILABLE&lt;/tt&gt;, &lt;tt&gt;INITIALISING&lt;/tt&gt; and &lt;tt&gt;INITIALISATION_FAILED&lt;/tt&gt;. Filtering by &lt;tt&gt;OFFLINE&lt;/tt&gt; repositories is not supported.&lt;br&gt;&lt;em&gt;Available since 5.13&lt;/em&gt; | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetRepositoriesRecentlyAccessed200Response**](GetRepositoriesRecentlyAccessed200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoriesRecentlyAccessed

> GetRepositoriesRecentlyAccessed200Response GetRepositoriesRecentlyAccessed(ctx).Permission(permission).Start(start).Limit(limit).Execute()

Get recently accessed repositories



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
	permission := "permission_example" // string | (optional) If specified, it must be a valid repository permission level name and will limit the resulting repository list to ones that the requesting user has the specified permission level to. If not specified, the default <code>REPO_READ</code> permission level will be assumed. (optional) (default to "REPO_READ")
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetRepositoriesRecentlyAccessed(context.Background()).Permission(permission).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRepositoriesRecentlyAccessed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoriesRecentlyAccessed`: GetRepositoriesRecentlyAccessed200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRepositoriesRecentlyAccessed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRecentlyAccessedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | **string** | (optional) If specified, it must be a valid repository permission level name and will limit the resulting repository list to ones that the requesting user has the specified permission level to. If not specified, the default &lt;code&gt;REPO_READ&lt;/code&gt; permission level will be assumed. | [default to &quot;REPO_READ&quot;]
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetRepositoriesRecentlyAccessed200Response**](GetRepositoriesRecentlyAccessed200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryHook1

> RestRepositoryHook GetRepositoryHook1(ctx, projectKey, hookKey, repositorySlug).Execute()

Get repository hook



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
	hookKey := "hookKey_example" // string | The hook key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetRepositoryHook1(context.Background(), projectKey, hookKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRepositoryHook1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryHook1`: RestRepositoryHook
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRepositoryHook1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryHook1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestRepositoryHook**](RestRepositoryHook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryHooks1

> GetRepositoryHooks1200Response GetRepositoryHooks1(ctx, projectKey, repositorySlug).Type_(type_).Start(start).Limit(limit).Execute()

Get repository hooks



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
	type_ := "type__example" // string | The optional type to filter by. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetRepositoryHooks1(context.Background(), projectKey, repositorySlug).Type_(type_).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRepositoryHooks1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryHooks1`: GetRepositoryHooks1200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRepositoryHooks1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryHooks1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | The optional type to filter by. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetRepositoryHooks1200Response**](GetRepositoryHooks1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestriction1

> RestRefRestriction GetRestriction1(ctx, projectKey, id, repositorySlug).Execute()

Get a ref restriction



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
	id := "id_example" // string | The restriction id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetRestriction1(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRestriction1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestriction1`: RestRefRestriction
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRestriction1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The restriction id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestriction1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestRefRestriction**](RestRefRestriction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestrictions1

> GetRestrictions1200Response GetRestrictions1(ctx, projectKey, repositorySlug).MatcherType(matcherType).MatcherId(matcherId).Type_(type_).Start(start).Limit(limit).Execute()

Search for ref restrictions



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
	matcherType := "matcherType_example" // string | Matcher type to filter on (optional)
	matcherId := "matcherId_example" // string | Matcher id to filter on. Requires the matcherType parameter to be specified also. (optional)
	type_ := "type__example" // string | Types of restrictions to filter on. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetRestrictions1(context.Background(), projectKey, repositorySlug).MatcherType(matcherType).MatcherId(matcherId).Type_(type_).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRestrictions1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestrictions1`: GetRestrictions1200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRestrictions1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestrictions1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **matcherType** | **string** | Matcher type to filter on | 
 **matcherId** | **string** | Matcher id to filter on. Requires the matcherType parameter to be specified also. | 
 **type_** | **string** | Types of restrictions to filter on. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetRestrictions1200Response**](GetRestrictions1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettings1

> ExampleSettings GetSettings1(ctx, projectKey, hookKey, repositorySlug).Execute()

Get repository hook settings



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
	hookKey := "hookKey_example" // string | The hook key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetSettings1(context.Background(), projectKey, hookKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings1`: ExampleSettings
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetSettings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ExampleSettings**](ExampleSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatistics1

> interface{} GetStatistics1(ctx, projectKey, webhookId, repositorySlug).Event(event).Execute()

Get webhook statistics



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
	webhookId := "webhookId_example" // string | ID of the webhook
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	event := "event_example" // string | The string ID of a specific event to retrieve the last invocation for. May be empty, in which case all events are considered (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetStatistics1(context.Background(), projectKey, webhookId, repositorySlug).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetStatistics1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatistics1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetStatistics1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | ID of the webhook | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatistics1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **event** | **string** | The string ID of a specific event to retrieve the last invocation for. May be empty, in which case all events are considered | 

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


## GetStatisticsSummary1

> interface{} GetStatisticsSummary1(ctx, projectKey, webhookId, repositorySlug).Execute()

Get webhook statistics summary



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
	webhookId := "webhookId_example" // string | ID of the webhook
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetStatisticsSummary1(context.Background(), projectKey, webhookId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetStatisticsSummary1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatisticsSummary1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetStatisticsSummary1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | ID of the webhook | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsSummary1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## GetTag

> RestTag GetTag(ctx, projectKey, name, repositorySlug).Execute()

Get tag



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
	name := "name_example" // string | The name of the tag to be retrieved.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetTag(context.Background(), projectKey, name, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTag`: RestTag
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**name** | **string** | The name of the tag to be retrieved. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestTag**](RestTag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> GetTags200Response GetTags(ctx, projectKey, repositorySlug).OrderBy(orderBy).FilterText(filterText).Start(start).Limit(limit).Execute()

Find tag



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
	orderBy := "orderBy_example" // string | Ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) (optional)
	filterText := "filterText_example" // string | The text to match on. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetTags(context.Background(), projectKey, repositorySlug).OrderBy(orderBy).FilterText(filterText).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: GetTags200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orderBy** | **string** | Ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) | 
 **filterText** | **string** | The text to match on. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetTags200Response**](GetTags200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook1

> RestWebhook GetWebhook1(ctx, projectKey, webhookId, repositorySlug).Statistics(statistics).Execute()

Get webhook



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
	webhookId := "webhookId_example" // string | ID of the webhook
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	statistics := "statistics_example" // string | <code>true</code> if statistics should be provided for the webhook (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetWebhook1(context.Background(), projectKey, webhookId, repositorySlug).Statistics(statistics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetWebhook1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook1`: RestWebhook
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetWebhook1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | ID of the webhook | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhook1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **statistics** | **string** | &lt;code&gt;true&lt;/code&gt; if statistics should be provided for the webhook | 

### Return type

[**RestWebhook**](RestWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MGetStatus

> RestRefSyncStatus MGetStatus(ctx, projectKey, repositorySlug).At(at).Execute()

Get synchronization status



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
	at := "at_example" // string | Retrieves the synchronization status for the specified ref within the repository, rather than for the entire repository (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.MGetStatus(context.Background(), projectKey, repositorySlug).At(at).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.MGetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MGetStatus`: RestRefSyncStatus
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.MGetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **at** | **string** | Retrieves the synchronization status for the specified ref within the repository, rather than for the entire repository | 

### Return type

[**RestRefSyncStatus**](RestRefSyncStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## React

> RestUserReaction React(ctx, projectKey, commentId, commitId, emoticon, repositorySlug).Execute()

React to a comment



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
	commentId := "commentId_example" // string | The comment id
	commitId := "commitId_example" // string | The commit id
	emoticon := "emoticon_example" // string | The emoticon to add
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.React(context.Background(), projectKey, commentId, commitId, emoticon, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.React``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `React`: RestUserReaction
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.React`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The comment id | 
**commitId** | **string** | The commit id | 
**emoticon** | **string** | The emoticon to add | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactRequest struct via the builder pattern


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


## RemoveConfiguration1

> RemoveConfiguration1(ctx, projectKey, scriptId, repositorySlug).Execute()

Remove a hook script



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
	scriptId := "scriptId_example" // string | The ID of the hook script
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.RemoveConfiguration1(context.Background(), projectKey, scriptId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.RemoveConfiguration1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**scriptId** | **string** | The ID of the hook script | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveConfiguration1Request struct via the builder pattern


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


## RemoveLabel

> RemoveLabel(ctx, projectKey, labelName, repositorySlug).Execute()

Remove repository label



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
	labelName := "labelName_example" // string | The label to remove
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.RemoveLabel(context.Background(), projectKey, labelName, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.RemoveLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**labelName** | **string** | The label to remove | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLabelRequest struct via the builder pattern


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


## SaveAttachmentMetadata

> SaveAttachmentMetadata(ctx, projectKey, attachmentId, repositorySlug).Body(body).Execute()

Save attachment metadata



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
	attachmentId := "attachmentId_example" // string | the attachment ID
	repositorySlug := "repositorySlug_example" // string | The repository slug
	body := "body_example" // string | The attachment metadata can be any valid JSON content (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.SaveAttachmentMetadata(context.Background(), projectKey, attachmentId, repositorySlug).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.SaveAttachmentMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**attachmentId** | **string** | the attachment ID | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveAttachmentMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | The attachment metadata can be any valid JSON content | 

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


## SearchWebhooks

> SearchWebhooks(ctx, projectKey, repositorySlug).ScopeType(scopeType).Event(event).Statistics(statistics).Execute()

Search webhooks



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
	scopeType := "scopeType_example" // string | Scopes to filter by. This parameter can be specified once e.g. \"scopeType=repository\", or twice e.g. \"scopeType=repository&scopeType=project\", to filter by more than one scope level.  (optional)
	event := "event_example" // string | List of <code>com.atlassian.webhooks.WebhookEvent</code> ids to filter for (optional)
	statistics := true // bool | <code>true</code> if statistics should be provided for all found webhooks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.SearchWebhooks(context.Background(), projectKey, repositorySlug).ScopeType(scopeType).Event(event).Statistics(statistics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.SearchWebhooks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSearchWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scopeType** | **string** | Scopes to filter by. This parameter can be specified once e.g. \&quot;scopeType&#x3D;repository\&quot;, or twice e.g. \&quot;scopeType&#x3D;repository&amp;scopeType&#x3D;project\&quot;, to filter by more than one scope level.  | 
 **event** | **string** | List of &lt;code&gt;com.atlassian.webhooks.WebhookEvent&lt;/code&gt; ids to filter for | 
 **statistics** | **bool** | &lt;code&gt;true&lt;/code&gt; if statistics should be provided for all found webhooks | 

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


## Set1

> RestAutoMergeRestrictedSettings Set1(ctx, projectKey, repositorySlug).RestAutoMergeSettingsRequest(restAutoMergeSettingsRequest).Execute()

Create or update the pull request auto-merge settings



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
	restAutoMergeSettingsRequest := *openapiclient.NewRestAutoMergeSettingsRequest() // RestAutoMergeSettingsRequest | The settings to create or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.Set1(context.Background(), projectKey, repositorySlug).RestAutoMergeSettingsRequest(restAutoMergeSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Set1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Set1`: RestAutoMergeRestrictedSettings
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.Set1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiSet1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restAutoMergeSettingsRequest** | [**RestAutoMergeSettingsRequest**](RestAutoMergeSettingsRequest.md) | The settings to create or update | 

### Return type

[**RestAutoMergeRestrictedSettings**](RestAutoMergeRestrictedSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAutoDeclineSettings1

> RestAutoDeclineSettings SetAutoDeclineSettings1(ctx, projectKey, repositorySlug).RestAutoDeclineSettingsRequest(restAutoDeclineSettingsRequest).Execute()

Create auto decline settings



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
	restAutoDeclineSettingsRequest := *openapiclient.NewRestAutoDeclineSettingsRequest() // RestAutoDeclineSettingsRequest | The settings to create or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.SetAutoDeclineSettings1(context.Background(), projectKey, repositorySlug).RestAutoDeclineSettingsRequest(restAutoDeclineSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.SetAutoDeclineSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAutoDeclineSettings1`: RestAutoDeclineSettings
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.SetAutoDeclineSettings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAutoDeclineSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restAutoDeclineSettingsRequest** | [**RestAutoDeclineSettingsRequest**](RestAutoDeclineSettingsRequest.md) | The settings to create or update | 

### Return type

[**RestAutoDeclineSettings**](RestAutoDeclineSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration1

> RestHookScriptConfig SetConfiguration1(ctx, projectKey, scriptId, repositorySlug).RestHookScriptTriggers(restHookScriptTriggers).Execute()

Create/update a hook script



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
	scriptId := "scriptId_example" // string | The ID of the hook script
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restHookScriptTriggers := *openapiclient.NewRestHookScriptTriggers() // RestHookScriptTriggers | The hook triggers for which the hook script should be run (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.SetConfiguration1(context.Background(), projectKey, scriptId, repositorySlug).RestHookScriptTriggers(restHookScriptTriggers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.SetConfiguration1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetConfiguration1`: RestHookScriptConfig
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.SetConfiguration1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**scriptId** | **string** | The ID of the hook script | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restHookScriptTriggers** | [**RestHookScriptTriggers**](RestHookScriptTriggers.md) | The hook triggers for which the hook script should be run | 

### Return type

[**RestHookScriptConfig**](RestHookScriptConfig.md)

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
	r, err := apiClient.RepositoryAPI.SetDefaultBranch1(context.Background(), projectKey, repositorySlug).RestBranch(restBranch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.SetDefaultBranch1``: %v\n", err)
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


## SetEnabled

> RestRefSyncStatus SetEnabled(ctx, projectKey, repositorySlug).RestRefSyncStatus(restRefSyncStatus).Execute()

Disable synchronization



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
	restRefSyncStatus := *openapiclient.NewRestRefSyncStatus() // RestRefSyncStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.SetEnabled(context.Background(), projectKey, repositorySlug).RestRefSyncStatus(restRefSyncStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.SetEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetEnabled`: RestRefSyncStatus
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.SetEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restRefSyncStatus** | [**RestRefSyncStatus**](RestRefSyncStatus.md) |  | 

### Return type

[**RestRefSyncStatus**](RestRefSyncStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSettings1

> ExampleSettings SetSettings1(ctx, projectKey, hookKey, repositorySlug).ExampleSettings(exampleSettings).Execute()

Update repository hook settings



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
	hookKey := "hookKey_example" // string | The hook key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	exampleSettings := *openapiclient.NewExampleSettings() // ExampleSettings | The raw settings. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.SetSettings1(context.Background(), projectKey, hookKey, repositorySlug).ExampleSettings(exampleSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.SetSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSettings1`: ExampleSettings
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.SetSettings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **exampleSettings** | [**ExampleSettings**](ExampleSettings.md) | The raw settings. | 

### Return type

[**ExampleSettings**](ExampleSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Stream

> ExampleFiles Stream(ctx, projectKey, repositorySlug).At(at).Execute()

Stream files



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
	at := "at_example" // string | The commit to use as the starting point when listing files and calculating modifications (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.Stream(context.Background(), projectKey, repositorySlug).At(at).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Stream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Stream`: ExampleFiles
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.Stream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **at** | **string** | The commit to use as the starting point when listing files and calculating modifications | 

### Return type

[**ExampleFiles**](ExampleFiles.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Stream1

> ExampleFiles Stream1(ctx, path, projectKey, repositorySlug).At(at).Execute()

Stream files with last modified commit in path



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
	path := "path_example" // string | The path within the repository whose files should be streamed
	projectKey := "projectKey_example" // string | The project key
	repositorySlug := "repositorySlug_example" // string | The repository slug
	at := "at_example" // string | The commit to use as the starting point when listing files and calculating modifications (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.Stream1(context.Background(), path, projectKey, repositorySlug).At(at).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Stream1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Stream1`: ExampleFiles
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.Stream1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path within the repository whose files should be streamed | 
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiStream1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **at** | **string** | The commit to use as the starting point when listing files and calculating modifications | 

### Return type

[**ExampleFiles**](ExampleFiles.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamChanges

> GetChanges1200Response StreamChanges(ctx, projectKey, repositorySlug).FromRepo(fromRepo).From(from).To(to).Start(start).Limit(limit).Execute()

Compare commits



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
	fromRepo := "fromRepo_example" // string | an optional parameter specifying the source repository containing the source commit if that commit is not present in the current repository; the repository can be specified by either its ID <em>fromRepo=42</em> or by its project key plus its repo slug separated by a slash: <em>fromRepo=projectKey/repoSlug</em> (optional)
	from := "from_example" // string | the source commit (can be a partial/full commit ID or qualified/unqualified ref name) (optional)
	to := "to_example" // string | the target commit (can be a partial/full commit ID or qualified/unqualified ref name) (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.StreamChanges(context.Background(), projectKey, repositorySlug).FromRepo(fromRepo).From(from).To(to).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamChanges`: GetChanges1200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.StreamChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromRepo** | **string** | an optional parameter specifying the source repository containing the source commit if that commit is not present in the current repository; the repository can be specified by either its ID &lt;em&gt;fromRepo&#x3D;42&lt;/em&gt; or by its project key plus its repo slug separated by a slash: &lt;em&gt;fromRepo&#x3D;projectKey/repoSlug&lt;/em&gt; | 
 **from** | **string** | the source commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **to** | **string** | the target commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetChanges1200Response**](GetChanges1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamCommits

> GetCommits200Response StreamCommits(ctx, projectKey, repositorySlug).FromRepo(fromRepo).From(from).To(to).Start(start).Limit(limit).Execute()

Get accessible commits



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
	fromRepo := "fromRepo_example" // string | an optional parameter specifying the source repository containing the source commit if that commit is not present in the current repository; the repository can be specified by either its ID <em>fromRepo=42</em> or by its project key plus its repo slug separated by a slash: <em>fromRepo=projectKey/repoSlug</em> (optional)
	from := "from_example" // string | the source commit (can be a partial/full commit ID or qualified/unqualified ref name) (optional)
	to := "to_example" // string | the target commit (can be a partial/full commit ID or qualified/unqualified ref name) (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.StreamCommits(context.Background(), projectKey, repositorySlug).FromRepo(fromRepo).From(from).To(to).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamCommits`: GetCommits200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.StreamCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromRepo** | **string** | an optional parameter specifying the source repository containing the source commit if that commit is not present in the current repository; the repository can be specified by either its ID &lt;em&gt;fromRepo&#x3D;42&lt;/em&gt; or by its project key plus its repo slug separated by a slash: &lt;em&gt;fromRepo&#x3D;projectKey/repoSlug&lt;/em&gt; | 
 **from** | **string** | the source commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **to** | **string** | the target commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
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


## StreamDiff

> RestDiff StreamDiff(ctx, commitId, repositorySlug, path, projectKey).SrcPath(srcPath).AvatarSize(avatarSize).Filter(filter).AvatarScheme(avatarScheme).ContextLines(contextLines).AutoSrcPath(autoSrcPath).Whitespace(whitespace).WithComments(withComments).Since(since).Execute()

Get diff between revisions



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
	commitId := "commitId_example" // string | The <i>full ID</i> of the commit within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug
	path := "path_example" // string | The path to the file which should be diffed (optional)
	projectKey := "projectKey_example" // string | The project key
	srcPath := "srcPath_example" // string | The source path for the file, if it was copied, moved or renamed (optional)
	avatarSize := "avatarSize_example" // string | If present the service adds avatar URLs for comment authors where the provided value specifies the desired avatar size in pixels. Not applicable if streaming raw diff (optional)
	filter := "filter_example" // string | Text used to filter files and lines (optional). Not applicable if streaming raw diff (optional)
	avatarScheme := "avatarScheme_example" // string | The security scheme for avatar URLs. If the scheme is not present then it is inherited from the request. It can be set to \"https\" to force the use of secure URLs. Not applicable if streaming raw diff (optional)
	contextLines := "contextLines_example" // string | The number of context lines to include around added/removed lines in the diff.Not applicable if streaming raw diff (optional)
	autoSrcPath := "autoSrcPath_example" // string | <code>true</code> to automatically try to find the source path when it's not provided, <code>false</code> otherwise. Requires the path to be provided. (optional)
	whitespace := "whitespace_example" // string | Optional whitespace flag which can be set to ignore-all (optional)
	withComments := "withComments_example" // string | <code>true</code> to embed comments in the diff (the default); otherwise <code>false</code> to stream the diff without comments. Not applicable if streaming raw diff (optional)
	since := "since_example" // string | The base revision to diff from. If omitted the parent revision of the until revision is used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.StreamDiff(context.Background(), commitId, repositorySlug, path, projectKey).SrcPath(srcPath).AvatarSize(avatarSize).Filter(filter).AvatarScheme(avatarScheme).ContextLines(contextLines).AutoSrcPath(autoSrcPath).Whitespace(whitespace).WithComments(withComments).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamDiff`: RestDiff
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.StreamDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitId** | **string** | The &lt;i&gt;full ID&lt;/i&gt; of the commit within the repository | 
**repositorySlug** | **string** | The repository slug | 
**path** | **string** | The path to the file which should be diffed (optional) | 
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **srcPath** | **string** | The source path for the file, if it was copied, moved or renamed | 
 **avatarSize** | **string** | If present the service adds avatar URLs for comment authors where the provided value specifies the desired avatar size in pixels. Not applicable if streaming raw diff | 
 **filter** | **string** | Text used to filter files and lines (optional). Not applicable if streaming raw diff | 
 **avatarScheme** | **string** | The security scheme for avatar URLs. If the scheme is not present then it is inherited from the request. It can be set to \&quot;https\&quot; to force the use of secure URLs. Not applicable if streaming raw diff | 
 **contextLines** | **string** | The number of context lines to include around added/removed lines in the diff.Not applicable if streaming raw diff | 
 **autoSrcPath** | **string** | &lt;code&gt;true&lt;/code&gt; to automatically try to find the source path when it&#39;s not provided, &lt;code&gt;false&lt;/code&gt; otherwise. Requires the path to be provided. | 
 **whitespace** | **string** | Optional whitespace flag which can be set to ignore-all | 
 **withComments** | **string** | &lt;code&gt;true&lt;/code&gt; to embed comments in the diff (the default); otherwise &lt;code&gt;false&lt;/code&gt; to stream the diff without comments. Not applicable if streaming raw diff | 
 **since** | **string** | The base revision to diff from. If omitted the parent revision of the until revision is used | 

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


## StreamDiff1

> RestDiff StreamDiff1(ctx, path, projectKey, repositorySlug).ContextLines(contextLines).FromRepo(fromRepo).SrcPath(srcPath).From(from).To(to).Whitespace(whitespace).Execute()

Get diff between commits



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
	path := "path_example" // string | the path to the file to diff (optional)
	projectKey := "projectKey_example" // string | The project key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	contextLines := "contextLines_example" // string | an optional number of context lines to include around each added or removed lines in the diff (optional)
	fromRepo := "fromRepo_example" // string | an optional parameter specifying the source repository containing the source commit if that commit is not present in the current repository; the repository can be specified by either its ID <em>fromRepo=42</em> or by its project key plus its repo slug separated by a slash: <em>fromRepo=projectKey/repoSlug</em> (optional)
	srcPath := "srcPath_example" // string | source path (optional)
	from := "from_example" // string | the source commit (can be a partial/full commit ID or qualified/unqualified ref name) (optional)
	to := "to_example" // string | the target commit (can be a partial/full commit ID or qualified/unqualified ref name) (optional)
	whitespace := "whitespace_example" // string | an optional whitespace flag which can be set to <code>ignore-all</code> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.StreamDiff1(context.Background(), path, projectKey, repositorySlug).ContextLines(contextLines).FromRepo(fromRepo).SrcPath(srcPath).From(from).To(to).Whitespace(whitespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamDiff1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamDiff1`: RestDiff
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.StreamDiff1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | the path to the file to diff (optional) | 
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamDiff1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contextLines** | **string** | an optional number of context lines to include around each added or removed lines in the diff | 
 **fromRepo** | **string** | an optional parameter specifying the source repository containing the source commit if that commit is not present in the current repository; the repository can be specified by either its ID &lt;em&gt;fromRepo&#x3D;42&lt;/em&gt; or by its project key plus its repo slug separated by a slash: &lt;em&gt;fromRepo&#x3D;projectKey/repoSlug&lt;/em&gt; | 
 **srcPath** | **string** | source path | 
 **from** | **string** | the source commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **to** | **string** | the target commit (can be a partial/full commit ID or qualified/unqualified ref name) | 
 **whitespace** | **string** | an optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 

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


## StreamFiles

> StreamFiles200Response StreamFiles(ctx, projectKey, repositorySlug).At(at).Start(start).Limit(limit).Execute()

Get files in directory



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
	at := "at_example" // string | The commit ID or ref (e.g. a branch or tag) to list the files at. If not specified the default branch will be used instead. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.StreamFiles(context.Background(), projectKey, repositorySlug).At(at).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamFiles`: StreamFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.StreamFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **at** | **string** | The commit ID or ref (e.g. a branch or tag) to list the files at. If not specified the default branch will be used instead. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**StreamFiles200Response**](StreamFiles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamFiles1

> StreamFiles200Response StreamFiles1(ctx, path, projectKey, repositorySlug).At(at).Start(start).Limit(limit).Execute()

Get files in directory



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
	path := "path_example" // string | The directory to list files for.
	projectKey := "projectKey_example" // string | The project key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	at := "at_example" // string | The commit ID or ref (e.g. a branch or tag) to list the files at. If not specified the default branch will be used instead. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.StreamFiles1(context.Background(), path, projectKey, repositorySlug).At(at).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamFiles1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamFiles1`: StreamFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.StreamFiles1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The directory to list files for. | 
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamFiles1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **at** | **string** | The commit ID or ref (e.g. a branch or tag) to list the files at. If not specified the default branch will be used instead. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**StreamFiles200Response**](StreamFiles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamPatch

> StreamPatch(ctx, projectKey, repositorySlug).Until(until).AllAncestors(allAncestors).Since(since).Execute()

Get patch content at revision



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
	until := "until_example" // string | The target revision from which to generate the patch (required) (optional)
	allAncestors := "allAncestors_example" // string | indicates whether or not to generate a patch which includes all the ancestors of the 'until' revision. If true, the value provided by 'since' is ignored. (optional)
	since := "since_example" // string | The base revision from which to generate the patch. This is only applicable when 'allAncestors' is false. If omitted the patch will represent one single commit, the 'until'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.StreamPatch(context.Background(), projectKey, repositorySlug).Until(until).AllAncestors(allAncestors).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamPatch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStreamPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **until** | **string** | The target revision from which to generate the patch (required) | 
 **allAncestors** | **string** | indicates whether or not to generate a patch which includes all the ancestors of the &#39;until&#39; revision. If true, the value provided by &#39;since&#39; is ignored. | 
 **since** | **string** | The base revision from which to generate the patch. This is only applicable when &#39;allAncestors&#39; is false. If omitted the patch will represent one single commit, the &#39;until&#39;. | 

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


## StreamRaw

> StreamRaw(ctx, path, projectKey, repositorySlug).At(at).Markup(markup).HtmlEscape(htmlEscape).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Execute()

Get raw content of a file at revision



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
	path := "path_example" // string | The file path to retrieve content from
	projectKey := "projectKey_example" // string | The project key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	at := "at_example" // string | A specific commit or ref to retrieve the raw content at, or the default branch if not specified (optional)
	markup := "markup_example" // string | If present or \"true\", triggers the raw content to be markup-rendered and returned as HTML; otherwise, if not specified, or any value other than \"true\", the content is streamed without markup (optional)
	htmlEscape := "htmlEscape_example" // string | (Optional) true if HTML should be escaped in the input markup, false otherwise. If not specified, the value of the markup.render.html.escape property, which is true by default, will be used (optional)
	includeHeadingId := "includeHeadingId_example" // string | (Optional) true if headings should contain an ID based on the heading content. If not specified, the value of the markup.render.headerids property, which is false by default, will be used (optional)
	hardwrap := "hardwrap_example" // string | (Optional) Whether the markup implementation should convert newlines to breaks. If not specified, the value of the markup.render.hardwrap property, which is true by default, will be used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.StreamRaw(context.Background(), path, projectKey, repositorySlug).At(at).Markup(markup).HtmlEscape(htmlEscape).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamRaw``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The file path to retrieve content from | 
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **at** | **string** | A specific commit or ref to retrieve the raw content at, or the default branch if not specified | 
 **markup** | **string** | If present or \&quot;true\&quot;, triggers the raw content to be markup-rendered and returned as HTML; otherwise, if not specified, or any value other than \&quot;true\&quot;, the content is streamed without markup | 
 **htmlEscape** | **string** | (Optional) true if HTML should be escaped in the input markup, false otherwise. If not specified, the value of the markup.render.html.escape property, which is true by default, will be used | 
 **includeHeadingId** | **string** | (Optional) true if headings should contain an ID based on the heading content. If not specified, the value of the markup.render.headerids property, which is false by default, will be used | 
 **hardwrap** | **string** | (Optional) Whether the markup implementation should convert newlines to breaks. If not specified, the value of the markup.render.hardwrap property, which is true by default, will be used | 

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


## StreamRawDiff

> StreamRawDiff(ctx, projectKey, repositorySlug).ContextLines(contextLines).SrcPath(srcPath).Until(until).Whitespace(whitespace).Since(since).Execute()

Get raw diff for path



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
	contextLines := "contextLines_example" // string | The number of context lines to include around added/removed lines in the diff (optional)
	srcPath := "srcPath_example" // string | The source path for the file, if it was copied, moved or renamed (optional)
	until := "until_example" // string | The target revision to diff to (required) (optional)
	whitespace := "whitespace_example" // string | Optional whitespace flag which can be set to <code>ignore-all</code> (optional)
	since := "since_example" // string | The base revision to diff from. If omitted the parent revision of the until revision is used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.StreamRawDiff(context.Background(), projectKey, repositorySlug).ContextLines(contextLines).SrcPath(srcPath).Until(until).Whitespace(whitespace).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamRawDiff``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStreamRawDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contextLines** | **string** | The number of context lines to include around added/removed lines in the diff | 
 **srcPath** | **string** | The source path for the file, if it was copied, moved or renamed | 
 **until** | **string** | The target revision to diff to (required) | 
 **whitespace** | **string** | Optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 
 **since** | **string** | The base revision to diff from. If omitted the parent revision of the until revision is used | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; qs=0.1, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamRawDiff1

> StreamRawDiff1(ctx, path, projectKey, repositorySlug).ContextLines(contextLines).SrcPath(srcPath).Until(until).Whitespace(whitespace).Since(since).Execute()

Get raw diff for path



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
	path := "path_example" // string | The path to the file which should be diffed (required)
	projectKey := "projectKey_example" // string | The project key.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	contextLines := "contextLines_example" // string | The number of context lines to include around added/removed lines in the diff (optional)
	srcPath := "srcPath_example" // string | The source path for the file, if it was copied, moved or renamed (optional)
	until := "until_example" // string | The target revision to diff to (required) (optional)
	whitespace := "whitespace_example" // string | Optional whitespace flag which can be set to <code>ignore-all</code> (optional)
	since := "since_example" // string | The base revision to diff from. If omitted the parent revision of the until revision is used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.StreamRawDiff1(context.Background(), path, projectKey, repositorySlug).ContextLines(contextLines).SrcPath(srcPath).Until(until).Whitespace(whitespace).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.StreamRawDiff1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** | The path to the file which should be diffed (required) | 
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamRawDiff1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contextLines** | **string** | The number of context lines to include around added/removed lines in the diff | 
 **srcPath** | **string** | The source path for the file, if it was copied, moved or renamed | 
 **until** | **string** | The target revision to diff to (required) | 
 **whitespace** | **string** | Optional whitespace flag which can be set to &lt;code&gt;ignore-all&lt;/code&gt; | 
 **since** | **string** | The base revision to diff from. If omitted the parent revision of the until revision is used | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; qs=0.1, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Synchronize

> RestRejectedRef Synchronize(ctx, projectKey, repositorySlug).RestRefSyncRequest(restRefSyncRequest).Execute()

Manual synchronization



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
	restRefSyncRequest := *openapiclient.NewRestRefSyncRequest() // RestRefSyncRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.Synchronize(context.Background(), projectKey, repositorySlug).RestRefSyncRequest(restRefSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Synchronize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Synchronize`: RestRejectedRef
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.Synchronize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSynchronizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restRefSyncRequest** | [**RestRefSyncRequest**](RestRefSyncRequest.md) |  | 

### Return type

[**RestRejectedRef**](RestRejectedRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWebhook1

> interface{} TestWebhook1(ctx, projectKey, repositorySlug).WebhookId(webhookId).SslVerificationRequired(sslVerificationRequired).Url(url).RestWebhookCredentials(restWebhookCredentials).Execute()

Test webhook



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
	webhookId := int32(56) // int32 |  (optional)
	sslVerificationRequired := "sslVerificationRequired_example" // string | Whether SSL verification is required for the specified webhook URL. Default value is  <code>true</code>. (optional)
	url := "url_example" // string | The url in which to connect to (optional)
	restWebhookCredentials := *openapiclient.NewRestWebhookCredentials() // RestWebhookCredentials | Basic authentication credentials, if required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.TestWebhook1(context.Background(), projectKey, repositorySlug).WebhookId(webhookId).SslVerificationRequired(sslVerificationRequired).Url(url).RestWebhookCredentials(restWebhookCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.TestWebhook1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestWebhook1`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.TestWebhook1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestWebhook1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **webhookId** | **int32** |  | 
 **sslVerificationRequired** | **string** | Whether SSL verification is required for the specified webhook URL. Default value is  &lt;code&gt;true&lt;/code&gt;. | 
 **url** | **string** | The url in which to connect to | 
 **restWebhookCredentials** | [**RestWebhookCredentials**](RestWebhookCredentials.md) | Basic authentication credentials, if required. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnReact

> UnReact(ctx, projectKey, commentId, commitId, emoticon, repositorySlug).Execute()

Remove a reaction from comment



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
	commentId := "commentId_example" // string | The comment id
	commitId := "commitId_example" // string | The commit id
	emoticon := "emoticon_example" // string | The emoticon to remove
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.UnReact(context.Background(), projectKey, commentId, commitId, emoticon, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.UnReact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commentId** | **string** | The comment id | 
**commitId** | **string** | The commit id | 
**emoticon** | **string** | The emoticon to remove | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnReactRequest struct via the builder pattern


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


## Unwatch

> Unwatch(ctx, projectKey, commitId, repositorySlug).Execute()

Stop watching commit



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
	commitId := "commitId_example" // string | The <i>full ID</i> of the commit within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.Unwatch(context.Background(), projectKey, commitId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Unwatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | The &lt;i&gt;full ID&lt;/i&gt; of the commit within the repository | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnwatchRequest struct via the builder pattern


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


## Unwatch2

> Unwatch2(ctx, projectKey, repositorySlug).Execute()

Stop watching repository



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
	r, err := apiClient.RepositoryAPI.Unwatch2(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Unwatch2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnwatch2Request struct via the builder pattern


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


## UpdateComment

> RestComment UpdateComment(ctx, projectKey, commentId, commitId, repositorySlug).RestComment(restComment).Execute()

Update a commit comment



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
	commentId := "commentId_example" // string | The ID of the comment to retrieve
	commitId := "commitId_example" // string | The <i>full ID</i> of the commit within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug
	restComment := *openapiclient.NewRestComment() // RestComment | The comment to update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.UpdateComment(context.Background(), projectKey, commentId, commitId, repositorySlug).RestComment(restComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.UpdateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComment`: RestComment
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.UpdateComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commentId** | **string** | The ID of the comment to retrieve | 
**commitId** | **string** | The &lt;i&gt;full ID&lt;/i&gt; of the commit within the repository | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **restComment** | [**RestComment**](RestComment.md) | The comment to update | 

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


## UpdateDefaultTask1

> RestDefaultTask UpdateDefaultTask1(ctx, projectKey, repositorySlug, taskId).RestDefaultTaskRequest(restDefaultTaskRequest).Execute()

Update a default task



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
	taskId := "taskId_example" // string | The ID of the default task
	restDefaultTaskRequest := *openapiclient.NewRestDefaultTaskRequest() // RestDefaultTaskRequest | The task to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.UpdateDefaultTask1(context.Background(), projectKey, repositorySlug, taskId).RestDefaultTaskRequest(restDefaultTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.UpdateDefaultTask1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDefaultTask1`: RestDefaultTask
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.UpdateDefaultTask1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 
**taskId** | **string** | The ID of the default task | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultTask1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restDefaultTaskRequest** | [**RestDefaultTaskRequest**](RestDefaultTaskRequest.md) | The task to be updated | 

### Return type

[**RestDefaultTask**](RestDefaultTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePullRequestSettings1

> RestRepositoryPullRequestSettings UpdatePullRequestSettings1(ctx, projectKey, repositorySlug).RestRepositoryPullRequestSettings(restRepositoryPullRequestSettings).Execute()

Update pull request settings



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
	restRepositoryPullRequestSettings := *openapiclient.NewRestRepositoryPullRequestSettings() // RestRepositoryPullRequestSettings | The updated settings. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.UpdatePullRequestSettings1(context.Background(), projectKey, repositorySlug).RestRepositoryPullRequestSettings(restRepositoryPullRequestSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.UpdatePullRequestSettings1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePullRequestSettings1`: RestRepositoryPullRequestSettings
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.UpdatePullRequestSettings1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePullRequestSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restRepositoryPullRequestSettings** | [**RestRepositoryPullRequestSettings**](RestRepositoryPullRequestSettings.md) | The updated settings. | 

### Return type

[**RestRepositoryPullRequestSettings**](RestRepositoryPullRequestSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook1

> RestWebhook UpdateWebhook1(ctx, projectKey, webhookId, repositorySlug).RestWebhook(restWebhook).Execute()

Update webhook



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
	webhookId := "webhookId_example" // string | Id of the existing webhook
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restWebhook := *openapiclient.NewRestWebhook() // RestWebhook | The representation of the updated values for the webhook (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.UpdateWebhook1(context.Background(), projectKey, webhookId, repositorySlug).RestWebhook(restWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.UpdateWebhook1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook1`: RestWebhook
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.UpdateWebhook1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | Id of the existing webhook | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhook1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restWebhook** | [**RestWebhook**](RestWebhook.md) | The representation of the updated values for the webhook | 

### Return type

[**RestWebhook**](RestWebhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Watch

> Watch(ctx, projectKey, commitId, repositorySlug).Execute()

Watch commit



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
	commitId := "commitId_example" // string | The <i>full ID</i> of the commit within the repository
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.Watch(context.Background(), projectKey, commitId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Watch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | The &lt;i&gt;full ID&lt;/i&gt; of the commit within the repository | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchRequest struct via the builder pattern


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


## Watch2

> Watch2(ctx, projectKey, repositorySlug).RestRepository(restRepository).Execute()

Watch repository



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
	restRepository := *openapiclient.NewRestRepository() // RestRepository | The repository to watch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RepositoryAPI.Watch2(context.Background(), projectKey, repositorySlug).RestRepository(restRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.Watch2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiWatch2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restRepository** | [**RestRepository**](RestRepository.md) | The repository to watch. | 

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


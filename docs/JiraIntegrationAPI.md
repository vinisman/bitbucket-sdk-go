# \JiraIntegrationAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIssue**](JiraIntegrationAPI.md#CreateIssue) | **Post** /jira/latest/comments/{commentId}/issues | Create Jira Issue
[**GetCommitsByIssueKey**](JiraIntegrationAPI.md#GetCommitsByIssueKey) | **Get** /jira/latest/issues/{issueKey}/commits | Get changesets for issue key
[**GetEnhancedEntityLinkForProject**](JiraIntegrationAPI.md#GetEnhancedEntityLinkForProject) | **Get** /jira/latest/projects/{projectKey}/primary-enhanced-entitylink | Get entity link
[**GetIssueKeysForPullRequest**](JiraIntegrationAPI.md#GetIssueKeysForPullRequest) | **Get** /jira/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/issues | Get issues for a pull request



## CreateIssue

> RestCommentJiraIssue CreateIssue(ctx, commentId).ApplicationId(applicationId).Body(body).Execute()

Create Jira Issue



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
	commentId := "commentId_example" // string | the comment to associate the created Jira issue to
	applicationId := "applicationId_example" // string | id of the Jira server (optional)
	body := "body_example" // string | A String representation of the JSON format Jira create issue request see: <a href=\"https://docs.atlassian.com/jira/REST/server/#api/2/issue-createIssue\">Jira REST API</a> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraIntegrationAPI.CreateIssue(context.Background(), commentId).ApplicationId(applicationId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationAPI.CreateIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssue`: RestCommentJiraIssue
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationAPI.CreateIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | the comment to associate the created Jira issue to | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationId** | **string** | id of the Jira server | 
 **body** | **string** | A String representation of the JSON format Jira create issue request see: &lt;a href&#x3D;\&quot;https://docs.atlassian.com/jira/REST/server/#api/2/issue-createIssue\&quot;&gt;Jira REST API&lt;/a&gt; | 

### Return type

[**RestCommentJiraIssue**](RestCommentJiraIssue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommitsByIssueKey

> GetCommitsByIssueKey200Response GetCommitsByIssueKey(ctx, issueKey).MaxChanges(maxChanges).Start(start).Limit(limit).Execute()

Get changesets for issue key



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
	issueKey := "issueKey_example" // string | The issue key to search by
	maxChanges := "maxChanges_example" // string | The maximum number of changes to retrieve for each changeset (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraIntegrationAPI.GetCommitsByIssueKey(context.Background(), issueKey).MaxChanges(maxChanges).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationAPI.GetCommitsByIssueKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommitsByIssueKey`: GetCommitsByIssueKey200Response
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationAPI.GetCommitsByIssueKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueKey** | **string** | The issue key to search by | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitsByIssueKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxChanges** | **string** | The maximum number of changes to retrieve for each changeset | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetCommitsByIssueKey200Response**](GetCommitsByIssueKey200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnhancedEntityLinkForProject

> RestEnhancedEntityLink GetEnhancedEntityLinkForProject(ctx, projectKey).Execute()

Get entity link



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraIntegrationAPI.GetEnhancedEntityLinkForProject(context.Background(), projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationAPI.GetEnhancedEntityLinkForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnhancedEntityLinkForProject`: RestEnhancedEntityLink
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationAPI.GetEnhancedEntityLinkForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnhancedEntityLinkForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestEnhancedEntityLink**](RestEnhancedEntityLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueKeysForPullRequest

> []RestJiraIssue GetIssueKeysForPullRequest(ctx, projectKey, pullRequestId, repositorySlug).Execute()

Get issues for a pull request



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
	pullRequestId := "pullRequestId_example" // string | The pull request id
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraIntegrationAPI.GetIssueKeysForPullRequest(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationAPI.GetIssueKeysForPullRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueKeysForPullRequest`: []RestJiraIssue
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationAPI.GetIssueKeysForPullRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**pullRequestId** | **string** | The pull request id | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueKeysForPullRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]RestJiraIssue**](RestJiraIssue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


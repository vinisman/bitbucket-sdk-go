# \DashboardAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPullRequestSuggestions**](DashboardAPI.md#GetPullRequestSuggestions) | **Get** /api/latest/dashboard/pull-request-suggestions | Get pull request suggestions
[**GetPullRequests1**](DashboardAPI.md#GetPullRequests1) | **Get** /api/latest/dashboard/pull-requests | Get pull requests for a user



## GetPullRequestSuggestions

> GetPullRequestSuggestions200Response GetPullRequestSuggestions(ctx).ChangesSince(changesSince).Limit(limit).Execute()

Get pull request suggestions



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
	changesSince := "changesSince_example" // string | restrict pull request suggestions to be based on events that occurred since some timein the past. This is expressed in seconds since \"now\". So to return suggestionsbased only on activity within the past 48 hours, pass a value of 172800. (optional)
	limit := "limit_example" // string | restricts the result set to return at most this many suggestions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetPullRequestSuggestions(context.Background()).ChangesSince(changesSince).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetPullRequestSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPullRequestSuggestions`: GetPullRequestSuggestions200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetPullRequestSuggestions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changesSince** | **string** | restrict pull request suggestions to be based on events that occurred since some timein the past. This is expressed in seconds since \&quot;now\&quot;. So to return suggestionsbased only on activity within the past 48 hours, pass a value of 172800. | 
 **limit** | **string** | restricts the result set to return at most this many suggestions. | 

### Return type

[**GetPullRequestSuggestions200Response**](GetPullRequestSuggestions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequests1

> GetPullRequests1200Response GetPullRequests1(ctx).ClosedSince(closedSince).Role(role).ParticipantStatus(participantStatus).State(state).User(user).Order(order).Start(start).Limit(limit).Execute()

Get pull requests for a user



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
	closedSince := "closedSince_example" // string | (optional, defaults to returning pull requests regardless of closed since date). Permits returning only pull requests with a closed timestamp set more recently that (now - closedSince). Units are in seconds. So for example if closed since 86400 is set only pull requests closed in the previous 24 hours will be returned. (optional)
	role := "role_example" // string | (optional, defaults to returning pull requests for any role). If a role is supplied only pull requests where the authenticated user is a participant in the given role will be returned. Either <strong>REVIEWER</strong>, <strong>AUTHOR</strong> or <strong>PARTICIPANT</strong>. (optional)
	participantStatus := "participantStatus_example" // string | (optional, defaults to returning pull requests with any participant status). A comma separated list of participant status. That is, one or more of <strong>UNAPPROVED</strong>, <strong>NEEDS_WORK</strong>, or <strong>APPROVED</strong>. (optional)
	state := "state_example" // string | (optional, defaults to returning pull requests in any state). If a state is supplied only pull requests in the specified state will be returned. Either <strong>OPEN</strong>, <strong>DECLINED</strong> or <strong>MERGED</strong>. Omit this parameter to return pull request in any state. (optional)
	user := "user_example" // string | The name of the involved user, defaults to the current user. (optional)
	order := "order_example" // string | (optional, defaults to <strong>NEWEST</strong>) the order/(s) to return pull requests in; can choose from <strong>OLDEST</strong> (as in: \"oldest first\"), <strong>NEWEST</strong>, <strong>DRAFT_STATUS</strong>, <strong>PARTICIPANT_STATUS</strong>, and/or <strong>CLOSED_DATE</strong>. Where <strong>CLOSED_DATE</strong> is specified and the result set includes pull requests that are not in the closed state, these pull requests will appear first in the result set, followed by most recently closed pull requests. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardAPI.GetPullRequests1(context.Background()).ClosedSince(closedSince).Role(role).ParticipantStatus(participantStatus).State(state).User(user).Order(order).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardAPI.GetPullRequests1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPullRequests1`: GetPullRequests1200Response
	fmt.Fprintf(os.Stdout, "Response from `DashboardAPI.GetPullRequests1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequests1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **closedSince** | **string** | (optional, defaults to returning pull requests regardless of closed since date). Permits returning only pull requests with a closed timestamp set more recently that (now - closedSince). Units are in seconds. So for example if closed since 86400 is set only pull requests closed in the previous 24 hours will be returned. | 
 **role** | **string** | (optional, defaults to returning pull requests for any role). If a role is supplied only pull requests where the authenticated user is a participant in the given role will be returned. Either &lt;strong&gt;REVIEWER&lt;/strong&gt;, &lt;strong&gt;AUTHOR&lt;/strong&gt; or &lt;strong&gt;PARTICIPANT&lt;/strong&gt;. | 
 **participantStatus** | **string** | (optional, defaults to returning pull requests with any participant status). A comma separated list of participant status. That is, one or more of &lt;strong&gt;UNAPPROVED&lt;/strong&gt;, &lt;strong&gt;NEEDS_WORK&lt;/strong&gt;, or &lt;strong&gt;APPROVED&lt;/strong&gt;. | 
 **state** | **string** | (optional, defaults to returning pull requests in any state). If a state is supplied only pull requests in the specified state will be returned. Either &lt;strong&gt;OPEN&lt;/strong&gt;, &lt;strong&gt;DECLINED&lt;/strong&gt; or &lt;strong&gt;MERGED&lt;/strong&gt;. Omit this parameter to return pull request in any state. | 
 **user** | **string** | The name of the involved user, defaults to the current user. | 
 **order** | **string** | (optional, defaults to &lt;strong&gt;NEWEST&lt;/strong&gt;) the order/(s) to return pull requests in; can choose from &lt;strong&gt;OLDEST&lt;/strong&gt; (as in: \&quot;oldest first\&quot;), &lt;strong&gt;NEWEST&lt;/strong&gt;, &lt;strong&gt;DRAFT_STATUS&lt;/strong&gt;, &lt;strong&gt;PARTICIPANT_STATUS&lt;/strong&gt;, and/or &lt;strong&gt;CLOSED_DATE&lt;/strong&gt;. Where &lt;strong&gt;CLOSED_DATE&lt;/strong&gt; is specified and the result set includes pull requests that are not in the closed state, these pull requests will appear first in the result set, followed by most recently closed pull requests. | 
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


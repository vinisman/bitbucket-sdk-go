/*
Bitbucket Data Center

This is the reference document for the Atlassian Bitbucket REST API. The REST API is for developers who want to:    - integrate Bitbucket with other applications;   - create scripts that interact with Bitbucket; or   - develop plugins that enhance the Bitbucket UI, using REST to interact with the backend.    You can read more about developing Bitbucket plugins in the [Bitbucket Developer Documentation](https://developer.atlassian.com/bitbucket/server/docs/latest/).

API version: 9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// JiraIntegrationAPIService JiraIntegrationAPI service
type JiraIntegrationAPIService service

type ApiCreateIssueRequest struct {
	ctx context.Context
	ApiService *JiraIntegrationAPIService
	commentId string
	applicationId *string
	body *string
}

// id of the Jira server
func (r ApiCreateIssueRequest) ApplicationId(applicationId string) ApiCreateIssueRequest {
	r.applicationId = &applicationId
	return r
}

// A String representation of the JSON format Jira create issue request see: &lt;a href&#x3D;\&quot;https://docs.atlassian.com/jira/REST/server/#api/2/issue-createIssue\&quot;&gt;Jira REST API&lt;/a&gt;
func (r ApiCreateIssueRequest) Body(body string) ApiCreateIssueRequest {
	r.body = &body
	return r
}

func (r ApiCreateIssueRequest) Execute() (*RestCommentJiraIssue, *http.Response, error) {
	return r.ApiService.CreateIssueExecute(r)
}

/*
CreateIssue Create Jira Issue

Create a Jira issue and associate it with a comment on a pull request.

This resource can only be used with comments on a pull request. Attempting to call this resource with a different type of comment (for example, a comment on a commit) will result in an error. 

 The authenticated user must have <strong>REPO_READ</strong> permission for the repository containing the comment to call this resource.

The JSON structure for the create issue format is specified by Jira's REST v2 API.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param commentId the comment to associate the created Jira issue to
 @return ApiCreateIssueRequest
*/
func (a *JiraIntegrationAPIService) CreateIssue(ctx context.Context, commentId string) ApiCreateIssueRequest {
	return ApiCreateIssueRequest{
		ApiService: a,
		ctx: ctx,
		commentId: commentId,
	}
}

// Execute executes the request
//  @return RestCommentJiraIssue
func (a *JiraIntegrationAPIService) CreateIssueExecute(r ApiCreateIssueRequest) (*RestCommentJiraIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RestCommentJiraIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraIntegrationAPIService.CreateIssue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jira/latest/comments/{commentId}/issues"
	localVarPath = strings.Replace(localVarPath, "{"+"commentId"+"}", url.PathEscape(parameterValueToString(r.commentId, "commentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.applicationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "applicationId", r.applicationId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetAllAccessTokens401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v GetAllAccessTokens401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCommitsByIssueKeyRequest struct {
	ctx context.Context
	ApiService *JiraIntegrationAPIService
	issueKey string
	maxChanges *string
	start *float32
	limit *float32
}

// The maximum number of changes to retrieve for each changeset
func (r ApiGetCommitsByIssueKeyRequest) MaxChanges(maxChanges string) ApiGetCommitsByIssueKeyRequest {
	r.maxChanges = &maxChanges
	return r
}

// Start number for the page (inclusive). If not passed, first page is assumed.
func (r ApiGetCommitsByIssueKeyRequest) Start(start float32) ApiGetCommitsByIssueKeyRequest {
	r.start = &start
	return r
}

// Number of items to return. If not passed, a page size of 25 is used.
func (r ApiGetCommitsByIssueKeyRequest) Limit(limit float32) ApiGetCommitsByIssueKeyRequest {
	r.limit = &limit
	return r
}

func (r ApiGetCommitsByIssueKeyRequest) Execute() (*GetCommitsByIssueKey200Response, *http.Response, error) {
	return r.ApiService.GetCommitsByIssueKeyExecute(r)
}

/*
GetCommitsByIssueKey Get changesets for issue key

Retrieve a page of changesets associated with the given issue key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueKey The issue key to search by
 @return ApiGetCommitsByIssueKeyRequest
*/
func (a *JiraIntegrationAPIService) GetCommitsByIssueKey(ctx context.Context, issueKey string) ApiGetCommitsByIssueKeyRequest {
	return ApiGetCommitsByIssueKeyRequest{
		ApiService: a,
		ctx: ctx,
		issueKey: issueKey,
	}
}

// Execute executes the request
//  @return GetCommitsByIssueKey200Response
func (a *JiraIntegrationAPIService) GetCommitsByIssueKeyExecute(r ApiGetCommitsByIssueKeyRequest) (*GetCommitsByIssueKey200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetCommitsByIssueKey200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraIntegrationAPIService.GetCommitsByIssueKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jira/latest/issues/{issueKey}/commits"
	localVarPath = strings.Replace(localVarPath, "{"+"issueKey"+"}", url.PathEscape(parameterValueToString(r.issueKey, "issueKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.maxChanges != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxChanges", r.maxChanges, "form", "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetEnhancedEntityLinkForProjectRequest struct {
	ctx context.Context
	ApiService *JiraIntegrationAPIService
	projectKey string
}

func (r ApiGetEnhancedEntityLinkForProjectRequest) Execute() (*RestEnhancedEntityLink, *http.Response, error) {
	return r.ApiService.GetEnhancedEntityLinkForProjectExecute(r)
}

/*
GetEnhancedEntityLinkForProject Get entity link

Retrieves the enchanced primary entitylink 

The authenticated user must have <strong>PROJECT_READ</strong> permission for the project having the primary enhanced entitylink. 



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey The project key
 @return ApiGetEnhancedEntityLinkForProjectRequest
*/
func (a *JiraIntegrationAPIService) GetEnhancedEntityLinkForProject(ctx context.Context, projectKey string) ApiGetEnhancedEntityLinkForProjectRequest {
	return ApiGetEnhancedEntityLinkForProjectRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
	}
}

// Execute executes the request
//  @return RestEnhancedEntityLink
func (a *JiraIntegrationAPIService) GetEnhancedEntityLinkForProjectExecute(r ApiGetEnhancedEntityLinkForProjectRequest) (*RestEnhancedEntityLink, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RestEnhancedEntityLink
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraIntegrationAPIService.GetEnhancedEntityLinkForProject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jira/latest/projects/{projectKey}/primary-enhanced-entitylink"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetIssueKeysForPullRequestRequest struct {
	ctx context.Context
	ApiService *JiraIntegrationAPIService
	projectKey string
	pullRequestId string
	repositorySlug string
}

func (r ApiGetIssueKeysForPullRequestRequest) Execute() ([]RestJiraIssue, *http.Response, error) {
	return r.ApiService.GetIssueKeysForPullRequestExecute(r)
}

/*
GetIssueKeysForPullRequest Get issues for a pull request

Retrieves Jira issue keys that are associated with the commits in the specified pull request. The number of commits checked for issues is limited to a default of 100.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectKey The project key
 @param pullRequestId The pull request id
 @param repositorySlug The repository slug
 @return ApiGetIssueKeysForPullRequestRequest
*/
func (a *JiraIntegrationAPIService) GetIssueKeysForPullRequest(ctx context.Context, projectKey string, pullRequestId string, repositorySlug string) ApiGetIssueKeysForPullRequestRequest {
	return ApiGetIssueKeysForPullRequestRequest{
		ApiService: a,
		ctx: ctx,
		projectKey: projectKey,
		pullRequestId: pullRequestId,
		repositorySlug: repositorySlug,
	}
}

// Execute executes the request
//  @return []RestJiraIssue
func (a *JiraIntegrationAPIService) GetIssueKeysForPullRequestExecute(r ApiGetIssueKeysForPullRequestRequest) ([]RestJiraIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RestJiraIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraIntegrationAPIService.GetIssueKeysForPullRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/jira/latest/projects/{projectKey}/repos/{repositorySlug}/pull-requests/{pullRequestId}/issues"
	localVarPath = strings.Replace(localVarPath, "{"+"projectKey"+"}", url.PathEscape(parameterValueToString(r.projectKey, "projectKey")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pullRequestId"+"}", url.PathEscape(parameterValueToString(r.pullRequestId, "pullRequestId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repositorySlug"+"}", url.PathEscape(parameterValueToString(r.repositorySlug, "repositorySlug")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

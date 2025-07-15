# \BuildsAndDeploymentsAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](BuildsAndDeploymentsAPI.md#Add) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/builds | Store a build status
[**AddAnnotations**](BuildsAndDeploymentsAPI.md#AddAnnotations) | **Post** /insights/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/reports/{key}/annotations | Add Code Insights annotations
[**CreateOrUpdateDeployment**](BuildsAndDeploymentsAPI.md#CreateOrUpdateDeployment) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/deployments | Create or update a deployment
[**CreateRequiredBuildsMergeCheck**](BuildsAndDeploymentsAPI.md#CreateRequiredBuildsMergeCheck) | **Post** /required-builds/latest/projects/{projectKey}/repos/{repositorySlug}/condition | Create a required builds merge check
[**Delete**](BuildsAndDeploymentsAPI.md#Delete) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/builds | Delete a specific build status
[**Delete1**](BuildsAndDeploymentsAPI.md#Delete1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/deployments | Delete a deployment
[**DeleteACodeInsightsReport**](BuildsAndDeploymentsAPI.md#DeleteACodeInsightsReport) | **Delete** /insights/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/reports/{key} | Delete a Code Insights report
[**DeleteAnnotations**](BuildsAndDeploymentsAPI.md#DeleteAnnotations) | **Delete** /insights/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/reports/{key}/annotations | Delete Code Insights annotations
[**DeleteRequiredBuildsMergeCheck**](BuildsAndDeploymentsAPI.md#DeleteRequiredBuildsMergeCheck) | **Delete** /required-builds/latest/projects/{projectKey}/repos/{repositorySlug}/condition/{id} | Delete a required builds merge check
[**Get**](BuildsAndDeploymentsAPI.md#Get) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/builds | Get a specific build status
[**Get1**](BuildsAndDeploymentsAPI.md#Get1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/deployments | Get a deployment
[**GetACodeInsightsReport**](BuildsAndDeploymentsAPI.md#GetACodeInsightsReport) | **Get** /insights/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/reports/{key} | Get a Code Insights report
[**GetAnnotations**](BuildsAndDeploymentsAPI.md#GetAnnotations) | **Get** /insights/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/reports/{key}/annotations | Get Code Insights annotations for a report
[**GetAnnotations1**](BuildsAndDeploymentsAPI.md#GetAnnotations1) | **Get** /insights/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/annotations | Get Code Insights annotations for a commit
[**GetBuildStatusStats**](BuildsAndDeploymentsAPI.md#GetBuildStatusStats) | **Get** /build-status/latest/commits/stats/{commitId} | Get build status statistics for commit
[**GetMultipleBuildStatusStats**](BuildsAndDeploymentsAPI.md#GetMultipleBuildStatusStats) | **Post** /build-status/latest/commits/stats | Get build status statistics for multiple commits
[**GetPageOfRequiredBuildsMergeChecks**](BuildsAndDeploymentsAPI.md#GetPageOfRequiredBuildsMergeChecks) | **Get** /required-builds/latest/projects/{projectKey}/repos/{repositorySlug}/conditions | Get required builds merge checks
[**GetReports**](BuildsAndDeploymentsAPI.md#GetReports) | **Get** /insights/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/reports | Get all Code Insights reports for a commit
[**SetACodeInsightsReport**](BuildsAndDeploymentsAPI.md#SetACodeInsightsReport) | **Put** /insights/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/reports/{key} | Create a Code Insights report
[**SetAnnotation**](BuildsAndDeploymentsAPI.md#SetAnnotation) | **Put** /insights/latest/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/reports/{key}/annotations/{externalId} | Create or replace a Code Insights annotation
[**UpdateRequiredBuildsMergeCheck**](BuildsAndDeploymentsAPI.md#UpdateRequiredBuildsMergeCheck) | **Put** /required-builds/latest/projects/{projectKey}/repos/{repositorySlug}/condition/{id} | Update a required builds merge check



## Add

> Add(ctx, projectKey, commitId, repositorySlug).RestBuildStatusSetRequest(restBuildStatusSetRequest).Execute()

Store a build status



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
	commitId := "commitId_example" // string | The commit.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restBuildStatusSetRequest := *openapiclient.NewRestBuildStatusSetRequest("TEST-REP123", "State_example", "https://bamboo.example.com/browse/TEST-REP1-3") // RestBuildStatusSetRequest | The contents of the build status request are: These fields are **required**:   - **key**: The string referring to this branch plan/job - **state**: The build status state, one of: \"SUCCESSFUL\", \"FAILED\", \"INPROGRESS\", \"CANCELLED\", \"UNKNOWN\" - **url**: URL referring to the build result page in the CI tool.   These fields are optional:   - **buildNumber** (optional): A unique identifier for this particular run of a plan< - **dateAdded** (optional): milliseconds since epoch. If not provided current date is used. - **description** (optional): Describes the build result - **duration** (optional): Duration of a completed build in milliseconds. - **name** (optional): A short string that describes the build plan - **parent** (optional): The identifier for the plan or job that ran the branch plan that produced this build status. - **ref** (optional): The fully qualified git reference e.g. refs/heads/master. - **testResults** (optional): A summary of the passed, failed and skipped tests.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAndDeploymentsAPI.Add(context.Background(), projectKey, commitId, repositorySlug).RestBuildStatusSetRequest(restBuildStatusSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.Add``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restBuildStatusSetRequest** | [**RestBuildStatusSetRequest**](RestBuildStatusSetRequest.md) | The contents of the build status request are: These fields are **required**:   - **key**: The string referring to this branch plan/job - **state**: The build status state, one of: \&quot;SUCCESSFUL\&quot;, \&quot;FAILED\&quot;, \&quot;INPROGRESS\&quot;, \&quot;CANCELLED\&quot;, \&quot;UNKNOWN\&quot; - **url**: URL referring to the build result page in the CI tool.   These fields are optional:   - **buildNumber** (optional): A unique identifier for this particular run of a plan&lt; - **dateAdded** (optional): milliseconds since epoch. If not provided current date is used. - **description** (optional): Describes the build result - **duration** (optional): Duration of a completed build in milliseconds. - **name** (optional): A short string that describes the build plan - **parent** (optional): The identifier for the plan or job that ran the branch plan that produced this build status. - **ref** (optional): The fully qualified git reference e.g. refs/heads/master. - **testResults** (optional): A summary of the passed, failed and skipped tests.  | 

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


## AddAnnotations

> AddAnnotations(ctx, projectKey, commitId, repositorySlug, key).RestBulkAddInsightAnnotationRequest(restBulkAddInsightAnnotationRequest).Execute()

Add Code Insights annotations



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
	commitId := "commitId_example" // string | The commit ID on which to record the annotation. This must be a full 40 character commit hash.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	key := "key_example" // string | The key of the report to which this annotation belongs.
	restBulkAddInsightAnnotationRequest := *openapiclient.NewRestBulkAddInsightAnnotationRequest() // RestBulkAddInsightAnnotationRequest | The annotations to add. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAndDeploymentsAPI.AddAnnotations(context.Background(), projectKey, commitId, repositorySlug, key).RestBulkAddInsightAnnotationRequest(restBulkAddInsightAnnotationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.AddAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit ID on which to record the annotation. This must be a full 40 character commit hash. | 
**repositorySlug** | **string** | The repository slug. | 
**key** | **string** | The key of the report to which this annotation belongs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAnnotationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **restBulkAddInsightAnnotationRequest** | [**RestBulkAddInsightAnnotationRequest**](RestBulkAddInsightAnnotationRequest.md) | The annotations to add. | 

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


## CreateOrUpdateDeployment

> RestDeployment CreateOrUpdateDeployment(ctx, projectKey, commitId, repositorySlug).RestDeploymentSetRequest(restDeploymentSetRequest).Execute()

Create or update a deployment



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
	commitId := "commitId_example" // string | the commitId that was deployed as indicated by the path
	repositorySlug := "repositorySlug_example" // string | The repository slug
	restDeploymentSetRequest := *openapiclient.NewRestDeploymentSetRequest(int64(2), "2nd deployment of commit 44bca31f4be to US East production", "US East marketing website production", *openapiclient.NewRestDeploymentEnvironment("US East Mirror", "us-east-mirror"), "marketing-us-prod", "SUCCESSFUL", "https://my-dep-tool/marketing-us-prod/2") // RestDeploymentSetRequest | the details of the deployment to create, as detailed by the request body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.CreateOrUpdateDeployment(context.Background(), projectKey, commitId, repositorySlug).RestDeploymentSetRequest(restDeploymentSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.CreateOrUpdateDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateDeployment`: RestDeployment
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.CreateOrUpdateDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | the commitId that was deployed as indicated by the path | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restDeploymentSetRequest** | [**RestDeploymentSetRequest**](RestDeploymentSetRequest.md) | the details of the deployment to create, as detailed by the request body | 

### Return type

[**RestDeployment**](RestDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequiredBuildsMergeCheck

> RestRequiredBuildCondition CreateRequiredBuildsMergeCheck(ctx, projectKey, repositorySlug).RestRequiredBuildConditionSetRequest(restRequiredBuildConditionSetRequest).Execute()

Create a required builds merge check



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
	projectKey := "projectKey_example" // string | The project that the repository belongs to
	repositorySlug := "repositorySlug_example" // string | The repository being used
	restRequiredBuildConditionSetRequest := *openapiclient.NewRestRequiredBuildConditionSetRequest([]string{"BuildParentKeys_example"}, *openapiclient.NewRestRefMatcher()) // RestRequiredBuildConditionSetRequest | The request specifying the required build parent keys, ref matcher and exemption matcher (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.CreateRequiredBuildsMergeCheck(context.Background(), projectKey, repositorySlug).RestRequiredBuildConditionSetRequest(restRequiredBuildConditionSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.CreateRequiredBuildsMergeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequiredBuildsMergeCheck`: RestRequiredBuildCondition
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.CreateRequiredBuildsMergeCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project that the repository belongs to | 
**repositorySlug** | **string** | The repository being used | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequiredBuildsMergeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restRequiredBuildConditionSetRequest** | [**RestRequiredBuildConditionSetRequest**](RestRequiredBuildConditionSetRequest.md) | The request specifying the required build parent keys, ref matcher and exemption matcher | 

### Return type

[**RestRequiredBuildCondition**](RestRequiredBuildCondition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, projectKey, commitId, repositorySlug).Key(key).Execute()

Delete a specific build status



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
	commitId := "commitId_example" // string | The commit.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	key := "key_example" // string | the key of the build status

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAndDeploymentsAPI.Delete(context.Background(), projectKey, commitId, repositorySlug).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **key** | **string** | the key of the build status | 

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


## Delete1

> Delete1(ctx, projectKey, commitId, repositorySlug).DeploymentSequenceNumber(deploymentSequenceNumber).Key(key).EnvironmentKey(environmentKey).Execute()

Delete a deployment



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
	commitId := "commitId_example" // string | the commitId that was deployed as indicated by the path
	repositorySlug := "repositorySlug_example" // string | The repository slug
	deploymentSequenceNumber := "deploymentSequenceNumber_example" // string | the sequence number of the deployment, as detailed by the query parameter (optional)
	key := "key_example" // string | the key of the deployment, as detailed by the query parameter (optional)
	environmentKey := "environmentKey_example" // string | the key of the environment, as detailed by the query parameter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAndDeploymentsAPI.Delete1(context.Background(), projectKey, commitId, repositorySlug).DeploymentSequenceNumber(deploymentSequenceNumber).Key(key).EnvironmentKey(environmentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.Delete1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | the commitId that was deployed as indicated by the path | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deploymentSequenceNumber** | **string** | the sequence number of the deployment, as detailed by the query parameter | 
 **key** | **string** | the key of the deployment, as detailed by the query parameter | 
 **environmentKey** | **string** | the key of the environment, as detailed by the query parameter | 

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


## DeleteACodeInsightsReport

> DeleteACodeInsightsReport(ctx, projectKey, commitId, repositorySlug, key).Execute()

Delete a Code Insights report



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
	commitId := "commitId_example" // string | The commit ID on which to record the annotation. This must be a full 40 character commit hash.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	key := "key_example" // string | The key of the report to which this annotation belongs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAndDeploymentsAPI.DeleteACodeInsightsReport(context.Background(), projectKey, commitId, repositorySlug, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.DeleteACodeInsightsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit ID on which to record the annotation. This must be a full 40 character commit hash. | 
**repositorySlug** | **string** | The repository slug. | 
**key** | **string** | The key of the report to which this annotation belongs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteACodeInsightsReportRequest struct via the builder pattern


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


## DeleteAnnotations

> DeleteAnnotations(ctx, projectKey, commitId, repositorySlug, key).ExternalId(externalId).Execute()

Delete Code Insights annotations



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
	commitId := "commitId_example" // string | The commit ID on which to record the annotation. This must be a full 40 character commit hash.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	key := "key_example" // string | The key of the report to which this annotation belongs.
	externalId := "externalId_example" // string | The external IDs for the annotations that are to be deleted. Can be specified more than once to delete by more than one external ID, or can be unspecified to delete all annotations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAndDeploymentsAPI.DeleteAnnotations(context.Background(), projectKey, commitId, repositorySlug, key).ExternalId(externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.DeleteAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit ID on which to record the annotation. This must be a full 40 character commit hash. | 
**repositorySlug** | **string** | The repository slug. | 
**key** | **string** | The key of the report to which this annotation belongs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnnotationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **externalId** | **string** | The external IDs for the annotations that are to be deleted. Can be specified more than once to delete by more than one external ID, or can be unspecified to delete all annotations. | 

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


## DeleteRequiredBuildsMergeCheck

> DeleteRequiredBuildsMergeCheck(ctx, projectKey, id, repositorySlug).Execute()

Delete a required builds merge check



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
	projectKey := "projectKey_example" // string | The project that the repository belongs to
	id := int64(789) // int64 | 
	repositorySlug := "repositorySlug_example" // string | The repository being used

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAndDeploymentsAPI.DeleteRequiredBuildsMergeCheck(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.DeleteRequiredBuildsMergeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project that the repository belongs to | 
**id** | **int64** |  | 
**repositorySlug** | **string** | The repository being used | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequiredBuildsMergeCheckRequest struct via the builder pattern


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


## Get

> RestBuildStatus Get(ctx, projectKey, commitId, repositorySlug).Key(key).Execute()

Get a specific build status



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
	commitId := "commitId_example" // string | The commit.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	key := "key_example" // string | the key of the build status

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.Get(context.Background(), projectKey, commitId, repositorySlug).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: RestBuildStatus
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **key** | **string** | the key of the build status | 

### Return type

[**RestBuildStatus**](RestBuildStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get1

> RestDeployment Get1(ctx, projectKey, commitId, repositorySlug).DeploymentSequenceNumber(deploymentSequenceNumber).Key(key).EnvironmentKey(environmentKey).Execute()

Get a deployment



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
	commitId := "commitId_example" // string | the commitId that was deployed as indicated by the query parameter
	repositorySlug := "repositorySlug_example" // string | The repository slug
	deploymentSequenceNumber := "deploymentSequenceNumber" // string | the sequence number of the deployment, as detailed by the query param (optional)
	key := "key_example" // string | the key of the deployment, as detailed by the query parameter (optional)
	environmentKey := "environmentKey_example" // string | the key of the environment, as detailed by the query parameter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.Get1(context.Background(), projectKey, commitId, repositorySlug).DeploymentSequenceNumber(deploymentSequenceNumber).Key(key).EnvironmentKey(environmentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.Get1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get1`: RestDeployment
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.Get1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**commitId** | **string** | the commitId that was deployed as indicated by the query parameter | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deploymentSequenceNumber** | **string** | the sequence number of the deployment, as detailed by the query param | 
 **key** | **string** | the key of the deployment, as detailed by the query parameter | 
 **environmentKey** | **string** | the key of the environment, as detailed by the query parameter | 

### Return type

[**RestDeployment**](RestDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetACodeInsightsReport

> RestInsightReport GetACodeInsightsReport(ctx, projectKey, commitId, repositorySlug, key).Execute()

Get a Code Insights report



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
	commitId := "commitId_example" // string | The commit ID on which to record the annotation. This must be a full 40 character commit hash.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	key := "key_example" // string | The report key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.GetACodeInsightsReport(context.Background(), projectKey, commitId, repositorySlug, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.GetACodeInsightsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetACodeInsightsReport`: RestInsightReport
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.GetACodeInsightsReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit ID on which to record the annotation. This must be a full 40 character commit hash. | 
**repositorySlug** | **string** | The repository slug. | 
**key** | **string** | The report key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetACodeInsightsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**RestInsightReport**](RestInsightReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnnotations

> RestInsightAnnotationsResponse GetAnnotations(ctx, projectKey, commitId, repositorySlug, key).Execute()

Get Code Insights annotations for a report



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
	commitId := "commitId_example" // string | The commit ID on which to record the annotation. This must be a full 40 character commit hash.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	key := "key_example" // string | The report key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.GetAnnotations(context.Background(), projectKey, commitId, repositorySlug, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.GetAnnotations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnnotations`: RestInsightAnnotationsResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.GetAnnotations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit ID on which to record the annotation. This must be a full 40 character commit hash. | 
**repositorySlug** | **string** | The repository slug. | 
**key** | **string** | The report key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnotationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**RestInsightAnnotationsResponse**](RestInsightAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnnotations1

> RestInsightAnnotationsResponse GetAnnotations1(ctx, projectKey, commitId, repositorySlug).Severity(severity).Path(path).ExternalId(externalId).Type_(type_).Key(key).Execute()

Get Code Insights annotations for a commit



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
	commitId := "commitId_example" // string | The commit ID on which to record the annotation. This must be a full 40 character commit hash.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	severity := "severity_example" // string | Return only annotations that have one of the given severities. Can be specified more than once to filter by more than one severity. Valid severities are <code>LOW</code>, <code>MEDIUM</code> and <code>HIGH</code>. (optional)
	path := "path_example" // string | Return only annotations that appear on one of the provided paths. Can be specified more than once to filter by more than one path. (optional)
	externalId := "externalId_example" // string | Return only annotations that have one of the provided external IDs. Can be specified more than once to filter by more than one external ID. (optional)
	type_ := "type__example" // string | Return only annotations that have one of the given types. Can be specified more than once to filter by multiple types. Valid types are <code>BUG</code>, <code>CODE_SMELL</code>, and <code>VULNERABILITY</code>. (optional)
	key := "key_example" // string | Return only annotations that belong to one of the provided report keys. Can be specified more than once to filter by more than one report (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.GetAnnotations1(context.Background(), projectKey, commitId, repositorySlug).Severity(severity).Path(path).ExternalId(externalId).Type_(type_).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.GetAnnotations1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnnotations1`: RestInsightAnnotationsResponse
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.GetAnnotations1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit ID on which to record the annotation. This must be a full 40 character commit hash. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnotations1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **severity** | **string** | Return only annotations that have one of the given severities. Can be specified more than once to filter by more than one severity. Valid severities are &lt;code&gt;LOW&lt;/code&gt;, &lt;code&gt;MEDIUM&lt;/code&gt; and &lt;code&gt;HIGH&lt;/code&gt;. | 
 **path** | **string** | Return only annotations that appear on one of the provided paths. Can be specified more than once to filter by more than one path. | 
 **externalId** | **string** | Return only annotations that have one of the provided external IDs. Can be specified more than once to filter by more than one external ID. | 
 **type_** | **string** | Return only annotations that have one of the given types. Can be specified more than once to filter by multiple types. Valid types are &lt;code&gt;BUG&lt;/code&gt;, &lt;code&gt;CODE_SMELL&lt;/code&gt;, and &lt;code&gt;VULNERABILITY&lt;/code&gt;. | 
 **key** | **string** | Return only annotations that belong to one of the provided report keys. Can be specified more than once to filter by more than one report | 

### Return type

[**RestInsightAnnotationsResponse**](RestInsightAnnotationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

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
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.GetBuildStatusStats(context.Background(), commitId).IncludeUnique(includeUnique).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.GetBuildStatusStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildStatusStats`: RestBuildStats
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.GetBuildStatusStats`: %v\n", resp)
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
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.GetMultipleBuildStatusStats(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.GetMultipleBuildStatusStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipleBuildStatusStats`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.GetMultipleBuildStatusStats`: %v\n", resp)
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


## GetPageOfRequiredBuildsMergeChecks

> GetPageOfRequiredBuildsMergeChecks200Response GetPageOfRequiredBuildsMergeChecks(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Execute()

Get required builds merge checks



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
	projectKey := "projectKey_example" // string | The project that the repository belongs to
	repositorySlug := "repositorySlug_example" // string | The repository being used
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.GetPageOfRequiredBuildsMergeChecks(context.Background(), projectKey, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.GetPageOfRequiredBuildsMergeChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPageOfRequiredBuildsMergeChecks`: GetPageOfRequiredBuildsMergeChecks200Response
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.GetPageOfRequiredBuildsMergeChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project that the repository belongs to | 
**repositorySlug** | **string** | The repository being used | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageOfRequiredBuildsMergeChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetPageOfRequiredBuildsMergeChecks200Response**](GetPageOfRequiredBuildsMergeChecks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReports

> GetReports200Response GetReports(ctx, projectKey, commitId, repositorySlug).Start(start).Limit(limit).Execute()

Get all Code Insights reports for a commit



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
	commitId := "commitId_example" // string | The commit ID on which to record the annotation. This must be a full 40 character commit hash.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.GetReports(context.Background(), projectKey, commitId, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.GetReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReports`: GetReports200Response
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.GetReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit ID on which to record the annotation. This must be a full 40 character commit hash. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetReports200Response**](GetReports200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetACodeInsightsReport

> RestInsightReport SetACodeInsightsReport(ctx, projectKey, commitId, repositorySlug, key).RestSetInsightReportRequest(restSetInsightReportRequest).Execute()

Create a Code Insights report



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
	commitId := "commitId_example" // string | The commit ID on which to record the annotation. This must be a full 40 character commit hash.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	key := "key_example" // string | A unique string representing the report as chosen by the reporter. This should be unique enough to not clash with other report's keys. To do this, we recommend namespacing the key using reverse DNS
	restSetInsightReportRequest := *openapiclient.NewRestSetInsightReportRequest([]openapiclient.RestInsightReportData{*openapiclient.NewRestInsightReportData()}, "report.title") // RestSetInsightReportRequest | The request object containing the details of the report to create (see example). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.SetACodeInsightsReport(context.Background(), projectKey, commitId, repositorySlug, key).RestSetInsightReportRequest(restSetInsightReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.SetACodeInsightsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetACodeInsightsReport`: RestInsightReport
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.SetACodeInsightsReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**commitId** | **string** | The commit ID on which to record the annotation. This must be a full 40 character commit hash. | 
**repositorySlug** | **string** | The repository slug. | 
**key** | **string** | A unique string representing the report as chosen by the reporter. This should be unique enough to not clash with other report&#39;s keys. To do this, we recommend namespacing the key using reverse DNS | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetACodeInsightsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **restSetInsightReportRequest** | [**RestSetInsightReportRequest**](RestSetInsightReportRequest.md) | The request object containing the details of the report to create (see example). | 

### Return type

[**RestInsightReport**](RestInsightReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAnnotation

> SetAnnotation(ctx, projectKey, externalId, commitId, repositorySlug, key).RestSingleAddInsightAnnotationRequest(restSingleAddInsightAnnotationRequest).Execute()

Create or replace a Code Insights annotation



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
	externalId := "externalId_example" // string | The external ID of the annotation that is to be updated or created
	commitId := "commitId_example" // string | The commit ID on which to record the annotation. This must be a full 40 character commit hash.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	key := "key_example" // string | The key of the report to which this annotation belongs
	restSingleAddInsightAnnotationRequest := *openapiclient.NewRestSingleAddInsightAnnotationRequest("This is a bug here because reasons", "MEDIUM") // RestSingleAddInsightAnnotationRequest | The new annotation that is to replace the existing one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuildsAndDeploymentsAPI.SetAnnotation(context.Background(), projectKey, externalId, commitId, repositorySlug, key).RestSingleAddInsightAnnotationRequest(restSingleAddInsightAnnotationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.SetAnnotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**externalId** | **string** | The external ID of the annotation that is to be updated or created | 
**commitId** | **string** | The commit ID on which to record the annotation. This must be a full 40 character commit hash. | 
**repositorySlug** | **string** | The repository slug. | 
**key** | **string** | The key of the report to which this annotation belongs | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAnnotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **restSingleAddInsightAnnotationRequest** | [**RestSingleAddInsightAnnotationRequest**](RestSingleAddInsightAnnotationRequest.md) | The new annotation that is to replace the existing one. | 

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


## UpdateRequiredBuildsMergeCheck

> RestRequiredBuildCondition UpdateRequiredBuildsMergeCheck(ctx, projectKey, id, repositorySlug).RestRequiredBuildConditionSetRequest(restRequiredBuildConditionSetRequest).Execute()

Update a required builds merge check



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
	projectKey := "projectKey_example" // string | The project that the repository belongs to
	id := int64(789) // int64 | 
	repositorySlug := "repositorySlug_example" // string | The repository being used
	restRequiredBuildConditionSetRequest := *openapiclient.NewRestRequiredBuildConditionSetRequest([]string{"BuildParentKeys_example"}, *openapiclient.NewRestRefMatcher()) // RestRequiredBuildConditionSetRequest | The request specifying the required build parent keys, ref matcher and exemption matcher (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAndDeploymentsAPI.UpdateRequiredBuildsMergeCheck(context.Background(), projectKey, id, repositorySlug).RestRequiredBuildConditionSetRequest(restRequiredBuildConditionSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAndDeploymentsAPI.UpdateRequiredBuildsMergeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRequiredBuildsMergeCheck`: RestRequiredBuildCondition
	fmt.Fprintf(os.Stdout, "Response from `BuildsAndDeploymentsAPI.UpdateRequiredBuildsMergeCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project that the repository belongs to | 
**id** | **int64** |  | 
**repositorySlug** | **string** | The repository being used | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequiredBuildsMergeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restRequiredBuildConditionSetRequest** | [**RestRequiredBuildConditionSetRequest**](RestRequiredBuildConditionSetRequest.md) | The request specifying the required build parent keys, ref matcher and exemption matcher | 

### Return type

[**RestRequiredBuildCondition**](RestRequiredBuildCondition.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


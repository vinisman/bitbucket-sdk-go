# \ProjectAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDefaultTask**](ProjectAPI.md#AddDefaultTask) | **Post** /default-tasks/latest/projects/{projectKey}/tasks | Add a default task
[**Create3**](ProjectAPI.md#Create3) | **Post** /api/latest/projects/{projectKey}/settings-restriction | Enforce project restriction
[**CreateProject**](ProjectAPI.md#CreateProject) | **Post** /api/latest/projects | Create a new project
[**CreateRepository**](ProjectAPI.md#CreateRepository) | **Post** /api/latest/projects/{projectKey}/repos | Create repository
[**CreateRestrictions**](ProjectAPI.md#CreateRestrictions) | **Post** /branch-permissions/latest/projects/{projectKey}/restrictions | Create multiple ref restrictions
[**CreateWebhook**](ProjectAPI.md#CreateWebhook) | **Post** /api/latest/projects/{projectKey}/webhooks | Create webhook
[**Delete4**](ProjectAPI.md#Delete4) | **Delete** /api/latest/projects/{projectKey}/settings/auto-merge | Delete pull request auto-merge settings
[**Delete9**](ProjectAPI.md#Delete9) | **Delete** /api/latest/projects/{projectKey}/settings-restriction | Stop enforcing project restriction
[**DeleteAllDefaultTasks**](ProjectAPI.md#DeleteAllDefaultTasks) | **Delete** /default-tasks/latest/projects/{projectKey}/tasks | Deletes all default tasks for the project
[**DeleteAutoDeclineSettings**](ProjectAPI.md#DeleteAutoDeclineSettings) | **Delete** /api/latest/projects/{projectKey}/settings/auto-decline | Delete auto decline settings
[**DeleteDefaultTask**](ProjectAPI.md#DeleteDefaultTask) | **Delete** /default-tasks/latest/projects/{projectKey}/tasks/{taskId} | Delete a specific default task
[**DeleteProject**](ProjectAPI.md#DeleteProject) | **Delete** /api/latest/projects/{projectKey} | Delete project
[**DeleteRepository**](ProjectAPI.md#DeleteRepository) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug} | Delete repository
[**DeleteRestriction**](ProjectAPI.md#DeleteRestriction) | **Delete** /branch-permissions/latest/projects/{projectKey}/restrictions/{id} | Delete a ref restriction
[**DeleteWebhook**](ProjectAPI.md#DeleteWebhook) | **Delete** /api/latest/projects/{projectKey}/webhooks/{webhookId} | Delete webhook
[**DisableHook**](ProjectAPI.md#DisableHook) | **Delete** /api/latest/projects/{projectKey}/settings/hooks/{hookKey}/enabled | Disable repository hook
[**EnableHook**](ProjectAPI.md#EnableHook) | **Put** /api/latest/projects/{projectKey}/settings/hooks/{hookKey}/enabled | Enable repository hook
[**FindWebhooks**](ProjectAPI.md#FindWebhooks) | **Get** /api/latest/projects/{projectKey}/webhooks | Find webhooks
[**ForkRepository**](ProjectAPI.md#ForkRepository) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug} | Fork repository
[**Get4**](ProjectAPI.md#Get4) | **Get** /api/latest/projects/{projectKey}/settings/auto-merge | Get pull request auto-merge settings
[**Get7**](ProjectAPI.md#Get7) | **Get** /api/latest/projects/{projectKey}/settings-restriction | Get enforcing project setting
[**GetAll**](ProjectAPI.md#GetAll) | **Get** /api/latest/projects/{projectKey}/settings-restriction/all | Get all enforcing project settings
[**GetAutoDeclineSettings**](ProjectAPI.md#GetAutoDeclineSettings) | **Get** /api/latest/projects/{projectKey}/settings/auto-decline | Get auto decline settings
[**GetAvatar**](ProjectAPI.md#GetAvatar) | **Get** /api/latest/hooks/{hookKey}/avatar | Get project avatar
[**GetConfigurations**](ProjectAPI.md#GetConfigurations) | **Get** /api/latest/projects/{projectKey}/hook-scripts | Get configured hook scripts
[**GetDefaultBranch2**](ProjectAPI.md#GetDefaultBranch2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/default-branch | Get repository default branch
[**GetDefaultTasks**](ProjectAPI.md#GetDefaultTasks) | **Get** /default-tasks/latest/projects/{projectKey}/tasks | Get a page of default tasks
[**GetForkedRepositories**](ProjectAPI.md#GetForkedRepositories) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/forks | Get repository forks
[**GetGroupsWithAnyPermission1**](ProjectAPI.md#GetGroupsWithAnyPermission1) | **Get** /api/latest/projects/{projectKey}/permissions/groups | Get groups with permission to project
[**GetGroupsWithoutAnyPermission1**](ProjectAPI.md#GetGroupsWithoutAnyPermission1) | **Get** /api/latest/projects/{projectKey}/permissions/groups/none | Get groups without project permission
[**GetLatestInvocation**](ProjectAPI.md#GetLatestInvocation) | **Get** /api/latest/projects/{projectKey}/webhooks/{webhookId}/latest | Get last webhook invocation details
[**GetProject**](ProjectAPI.md#GetProject) | **Get** /api/latest/projects/{projectKey} | Get a project
[**GetProjectAvatar**](ProjectAPI.md#GetProjectAvatar) | **Get** /api/latest/projects/{projectKey}/avatar.png | Get avatar for project
[**GetProjects**](ProjectAPI.md#GetProjects) | **Get** /api/latest/projects | Get projects
[**GetPullRequestSettings**](ProjectAPI.md#GetPullRequestSettings) | **Get** /api/latest/projects/{projectKey}/settings/pull-requests/{scmId} | Get merge strategy
[**GetRelatedRepositories**](ProjectAPI.md#GetRelatedRepositories) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/related | Get related repository
[**GetRepositories**](ProjectAPI.md#GetRepositories) | **Get** /api/latest/projects/{projectKey}/repos | Get repositories for project
[**GetRepository**](ProjectAPI.md#GetRepository) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug} | Get repository
[**GetRepositoryHook**](ProjectAPI.md#GetRepositoryHook) | **Get** /api/latest/projects/{projectKey}/settings/hooks/{hookKey} | Get a repository hook
[**GetRepositoryHooks**](ProjectAPI.md#GetRepositoryHooks) | **Get** /api/latest/projects/{projectKey}/settings/hooks | Get repository hooks
[**GetRestriction**](ProjectAPI.md#GetRestriction) | **Get** /branch-permissions/latest/projects/{projectKey}/restrictions/{id} | Get a ref restriction
[**GetRestrictions**](ProjectAPI.md#GetRestrictions) | **Get** /branch-permissions/latest/projects/{projectKey}/restrictions | Search for ref restrictions
[**GetSettings**](ProjectAPI.md#GetSettings) | **Get** /api/latest/projects/{projectKey}/settings/hooks/{hookKey}/settings | Get repository hook settings
[**GetStatistics**](ProjectAPI.md#GetStatistics) | **Get** /api/latest/projects/{projectKey}/webhooks/{webhookId}/statistics | Get webhook statistics
[**GetStatisticsSummary**](ProjectAPI.md#GetStatisticsSummary) | **Get** /api/latest/projects/{projectKey}/webhooks/{webhookId}/statistics/summary | Get webhook statistics summary
[**GetUsersWithAnyPermission1**](ProjectAPI.md#GetUsersWithAnyPermission1) | **Get** /api/latest/projects/{projectKey}/permissions/users | Get users with permission to project
[**GetUsersWithoutPermission**](ProjectAPI.md#GetUsersWithoutPermission) | **Get** /api/latest/projects/{projectKey}/permissions/users/none | Get users without project permission
[**GetWebhook**](ProjectAPI.md#GetWebhook) | **Get** /api/latest/projects/{projectKey}/webhooks/{webhookId} | Get webhook
[**HasAllUserPermission**](ProjectAPI.md#HasAllUserPermission) | **Get** /api/latest/projects/{projectKey}/permissions/{permission}/all | Check default project permission
[**ModifyAllUserPermission**](ProjectAPI.md#ModifyAllUserPermission) | **Post** /api/latest/projects/{projectKey}/permissions/{permission}/all | Grant project permission
[**RemoveConfiguration**](ProjectAPI.md#RemoveConfiguration) | **Delete** /api/latest/projects/{projectKey}/hook-scripts/{scriptId} | Remove a hook script
[**RetryCreateRepository**](ProjectAPI.md#RetryCreateRepository) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/recreate | Retry repository creation
[**RevokePermissions**](ProjectAPI.md#RevokePermissions) | **Delete** /api/latest/projects/{projectKey}/permissions | Revoke project permissions
[**RevokePermissionsForGroup1**](ProjectAPI.md#RevokePermissionsForGroup1) | **Delete** /api/latest/projects/{projectKey}/permissions/groups | Revoke group project permission
[**RevokePermissionsForUser1**](ProjectAPI.md#RevokePermissionsForUser1) | **Delete** /api/latest/projects/{projectKey}/permissions/users | Revoke user project permission
[**SearchPermissions**](ProjectAPI.md#SearchPermissions) | **Get** /api/latest/projects/{projectKey}/permissions/search | Search project permissions
[**Set**](ProjectAPI.md#Set) | **Put** /api/latest/projects/{projectKey}/settings/auto-merge | Create or update the pull request auto-merge settings
[**SetAutoDeclineSettings**](ProjectAPI.md#SetAutoDeclineSettings) | **Put** /api/latest/projects/{projectKey}/settings/auto-decline | Create/Update auto decline settings
[**SetConfiguration**](ProjectAPI.md#SetConfiguration) | **Put** /api/latest/projects/{projectKey}/hook-scripts/{scriptId} | Create/update a hook script
[**SetDefaultBranch2**](ProjectAPI.md#SetDefaultBranch2) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/default-branch | Update default branch for repository
[**SetPermissionForGroups1**](ProjectAPI.md#SetPermissionForGroups1) | **Put** /api/latest/projects/{projectKey}/permissions/groups | Update group project permission
[**SetPermissionForUsers1**](ProjectAPI.md#SetPermissionForUsers1) | **Put** /api/latest/projects/{projectKey}/permissions/users | Update user project permission
[**SetSettings**](ProjectAPI.md#SetSettings) | **Put** /api/latest/projects/{projectKey}/settings/hooks/{hookKey}/settings | Update repository hook settings
[**StreamContributing**](ProjectAPI.md#StreamContributing) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/contributing | Get repository contributing guidelines
[**StreamLicense**](ProjectAPI.md#StreamLicense) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/license | Get repository license
[**StreamReadme**](ProjectAPI.md#StreamReadme) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/readme | Get repository readme
[**TestWebhook**](ProjectAPI.md#TestWebhook) | **Post** /api/latest/projects/{projectKey}/webhooks/test | Test webhook
[**UpdateDefaultTask**](ProjectAPI.md#UpdateDefaultTask) | **Put** /default-tasks/latest/projects/{projectKey}/tasks/{taskId} | Update a default task
[**UpdateProject**](ProjectAPI.md#UpdateProject) | **Put** /api/latest/projects/{projectKey} | Update project
[**UpdatePullRequestSettings**](ProjectAPI.md#UpdatePullRequestSettings) | **Post** /api/latest/projects/{projectKey}/settings/pull-requests/{scmId} | Update merge strategy
[**UpdateRepository**](ProjectAPI.md#UpdateRepository) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug} | Update repository
[**UpdateWebhook**](ProjectAPI.md#UpdateWebhook) | **Put** /api/latest/projects/{projectKey}/webhooks/{webhookId} | Update webhook
[**UploadAvatar**](ProjectAPI.md#UploadAvatar) | **Post** /api/latest/projects/{projectKey}/avatar.png | Update project avatar



## AddDefaultTask

> RestDefaultTask AddDefaultTask(ctx, projectKey).RestDefaultTaskRequest(restDefaultTaskRequest).Execute()

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
	restDefaultTaskRequest := *openapiclient.NewRestDefaultTaskRequest() // RestDefaultTaskRequest | The task to be added

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.AddDefaultTask(context.Background(), projectKey).RestDefaultTaskRequest(restDefaultTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.AddDefaultTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDefaultTask`: RestDefaultTask
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.AddDefaultTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDefaultTaskRequest struct via the builder pattern


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


## Create3

> RestProjectSettingsRestriction Create3(ctx, projectKey).RestProjectSettingsRestrictionRequest(restProjectSettingsRestrictionRequest).Execute()

Enforce project restriction



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
	restProjectSettingsRestrictionRequest := *openapiclient.NewRestProjectSettingsRestrictionRequest("my-admin-feature", "org.featuredeveloper") // RestProjectSettingsRestrictionRequest | The project settings restriction to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.Create3(context.Background(), projectKey).RestProjectSettingsRestrictionRequest(restProjectSettingsRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.Create3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create3`: RestProjectSettingsRestriction
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.Create3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreate3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restProjectSettingsRestrictionRequest** | [**RestProjectSettingsRestrictionRequest**](RestProjectSettingsRestrictionRequest.md) | The project settings restriction to create | 

### Return type

[**RestProjectSettingsRestriction**](RestProjectSettingsRestriction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> RestProject CreateProject(ctx).RestProject(restProject).Execute()

Create a new project



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
	restProject := *openapiclient.NewRestProject() // RestProject | The project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateProject(context.Background()).RestProject(restProject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProject`: RestProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restProject** | [**RestProject**](RestProject.md) | The project. | 

### Return type

[**RestProject**](RestProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRepository

> RestRepository CreateRepository(ctx, projectKey).RestRepository(restRepository).Execute()

Create repository



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
	restRepository := *openapiclient.NewRestRepository() // RestRepository | The repository (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateRepository(context.Background(), projectKey).RestRepository(restRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRepository`: RestRepository
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restRepository** | [**RestRepository**](RestRepository.md) | The repository | 

### Return type

[**RestRepository**](RestRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRestrictions

> RestRefRestriction CreateRestrictions(ctx, projectKey).RestRestrictionRequest(restRestrictionRequest).Execute()

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
	restRestrictionRequest := []openapiclient.RestRestrictionRequest{*openapiclient.NewRestRestrictionRequest()} // []RestRestrictionRequest | The request containing a list of the details of the restrictions to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateRestrictions(context.Background(), projectKey).RestRestrictionRequest(restRestrictionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRestrictions`: RestRefRestriction
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRestrictionsRequest struct via the builder pattern


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


## CreateWebhook

> RestWebhook CreateWebhook(ctx, projectKey).RestWebhook(restWebhook).Execute()

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
	restWebhook := *openapiclient.NewRestWebhook() // RestWebhook | The webhook to be created for this project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.CreateWebhook(context.Background(), projectKey).RestWebhook(restWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: RestWebhook
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restWebhook** | [**RestWebhook**](RestWebhook.md) | The webhook to be created for this project. | 

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


## Delete4

> Delete4(ctx, projectKey).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.Delete4(context.Background(), projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.Delete4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete4Request struct via the builder pattern


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


## Delete9

> Delete9(ctx, projectKey).Namespace(namespace).FeatureKey(featureKey).ComponentKey(componentKey).Execute()

Stop enforcing project restriction



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
	namespace := "namespace_example" // string | A namespace used to identify the provider of the feature
	featureKey := "featureKey_example" // string | A key to uniquely identify the feature within the provided namespace
	componentKey := "componentKey_example" // string | A key to uniquely identify individually restrictable subcomponents of a feature within the provided feature key and namespace (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.Delete9(context.Background(), projectKey).Namespace(namespace).FeatureKey(featureKey).ComponentKey(componentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.Delete9``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete9Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | A namespace used to identify the provider of the feature | 
 **featureKey** | **string** | A key to uniquely identify the feature within the provided namespace | 
 **componentKey** | **string** | A key to uniquely identify individually restrictable subcomponents of a feature within the provided feature key and namespace | 

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


## DeleteAllDefaultTasks

> DeleteAllDefaultTasks(ctx, projectKey).Execute()

Deletes all default tasks for the project



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
	r, err := apiClient.ProjectAPI.DeleteAllDefaultTasks(context.Background(), projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteAllDefaultTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllDefaultTasksRequest struct via the builder pattern


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


## DeleteAutoDeclineSettings

> DeleteAutoDeclineSettings(ctx, projectKey).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.DeleteAutoDeclineSettings(context.Background(), projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteAutoDeclineSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoDeclineSettingsRequest struct via the builder pattern


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


## DeleteDefaultTask

> DeleteDefaultTask(ctx, projectKey, taskId).Execute()

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
	taskId := "taskId_example" // string | The ID of the default task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.DeleteDefaultTask(context.Background(), projectKey, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteDefaultTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**taskId** | **string** | The ID of the default task | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultTaskRequest struct via the builder pattern


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


## DeleteProject

> DeleteProject(ctx, projectKey).Execute()

Delete project



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
	r, err := apiClient.ProjectAPI.DeleteProject(context.Background(), projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


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


## DeleteRepository

> DeleteRepository(ctx, projectKey, repositorySlug).Execute()

Delete repository



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
	r, err := apiClient.ProjectAPI.DeleteRepository(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteRepository``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


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


## DeleteRestriction

> DeleteRestriction(ctx, projectKey, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.DeleteRestriction(context.Background(), projectKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteRestriction``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRestrictionRequest struct via the builder pattern


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


## DeleteWebhook

> DeleteWebhook(ctx, projectKey, webhookId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.DeleteWebhook(context.Background(), projectKey, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteWebhook``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## DisableHook

> RestRepositoryHook DisableHook(ctx, projectKey, hookKey).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.DisableHook(context.Background(), projectKey, hookKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DisableHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableHook`: RestRepositoryHook
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.DisableHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableHookRequest struct via the builder pattern


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


## EnableHook

> RestRepositoryHook EnableHook(ctx, projectKey, hookKey).ContentLength(contentLength).Execute()

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
	contentLength := int64(789) // int64 | The content length. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.EnableHook(context.Background(), projectKey, hookKey).ContentLength(contentLength).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.EnableHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableHook`: RestRepositoryHook
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.EnableHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentLength** | **int64** | The content length. | 

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


## FindWebhooks

> FindWebhooks(ctx, projectKey).Event(event).Statistics(statistics).Execute()

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
	event := "event_example" // string | List of <code>com.atlassian.webhooks.WebhookEvent</code> IDs to filter for (optional)
	statistics := true // bool | <code>true</code> if statistics should be provided for all found webhooks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.FindWebhooks(context.Background(), projectKey).Event(event).Statistics(statistics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.FindWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWebhooksRequest struct via the builder pattern


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


## ForkRepository

> RestRepository ForkRepository(ctx, projectKey, repositorySlug).RestRepository(restRepository).Execute()

Fork repository



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
	restRepository := *openapiclient.NewRestRepository() // RestRepository | The rest fork. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.ForkRepository(context.Background(), projectKey, repositorySlug).RestRepository(restRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ForkRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForkRepository`: RestRepository
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ForkRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiForkRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restRepository** | [**RestRepository**](RestRepository.md) | The rest fork. | 

### Return type

[**RestRepository**](RestRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get4

> RestAutoMergeRestrictedSettings Get4(ctx, projectKey).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.Get4(context.Background(), projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.Get4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get4`: RestAutoMergeRestrictedSettings
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.Get4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet4Request struct via the builder pattern


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


## Get7

> RestProjectSettingsRestriction Get7(ctx, projectKey).Namespace(namespace).FeatureKey(featureKey).ComponentKey(componentKey).Execute()

Get enforcing project setting



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
	namespace := "namespace_example" // string | The namespace used to identify the provider of the feature
	featureKey := "featureKey_example" // string | The feature key to uniquely identify the feature within the provided namespace
	componentKey := "componentKey_example" // string | The component key to uniquely identify individually restrictable subcomponents of a feature within the provided feature key and namespace (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.Get7(context.Background(), projectKey).Namespace(namespace).FeatureKey(featureKey).ComponentKey(componentKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.Get7``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get7`: RestProjectSettingsRestriction
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.Get7`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | The namespace used to identify the provider of the feature | 
 **featureKey** | **string** | The feature key to uniquely identify the feature within the provided namespace | 
 **componentKey** | **string** | The component key to uniquely identify individually restrictable subcomponents of a feature within the provided feature key and namespace | 

### Return type

[**RestProjectSettingsRestriction**](RestProjectSettingsRestriction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> GetAll200Response GetAll(ctx, projectKey).Namespace(namespace).FeatureKey(featureKey).Start(start).Limit(limit).Execute()

Get all enforcing project settings



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
	namespace := "namespace_example" // string | A namespace used to identify the provider of the feature
	featureKey := "featureKey_example" // string | A key to uniquely identify the feature within the provided namespace
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetAll(context.Background(), projectKey).Namespace(namespace).FeatureKey(featureKey).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAll`: GetAll200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** | A namespace used to identify the provider of the feature | 
 **featureKey** | **string** | A key to uniquely identify the feature within the provided namespace | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetAll200Response**](GetAll200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoDeclineSettings

> RestAutoDeclineSettings GetAutoDeclineSettings(ctx, projectKey).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetAutoDeclineSettings(context.Background(), projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetAutoDeclineSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoDeclineSettings`: RestAutoDeclineSettings
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetAutoDeclineSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoDeclineSettingsRequest struct via the builder pattern


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


## GetAvatar

> GetAvatar(ctx, hookKey).Version(version).Execute()

Get project avatar



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
	hookKey := "hookKey_example" // string | The complete module key of the hook module.
	version := "version_example" // string | (optional) Version used for HTTP caching only - any non-blank version will result in a large max-age Cache-Control header. Note that this does not affect the Last-Modified header. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.GetAvatar(context.Background(), hookKey).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hookKey** | **string** | The complete module key of the hook module. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | (optional) Version used for HTTP caching only - any non-blank version will result in a large max-age Cache-Control header. Note that this does not affect the Last-Modified header. | 

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


## GetConfigurations

> GetConfigurations200Response GetConfigurations(ctx, projectKey).Start(start).Limit(limit).Execute()

Get configured hook scripts



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
	resp, r, err := apiClient.ProjectAPI.GetConfigurations(context.Background(), projectKey).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurations`: GetConfigurations200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationsRequest struct via the builder pattern


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


## GetDefaultBranch2

> RestMinimalRef GetDefaultBranch2(ctx, projectKey, repositorySlug).Execute()

Get repository default branch



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
	resp, r, err := apiClient.ProjectAPI.GetDefaultBranch2(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetDefaultBranch2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultBranch2`: RestMinimalRef
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetDefaultBranch2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultBranch2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestMinimalRef**](RestMinimalRef.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultTasks

> GetDefaultTasks1200Response GetDefaultTasks(ctx, projectKey).Markup(markup).Start(start).Limit(limit).Execute()

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
	markup := "markup_example" // string | If present or \"true\", includes a markup-rendered description (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetDefaultTasks(context.Background(), projectKey).Markup(markup).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetDefaultTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultTasks`: GetDefaultTasks1200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetDefaultTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **markup** | **string** | If present or \&quot;true\&quot;, includes a markup-rendered description | 
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


## GetForkedRepositories

> GetRepositoriesRecentlyAccessed200Response GetForkedRepositories(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Execute()

Get repository forks



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
	resp, r, err := apiClient.ProjectAPI.GetForkedRepositories(context.Background(), projectKey, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetForkedRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForkedRepositories`: GetRepositoriesRecentlyAccessed200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetForkedRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForkedRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetGroupsWithAnyPermission1

> GetGroupsWithAnyPermission200Response GetGroupsWithAnyPermission1(ctx, projectKey).Filter(filter).Start(start).Limit(limit).Execute()

Get groups with permission to project



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
	filter := "filter_example" // string | If specified only group names containing the supplied string will be returned (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetGroupsWithAnyPermission1(context.Background(), projectKey).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetGroupsWithAnyPermission1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithAnyPermission1`: GetGroupsWithAnyPermission200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetGroupsWithAnyPermission1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithAnyPermission1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | If specified only group names containing the supplied string will be returned | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetGroupsWithAnyPermission200Response**](GetGroupsWithAnyPermission200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsWithoutAnyPermission1

> GetGroups1200Response GetGroupsWithoutAnyPermission1(ctx, projectKey).Filter(filter).Start(start).Limit(limit).Execute()

Get groups without project permission



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
	filter := "filter_example" // string | If specified only group names containing the supplied string will be returned (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetGroupsWithoutAnyPermission1(context.Background(), projectKey).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetGroupsWithoutAnyPermission1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithoutAnyPermission1`: GetGroups1200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetGroupsWithoutAnyPermission1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithoutAnyPermission1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | If specified only group names containing the supplied string will be returned | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetGroups1200Response**](GetGroups1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestInvocation

> RestDetailedInvocation GetLatestInvocation(ctx, projectKey, webhookId).Event(event).Outcome(outcome).Execute()

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
	event := "event_example" // string | The string ID of a specific event to retrieve the last invocation for. (optional)
	outcome := "outcome_example" // string | The outcome to filter for. Can be SUCCESS, FAILURE, ERROR. None specified means that the all will be considered (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetLatestInvocation(context.Background(), projectKey, webhookId).Event(event).Outcome(outcome).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetLatestInvocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestInvocation`: RestDetailedInvocation
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetLatestInvocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | ID of the webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestInvocationRequest struct via the builder pattern


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


## GetProject

> RestProject GetProject(ctx, projectKey).Execute()

Get a project



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
	resp, r, err := apiClient.ProjectAPI.GetProject(context.Background(), projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProject`: RestProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestProject**](RestProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectAvatar

> GetProjectAvatar(ctx, projectKey).S(s).Execute()

Get avatar for project



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
	s := "s_example" // string | The desired size of the image. The server will return an image as close as possible to the specified size. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.GetProjectAvatar(context.Background(), projectKey).S(s).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s** | **string** | The desired size of the image. The server will return an image as close as possible to the specified size. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> GetProjects200Response GetProjects(ctx).Name(name).Permission(permission).Start(start).Limit(limit).Execute()

Get projects



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
	name := "name_example" // string | Name to filter by. (optional)
	permission := "permission_example" // string | Permission to filter by (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjects(context.Background()).Name(name).Permission(permission).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjects`: GetProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name to filter by. | 
 **permission** | **string** | Permission to filter by | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetProjects200Response**](GetProjects200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPullRequestSettings

> RestPullRequestSettings GetPullRequestSettings(ctx, projectKey, scmId).Execute()

Get merge strategy



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
	scmId := "scmId_example" // string | The SCM to get strategies for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetPullRequestSettings(context.Background(), projectKey, scmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetPullRequestSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPullRequestSettings`: RestPullRequestSettings
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetPullRequestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**scmId** | **string** | The SCM to get strategies for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullRequestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestPullRequestSettings**](RestPullRequestSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelatedRepositories

> GetRepositoriesRecentlyAccessed200Response GetRelatedRepositories(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Execute()

Get related repository



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
	resp, r, err := apiClient.ProjectAPI.GetRelatedRepositories(context.Background(), projectKey, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetRelatedRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRelatedRepositories`: GetRepositoriesRecentlyAccessed200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetRelatedRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelatedRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetRepositories

> GetRepositoriesRecentlyAccessed200Response GetRepositories(ctx, projectKey).Start(start).Limit(limit).Execute()

Get repositories for project



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
	resp, r, err := apiClient.ProjectAPI.GetRepositories(context.Background(), projectKey).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositories`: GetRepositoriesRecentlyAccessed200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetRepository

> RestRepository GetRepository(ctx, projectKey, repositorySlug).Execute()

Get repository



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
	resp, r, err := apiClient.ProjectAPI.GetRepository(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository`: RestRepository
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestRepository**](RestRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryHook

> RestRepositoryHook GetRepositoryHook(ctx, projectKey, hookKey).Execute()

Get a repository hook



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetRepositoryHook(context.Background(), projectKey, hookKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetRepositoryHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryHook`: RestRepositoryHook
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetRepositoryHook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryHookRequest struct via the builder pattern


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


## GetRepositoryHooks

> GetRepositoryHooks1200Response GetRepositoryHooks(ctx, projectKey).Type_(type_).Start(start).Limit(limit).Execute()

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
	type_ := "type__example" // string | The optional type to filter by. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetRepositoryHooks(context.Background(), projectKey).Type_(type_).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetRepositoryHooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryHooks`: GetRepositoryHooks1200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetRepositoryHooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryHooksRequest struct via the builder pattern


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


## GetRestriction

> RestRefRestriction GetRestriction(ctx, projectKey, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetRestriction(context.Background(), projectKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestriction`: RestRefRestriction
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The restriction id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestrictionRequest struct via the builder pattern


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


## GetRestrictions

> GetRestrictions1200Response GetRestrictions(ctx, projectKey).MatcherType(matcherType).MatcherId(matcherId).Type_(type_).Start(start).Limit(limit).Execute()

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
	matcherType := "matcherType_example" // string | Matcher type to filter on (optional)
	matcherId := "matcherId_example" // string | Matcher id to filter on. Requires the matcherType parameter to be specified also. (optional)
	type_ := "type__example" // string | Types of restrictions to filter on. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetRestrictions(context.Background(), projectKey).MatcherType(matcherType).MatcherId(matcherId).Type_(type_).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestrictions`: GetRestrictions1200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestrictionsRequest struct via the builder pattern


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


## GetSettings

> ExampleSettings GetSettings(ctx, projectKey, hookKey).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetSettings(context.Background(), projectKey, hookKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings`: ExampleSettings
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The hook key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


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


## GetStatistics

> interface{} GetStatistics(ctx, projectKey, webhookId).Event(event).Execute()

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
	event := "event_example" // string | The string ID of a specific event to retrieve the last invocation for. May be empty, in which case all events are considered (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetStatistics(context.Background(), projectKey, webhookId).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatistics`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | ID of the webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsRequest struct via the builder pattern


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


## GetStatisticsSummary

> interface{} GetStatisticsSummary(ctx, projectKey, webhookId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetStatisticsSummary(context.Background(), projectKey, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetStatisticsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatisticsSummary`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetStatisticsSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | ID of the webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticsSummaryRequest struct via the builder pattern


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


## GetUsersWithAnyPermission1

> GetUsersWithAnyPermission1200Response GetUsersWithAnyPermission1(ctx, projectKey).Filter(filter).Start(start).Limit(limit).Execute()

Get users with permission to project



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
	filter := "filter_example" // string | If specified only user names containing the supplied string will be returned (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetUsersWithAnyPermission1(context.Background(), projectKey).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetUsersWithAnyPermission1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersWithAnyPermission1`: GetUsersWithAnyPermission1200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetUsersWithAnyPermission1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersWithAnyPermission1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | If specified only user names containing the supplied string will be returned | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetUsersWithAnyPermission1200Response**](GetUsersWithAnyPermission1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersWithoutPermission

> GetUsersWithoutAnyPermission200Response GetUsersWithoutPermission(ctx, projectKey).Filter(filter).Start(start).Limit(limit).Execute()

Get users without project permission



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
	filter := "filter_example" // string | If specified only user names containing the supplied string will be returned (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetUsersWithoutPermission(context.Background(), projectKey).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetUsersWithoutPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersWithoutPermission`: GetUsersWithoutAnyPermission200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetUsersWithoutPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersWithoutPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | If specified only user names containing the supplied string will be returned | 
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


## GetWebhook

> RestWebhook GetWebhook(ctx, projectKey, webhookId).Statistics(statistics).Execute()

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
	statistics := "statistics_example" // string | <code>true</code> if statistics should be provided for the webhook (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetWebhook(context.Background(), projectKey, webhookId).Statistics(statistics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: RestWebhook
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | ID of the webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


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


## HasAllUserPermission

> RestPermitted HasAllUserPermission(ctx, projectKey, permission).Execute()

Check default project permission



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
	permission := "permission_example" // string | The permission to grant. Available project permissions are:  - PROJECT_READ - PROJECT_WRITE - PROJECT_ADMIN   

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.HasAllUserPermission(context.Background(), projectKey, permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.HasAllUserPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HasAllUserPermission`: RestPermitted
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.HasAllUserPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**permission** | **string** | The permission to grant. Available project permissions are:  - PROJECT_READ - PROJECT_WRITE - PROJECT_ADMIN    | 

### Other Parameters

Other parameters are passed through a pointer to a apiHasAllUserPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestPermitted**](RestPermitted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAllUserPermission

> ModifyAllUserPermission(ctx, projectKey, permission).Allow(allow).Execute()

Grant project permission



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
	permission := "permission_example" // string | The permission to grant. Available project permissions are:  - PROJECT_READ - PROJECT_WRITE - PROJECT_ADMIN   
	allow := "allow_example" // string | <em>true</em> to grant the specified permission to all users, or <em>false</em> to revoke it (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.ModifyAllUserPermission(context.Background(), projectKey, permission).Allow(allow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ModifyAllUserPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**permission** | **string** | The permission to grant. Available project permissions are:  - PROJECT_READ - PROJECT_WRITE - PROJECT_ADMIN    | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyAllUserPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allow** | **string** | &lt;em&gt;true&lt;/em&gt; to grant the specified permission to all users, or &lt;em&gt;false&lt;/em&gt; to revoke it | 

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


## RemoveConfiguration

> RemoveConfiguration(ctx, projectKey, scriptId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.RemoveConfiguration(context.Background(), projectKey, scriptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.RemoveConfiguration``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveConfigurationRequest struct via the builder pattern


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


## RetryCreateRepository

> RestRepository RetryCreateRepository(ctx, projectKey, repositorySlug).Execute()

Retry repository creation



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
	resp, r, err := apiClient.ProjectAPI.RetryCreateRepository(context.Background(), projectKey, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.RetryCreateRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryCreateRepository`: RestRepository
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.RetryCreateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryCreateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestRepository**](RestRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePermissions

> RevokePermissions(ctx, projectKey).User(user).Group(group).Execute()

Revoke project permissions



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
	user := "user_example" // string | The names of the users (optional)
	group := "group_example" // string | The names of the groups (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.RevokePermissions(context.Background(), projectKey).User(user).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.RevokePermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | **string** | The names of the users | 
 **group** | **string** | The names of the groups | 

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


## RevokePermissionsForGroup1

> RevokePermissionsForGroup1(ctx, projectKey).Name(name).Execute()

Revoke group project permission



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
	name := "name_example" // string | The name of the group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.RevokePermissionsForGroup1(context.Background(), projectKey).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.RevokePermissionsForGroup1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionsForGroup1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the group | 

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


## RevokePermissionsForUser1

> RevokePermissionsForUser1(ctx, projectKey).Name(name).Execute()

Revoke user project permission



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
	name := "name_example" // string | The name of the user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.RevokePermissionsForUser1(context.Background(), projectKey).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.RevokePermissionsForUser1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionsForUser1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the user | 

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


## SearchPermissions

> SearchPermissions(ctx, projectKey).Permission(permission).FilterText(filterText).Type_(type_).Execute()

Search project permissions



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
	permission := "permission_example" // string | Permissions to filter by. See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+project+permissions)for a detailed explanation of what each permission entails. This parameter can be specified multiple times to filter by more than one permission, and can contain global and project permissions.   (optional)
	filterText := "filterText_example" // string | Name of the user or group to filter the name of (optional)
	type_ := "type__example" // string | Type of entity (user or group)Valid entity types are:  - USER- GROUP (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.SearchPermissions(context.Background(), projectKey).Permission(permission).FilterText(filterText).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.SearchPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permission** | **string** | Permissions to filter by. See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+project+permissions)for a detailed explanation of what each permission entails. This parameter can be specified multiple times to filter by more than one permission, and can contain global and project permissions.   | 
 **filterText** | **string** | Name of the user or group to filter the name of | 
 **type_** | **string** | Type of entity (user or group)Valid entity types are:  - USER- GROUP | 

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


## Set

> RestAutoMergeRestrictedSettings Set(ctx, projectKey).RestAutoMergeProjectSettingsRequest(restAutoMergeProjectSettingsRequest).Execute()

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
	restAutoMergeProjectSettingsRequest := *openapiclient.NewRestAutoMergeProjectSettingsRequest() // RestAutoMergeProjectSettingsRequest | The settings to create or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.Set(context.Background(), projectKey).RestAutoMergeProjectSettingsRequest(restAutoMergeProjectSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.Set``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Set`: RestAutoMergeRestrictedSettings
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.Set`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restAutoMergeProjectSettingsRequest** | [**RestAutoMergeProjectSettingsRequest**](RestAutoMergeProjectSettingsRequest.md) | The settings to create or update | 

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


## SetAutoDeclineSettings

> RestAutoDeclineSettings SetAutoDeclineSettings(ctx, projectKey).RestAutoDeclineSettingsRequest(restAutoDeclineSettingsRequest).Execute()

Create/Update auto decline settings



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
	restAutoDeclineSettingsRequest := *openapiclient.NewRestAutoDeclineSettingsRequest() // RestAutoDeclineSettingsRequest | The settings to create or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.SetAutoDeclineSettings(context.Background(), projectKey).RestAutoDeclineSettingsRequest(restAutoDeclineSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.SetAutoDeclineSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAutoDeclineSettings`: RestAutoDeclineSettings
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.SetAutoDeclineSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAutoDeclineSettingsRequest struct via the builder pattern


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


## SetConfiguration

> RestHookScriptConfig SetConfiguration(ctx, projectKey, scriptId).RestHookScriptTriggers(restHookScriptTriggers).Execute()

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
	restHookScriptTriggers := *openapiclient.NewRestHookScriptTriggers() // RestHookScriptTriggers | The hook triggers for which the hook script should be run (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.SetConfiguration(context.Background(), projectKey, scriptId).RestHookScriptTriggers(restHookScriptTriggers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.SetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetConfiguration`: RestHookScriptConfig
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.SetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**scriptId** | **string** | The ID of the hook script | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetConfigurationRequest struct via the builder pattern


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


## SetDefaultBranch2

> SetDefaultBranch2(ctx, projectKey, repositorySlug).RestBranch(restBranch).Execute()

Update default branch for repository



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
	r, err := apiClient.ProjectAPI.SetDefaultBranch2(context.Background(), projectKey, repositorySlug).RestBranch(restBranch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.SetDefaultBranch2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetDefaultBranch2Request struct via the builder pattern


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


## SetPermissionForGroups1

> SetPermissionForGroups1(ctx, projectKey).Name(name).Permission(permission).Execute()

Update group project permission



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
	name := "name_example" // string | The names of the groups (optional)
	permission := "permission_example" // string | The permission to grant.See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+project+permissions)for a detailed explanation of what each permission entails. Available project permissions are:  - PROJECT_READ - PROJECT_WRITE - PROJECT_ADMIN    (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.SetPermissionForGroups1(context.Background(), projectKey).Name(name).Permission(permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.SetPermissionForGroups1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPermissionForGroups1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The names of the groups | 
 **permission** | **string** | The permission to grant.See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+project+permissions)for a detailed explanation of what each permission entails. Available project permissions are:  - PROJECT_READ - PROJECT_WRITE - PROJECT_ADMIN    | 

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


## SetPermissionForUsers1

> SetPermissionForUsers1(ctx, projectKey).Name(name).Permission(permission).Execute()

Update user project permission



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
	name := "name_example" // string | The names of the users (optional)
	permission := "permission_example" // string | The permission to grant.See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+project+permissions)for a detailed explanation of what each permission entails. Available project permissions are:  - PROJECT_READ - PROJECT_WRITE - PROJECT_ADMIN    (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.SetPermissionForUsers1(context.Background(), projectKey).Name(name).Permission(permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.SetPermissionForUsers1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPermissionForUsers1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The names of the users | 
 **permission** | **string** | The permission to grant.See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+project+permissions)for a detailed explanation of what each permission entails. Available project permissions are:  - PROJECT_READ - PROJECT_WRITE - PROJECT_ADMIN    | 

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


## SetSettings

> ExampleSettings SetSettings(ctx, projectKey, hookKey).ExampleSettings(exampleSettings).Execute()

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
	hookKey := "hookKey_example" // string | The complete module key of the hook module.
	exampleSettings := *openapiclient.NewExampleSettings() // ExampleSettings | The raw settings. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.SetSettings(context.Background(), projectKey, hookKey).ExampleSettings(exampleSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.SetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSettings`: ExampleSettings
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.SetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**hookKey** | **string** | The complete module key of the hook module. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSettingsRequest struct via the builder pattern


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


## StreamContributing

> StreamContributing(ctx, projectKey, repositorySlug).At(at).Markup(markup).HtmlEscape(htmlEscape).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Execute()

Get repository contributing guidelines



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
	at := "at_example" // string | A specific commit or ref to retrieve the guidelines at, or the default branch if not specified (optional)
	markup := "markup_example" // string | If present or <code>\"true\"</code>, triggers the raw content to be markup-rendered and returned as HTML; otherwise, if not specified, or any value other than <code>\"true\"</code>, the content is streamed without markup (optional)
	htmlEscape := "htmlEscape_example" // string | (Optional) true if HTML should be escaped in the input markup, false otherwise. If not specified, the value of the <code>markup.render.html.escape</code> property, which is <code>true</code> by default, will be used (optional)
	includeHeadingId := "includeHeadingId_example" // string | (Optional) true if headings should contain an ID based on the heading content. If not specified, the value of the <code>markup.render.headerids</code> property, which is false by default, will be used (optional)
	hardwrap := "hardwrap_example" // string | (Optional) Whether the markup implementation should convert newlines to breaks. If not specified, the value of the <code>markup.render.hardwrap</code> property, which is <code>true</code> by default, will be used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.StreamContributing(context.Background(), projectKey, repositorySlug).At(at).Markup(markup).HtmlEscape(htmlEscape).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.StreamContributing``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStreamContributingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **at** | **string** | A specific commit or ref to retrieve the guidelines at, or the default branch if not specified | 
 **markup** | **string** | If present or &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, triggers the raw content to be markup-rendered and returned as HTML; otherwise, if not specified, or any value other than &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, the content is streamed without markup | 
 **htmlEscape** | **string** | (Optional) true if HTML should be escaped in the input markup, false otherwise. If not specified, the value of the &lt;code&gt;markup.render.html.escape&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default, will be used | 
 **includeHeadingId** | **string** | (Optional) true if headings should contain an ID based on the heading content. If not specified, the value of the &lt;code&gt;markup.render.headerids&lt;/code&gt; property, which is false by default, will be used | 
 **hardwrap** | **string** | (Optional) Whether the markup implementation should convert newlines to breaks. If not specified, the value of the &lt;code&gt;markup.render.hardwrap&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default, will be used | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamLicense

> StreamLicense(ctx, projectKey, repositorySlug).At(at).Markup(markup).HtmlEscape(htmlEscape).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Execute()

Get repository license



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
	at := "at_example" // string | A specific commit or ref to retrieve the guidelines at, or the default branch if not specified (optional)
	markup := "markup_example" // string | If present or <code>\"true\"</code>, triggers the raw content to be markup-rendered and returned as HTML; otherwise, if not specified, or any value other than <code>\"true\"</code>, the content is streamed without markup (optional)
	htmlEscape := "htmlEscape_example" // string | (Optional) true if HTML should be escaped in the input markup, false otherwise. If not specified, the value of the <code>markup.render.html.escape</code> property, which is <code>true</code> by default, will be used (optional)
	includeHeadingId := "includeHeadingId_example" // string | (Optional) true if headings should contain an ID based on the heading content. If not specified, the value of the <code>markup.render.headerids</code> property, which is false by default, will be used (optional)
	hardwrap := "hardwrap_example" // string | (Optional) Whether the markup implementation should convert newlines to breaks. If not specified, the value of the <code>markup.render.hardwrap</code> property, which is <code>true</code> by default, will be used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.StreamLicense(context.Background(), projectKey, repositorySlug).At(at).Markup(markup).HtmlEscape(htmlEscape).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.StreamLicense``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStreamLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **at** | **string** | A specific commit or ref to retrieve the guidelines at, or the default branch if not specified | 
 **markup** | **string** | If present or &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, triggers the raw content to be markup-rendered and returned as HTML; otherwise, if not specified, or any value other than &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, the content is streamed without markup | 
 **htmlEscape** | **string** | (Optional) true if HTML should be escaped in the input markup, false otherwise. If not specified, the value of the &lt;code&gt;markup.render.html.escape&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default, will be used | 
 **includeHeadingId** | **string** | (Optional) true if headings should contain an ID based on the heading content. If not specified, the value of the &lt;code&gt;markup.render.headerids&lt;/code&gt; property, which is false by default, will be used | 
 **hardwrap** | **string** | (Optional) Whether the markup implementation should convert newlines to breaks. If not specified, the value of the &lt;code&gt;markup.render.hardwrap&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default, will be used | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamReadme

> StreamReadme(ctx, projectKey, repositorySlug).At(at).Markup(markup).HtmlEscape(htmlEscape).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Execute()

Get repository readme



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
	at := "at_example" // string | A specific commit or ref to retrieve the guidelines at, or the default branch if not specified (optional)
	markup := "markup_example" // string | If present or <code>\"true\"</code>, triggers the raw content to be markup-rendered and returned as HTML; otherwise, if not specified, or any value other than <code>\"true\"</code>, the content is streamed without markup (optional)
	htmlEscape := "htmlEscape_example" // string | (Optional) true if HTML should be escaped in the input markup, false otherwise. If not specified, the value of the <code>markup.render.html.escape</code> property, which is <code>true</code> by default, will be used (optional)
	includeHeadingId := "includeHeadingId_example" // string | (Optional) true if headings should contain an ID based on the heading content. If not specified, the value of the <code>markup.render.headerids</code> property, which is false by default, will be used (optional)
	hardwrap := "hardwrap_example" // string | (Optional) Whether the markup implementation should convert newlines to breaks. If not specified, the value of the <code>markup.render.hardwrap</code> property, which is <code>true</code> by default, will be used (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.StreamReadme(context.Background(), projectKey, repositorySlug).At(at).Markup(markup).HtmlEscape(htmlEscape).IncludeHeadingId(includeHeadingId).Hardwrap(hardwrap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.StreamReadme``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStreamReadmeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **at** | **string** | A specific commit or ref to retrieve the guidelines at, or the default branch if not specified | 
 **markup** | **string** | If present or &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, triggers the raw content to be markup-rendered and returned as HTML; otherwise, if not specified, or any value other than &lt;code&gt;\&quot;true\&quot;&lt;/code&gt;, the content is streamed without markup | 
 **htmlEscape** | **string** | (Optional) true if HTML should be escaped in the input markup, false otherwise. If not specified, the value of the &lt;code&gt;markup.render.html.escape&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default, will be used | 
 **includeHeadingId** | **string** | (Optional) true if headings should contain an ID based on the heading content. If not specified, the value of the &lt;code&gt;markup.render.headerids&lt;/code&gt; property, which is false by default, will be used | 
 **hardwrap** | **string** | (Optional) Whether the markup implementation should convert newlines to breaks. If not specified, the value of the &lt;code&gt;markup.render.hardwrap&lt;/code&gt; property, which is &lt;code&gt;true&lt;/code&gt; by default, will be used | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWebhook

> interface{} TestWebhook(ctx, projectKey).WebhookId(webhookId).SslVerificationRequired(sslVerificationRequired).Url(url).RestWebhookCredentials(restWebhookCredentials).Execute()

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
	webhookId := int32(56) // int32 |  (optional)
	sslVerificationRequired := true // bool |  (optional) (default to true)
	url := "url_example" // string | The url in which to connect to (optional)
	restWebhookCredentials := *openapiclient.NewRestWebhookCredentials() // RestWebhookCredentials | Basic authentication credentials, if required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.TestWebhook(context.Background(), projectKey).WebhookId(webhookId).SslVerificationRequired(sslVerificationRequired).Url(url).RestWebhookCredentials(restWebhookCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.TestWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestWebhook`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.TestWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookId** | **int32** |  | 
 **sslVerificationRequired** | **bool** |  | [default to true]
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


## UpdateDefaultTask

> RestDefaultTask UpdateDefaultTask(ctx, projectKey, taskId).RestDefaultTaskRequest(restDefaultTaskRequest).Execute()

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
	taskId := "taskId_example" // string | The ID of the default task
	restDefaultTaskRequest := *openapiclient.NewRestDefaultTaskRequest() // RestDefaultTaskRequest | The task to be updated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.UpdateDefaultTask(context.Background(), projectKey, taskId).RestDefaultTaskRequest(restDefaultTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.UpdateDefaultTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDefaultTask`: RestDefaultTask
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.UpdateDefaultTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**taskId** | **string** | The ID of the default task | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultTaskRequest struct via the builder pattern


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


## UpdateProject

> RestProject UpdateProject(ctx, projectKey).RestProject(restProject).Execute()

Update project



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
	restProject := *openapiclient.NewRestProject() // RestProject | Project parameters to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.UpdateProject(context.Background(), projectKey).RestProject(restProject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.UpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProject`: RestProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restProject** | [**RestProject**](RestProject.md) | Project parameters to update. | 

### Return type

[**RestProject**](RestProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePullRequestSettings

> RestPullRequestSettings UpdatePullRequestSettings(ctx, projectKey, scmId).RestPullRequestSettings(restPullRequestSettings).Execute()

Update merge strategy



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
	scmId := "scmId_example" // string | The SCM to get strategies for.
	restPullRequestSettings := *openapiclient.NewRestPullRequestSettings() // RestPullRequestSettings | The settings. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.UpdatePullRequestSettings(context.Background(), projectKey, scmId).RestPullRequestSettings(restPullRequestSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.UpdatePullRequestSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePullRequestSettings`: RestPullRequestSettings
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.UpdatePullRequestSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**scmId** | **string** | The SCM to get strategies for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePullRequestSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restPullRequestSettings** | [**RestPullRequestSettings**](RestPullRequestSettings.md) | The settings. | 

### Return type

[**RestPullRequestSettings**](RestPullRequestSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepository

> RestRepository UpdateRepository(ctx, projectKey, repositorySlug).RestRepository(restRepository).Execute()

Update repository



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
	restRepository := *openapiclient.NewRestRepository() // RestRepository | The updated repository. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.UpdateRepository(context.Background(), projectKey, repositorySlug).RestRepository(restRepository).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.UpdateRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRepository`: RestRepository
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.UpdateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restRepository** | [**RestRepository**](RestRepository.md) | The updated repository. | 

### Return type

[**RestRepository**](RestRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> RestWebhook UpdateWebhook(ctx, projectKey, webhookId).RestWebhook(restWebhook).Execute()

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
	restWebhook := *openapiclient.NewRestWebhook() // RestWebhook | The representation of the updated values for the webhook (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.UpdateWebhook(context.Background(), projectKey, webhookId).RestWebhook(restWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: RestWebhook
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**webhookId** | **string** | Id of the existing webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


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


## UploadAvatar

> UploadAvatar(ctx, projectKey).Avatar(avatar).Execute()

Update project avatar



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
	avatar := os.NewFile(1234, "some_file") // *os.File | The avatar file to upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.UploadAvatar(context.Background(), projectKey).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.UploadAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **avatar** | ***os.File** | The avatar file to upload. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


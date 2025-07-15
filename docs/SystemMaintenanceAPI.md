# \SystemMaintenanceAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelExportJob**](SystemMaintenanceAPI.md#CancelExportJob) | **Post** /api/latest/migration/exports/{jobId}/cancel | Cancel export job
[**CancelImportJob**](SystemMaintenanceAPI.md#CancelImportJob) | **Post** /api/latest/migration/imports/{jobId}/cancel | Cancel import job
[**CancelMeshMigrationJob**](SystemMaintenanceAPI.md#CancelMeshMigrationJob) | **Post** /api/latest/migration/mesh/{jobId}/cancel | Cancel Mesh migration job
[**ClearDefaultBranch**](SystemMaintenanceAPI.md#ClearDefaultBranch) | **Delete** /api/latest/admin/default-branch | Clear default branch
[**ClearSenderAddress**](SystemMaintenanceAPI.md#ClearSenderAddress) | **Delete** /api/latest/admin/mail-server/sender-address | Update mail configuration
[**Connectivity**](SystemMaintenanceAPI.md#Connectivity) | **Get** /api/latest/admin/git/mesh/diagnostics/connectivity | Generate Mesh connectivity report
[**CreateHookScript**](SystemMaintenanceAPI.md#CreateHookScript) | **Post** /api/latest/hook-scripts | Create a new hook script
[**Delete2**](SystemMaintenanceAPI.md#Delete2) | **Delete** /api/latest/admin/git/mesh/nodes/{id} | Delete Mesh node
[**Delete8**](SystemMaintenanceAPI.md#Delete8) | **Delete** /api/latest/admin/rate-limit/settings/users/{userSlug} | Delete user specific rate limit settings
[**DeleteAvatar**](SystemMaintenanceAPI.md#DeleteAvatar) | **Delete** /api/latest/users/{userSlug}/avatar.png | Delete user avatar
[**DeleteBanner**](SystemMaintenanceAPI.md#DeleteBanner) | **Delete** /api/latest/admin/banner | Delete announcement banner
[**DeleteHookScript**](SystemMaintenanceAPI.md#DeleteHookScript) | **Delete** /api/latest/hook-scripts/{scriptId} | Delete a hook script.
[**DeleteMailConfig**](SystemMaintenanceAPI.md#DeleteMailConfig) | **Delete** /api/latest/admin/mail-server | Delete mail configuration
[**DismissRetentionConfigReviewNotification**](SystemMaintenanceAPI.md#DismissRetentionConfigReviewNotification) | **Delete** /audit/latest/notification-settings/retention-config-review | Dismiss retention config notification
[**Get2**](SystemMaintenanceAPI.md#Get2) | **Get** /api/latest/admin/license | Get license details
[**Get6**](SystemMaintenanceAPI.md#Get6) | **Get** /api/latest/admin/rate-limit/settings/users/{userSlug} | Get user specific rate limit settings
[**GetActiveMeshMigrationSummary**](SystemMaintenanceAPI.md#GetActiveMeshMigrationSummary) | **Get** /api/latest/migration/mesh/summary | Get summary for Mesh migration job
[**GetAllMeshMigrationSummaries**](SystemMaintenanceAPI.md#GetAllMeshMigrationSummaries) | **Get** /api/latest/migration/mesh/summaries | Get all Mesh migration job summaries
[**GetAllRateLimitSettings**](SystemMaintenanceAPI.md#GetAllRateLimitSettings) | **Get** /api/latest/admin/rate-limit/settings/users | Get rate limit settings for user
[**GetAllRegisteredMeshNodes**](SystemMaintenanceAPI.md#GetAllRegisteredMeshNodes) | **Get** /api/latest/admin/git/mesh/nodes | Get all registered Mesh nodes
[**GetApplicationProperties**](SystemMaintenanceAPI.md#GetApplicationProperties) | **Get** /api/latest/application-properties | Get application properties
[**GetBanner**](SystemMaintenanceAPI.md#GetBanner) | **Get** /api/latest/admin/banner | Get announcement banner
[**GetControlPlanePublicKey**](SystemMaintenanceAPI.md#GetControlPlanePublicKey) | **Get** /api/latest/admin/git/mesh/config/control-plane.pem | Get the control plane PEM
[**GetDefaultBranch**](SystemMaintenanceAPI.md#GetDefaultBranch) | **Get** /api/latest/admin/default-branch | Get the default branch
[**GetExportJob**](SystemMaintenanceAPI.md#GetExportJob) | **Get** /api/latest/migration/exports/{jobId} | Get export job details
[**GetExportJobMessages**](SystemMaintenanceAPI.md#GetExportJobMessages) | **Get** /api/latest/migration/exports/{jobId}/messages | Get job messages
[**GetGlobalSettings**](SystemMaintenanceAPI.md#GetGlobalSettings) | **Get** /admin | Get global SSH key settings
[**GetHistory**](SystemMaintenanceAPI.md#GetHistory) | **Get** /api/latest/admin/rate-limit/history | Get rate limit history
[**GetHookScript**](SystemMaintenanceAPI.md#GetHookScript) | **Get** /api/latest/hook-scripts/{scriptId} | Get a hook script
[**GetImportJob**](SystemMaintenanceAPI.md#GetImportJob) | **Get** /api/latest/migration/imports/{jobId} | Get import job status
[**GetImportJobMessages**](SystemMaintenanceAPI.md#GetImportJobMessages) | **Get** /api/latest/migration/imports/{jobId}/messages | Get import job messages
[**GetInformation**](SystemMaintenanceAPI.md#GetInformation) | **Get** /api/latest/admin/cluster | Get cluster node information
[**GetLabel**](SystemMaintenanceAPI.md#GetLabel) | **Get** /api/latest/labels/{labelName} | Get label
[**GetLabelables**](SystemMaintenanceAPI.md#GetLabelables) | **Get** /api/latest/labels/{labelName}/labeled | Get labelables for label
[**GetLabels**](SystemMaintenanceAPI.md#GetLabels) | **Get** /api/latest/labels | Get all labels
[**GetLevel**](SystemMaintenanceAPI.md#GetLevel) | **Get** /api/latest/logs/logger/{loggerName} | Get current log level
[**GetMailConfig**](SystemMaintenanceAPI.md#GetMailConfig) | **Get** /api/latest/admin/mail-server | Get mail configuration
[**GetMeshMigrationJob**](SystemMaintenanceAPI.md#GetMeshMigrationJob) | **Get** /api/latest/migration/mesh/{jobId} | Get Mesh migration job details
[**GetMeshMigrationJobMessages**](SystemMaintenanceAPI.md#GetMeshMigrationJobMessages) | **Get** /api/latest/migration/mesh/{jobId}/messages | Get Mesh migration job messages
[**GetMeshMigrationJobSummary**](SystemMaintenanceAPI.md#GetMeshMigrationJobSummary) | **Get** /api/latest/migration/mesh/{jobId}/summary | Get Mesh migration job summary
[**GetRegisteredMeshNodeById**](SystemMaintenanceAPI.md#GetRegisteredMeshNodeById) | **Get** /api/latest/admin/git/mesh/nodes/{id} | Get Mesh node
[**GetRepositoryArchivePolicy**](SystemMaintenanceAPI.md#GetRepositoryArchivePolicy) | **Get** /policies/latest/admin/repos/archive | Get repository archive policy
[**GetRepositoryDeletePolicy**](SystemMaintenanceAPI.md#GetRepositoryDeletePolicy) | **Get** /policies/latest/admin/repos/delete | Get repository delete policy
[**GetRootLevel**](SystemMaintenanceAPI.md#GetRootLevel) | **Get** /api/latest/logs/rootLogger | Get root log level
[**GetSenderAddress**](SystemMaintenanceAPI.md#GetSenderAddress) | **Get** /api/latest/admin/mail-server/sender-address | Get server mail address
[**GetSettings2**](SystemMaintenanceAPI.md#GetSettings2) | **Get** /api/latest/logs/settings | Get debug logging and profiling
[**GetSettings3**](SystemMaintenanceAPI.md#GetSettings3) | **Get** /api/latest/admin/rate-limit/settings | Get rate limit settings
[**GetSupportZip**](SystemMaintenanceAPI.md#GetSupportZip) | **Get** /api/latest/admin/git/mesh/support-zips/{id} | Get support zip for node
[**GetSupportZips**](SystemMaintenanceAPI.md#GetSupportZips) | **Get** /api/latest/admin/git/mesh/support-zips | Get support zips for all Mesh nodes
[**GetSupportedKeyTypes**](SystemMaintenanceAPI.md#GetSupportedKeyTypes) | **Get** /admin/supported-key-types | Get supported SSH key algorithms and lengths
[**GetUser**](SystemMaintenanceAPI.md#GetUser) | **Get** /api/latest/users/{userSlug} | Get user
[**GetUserSettings**](SystemMaintenanceAPI.md#GetUserSettings) | **Get** /api/latest/users/{userSlug}/settings | Get user settings
[**GetUsers2**](SystemMaintenanceAPI.md#GetUsers2) | **Get** /api/latest/users | Get all users
[**PreviewExport**](SystemMaintenanceAPI.md#PreviewExport) | **Post** /api/latest/migration/exports/preview | Preview export
[**PreviewMeshMigration**](SystemMaintenanceAPI.md#PreviewMeshMigration) | **Post** /api/latest/migration/mesh/preview | Preview Mesh migration
[**Read**](SystemMaintenanceAPI.md#Read) | **Get** /api/latest/hook-scripts/{scriptId}/content | Get hook script content
[**RegisterNewMeshNode**](SystemMaintenanceAPI.md#RegisterNewMeshNode) | **Post** /api/latest/admin/git/mesh/nodes | Register new Mesh node
[**SearchMeshMigrationRepos**](SystemMaintenanceAPI.md#SearchMeshMigrationRepos) | **Get** /api/latest/migration/mesh/repos | Find repositories by Mesh migration state
[**Set2**](SystemMaintenanceAPI.md#Set2) | **Post** /api/latest/admin/rate-limit/settings/users | Set rate limit settings for users
[**Set3**](SystemMaintenanceAPI.md#Set3) | **Put** /api/latest/admin/rate-limit/settings/users/{userSlug} | Set rate limit settings for user
[**SetBanner**](SystemMaintenanceAPI.md#SetBanner) | **Put** /api/latest/admin/banner | Update/Set announcement banner
[**SetDefaultBranch**](SystemMaintenanceAPI.md#SetDefaultBranch) | **Put** /api/latest/admin/default-branch | Update/Set default branch
[**SetLevel**](SystemMaintenanceAPI.md#SetLevel) | **Put** /api/latest/logs/logger/{loggerName}/{levelName} | Set log level
[**SetMailConfig**](SystemMaintenanceAPI.md#SetMailConfig) | **Put** /api/latest/admin/mail-server | Update mail configuration
[**SetRepositoryArchivePolicy**](SystemMaintenanceAPI.md#SetRepositoryArchivePolicy) | **Put** /policies/latest/admin/repos/archive | Update repository archive policy
[**SetRepositoryDeletePolicy**](SystemMaintenanceAPI.md#SetRepositoryDeletePolicy) | **Put** /policies/latest/admin/repos/delete | Update the repository delete policy
[**SetRootLevel**](SystemMaintenanceAPI.md#SetRootLevel) | **Put** /api/latest/logs/rootLogger/{levelName} | Set root log level
[**SetSenderAddress**](SystemMaintenanceAPI.md#SetSenderAddress) | **Put** /api/latest/admin/mail-server/sender-address | Update server mail address
[**SetSettings2**](SystemMaintenanceAPI.md#SetSettings2) | **Put** /api/latest/logs/settings | Set debug logging and profiling
[**SetSettings3**](SystemMaintenanceAPI.md#SetSettings3) | **Put** /api/latest/admin/rate-limit/settings | Set rate limit
[**StartExport**](SystemMaintenanceAPI.md#StartExport) | **Post** /api/latest/migration/exports | Start export job
[**StartImport**](SystemMaintenanceAPI.md#StartImport) | **Post** /api/latest/migration/imports | Start import job
[**StartMeshMigration**](SystemMaintenanceAPI.md#StartMeshMigration) | **Post** /api/latest/migration/mesh | Start Mesh migration job
[**UpdateGlobalSettings**](SystemMaintenanceAPI.md#UpdateGlobalSettings) | **Put** /admin | Update global SSH key settings
[**UpdateHookScript**](SystemMaintenanceAPI.md#UpdateHookScript) | **Put** /api/latest/hook-scripts/{scriptId} | Update a hook script
[**UpdateLicense**](SystemMaintenanceAPI.md#UpdateLicense) | **Post** /api/latest/admin/license | Update license
[**UpdateMeshNode**](SystemMaintenanceAPI.md#UpdateMeshNode) | **Put** /api/latest/admin/git/mesh/nodes/{id} | Update Mesh node
[**UpdateSettings**](SystemMaintenanceAPI.md#UpdateSettings) | **Post** /api/latest/users/{userSlug}/settings | Update user settings
[**UpdateUserDetails1**](SystemMaintenanceAPI.md#UpdateUserDetails1) | **Put** /api/latest/users | Update user details
[**UpdateUserPassword1**](SystemMaintenanceAPI.md#UpdateUserPassword1) | **Put** /api/latest/users/credentials | Set password
[**UploadAvatar1**](SystemMaintenanceAPI.md#UploadAvatar1) | **Post** /api/latest/users/{userSlug}/avatar.png | Update user avatar



## CancelExportJob

> CancelExportJob(ctx, jobId).Execute()

Cancel export job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | the ID of the job to cancel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.CancelExportJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.CancelExportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | the ID of the job to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelExportJobRequest struct via the builder pattern


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


## CancelImportJob

> CancelImportJob(ctx, jobId).Execute()

Cancel import job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | the ID of the job to cancel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.CancelImportJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.CancelImportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | the ID of the job to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelImportJobRequest struct via the builder pattern


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


## CancelMeshMigrationJob

> CancelMeshMigrationJob(ctx, jobId).Execute()

Cancel Mesh migration job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | The ID of the job to cancel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.CancelMeshMigrationJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.CancelMeshMigrationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelMeshMigrationJobRequest struct via the builder pattern


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


## ClearDefaultBranch

> ClearDefaultBranch(ctx).Execute()

Clear default branch



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.ClearDefaultBranch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.ClearDefaultBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearDefaultBranchRequest struct via the builder pattern


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


## ClearSenderAddress

> ClearSenderAddress(ctx).Execute()

Update mail configuration



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.ClearSenderAddress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.ClearSenderAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearSenderAddressRequest struct via the builder pattern


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


## Connectivity

> RestMeshConnectivityReport Connectivity(ctx).Execute()

Generate Mesh connectivity report



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.Connectivity(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.Connectivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Connectivity`: RestMeshConnectivityReport
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.Connectivity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConnectivityRequest struct via the builder pattern


### Return type

[**RestMeshConnectivityReport**](RestMeshConnectivityReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHookScript

> RestHookScript CreateHookScript(ctx).Content(content).Description(description).Name(name).Type_(type_).Execute()

Create a new hook script



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	content := "content_example" // string | The hook script contents. (optional)
	description := "description_example" // string | A description of the hook script (useful when querying registered hook scripts). (optional)
	name := "name_example" // string | The name of the hook script (useful when querying registered hook scripts). (optional)
	type_ := "type__example" // string | The type of hook script; supported values are \\\"PRE\\\" for pre-receive hooks and \\\"POST\\\" for post-receive hooks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.CreateHookScript(context.Background()).Content(content).Description(description).Name(name).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.CreateHookScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHookScript`: RestHookScript
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.CreateHookScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHookScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **content** | **string** | The hook script contents. | 
 **description** | **string** | A description of the hook script (useful when querying registered hook scripts). | 
 **name** | **string** | The name of the hook script (useful when querying registered hook scripts). | 
 **type_** | **string** | The type of hook script; supported values are \\\&quot;PRE\\\&quot; for pre-receive hooks and \\\&quot;POST\\\&quot; for post-receive hooks. | 

### Return type

[**RestHookScript**](RestHookScript.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete2

> Delete2(ctx, id).Force(force).Execute()

Delete Mesh node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := int64(789) // int64 | 
	force := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.Delete2(context.Background(), id).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.Delete2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** |  | [default to false]

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


## Delete8

> Delete8(ctx, userSlug).Execute()

Delete user specific rate limit settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userSlug := "userSlug_example" // string | The user slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.Delete8(context.Background(), userSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.Delete8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete8Request struct via the builder pattern


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


## DeleteAvatar

> RestNamedLink DeleteAvatar(ctx, userSlug).Execute()

Delete user avatar



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userSlug := "userSlug_example" // string | The user slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.DeleteAvatar(context.Background(), userSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.DeleteAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAvatar`: RestNamedLink
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.DeleteAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestNamedLink**](RestNamedLink.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBanner

> DeleteBanner(ctx).Execute()

Delete announcement banner



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.DeleteBanner(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.DeleteBanner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBannerRequest struct via the builder pattern


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


## DeleteHookScript

> DeleteHookScript(ctx, scriptId).Execute()

Delete a hook script.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	scriptId := "scriptId_example" // string | The ID of the hook script to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.DeleteHookScript(context.Background(), scriptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.DeleteHookScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptId** | **string** | The ID of the hook script to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHookScriptRequest struct via the builder pattern


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


## DeleteMailConfig

> DeleteMailConfig(ctx).Execute()

Delete mail configuration



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.DeleteMailConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.DeleteMailConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailConfigRequest struct via the builder pattern


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


## DismissRetentionConfigReviewNotification

> DismissRetentionConfigReviewNotification(ctx).Execute()

Dismiss retention config notification



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.DismissRetentionConfigReviewNotification(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.DismissRetentionConfigReviewNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDismissRetentionConfigReviewNotificationRequest struct via the builder pattern


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


## Get2

> RestBitbucketLicense Get2(ctx).Execute()

Get license details



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.Get2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.Get2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get2`: RestBitbucketLicense
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.Get2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGet2Request struct via the builder pattern


### Return type

[**RestBitbucketLicense**](RestBitbucketLicense.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get6

> RestUserRateLimitSettings Get6(ctx, userSlug).Execute()

Get user specific rate limit settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userSlug := "userSlug_example" // string | The user slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.Get6(context.Background(), userSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.Get6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get6`: RestUserRateLimitSettings
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.Get6`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestUserRateLimitSettings**](RestUserRateLimitSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveMeshMigrationSummary

> RestMeshMigrationSummary GetActiveMeshMigrationSummary(ctx).Execute()

Get summary for Mesh migration job



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetActiveMeshMigrationSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetActiveMeshMigrationSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveMeshMigrationSummary`: RestMeshMigrationSummary
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetActiveMeshMigrationSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveMeshMigrationSummaryRequest struct via the builder pattern


### Return type

[**RestMeshMigrationSummary**](RestMeshMigrationSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllMeshMigrationSummaries

> GetAllMeshMigrationSummaries200Response GetAllMeshMigrationSummaries(ctx).Start(start).Limit(limit).Execute()

Get all Mesh migration job summaries



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetAllMeshMigrationSummaries(context.Background()).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetAllMeshMigrationSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMeshMigrationSummaries`: GetAllMeshMigrationSummaries200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetAllMeshMigrationSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMeshMigrationSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetAllMeshMigrationSummaries200Response**](GetAllMeshMigrationSummaries200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRateLimitSettings

> GetAllRateLimitSettings200Response GetAllRateLimitSettings(ctx).Filter(filter).Start(start).Limit(limit).Execute()

Get rate limit settings for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "filter_example" // string | Optional filter (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetAllRateLimitSettings(context.Background()).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetAllRateLimitSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllRateLimitSettings`: GetAllRateLimitSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetAllRateLimitSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRateLimitSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Optional filter | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetAllRateLimitSettings200Response**](GetAllRateLimitSettings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRegisteredMeshNodes

> RestMeshNode GetAllRegisteredMeshNodes(ctx).Execute()

Get all registered Mesh nodes



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetAllRegisteredMeshNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetAllRegisteredMeshNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllRegisteredMeshNodes`: RestMeshNode
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetAllRegisteredMeshNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRegisteredMeshNodesRequest struct via the builder pattern


### Return type

[**RestMeshNode**](RestMeshNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationProperties

> RestApplicationProperties GetApplicationProperties(ctx).Execute()

Get application properties



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetApplicationProperties(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetApplicationProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationProperties`: RestApplicationProperties
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetApplicationProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPropertiesRequest struct via the builder pattern


### Return type

[**RestApplicationProperties**](RestApplicationProperties.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBanner

> RestAnnouncementBanner GetBanner(ctx).Execute()

Get announcement banner



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetBanner(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetBanner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBanner`: RestAnnouncementBanner
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetBanner`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBannerRequest struct via the builder pattern


### Return type

[**RestAnnouncementBanner**](RestAnnouncementBanner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetControlPlanePublicKey

> GetControlPlanePublicKey(ctx).Execute()

Get the control plane PEM



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.GetControlPlanePublicKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetControlPlanePublicKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetControlPlanePublicKeyRequest struct via the builder pattern


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


## GetDefaultBranch

> GetDefaultBranch(ctx).Execute()

Get the default branch



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.GetDefaultBranch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetDefaultBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultBranchRequest struct via the builder pattern


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


## GetExportJob

> RestJob GetExportJob(ctx, jobId).Execute()

Get export job details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | the ID of the job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetExportJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetExportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExportJob`: RestJob
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetExportJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | the ID of the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestJob**](RestJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportJobMessages

> GetExportJobMessages200Response GetExportJobMessages(ctx, jobId).Severity(severity).Subject(subject).Start(start).Limit(limit).Execute()

Get job messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | The ID of the job
	severity := "severity_example" // string | The severity to include in the results (optional)
	subject := "subject_example" // string | The subject (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetExportJobMessages(context.Background(), jobId).Severity(severity).Subject(subject).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetExportJobMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExportJobMessages`: GetExportJobMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetExportJobMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportJobMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **severity** | **string** | The severity to include in the results | 
 **subject** | **string** | The subject | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetExportJobMessages200Response**](GetExportJobMessages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalSettings

> RestSshKeySettings GetGlobalSettings(ctx).Execute()

Get global SSH key settings



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetGlobalSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetGlobalSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalSettings`: RestSshKeySettings
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetGlobalSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalSettingsRequest struct via the builder pattern


### Return type

[**RestSshKeySettings**](RestSshKeySettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistory

> GetHistory200Response GetHistory(ctx).Order(order).Start(start).Limit(limit).Execute()

Get rate limit history



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	order := "order_example" // string | An optional sort category to arrange the results in descending order (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetHistory(context.Background()).Order(order).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistory`: GetHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | An optional sort category to arrange the results in descending order | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetHistory200Response**](GetHistory200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHookScript

> RestHookScript GetHookScript(ctx, scriptId).Execute()

Get a hook script



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	scriptId := "scriptId_example" // string | The ID of the hook script to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetHookScript(context.Background(), scriptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetHookScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHookScript`: RestHookScript
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetHookScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptId** | **string** | The ID of the hook script to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHookScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestHookScript**](RestHookScript.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportJob

> RestJob GetImportJob(ctx, jobId).Execute()

Get import job status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | The ID of the job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetImportJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetImportJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportJob`: RestJob
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetImportJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestJob**](RestJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportJobMessages

> GetExportJobMessages200Response GetImportJobMessages(ctx, jobId).Severity(severity).Subject(subject).Start(start).Limit(limit).Execute()

Get import job messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | The ID of the job
	severity := "severity_example" // string | The severity to include in the results (optional)
	subject := "subject_example" // string | The subject (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetImportJobMessages(context.Background(), jobId).Severity(severity).Subject(subject).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetImportJobMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportJobMessages`: GetExportJobMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetImportJobMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportJobMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **severity** | **string** | The severity to include in the results | 
 **subject** | **string** | The subject | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetExportJobMessages200Response**](GetExportJobMessages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInformation

> RestClusterInformation GetInformation(ctx).Execute()

Get cluster node information



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetInformation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInformation`: RestClusterInformation
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInformationRequest struct via the builder pattern


### Return type

[**RestClusterInformation**](RestClusterInformation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabel

> RestLabel GetLabel(ctx, labelName).Execute()

Get label



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	labelName := "labelName_example" // string | the label name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetLabel(context.Background(), labelName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabel`: RestLabel
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelName** | **string** | the label name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelRequest struct via the builder pattern


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


## GetLabelables

> GetLabelables200Response GetLabelables(ctx, labelName).Type_(type_).Start(start).Limit(limit).Execute()

Get labelables for label



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	labelName := "labelName_example" // string | The page of labelables.
	type_ := "type__example" // string |  the type of labelables to be returned. Supported values: REPOSITORY (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetLabelables(context.Background(), labelName).Type_(type_).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetLabelables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabelables`: GetLabelables200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetLabelables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelName** | **string** | The page of labelables. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** |  the type of labelables to be returned. Supported values: REPOSITORY | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetLabelables200Response**](GetLabelables200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabels

> GetLabels200Response GetLabels(ctx).Prefix(prefix).Start(start).Limit(limit).Execute()

Get all labels



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	prefix := "prefix_example" // string | (optional) prefix to filter the labels on. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetLabels(context.Background()).Prefix(prefix).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabels`: GetLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** | (optional) prefix to filter the labels on. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetLabels200Response**](GetLabels200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLevel

> RestLogLevel GetLevel(ctx, loggerName).Execute()

Get current log level



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	loggerName := "loggerName_example" // string | The name of the logger.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetLevel(context.Background(), loggerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLevel`: RestLogLevel
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loggerName** | **string** | The name of the logger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestLogLevel**](RestLogLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMailConfig

> RestMailConfiguration GetMailConfig(ctx).Execute()

Get mail configuration



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetMailConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetMailConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMailConfig`: RestMailConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetMailConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailConfigRequest struct via the builder pattern


### Return type

[**RestMailConfiguration**](RestMailConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeshMigrationJob

> GetMeshMigrationJob(ctx, jobId).Execute()

Get Mesh migration job details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | The ID of the job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.GetMeshMigrationJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetMeshMigrationJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeshMigrationJobRequest struct via the builder pattern


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


## GetMeshMigrationJobMessages

> GetExportJobMessages200Response GetMeshMigrationJobMessages(ctx, jobId).Severity(severity).Subject(subject).Start(start).Limit(limit).Execute()

Get Mesh migration job messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | The ID of the job
	severity := "severity_example" // string | The severity to include in the results (optional)
	subject := "subject_example" // string | The subject (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetMeshMigrationJobMessages(context.Background(), jobId).Severity(severity).Subject(subject).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetMeshMigrationJobMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeshMigrationJobMessages`: GetExportJobMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetMeshMigrationJobMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeshMigrationJobMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **severity** | **string** | The severity to include in the results | 
 **subject** | **string** | The subject | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetExportJobMessages200Response**](GetExportJobMessages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeshMigrationJobSummary

> RestMeshMigrationSummary GetMeshMigrationJobSummary(ctx, jobId).Execute()

Get Mesh migration job summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	jobId := "jobId_example" // string | The ID of the job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetMeshMigrationJobSummary(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetMeshMigrationJobSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeshMigrationJobSummary`: RestMeshMigrationSummary
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetMeshMigrationJobSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The ID of the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeshMigrationJobSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestMeshMigrationSummary**](RestMeshMigrationSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredMeshNodeById

> RestMeshNode GetRegisteredMeshNodeById(ctx, id).Execute()

Get Mesh node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "id_example" // string | The ID of the Mesh node.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetRegisteredMeshNodeById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetRegisteredMeshNodeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegisteredMeshNodeById`: RestMeshNode
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetRegisteredMeshNodeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Mesh node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredMeshNodeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestMeshNode**](RestMeshNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryArchivePolicy

> RestRepositoryPolicy GetRepositoryArchivePolicy(ctx).Execute()

Get repository archive policy



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetRepositoryArchivePolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetRepositoryArchivePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryArchivePolicy`: RestRepositoryPolicy
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetRepositoryArchivePolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryArchivePolicyRequest struct via the builder pattern


### Return type

[**RestRepositoryPolicy**](RestRepositoryPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryDeletePolicy

> RestRepositoryPolicy GetRepositoryDeletePolicy(ctx).Execute()

Get repository delete policy



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetRepositoryDeletePolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetRepositoryDeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryDeletePolicy`: RestRepositoryPolicy
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetRepositoryDeletePolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryDeletePolicyRequest struct via the builder pattern


### Return type

[**RestRepositoryPolicy**](RestRepositoryPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootLevel

> RestLogLevel GetRootLevel(ctx).Execute()

Get root log level



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetRootLevel(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetRootLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRootLevel`: RestLogLevel
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetRootLevel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootLevelRequest struct via the builder pattern


### Return type

[**RestLogLevel**](RestLogLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSenderAddress

> GetSenderAddress(ctx).Execute()

Get server mail address



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.GetSenderAddress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetSenderAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSenderAddressRequest struct via the builder pattern


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


## GetSettings2

> RestLoggingSettings GetSettings2(ctx).Execute()

Get debug logging and profiling



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetSettings2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetSettings2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings2`: RestLoggingSettings
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetSettings2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettings2Request struct via the builder pattern


### Return type

[**RestLoggingSettings**](RestLoggingSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettings3

> RestRateLimitSettings GetSettings3(ctx).Execute()

Get rate limit settings



### Example

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
	resp, r, err := apiClient.SystemMaintenanceAPI.GetSettings3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetSettings3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings3`: RestRateLimitSettings
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetSettings3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettings3Request struct via the builder pattern


### Return type

[**RestRateLimitSettings**](RestRateLimitSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportZip

> GetSupportZip(ctx, id).Execute()

Get support zip for node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "id_example" // string | The ID of the Mesh node.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.GetSupportZip(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetSupportZip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Mesh node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportZips

> GetSupportZips(ctx).Execute()

Get support zips for all Mesh nodes



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.GetSupportZips(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetSupportZips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportZipsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedKeyTypes

> GetSupportedKeyTypes(ctx).Execute()

Get supported SSH key algorithms and lengths



### Example

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
	r, err := apiClient.SystemMaintenanceAPI.GetSupportedKeyTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetSupportedKeyTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedKeyTypesRequest struct via the builder pattern


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


## GetUser

> RestApplicationUser GetUser(ctx, userSlug).Execute()

Get user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userSlug := "userSlug_example" // string | The user slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetUser(context.Background(), userSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: RestApplicationUser
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestApplicationUser**](RestApplicationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettings

> ExampleSettingsMap GetUserSettings(ctx, userSlug).Execute()

Get user settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userSlug := "userSlug_example" // string | The user slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetUserSettings(context.Background(), userSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSettings`: ExampleSettingsMap
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetUserSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExampleSettingsMap**](ExampleSettingsMap.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers2

> RestApplicationUser GetUsers2(ctx).Filter(filter).PermissionN(permissionN).Permission(permission).Group(group).Execute()

Get all users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "filter_example" // string | Return only users, whose username, name or email address <i>contain</i> the <code> filter</code> value (optional)
	permissionN := "permissionN_example" // string | The \"root\" of a single permission filter, similar to the <code>permission</code> parameter, where \"N\" is a natural number starting from 1. This allows clients to specify multiple permission filters, by providing consecutive filters as <code>permission.1</code>, <code>permission.2</code> etc. Note that the filters numbering has to start with 1 and be continuous for all filters to be processed. The total allowed number of permission filters is 50 and all filters exceeding that limit will be dropped. See the section \"Permission Filters\" above for more details on how the permission filters are processed. (optional)
	permission := "permission_example" // string | The \"root\" of a permission filter, whose value must be a valid global, project, or repository permission. Additional filter parameters referring to this filter that specify the resource (project or repository) to apply the filter to must be prefixed with <code>permission.</code>. See the section \"Permission Filters\" above for more details. (optional)
	group := "group_example" // string | return only users who are members of the given group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.GetUsers2(context.Background()).Filter(filter).PermissionN(permissionN).Permission(permission).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.GetUsers2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers2`: RestApplicationUser
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.GetUsers2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsers2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Return only users, whose username, name or email address &lt;i&gt;contain&lt;/i&gt; the &lt;code&gt; filter&lt;/code&gt; value | 
 **permissionN** | **string** | The \&quot;root\&quot; of a single permission filter, similar to the &lt;code&gt;permission&lt;/code&gt; parameter, where \&quot;N\&quot; is a natural number starting from 1. This allows clients to specify multiple permission filters, by providing consecutive filters as &lt;code&gt;permission.1&lt;/code&gt;, &lt;code&gt;permission.2&lt;/code&gt; etc. Note that the filters numbering has to start with 1 and be continuous for all filters to be processed. The total allowed number of permission filters is 50 and all filters exceeding that limit will be dropped. See the section \&quot;Permission Filters\&quot; above for more details on how the permission filters are processed. | 
 **permission** | **string** | The \&quot;root\&quot; of a permission filter, whose value must be a valid global, project, or repository permission. Additional filter parameters referring to this filter that specify the resource (project or repository) to apply the filter to must be prefixed with &lt;code&gt;permission.&lt;/code&gt;. See the section \&quot;Permission Filters\&quot; above for more details. | 
 **group** | **string** | return only users who are members of the given group | 

### Return type

[**RestApplicationUser**](RestApplicationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewExport

> RestScopesExample PreviewExport(ctx).RestExportRequest(restExportRequest).Execute()

Preview export



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restExportRequest := *openapiclient.NewRestExportRequest(*openapiclient.NewRestExportRequestRepositoriesRequest([]openapiclient.RestRepositorySelector{*openapiclient.NewRestRepositorySelector()})) // RestExportRequest | the export request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.PreviewExport(context.Background()).RestExportRequest(restExportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.PreviewExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewExport`: RestScopesExample
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.PreviewExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restExportRequest** | [**RestExportRequest**](RestExportRequest.md) | the export request | 

### Return type

[**RestScopesExample**](RestScopesExample.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewMeshMigration

> ExamplePreviewMigration PreviewMeshMigration(ctx).RestMeshMigrationRequest(restMeshMigrationRequest).Execute()

Preview Mesh migration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restMeshMigrationRequest := *openapiclient.NewRestMeshMigrationRequest() // RestMeshMigrationRequest | The export request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.PreviewMeshMigration(context.Background()).RestMeshMigrationRequest(restMeshMigrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.PreviewMeshMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewMeshMigration`: ExamplePreviewMigration
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.PreviewMeshMigration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewMeshMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restMeshMigrationRequest** | [**RestMeshMigrationRequest**](RestMeshMigrationRequest.md) | The export request | 

### Return type

[**ExamplePreviewMigration**](ExamplePreviewMigration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> Read(ctx, scriptId).Execute()

Get hook script content



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	scriptId := "scriptId_example" // string | The ID of the hook script

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.Read(context.Background(), scriptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptId** | **string** | The ID of the hook script | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, text/plain;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterNewMeshNode

> RestMeshNode RegisterNewMeshNode(ctx).RestMeshNode(restMeshNode).Execute()

Register new Mesh node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restMeshNode := *openapiclient.NewRestMeshNode() // RestMeshNode | The request specifying the new Mesh node. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.RegisterNewMeshNode(context.Background()).RestMeshNode(restMeshNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.RegisterNewMeshNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterNewMeshNode`: RestMeshNode
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.RegisterNewMeshNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterNewMeshNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restMeshNode** | [**RestMeshNode**](RestMeshNode.md) | The request specifying the new Mesh node. | 

### Return type

[**RestMeshNode**](RestMeshNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMeshMigrationRepos

> SearchMeshMigrationRepos200Response SearchMeshMigrationRepos(ctx).MigrationId(migrationId).ProjectKey(projectKey).Name(name).State(state).Remote(remote).Start(start).Limit(limit).Execute()

Find repositories by Mesh migration state



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	migrationId := "migrationId_example" // string | (optional) The currently active migration job. If not passed, this is looked up internally. (optional)
	projectKey := "projectKey_example" // string | (optional) The project key. Can be specified more than once to filter by more than one project. (optional)
	name := "name_example" // string | (optional) The repository name (optional)
	state := "state_example" // string | (optional) If a migration is active, the MeshMigrationQueueState state to filter results by. Can be specified more than once to filter by more than one state. (optional)
	remote := "remote_example" // string | (optional) Whether the repository has been fully migrated to Mesh. If not present, all repositories are considered regardless of where they're located. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.SearchMeshMigrationRepos(context.Background()).MigrationId(migrationId).ProjectKey(projectKey).Name(name).State(state).Remote(remote).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SearchMeshMigrationRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMeshMigrationRepos`: SearchMeshMigrationRepos200Response
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.SearchMeshMigrationRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMeshMigrationReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **migrationId** | **string** | (optional) The currently active migration job. If not passed, this is looked up internally. | 
 **projectKey** | **string** | (optional) The project key. Can be specified more than once to filter by more than one project. | 
 **name** | **string** | (optional) The repository name | 
 **state** | **string** | (optional) If a migration is active, the MeshMigrationQueueState state to filter results by. Can be specified more than once to filter by more than one state. | 
 **remote** | **string** | (optional) Whether the repository has been fully migrated to Mesh. If not present, all repositories are considered regardless of where they&#39;re located. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**SearchMeshMigrationRepos200Response**](SearchMeshMigrationRepos200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Set2

> RestUserRateLimitSettings Set2(ctx).RestBulkUserRateLimitSettingsUpdateRequest(restBulkUserRateLimitSettingsUpdateRequest).Execute()

Set rate limit settings for users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restBulkUserRateLimitSettingsUpdateRequest := *openapiclient.NewRestBulkUserRateLimitSettingsUpdateRequest() // RestBulkUserRateLimitSettingsUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.Set2(context.Background()).RestBulkUserRateLimitSettingsUpdateRequest(restBulkUserRateLimitSettingsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.Set2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Set2`: RestUserRateLimitSettings
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.Set2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSet2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restBulkUserRateLimitSettingsUpdateRequest** | [**RestBulkUserRateLimitSettingsUpdateRequest**](RestBulkUserRateLimitSettingsUpdateRequest.md) |  | 

### Return type

[**RestUserRateLimitSettings**](RestUserRateLimitSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Set3

> RestUserRateLimitSettings Set3(ctx, userSlug).RestUserRateLimitSettingsUpdateRequest(restUserRateLimitSettingsUpdateRequest).Execute()

Set rate limit settings for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userSlug := "userSlug_example" // string | The user slug.
	restUserRateLimitSettingsUpdateRequest := *openapiclient.NewRestUserRateLimitSettingsUpdateRequest() // RestUserRateLimitSettingsUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.Set3(context.Background(), userSlug).RestUserRateLimitSettingsUpdateRequest(restUserRateLimitSettingsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.Set3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Set3`: RestUserRateLimitSettings
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.Set3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSet3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restUserRateLimitSettingsUpdateRequest** | [**RestUserRateLimitSettingsUpdateRequest**](RestUserRateLimitSettingsUpdateRequest.md) |  | 

### Return type

[**RestUserRateLimitSettings**](RestUserRateLimitSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBanner

> SetBanner(ctx).SetBannerRequest(setBannerRequest).Execute()

Update/Set announcement banner



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	setBannerRequest := *openapiclient.NewSetBannerRequest("Audience_example") // SetBannerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.SetBanner(context.Background()).SetBannerRequest(setBannerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetBanner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetBannerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setBannerRequest** | [**SetBannerRequest**](SetBannerRequest.md) |  | 

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


## SetDefaultBranch

> SetDefaultBranch(ctx).SetDefaultBranchRequest(setDefaultBranchRequest).Execute()

Update/Set default branch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	setDefaultBranchRequest := *openapiclient.NewSetDefaultBranchRequest() // SetDefaultBranchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.SetDefaultBranch(context.Background()).SetDefaultBranchRequest(setDefaultBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetDefaultBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultBranchRequest** | [**SetDefaultBranchRequest**](SetDefaultBranchRequest.md) |  | 

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


## SetLevel

> SetLevel(ctx, levelName, loggerName).Execute()

Set log level



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	levelName := "levelName_example" // string | The level to set the logger to. Either TRACE, DEBUG, INFO, WARN or ERROR
	loggerName := "loggerName_example" // string | The name of the logger.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.SetLevel(context.Background(), levelName, loggerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**levelName** | **string** | The level to set the logger to. Either TRACE, DEBUG, INFO, WARN or ERROR | 
**loggerName** | **string** | The name of the logger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLevelRequest struct via the builder pattern


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


## SetMailConfig

> RestMailConfiguration SetMailConfig(ctx).SetMailConfigRequest(setMailConfigRequest).Execute()

Update mail configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	setMailConfigRequest := *openapiclient.NewSetMailConfigRequest() // SetMailConfigRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.SetMailConfig(context.Background()).SetMailConfigRequest(setMailConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetMailConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMailConfig`: RestMailConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.SetMailConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMailConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setMailConfigRequest** | [**SetMailConfigRequest**](SetMailConfigRequest.md) |  | 

### Return type

[**RestMailConfiguration**](RestMailConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRepositoryArchivePolicy

> RestRepositoryPolicy SetRepositoryArchivePolicy(ctx).RestRepositoryPolicy(restRepositoryPolicy).Execute()

Update repository archive policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restRepositoryPolicy := *openapiclient.NewRestRepositoryPolicy() // RestRepositoryPolicy | The request containing the details of the policy. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.SetRepositoryArchivePolicy(context.Background()).RestRepositoryPolicy(restRepositoryPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetRepositoryArchivePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRepositoryArchivePolicy`: RestRepositoryPolicy
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.SetRepositoryArchivePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRepositoryArchivePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restRepositoryPolicy** | [**RestRepositoryPolicy**](RestRepositoryPolicy.md) | The request containing the details of the policy. | 

### Return type

[**RestRepositoryPolicy**](RestRepositoryPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRepositoryDeletePolicy

> RestRepositoryPolicy SetRepositoryDeletePolicy(ctx).RestRepositoryPolicy(restRepositoryPolicy).Execute()

Update the repository delete policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restRepositoryPolicy := *openapiclient.NewRestRepositoryPolicy() // RestRepositoryPolicy | The request containing the details of the policy. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.SetRepositoryDeletePolicy(context.Background()).RestRepositoryPolicy(restRepositoryPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetRepositoryDeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRepositoryDeletePolicy`: RestRepositoryPolicy
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.SetRepositoryDeletePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRepositoryDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restRepositoryPolicy** | [**RestRepositoryPolicy**](RestRepositoryPolicy.md) | The request containing the details of the policy. | 

### Return type

[**RestRepositoryPolicy**](RestRepositoryPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRootLevel

> SetRootLevel(ctx, levelName).Execute()

Set root log level



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	levelName := "levelName_example" // string | the level to set the logger to. Either TRACE, DEBUG, INFO, WARN or ERROR

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.SetRootLevel(context.Background(), levelName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetRootLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**levelName** | **string** | the level to set the logger to. Either TRACE, DEBUG, INFO, WARN or ERROR | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRootLevelRequest struct via the builder pattern


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


## SetSenderAddress

> SetSenderAddress(ctx).Body(body).Execute()

Update server mail address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.SetSenderAddress(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetSenderAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSenderAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## SetSettings2

> RestLoggingSettings SetSettings2(ctx).SetSettings2Request(setSettings2Request).Execute()

Set debug logging and profiling



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	setSettings2Request := *openapiclient.NewSetSettings2Request(false, false) // SetSettings2Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.SetSettings2(context.Background()).SetSettings2Request(setSettings2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetSettings2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSettings2`: RestLoggingSettings
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.SetSettings2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSettings2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setSettings2Request** | [**SetSettings2Request**](SetSettings2Request.md) |  | 

### Return type

[**RestLoggingSettings**](RestLoggingSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSettings3

> RestRateLimitSettings SetSettings3(ctx).RestRateLimitSettings(restRateLimitSettings).Execute()

Set rate limit



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restRateLimitSettings := *openapiclient.NewRestRateLimitSettings() // RestRateLimitSettings | Sets the rate limit settings for the instance.  The authenticated user must have <strong>ADMIN</strong> permission to call this resource. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.SetSettings3(context.Background()).RestRateLimitSettings(restRateLimitSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.SetSettings3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSettings3`: RestRateLimitSettings
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.SetSettings3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSettings3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restRateLimitSettings** | [**RestRateLimitSettings**](RestRateLimitSettings.md) | Sets the rate limit settings for the instance.  The authenticated user must have &lt;strong&gt;ADMIN&lt;/strong&gt; permission to call this resource. | 

### Return type

[**RestRateLimitSettings**](RestRateLimitSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartExport

> RestJob StartExport(ctx).RestExportRequest(restExportRequest).Execute()

Start export job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restExportRequest := *openapiclient.NewRestExportRequest(*openapiclient.NewRestExportRequestRepositoriesRequest([]openapiclient.RestRepositorySelector{*openapiclient.NewRestRepositorySelector()})) // RestExportRequest | The request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.StartExport(context.Background()).RestExportRequest(restExportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.StartExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartExport`: RestJob
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.StartExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restExportRequest** | [**RestExportRequest**](RestExportRequest.md) | The request | 

### Return type

[**RestJob**](RestJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartImport

> RestJob StartImport(ctx).RestImportRequest(restImportRequest).Execute()

Start import job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restImportRequest := *openapiclient.NewRestImportRequest() // RestImportRequest | The request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.StartImport(context.Background()).RestImportRequest(restImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.StartImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartImport`: RestJob
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.StartImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restImportRequest** | [**RestImportRequest**](RestImportRequest.md) | The request | 

### Return type

[**RestJob**](RestJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartMeshMigration

> RestJob StartMeshMigration(ctx).StartMeshMigrationRequest(startMeshMigrationRequest).Execute()

Start Mesh migration job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	startMeshMigrationRequest := *openapiclient.NewStartMeshMigrationRequest() // StartMeshMigrationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.StartMeshMigration(context.Background()).StartMeshMigrationRequest(startMeshMigrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.StartMeshMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartMeshMigration`: RestJob
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.StartMeshMigration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartMeshMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startMeshMigrationRequest** | [**StartMeshMigrationRequest**](StartMeshMigrationRequest.md) |  | 

### Return type

[**RestJob**](RestJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalSettings

> UpdateGlobalSettings(ctx).RestSshKeySettings(restSshKeySettings).Execute()

Update global SSH key settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restSshKeySettings := *openapiclient.NewRestSshKeySettings() // RestSshKeySettings | A request containing expiry length to be set for SSH keys and a list of SSH key type restrictions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.UpdateGlobalSettings(context.Background()).RestSshKeySettings(restSshKeySettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.UpdateGlobalSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restSshKeySettings** | [**RestSshKeySettings**](RestSshKeySettings.md) | A request containing expiry length to be set for SSH keys and a list of SSH key type restrictions. | 

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


## UpdateHookScript

> RestHookScript UpdateHookScript(ctx, scriptId).ExamplePutMultipartFormData(examplePutMultipartFormData).Execute()

Update a hook script



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	scriptId := "scriptId_example" // string | The ID of the hook script
	examplePutMultipartFormData := *openapiclient.NewExamplePutMultipartFormData() // ExamplePutMultipartFormData | The multipart form data containing the hook script (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.UpdateHookScript(context.Background(), scriptId).ExamplePutMultipartFormData(examplePutMultipartFormData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.UpdateHookScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHookScript`: RestHookScript
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.UpdateHookScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptId** | **string** | The ID of the hook script | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHookScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **examplePutMultipartFormData** | [**ExamplePutMultipartFormData**](ExamplePutMultipartFormData.md) | The multipart form data containing the hook script | 

### Return type

[**RestHookScript**](RestHookScript.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicense

> RestBitbucketLicense UpdateLicense(ctx).RestBitbucketLicense(restBitbucketLicense).Execute()

Update license



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restBitbucketLicense := *openapiclient.NewRestBitbucketLicense() // RestBitbucketLicense | a JSON payload containing the encoded license to apply (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.UpdateLicense(context.Background()).RestBitbucketLicense(restBitbucketLicense).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.UpdateLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLicense`: RestBitbucketLicense
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.UpdateLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restBitbucketLicense** | [**RestBitbucketLicense**](RestBitbucketLicense.md) | a JSON payload containing the encoded license to apply | 

### Return type

[**RestBitbucketLicense**](RestBitbucketLicense.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMeshNode

> RestMeshNode UpdateMeshNode(ctx, id).RestMeshNode(restMeshNode).Execute()

Update Mesh node



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "id_example" // string | The ID of the Mesh node to update.
	restMeshNode := *openapiclient.NewRestMeshNode() // RestMeshNode | The request specifying the updated Mesh node. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.UpdateMeshNode(context.Background(), id).RestMeshNode(restMeshNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.UpdateMeshNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMeshNode`: RestMeshNode
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.UpdateMeshNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Mesh node to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMeshNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restMeshNode** | [**RestMeshNode**](RestMeshNode.md) | The request specifying the updated Mesh node. | 

### Return type

[**RestMeshNode**](RestMeshNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettings

> UpdateSettings(ctx, userSlug).ExampleSettingsMap(exampleSettingsMap).Execute()

Update user settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userSlug := "userSlug_example" // string | The user slug.
	exampleSettingsMap := *openapiclient.NewExampleSettingsMap() // ExampleSettingsMap | A map with the UserSettings entries which must be updated. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.UpdateSettings(context.Background(), userSlug).ExampleSettingsMap(exampleSettingsMap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.UpdateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exampleSettingsMap** | [**ExampleSettingsMap**](ExampleSettingsMap.md) | A map with the UserSettings entries which must be updated. | 

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


## UpdateUserDetails1

> RestApplicationUser UpdateUserDetails1(ctx).UserUpdateWithCredentials(userUpdateWithCredentials).Execute()

Update user details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userUpdateWithCredentials := *openapiclient.NewUserUpdateWithCredentials() // UserUpdateWithCredentials | The user update details (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemMaintenanceAPI.UpdateUserDetails1(context.Background()).UserUpdateWithCredentials(userUpdateWithCredentials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.UpdateUserDetails1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserDetails1`: RestApplicationUser
	fmt.Fprintf(os.Stdout, "Response from `SystemMaintenanceAPI.UpdateUserDetails1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserDetails1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userUpdateWithCredentials** | [**UserUpdateWithCredentials**](UserUpdateWithCredentials.md) | The user update details | 

### Return type

[**RestApplicationUser**](RestApplicationUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPassword1

> UpdateUserPassword1(ctx).UserPasswordUpdate(userPasswordUpdate).Execute()

Set password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userPasswordUpdate := *openapiclient.NewUserPasswordUpdate() // UserPasswordUpdate | The password update details (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.UpdateUserPassword1(context.Background()).UserPasswordUpdate(userPasswordUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.UpdateUserPassword1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPassword1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPasswordUpdate** | [**UserPasswordUpdate**](UserPasswordUpdate.md) | The password update details | 

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


## UploadAvatar1

> UploadAvatar1(ctx, userSlug).XAtlassianToken(xAtlassianToken).Avatar(avatar).Execute()

Update user avatar



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userSlug := "userSlug_example" // string | The user slug
	xAtlassianToken := "no-check" // string | This resource has Cross-Site Request Forgery (XSRF) protection. To allow the request to pass the XSRF check the caller needs to send an <code>X-Atlassian-Token</code> HTTP header with the value <code>no-check</code>. (optional)
	avatar := os.NewFile(1234, "some_file") // *os.File | The avatar file to upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SystemMaintenanceAPI.UploadAvatar1(context.Background(), userSlug).XAtlassianToken(xAtlassianToken).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemMaintenanceAPI.UploadAvatar1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadAvatar1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAtlassianToken** | **string** | This resource has Cross-Site Request Forgery (XSRF) protection. To allow the request to pass the XSRF check the caller needs to send an &lt;code&gt;X-Atlassian-Token&lt;/code&gt; HTTP header with the value &lt;code&gt;no-check&lt;/code&gt;. | 
 **avatar** | ***os.File** | The avatar file to upload. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


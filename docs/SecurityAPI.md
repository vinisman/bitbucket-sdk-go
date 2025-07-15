# \SecurityAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExemptRepo**](SecurityAPI.md#AddExemptRepo) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/exempt | Exempt a repo from secret scanning
[**AddKey**](SecurityAPI.md#AddKey) | **Post** /gpg/latest/keys | Create a GPG key
[**BulkAddExemptRepositories**](SecurityAPI.md#BulkAddExemptRepositories) | **Post** /api/latest/secret-scanning/exempt | Bulk exempt repos from secret scanning
[**BulkAddExemptRepositories1**](SecurityAPI.md#BulkAddExemptRepositories1) | **Post** /api/latest/projects/{projectKey}/secret-scanning/exempt | Bulk exempt repos from secret scanning
[**CreateAllowlistRule**](SecurityAPI.md#CreateAllowlistRule) | **Post** /api/latest/projects/{projectKey}/secret-scanning/allowlist | Create project secret scanning allowlist rule
[**CreateAllowlistRule1**](SecurityAPI.md#CreateAllowlistRule1) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/allowlist | Create repository secret scanning allowlist rule
[**CreateCertificate**](SecurityAPI.md#CreateCertificate) | **Post** /api/latest/signing/x509-certificates | Create an X.509 certificate
[**CreateRule**](SecurityAPI.md#CreateRule) | **Post** /api/latest/projects/{projectKey}/secret-scanning/rules | Create project secret scanning rule
[**CreateRule1**](SecurityAPI.md#CreateRule1) | **Post** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/rules | Create repository secret scanning rule
[**CreateRule2**](SecurityAPI.md#CreateRule2) | **Post** /api/latest/secret-scanning/rules | Create global secret scanning rule
[**DeleteAllowlistRule**](SecurityAPI.md#DeleteAllowlistRule) | **Delete** /api/latest/projects/{projectKey}/secret-scanning/allowlist/{id} | Delete a project secret scanning allowlist rule
[**DeleteAllowlistRule1**](SecurityAPI.md#DeleteAllowlistRule1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/allowlist/{id} | Delete a repository secret scanning allowlist rule
[**DeleteCertificate**](SecurityAPI.md#DeleteCertificate) | **Delete** /api/latest/signing/x509-certificates/{id} | Delete an X.509 certificate
[**DeleteExemptRepo**](SecurityAPI.md#DeleteExemptRepo) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/exempt | Delete an exempt repository
[**DeleteForUser**](SecurityAPI.md#DeleteForUser) | **Delete** /gpg/latest/keys | Delete all GPG keys for user
[**DeleteInactiveKeys**](SecurityAPI.md#DeleteInactiveKeys) | **Delete** /secrets/1.0/keys/inactive | Delete inactive AES key(s)
[**DeleteKey**](SecurityAPI.md#DeleteKey) | **Delete** /gpg/latest/keys/{fingerprintOrId} | Delete a GPG key
[**DeleteRule**](SecurityAPI.md#DeleteRule) | **Delete** /api/latest/projects/{projectKey}/secret-scanning/rules/{id} | Delete a project secret scanning rule
[**DeleteRule1**](SecurityAPI.md#DeleteRule1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/rules/{id} | Delete a repository secret scanning rule
[**DeleteRule2**](SecurityAPI.md#DeleteRule2) | **Delete** /api/latest/secret-scanning/rules/{id} | Delete a global secret scanning rule
[**EditAllowlistRule**](SecurityAPI.md#EditAllowlistRule) | **Put** /api/latest/projects/{projectKey}/secret-scanning/allowlist/{id} | Edit an existing project secret scanning allowlist rule
[**EditAllowlistRule1**](SecurityAPI.md#EditAllowlistRule1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/allowlist/{id} | Edit an existing repository secret scanning allowlist rule
[**EditRule**](SecurityAPI.md#EditRule) | **Put** /api/latest/projects/{projectKey}/secret-scanning/rules/{id} | Edit an existing project secret scanning rule
[**EditRule1**](SecurityAPI.md#EditRule1) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/rules/{id} | Edit an existing repository secret scanning rule
[**EditRule2**](SecurityAPI.md#EditRule2) | **Put** /api/latest/secret-scanning/rules/{id} | Edit a global secret scanning rule.
[**FindExemptReposByProject**](SecurityAPI.md#FindExemptReposByProject) | **Get** /api/latest/projects/{projectKey}/secret-scanning/exempt | Find repos exempt from secret scanning for a project
[**FindExemptReposByScope**](SecurityAPI.md#FindExemptReposByScope) | **Get** /api/latest/secret-scanning/exempt | Find all repos exempt from secret scan
[**GetAllCertificates**](SecurityAPI.md#GetAllCertificates) | **Get** /api/latest/signing/x509-certificates | Get all X.509 certificates
[**GetAllowlistRule**](SecurityAPI.md#GetAllowlistRule) | **Get** /api/latest/projects/{projectKey}/secret-scanning/allowlist/{id} | Get a project secret scanning allowlist rule
[**GetAllowlistRule1**](SecurityAPI.md#GetAllowlistRule1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/allowlist/{id} | Get a repository secret scanning allowlist rule
[**GetInactiveKeys**](SecurityAPI.md#GetInactiveKeys) | **Get** /secrets/1.0/keys/inactive | Retrieve inactive AES key(s)
[**GetKeysForUser**](SecurityAPI.md#GetKeysForUser) | **Get** /gpg/latest/keys | Get all GPG keys
[**GetRule**](SecurityAPI.md#GetRule) | **Get** /api/latest/projects/{projectKey}/secret-scanning/rules/{id} | Get a project secret scanning rule
[**GetRule1**](SecurityAPI.md#GetRule1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/rules/{id} | Get a repository secret scanning rule
[**GetRule2**](SecurityAPI.md#GetRule2) | **Get** /api/latest/secret-scanning/rules/{id} | Get a global secret scanning rule
[**GetSystemSigningConfiguration**](SecurityAPI.md#GetSystemSigningConfiguration) | **Get** /api/latest/system-signing/configuration | Get system signing configuration
[**IsRepoExempt**](SecurityAPI.md#IsRepoExempt) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/exempt | Get whether a repository is exempt
[**RotateKey**](SecurityAPI.md#RotateKey) | **Post** /secrets/1.0/keys/rotate | Rotate the current AES key
[**Search1**](SecurityAPI.md#Search1) | **Get** /api/latest/projects/{projectKey}/secret-scanning/rules | Find project secret scanning rules
[**Search2**](SecurityAPI.md#Search2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/allowlist | Find repository secret scanning allowlist rules
[**Search3**](SecurityAPI.md#Search3) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/secret-scanning/rules | Find repository secret scanning rules
[**Search4**](SecurityAPI.md#Search4) | **Get** /api/latest/secret-scanning/rules | Find global secret scanning rules
[**SearchAllowlistRule**](SecurityAPI.md#SearchAllowlistRule) | **Get** /api/latest/projects/{projectKey}/secret-scanning/allowlist | Find project secret scanning allowlist rules
[**UpdateCertificateRevocationListEntries**](SecurityAPI.md#UpdateCertificateRevocationListEntries) | **Put** /api/latest/signing/x509-certificates/crl/{id} | Update X.509 CRL entries
[**UpdateSystemSigningConfiguration**](SecurityAPI.md#UpdateSystemSigningConfiguration) | **Post** /api/latest/system-signing/configuration | Update system signing configuration



## AddExemptRepo

> AddExemptRepo(ctx, repositorySlug, projectKey).Execute()

Exempt a repo from secret scanning



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	projectKey := "projectKey_example" // string | The project key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.AddExemptRepo(context.Background(), repositorySlug, projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.AddExemptRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositorySlug** | **string** | The repository slug. | 
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddExemptRepoRequest struct via the builder pattern


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


## AddKey

> RestGpgKey AddKey(ctx).User(user).RestGpgKey(restGpgKey).Execute()

Create a GPG key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	user := "user_example" // string | The name of the user to add a key for (optional; requires ADMIN permission or higher). (optional)
	restGpgKey := *openapiclient.NewRestGpgKey() // RestGpgKey | The request body. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.AddKey(context.Background()).User(user).RestGpgKey(restGpgKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.AddKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddKey`: RestGpgKey
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.AddKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | The name of the user to add a key for (optional; requires ADMIN permission or higher). | 
 **restGpgKey** | [**RestGpgKey**](RestGpgKey.md) | The request body. | 

### Return type

[**RestGpgKey**](RestGpgKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkAddExemptRepositories

> BulkAddExemptRepositories(ctx).RestRepositorySelector(restRepositorySelector).Execute()

Bulk exempt repos from secret scanning



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restRepositorySelector := []openapiclient.RestRepositorySelector{*openapiclient.NewRestRepositorySelector()} // []RestRepositorySelector |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.BulkAddExemptRepositories(context.Background()).RestRepositorySelector(restRepositorySelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.BulkAddExemptRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAddExemptRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restRepositorySelector** | [**[]RestRepositorySelector**](RestRepositorySelector.md) |  | 

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


## BulkAddExemptRepositories1

> BulkAddExemptRepositories1(ctx, projectKey).RestRepositorySelector(restRepositorySelector).Execute()

Bulk exempt repos from secret scanning



### Example

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
	restRepositorySelector := []openapiclient.RestRepositorySelector{*openapiclient.NewRestRepositorySelector()} // []RestRepositorySelector |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.BulkAddExemptRepositories1(context.Background(), projectKey).RestRepositorySelector(restRepositorySelector).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.BulkAddExemptRepositories1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBulkAddExemptRepositories1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restRepositorySelector** | [**[]RestRepositorySelector**](RestRepositorySelector.md) |  | 

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


## CreateAllowlistRule

> RestSecretScanningAllowlistRule CreateAllowlistRule(ctx, projectKey).RestSecretScanningAllowlistRuleSetRequest(restSecretScanningAllowlistRuleSetRequest).Execute()

Create project secret scanning allowlist rule



### Example

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
	restSecretScanningAllowlistRuleSetRequest := *openapiclient.NewRestSecretScanningAllowlistRuleSetRequest() // RestSecretScanningAllowlistRuleSetRequest | Allowlist rule to create, either the line regular expression or the path regular expression must be present

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateAllowlistRule(context.Background(), projectKey).RestSecretScanningAllowlistRuleSetRequest(restSecretScanningAllowlistRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateAllowlistRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAllowlistRule`: RestSecretScanningAllowlistRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateAllowlistRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAllowlistRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restSecretScanningAllowlistRuleSetRequest** | [**RestSecretScanningAllowlistRuleSetRequest**](RestSecretScanningAllowlistRuleSetRequest.md) | Allowlist rule to create, either the line regular expression or the path regular expression must be present | 

### Return type

[**RestSecretScanningAllowlistRule**](RestSecretScanningAllowlistRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAllowlistRule1

> RestSecretScanningAllowlistRule CreateAllowlistRule1(ctx, projectKey, repositorySlug).RestSecretScanningAllowlistRuleSetRequest(restSecretScanningAllowlistRuleSetRequest).Execute()

Create repository secret scanning allowlist rule



### Example

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
	restSecretScanningAllowlistRuleSetRequest := *openapiclient.NewRestSecretScanningAllowlistRuleSetRequest() // RestSecretScanningAllowlistRuleSetRequest | Allowlist rule to create, either the line regular expression or the path regular expression must be present

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateAllowlistRule1(context.Background(), projectKey, repositorySlug).RestSecretScanningAllowlistRuleSetRequest(restSecretScanningAllowlistRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateAllowlistRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAllowlistRule1`: RestSecretScanningAllowlistRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateAllowlistRule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAllowlistRule1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restSecretScanningAllowlistRuleSetRequest** | [**RestSecretScanningAllowlistRuleSetRequest**](RestSecretScanningAllowlistRuleSetRequest.md) | Allowlist rule to create, either the line regular expression or the path regular expression must be present | 

### Return type

[**RestSecretScanningAllowlistRule**](RestSecretScanningAllowlistRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificate

> RestX509Certificate CreateCertificate(ctx).Certificate(certificate).Execute()

Create an X.509 certificate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	certificate := os.NewFile(1234, "some_file") // *os.File | The X.509 certificate file to upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateCertificate(context.Background()).Certificate(certificate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCertificate`: RestX509Certificate
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificate** | ***os.File** | The X.509 certificate file to upload. | 

### Return type

[**RestX509Certificate**](RestX509Certificate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRule

> RestSecretScanningRule CreateRule(ctx, projectKey).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()

Create project secret scanning rule



### Example

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
	restSecretScanningRuleSetRequest := *openapiclient.NewRestSecretScanningRuleSetRequest() // RestSecretScanningRuleSetRequest | Rule to create, either the line regular expression or the path regular expression must be present

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateRule(context.Background(), projectKey).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRule`: RestSecretScanningRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restSecretScanningRuleSetRequest** | [**RestSecretScanningRuleSetRequest**](RestSecretScanningRuleSetRequest.md) | Rule to create, either the line regular expression or the path regular expression must be present | 

### Return type

[**RestSecretScanningRule**](RestSecretScanningRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRule1

> RestSecretScanningRule CreateRule1(ctx, projectKey, repositorySlug).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()

Create repository secret scanning rule



### Example

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
	restSecretScanningRuleSetRequest := *openapiclient.NewRestSecretScanningRuleSetRequest() // RestSecretScanningRuleSetRequest | Rule to create, either the line regular expression or the path regular expression must be present

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateRule1(context.Background(), projectKey, repositorySlug).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRule1`: RestSecretScanningRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateRule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRule1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restSecretScanningRuleSetRequest** | [**RestSecretScanningRuleSetRequest**](RestSecretScanningRuleSetRequest.md) | Rule to create, either the line regular expression or the path regular expression must be present | 

### Return type

[**RestSecretScanningRule**](RestSecretScanningRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRule2

> RestSecretScanningRule CreateRule2(ctx).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()

Create global secret scanning rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	restSecretScanningRuleSetRequest := *openapiclient.NewRestSecretScanningRuleSetRequest() // RestSecretScanningRuleSetRequest | Rule to create, either the line regular expression or the path regular expression must be present

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.CreateRule2(context.Background()).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.CreateRule2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRule2`: RestSecretScanningRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.CreateRule2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRule2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restSecretScanningRuleSetRequest** | [**RestSecretScanningRuleSetRequest**](RestSecretScanningRuleSetRequest.md) | Rule to create, either the line regular expression or the path regular expression must be present | 

### Return type

[**RestSecretScanningRule**](RestSecretScanningRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllowlistRule

> DeleteAllowlistRule(ctx, projectKey, id).Execute()

Delete a project secret scanning allowlist rule



### Example

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
	id := "7" // string | The allowlist rule id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteAllowlistRule(context.Background(), projectKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteAllowlistRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The allowlist rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllowlistRuleRequest struct via the builder pattern


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


## DeleteAllowlistRule1

> DeleteAllowlistRule1(ctx, projectKey, id, repositorySlug).Execute()

Delete a repository secret scanning allowlist rule



### Example

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
	id := "7" // string | The allowlist rule id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteAllowlistRule1(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteAllowlistRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The allowlist rule id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllowlistRule1Request struct via the builder pattern


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


## DeleteCertificate

> RestX509Certificate DeleteCertificate(ctx, id).Execute()

Delete an X.509 certificate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "id_example" // string | The ID of the X.509 certificate.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.DeleteCertificate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertificate`: RestX509Certificate
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.DeleteCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the X.509 certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestX509Certificate**](RestX509Certificate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteExemptRepo

> DeleteExemptRepo(ctx, repositorySlug, projectKey).Execute()

Delete an exempt repository



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	projectKey := "projectKey_example" // string | The project key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteExemptRepo(context.Background(), repositorySlug, projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteExemptRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositorySlug** | **string** | The repository slug. | 
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExemptRepoRequest struct via the builder pattern


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


## DeleteForUser

> DeleteForUser(ctx).User(user).Execute()

Delete all GPG keys for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	user := "user_example" // string | The username of the user to delete the keys for. If no username is specified, the GPG keys will be deleted for the currently authenticated user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteForUser(context.Background()).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | The username of the user to delete the keys for. If no username is specified, the GPG keys will be deleted for the currently authenticated user. | 

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


## DeleteInactiveKeys

> DeleteInactiveKeys(ctx).Execute()

Delete inactive AES key(s)



### Example

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
	r, err := apiClient.SecurityAPI.DeleteInactiveKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteInactiveKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInactiveKeysRequest struct via the builder pattern


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


## DeleteKey

> DeleteKey(ctx, fingerprintOrId).Execute()

Delete a GPG key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	fingerprintOrId := "fingerprintOrId_example" // string | The GPG fingerprint or ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteKey(context.Background(), fingerprintOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fingerprintOrId** | **string** | The GPG fingerprint or ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


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


## DeleteRule

> DeleteRule(ctx, projectKey, id).Execute()

Delete a project secret scanning rule



### Example

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
	id := "7" // string | The rule id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteRule(context.Background(), projectKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## DeleteRule1

> DeleteRule1(ctx, projectKey, id, repositorySlug).Execute()

Delete a repository secret scanning rule



### Example

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
	id := "7" // string | The rule id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteRule1(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The rule id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRule1Request struct via the builder pattern


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


## DeleteRule2

> DeleteRule2(ctx, id).Execute()

Delete a global secret scanning rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "7" // string | The rule id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.DeleteRule2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.DeleteRule2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRule2Request struct via the builder pattern


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


## EditAllowlistRule

> RestSecretScanningAllowlistRule EditAllowlistRule(ctx, projectKey, id).RestSecretScanningAllowlistRuleSetRequest(restSecretScanningAllowlistRuleSetRequest).Execute()

Edit an existing project secret scanning allowlist rule



### Example

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
	id := "7" // string | The allowlist rule id.
	restSecretScanningAllowlistRuleSetRequest := *openapiclient.NewRestSecretScanningAllowlistRuleSetRequest() // RestSecretScanningAllowlistRuleSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.EditAllowlistRule(context.Background(), projectKey, id).RestSecretScanningAllowlistRuleSetRequest(restSecretScanningAllowlistRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.EditAllowlistRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAllowlistRule`: RestSecretScanningAllowlistRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.EditAllowlistRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The allowlist rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAllowlistRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restSecretScanningAllowlistRuleSetRequest** | [**RestSecretScanningAllowlistRuleSetRequest**](RestSecretScanningAllowlistRuleSetRequest.md) |  | 

### Return type

[**RestSecretScanningAllowlistRule**](RestSecretScanningAllowlistRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAllowlistRule1

> RestSecretScanningAllowlistRule EditAllowlistRule1(ctx, projectKey, id, repositorySlug).RestSecretScanningAllowlistRuleSetRequest(restSecretScanningAllowlistRuleSetRequest).Execute()

Edit an existing repository secret scanning allowlist rule



### Example

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
	id := "7" // string | The allowlist rule id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restSecretScanningAllowlistRuleSetRequest := *openapiclient.NewRestSecretScanningAllowlistRuleSetRequest() // RestSecretScanningAllowlistRuleSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.EditAllowlistRule1(context.Background(), projectKey, id, repositorySlug).RestSecretScanningAllowlistRuleSetRequest(restSecretScanningAllowlistRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.EditAllowlistRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAllowlistRule1`: RestSecretScanningAllowlistRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.EditAllowlistRule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The allowlist rule id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAllowlistRule1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restSecretScanningAllowlistRuleSetRequest** | [**RestSecretScanningAllowlistRuleSetRequest**](RestSecretScanningAllowlistRuleSetRequest.md) |  | 

### Return type

[**RestSecretScanningAllowlistRule**](RestSecretScanningAllowlistRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRule

> RestSecretScanningRule EditRule(ctx, projectKey, id).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()

Edit an existing project secret scanning rule



### Example

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
	id := "7" // string | The rule id.
	restSecretScanningRuleSetRequest := *openapiclient.NewRestSecretScanningRuleSetRequest() // RestSecretScanningRuleSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.EditRule(context.Background(), projectKey, id).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.EditRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRule`: RestSecretScanningRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.EditRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restSecretScanningRuleSetRequest** | [**RestSecretScanningRuleSetRequest**](RestSecretScanningRuleSetRequest.md) |  | 

### Return type

[**RestSecretScanningRule**](RestSecretScanningRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRule1

> RestSecretScanningRule EditRule1(ctx, projectKey, id, repositorySlug).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()

Edit an existing repository secret scanning rule



### Example

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
	id := "7" // string | The rule id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restSecretScanningRuleSetRequest := *openapiclient.NewRestSecretScanningRuleSetRequest() // RestSecretScanningRuleSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.EditRule1(context.Background(), projectKey, id, repositorySlug).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.EditRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRule1`: RestSecretScanningRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.EditRule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The rule id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRule1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restSecretScanningRuleSetRequest** | [**RestSecretScanningRuleSetRequest**](RestSecretScanningRuleSetRequest.md) |  | 

### Return type

[**RestSecretScanningRule**](RestSecretScanningRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRule2

> RestSecretScanningRule EditRule2(ctx, id).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()

Edit a global secret scanning rule.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "7" // string | The rule id.
	restSecretScanningRuleSetRequest := *openapiclient.NewRestSecretScanningRuleSetRequest() // RestSecretScanningRuleSetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.EditRule2(context.Background(), id).RestSecretScanningRuleSetRequest(restSecretScanningRuleSetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.EditRule2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRule2`: RestSecretScanningRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.EditRule2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRule2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restSecretScanningRuleSetRequest** | [**RestSecretScanningRuleSetRequest**](RestSecretScanningRuleSetRequest.md) |  | 

### Return type

[**RestSecretScanningRule**](RestSecretScanningRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindExemptReposByProject

> GetRepositoriesRecentlyAccessed200Response FindExemptReposByProject(ctx, projectKey).Order(order).Start(start).Limit(limit).Execute()

Find repos exempt from secret scanning for a project



### Example

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
	order := "order_example" // string | Order by project name followed by repository name either ascending or descending, defaults to ascending. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.FindExemptReposByProject(context.Background(), projectKey).Order(order).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.FindExemptReposByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindExemptReposByProject`: GetRepositoriesRecentlyAccessed200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.FindExemptReposByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindExemptReposByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | **string** | Order by project name followed by repository name either ascending or descending, defaults to ascending. | 
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


## FindExemptReposByScope

> GetRepositoriesRecentlyAccessed200Response FindExemptReposByScope(ctx).Order(order).Start(start).Limit(limit).Execute()

Find all repos exempt from secret scan



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	order := "order_example" // string | Order by project name followed by repository name either ascending or descending, defaults to ascending. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.FindExemptReposByScope(context.Background()).Order(order).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.FindExemptReposByScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindExemptReposByScope`: GetRepositoriesRecentlyAccessed200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.FindExemptReposByScope`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindExemptReposByScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Order by project name followed by repository name either ascending or descending, defaults to ascending. | 
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


## GetAllCertificates

> RestX509Certificate GetAllCertificates(ctx).Execute()

Get all X.509 certificates



### Example

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
	resp, r, err := apiClient.SecurityAPI.GetAllCertificates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetAllCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCertificates`: RestX509Certificate
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetAllCertificates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCertificatesRequest struct via the builder pattern


### Return type

[**RestX509Certificate**](RestX509Certificate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowlistRule

> RestSecretScanningAllowlistRule GetAllowlistRule(ctx, projectKey, id).Execute()

Get a project secret scanning allowlist rule



### Example

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
	id := "7" // string | The allowlist rule id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetAllowlistRule(context.Background(), projectKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetAllowlistRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllowlistRule`: RestSecretScanningAllowlistRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetAllowlistRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The allowlist rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowlistRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestSecretScanningAllowlistRule**](RestSecretScanningAllowlistRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowlistRule1

> RestSecretScanningAllowlistRule GetAllowlistRule1(ctx, projectKey, id, repositorySlug).Execute()

Get a repository secret scanning allowlist rule



### Example

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
	id := "7" // string | The allowlist rule id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetAllowlistRule1(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetAllowlistRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllowlistRule1`: RestSecretScanningAllowlistRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetAllowlistRule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The allowlist rule id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowlistRule1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestSecretScanningAllowlistRule**](RestSecretScanningAllowlistRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInactiveKeys

> GetInactiveKeys(ctx).Execute()

Retrieve inactive AES key(s)



### Example

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
	r, err := apiClient.SecurityAPI.GetInactiveKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetInactiveKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInactiveKeysRequest struct via the builder pattern


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


## GetKeysForUser

> GetKeysForUser200Response GetKeysForUser(ctx).User(user).Start(start).Limit(limit).Execute()

Get all GPG keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	user := "user_example" // string | The name of the user to get keys for (optional; requires ADMIN permission or higher). (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetKeysForUser(context.Background()).User(user).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetKeysForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeysForUser`: GetKeysForUser200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetKeysForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKeysForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | The name of the user to get keys for (optional; requires ADMIN permission or higher). | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetKeysForUser200Response**](GetKeysForUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRule

> RestSecretScanningRule GetRule(ctx, projectKey, id).Execute()

Get a project secret scanning rule



### Example

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
	id := "7" // string | The rule id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetRule(context.Background(), projectKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRule`: RestSecretScanningRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestSecretScanningRule**](RestSecretScanningRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRule1

> RestSecretScanningRule GetRule1(ctx, projectKey, id, repositorySlug).Execute()

Get a repository secret scanning rule



### Example

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
	id := "7" // string | The rule id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetRule1(context.Background(), projectKey, id, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetRule1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRule1`: RestSecretScanningRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetRule1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**id** | **string** | The rule id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRule1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestSecretScanningRule**](RestSecretScanningRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRule2

> RestSecretScanningRule GetRule2(ctx, id).Execute()

Get a global secret scanning rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "7" // string | The rule id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.GetRule2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetRule2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRule2`: RestSecretScanningRule
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetRule2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRule2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestSecretScanningRule**](RestSecretScanningRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemSigningConfiguration

> RestSystemSigningConfiguration GetSystemSigningConfiguration(ctx).Execute()

Get system signing configuration



### Example

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
	resp, r, err := apiClient.SecurityAPI.GetSystemSigningConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.GetSystemSigningConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemSigningConfiguration`: RestSystemSigningConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.GetSystemSigningConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemSigningConfigurationRequest struct via the builder pattern


### Return type

[**RestSystemSigningConfiguration**](RestSystemSigningConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsRepoExempt

> IsRepoExempt(ctx, repositorySlug, projectKey).Execute()

Get whether a repository is exempt



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	projectKey := "projectKey_example" // string | The project key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.IsRepoExempt(context.Background(), repositorySlug, projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.IsRepoExempt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositorySlug** | **string** | The repository slug. | 
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsRepoExemptRequest struct via the builder pattern


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


## RotateKey

> RotateKey(ctx).Execute()

Rotate the current AES key



### Example

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
	r, err := apiClient.SecurityAPI.RotateKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.RotateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRotateKeyRequest struct via the builder pattern


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


## Search1

> Search3200Response Search1(ctx, projectKey).Filter(filter).Order(order).Start(start).Limit(limit).Execute()

Find project secret scanning rules



### Example

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
	filter := "Access" // string | Filter names by the provided text (optional)
	order := "order_example" // string | Order by (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.Search1(context.Background(), projectKey).Filter(filter).Order(order).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.Search1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search1`: Search3200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.Search1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearch1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter names by the provided text | 
 **order** | **string** | Order by | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**Search3200Response**](Search3200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search2

> Search2200Response Search2(ctx, projectKey, repositorySlug).Filter(filter).Order(order).Start(start).Limit(limit).Execute()

Find repository secret scanning allowlist rules



### Example

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
	filter := "Access" // string | Filter names by the provided text (optional)
	order := "order_example" // string | Order by (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.Search2(context.Background(), projectKey, repositorySlug).Filter(filter).Order(order).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.Search2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search2`: Search2200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.Search2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearch2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | Filter names by the provided text | 
 **order** | **string** | Order by | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**Search2200Response**](Search2200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search3

> Search3200Response Search3(ctx, repositorySlug, projectKey).Filter(filter).Order(order).Start(start).Limit(limit).Execute()

Find repository secret scanning rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	projectKey := "projectKey_example" // string | The project key.
	filter := "Access" // string | Filter names by the provided text (optional)
	order := "order_example" // string | Order by (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.Search3(context.Background(), repositorySlug, projectKey).Filter(filter).Order(order).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.Search3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search3`: Search3200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.Search3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositorySlug** | **string** | The repository slug. | 
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearch3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | Filter names by the provided text | 
 **order** | **string** | Order by | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**Search3200Response**](Search3200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search4

> Search3200Response Search4(ctx).Filter(filter).Order(order).Start(start).Limit(limit).Execute()

Find global secret scanning rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "Access" // string | Filter by rule name (optional)
	order := "order_example" // string | Order by (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.Search4(context.Background()).Filter(filter).Order(order).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.Search4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search4`: Search3200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.Search4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearch4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter by rule name | 
 **order** | **string** | Order by | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**Search3200Response**](Search3200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAllowlistRule

> Search2200Response SearchAllowlistRule(ctx, projectKey).Filter(filter).Order(order).Start(start).Limit(limit).Execute()

Find project secret scanning allowlist rules



### Example

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
	filter := "Access" // string | Filter names by the provided text (optional)
	order := "order_example" // string | Order by (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.SearchAllowlistRule(context.Background(), projectKey).Filter(filter).Order(order).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.SearchAllowlistRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAllowlistRule`: Search2200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.SearchAllowlistRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAllowlistRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter names by the provided text | 
 **order** | **string** | Order by | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**Search2200Response**](Search2200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificateRevocationListEntries

> UpdateCertificateRevocationListEntries(ctx, id).Execute()

Update X.509 CRL entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "id_example" // string | The ID of the issuer certificate.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.UpdateCertificateRevocationListEntries(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UpdateCertificateRevocationListEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issuer certificate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificateRevocationListEntriesRequest struct via the builder pattern


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


## UpdateSystemSigningConfiguration

> RestSystemSigningConfiguration UpdateSystemSigningConfiguration(ctx).UpdateSystemSigningConfigurationRequest(updateSystemSigningConfigurationRequest).Execute()

Update system signing configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	updateSystemSigningConfigurationRequest := *openapiclient.NewUpdateSystemSigningConfigurationRequest() // UpdateSystemSigningConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.UpdateSystemSigningConfiguration(context.Background()).UpdateSystemSigningConfigurationRequest(updateSystemSigningConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.UpdateSystemSigningConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSystemSigningConfiguration`: RestSystemSigningConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.UpdateSystemSigningConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemSigningConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSystemSigningConfigurationRequest** | [**UpdateSystemSigningConfigurationRequest**](UpdateSystemSigningConfigurationRequest.md) |  | 

### Return type

[**RestSystemSigningConfiguration**](RestSystemSigningConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AuthenticationAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AGet**](AuthenticationAPI.md#AGet) | **Get** /basicauth/latest/config | Get basic auth configuration
[**AddForProject**](AuthenticationAPI.md#AddForProject) | **Post** /keys/latest/projects/{projectKey}/ssh | Add project SSH key
[**AddForRepository**](AuthenticationAPI.md#AddForRepository) | **Post** /keys/latest/projects/{projectKey}/repos/{repositorySlug}/ssh | Add repository SSH key
[**AddIdp**](AuthenticationAPI.md#AddIdp) | **Post** /authconfig/latest/idps | Create IdP configuration
[**AddSshKey**](AuthenticationAPI.md#AddSshKey) | **Post** /ssh/latest/keys | Add SSH key for user
[**Authenticate**](AuthenticationAPI.md#Authenticate) | **Post** /tsv/latest/authenticate | Authenticate with 2SV
[**AuthenticateWithRecoveryCode**](AuthenticationAPI.md#AuthenticateWithRecoveryCode) | **Post** /tsv/latest/authenticate/recovery-code | Authenticate using recovery code
[**CompleteAuthenticationChange**](AuthenticationAPI.md#CompleteAuthenticationChange) | **Post** /tsv/latest/totp/complete-enrollment-update | Complete authentication app update for 2SV
[**CompleteEnforcedEnrollment**](AuthenticationAPI.md#CompleteEnforcedEnrollment) | **Post** /tsv/latest/totp/complete-enforced-enrollment | Complete enforced enrollment in 2SV
[**CompleteVoluntaryEnrollment**](AuthenticationAPI.md#CompleteVoluntaryEnrollment) | **Post** /tsv/latest/totp/complete-voluntary-enrollment | Complete voluntary enrollment in 2SV
[**CreateAccessToken**](AuthenticationAPI.md#CreateAccessToken) | **Put** /access-tokens/latest/projects/{projectKey} | Create project HTTP token
[**CreateAccessToken1**](AuthenticationAPI.md#CreateAccessToken1) | **Put** /access-tokens/latest/projects/{projectKey}/repos/{repositorySlug} | Create repository HTTP token
[**CreateAccessToken2**](AuthenticationAPI.md#CreateAccessToken2) | **Put** /access-tokens/latest/users/{userSlug} | Create personal HTTP token
[**DeleteById**](AuthenticationAPI.md#DeleteById) | **Delete** /access-tokens/latest/projects/{projectKey}/{tokenId} | Delete a HTTP token
[**DeleteById1**](AuthenticationAPI.md#DeleteById1) | **Delete** /access-tokens/latest/projects/{projectKey}/repos/{repositorySlug}/{tokenId} | Delete a HTTP token
[**DeleteById2**](AuthenticationAPI.md#DeleteById2) | **Delete** /access-tokens/latest/users/{userSlug}/{tokenId} | Delete a HTTP token
[**DeleteSshKey**](AuthenticationAPI.md#DeleteSshKey) | **Delete** /ssh/latest/keys/{keyId} | Remove SSH key
[**DeleteSshKeys**](AuthenticationAPI.md#DeleteSshKeys) | **Delete** /ssh/latest/keys | Delete all user SSH key
[**ElevatePermissionsWithPassword**](AuthenticationAPI.md#ElevatePermissionsWithPassword) | **Post** /tsv/latest/elevate-permissions/password | Create elevated session with password
[**ElevatePermissionsWithRecoveryCode**](AuthenticationAPI.md#ElevatePermissionsWithRecoveryCode) | **Post** /tsv/latest/elevate-permissions/recovery-code | Create elevated session with recovery code
[**ElevatePermissionsWithTotp**](AuthenticationAPI.md#ElevatePermissionsWithTotp) | **Post** /tsv/latest/elevate-permissions/totp | Create elevated session with TOTP
[**GetAllAccessTokens**](AuthenticationAPI.md#GetAllAccessTokens) | **Get** /access-tokens/latest/projects/{projectKey} | Get project HTTP tokens
[**GetAllAccessTokens1**](AuthenticationAPI.md#GetAllAccessTokens1) | **Get** /access-tokens/latest/projects/{projectKey}/repos/{repositorySlug} | Get repository HTTP tokens
[**GetAllAccessTokens2**](AuthenticationAPI.md#GetAllAccessTokens2) | **Get** /access-tokens/latest/users/{userSlug} | Get personal HTTP tokens
[**GetById**](AuthenticationAPI.md#GetById) | **Get** /access-tokens/latest/projects/{projectKey}/{tokenId} | Get HTTP token by ID
[**GetById1**](AuthenticationAPI.md#GetById1) | **Get** /access-tokens/latest/projects/{projectKey}/repos/{repositorySlug}/{tokenId} | Get HTTP token by ID
[**GetById2**](AuthenticationAPI.md#GetById2) | **Get** /access-tokens/latest/users/{userSlug}/{tokenId} | Get HTTP token by ID
[**GetCaptchaData**](AuthenticationAPI.md#GetCaptchaData) | **Get** /tsv/latest/authenticate/captcha | Get CAPTCHA challenge
[**GetConfig**](AuthenticationAPI.md#GetConfig) | **Get** /authconfig/latest/sso | Get SSO configuration
[**GetElevatedPermissionStatus**](AuthenticationAPI.md#GetElevatedPermissionStatus) | **Get** /tsv/latest/elevate-permissions | Get elevated session status
[**GetForProject**](AuthenticationAPI.md#GetForProject) | **Get** /keys/latest/projects/{projectKey}/ssh/{keyId} | Get project SSH key
[**GetForProjects**](AuthenticationAPI.md#GetForProjects) | **Get** /keys/latest/ssh/{keyId}/projects | Get project SSH keys
[**GetForRepositories**](AuthenticationAPI.md#GetForRepositories) | **Get** /keys/latest/ssh/{keyId}/repos | Get repository SSH key
[**GetForRepository**](AuthenticationAPI.md#GetForRepository) | **Get** /keys/latest/projects/{projectKey}/repos/{repositorySlug}/ssh/{keyId} | Get repository SSH key
[**GetForRepository1**](AuthenticationAPI.md#GetForRepository1) | **Get** /keys/latest/projects/{projectKey}/repos/{repositorySlug}/ssh | Get repository SSH keys
[**GetIdp**](AuthenticationAPI.md#GetIdp) | **Get** /authconfig/latest/idps/{id} | Get IdP configuration
[**GetIdps**](AuthenticationAPI.md#GetIdps) | **Get** /authconfig/latest/idps | Get all configured IdPs
[**GetJitProvisionedUsers**](AuthenticationAPI.md#GetJitProvisionedUsers) | **Get** /authconfig/latest/jit-users | Get all JIT provisioned users
[**GetLoginOptions**](AuthenticationAPI.md#GetLoginOptions) | **Get** /authconfig/latest/login-options | Get available login options
[**GetSshKey**](AuthenticationAPI.md#GetSshKey) | **Get** /ssh/latest/keys/{keyId} | Get SSH key for user by keyId
[**GetSshKeys**](AuthenticationAPI.md#GetSshKeys) | **Get** /ssh/latest/keys | Get SSH keys for user
[**GetSshKeysForProject**](AuthenticationAPI.md#GetSshKeysForProject) | **Get** /keys/latest/projects/{projectKey}/ssh | Get SSH key
[**GetSsoManagementStatus**](AuthenticationAPI.md#GetSsoManagementStatus) | **Get** /tsv/latest/sso-management-status | Get SSO management status
[**GetStatus**](AuthenticationAPI.md#GetStatus) | **Get** /tsv/latest/status | Get two-step verification status
[**Put**](AuthenticationAPI.md#Put) | **Put** /basicauth/latest/config | Update basic auth configuration
[**RemoveIdp**](AuthenticationAPI.md#RemoveIdp) | **Delete** /authconfig/latest/idps/{id} | Delete IdP configuration
[**RevokeForProject**](AuthenticationAPI.md#RevokeForProject) | **Delete** /keys/latest/projects/{projectKey}/ssh/{keyId} | Revoke project SSH key
[**RevokeForRepository**](AuthenticationAPI.md#RevokeForRepository) | **Delete** /keys/latest/projects/{projectKey}/repos/{repositorySlug}/ssh/{keyId} | Revoke repository SSH key
[**RevokeMany**](AuthenticationAPI.md#RevokeMany) | **Delete** /keys/latest/ssh/{keyId} | Revoke project SSH key
[**RotateRecoverCode**](AuthenticationAPI.md#RotateRecoverCode) | **Post** /tsv/latest/totp/recovery-code/rotate | Rotate recovery code
[**SshSettings**](AuthenticationAPI.md#SshSettings) | **Get** /ssh/latest/settings | Get SSH settings
[**StartEnforcedEnrollment**](AuthenticationAPI.md#StartEnforcedEnrollment) | **Post** /tsv/latest/totp/start-enforced-enrollment | Start enforced enrollment in 2SV
[**StartEnrollmentUpdate**](AuthenticationAPI.md#StartEnrollmentUpdate) | **Post** /tsv/latest/totp/start-enrollment-update | Start authentication app update for 2SV
[**StartVoluntaryEnrollment**](AuthenticationAPI.md#StartVoluntaryEnrollment) | **Post** /tsv/latest/totp/start-voluntary-enrollment | Start voluntary enrollment in 2SV
[**Unenroll**](AuthenticationAPI.md#Unenroll) | **Delete** /tsv/latest/totp/unenroll | Uneroll current user from two-step verification
[**UnenrollUser**](AuthenticationAPI.md#UnenrollUser) | **Delete** /tsv/latest/totp/unenroll/user/{userName} | Unenroll specific user from two-step verification
[**UpdateAccessToken**](AuthenticationAPI.md#UpdateAccessToken) | **Post** /access-tokens/latest/projects/{projectKey}/{tokenId} | Update HTTP token
[**UpdateAccessToken1**](AuthenticationAPI.md#UpdateAccessToken1) | **Post** /access-tokens/latest/projects/{projectKey}/repos/{repositorySlug}/{tokenId} | Update HTTP token
[**UpdateAccessToken2**](AuthenticationAPI.md#UpdateAccessToken2) | **Post** /access-tokens/latest/users/{userSlug}/{tokenId} | Update HTTP token
[**UpdateConfig**](AuthenticationAPI.md#UpdateConfig) | **Patch** /authconfig/latest/sso | Update SSO configuration
[**UpdateIdp**](AuthenticationAPI.md#UpdateIdp) | **Patch** /authconfig/latest/idps/{id} | Update IdP configuration
[**UpdatePermission**](AuthenticationAPI.md#UpdatePermission) | **Put** /keys/latest/projects/{projectKey}/ssh/{keyId}/permission/{permission} | Update project SSH key permission
[**UpdatePermission1**](AuthenticationAPI.md#UpdatePermission1) | **Put** /keys/latest/projects/{projectKey}/repos/{repositorySlug}/ssh/{keyId}/permission/{permission} | Update repository SSH key permission
[**VerifyCode**](AuthenticationAPI.md#VerifyCode) | **Post** /tsv/latest/authenticate/totp-code | Authenticate using TOTP code



## AGet

> BasicAuthConfigEntity AGet(ctx).Execute()

Get basic auth configuration



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.AGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AGet`: BasicAuthConfigEntity
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAGetRequest struct via the builder pattern


### Return type

[**BasicAuthConfigEntity**](BasicAuthConfigEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddForProject

> RestSshAccessKey AddForProject(ctx, projectKey).RestSshAccessKey(restSshAccessKey).Execute()

Add project SSH key



### Example

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
	restSshAccessKey := *openapiclient.NewRestSshAccessKey() // RestSshAccessKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AddForProject(context.Background(), projectKey).RestSshAccessKey(restSshAccessKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AddForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddForProject`: RestSshAccessKey
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AddForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restSshAccessKey** | [**RestSshAccessKey**](RestSshAccessKey.md) |  | 

### Return type

[**RestSshAccessKey**](RestSshAccessKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddForRepository

> RestSshAccessKey AddForRepository(ctx, projectKey, repositorySlug).RestSshAccessKey(restSshAccessKey).Execute()

Add repository SSH key



### Example

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
	restSshAccessKey := *openapiclient.NewRestSshAccessKey() // RestSshAccessKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AddForRepository(context.Background(), projectKey, repositorySlug).RestSshAccessKey(restSshAccessKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AddForRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddForRepository`: RestSshAccessKey
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AddForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restSshAccessKey** | [**RestSshAccessKey**](RestSshAccessKey.md) |  | 

### Return type

[**RestSshAccessKey**](RestSshAccessKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddIdp

> IdpConfigEntity AddIdp(ctx).IdpConfigEntity(idpConfigEntity).Execute()

Create IdP configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	idpConfigEntity := *openapiclient.NewIdpConfigEntity() // IdpConfigEntity | The configuration of the new IdP to add. The ID must be null. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AddIdp(context.Background()).IdpConfigEntity(idpConfigEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AddIdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddIdp`: IdpConfigEntity
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AddIdp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idpConfigEntity** | [**IdpConfigEntity**](IdpConfigEntity.md) | The configuration of the new IdP to add. The ID must be null. | 

### Return type

[**IdpConfigEntity**](IdpConfigEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSshKey

> RestSshKey AddSshKey(ctx).User(user).AddSshKeyRequest(addSshKeyRequest).Execute()

Add SSH key for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	user := *openapiclient.NewRestSshKey() // RestSshKey | the username of the user to add the SSH key for. If no username is specified, the SSH key will be added for the current authenticated user. (optional)
	addSshKeyRequest := *openapiclient.NewAddSshKeyRequest() // AddSshKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AddSshKey(context.Background()).User(user).AddSshKeyRequest(addSshKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AddSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSshKey`: RestSshKey
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AddSshKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**RestSshKey**](RestSshKey.md) | the username of the user to add the SSH key for. If no username is specified, the SSH key will be added for the current authenticated user. | 
 **addSshKeyRequest** | [**AddSshKeyRequest**](AddSshKeyRequest.md) |  | 

### Return type

[**RestSshKey**](RestSshKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Authenticate

> AuthenticationResponse Authenticate(ctx).AuthenticationEntity(authenticationEntity).Execute()

Authenticate with 2SV



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	authenticationEntity := *openapiclient.NewAuthenticationEntity() // AuthenticationEntity |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Authenticate(context.Background()).AuthenticationEntity(authenticationEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Authenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Authenticate`: AuthenticationResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Authenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationEntity** | [**AuthenticationEntity**](AuthenticationEntity.md) |  | 

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateWithRecoveryCode

> AuthenticationResponse AuthenticateWithRecoveryCode(ctx).TotpRecoveryCodeAuthenticationDTO(totpRecoveryCodeAuthenticationDTO).Execute()

Authenticate using recovery code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	totpRecoveryCodeAuthenticationDTO := *openapiclient.NewTotpRecoveryCodeAuthenticationDTO() // TotpRecoveryCodeAuthenticationDTO | A request containing a recovery code for the specified user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.AuthenticateWithRecoveryCode(context.Background()).TotpRecoveryCodeAuthenticationDTO(totpRecoveryCodeAuthenticationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthenticateWithRecoveryCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticateWithRecoveryCode`: AuthenticationResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthenticateWithRecoveryCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateWithRecoveryCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **totpRecoveryCodeAuthenticationDTO** | [**TotpRecoveryCodeAuthenticationDTO**](TotpRecoveryCodeAuthenticationDTO.md) | A request containing a recovery code for the specified user. | 

### Return type

[**AuthenticationResponse**](AuthenticationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteAuthenticationChange

> TotpUserEnrollmentDTO CompleteAuthenticationChange(ctx).TotpCodeVerificationDTO(totpCodeVerificationDTO).Execute()

Complete authentication app update for 2SV



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	totpCodeVerificationDTO := *openapiclient.NewTotpCodeVerificationDTO() // TotpCodeVerificationDTO | A request containing a TOTP code for the given user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.CompleteAuthenticationChange(context.Background()).TotpCodeVerificationDTO(totpCodeVerificationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CompleteAuthenticationChange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteAuthenticationChange`: TotpUserEnrollmentDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.CompleteAuthenticationChange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteAuthenticationChangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **totpCodeVerificationDTO** | [**TotpCodeVerificationDTO**](TotpCodeVerificationDTO.md) | A request containing a TOTP code for the given user. | 

### Return type

[**TotpUserEnrollmentDTO**](TotpUserEnrollmentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteEnforcedEnrollment

> TotpRecoveryCodeDTO CompleteEnforcedEnrollment(ctx).TotpCodeVerificationDTO(totpCodeVerificationDTO).Execute()

Complete enforced enrollment in 2SV



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	totpCodeVerificationDTO := *openapiclient.NewTotpCodeVerificationDTO() // TotpCodeVerificationDTO | A request containing a TOTP code for the given user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.CompleteEnforcedEnrollment(context.Background()).TotpCodeVerificationDTO(totpCodeVerificationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CompleteEnforcedEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteEnforcedEnrollment`: TotpRecoveryCodeDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.CompleteEnforcedEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteEnforcedEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **totpCodeVerificationDTO** | [**TotpCodeVerificationDTO**](TotpCodeVerificationDTO.md) | A request containing a TOTP code for the given user. | 

### Return type

[**TotpRecoveryCodeDTO**](TotpRecoveryCodeDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteVoluntaryEnrollment

> TotpUserEnrollmentDTO CompleteVoluntaryEnrollment(ctx).TotpCodeVerificationDTO(totpCodeVerificationDTO).Execute()

Complete voluntary enrollment in 2SV



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	totpCodeVerificationDTO := *openapiclient.NewTotpCodeVerificationDTO() // TotpCodeVerificationDTO | A request containing a TOTP code for the given user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.CompleteVoluntaryEnrollment(context.Background()).TotpCodeVerificationDTO(totpCodeVerificationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CompleteVoluntaryEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteVoluntaryEnrollment`: TotpUserEnrollmentDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.CompleteVoluntaryEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteVoluntaryEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **totpCodeVerificationDTO** | [**TotpCodeVerificationDTO**](TotpCodeVerificationDTO.md) | A request containing a TOTP code for the given user. | 

### Return type

[**TotpUserEnrollmentDTO**](TotpUserEnrollmentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessToken

> RestRawAccessToken CreateAccessToken(ctx, projectKey).RestAccessTokenRequest(restAccessTokenRequest).Execute()

Create project HTTP token



### Example

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
	restAccessTokenRequest := *openapiclient.NewRestAccessTokenRequest() // RestAccessTokenRequest | The request containing the details of the access token to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.CreateAccessToken(context.Background(), projectKey).RestAccessTokenRequest(restAccessTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CreateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessToken`: RestRawAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.CreateAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restAccessTokenRequest** | [**RestAccessTokenRequest**](RestAccessTokenRequest.md) | The request containing the details of the access token to create. | 

### Return type

[**RestRawAccessToken**](RestRawAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessToken1

> RestRawAccessToken CreateAccessToken1(ctx, projectKey, repositorySlug).RestAccessTokenRequest(restAccessTokenRequest).Execute()

Create repository HTTP token



### Example

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
	restAccessTokenRequest := *openapiclient.NewRestAccessTokenRequest() // RestAccessTokenRequest | The request containing the details of the access token to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.CreateAccessToken1(context.Background(), projectKey, repositorySlug).RestAccessTokenRequest(restAccessTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CreateAccessToken1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessToken1`: RestRawAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.CreateAccessToken1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessToken1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restAccessTokenRequest** | [**RestAccessTokenRequest**](RestAccessTokenRequest.md) | The request containing the details of the access token to create. | 

### Return type

[**RestRawAccessToken**](RestRawAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessToken2

> RestRawAccessToken CreateAccessToken2(ctx, userSlug).RestAccessTokenRequest(restAccessTokenRequest).Execute()

Create personal HTTP token



### Example

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
	restAccessTokenRequest := *openapiclient.NewRestAccessTokenRequest() // RestAccessTokenRequest | The request containing the details of the access token to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.CreateAccessToken2(context.Background(), userSlug).RestAccessTokenRequest(restAccessTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CreateAccessToken2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessToken2`: RestRawAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.CreateAccessToken2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessToken2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restAccessTokenRequest** | [**RestAccessTokenRequest**](RestAccessTokenRequest.md) | The request containing the details of the access token to create. | 

### Return type

[**RestRawAccessToken**](RestRawAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteById

> DeleteById(ctx, projectKey, tokenId).Execute()

Delete a HTTP token



### Example

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
	tokenId := "tokenId_example" // string | The token id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.DeleteById(context.Background(), projectKey, tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeleteById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**tokenId** | **string** | The token id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByIdRequest struct via the builder pattern


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


## DeleteById1

> DeleteById1(ctx, projectKey, tokenId, repositorySlug).Execute()

Delete a HTTP token



### Example

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
	tokenId := "tokenId_example" // string | The token id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.DeleteById1(context.Background(), projectKey, tokenId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeleteById1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**tokenId** | **string** | The token id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteById1Request struct via the builder pattern


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


## DeleteById2

> DeleteById2(ctx, tokenId, userSlug).Execute()

Delete a HTTP token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	tokenId := "tokenId_example" // string | The token id.
	userSlug := "userSlug_example" // string | The user slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.DeleteById2(context.Background(), tokenId, userSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeleteById2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | The token id. | 
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteById2Request struct via the builder pattern


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


## DeleteSshKey

> DeleteSshKey(ctx, keyId).Execute()

Remove SSH key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	keyId := "keyId_example" // string | the id of the key to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.DeleteSshKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeleteSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | the id of the key to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshKeyRequest struct via the builder pattern


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


## DeleteSshKeys

> DeleteSshKeys(ctx).UserName(userName).User(user).Execute()

Delete all user SSH key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userName := "userName_example" // string | the username of the user to delete the keys for. If no username is specified, the SSH keys will be deleted for the current authenticated user. (optional)
	user := "user_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.DeleteSshKeys(context.Background()).UserName(userName).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeleteSshKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSshKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | the username of the user to delete the keys for. If no username is specified, the SSH keys will be deleted for the current authenticated user. | 
 **user** | **string** |  | 

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


## ElevatePermissionsWithPassword

> ElevatePermissionsWithPassword(ctx).ActionType(actionType).TotpElevationRestDTO(totpElevationRestDTO).Execute()

Create elevated session with password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	actionType := "actionType_example" // string | The type of action being performed. (optional)
	totpElevationRestDTO := *openapiclient.NewTotpElevationRestDTO() // TotpElevationRestDTO | A request containing the password for the currently authenticated user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.ElevatePermissionsWithPassword(context.Background()).ActionType(actionType).TotpElevationRestDTO(totpElevationRestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ElevatePermissionsWithPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiElevatePermissionsWithPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionType** | **string** | The type of action being performed. | 
 **totpElevationRestDTO** | [**TotpElevationRestDTO**](TotpElevationRestDTO.md) | A request containing the password for the currently authenticated user. | 

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


## ElevatePermissionsWithRecoveryCode

> TotpRecoveryCodeDTO ElevatePermissionsWithRecoveryCode(ctx).ActionType(actionType).TotpRecoveryCodeDTO(totpRecoveryCodeDTO).Execute()

Create elevated session with recovery code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	actionType := "actionType_example" // string | The type of action being performed. (optional)
	totpRecoveryCodeDTO := *openapiclient.NewTotpRecoveryCodeDTO() // TotpRecoveryCodeDTO | A request containing a recovery code for the currently authenticated user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.ElevatePermissionsWithRecoveryCode(context.Background()).ActionType(actionType).TotpRecoveryCodeDTO(totpRecoveryCodeDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ElevatePermissionsWithRecoveryCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ElevatePermissionsWithRecoveryCode`: TotpRecoveryCodeDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.ElevatePermissionsWithRecoveryCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiElevatePermissionsWithRecoveryCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionType** | **string** | The type of action being performed. | 
 **totpRecoveryCodeDTO** | [**TotpRecoveryCodeDTO**](TotpRecoveryCodeDTO.md) | A request containing a recovery code for the currently authenticated user. | 

### Return type

[**TotpRecoveryCodeDTO**](TotpRecoveryCodeDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ElevatePermissionsWithTotp

> ElevatePermissionsWithTotp(ctx).ActionType(actionType).TotpElevationRestDTO(totpElevationRestDTO).Execute()

Create elevated session with TOTP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	actionType := "actionType_example" // string | The type of action being performed. (optional)
	totpElevationRestDTO := *openapiclient.NewTotpElevationRestDTO() // TotpElevationRestDTO | A request containing a TOTP code for the given user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.ElevatePermissionsWithTotp(context.Background()).ActionType(actionType).TotpElevationRestDTO(totpElevationRestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ElevatePermissionsWithTotp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiElevatePermissionsWithTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionType** | **string** | The type of action being performed. | 
 **totpElevationRestDTO** | [**TotpElevationRestDTO**](TotpElevationRestDTO.md) | A request containing a TOTP code for the given user. | 

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


## GetAllAccessTokens

> GetAllAccessTokens200Response GetAllAccessTokens(ctx, projectKey).Start(start).Limit(limit).Execute()

Get project HTTP tokens



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.GetAllAccessTokens(context.Background(), projectKey).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetAllAccessTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAccessTokens`: GetAllAccessTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetAllAccessTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccessTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetAllAccessTokens200Response**](GetAllAccessTokens200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAccessTokens1

> GetAllAccessTokens200Response GetAllAccessTokens1(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Execute()

Get repository HTTP tokens



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.GetAllAccessTokens1(context.Background(), projectKey, repositorySlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetAllAccessTokens1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAccessTokens1`: GetAllAccessTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetAllAccessTokens1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccessTokens1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetAllAccessTokens200Response**](GetAllAccessTokens200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAccessTokens2

> GetAllAccessTokens200Response GetAllAccessTokens2(ctx, userSlug).Start(start).Limit(limit).Execute()

Get personal HTTP tokens



### Example

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
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetAllAccessTokens2(context.Background(), userSlug).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetAllAccessTokens2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAccessTokens2`: GetAllAccessTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetAllAccessTokens2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccessTokens2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetAllAccessTokens200Response**](GetAllAccessTokens200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetById

> RestAccessToken GetById(ctx, projectKey, tokenId).Execute()

Get HTTP token by ID



### Example

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
	tokenId := "tokenId_example" // string | The token id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetById(context.Background(), projectKey, tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetById`: RestAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**tokenId** | **string** | The token id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestAccessToken**](RestAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetById1

> RestAccessToken GetById1(ctx, projectKey, tokenId, repositorySlug).Execute()

Get HTTP token by ID



### Example

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
	tokenId := "tokenId_example" // string | The token id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetById1(context.Background(), projectKey, tokenId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetById1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetById1`: RestAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetById1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**tokenId** | **string** | The token id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetById1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestAccessToken**](RestAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetById2

> RestAccessToken GetById2(ctx, tokenId, userSlug).Execute()

Get HTTP token by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	tokenId := "tokenId_example" // string | The token id.
	userSlug := "userSlug_example" // string | The user slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetById2(context.Background(), tokenId, userSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetById2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetById2`: RestAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetById2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | The token id. | 
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetById2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestAccessToken**](RestAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptchaData

> CaptchaDataEntity GetCaptchaData(ctx).Execute()

Get CAPTCHA challenge



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.GetCaptchaData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetCaptchaData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptchaData`: CaptchaDataEntity
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetCaptchaData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptchaDataRequest struct via the builder pattern


### Return type

[**CaptchaDataEntity**](CaptchaDataEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfig

> SsoConfigEntity GetConfig(ctx).Execute()

Get SSO configuration



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: SsoConfigEntity
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


### Return type

[**SsoConfigEntity**](SsoConfigEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElevatedPermissionStatus

> GetElevatedPermissionStatus(ctx).ActionType(actionType).Execute()

Get elevated session status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	actionType := "actionType_example" // string | The type of action being performed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.GetElevatedPermissionStatus(context.Background()).ActionType(actionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetElevatedPermissionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetElevatedPermissionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionType** | **string** | The type of action being performed. | 

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


## GetForProject

> RestSshAccessKey GetForProject(ctx, projectKey, keyId).Execute()

Get project SSH key



### Example

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
	keyId := "keyId_example" // string | The key id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetForProject(context.Background(), projectKey, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForProject`: RestSshAccessKey
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**keyId** | **string** | The key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RestSshAccessKey**](RestSshAccessKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForProjects

> GetForProjects(ctx, keyId).Execute()

Get project SSH keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	keyId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.GetForProjects(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetForProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForProjectsRequest struct via the builder pattern


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


## GetForRepositories

> GetForRepositories(ctx, keyId).WithRestrictions(withRestrictions).Execute()

Get repository SSH key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	keyId := "keyId_example" // string | The key id
	withRestrictions := "withRestrictions_example" // string | Include the readOnly field. The `readOnly` field is contextual for the user making the request. `readOnly` returns true if there is a restriction and the user does not have`PROJECT_ADMIN` access for the repository the key is associated with. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.GetForRepositories(context.Background(), keyId).WithRestrictions(withRestrictions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetForRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withRestrictions** | **string** | Include the readOnly field. The &#x60;readOnly&#x60; field is contextual for the user making the request. &#x60;readOnly&#x60; returns true if there is a restriction and the user does not have&#x60;PROJECT_ADMIN&#x60; access for the repository the key is associated with. | 

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


## GetForRepository

> RestSshAccessKey GetForRepository(ctx, projectKey, keyId, repositorySlug).Execute()

Get repository SSH key



### Example

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
	keyId := "keyId_example" // string | The key id
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetForRepository(context.Background(), projectKey, keyId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetForRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForRepository`: RestSshAccessKey
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetForRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**keyId** | **string** | The key id | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestSshAccessKey**](RestSshAccessKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForRepository1

> GetForRepository1200Response GetForRepository1(ctx, projectKey, repositorySlug).Filter(filter).Effective(effective).MinimumPermission(minimumPermission).Permission(permission).Start(start).Limit(limit).Execute()

Get repository SSH keys



### Example

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
	filter := "filter_example" // string | If specified only SSH access keys with a label prefixed with the supplied string will be returned (optional)
	effective := "effective_example" // string | Controls whether SSH access keys configured at the project level should be included in the results or not. When set to <code>true</code> all keys that have <em>access</em> to the repository (including project level keys) are included in the results. When set to <code>false</code>, only access keys configured for the specified <code>repository</code> are considered. Default is <code>false</code>. (optional)
	minimumPermission := "minimumPermission_example" // string | If specified only SSH access keys with at least the supplied permission will be returned. Default is <code>Permission.REPO_READ</code>. (optional)
	permission := "permission_example" // string |  (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetForRepository1(context.Background(), projectKey, repositorySlug).Filter(filter).Effective(effective).MinimumPermission(minimumPermission).Permission(permission).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetForRepository1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForRepository1`: GetForRepository1200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetForRepository1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForRepository1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | If specified only SSH access keys with a label prefixed with the supplied string will be returned | 
 **effective** | **string** | Controls whether SSH access keys configured at the project level should be included in the results or not. When set to &lt;code&gt;true&lt;/code&gt; all keys that have &lt;em&gt;access&lt;/em&gt; to the repository (including project level keys) are included in the results. When set to &lt;code&gt;false&lt;/code&gt;, only access keys configured for the specified &lt;code&gt;repository&lt;/code&gt; are considered. Default is &lt;code&gt;false&lt;/code&gt;. | 
 **minimumPermission** | **string** | If specified only SSH access keys with at least the supplied permission will be returned. Default is &lt;code&gt;Permission.REPO_READ&lt;/code&gt;. | 
 **permission** | **string** |  | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetForRepository1200Response**](GetForRepository1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdp

> IdpConfigEntity GetIdp(ctx, id).Execute()

Get IdP configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "id_example" // string | The ID of the IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetIdp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetIdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdp`: IdpConfigEntity
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdpConfigEntity**](IdpConfigEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdps

> GetIdps200Response GetIdps(ctx).Start(start).Limit(limit).Execute()

Get all configured IdPs



### Example

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
	limit := float32(50) // float32 | Number of items to return. If not passed, a page size of 50 is used. A limit of -1 means that the request will fetch all results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetIdps(context.Background()).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetIdps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdps`: GetIdps200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetIdps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 50 is used. A limit of -1 means that the request will fetch all results. | 

### Return type

[**GetIdps200Response**](GetIdps200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJitProvisionedUsers

> JitUserEntity GetJitProvisionedUsers(ctx).Execute()

Get all JIT provisioned users



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.GetJitProvisionedUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetJitProvisionedUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJitProvisionedUsers`: JitUserEntity
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetJitProvisionedUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJitProvisionedUsersRequest struct via the builder pattern


### Return type

[**JitUserEntity**](JitUserEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginOptions

> GetLoginOptions200Response GetLoginOptions(ctx).Start(start).Limit(limit).Execute()

Get available login options



### Example

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
	limit := float32(50) // float32 | Number of items to return. If not passed, a page size of 50 is used. A limit of -1 means that the request will fetch all results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetLoginOptions(context.Background()).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetLoginOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoginOptions`: GetLoginOptions200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetLoginOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 50 is used. A limit of -1 means that the request will fetch all results. | 

### Return type

[**GetLoginOptions200Response**](GetLoginOptions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshKey

> RestSshKey GetSshKey(ctx, keyId).Execute()

Get SSH key for user by keyId



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	keyId := "keyId_example" // string | the ID of the key to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetSshKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshKey`: RestSshKey
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetSshKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | the ID of the key to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestSshKey**](RestSshKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshKeys

> GetSshKeys200Response GetSshKeys(ctx).UserName(userName).User(user).Start(start).Limit(limit).Execute()

Get SSH keys for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userName := "userName_example" // string | the username of the user to retrieve the keys for. If no username is specified, the SSH keys will be retrieved for the current authenticated user. (optional)
	user := "user_example" // string |  (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetSshKeys(context.Background()).UserName(userName).User(user).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetSshKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshKeys`: GetSshKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetSshKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | the username of the user to retrieve the keys for. If no username is specified, the SSH keys will be retrieved for the current authenticated user. | 
 **user** | **string** |  | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetSshKeys200Response**](GetSshKeys200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSshKeysForProject

> GetForRepository1200Response GetSshKeysForProject(ctx, projectKey).Filter(filter).Permission(permission).Start(start).Limit(limit).Execute()

Get SSH key



### Example

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
	filter := "filter_example" // string | If specified only SSH access keys with a label prefixed with the supplied string will be returned. (optional)
	permission := "permission_example" // string | If specified only SSH access keys with at least the supplied permission will be returned Default is PROJECT_READ. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetSshKeysForProject(context.Background(), projectKey).Filter(filter).Permission(permission).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetSshKeysForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSshKeysForProject`: GetForRepository1200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetSshKeysForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSshKeysForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | If specified only SSH access keys with a label prefixed with the supplied string will be returned. | 
 **permission** | **string** | If specified only SSH access keys with at least the supplied permission will be returned Default is PROJECT_READ. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetForRepository1200Response**](GetForRepository1200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsoManagementStatus

> SsoManagementStatusDTO GetSsoManagementStatus(ctx).Execute()

Get SSO management status



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.GetSsoManagementStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetSsoManagementStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsoManagementStatus`: SsoManagementStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetSsoManagementStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsoManagementStatusRequest struct via the builder pattern


### Return type

[**SsoManagementStatusDTO**](SsoManagementStatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> StatusDTO GetStatus(ctx).Execute()

Get two-step verification status



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.GetStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: StatusDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**StatusDTO**](StatusDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Put

> Put(ctx).BasicAuthConfigEntity(basicAuthConfigEntity).Execute()

Update basic auth configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	basicAuthConfigEntity := *openapiclient.NewBasicAuthConfigEntity() // BasicAuthConfigEntity | A request containing the new basic authentication configuration. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.Put(context.Background()).BasicAuthConfigEntity(basicAuthConfigEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Put``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **basicAuthConfigEntity** | [**BasicAuthConfigEntity**](BasicAuthConfigEntity.md) | A request containing the new basic authentication configuration. | 

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


## RemoveIdp

> IdpConfigEntity RemoveIdp(ctx, id).Execute()

Delete IdP configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "id_example" // string | The ID of the IdP

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.RemoveIdp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.RemoveIdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveIdp`: IdpConfigEntity
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.RemoveIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdpConfigEntity**](IdpConfigEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeForProject

> RevokeForProject(ctx, projectKey, keyId).Execute()

Revoke project SSH key



### Example

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
	keyId := "keyId_example" // string | The key id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.RevokeForProject(context.Background(), projectKey, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.RevokeForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**keyId** | **string** | The key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeForProjectRequest struct via the builder pattern


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


## RevokeForRepository

> RevokeForRepository(ctx, projectKey, keyId, repositorySlug).Execute()

Revoke repository SSH key



### Example

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
	keyId := "keyId_example" // string | The key id
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.RevokeForRepository(context.Background(), projectKey, keyId, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.RevokeForRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**keyId** | **string** | The key id | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeForRepositoryRequest struct via the builder pattern


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


## RevokeMany

> RevokeMany(ctx, keyId).Execute()

Revoke project SSH key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	keyId := "keyId_example" // string | The identifier of the SSH key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.RevokeMany(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.RevokeMany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | The identifier of the SSH key | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeManyRequest struct via the builder pattern


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


## RotateRecoverCode

> TotpRecoveryCodeDTO RotateRecoverCode(ctx).Execute()

Rotate recovery code



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.RotateRecoverCode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.RotateRecoverCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateRecoverCode`: TotpRecoveryCodeDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.RotateRecoverCode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRotateRecoverCodeRequest struct via the builder pattern


### Return type

[**TotpRecoveryCodeDTO**](TotpRecoveryCodeDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SshSettings

> RestSshSettings SshSettings(ctx).Execute()

Get SSH settings



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.SshSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.SshSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SshSettings`: RestSshSettings
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.SshSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSshSettingsRequest struct via the builder pattern


### Return type

[**RestSshSettings**](RestSshSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartEnforcedEnrollment

> TotpUserEnrollmentDTO StartEnforcedEnrollment(ctx).ConversationDTO(conversationDTO).Execute()

Start enforced enrollment in 2SV



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	conversationDTO := *openapiclient.NewConversationDTO() // ConversationDTO |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.StartEnforcedEnrollment(context.Background()).ConversationDTO(conversationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.StartEnforcedEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartEnforcedEnrollment`: TotpUserEnrollmentDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.StartEnforcedEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartEnforcedEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversationDTO** | [**ConversationDTO**](ConversationDTO.md) |  | 

### Return type

[**TotpUserEnrollmentDTO**](TotpUserEnrollmentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartEnrollmentUpdate

> TotpUserEnrollmentDTO StartEnrollmentUpdate(ctx).Execute()

Start authentication app update for 2SV



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.StartEnrollmentUpdate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.StartEnrollmentUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartEnrollmentUpdate`: TotpUserEnrollmentDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.StartEnrollmentUpdate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartEnrollmentUpdateRequest struct via the builder pattern


### Return type

[**TotpUserEnrollmentDTO**](TotpUserEnrollmentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVoluntaryEnrollment

> TotpUserEnrollmentDTO StartVoluntaryEnrollment(ctx).Execute()

Start voluntary enrollment in 2SV



### Example

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
	resp, r, err := apiClient.AuthenticationAPI.StartVoluntaryEnrollment(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.StartVoluntaryEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartVoluntaryEnrollment`: TotpUserEnrollmentDTO
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.StartVoluntaryEnrollment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartVoluntaryEnrollmentRequest struct via the builder pattern


### Return type

[**TotpUserEnrollmentDTO**](TotpUserEnrollmentDTO.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unenroll

> Unenroll(ctx).Execute()

Uneroll current user from two-step verification



### Example

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
	r, err := apiClient.AuthenticationAPI.Unenroll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Unenroll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUnenrollRequest struct via the builder pattern


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


## UnenrollUser

> UnenrollUser(ctx, userName).Execute()

Unenroll specific user from two-step verification



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userName := "userName_example" // string | username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.UnenrollUser(context.Background(), userName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UnenrollUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userName** | **string** | username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnenrollUserRequest struct via the builder pattern


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


## UpdateAccessToken

> RestAccessToken UpdateAccessToken(ctx, projectKey, tokenId).RestAccessTokenRequest(restAccessTokenRequest).Execute()

Update HTTP token



### Example

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
	tokenId := "tokenId_example" // string | The token id.
	restAccessTokenRequest := *openapiclient.NewRestAccessTokenRequest() // RestAccessTokenRequest | The request containing the details of the access token to modify (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdateAccessToken(context.Background(), projectKey, tokenId).RestAccessTokenRequest(restAccessTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessToken`: RestAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdateAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**tokenId** | **string** | The token id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restAccessTokenRequest** | [**RestAccessTokenRequest**](RestAccessTokenRequest.md) | The request containing the details of the access token to modify | 

### Return type

[**RestAccessToken**](RestAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessToken1

> RestAccessToken UpdateAccessToken1(ctx, projectKey, tokenId, repositorySlug).RestAccessTokenRequest(restAccessTokenRequest).Execute()

Update HTTP token



### Example

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
	tokenId := "tokenId_example" // string | The token id.
	repositorySlug := "repositorySlug_example" // string | The repository slug.
	restAccessTokenRequest := *openapiclient.NewRestAccessTokenRequest() // RestAccessTokenRequest | The request containing the details of the access token to modify (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdateAccessToken1(context.Background(), projectKey, tokenId, repositorySlug).RestAccessTokenRequest(restAccessTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdateAccessToken1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessToken1`: RestAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdateAccessToken1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**tokenId** | **string** | The token id. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessToken1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restAccessTokenRequest** | [**RestAccessTokenRequest**](RestAccessTokenRequest.md) | The request containing the details of the access token to modify | 

### Return type

[**RestAccessToken**](RestAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessToken2

> RestAccessToken UpdateAccessToken2(ctx, tokenId, userSlug).RestAccessTokenRequest(restAccessTokenRequest).Execute()

Update HTTP token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	tokenId := "tokenId_example" // string | The token id.
	userSlug := "userSlug_example" // string | The user slug.
	restAccessTokenRequest := *openapiclient.NewRestAccessTokenRequest() // RestAccessTokenRequest | The request containing the details of the access token to modify (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdateAccessToken2(context.Background(), tokenId, userSlug).RestAccessTokenRequest(restAccessTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdateAccessToken2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessToken2`: RestAccessToken
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdateAccessToken2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | The token id. | 
**userSlug** | **string** | The user slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessToken2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restAccessTokenRequest** | [**RestAccessTokenRequest**](RestAccessTokenRequest.md) | The request containing the details of the access token to modify | 

### Return type

[**RestAccessToken**](RestAccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfig

> SsoConfigEntity UpdateConfig(ctx).SsoConfigEntity(ssoConfigEntity).Execute()

Update SSO configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	ssoConfigEntity := *openapiclient.NewSsoConfigEntity() // SsoConfigEntity | A request containing the SSO configuration to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdateConfig(context.Background()).SsoConfigEntity(ssoConfigEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdateConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfig`: SsoConfigEntity
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdateConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoConfigEntity** | [**SsoConfigEntity**](SsoConfigEntity.md) | A request containing the SSO configuration to update. | 

### Return type

[**SsoConfigEntity**](SsoConfigEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdp

> IdpConfigEntity UpdateIdp(ctx, id).IdpConfigEntity(idpConfigEntity).Execute()

Update IdP configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	id := "id_example" // string | The ID of the IdP
	idpConfigEntity := *openapiclient.NewIdpConfigEntity() // IdpConfigEntity | A request containing the IdP configuration to update. The ID must either be null or equal to the ID specified in the path. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdateIdp(context.Background(), id).IdpConfigEntity(idpConfigEntity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdateIdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIdp`: IdpConfigEntity
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdateIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the IdP | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idpConfigEntity** | [**IdpConfigEntity**](IdpConfigEntity.md) | A request containing the IdP configuration to update. The ID must either be null or equal to the ID specified in the path. | 

### Return type

[**IdpConfigEntity**](IdpConfigEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePermission

> RestSshAccessKey UpdatePermission(ctx, projectKey, keyId, permission).Execute()

Update project SSH key permission



### Example

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
	keyId := "keyId_example" // string | The newly created access key
	permission := "permission_example" // string | The new permission to be granted to the SSH key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdatePermission(context.Background(), projectKey, keyId, permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdatePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePermission`: RestSshAccessKey
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdatePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**keyId** | **string** | The newly created access key | 
**permission** | **string** | The new permission to be granted to the SSH key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RestSshAccessKey**](RestSshAccessKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePermission1

> RestSshAccessKey UpdatePermission1(ctx, projectKey, keyId, permission, repositorySlug).Execute()

Update repository SSH key permission



### Example

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
	keyId := "keyId_example" // string | The newly created access key
	permission := "permission_example" // string | The new permission to be granted to the SSH key
	repositorySlug := "repositorySlug_example" // string | The repository slug

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.UpdatePermission1(context.Background(), projectKey, keyId, permission, repositorySlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.UpdatePermission1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePermission1`: RestSshAccessKey
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.UpdatePermission1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key | 
**keyId** | **string** | The newly created access key | 
**permission** | **string** | The new permission to be granted to the SSH key | 
**repositorySlug** | **string** | The repository slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePermission1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**RestSshAccessKey**](RestSshAccessKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyCode

> VerifyCode(ctx).TotpCodeVerificationDTO(totpCodeVerificationDTO).Execute()

Authenticate using TOTP code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	totpCodeVerificationDTO := *openapiclient.NewTotpCodeVerificationDTO() // TotpCodeVerificationDTO | A request containing a TOTP code for the given user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.VerifyCode(context.Background()).TotpCodeVerificationDTO(totpCodeVerificationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.VerifyCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **totpCodeVerificationDTO** | [**TotpCodeVerificationDTO**](TotpCodeVerificationDTO.md) | A request containing a TOTP code for the given user. | 

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


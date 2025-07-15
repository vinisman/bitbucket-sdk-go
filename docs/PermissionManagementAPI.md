# \PermissionManagementAPI

All URIs are relative to *http://example.com:7990/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupToUser**](PermissionManagementAPI.md#AddGroupToUser) | **Post** /api/latest/admin/users/add-group | Add user to group
[**AddUserToGroup**](PermissionManagementAPI.md#AddUserToGroup) | **Post** /api/latest/admin/groups/add-user | Add user to group
[**AddUserToGroups**](PermissionManagementAPI.md#AddUserToGroups) | **Post** /api/latest/admin/users/add-groups | Add user to groups
[**AddUsersToGroup**](PermissionManagementAPI.md#AddUsersToGroup) | **Post** /api/latest/admin/groups/add-users | Add multiple users to group
[**ClearUserCaptchaChallenge**](PermissionManagementAPI.md#ClearUserCaptchaChallenge) | **Delete** /api/latest/admin/users/captcha | Clear CAPTCHA for user
[**CreateGroup**](PermissionManagementAPI.md#CreateGroup) | **Post** /api/latest/admin/groups | Create group
[**CreateUser**](PermissionManagementAPI.md#CreateUser) | **Post** /api/latest/admin/users | Create user
[**DeleteGroup**](PermissionManagementAPI.md#DeleteGroup) | **Delete** /api/latest/admin/groups | Remove group
[**DeleteUser**](PermissionManagementAPI.md#DeleteUser) | **Delete** /api/latest/admin/users | Remove user
[**EraseUser**](PermissionManagementAPI.md#EraseUser) | **Post** /api/latest/admin/users/erasure | Erase user information
[**FindGroupsForUser**](PermissionManagementAPI.md#FindGroupsForUser) | **Get** /api/latest/admin/users/more-members | Get groups for user
[**FindOtherGroupsForUser**](PermissionManagementAPI.md#FindOtherGroupsForUser) | **Get** /api/latest/admin/users/more-non-members | Find other groups for user
[**FindUsersInGroup**](PermissionManagementAPI.md#FindUsersInGroup) | **Get** /api/latest/admin/groups/more-members | Get group members
[**FindUsersNotInGroup**](PermissionManagementAPI.md#FindUsersNotInGroup) | **Get** /api/latest/admin/groups/more-non-members | Get members not in group
[**GetGroups**](PermissionManagementAPI.md#GetGroups) | **Get** /api/latest/groups | Get group names
[**GetGroups1**](PermissionManagementAPI.md#GetGroups1) | **Get** /api/latest/admin/groups | Get groups
[**GetGroupsWithAnyPermission**](PermissionManagementAPI.md#GetGroupsWithAnyPermission) | **Get** /api/latest/admin/permissions/groups | Get groups with a global permission
[**GetGroupsWithAnyPermission2**](PermissionManagementAPI.md#GetGroupsWithAnyPermission2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions/groups | Get groups with permission to repository
[**GetGroupsWithoutAnyPermission**](PermissionManagementAPI.md#GetGroupsWithoutAnyPermission) | **Get** /api/latest/admin/permissions/groups/none | Get groups with no global permission
[**GetGroupsWithoutAnyPermission2**](PermissionManagementAPI.md#GetGroupsWithoutAnyPermission2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions/groups/none | Get groups without repository permission
[**GetUserDirectories**](PermissionManagementAPI.md#GetUserDirectories) | **Get** /api/latest/admin/user-directories | Get directories
[**GetUsers1**](PermissionManagementAPI.md#GetUsers1) | **Get** /api/latest/admin/users | Get users
[**GetUsersWithAnyPermission**](PermissionManagementAPI.md#GetUsersWithAnyPermission) | **Get** /api/latest/admin/permissions/users | Get users with a global permission
[**GetUsersWithAnyPermission2**](PermissionManagementAPI.md#GetUsersWithAnyPermission2) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions/users | Get users with permission to repository
[**GetUsersWithoutAnyPermission**](PermissionManagementAPI.md#GetUsersWithoutAnyPermission) | **Get** /api/latest/admin/permissions/users/none | Get users with no global permission
[**GetUsersWithoutPermission1**](PermissionManagementAPI.md#GetUsersWithoutPermission1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions/users/none | Get users without repository permission
[**RemoveGroupFromUser**](PermissionManagementAPI.md#RemoveGroupFromUser) | **Post** /api/latest/admin/users/remove-group | Remove user from group
[**RemoveUserFromGroup**](PermissionManagementAPI.md#RemoveUserFromGroup) | **Post** /api/latest/admin/groups/remove-user | Remove user from group
[**RenameUser**](PermissionManagementAPI.md#RenameUser) | **Post** /api/latest/admin/users/rename | Rename user
[**RevokePermissions1**](PermissionManagementAPI.md#RevokePermissions1) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions | Revoke all repository permissions for users and groups
[**RevokePermissionsForGroup**](PermissionManagementAPI.md#RevokePermissionsForGroup) | **Delete** /api/latest/admin/permissions/groups | Revoke all global permissions for group
[**RevokePermissionsForGroup2**](PermissionManagementAPI.md#RevokePermissionsForGroup2) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions/groups | Revoke group repository permission
[**RevokePermissionsForUser**](PermissionManagementAPI.md#RevokePermissionsForUser) | **Delete** /api/latest/admin/permissions/users | Revoke all global permissions for user
[**RevokePermissionsForUser2**](PermissionManagementAPI.md#RevokePermissionsForUser2) | **Delete** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions/users | Revoke user repository permission
[**SearchPermissions1**](PermissionManagementAPI.md#SearchPermissions1) | **Get** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions/search | Search repository permissions
[**SetPermissionForGroup**](PermissionManagementAPI.md#SetPermissionForGroup) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions/groups | Update group repository permission
[**SetPermissionForGroups**](PermissionManagementAPI.md#SetPermissionForGroups) | **Put** /api/latest/admin/permissions/groups | Update global permission for group
[**SetPermissionForUser**](PermissionManagementAPI.md#SetPermissionForUser) | **Put** /api/latest/projects/{projectKey}/repos/{repositorySlug}/permissions/users | Update user repository permission
[**SetPermissionForUsers**](PermissionManagementAPI.md#SetPermissionForUsers) | **Put** /api/latest/admin/permissions/users | Update global permission for user
[**UpdateUserDetails**](PermissionManagementAPI.md#UpdateUserDetails) | **Put** /api/latest/admin/users | Update user details
[**UpdateUserPassword**](PermissionManagementAPI.md#UpdateUserPassword) | **Put** /api/latest/admin/users/credentials | Set password for user
[**ValidateErasable**](PermissionManagementAPI.md#ValidateErasable) | **Get** /api/latest/admin/users/erasure | Check user removal



## AddGroupToUser

> AddGroupToUser(ctx).GroupPickerContext(groupPickerContext).Execute()

Add user to group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	groupPickerContext := *openapiclient.NewGroupPickerContext() // GroupPickerContext |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.AddGroupToUser(context.Background()).GroupPickerContext(groupPickerContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.AddGroupToUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupToUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupPickerContext** | [**GroupPickerContext**](GroupPickerContext.md) |  | 

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


## AddUserToGroup

> AddUserToGroup(ctx).UserPickerContext(userPickerContext).Execute()

Add user to group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userPickerContext := *openapiclient.NewUserPickerContext() // UserPickerContext |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.AddUserToGroup(context.Background()).UserPickerContext(userPickerContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.AddUserToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPickerContext** | [**UserPickerContext**](UserPickerContext.md) |  | 

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


## AddUserToGroups

> AddUserToGroups(ctx).UserAndGroups(userAndGroups).Execute()

Add user to groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userAndGroups := *openapiclient.NewUserAndGroups([]string{"Groups_example"}) // UserAndGroups |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.AddUserToGroups(context.Background()).UserAndGroups(userAndGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.AddUserToGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAndGroups** | [**UserAndGroups**](UserAndGroups.md) |  | 

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


## AddUsersToGroup

> AddUsersToGroup(ctx).GroupAndUsers(groupAndUsers).Execute()

Add multiple users to group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	groupAndUsers := *openapiclient.NewGroupAndUsers([]string{"Users_example"}) // GroupAndUsers |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.AddUsersToGroup(context.Background()).GroupAndUsers(groupAndUsers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.AddUsersToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUsersToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupAndUsers** | [**GroupAndUsers**](GroupAndUsers.md) |  | 

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


## ClearUserCaptchaChallenge

> ClearUserCaptchaChallenge(ctx).Name(name).Execute()

Clear CAPTCHA for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := "name_example" // string | The username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.ClearUserCaptchaChallenge(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.ClearUserCaptchaChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearUserCaptchaChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The username | 

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


## CreateGroup

> RestDetailedGroup CreateGroup(ctx).Name(name).Execute()

Create group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := "name_example" // string | Name of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.CreateGroup(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: RestDetailedGroup
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of the group. | 

### Return type

[**RestDetailedGroup**](RestDetailedGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> CreateUser(ctx).EmailAddress(emailAddress).DisplayName(displayName).Name(name).Password(password).AddToDefaultGroup(addToDefaultGroup).Notify(notify).Execute()

Create user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	emailAddress := "emailAddress_example" // string | The e-mail address for the new user.
	displayName := "displayName_example" // string | The display name for the new user.
	name := "name_example" // string | The username for the new user.
	password := "password_example" // string | The password for the new user. Required if the <code>notify</code> parameter is not present or is set to <code>false</false> (optional)
	addToDefaultGroup := true // bool | Set <code>true</code> to add the user to the default group, which can be used to grant them a set of initial permissions; otherwise, <code>false</code> to not add them to a group. (optional) (default to true)
	notify := true // bool | If present and not <code>false</code> instead of requiring a password, the create user will be notified via email their account has been created and requires a password to be reset. This option can only be used if a mail server has been configured. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.CreateUser(context.Background()).EmailAddress(emailAddress).DisplayName(displayName).Name(name).Password(password).AddToDefaultGroup(addToDefaultGroup).Notify(notify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailAddress** | **string** | The e-mail address for the new user. | 
 **displayName** | **string** | The display name for the new user. | 
 **name** | **string** | The username for the new user. | 
 **password** | **string** | The password for the new user. Required if the &lt;code&gt;notify&lt;/code&gt; parameter is not present or is set to &lt;code&gt;false&lt;/false&gt; | 
 **addToDefaultGroup** | **bool** | Set &lt;code&gt;true&lt;/code&gt; to add the user to the default group, which can be used to grant them a set of initial permissions; otherwise, &lt;code&gt;false&lt;/code&gt; to not add them to a group. | [default to true]
 **notify** | **bool** | If present and not &lt;code&gt;false&lt;/code&gt; instead of requiring a password, the create user will be notified via email their account has been created and requires a password to be reset. This option can only be used if a mail server has been configured. | 

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


## DeleteGroup

> RestDetailedGroup DeleteGroup(ctx).Name(name).Execute()

Remove group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := "name_example" // string | The name identifying the group to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.DeleteGroup(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroup`: RestDetailedGroup
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.DeleteGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name identifying the group to delete. | 

### Return type

[**RestDetailedGroup**](RestDetailedGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> RestDetailedUser DeleteUser(ctx).Name(name).Execute()

Remove user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := "name_example" // string | The username identifying the user to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.DeleteUser(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUser`: RestDetailedUser
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The username identifying the user to delete. | 

### Return type

[**RestDetailedUser**](RestDetailedUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EraseUser

> RestErasedUser EraseUser(ctx).Name(name).Execute()

Erase user information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := "name_example" // string | The username identifying the user to erase.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.EraseUser(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.EraseUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EraseUser`: RestErasedUser
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.EraseUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEraseUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The username identifying the user to erase. | 

### Return type

[**RestErasedUser**](RestErasedUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindGroupsForUser

> FindUsersInGroup200Response FindGroupsForUser(ctx).Context(context).Filter(filter).Start(start).Limit(limit).Execute()

Get groups for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	context := "context_example" // string | The group which should be used to locate members.
	filter := "filter_example" // string | If specified only users with usernames, display names or email addresses containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.FindGroupsForUser(context.Background()).Context(context).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.FindGroupsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindGroupsForUser`: FindUsersInGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.FindGroupsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindGroupsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string** | The group which should be used to locate members. | 
 **filter** | **string** | If specified only users with usernames, display names or email addresses containing the supplied string will be returned. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**FindUsersInGroup200Response**](FindUsersInGroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindOtherGroupsForUser

> GetGroups1200Response FindOtherGroupsForUser(ctx).Context(context).Filter(filter).Start(start).Limit(limit).Execute()

Find other groups for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	context := "context_example" // string | The user which should be used to locate groups.
	filter := "filter_example" // string | If specified only groups with names containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.FindOtherGroupsForUser(context.Background()).Context(context).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.FindOtherGroupsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindOtherGroupsForUser`: GetGroups1200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.FindOtherGroupsForUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindOtherGroupsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string** | The user which should be used to locate groups. | 
 **filter** | **string** | If specified only groups with names containing the supplied string will be returned. | 
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


## FindUsersInGroup

> FindUsersInGroup200Response FindUsersInGroup(ctx).Context(context).Filter(filter).Start(start).Limit(limit).Execute()

Get group members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	context := "context_example" // string | The group which should be used to locate members.
	filter := "filter_example" // string | If specified only users with usernames, display names or email addresses containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.FindUsersInGroup(context.Background()).Context(context).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.FindUsersInGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUsersInGroup`: FindUsersInGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.FindUsersInGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersInGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string** | The group which should be used to locate members. | 
 **filter** | **string** | If specified only users with usernames, display names or email addresses containing the supplied string will be returned. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**FindUsersInGroup200Response**](FindUsersInGroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsersNotInGroup

> FindUsersInGroup200Response FindUsersNotInGroup(ctx).Context(context).Filter(filter).Start(start).Limit(limit).Execute()

Get members not in group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	context := "context_example" // string | The group which should be used to locate members.
	filter := "filter_example" // string | If specified only users with usernames, display names or email addresses containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.FindUsersNotInGroup(context.Background()).Context(context).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.FindUsersNotInGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUsersNotInGroup`: FindUsersInGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.FindUsersNotInGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersNotInGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **string** | The group which should be used to locate members. | 
 **filter** | **string** | If specified only users with usernames, display names or email addresses containing the supplied string will be returned. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**FindUsersInGroup200Response**](FindUsersInGroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> GetGroups200Response GetGroups(ctx).Filter(filter).Start(start).Limit(limit).Execute()

Get group names



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "filter_example" // string |  (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetGroups(context.Background()).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroups`: GetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |  | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**GetGroups200Response**](GetGroups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups1

> GetGroups1200Response GetGroups1(ctx).Filter(filter).Start(start).Limit(limit).Execute()

Get groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "filter_example" // string | If specified only group names containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetGroups1(context.Background()).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetGroups1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroups1`: GetGroups1200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetGroups1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroups1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | If specified only group names containing the supplied string will be returned. | 
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


## GetGroupsWithAnyPermission

> GetGroupsWithAnyPermission200Response GetGroupsWithAnyPermission(ctx).Filter(filter).Start(start).Limit(limit).Execute()

Get groups with a global permission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "filter_example" // string | If specified only group names containing the supplied string will be returned (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetGroupsWithAnyPermission(context.Background()).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetGroupsWithAnyPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithAnyPermission`: GetGroupsWithAnyPermission200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetGroupsWithAnyPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithAnyPermissionRequest struct via the builder pattern


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


## GetGroupsWithAnyPermission2

> GetGroupsWithAnyPermission200Response GetGroupsWithAnyPermission2(ctx, projectKey, repositorySlug).Filter(filter).Start(start).Limit(limit).Execute()

Get groups with permission to repository



### Example

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
	filter := "filter_example" // string | If specified only group names containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetGroupsWithAnyPermission2(context.Background(), projectKey, repositorySlug).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetGroupsWithAnyPermission2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithAnyPermission2`: GetGroupsWithAnyPermission200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetGroupsWithAnyPermission2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithAnyPermission2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | If specified only group names containing the supplied string will be returned. | 
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


## GetGroupsWithoutAnyPermission

> GetGroups1200Response GetGroupsWithoutAnyPermission(ctx).Filter(filter).Start(start).Limit(limit).Execute()

Get groups with no global permission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "filter_example" // string | If specified only user names containing the supplied string will be returned (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetGroupsWithoutAnyPermission(context.Background()).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetGroupsWithoutAnyPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithoutAnyPermission`: GetGroups1200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetGroupsWithoutAnyPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithoutAnyPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | If specified only user names containing the supplied string will be returned | 
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


## GetGroupsWithoutAnyPermission2

> GetGroups1200Response GetGroupsWithoutAnyPermission2(ctx, projectKey, repositorySlug).Filter(filter).Start(start).Limit(limit).Execute()

Get groups without repository permission



### Example

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
	filter := "filter_example" // string | If specified only group names containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetGroupsWithoutAnyPermission2(context.Background(), projectKey, repositorySlug).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetGroupsWithoutAnyPermission2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithoutAnyPermission2`: GetGroups1200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetGroupsWithoutAnyPermission2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithoutAnyPermission2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | If specified only group names containing the supplied string will be returned. | 
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


## GetUserDirectories

> RestUserDirectory GetUserDirectories(ctx).IncludeInactive(includeInactive).Execute()

Get directories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	includeInactive := "includeInactive_example" // string | Set <code>true</code> to include inactive directories; otherwise, <code>false</code> to only return active directories. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetUserDirectories(context.Background()).IncludeInactive(includeInactive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetUserDirectories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserDirectories`: RestUserDirectory
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetUserDirectories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDirectoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeInactive** | **string** | Set &lt;code&gt;true&lt;/code&gt; to include inactive directories; otherwise, &lt;code&gt;false&lt;/code&gt; to only return active directories. | 

### Return type

[**RestUserDirectory**](RestUserDirectory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers1

> FindUsersInGroup200Response GetUsers1(ctx).Filter(filter).Start(start).Limit(limit).Execute()

Get users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "filter_example" // string | If specified only users with usernames, display name or email addresses containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetUsers1(context.Background()).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetUsers1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers1`: FindUsersInGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetUsers1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsers1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | If specified only users with usernames, display name or email addresses containing the supplied string will be returned. | 
 **start** | **float32** | Start number for the page (inclusive). If not passed, first page is assumed. | 
 **limit** | **float32** | Number of items to return. If not passed, a page size of 25 is used. | 

### Return type

[**FindUsersInGroup200Response**](FindUsersInGroup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersWithAnyPermission

> GetGroupsWithAnyPermission200Response GetUsersWithAnyPermission(ctx).Filter(filter).Start(start).Limit(limit).Execute()

Get users with a global permission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "filter_example" // string | If specified only user names containing the supplied string will be returned (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetUsersWithAnyPermission(context.Background()).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetUsersWithAnyPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersWithAnyPermission`: GetGroupsWithAnyPermission200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetUsersWithAnyPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersWithAnyPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | If specified only user names containing the supplied string will be returned | 
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


## GetUsersWithAnyPermission2

> GetUsersWithAnyPermission1200Response GetUsersWithAnyPermission2(ctx, projectKey, repositorySlug).Filter(filter).Start(start).Limit(limit).Execute()

Get users with permission to repository



### Example

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
	filter := "filter_example" // string | If specified only user names containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetUsersWithAnyPermission2(context.Background(), projectKey, repositorySlug).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetUsersWithAnyPermission2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersWithAnyPermission2`: GetUsersWithAnyPermission1200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetUsersWithAnyPermission2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersWithAnyPermission2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | If specified only user names containing the supplied string will be returned. | 
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


## GetUsersWithoutAnyPermission

> GetUsersWithoutAnyPermission200Response GetUsersWithoutAnyPermission(ctx).Filter(filter).Start(start).Limit(limit).Execute()

Get users with no global permission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	filter := "filter_example" // string | If specified only user names containing the supplied string will be returned (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetUsersWithoutAnyPermission(context.Background()).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetUsersWithoutAnyPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersWithoutAnyPermission`: GetUsersWithoutAnyPermission200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetUsersWithoutAnyPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersWithoutAnyPermissionRequest struct via the builder pattern


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


## GetUsersWithoutPermission1

> GetUsersWithoutAnyPermission200Response GetUsersWithoutPermission1(ctx, projectKey, repositorySlug).Filter(filter).Start(start).Limit(limit).Execute()

Get users without repository permission



### Example

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
	filter := "filter_example" // string | If specified only user names containing the supplied string will be returned. (optional)
	start := float32(0) // float32 | Start number for the page (inclusive). If not passed, first page is assumed. (optional)
	limit := float32(25) // float32 | Number of items to return. If not passed, a page size of 25 is used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.GetUsersWithoutPermission1(context.Background(), projectKey, repositorySlug).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.GetUsersWithoutPermission1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersWithoutPermission1`: GetUsersWithoutAnyPermission200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.GetUsersWithoutPermission1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** | The project key. | 
**repositorySlug** | **string** | The repository slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersWithoutPermission1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | If specified only user names containing the supplied string will be returned. | 
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


## RemoveGroupFromUser

> RemoveGroupFromUser(ctx).GroupPickerContext(groupPickerContext).Execute()

Remove user from group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	groupPickerContext := *openapiclient.NewGroupPickerContext() // GroupPickerContext |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.RemoveGroupFromUser(context.Background()).GroupPickerContext(groupPickerContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.RemoveGroupFromUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupFromUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupPickerContext** | [**GroupPickerContext**](GroupPickerContext.md) |  | 

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


## RemoveUserFromGroup

> RemoveUserFromGroup(ctx).UserPickerContext(userPickerContext).Execute()

Remove user from group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userPickerContext := *openapiclient.NewUserPickerContext() // UserPickerContext |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.RemoveUserFromGroup(context.Background()).UserPickerContext(userPickerContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.RemoveUserFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userPickerContext** | [**UserPickerContext**](UserPickerContext.md) |  | 

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


## RenameUser

> RestDetailedUser RenameUser(ctx).UserRename(userRename).Execute()

Rename user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	userRename := *openapiclient.NewUserRename() // UserRename |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.RenameUser(context.Background()).UserRename(userRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.RenameUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameUser`: RestDetailedUser
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.RenameUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRename** | [**UserRename**](UserRename.md) |  | 

### Return type

[**RestDetailedUser**](RestDetailedUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePermissions1

> RevokePermissions1(ctx, projectKey, repositorySlug).User(user).Group(group).Execute()

Revoke all repository permissions for users and groups



### Example

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
	user := "user_example" // string | The names of the users (optional)
	group := "group_example" // string | The names of the groups (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.RevokePermissions1(context.Background(), projectKey, repositorySlug).User(user).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.RevokePermissions1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRevokePermissions1Request struct via the builder pattern


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


## RevokePermissionsForGroup

> RevokePermissionsForGroup(ctx).Name(name).Execute()

Revoke all global permissions for group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := "name_example" // string | The name of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.RevokePermissionsForGroup(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.RevokePermissionsForGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionsForGroupRequest struct via the builder pattern


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


## RevokePermissionsForGroup2

> RevokePermissionsForGroup2(ctx, projectKey, repositorySlug).Name(name).Execute()

Revoke group repository permission



### Example

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
	name := "name_example" // string | The name of the group.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.RevokePermissionsForGroup2(context.Background(), projectKey, repositorySlug).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.RevokePermissionsForGroup2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRevokePermissionsForGroup2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the group. | 


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


## RevokePermissionsForUser

> RevokePermissionsForUser(ctx).Name(name).Execute()

Revoke all global permissions for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := "name_example" // string | The name of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.RevokePermissionsForUser(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.RevokePermissionsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionsForUserRequest struct via the builder pattern


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


## RevokePermissionsForUser2

> RevokePermissionsForUser2(ctx, projectKey, repositorySlug).Name(name).Execute()

Revoke user repository permission



### Example

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
	name := "name_example" // string | The name of the user.
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.RevokePermissionsForUser2(context.Background(), projectKey, repositorySlug).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.RevokePermissionsForUser2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRevokePermissionsForUser2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | The name of the user. | 


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


## SearchPermissions1

> SearchPermissions1(ctx, projectKey, repositorySlug).Permission(permission).FilterText(filterText).Type_(type_).Execute()

Search repository permissions



### Example

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
	permission := "permission_example" // string | Permissions to filter by. See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+repository+permissions)for a detailed explanation of what each permission entails. This parameter can be specified multiple times to filter by more than one permission, and can contain repository, project, and global permissions.   (optional)
	filterText := "filterText_example" // string | Name of the user or group to filter the name of (optional)
	type_ := "type__example" // string | Type of entity (user or group)Valid entity types are:  - USER- GROUP (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.SearchPermissions1(context.Background(), projectKey, repositorySlug).Permission(permission).FilterText(filterText).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.SearchPermissions1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSearchPermissions1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **permission** | **string** | Permissions to filter by. See the [permissions documentation](https://confluence.atlassian.com/display/BitbucketServer/Using+repository+permissions)for a detailed explanation of what each permission entails. This parameter can be specified multiple times to filter by more than one permission, and can contain repository, project, and global permissions.   | 
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


## SetPermissionForGroup

> SetPermissionForGroup(ctx, projectKey, repositorySlug).Name(name).Permission(permission).Execute()

Update group repository permission



### Example

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
	name := []string{"Inner_example"} // []string | The names of the groups.
	permission := "permission_example" // string | The permission to grant
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.SetPermissionForGroup(context.Background(), projectKey, repositorySlug).Name(name).Permission(permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.SetPermissionForGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetPermissionForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **[]string** | The names of the groups. | 
 **permission** | **string** | The permission to grant | 


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


## SetPermissionForGroups

> SetPermissionForGroups(ctx).Name(name).Permission(permission).Execute()

Update global permission for group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := []string{"Inner_example"} // []string | The names of the groups
	permission := "permission_example" // string | The permission to grant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.SetPermissionForGroups(context.Background()).Name(name).Permission(permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.SetPermissionForGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPermissionForGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** | The names of the groups | 
 **permission** | **string** | The permission to grant | 

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


## SetPermissionForUser

> SetPermissionForUser(ctx, projectKey, repositorySlug).Name(name).Permission(permission).Execute()

Update user repository permission



### Example

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
	name := []string{"Inner_example"} // []string | The names of the users.
	permission := "permission_example" // string | The permission to grant
	repositorySlug := "repositorySlug_example" // string | The repository slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.SetPermissionForUser(context.Background(), projectKey, repositorySlug).Name(name).Permission(permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.SetPermissionForUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetPermissionForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **[]string** | The names of the users. | 
 **permission** | **string** | The permission to grant | 


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


## SetPermissionForUsers

> SetPermissionForUsers(ctx).Name(name).Permission(permission).Execute()

Update global permission for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := []string{"Inner_example"} // []string | The names of the users
	permission := "permission_example" // string | The permission to grant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.SetPermissionForUsers(context.Background()).Name(name).Permission(permission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.SetPermissionForUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPermissionForUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** | The names of the users | 
 **permission** | **string** | The permission to grant | 

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


## UpdateUserDetails

> RestDetailedUser UpdateUserDetails(ctx).UserUpdate(userUpdate).Execute()

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
	userUpdate := *openapiclient.NewUserUpdate() // UserUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionManagementAPI.UpdateUserDetails(context.Background()).UserUpdate(userUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.UpdateUserDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserDetails`: RestDetailedUser
	fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.UpdateUserDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userUpdate** | [**UserUpdate**](UserUpdate.md) |  | 

### Return type

[**RestDetailedUser**](RestDetailedUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPassword

> UpdateUserPassword(ctx).AdminPasswordUpdate(adminPasswordUpdate).Execute()

Set password for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	adminPasswordUpdate := *openapiclient.NewAdminPasswordUpdate() // AdminPasswordUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.UpdateUserPassword(context.Background()).AdminPasswordUpdate(adminPasswordUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.UpdateUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminPasswordUpdate** | [**AdminPasswordUpdate**](AdminPasswordUpdate.md) |  | 

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


## ValidateErasable

> ValidateErasable(ctx).Name(name).Execute()

Check user removal



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func main() {
	name := "name_example" // string | The username of the user to validate erasability for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionManagementAPI.ValidateErasable(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.ValidateErasable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateErasableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The username of the user to validate erasability for. | 

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


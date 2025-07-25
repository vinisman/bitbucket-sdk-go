/*
Bitbucket Data Center

This is the reference document for the Atlassian Bitbucket REST API. The REST API is for developers who want to:    - integrate Bitbucket with other applications;   - create scripts that interact with Bitbucket; or   - develop plugins that enhance the Bitbucket UI, using REST to interact with the backend.    You can read more about developing Bitbucket plugins in the [Bitbucket Developer Documentation](https://developer.atlassian.com/bitbucket/server/docs/latest/).

API version: 9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RestPullRequestActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestPullRequestActivity{}

// RestPullRequestActivity struct for RestPullRequestActivity
type RestPullRequestActivity struct {
	Action *string `json:"action,omitempty"`
	CreatedDate *int64 `json:"createdDate,omitempty"`
	Id *int64 `json:"id,omitempty"`
	User *RestCommentAnchorPullRequestAuthorUser `json:"user,omitempty"`
}

// NewRestPullRequestActivity instantiates a new RestPullRequestActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestPullRequestActivity() *RestPullRequestActivity {
	this := RestPullRequestActivity{}
	return &this
}

// NewRestPullRequestActivityWithDefaults instantiates a new RestPullRequestActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestPullRequestActivityWithDefaults() *RestPullRequestActivity {
	this := RestPullRequestActivity{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RestPullRequestActivity) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPullRequestActivity) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RestPullRequestActivity) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RestPullRequestActivity) SetAction(v string) {
	o.Action = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *RestPullRequestActivity) GetCreatedDate() int64 {
	if o == nil || IsNil(o.CreatedDate) {
		var ret int64
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPullRequestActivity) GetCreatedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *RestPullRequestActivity) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given int64 and assigns it to the CreatedDate field.
func (o *RestPullRequestActivity) SetCreatedDate(v int64) {
	o.CreatedDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RestPullRequestActivity) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPullRequestActivity) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RestPullRequestActivity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RestPullRequestActivity) SetId(v int64) {
	o.Id = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RestPullRequestActivity) GetUser() RestCommentAnchorPullRequestAuthorUser {
	if o == nil || IsNil(o.User) {
		var ret RestCommentAnchorPullRequestAuthorUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPullRequestActivity) GetUserOk() (*RestCommentAnchorPullRequestAuthorUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RestPullRequestActivity) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given RestCommentAnchorPullRequestAuthorUser and assigns it to the User field.
func (o *RestPullRequestActivity) SetUser(v RestCommentAnchorPullRequestAuthorUser) {
	o.User = &v
}

func (o RestPullRequestActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestPullRequestActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableRestPullRequestActivity struct {
	value *RestPullRequestActivity
	isSet bool
}

func (v NullableRestPullRequestActivity) Get() *RestPullRequestActivity {
	return v.value
}

func (v *NullableRestPullRequestActivity) Set(val *RestPullRequestActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableRestPullRequestActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableRestPullRequestActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestPullRequestActivity(val *RestPullRequestActivity) *NullableRestPullRequestActivity {
	return &NullableRestPullRequestActivity{value: val, isSet: true}
}

func (v NullableRestPullRequestActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestPullRequestActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



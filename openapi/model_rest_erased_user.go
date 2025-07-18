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

// checks if the RestErasedUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestErasedUser{}

// RestErasedUser struct for RestErasedUser
type RestErasedUser struct {
	NewIdentifier *string `json:"newIdentifier,omitempty"`
}

// NewRestErasedUser instantiates a new RestErasedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestErasedUser() *RestErasedUser {
	this := RestErasedUser{}
	return &this
}

// NewRestErasedUserWithDefaults instantiates a new RestErasedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestErasedUserWithDefaults() *RestErasedUser {
	this := RestErasedUser{}
	return &this
}

// GetNewIdentifier returns the NewIdentifier field value if set, zero value otherwise.
func (o *RestErasedUser) GetNewIdentifier() string {
	if o == nil || IsNil(o.NewIdentifier) {
		var ret string
		return ret
	}
	return *o.NewIdentifier
}

// GetNewIdentifierOk returns a tuple with the NewIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestErasedUser) GetNewIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.NewIdentifier) {
		return nil, false
	}
	return o.NewIdentifier, true
}

// HasNewIdentifier returns a boolean if a field has been set.
func (o *RestErasedUser) HasNewIdentifier() bool {
	if o != nil && !IsNil(o.NewIdentifier) {
		return true
	}

	return false
}

// SetNewIdentifier gets a reference to the given string and assigns it to the NewIdentifier field.
func (o *RestErasedUser) SetNewIdentifier(v string) {
	o.NewIdentifier = &v
}

func (o RestErasedUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestErasedUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewIdentifier) {
		toSerialize["newIdentifier"] = o.NewIdentifier
	}
	return toSerialize, nil
}

type NullableRestErasedUser struct {
	value *RestErasedUser
	isSet bool
}

func (v NullableRestErasedUser) Get() *RestErasedUser {
	return v.value
}

func (v *NullableRestErasedUser) Set(val *RestErasedUser) {
	v.value = val
	v.isSet = true
}

func (v NullableRestErasedUser) IsSet() bool {
	return v.isSet
}

func (v *NullableRestErasedUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestErasedUser(val *RestErasedUser) *NullableRestErasedUser {
	return &NullableRestErasedUser{value: val, isSet: true}
}

func (v NullableRestErasedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestErasedUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Bitbucket Data Center

This is the reference document for the Atlassian Bitbucket REST API. The REST API is for developers who want to:    - integrate Bitbucket with other applications;   - create scripts that interact with Bitbucket; or   - develop plugins that enhance the Bitbucket UI, using REST to interact with the backend.    You can read more about developing Bitbucket plugins in the [Bitbucket Developer Documentation](https://developer.atlassian.com/bitbucket/server/docs/latest/).

API version: 9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GroupAndUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupAndUsers{}

// GroupAndUsers struct for GroupAndUsers
type GroupAndUsers struct {
	Group *string `json:"group,omitempty"`
	Users []string `json:"users"`
}

type _GroupAndUsers GroupAndUsers

// NewGroupAndUsers instantiates a new GroupAndUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupAndUsers(users []string) *GroupAndUsers {
	this := GroupAndUsers{}
	this.Users = users
	return &this
}

// NewGroupAndUsersWithDefaults instantiates a new GroupAndUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupAndUsersWithDefaults() *GroupAndUsers {
	this := GroupAndUsers{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *GroupAndUsers) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupAndUsers) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *GroupAndUsers) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *GroupAndUsers) SetGroup(v string) {
	o.Group = &v
}

// GetUsers returns the Users field value
func (o *GroupAndUsers) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *GroupAndUsers) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *GroupAndUsers) SetUsers(v []string) {
	o.Users = v
}

func (o GroupAndUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupAndUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *GroupAndUsers) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGroupAndUsers := _GroupAndUsers{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGroupAndUsers)

	if err != nil {
		return err
	}

	*o = GroupAndUsers(varGroupAndUsers)

	return err
}

type NullableGroupAndUsers struct {
	value *GroupAndUsers
	isSet bool
}

func (v NullableGroupAndUsers) Get() *GroupAndUsers {
	return v.value
}

func (v *NullableGroupAndUsers) Set(val *GroupAndUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupAndUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupAndUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupAndUsers(val *GroupAndUsers) *NullableGroupAndUsers {
	return &NullableGroupAndUsers{value: val, isSet: true}
}

func (v NullableGroupAndUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupAndUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



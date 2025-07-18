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

// checks if the RestDefaultBranch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestDefaultBranch{}

// RestDefaultBranch struct for RestDefaultBranch
type RestDefaultBranch struct {
	Id *string `json:"id,omitempty"`
}

// NewRestDefaultBranch instantiates a new RestDefaultBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestDefaultBranch() *RestDefaultBranch {
	this := RestDefaultBranch{}
	return &this
}

// NewRestDefaultBranchWithDefaults instantiates a new RestDefaultBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestDefaultBranchWithDefaults() *RestDefaultBranch {
	this := RestDefaultBranch{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RestDefaultBranch) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestDefaultBranch) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RestDefaultBranch) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RestDefaultBranch) SetId(v string) {
	o.Id = &v
}

func (o RestDefaultBranch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestDefaultBranch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableRestDefaultBranch struct {
	value *RestDefaultBranch
	isSet bool
}

func (v NullableRestDefaultBranch) Get() *RestDefaultBranch {
	return v.value
}

func (v *NullableRestDefaultBranch) Set(val *RestDefaultBranch) {
	v.value = val
	v.isSet = true
}

func (v NullableRestDefaultBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableRestDefaultBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestDefaultBranch(val *RestDefaultBranch) *NullableRestDefaultBranch {
	return &NullableRestDefaultBranch{value: val, isSet: true}
}

func (v NullableRestDefaultBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestDefaultBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



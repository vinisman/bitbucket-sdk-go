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

// checks if the RestAccessTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestAccessTokenRequest{}

// RestAccessTokenRequest struct for RestAccessTokenRequest
type RestAccessTokenRequest struct {
	ExpiryDays *int32 `json:"expiryDays,omitempty"`
	Name *string `json:"name,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// NewRestAccessTokenRequest instantiates a new RestAccessTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestAccessTokenRequest() *RestAccessTokenRequest {
	this := RestAccessTokenRequest{}
	return &this
}

// NewRestAccessTokenRequestWithDefaults instantiates a new RestAccessTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestAccessTokenRequestWithDefaults() *RestAccessTokenRequest {
	this := RestAccessTokenRequest{}
	return &this
}

// GetExpiryDays returns the ExpiryDays field value if set, zero value otherwise.
func (o *RestAccessTokenRequest) GetExpiryDays() int32 {
	if o == nil || IsNil(o.ExpiryDays) {
		var ret int32
		return ret
	}
	return *o.ExpiryDays
}

// GetExpiryDaysOk returns a tuple with the ExpiryDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestAccessTokenRequest) GetExpiryDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiryDays) {
		return nil, false
	}
	return o.ExpiryDays, true
}

// HasExpiryDays returns a boolean if a field has been set.
func (o *RestAccessTokenRequest) HasExpiryDays() bool {
	if o != nil && !IsNil(o.ExpiryDays) {
		return true
	}

	return false
}

// SetExpiryDays gets a reference to the given int32 and assigns it to the ExpiryDays field.
func (o *RestAccessTokenRequest) SetExpiryDays(v int32) {
	o.ExpiryDays = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RestAccessTokenRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestAccessTokenRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RestAccessTokenRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RestAccessTokenRequest) SetName(v string) {
	o.Name = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RestAccessTokenRequest) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestAccessTokenRequest) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RestAccessTokenRequest) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *RestAccessTokenRequest) SetPermissions(v []string) {
	o.Permissions = v
}

func (o RestAccessTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestAccessTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiryDays) {
		toSerialize["expiryDays"] = o.ExpiryDays
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableRestAccessTokenRequest struct {
	value *RestAccessTokenRequest
	isSet bool
}

func (v NullableRestAccessTokenRequest) Get() *RestAccessTokenRequest {
	return v.value
}

func (v *NullableRestAccessTokenRequest) Set(val *RestAccessTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestAccessTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestAccessTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestAccessTokenRequest(val *RestAccessTokenRequest) *NullableRestAccessTokenRequest {
	return &NullableRestAccessTokenRequest{value: val, isSet: true}
}

func (v NullableRestAccessTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestAccessTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



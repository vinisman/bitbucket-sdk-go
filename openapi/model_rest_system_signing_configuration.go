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

// checks if the RestSystemSigningConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestSystemSigningConfiguration{}

// RestSystemSigningConfiguration struct for RestSystemSigningConfiguration
type RestSystemSigningConfiguration struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewRestSystemSigningConfiguration instantiates a new RestSystemSigningConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestSystemSigningConfiguration() *RestSystemSigningConfiguration {
	this := RestSystemSigningConfiguration{}
	return &this
}

// NewRestSystemSigningConfigurationWithDefaults instantiates a new RestSystemSigningConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestSystemSigningConfigurationWithDefaults() *RestSystemSigningConfiguration {
	this := RestSystemSigningConfiguration{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RestSystemSigningConfiguration) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestSystemSigningConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RestSystemSigningConfiguration) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RestSystemSigningConfiguration) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o RestSystemSigningConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestSystemSigningConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableRestSystemSigningConfiguration struct {
	value *RestSystemSigningConfiguration
	isSet bool
}

func (v NullableRestSystemSigningConfiguration) Get() *RestSystemSigningConfiguration {
	return v.value
}

func (v *NullableRestSystemSigningConfiguration) Set(val *RestSystemSigningConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableRestSystemSigningConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableRestSystemSigningConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestSystemSigningConfiguration(val *RestSystemSigningConfiguration) *NullableRestSystemSigningConfiguration {
	return &NullableRestSystemSigningConfiguration{value: val, isSet: true}
}

func (v NullableRestSystemSigningConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestSystemSigningConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



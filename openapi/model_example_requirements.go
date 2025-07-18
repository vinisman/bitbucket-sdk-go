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

// checks if the ExampleRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExampleRequirements{}

// ExampleRequirements struct for ExampleRequirements
type ExampleRequirements struct {
	Count *string `json:"count,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewExampleRequirements instantiates a new ExampleRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExampleRequirements() *ExampleRequirements {
	this := ExampleRequirements{}
	return &this
}

// NewExampleRequirementsWithDefaults instantiates a new ExampleRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExampleRequirementsWithDefaults() *ExampleRequirements {
	this := ExampleRequirements{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ExampleRequirements) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleRequirements) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ExampleRequirements) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *ExampleRequirements) SetCount(v string) {
	o.Count = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ExampleRequirements) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleRequirements) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ExampleRequirements) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ExampleRequirements) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ExampleRequirements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExampleRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableExampleRequirements struct {
	value *ExampleRequirements
	isSet bool
}

func (v NullableExampleRequirements) Get() *ExampleRequirements {
	return v.value
}

func (v *NullableExampleRequirements) Set(val *ExampleRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableExampleRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableExampleRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExampleRequirements(val *ExampleRequirements) *NullableExampleRequirements {
	return &NullableExampleRequirements{value: val, isSet: true}
}

func (v NullableExampleRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExampleRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



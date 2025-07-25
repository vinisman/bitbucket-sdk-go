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

// checks if the RestRepositoryPullRequestSettingsRequiredApprovers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestRepositoryPullRequestSettingsRequiredApprovers{}

// RestRepositoryPullRequestSettingsRequiredApprovers struct for RestRepositoryPullRequestSettingsRequiredApprovers
type RestRepositoryPullRequestSettingsRequiredApprovers struct {
	Count *string `json:"count,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewRestRepositoryPullRequestSettingsRequiredApprovers instantiates a new RestRepositoryPullRequestSettingsRequiredApprovers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestRepositoryPullRequestSettingsRequiredApprovers() *RestRepositoryPullRequestSettingsRequiredApprovers {
	this := RestRepositoryPullRequestSettingsRequiredApprovers{}
	return &this
}

// NewRestRepositoryPullRequestSettingsRequiredApproversWithDefaults instantiates a new RestRepositoryPullRequestSettingsRequiredApprovers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestRepositoryPullRequestSettingsRequiredApproversWithDefaults() *RestRepositoryPullRequestSettingsRequiredApprovers {
	this := RestRepositoryPullRequestSettingsRequiredApprovers{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RestRepositoryPullRequestSettingsRequiredApprovers) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestRepositoryPullRequestSettingsRequiredApprovers) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RestRepositoryPullRequestSettingsRequiredApprovers) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *RestRepositoryPullRequestSettingsRequiredApprovers) SetCount(v string) {
	o.Count = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RestRepositoryPullRequestSettingsRequiredApprovers) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestRepositoryPullRequestSettingsRequiredApprovers) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RestRepositoryPullRequestSettingsRequiredApprovers) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RestRepositoryPullRequestSettingsRequiredApprovers) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o RestRepositoryPullRequestSettingsRequiredApprovers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestRepositoryPullRequestSettingsRequiredApprovers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableRestRepositoryPullRequestSettingsRequiredApprovers struct {
	value *RestRepositoryPullRequestSettingsRequiredApprovers
	isSet bool
}

func (v NullableRestRepositoryPullRequestSettingsRequiredApprovers) Get() *RestRepositoryPullRequestSettingsRequiredApprovers {
	return v.value
}

func (v *NullableRestRepositoryPullRequestSettingsRequiredApprovers) Set(val *RestRepositoryPullRequestSettingsRequiredApprovers) {
	v.value = val
	v.isSet = true
}

func (v NullableRestRepositoryPullRequestSettingsRequiredApprovers) IsSet() bool {
	return v.isSet
}

func (v *NullableRestRepositoryPullRequestSettingsRequiredApprovers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestRepositoryPullRequestSettingsRequiredApprovers(val *RestRepositoryPullRequestSettingsRequiredApprovers) *NullableRestRepositoryPullRequestSettingsRequiredApprovers {
	return &NullableRestRepositoryPullRequestSettingsRequiredApprovers{value: val, isSet: true}
}

func (v NullableRestRepositoryPullRequestSettingsRequiredApprovers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestRepositoryPullRequestSettingsRequiredApprovers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



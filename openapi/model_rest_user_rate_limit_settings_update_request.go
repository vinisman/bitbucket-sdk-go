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

// checks if the RestUserRateLimitSettingsUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestUserRateLimitSettingsUpdateRequest{}

// RestUserRateLimitSettingsUpdateRequest struct for RestUserRateLimitSettingsUpdateRequest
type RestUserRateLimitSettingsUpdateRequest struct {
	Settings *RestBulkUserRateLimitSettingsUpdateRequestSettings `json:"settings,omitempty"`
	Whitelisted *bool `json:"whitelisted,omitempty"`
}

// NewRestUserRateLimitSettingsUpdateRequest instantiates a new RestUserRateLimitSettingsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestUserRateLimitSettingsUpdateRequest() *RestUserRateLimitSettingsUpdateRequest {
	this := RestUserRateLimitSettingsUpdateRequest{}
	return &this
}

// NewRestUserRateLimitSettingsUpdateRequestWithDefaults instantiates a new RestUserRateLimitSettingsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestUserRateLimitSettingsUpdateRequestWithDefaults() *RestUserRateLimitSettingsUpdateRequest {
	this := RestUserRateLimitSettingsUpdateRequest{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *RestUserRateLimitSettingsUpdateRequest) GetSettings() RestBulkUserRateLimitSettingsUpdateRequestSettings {
	if o == nil || IsNil(o.Settings) {
		var ret RestBulkUserRateLimitSettingsUpdateRequestSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestUserRateLimitSettingsUpdateRequest) GetSettingsOk() (*RestBulkUserRateLimitSettingsUpdateRequestSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *RestUserRateLimitSettingsUpdateRequest) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given RestBulkUserRateLimitSettingsUpdateRequestSettings and assigns it to the Settings field.
func (o *RestUserRateLimitSettingsUpdateRequest) SetSettings(v RestBulkUserRateLimitSettingsUpdateRequestSettings) {
	o.Settings = &v
}

// GetWhitelisted returns the Whitelisted field value if set, zero value otherwise.
func (o *RestUserRateLimitSettingsUpdateRequest) GetWhitelisted() bool {
	if o == nil || IsNil(o.Whitelisted) {
		var ret bool
		return ret
	}
	return *o.Whitelisted
}

// GetWhitelistedOk returns a tuple with the Whitelisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestUserRateLimitSettingsUpdateRequest) GetWhitelistedOk() (*bool, bool) {
	if o == nil || IsNil(o.Whitelisted) {
		return nil, false
	}
	return o.Whitelisted, true
}

// HasWhitelisted returns a boolean if a field has been set.
func (o *RestUserRateLimitSettingsUpdateRequest) HasWhitelisted() bool {
	if o != nil && !IsNil(o.Whitelisted) {
		return true
	}

	return false
}

// SetWhitelisted gets a reference to the given bool and assigns it to the Whitelisted field.
func (o *RestUserRateLimitSettingsUpdateRequest) SetWhitelisted(v bool) {
	o.Whitelisted = &v
}

func (o RestUserRateLimitSettingsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestUserRateLimitSettingsUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Whitelisted) {
		toSerialize["whitelisted"] = o.Whitelisted
	}
	return toSerialize, nil
}

type NullableRestUserRateLimitSettingsUpdateRequest struct {
	value *RestUserRateLimitSettingsUpdateRequest
	isSet bool
}

func (v NullableRestUserRateLimitSettingsUpdateRequest) Get() *RestUserRateLimitSettingsUpdateRequest {
	return v.value
}

func (v *NullableRestUserRateLimitSettingsUpdateRequest) Set(val *RestUserRateLimitSettingsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestUserRateLimitSettingsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestUserRateLimitSettingsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestUserRateLimitSettingsUpdateRequest(val *RestUserRateLimitSettingsUpdateRequest) *NullableRestUserRateLimitSettingsUpdateRequest {
	return &NullableRestUserRateLimitSettingsUpdateRequest{value: val, isSet: true}
}

func (v NullableRestUserRateLimitSettingsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestUserRateLimitSettingsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



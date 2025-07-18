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

// checks if the RestBulkUserRateLimitSettingsUpdateRequestSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestBulkUserRateLimitSettingsUpdateRequestSettings{}

// RestBulkUserRateLimitSettingsUpdateRequestSettings struct for RestBulkUserRateLimitSettingsUpdateRequestSettings
type RestBulkUserRateLimitSettingsUpdateRequestSettings struct {
	Capacity *int32 `json:"capacity,omitempty"`
	FillRate *int32 `json:"fillRate,omitempty"`
}

// NewRestBulkUserRateLimitSettingsUpdateRequestSettings instantiates a new RestBulkUserRateLimitSettingsUpdateRequestSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestBulkUserRateLimitSettingsUpdateRequestSettings() *RestBulkUserRateLimitSettingsUpdateRequestSettings {
	this := RestBulkUserRateLimitSettingsUpdateRequestSettings{}
	return &this
}

// NewRestBulkUserRateLimitSettingsUpdateRequestSettingsWithDefaults instantiates a new RestBulkUserRateLimitSettingsUpdateRequestSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestBulkUserRateLimitSettingsUpdateRequestSettingsWithDefaults() *RestBulkUserRateLimitSettingsUpdateRequestSettings {
	this := RestBulkUserRateLimitSettingsUpdateRequestSettings{}
	return &this
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *RestBulkUserRateLimitSettingsUpdateRequestSettings) GetCapacity() int32 {
	if o == nil || IsNil(o.Capacity) {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBulkUserRateLimitSettingsUpdateRequestSettings) GetCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *RestBulkUserRateLimitSettingsUpdateRequestSettings) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *RestBulkUserRateLimitSettingsUpdateRequestSettings) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetFillRate returns the FillRate field value if set, zero value otherwise.
func (o *RestBulkUserRateLimitSettingsUpdateRequestSettings) GetFillRate() int32 {
	if o == nil || IsNil(o.FillRate) {
		var ret int32
		return ret
	}
	return *o.FillRate
}

// GetFillRateOk returns a tuple with the FillRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBulkUserRateLimitSettingsUpdateRequestSettings) GetFillRateOk() (*int32, bool) {
	if o == nil || IsNil(o.FillRate) {
		return nil, false
	}
	return o.FillRate, true
}

// HasFillRate returns a boolean if a field has been set.
func (o *RestBulkUserRateLimitSettingsUpdateRequestSettings) HasFillRate() bool {
	if o != nil && !IsNil(o.FillRate) {
		return true
	}

	return false
}

// SetFillRate gets a reference to the given int32 and assigns it to the FillRate field.
func (o *RestBulkUserRateLimitSettingsUpdateRequestSettings) SetFillRate(v int32) {
	o.FillRate = &v
}

func (o RestBulkUserRateLimitSettingsUpdateRequestSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestBulkUserRateLimitSettingsUpdateRequestSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !IsNil(o.FillRate) {
		toSerialize["fillRate"] = o.FillRate
	}
	return toSerialize, nil
}

type NullableRestBulkUserRateLimitSettingsUpdateRequestSettings struct {
	value *RestBulkUserRateLimitSettingsUpdateRequestSettings
	isSet bool
}

func (v NullableRestBulkUserRateLimitSettingsUpdateRequestSettings) Get() *RestBulkUserRateLimitSettingsUpdateRequestSettings {
	return v.value
}

func (v *NullableRestBulkUserRateLimitSettingsUpdateRequestSettings) Set(val *RestBulkUserRateLimitSettingsUpdateRequestSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRestBulkUserRateLimitSettingsUpdateRequestSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRestBulkUserRateLimitSettingsUpdateRequestSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestBulkUserRateLimitSettingsUpdateRequestSettings(val *RestBulkUserRateLimitSettingsUpdateRequestSettings) *NullableRestBulkUserRateLimitSettingsUpdateRequestSettings {
	return &NullableRestBulkUserRateLimitSettingsUpdateRequestSettings{value: val, isSet: true}
}

func (v NullableRestBulkUserRateLimitSettingsUpdateRequestSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestBulkUserRateLimitSettingsUpdateRequestSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



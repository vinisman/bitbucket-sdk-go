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

// checks if the RestChangeConflict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestChangeConflict{}

// RestChangeConflict struct for RestChangeConflict
type RestChangeConflict struct {
	OurChange *RestChangeConflictOurChange `json:"ourChange,omitempty"`
	TheirChange *RestChangeConflictOurChange `json:"theirChange,omitempty"`
}

// NewRestChangeConflict instantiates a new RestChangeConflict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestChangeConflict() *RestChangeConflict {
	this := RestChangeConflict{}
	return &this
}

// NewRestChangeConflictWithDefaults instantiates a new RestChangeConflict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestChangeConflictWithDefaults() *RestChangeConflict {
	this := RestChangeConflict{}
	return &this
}

// GetOurChange returns the OurChange field value if set, zero value otherwise.
func (o *RestChangeConflict) GetOurChange() RestChangeConflictOurChange {
	if o == nil || IsNil(o.OurChange) {
		var ret RestChangeConflictOurChange
		return ret
	}
	return *o.OurChange
}

// GetOurChangeOk returns a tuple with the OurChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangeConflict) GetOurChangeOk() (*RestChangeConflictOurChange, bool) {
	if o == nil || IsNil(o.OurChange) {
		return nil, false
	}
	return o.OurChange, true
}

// HasOurChange returns a boolean if a field has been set.
func (o *RestChangeConflict) HasOurChange() bool {
	if o != nil && !IsNil(o.OurChange) {
		return true
	}

	return false
}

// SetOurChange gets a reference to the given RestChangeConflictOurChange and assigns it to the OurChange field.
func (o *RestChangeConflict) SetOurChange(v RestChangeConflictOurChange) {
	o.OurChange = &v
}

// GetTheirChange returns the TheirChange field value if set, zero value otherwise.
func (o *RestChangeConflict) GetTheirChange() RestChangeConflictOurChange {
	if o == nil || IsNil(o.TheirChange) {
		var ret RestChangeConflictOurChange
		return ret
	}
	return *o.TheirChange
}

// GetTheirChangeOk returns a tuple with the TheirChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangeConflict) GetTheirChangeOk() (*RestChangeConflictOurChange, bool) {
	if o == nil || IsNil(o.TheirChange) {
		return nil, false
	}
	return o.TheirChange, true
}

// HasTheirChange returns a boolean if a field has been set.
func (o *RestChangeConflict) HasTheirChange() bool {
	if o != nil && !IsNil(o.TheirChange) {
		return true
	}

	return false
}

// SetTheirChange gets a reference to the given RestChangeConflictOurChange and assigns it to the TheirChange field.
func (o *RestChangeConflict) SetTheirChange(v RestChangeConflictOurChange) {
	o.TheirChange = &v
}

func (o RestChangeConflict) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestChangeConflict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OurChange) {
		toSerialize["ourChange"] = o.OurChange
	}
	if !IsNil(o.TheirChange) {
		toSerialize["theirChange"] = o.TheirChange
	}
	return toSerialize, nil
}

type NullableRestChangeConflict struct {
	value *RestChangeConflict
	isSet bool
}

func (v NullableRestChangeConflict) Get() *RestChangeConflict {
	return v.value
}

func (v *NullableRestChangeConflict) Set(val *RestChangeConflict) {
	v.value = val
	v.isSet = true
}

func (v NullableRestChangeConflict) IsSet() bool {
	return v.isSet
}

func (v *NullableRestChangeConflict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestChangeConflict(val *RestChangeConflict) *NullableRestChangeConflict {
	return &NullableRestChangeConflict{value: val, isSet: true}
}

func (v NullableRestChangeConflict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestChangeConflict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the RestSyncProgress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestSyncProgress{}

// RestSyncProgress struct for RestSyncProgress
type RestSyncProgress struct {
	Discovering *bool `json:"discovering,omitempty"`
	SyncedRepos *int32 `json:"syncedRepos,omitempty"`
	TotalRepos *int32 `json:"totalRepos,omitempty"`
}

// NewRestSyncProgress instantiates a new RestSyncProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestSyncProgress() *RestSyncProgress {
	this := RestSyncProgress{}
	return &this
}

// NewRestSyncProgressWithDefaults instantiates a new RestSyncProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestSyncProgressWithDefaults() *RestSyncProgress {
	this := RestSyncProgress{}
	return &this
}

// GetDiscovering returns the Discovering field value if set, zero value otherwise.
func (o *RestSyncProgress) GetDiscovering() bool {
	if o == nil || IsNil(o.Discovering) {
		var ret bool
		return ret
	}
	return *o.Discovering
}

// GetDiscoveringOk returns a tuple with the Discovering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestSyncProgress) GetDiscoveringOk() (*bool, bool) {
	if o == nil || IsNil(o.Discovering) {
		return nil, false
	}
	return o.Discovering, true
}

// HasDiscovering returns a boolean if a field has been set.
func (o *RestSyncProgress) HasDiscovering() bool {
	if o != nil && !IsNil(o.Discovering) {
		return true
	}

	return false
}

// SetDiscovering gets a reference to the given bool and assigns it to the Discovering field.
func (o *RestSyncProgress) SetDiscovering(v bool) {
	o.Discovering = &v
}

// GetSyncedRepos returns the SyncedRepos field value if set, zero value otherwise.
func (o *RestSyncProgress) GetSyncedRepos() int32 {
	if o == nil || IsNil(o.SyncedRepos) {
		var ret int32
		return ret
	}
	return *o.SyncedRepos
}

// GetSyncedReposOk returns a tuple with the SyncedRepos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestSyncProgress) GetSyncedReposOk() (*int32, bool) {
	if o == nil || IsNil(o.SyncedRepos) {
		return nil, false
	}
	return o.SyncedRepos, true
}

// HasSyncedRepos returns a boolean if a field has been set.
func (o *RestSyncProgress) HasSyncedRepos() bool {
	if o != nil && !IsNil(o.SyncedRepos) {
		return true
	}

	return false
}

// SetSyncedRepos gets a reference to the given int32 and assigns it to the SyncedRepos field.
func (o *RestSyncProgress) SetSyncedRepos(v int32) {
	o.SyncedRepos = &v
}

// GetTotalRepos returns the TotalRepos field value if set, zero value otherwise.
func (o *RestSyncProgress) GetTotalRepos() int32 {
	if o == nil || IsNil(o.TotalRepos) {
		var ret int32
		return ret
	}
	return *o.TotalRepos
}

// GetTotalReposOk returns a tuple with the TotalRepos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestSyncProgress) GetTotalReposOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRepos) {
		return nil, false
	}
	return o.TotalRepos, true
}

// HasTotalRepos returns a boolean if a field has been set.
func (o *RestSyncProgress) HasTotalRepos() bool {
	if o != nil && !IsNil(o.TotalRepos) {
		return true
	}

	return false
}

// SetTotalRepos gets a reference to the given int32 and assigns it to the TotalRepos field.
func (o *RestSyncProgress) SetTotalRepos(v int32) {
	o.TotalRepos = &v
}

func (o RestSyncProgress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestSyncProgress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discovering) {
		toSerialize["discovering"] = o.Discovering
	}
	if !IsNil(o.SyncedRepos) {
		toSerialize["syncedRepos"] = o.SyncedRepos
	}
	if !IsNil(o.TotalRepos) {
		toSerialize["totalRepos"] = o.TotalRepos
	}
	return toSerialize, nil
}

type NullableRestSyncProgress struct {
	value *RestSyncProgress
	isSet bool
}

func (v NullableRestSyncProgress) Get() *RestSyncProgress {
	return v.value
}

func (v *NullableRestSyncProgress) Set(val *RestSyncProgress) {
	v.value = val
	v.isSet = true
}

func (v NullableRestSyncProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableRestSyncProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestSyncProgress(val *RestSyncProgress) *NullableRestSyncProgress {
	return &NullableRestSyncProgress{value: val, isSet: true}
}

func (v NullableRestSyncProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestSyncProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



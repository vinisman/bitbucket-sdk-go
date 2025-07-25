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

// checks if the RestPullRequestDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestPullRequestDeleteRequest{}

// RestPullRequestDeleteRequest struct for RestPullRequestDeleteRequest
type RestPullRequestDeleteRequest struct {
	Version *int32 `json:"version,omitempty"`
}

// NewRestPullRequestDeleteRequest instantiates a new RestPullRequestDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestPullRequestDeleteRequest() *RestPullRequestDeleteRequest {
	this := RestPullRequestDeleteRequest{}
	return &this
}

// NewRestPullRequestDeleteRequestWithDefaults instantiates a new RestPullRequestDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestPullRequestDeleteRequestWithDefaults() *RestPullRequestDeleteRequest {
	this := RestPullRequestDeleteRequest{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RestPullRequestDeleteRequest) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPullRequestDeleteRequest) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RestPullRequestDeleteRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *RestPullRequestDeleteRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o RestPullRequestDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestPullRequestDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableRestPullRequestDeleteRequest struct {
	value *RestPullRequestDeleteRequest
	isSet bool
}

func (v NullableRestPullRequestDeleteRequest) Get() *RestPullRequestDeleteRequest {
	return v.value
}

func (v *NullableRestPullRequestDeleteRequest) Set(val *RestPullRequestDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestPullRequestDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestPullRequestDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestPullRequestDeleteRequest(val *RestPullRequestDeleteRequest) *NullableRestPullRequestDeleteRequest {
	return &NullableRestPullRequestDeleteRequest{value: val, isSet: true}
}

func (v NullableRestPullRequestDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestPullRequestDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



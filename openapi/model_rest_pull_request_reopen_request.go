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

// checks if the RestPullRequestReopenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestPullRequestReopenRequest{}

// RestPullRequestReopenRequest struct for RestPullRequestReopenRequest
type RestPullRequestReopenRequest struct {
	Version *int32 `json:"version,omitempty"`
}

// NewRestPullRequestReopenRequest instantiates a new RestPullRequestReopenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestPullRequestReopenRequest() *RestPullRequestReopenRequest {
	this := RestPullRequestReopenRequest{}
	return &this
}

// NewRestPullRequestReopenRequestWithDefaults instantiates a new RestPullRequestReopenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestPullRequestReopenRequestWithDefaults() *RestPullRequestReopenRequest {
	this := RestPullRequestReopenRequest{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RestPullRequestReopenRequest) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPullRequestReopenRequest) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RestPullRequestReopenRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *RestPullRequestReopenRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o RestPullRequestReopenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestPullRequestReopenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableRestPullRequestReopenRequest struct {
	value *RestPullRequestReopenRequest
	isSet bool
}

func (v NullableRestPullRequestReopenRequest) Get() *RestPullRequestReopenRequest {
	return v.value
}

func (v *NullableRestPullRequestReopenRequest) Set(val *RestPullRequestReopenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestPullRequestReopenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestPullRequestReopenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestPullRequestReopenRequest(val *RestPullRequestReopenRequest) *NullableRestPullRequestReopenRequest {
	return &NullableRestPullRequestReopenRequest{value: val, isSet: true}
}

func (v NullableRestPullRequestReopenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestPullRequestReopenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



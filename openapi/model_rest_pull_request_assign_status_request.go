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

// checks if the RestPullRequestAssignStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestPullRequestAssignStatusRequest{}

// RestPullRequestAssignStatusRequest struct for RestPullRequestAssignStatusRequest
type RestPullRequestAssignStatusRequest struct {
	LastReviewedCommit *string `json:"lastReviewedCommit,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewRestPullRequestAssignStatusRequest instantiates a new RestPullRequestAssignStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestPullRequestAssignStatusRequest() *RestPullRequestAssignStatusRequest {
	this := RestPullRequestAssignStatusRequest{}
	return &this
}

// NewRestPullRequestAssignStatusRequestWithDefaults instantiates a new RestPullRequestAssignStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestPullRequestAssignStatusRequestWithDefaults() *RestPullRequestAssignStatusRequest {
	this := RestPullRequestAssignStatusRequest{}
	return &this
}

// GetLastReviewedCommit returns the LastReviewedCommit field value if set, zero value otherwise.
func (o *RestPullRequestAssignStatusRequest) GetLastReviewedCommit() string {
	if o == nil || IsNil(o.LastReviewedCommit) {
		var ret string
		return ret
	}
	return *o.LastReviewedCommit
}

// GetLastReviewedCommitOk returns a tuple with the LastReviewedCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPullRequestAssignStatusRequest) GetLastReviewedCommitOk() (*string, bool) {
	if o == nil || IsNil(o.LastReviewedCommit) {
		return nil, false
	}
	return o.LastReviewedCommit, true
}

// HasLastReviewedCommit returns a boolean if a field has been set.
func (o *RestPullRequestAssignStatusRequest) HasLastReviewedCommit() bool {
	if o != nil && !IsNil(o.LastReviewedCommit) {
		return true
	}

	return false
}

// SetLastReviewedCommit gets a reference to the given string and assigns it to the LastReviewedCommit field.
func (o *RestPullRequestAssignStatusRequest) SetLastReviewedCommit(v string) {
	o.LastReviewedCommit = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RestPullRequestAssignStatusRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPullRequestAssignStatusRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RestPullRequestAssignStatusRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RestPullRequestAssignStatusRequest) SetStatus(v string) {
	o.Status = &v
}

func (o RestPullRequestAssignStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestPullRequestAssignStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastReviewedCommit) {
		toSerialize["lastReviewedCommit"] = o.LastReviewedCommit
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableRestPullRequestAssignStatusRequest struct {
	value *RestPullRequestAssignStatusRequest
	isSet bool
}

func (v NullableRestPullRequestAssignStatusRequest) Get() *RestPullRequestAssignStatusRequest {
	return v.value
}

func (v *NullableRestPullRequestAssignStatusRequest) Set(val *RestPullRequestAssignStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestPullRequestAssignStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestPullRequestAssignStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestPullRequestAssignStatusRequest(val *RestPullRequestAssignStatusRequest) *NullableRestPullRequestAssignStatusRequest {
	return &NullableRestPullRequestAssignStatusRequest{value: val, isSet: true}
}

func (v NullableRestPullRequestAssignStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestPullRequestAssignStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



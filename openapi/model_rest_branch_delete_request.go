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

// checks if the RestBranchDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestBranchDeleteRequest{}

// RestBranchDeleteRequest struct for RestBranchDeleteRequest
type RestBranchDeleteRequest struct {
	// Don't actually delete the ref name, just do a dry run
	DryRun *bool `json:"dryRun,omitempty"`
	// Commit ID that the provided ref name is expected to point to
	EndPoint *string `json:"endPoint,omitempty"`
	// Name of the ref to be deleted
	Name *string `json:"name,omitempty"`
}

// NewRestBranchDeleteRequest instantiates a new RestBranchDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestBranchDeleteRequest() *RestBranchDeleteRequest {
	this := RestBranchDeleteRequest{}
	return &this
}

// NewRestBranchDeleteRequestWithDefaults instantiates a new RestBranchDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestBranchDeleteRequestWithDefaults() *RestBranchDeleteRequest {
	this := RestBranchDeleteRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *RestBranchDeleteRequest) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBranchDeleteRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *RestBranchDeleteRequest) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *RestBranchDeleteRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetEndPoint returns the EndPoint field value if set, zero value otherwise.
func (o *RestBranchDeleteRequest) GetEndPoint() string {
	if o == nil || IsNil(o.EndPoint) {
		var ret string
		return ret
	}
	return *o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBranchDeleteRequest) GetEndPointOk() (*string, bool) {
	if o == nil || IsNil(o.EndPoint) {
		return nil, false
	}
	return o.EndPoint, true
}

// HasEndPoint returns a boolean if a field has been set.
func (o *RestBranchDeleteRequest) HasEndPoint() bool {
	if o != nil && !IsNil(o.EndPoint) {
		return true
	}

	return false
}

// SetEndPoint gets a reference to the given string and assigns it to the EndPoint field.
func (o *RestBranchDeleteRequest) SetEndPoint(v string) {
	o.EndPoint = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RestBranchDeleteRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBranchDeleteRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RestBranchDeleteRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RestBranchDeleteRequest) SetName(v string) {
	o.Name = &v
}

func (o RestBranchDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestBranchDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DryRun) {
		toSerialize["dryRun"] = o.DryRun
	}
	if !IsNil(o.EndPoint) {
		toSerialize["endPoint"] = o.EndPoint
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableRestBranchDeleteRequest struct {
	value *RestBranchDeleteRequest
	isSet bool
}

func (v NullableRestBranchDeleteRequest) Get() *RestBranchDeleteRequest {
	return v.value
}

func (v *NullableRestBranchDeleteRequest) Set(val *RestBranchDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestBranchDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestBranchDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestBranchDeleteRequest(val *RestBranchDeleteRequest) *NullableRestBranchDeleteRequest {
	return &NullableRestBranchDeleteRequest{value: val, isSet: true}
}

func (v NullableRestBranchDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestBranchDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



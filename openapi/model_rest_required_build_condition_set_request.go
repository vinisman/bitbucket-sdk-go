/*
Bitbucket Data Center

This is the reference document for the Atlassian Bitbucket REST API. The REST API is for developers who want to:    - integrate Bitbucket with other applications;   - create scripts that interact with Bitbucket; or   - develop plugins that enhance the Bitbucket UI, using REST to interact with the backend.    You can read more about developing Bitbucket plugins in the [Bitbucket Developer Documentation](https://developer.atlassian.com/bitbucket/server/docs/latest/).

API version: 9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RestRequiredBuildConditionSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestRequiredBuildConditionSetRequest{}

// RestRequiredBuildConditionSetRequest struct for RestRequiredBuildConditionSetRequest
type RestRequiredBuildConditionSetRequest struct {
	// A non-empty list of build parent keys that require green builds for this merge check to pass
	BuildParentKeys []string `json:"buildParentKeys"`
	ExemptRefMatcher *UpdatePullRequestCondition1RequestSourceMatcher `json:"exemptRefMatcher,omitempty"`
	RefMatcher RestRefMatcher `json:"refMatcher"`
}

type _RestRequiredBuildConditionSetRequest RestRequiredBuildConditionSetRequest

// NewRestRequiredBuildConditionSetRequest instantiates a new RestRequiredBuildConditionSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestRequiredBuildConditionSetRequest(buildParentKeys []string, refMatcher RestRefMatcher) *RestRequiredBuildConditionSetRequest {
	this := RestRequiredBuildConditionSetRequest{}
	this.BuildParentKeys = buildParentKeys
	this.RefMatcher = refMatcher
	return &this
}

// NewRestRequiredBuildConditionSetRequestWithDefaults instantiates a new RestRequiredBuildConditionSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestRequiredBuildConditionSetRequestWithDefaults() *RestRequiredBuildConditionSetRequest {
	this := RestRequiredBuildConditionSetRequest{}
	return &this
}

// GetBuildParentKeys returns the BuildParentKeys field value
func (o *RestRequiredBuildConditionSetRequest) GetBuildParentKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BuildParentKeys
}

// GetBuildParentKeysOk returns a tuple with the BuildParentKeys field value
// and a boolean to check if the value has been set.
func (o *RestRequiredBuildConditionSetRequest) GetBuildParentKeysOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildParentKeys, true
}

// SetBuildParentKeys sets field value
func (o *RestRequiredBuildConditionSetRequest) SetBuildParentKeys(v []string) {
	o.BuildParentKeys = v
}

// GetExemptRefMatcher returns the ExemptRefMatcher field value if set, zero value otherwise.
func (o *RestRequiredBuildConditionSetRequest) GetExemptRefMatcher() UpdatePullRequestCondition1RequestSourceMatcher {
	if o == nil || IsNil(o.ExemptRefMatcher) {
		var ret UpdatePullRequestCondition1RequestSourceMatcher
		return ret
	}
	return *o.ExemptRefMatcher
}

// GetExemptRefMatcherOk returns a tuple with the ExemptRefMatcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestRequiredBuildConditionSetRequest) GetExemptRefMatcherOk() (*UpdatePullRequestCondition1RequestSourceMatcher, bool) {
	if o == nil || IsNil(o.ExemptRefMatcher) {
		return nil, false
	}
	return o.ExemptRefMatcher, true
}

// HasExemptRefMatcher returns a boolean if a field has been set.
func (o *RestRequiredBuildConditionSetRequest) HasExemptRefMatcher() bool {
	if o != nil && !IsNil(o.ExemptRefMatcher) {
		return true
	}

	return false
}

// SetExemptRefMatcher gets a reference to the given UpdatePullRequestCondition1RequestSourceMatcher and assigns it to the ExemptRefMatcher field.
func (o *RestRequiredBuildConditionSetRequest) SetExemptRefMatcher(v UpdatePullRequestCondition1RequestSourceMatcher) {
	o.ExemptRefMatcher = &v
}

// GetRefMatcher returns the RefMatcher field value
func (o *RestRequiredBuildConditionSetRequest) GetRefMatcher() RestRefMatcher {
	if o == nil {
		var ret RestRefMatcher
		return ret
	}

	return o.RefMatcher
}

// GetRefMatcherOk returns a tuple with the RefMatcher field value
// and a boolean to check if the value has been set.
func (o *RestRequiredBuildConditionSetRequest) GetRefMatcherOk() (*RestRefMatcher, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefMatcher, true
}

// SetRefMatcher sets field value
func (o *RestRequiredBuildConditionSetRequest) SetRefMatcher(v RestRefMatcher) {
	o.RefMatcher = v
}

func (o RestRequiredBuildConditionSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestRequiredBuildConditionSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buildParentKeys"] = o.BuildParentKeys
	if !IsNil(o.ExemptRefMatcher) {
		toSerialize["exemptRefMatcher"] = o.ExemptRefMatcher
	}
	toSerialize["refMatcher"] = o.RefMatcher
	return toSerialize, nil
}

func (o *RestRequiredBuildConditionSetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buildParentKeys",
		"refMatcher",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRestRequiredBuildConditionSetRequest := _RestRequiredBuildConditionSetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestRequiredBuildConditionSetRequest)

	if err != nil {
		return err
	}

	*o = RestRequiredBuildConditionSetRequest(varRestRequiredBuildConditionSetRequest)

	return err
}

type NullableRestRequiredBuildConditionSetRequest struct {
	value *RestRequiredBuildConditionSetRequest
	isSet bool
}

func (v NullableRestRequiredBuildConditionSetRequest) Get() *RestRequiredBuildConditionSetRequest {
	return v.value
}

func (v *NullableRestRequiredBuildConditionSetRequest) Set(val *RestRequiredBuildConditionSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestRequiredBuildConditionSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestRequiredBuildConditionSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestRequiredBuildConditionSetRequest(val *RestRequiredBuildConditionSetRequest) *NullableRestRequiredBuildConditionSetRequest {
	return &NullableRestRequiredBuildConditionSetRequest{value: val, isSet: true}
}

func (v NullableRestRequiredBuildConditionSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestRequiredBuildConditionSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



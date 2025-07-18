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

// checks if the RestBuildStatusSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestBuildStatusSetRequest{}

// RestBuildStatusSetRequest struct for RestBuildStatusSetRequest
type RestBuildStatusSetRequest struct {
	BuildNumber *string `json:"buildNumber,omitempty"`
	Description *string `json:"description,omitempty"`
	Duration *int64 `json:"duration,omitempty"`
	Key string `json:"key"`
	LastUpdated *int64 `json:"lastUpdated,omitempty"`
	Name *string `json:"name,omitempty"`
	Parent *string `json:"parent,omitempty"`
	Ref *string `json:"ref,omitempty" validate:"regexp=^refs\\/.*"`
	State string `json:"state"`
	TestResults *RestBuildStatusTestResults `json:"testResults,omitempty"`
	Url string `json:"url"`
}

type _RestBuildStatusSetRequest RestBuildStatusSetRequest

// NewRestBuildStatusSetRequest instantiates a new RestBuildStatusSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestBuildStatusSetRequest(key string, state string, url string) *RestBuildStatusSetRequest {
	this := RestBuildStatusSetRequest{}
	this.Key = key
	this.State = state
	this.Url = url
	return &this
}

// NewRestBuildStatusSetRequestWithDefaults instantiates a new RestBuildStatusSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestBuildStatusSetRequestWithDefaults() *RestBuildStatusSetRequest {
	this := RestBuildStatusSetRequest{}
	return &this
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *RestBuildStatusSetRequest) GetBuildNumber() string {
	if o == nil || IsNil(o.BuildNumber) {
		var ret string
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetBuildNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BuildNumber) {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *RestBuildStatusSetRequest) HasBuildNumber() bool {
	if o != nil && !IsNil(o.BuildNumber) {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given string and assigns it to the BuildNumber field.
func (o *RestBuildStatusSetRequest) SetBuildNumber(v string) {
	o.BuildNumber = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RestBuildStatusSetRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RestBuildStatusSetRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RestBuildStatusSetRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *RestBuildStatusSetRequest) GetDuration() int64 {
	if o == nil || IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *RestBuildStatusSetRequest) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *RestBuildStatusSetRequest) SetDuration(v int64) {
	o.Duration = &v
}

// GetKey returns the Key field value
func (o *RestBuildStatusSetRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *RestBuildStatusSetRequest) SetKey(v string) {
	o.Key = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *RestBuildStatusSetRequest) GetLastUpdated() int64 {
	if o == nil || IsNil(o.LastUpdated) {
		var ret int64
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetLastUpdatedOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *RestBuildStatusSetRequest) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given int64 and assigns it to the LastUpdated field.
func (o *RestBuildStatusSetRequest) SetLastUpdated(v int64) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RestBuildStatusSetRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RestBuildStatusSetRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RestBuildStatusSetRequest) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *RestBuildStatusSetRequest) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *RestBuildStatusSetRequest) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *RestBuildStatusSetRequest) SetParent(v string) {
	o.Parent = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *RestBuildStatusSetRequest) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *RestBuildStatusSetRequest) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *RestBuildStatusSetRequest) SetRef(v string) {
	o.Ref = &v
}

// GetState returns the State field value
func (o *RestBuildStatusSetRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *RestBuildStatusSetRequest) SetState(v string) {
	o.State = v
}

// GetTestResults returns the TestResults field value if set, zero value otherwise.
func (o *RestBuildStatusSetRequest) GetTestResults() RestBuildStatusTestResults {
	if o == nil || IsNil(o.TestResults) {
		var ret RestBuildStatusTestResults
		return ret
	}
	return *o.TestResults
}

// GetTestResultsOk returns a tuple with the TestResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetTestResultsOk() (*RestBuildStatusTestResults, bool) {
	if o == nil || IsNil(o.TestResults) {
		return nil, false
	}
	return o.TestResults, true
}

// HasTestResults returns a boolean if a field has been set.
func (o *RestBuildStatusSetRequest) HasTestResults() bool {
	if o != nil && !IsNil(o.TestResults) {
		return true
	}

	return false
}

// SetTestResults gets a reference to the given RestBuildStatusTestResults and assigns it to the TestResults field.
func (o *RestBuildStatusSetRequest) SetTestResults(v RestBuildStatusTestResults) {
	o.TestResults = &v
}

// GetUrl returns the Url field value
func (o *RestBuildStatusSetRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RestBuildStatusSetRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RestBuildStatusSetRequest) SetUrl(v string) {
	o.Url = v
}

func (o RestBuildStatusSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestBuildStatusSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildNumber) {
		toSerialize["buildNumber"] = o.BuildNumber
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	toSerialize["state"] = o.State
	if !IsNil(o.TestResults) {
		toSerialize["testResults"] = o.TestResults
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *RestBuildStatusSetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"state",
		"url",
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

	varRestBuildStatusSetRequest := _RestBuildStatusSetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestBuildStatusSetRequest)

	if err != nil {
		return err
	}

	*o = RestBuildStatusSetRequest(varRestBuildStatusSetRequest)

	return err
}

type NullableRestBuildStatusSetRequest struct {
	value *RestBuildStatusSetRequest
	isSet bool
}

func (v NullableRestBuildStatusSetRequest) Get() *RestBuildStatusSetRequest {
	return v.value
}

func (v *NullableRestBuildStatusSetRequest) Set(val *RestBuildStatusSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestBuildStatusSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestBuildStatusSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestBuildStatusSetRequest(val *RestBuildStatusSetRequest) *NullableRestBuildStatusSetRequest {
	return &NullableRestBuildStatusSetRequest{value: val, isSet: true}
}

func (v NullableRestBuildStatusSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestBuildStatusSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



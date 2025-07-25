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

// checks if the ExampleFiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExampleFiles{}

// ExampleFiles struct for ExampleFiles
type ExampleFiles struct {
	Files *ExampleJsonLastModifiedCallback `json:"files,omitempty"`
}

// NewExampleFiles instantiates a new ExampleFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExampleFiles() *ExampleFiles {
	this := ExampleFiles{}
	return &this
}

// NewExampleFilesWithDefaults instantiates a new ExampleFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExampleFilesWithDefaults() *ExampleFiles {
	this := ExampleFiles{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *ExampleFiles) GetFiles() ExampleJsonLastModifiedCallback {
	if o == nil || IsNil(o.Files) {
		var ret ExampleJsonLastModifiedCallback
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExampleFiles) GetFilesOk() (*ExampleJsonLastModifiedCallback, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ExampleFiles) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given ExampleJsonLastModifiedCallback and assigns it to the Files field.
func (o *ExampleFiles) SetFiles(v ExampleJsonLastModifiedCallback) {
	o.Files = &v
}

func (o ExampleFiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExampleFiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	return toSerialize, nil
}

type NullableExampleFiles struct {
	value *ExampleFiles
	isSet bool
}

func (v NullableExampleFiles) Get() *ExampleFiles {
	return v.value
}

func (v *NullableExampleFiles) Set(val *ExampleFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableExampleFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableExampleFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExampleFiles(val *ExampleFiles) *NullableExampleFiles {
	return &NullableExampleFiles{value: val, isSet: true}
}

func (v NullableExampleFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExampleFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



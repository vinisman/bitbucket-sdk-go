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

// checks if the EnrichedRepositoryProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrichedRepositoryProperties{}

// EnrichedRepositoryProperties struct for EnrichedRepositoryProperties
type EnrichedRepositoryProperties struct {
	ContentHash *string `json:"contentHash,omitempty"`
	DefaultBranchId *string `json:"defaultBranchId,omitempty"`
	MetadataHash *string `json:"metadataHash,omitempty"`
}

// NewEnrichedRepositoryProperties instantiates a new EnrichedRepositoryProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrichedRepositoryProperties() *EnrichedRepositoryProperties {
	this := EnrichedRepositoryProperties{}
	return &this
}

// NewEnrichedRepositoryPropertiesWithDefaults instantiates a new EnrichedRepositoryProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrichedRepositoryPropertiesWithDefaults() *EnrichedRepositoryProperties {
	this := EnrichedRepositoryProperties{}
	return &this
}

// GetContentHash returns the ContentHash field value if set, zero value otherwise.
func (o *EnrichedRepositoryProperties) GetContentHash() string {
	if o == nil || IsNil(o.ContentHash) {
		var ret string
		return ret
	}
	return *o.ContentHash
}

// GetContentHashOk returns a tuple with the ContentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrichedRepositoryProperties) GetContentHashOk() (*string, bool) {
	if o == nil || IsNil(o.ContentHash) {
		return nil, false
	}
	return o.ContentHash, true
}

// HasContentHash returns a boolean if a field has been set.
func (o *EnrichedRepositoryProperties) HasContentHash() bool {
	if o != nil && !IsNil(o.ContentHash) {
		return true
	}

	return false
}

// SetContentHash gets a reference to the given string and assigns it to the ContentHash field.
func (o *EnrichedRepositoryProperties) SetContentHash(v string) {
	o.ContentHash = &v
}

// GetDefaultBranchId returns the DefaultBranchId field value if set, zero value otherwise.
func (o *EnrichedRepositoryProperties) GetDefaultBranchId() string {
	if o == nil || IsNil(o.DefaultBranchId) {
		var ret string
		return ret
	}
	return *o.DefaultBranchId
}

// GetDefaultBranchIdOk returns a tuple with the DefaultBranchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrichedRepositoryProperties) GetDefaultBranchIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBranchId) {
		return nil, false
	}
	return o.DefaultBranchId, true
}

// HasDefaultBranchId returns a boolean if a field has been set.
func (o *EnrichedRepositoryProperties) HasDefaultBranchId() bool {
	if o != nil && !IsNil(o.DefaultBranchId) {
		return true
	}

	return false
}

// SetDefaultBranchId gets a reference to the given string and assigns it to the DefaultBranchId field.
func (o *EnrichedRepositoryProperties) SetDefaultBranchId(v string) {
	o.DefaultBranchId = &v
}

// GetMetadataHash returns the MetadataHash field value if set, zero value otherwise.
func (o *EnrichedRepositoryProperties) GetMetadataHash() string {
	if o == nil || IsNil(o.MetadataHash) {
		var ret string
		return ret
	}
	return *o.MetadataHash
}

// GetMetadataHashOk returns a tuple with the MetadataHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrichedRepositoryProperties) GetMetadataHashOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataHash) {
		return nil, false
	}
	return o.MetadataHash, true
}

// HasMetadataHash returns a boolean if a field has been set.
func (o *EnrichedRepositoryProperties) HasMetadataHash() bool {
	if o != nil && !IsNil(o.MetadataHash) {
		return true
	}

	return false
}

// SetMetadataHash gets a reference to the given string and assigns it to the MetadataHash field.
func (o *EnrichedRepositoryProperties) SetMetadataHash(v string) {
	o.MetadataHash = &v
}

func (o EnrichedRepositoryProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrichedRepositoryProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentHash) {
		toSerialize["contentHash"] = o.ContentHash
	}
	if !IsNil(o.DefaultBranchId) {
		toSerialize["defaultBranchId"] = o.DefaultBranchId
	}
	if !IsNil(o.MetadataHash) {
		toSerialize["metadataHash"] = o.MetadataHash
	}
	return toSerialize, nil
}

type NullableEnrichedRepositoryProperties struct {
	value *EnrichedRepositoryProperties
	isSet bool
}

func (v NullableEnrichedRepositoryProperties) Get() *EnrichedRepositoryProperties {
	return v.value
}

func (v *NullableEnrichedRepositoryProperties) Set(val *EnrichedRepositoryProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrichedRepositoryProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrichedRepositoryProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrichedRepositoryProperties(val *EnrichedRepositoryProperties) *NullableEnrichedRepositoryProperties {
	return &NullableEnrichedRepositoryProperties{value: val, isSet: true}
}

func (v NullableEnrichedRepositoryProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrichedRepositoryProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



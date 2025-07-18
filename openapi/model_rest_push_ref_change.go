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

// checks if the RestPushRefChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestPushRefChange{}

// RestPushRefChange struct for RestPushRefChange
type RestPushRefChange struct {
	FromHash *string `json:"fromHash,omitempty"`
	Ref *RestPullRequestRebaseResultRefChangeRef `json:"ref,omitempty"`
	RefId *string `json:"refId,omitempty"`
	ToHash *string `json:"toHash,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedType *string `json:"updatedType,omitempty"`
}

// NewRestPushRefChange instantiates a new RestPushRefChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestPushRefChange() *RestPushRefChange {
	this := RestPushRefChange{}
	return &this
}

// NewRestPushRefChangeWithDefaults instantiates a new RestPushRefChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestPushRefChangeWithDefaults() *RestPushRefChange {
	this := RestPushRefChange{}
	return &this
}

// GetFromHash returns the FromHash field value if set, zero value otherwise.
func (o *RestPushRefChange) GetFromHash() string {
	if o == nil || IsNil(o.FromHash) {
		var ret string
		return ret
	}
	return *o.FromHash
}

// GetFromHashOk returns a tuple with the FromHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPushRefChange) GetFromHashOk() (*string, bool) {
	if o == nil || IsNil(o.FromHash) {
		return nil, false
	}
	return o.FromHash, true
}

// HasFromHash returns a boolean if a field has been set.
func (o *RestPushRefChange) HasFromHash() bool {
	if o != nil && !IsNil(o.FromHash) {
		return true
	}

	return false
}

// SetFromHash gets a reference to the given string and assigns it to the FromHash field.
func (o *RestPushRefChange) SetFromHash(v string) {
	o.FromHash = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *RestPushRefChange) GetRef() RestPullRequestRebaseResultRefChangeRef {
	if o == nil || IsNil(o.Ref) {
		var ret RestPullRequestRebaseResultRefChangeRef
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPushRefChange) GetRefOk() (*RestPullRequestRebaseResultRefChangeRef, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *RestPushRefChange) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given RestPullRequestRebaseResultRefChangeRef and assigns it to the Ref field.
func (o *RestPushRefChange) SetRef(v RestPullRequestRebaseResultRefChangeRef) {
	o.Ref = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *RestPushRefChange) GetRefId() string {
	if o == nil || IsNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPushRefChange) GetRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *RestPushRefChange) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *RestPushRefChange) SetRefId(v string) {
	o.RefId = &v
}

// GetToHash returns the ToHash field value if set, zero value otherwise.
func (o *RestPushRefChange) GetToHash() string {
	if o == nil || IsNil(o.ToHash) {
		var ret string
		return ret
	}
	return *o.ToHash
}

// GetToHashOk returns a tuple with the ToHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPushRefChange) GetToHashOk() (*string, bool) {
	if o == nil || IsNil(o.ToHash) {
		return nil, false
	}
	return o.ToHash, true
}

// HasToHash returns a boolean if a field has been set.
func (o *RestPushRefChange) HasToHash() bool {
	if o != nil && !IsNil(o.ToHash) {
		return true
	}

	return false
}

// SetToHash gets a reference to the given string and assigns it to the ToHash field.
func (o *RestPushRefChange) SetToHash(v string) {
	o.ToHash = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RestPushRefChange) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPushRefChange) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RestPushRefChange) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RestPushRefChange) SetType(v string) {
	o.Type = &v
}

// GetUpdatedType returns the UpdatedType field value if set, zero value otherwise.
func (o *RestPushRefChange) GetUpdatedType() string {
	if o == nil || IsNil(o.UpdatedType) {
		var ret string
		return ret
	}
	return *o.UpdatedType
}

// GetUpdatedTypeOk returns a tuple with the UpdatedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestPushRefChange) GetUpdatedTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedType) {
		return nil, false
	}
	return o.UpdatedType, true
}

// HasUpdatedType returns a boolean if a field has been set.
func (o *RestPushRefChange) HasUpdatedType() bool {
	if o != nil && !IsNil(o.UpdatedType) {
		return true
	}

	return false
}

// SetUpdatedType gets a reference to the given string and assigns it to the UpdatedType field.
func (o *RestPushRefChange) SetUpdatedType(v string) {
	o.UpdatedType = &v
}

func (o RestPushRefChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestPushRefChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromHash) {
		toSerialize["fromHash"] = o.FromHash
	}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}
	if !IsNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !IsNil(o.ToHash) {
		toSerialize["toHash"] = o.ToHash
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedType) {
		toSerialize["updatedType"] = o.UpdatedType
	}
	return toSerialize, nil
}

type NullableRestPushRefChange struct {
	value *RestPushRefChange
	isSet bool
}

func (v NullableRestPushRefChange) Get() *RestPushRefChange {
	return v.value
}

func (v *NullableRestPushRefChange) Set(val *RestPushRefChange) {
	v.value = val
	v.isSet = true
}

func (v NullableRestPushRefChange) IsSet() bool {
	return v.isSet
}

func (v *NullableRestPushRefChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestPushRefChange(val *RestPushRefChange) *NullableRestPushRefChange {
	return &NullableRestPushRefChange{value: val, isSet: true}
}

func (v NullableRestPushRefChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestPushRefChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



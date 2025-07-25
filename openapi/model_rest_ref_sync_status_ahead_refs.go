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

// checks if the RestRefSyncStatusAheadRefs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestRefSyncStatusAheadRefs{}

// RestRefSyncStatusAheadRefs struct for RestRefSyncStatusAheadRefs
type RestRefSyncStatusAheadRefs struct {
	DisplayId *string `json:"displayId,omitempty"`
	Id *string `json:"id,omitempty"`
	State *string `json:"state,omitempty"`
	Tag *bool `json:"tag,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewRestRefSyncStatusAheadRefs instantiates a new RestRefSyncStatusAheadRefs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestRefSyncStatusAheadRefs() *RestRefSyncStatusAheadRefs {
	this := RestRefSyncStatusAheadRefs{}
	return &this
}

// NewRestRefSyncStatusAheadRefsWithDefaults instantiates a new RestRefSyncStatusAheadRefs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestRefSyncStatusAheadRefsWithDefaults() *RestRefSyncStatusAheadRefs {
	this := RestRefSyncStatusAheadRefs{}
	return &this
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *RestRefSyncStatusAheadRefs) GetDisplayId() string {
	if o == nil || IsNil(o.DisplayId) {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestRefSyncStatusAheadRefs) GetDisplayIdOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayId) {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *RestRefSyncStatusAheadRefs) HasDisplayId() bool {
	if o != nil && !IsNil(o.DisplayId) {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *RestRefSyncStatusAheadRefs) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RestRefSyncStatusAheadRefs) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestRefSyncStatusAheadRefs) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RestRefSyncStatusAheadRefs) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RestRefSyncStatusAheadRefs) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RestRefSyncStatusAheadRefs) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestRefSyncStatusAheadRefs) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RestRefSyncStatusAheadRefs) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RestRefSyncStatusAheadRefs) SetState(v string) {
	o.State = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *RestRefSyncStatusAheadRefs) GetTag() bool {
	if o == nil || IsNil(o.Tag) {
		var ret bool
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestRefSyncStatusAheadRefs) GetTagOk() (*bool, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *RestRefSyncStatusAheadRefs) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given bool and assigns it to the Tag field.
func (o *RestRefSyncStatusAheadRefs) SetTag(v bool) {
	o.Tag = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RestRefSyncStatusAheadRefs) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestRefSyncStatusAheadRefs) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RestRefSyncStatusAheadRefs) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RestRefSyncStatusAheadRefs) SetType(v string) {
	o.Type = &v
}

func (o RestRefSyncStatusAheadRefs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestRefSyncStatusAheadRefs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayId) {
		toSerialize["displayId"] = o.DisplayId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableRestRefSyncStatusAheadRefs struct {
	value *RestRefSyncStatusAheadRefs
	isSet bool
}

func (v NullableRestRefSyncStatusAheadRefs) Get() *RestRefSyncStatusAheadRefs {
	return v.value
}

func (v *NullableRestRefSyncStatusAheadRefs) Set(val *RestRefSyncStatusAheadRefs) {
	v.value = val
	v.isSet = true
}

func (v NullableRestRefSyncStatusAheadRefs) IsSet() bool {
	return v.isSet
}

func (v *NullableRestRefSyncStatusAheadRefs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestRefSyncStatusAheadRefs(val *RestRefSyncStatusAheadRefs) *NullableRestRefSyncStatusAheadRefs {
	return &NullableRestRefSyncStatusAheadRefs{value: val, isSet: true}
}

func (v NullableRestRefSyncStatusAheadRefs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestRefSyncStatusAheadRefs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



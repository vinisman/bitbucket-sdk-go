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

// checks if the RestNamedLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestNamedLink{}

// RestNamedLink struct for RestNamedLink
type RestNamedLink struct {
	Href *string `json:"href,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewRestNamedLink instantiates a new RestNamedLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestNamedLink() *RestNamedLink {
	this := RestNamedLink{}
	return &this
}

// NewRestNamedLinkWithDefaults instantiates a new RestNamedLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestNamedLinkWithDefaults() *RestNamedLink {
	this := RestNamedLink{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RestNamedLink) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestNamedLink) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RestNamedLink) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RestNamedLink) SetHref(v string) {
	o.Href = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RestNamedLink) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestNamedLink) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RestNamedLink) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RestNamedLink) SetName(v string) {
	o.Name = &v
}

func (o RestNamedLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestNamedLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableRestNamedLink struct {
	value *RestNamedLink
	isSet bool
}

func (v NullableRestNamedLink) Get() *RestNamedLink {
	return v.value
}

func (v *NullableRestNamedLink) Set(val *RestNamedLink) {
	v.value = val
	v.isSet = true
}

func (v NullableRestNamedLink) IsSet() bool {
	return v.isSet
}

func (v *NullableRestNamedLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestNamedLink(val *RestNamedLink) *NullableRestNamedLink {
	return &NullableRestNamedLink{value: val, isSet: true}
}

func (v NullableRestNamedLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestNamedLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



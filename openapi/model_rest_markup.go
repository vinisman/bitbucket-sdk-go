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

// checks if the RestMarkup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestMarkup{}

// RestMarkup struct for RestMarkup
type RestMarkup struct {
	Html *string `json:"html,omitempty"`
}

// NewRestMarkup instantiates a new RestMarkup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestMarkup() *RestMarkup {
	this := RestMarkup{}
	return &this
}

// NewRestMarkupWithDefaults instantiates a new RestMarkup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestMarkupWithDefaults() *RestMarkup {
	this := RestMarkup{}
	return &this
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *RestMarkup) GetHtml() string {
	if o == nil || IsNil(o.Html) {
		var ret string
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestMarkup) GetHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.Html) {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *RestMarkup) HasHtml() bool {
	if o != nil && !IsNil(o.Html) {
		return true
	}

	return false
}

// SetHtml gets a reference to the given string and assigns it to the Html field.
func (o *RestMarkup) SetHtml(v string) {
	o.Html = &v
}

func (o RestMarkup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestMarkup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Html) {
		toSerialize["html"] = o.Html
	}
	return toSerialize, nil
}

type NullableRestMarkup struct {
	value *RestMarkup
	isSet bool
}

func (v NullableRestMarkup) Get() *RestMarkup {
	return v.value
}

func (v *NullableRestMarkup) Set(val *RestMarkup) {
	v.value = val
	v.isSet = true
}

func (v NullableRestMarkup) IsSet() bool {
	return v.isSet
}

func (v *NullableRestMarkup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestMarkup(val *RestMarkup) *NullableRestMarkup {
	return &NullableRestMarkup{value: val, isSet: true}
}

func (v NullableRestMarkup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestMarkup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



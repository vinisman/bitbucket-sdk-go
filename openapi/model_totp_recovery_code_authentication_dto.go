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

// checks if the TotpRecoveryCodeAuthenticationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TotpRecoveryCodeAuthenticationDTO{}

// TotpRecoveryCodeAuthenticationDTO struct for TotpRecoveryCodeAuthenticationDTO
type TotpRecoveryCodeAuthenticationDTO struct {
	ConversationId *string `json:"conversationId,omitempty"`
	RecoveryCode *string `json:"recoveryCode,omitempty"`
}

// NewTotpRecoveryCodeAuthenticationDTO instantiates a new TotpRecoveryCodeAuthenticationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotpRecoveryCodeAuthenticationDTO() *TotpRecoveryCodeAuthenticationDTO {
	this := TotpRecoveryCodeAuthenticationDTO{}
	return &this
}

// NewTotpRecoveryCodeAuthenticationDTOWithDefaults instantiates a new TotpRecoveryCodeAuthenticationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotpRecoveryCodeAuthenticationDTOWithDefaults() *TotpRecoveryCodeAuthenticationDTO {
	this := TotpRecoveryCodeAuthenticationDTO{}
	return &this
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *TotpRecoveryCodeAuthenticationDTO) GetConversationId() string {
	if o == nil || IsNil(o.ConversationId) {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TotpRecoveryCodeAuthenticationDTO) GetConversationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConversationId) {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *TotpRecoveryCodeAuthenticationDTO) HasConversationId() bool {
	if o != nil && !IsNil(o.ConversationId) {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *TotpRecoveryCodeAuthenticationDTO) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetRecoveryCode returns the RecoveryCode field value if set, zero value otherwise.
func (o *TotpRecoveryCodeAuthenticationDTO) GetRecoveryCode() string {
	if o == nil || IsNil(o.RecoveryCode) {
		var ret string
		return ret
	}
	return *o.RecoveryCode
}

// GetRecoveryCodeOk returns a tuple with the RecoveryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TotpRecoveryCodeAuthenticationDTO) GetRecoveryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RecoveryCode) {
		return nil, false
	}
	return o.RecoveryCode, true
}

// HasRecoveryCode returns a boolean if a field has been set.
func (o *TotpRecoveryCodeAuthenticationDTO) HasRecoveryCode() bool {
	if o != nil && !IsNil(o.RecoveryCode) {
		return true
	}

	return false
}

// SetRecoveryCode gets a reference to the given string and assigns it to the RecoveryCode field.
func (o *TotpRecoveryCodeAuthenticationDTO) SetRecoveryCode(v string) {
	o.RecoveryCode = &v
}

func (o TotpRecoveryCodeAuthenticationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TotpRecoveryCodeAuthenticationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConversationId) {
		toSerialize["conversationId"] = o.ConversationId
	}
	if !IsNil(o.RecoveryCode) {
		toSerialize["recoveryCode"] = o.RecoveryCode
	}
	return toSerialize, nil
}

type NullableTotpRecoveryCodeAuthenticationDTO struct {
	value *TotpRecoveryCodeAuthenticationDTO
	isSet bool
}

func (v NullableTotpRecoveryCodeAuthenticationDTO) Get() *TotpRecoveryCodeAuthenticationDTO {
	return v.value
}

func (v *NullableTotpRecoveryCodeAuthenticationDTO) Set(val *TotpRecoveryCodeAuthenticationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTotpRecoveryCodeAuthenticationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTotpRecoveryCodeAuthenticationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotpRecoveryCodeAuthenticationDTO(val *TotpRecoveryCodeAuthenticationDTO) *NullableTotpRecoveryCodeAuthenticationDTO {
	return &NullableTotpRecoveryCodeAuthenticationDTO{value: val, isSet: true}
}

func (v NullableTotpRecoveryCodeAuthenticationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotpRecoveryCodeAuthenticationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



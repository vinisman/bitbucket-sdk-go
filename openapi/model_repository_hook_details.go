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

// checks if the RepositoryHookDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepositoryHookDetails{}

// RepositoryHookDetails struct for RepositoryHookDetails
type RepositoryHookDetails struct {
	ConfigFormKey *string `json:"configFormKey,omitempty"`
	Description *string `json:"description,omitempty"`
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	SupportedScopes []string `json:"supportedScopes,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewRepositoryHookDetails instantiates a new RepositoryHookDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoryHookDetails() *RepositoryHookDetails {
	this := RepositoryHookDetails{}
	return &this
}

// NewRepositoryHookDetailsWithDefaults instantiates a new RepositoryHookDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryHookDetailsWithDefaults() *RepositoryHookDetails {
	this := RepositoryHookDetails{}
	return &this
}

// GetConfigFormKey returns the ConfigFormKey field value if set, zero value otherwise.
func (o *RepositoryHookDetails) GetConfigFormKey() string {
	if o == nil || IsNil(o.ConfigFormKey) {
		var ret string
		return ret
	}
	return *o.ConfigFormKey
}

// GetConfigFormKeyOk returns a tuple with the ConfigFormKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryHookDetails) GetConfigFormKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigFormKey) {
		return nil, false
	}
	return o.ConfigFormKey, true
}

// HasConfigFormKey returns a boolean if a field has been set.
func (o *RepositoryHookDetails) HasConfigFormKey() bool {
	if o != nil && !IsNil(o.ConfigFormKey) {
		return true
	}

	return false
}

// SetConfigFormKey gets a reference to the given string and assigns it to the ConfigFormKey field.
func (o *RepositoryHookDetails) SetConfigFormKey(v string) {
	o.ConfigFormKey = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RepositoryHookDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryHookDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RepositoryHookDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RepositoryHookDetails) SetDescription(v string) {
	o.Description = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RepositoryHookDetails) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryHookDetails) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RepositoryHookDetails) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RepositoryHookDetails) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RepositoryHookDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryHookDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RepositoryHookDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RepositoryHookDetails) SetName(v string) {
	o.Name = &v
}

// GetSupportedScopes returns the SupportedScopes field value if set, zero value otherwise.
func (o *RepositoryHookDetails) GetSupportedScopes() []string {
	if o == nil || IsNil(o.SupportedScopes) {
		var ret []string
		return ret
	}
	return o.SupportedScopes
}

// GetSupportedScopesOk returns a tuple with the SupportedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryHookDetails) GetSupportedScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedScopes) {
		return nil, false
	}
	return o.SupportedScopes, true
}

// HasSupportedScopes returns a boolean if a field has been set.
func (o *RepositoryHookDetails) HasSupportedScopes() bool {
	if o != nil && !IsNil(o.SupportedScopes) {
		return true
	}

	return false
}

// SetSupportedScopes gets a reference to the given []string and assigns it to the SupportedScopes field.
func (o *RepositoryHookDetails) SetSupportedScopes(v []string) {
	o.SupportedScopes = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RepositoryHookDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryHookDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RepositoryHookDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RepositoryHookDetails) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RepositoryHookDetails) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepositoryHookDetails) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RepositoryHookDetails) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RepositoryHookDetails) SetVersion(v string) {
	o.Version = &v
}

func (o RepositoryHookDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepositoryHookDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigFormKey) {
		toSerialize["configFormKey"] = o.ConfigFormKey
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SupportedScopes) {
		toSerialize["supportedScopes"] = o.SupportedScopes
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableRepositoryHookDetails struct {
	value *RepositoryHookDetails
	isSet bool
}

func (v NullableRepositoryHookDetails) Get() *RepositoryHookDetails {
	return v.value
}

func (v *NullableRepositoryHookDetails) Set(val *RepositoryHookDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoryHookDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoryHookDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoryHookDetails(val *RepositoryHookDetails) *NullableRepositoryHookDetails {
	return &NullableRepositoryHookDetails{value: val, isSet: true}
}

func (v NullableRepositoryHookDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoryHookDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the RestEnhancedEntityLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestEnhancedEntityLink{}

// RestEnhancedEntityLink struct for RestEnhancedEntityLink
type RestEnhancedEntityLink struct {
	ApplicationLinkId *string `json:"applicationLinkId,omitempty"`
	DisplayUrl *string `json:"displayUrl,omitempty"`
	ProjectId *int64 `json:"projectId,omitempty"`
	ProjectKey *string `json:"projectKey,omitempty"`
	ProjectName *string `json:"projectName,omitempty"`
}

// NewRestEnhancedEntityLink instantiates a new RestEnhancedEntityLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestEnhancedEntityLink() *RestEnhancedEntityLink {
	this := RestEnhancedEntityLink{}
	return &this
}

// NewRestEnhancedEntityLinkWithDefaults instantiates a new RestEnhancedEntityLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestEnhancedEntityLinkWithDefaults() *RestEnhancedEntityLink {
	this := RestEnhancedEntityLink{}
	return &this
}

// GetApplicationLinkId returns the ApplicationLinkId field value if set, zero value otherwise.
func (o *RestEnhancedEntityLink) GetApplicationLinkId() string {
	if o == nil || IsNil(o.ApplicationLinkId) {
		var ret string
		return ret
	}
	return *o.ApplicationLinkId
}

// GetApplicationLinkIdOk returns a tuple with the ApplicationLinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestEnhancedEntityLink) GetApplicationLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationLinkId) {
		return nil, false
	}
	return o.ApplicationLinkId, true
}

// HasApplicationLinkId returns a boolean if a field has been set.
func (o *RestEnhancedEntityLink) HasApplicationLinkId() bool {
	if o != nil && !IsNil(o.ApplicationLinkId) {
		return true
	}

	return false
}

// SetApplicationLinkId gets a reference to the given string and assigns it to the ApplicationLinkId field.
func (o *RestEnhancedEntityLink) SetApplicationLinkId(v string) {
	o.ApplicationLinkId = &v
}

// GetDisplayUrl returns the DisplayUrl field value if set, zero value otherwise.
func (o *RestEnhancedEntityLink) GetDisplayUrl() string {
	if o == nil || IsNil(o.DisplayUrl) {
		var ret string
		return ret
	}
	return *o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestEnhancedEntityLink) GetDisplayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUrl) {
		return nil, false
	}
	return o.DisplayUrl, true
}

// HasDisplayUrl returns a boolean if a field has been set.
func (o *RestEnhancedEntityLink) HasDisplayUrl() bool {
	if o != nil && !IsNil(o.DisplayUrl) {
		return true
	}

	return false
}

// SetDisplayUrl gets a reference to the given string and assigns it to the DisplayUrl field.
func (o *RestEnhancedEntityLink) SetDisplayUrl(v string) {
	o.DisplayUrl = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *RestEnhancedEntityLink) GetProjectId() int64 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int64
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestEnhancedEntityLink) GetProjectIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *RestEnhancedEntityLink) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int64 and assigns it to the ProjectId field.
func (o *RestEnhancedEntityLink) SetProjectId(v int64) {
	o.ProjectId = &v
}

// GetProjectKey returns the ProjectKey field value if set, zero value otherwise.
func (o *RestEnhancedEntityLink) GetProjectKey() string {
	if o == nil || IsNil(o.ProjectKey) {
		var ret string
		return ret
	}
	return *o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestEnhancedEntityLink) GetProjectKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectKey) {
		return nil, false
	}
	return o.ProjectKey, true
}

// HasProjectKey returns a boolean if a field has been set.
func (o *RestEnhancedEntityLink) HasProjectKey() bool {
	if o != nil && !IsNil(o.ProjectKey) {
		return true
	}

	return false
}

// SetProjectKey gets a reference to the given string and assigns it to the ProjectKey field.
func (o *RestEnhancedEntityLink) SetProjectKey(v string) {
	o.ProjectKey = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *RestEnhancedEntityLink) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestEnhancedEntityLink) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *RestEnhancedEntityLink) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *RestEnhancedEntityLink) SetProjectName(v string) {
	o.ProjectName = &v
}

func (o RestEnhancedEntityLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestEnhancedEntityLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationLinkId) {
		toSerialize["applicationLinkId"] = o.ApplicationLinkId
	}
	if !IsNil(o.DisplayUrl) {
		toSerialize["displayUrl"] = o.DisplayUrl
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.ProjectKey) {
		toSerialize["projectKey"] = o.ProjectKey
	}
	if !IsNil(o.ProjectName) {
		toSerialize["projectName"] = o.ProjectName
	}
	return toSerialize, nil
}

type NullableRestEnhancedEntityLink struct {
	value *RestEnhancedEntityLink
	isSet bool
}

func (v NullableRestEnhancedEntityLink) Get() *RestEnhancedEntityLink {
	return v.value
}

func (v *NullableRestEnhancedEntityLink) Set(val *RestEnhancedEntityLink) {
	v.value = val
	v.isSet = true
}

func (v NullableRestEnhancedEntityLink) IsSet() bool {
	return v.isSet
}

func (v *NullableRestEnhancedEntityLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestEnhancedEntityLink(val *RestEnhancedEntityLink) *NullableRestEnhancedEntityLink {
	return &NullableRestEnhancedEntityLink{value: val, isSet: true}
}

func (v NullableRestEnhancedEntityLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestEnhancedEntityLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



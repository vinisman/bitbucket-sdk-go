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

// checks if the GetIdps200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIdps200Response{}

// GetIdps200Response struct for GetIdps200Response
type GetIdps200Response struct {
	IsLastPage *bool `json:"isLastPage,omitempty"`
	Limit *float32 `json:"limit,omitempty"`
	Results []IdpConfigEntity `json:"results,omitempty"`
	Size *float32 `json:"size,omitempty"`
	Start *int32 `json:"start,omitempty"`
}

// NewGetIdps200Response instantiates a new GetIdps200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIdps200Response() *GetIdps200Response {
	this := GetIdps200Response{}
	return &this
}

// NewGetIdps200ResponseWithDefaults instantiates a new GetIdps200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIdps200ResponseWithDefaults() *GetIdps200Response {
	this := GetIdps200Response{}
	return &this
}

// GetIsLastPage returns the IsLastPage field value if set, zero value otherwise.
func (o *GetIdps200Response) GetIsLastPage() bool {
	if o == nil || IsNil(o.IsLastPage) {
		var ret bool
		return ret
	}
	return *o.IsLastPage
}

// GetIsLastPageOk returns a tuple with the IsLastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdps200Response) GetIsLastPageOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLastPage) {
		return nil, false
	}
	return o.IsLastPage, true
}

// HasIsLastPage returns a boolean if a field has been set.
func (o *GetIdps200Response) HasIsLastPage() bool {
	if o != nil && !IsNil(o.IsLastPage) {
		return true
	}

	return false
}

// SetIsLastPage gets a reference to the given bool and assigns it to the IsLastPage field.
func (o *GetIdps200Response) SetIsLastPage(v bool) {
	o.IsLastPage = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetIdps200Response) GetLimit() float32 {
	if o == nil || IsNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdps200Response) GetLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetIdps200Response) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *GetIdps200Response) SetLimit(v float32) {
	o.Limit = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetIdps200Response) GetResults() []IdpConfigEntity {
	if o == nil || IsNil(o.Results) {
		var ret []IdpConfigEntity
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdps200Response) GetResultsOk() ([]IdpConfigEntity, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetIdps200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []IdpConfigEntity and assigns it to the Results field.
func (o *GetIdps200Response) SetResults(v []IdpConfigEntity) {
	o.Results = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GetIdps200Response) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdps200Response) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GetIdps200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *GetIdps200Response) SetSize(v float32) {
	o.Size = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *GetIdps200Response) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIdps200Response) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *GetIdps200Response) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *GetIdps200Response) SetStart(v int32) {
	o.Start = &v
}

func (o GetIdps200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIdps200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsLastPage) {
		toSerialize["isLastPage"] = o.IsLastPage
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return toSerialize, nil
}

type NullableGetIdps200Response struct {
	value *GetIdps200Response
	isSet bool
}

func (v NullableGetIdps200Response) Get() *GetIdps200Response {
	return v.value
}

func (v *NullableGetIdps200Response) Set(val *GetIdps200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIdps200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIdps200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIdps200Response(val *GetIdps200Response) *NullableGetIdps200Response {
	return &NullableGetIdps200Response{value: val, isSet: true}
}

func (v NullableGetIdps200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIdps200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the RestDiffHunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestDiffHunk{}

// RestDiffHunk struct for RestDiffHunk
type RestDiffHunk struct {
	Context *string `json:"context,omitempty"`
	DestinationLine *int32 `json:"destinationLine,omitempty"`
	DestinationSpan *int32 `json:"destinationSpan,omitempty"`
	Segments []RestDiffSegment `json:"segments,omitempty"`
	SourceLine *int32 `json:"sourceLine,omitempty"`
	SourceSpan *int32 `json:"sourceSpan,omitempty"`
	Truncated *bool `json:"truncated,omitempty"`
}

// NewRestDiffHunk instantiates a new RestDiffHunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestDiffHunk() *RestDiffHunk {
	this := RestDiffHunk{}
	return &this
}

// NewRestDiffHunkWithDefaults instantiates a new RestDiffHunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestDiffHunkWithDefaults() *RestDiffHunk {
	this := RestDiffHunk{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RestDiffHunk) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestDiffHunk) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RestDiffHunk) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *RestDiffHunk) SetContext(v string) {
	o.Context = &v
}

// GetDestinationLine returns the DestinationLine field value if set, zero value otherwise.
func (o *RestDiffHunk) GetDestinationLine() int32 {
	if o == nil || IsNil(o.DestinationLine) {
		var ret int32
		return ret
	}
	return *o.DestinationLine
}

// GetDestinationLineOk returns a tuple with the DestinationLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestDiffHunk) GetDestinationLineOk() (*int32, bool) {
	if o == nil || IsNil(o.DestinationLine) {
		return nil, false
	}
	return o.DestinationLine, true
}

// HasDestinationLine returns a boolean if a field has been set.
func (o *RestDiffHunk) HasDestinationLine() bool {
	if o != nil && !IsNil(o.DestinationLine) {
		return true
	}

	return false
}

// SetDestinationLine gets a reference to the given int32 and assigns it to the DestinationLine field.
func (o *RestDiffHunk) SetDestinationLine(v int32) {
	o.DestinationLine = &v
}

// GetDestinationSpan returns the DestinationSpan field value if set, zero value otherwise.
func (o *RestDiffHunk) GetDestinationSpan() int32 {
	if o == nil || IsNil(o.DestinationSpan) {
		var ret int32
		return ret
	}
	return *o.DestinationSpan
}

// GetDestinationSpanOk returns a tuple with the DestinationSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestDiffHunk) GetDestinationSpanOk() (*int32, bool) {
	if o == nil || IsNil(o.DestinationSpan) {
		return nil, false
	}
	return o.DestinationSpan, true
}

// HasDestinationSpan returns a boolean if a field has been set.
func (o *RestDiffHunk) HasDestinationSpan() bool {
	if o != nil && !IsNil(o.DestinationSpan) {
		return true
	}

	return false
}

// SetDestinationSpan gets a reference to the given int32 and assigns it to the DestinationSpan field.
func (o *RestDiffHunk) SetDestinationSpan(v int32) {
	o.DestinationSpan = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *RestDiffHunk) GetSegments() []RestDiffSegment {
	if o == nil || IsNil(o.Segments) {
		var ret []RestDiffSegment
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestDiffHunk) GetSegmentsOk() ([]RestDiffSegment, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *RestDiffHunk) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []RestDiffSegment and assigns it to the Segments field.
func (o *RestDiffHunk) SetSegments(v []RestDiffSegment) {
	o.Segments = v
}

// GetSourceLine returns the SourceLine field value if set, zero value otherwise.
func (o *RestDiffHunk) GetSourceLine() int32 {
	if o == nil || IsNil(o.SourceLine) {
		var ret int32
		return ret
	}
	return *o.SourceLine
}

// GetSourceLineOk returns a tuple with the SourceLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestDiffHunk) GetSourceLineOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceLine) {
		return nil, false
	}
	return o.SourceLine, true
}

// HasSourceLine returns a boolean if a field has been set.
func (o *RestDiffHunk) HasSourceLine() bool {
	if o != nil && !IsNil(o.SourceLine) {
		return true
	}

	return false
}

// SetSourceLine gets a reference to the given int32 and assigns it to the SourceLine field.
func (o *RestDiffHunk) SetSourceLine(v int32) {
	o.SourceLine = &v
}

// GetSourceSpan returns the SourceSpan field value if set, zero value otherwise.
func (o *RestDiffHunk) GetSourceSpan() int32 {
	if o == nil || IsNil(o.SourceSpan) {
		var ret int32
		return ret
	}
	return *o.SourceSpan
}

// GetSourceSpanOk returns a tuple with the SourceSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestDiffHunk) GetSourceSpanOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceSpan) {
		return nil, false
	}
	return o.SourceSpan, true
}

// HasSourceSpan returns a boolean if a field has been set.
func (o *RestDiffHunk) HasSourceSpan() bool {
	if o != nil && !IsNil(o.SourceSpan) {
		return true
	}

	return false
}

// SetSourceSpan gets a reference to the given int32 and assigns it to the SourceSpan field.
func (o *RestDiffHunk) SetSourceSpan(v int32) {
	o.SourceSpan = &v
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *RestDiffHunk) GetTruncated() bool {
	if o == nil || IsNil(o.Truncated) {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestDiffHunk) GetTruncatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Truncated) {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *RestDiffHunk) HasTruncated() bool {
	if o != nil && !IsNil(o.Truncated) {
		return true
	}

	return false
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *RestDiffHunk) SetTruncated(v bool) {
	o.Truncated = &v
}

func (o RestDiffHunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestDiffHunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.DestinationLine) {
		toSerialize["destinationLine"] = o.DestinationLine
	}
	if !IsNil(o.DestinationSpan) {
		toSerialize["destinationSpan"] = o.DestinationSpan
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.SourceLine) {
		toSerialize["sourceLine"] = o.SourceLine
	}
	if !IsNil(o.SourceSpan) {
		toSerialize["sourceSpan"] = o.SourceSpan
	}
	if !IsNil(o.Truncated) {
		toSerialize["truncated"] = o.Truncated
	}
	return toSerialize, nil
}

type NullableRestDiffHunk struct {
	value *RestDiffHunk
	isSet bool
}

func (v NullableRestDiffHunk) Get() *RestDiffHunk {
	return v.value
}

func (v *NullableRestDiffHunk) Set(val *RestDiffHunk) {
	v.value = val
	v.isSet = true
}

func (v NullableRestDiffHunk) IsSet() bool {
	return v.isSet
}

func (v *NullableRestDiffHunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestDiffHunk(val *RestDiffHunk) *NullableRestDiffHunk {
	return &NullableRestDiffHunk{value: val, isSet: true}
}

func (v NullableRestDiffHunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestDiffHunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the RestChangesetToCommit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestChangesetToCommit{}

// RestChangesetToCommit struct for RestChangesetToCommit
type RestChangesetToCommit struct {
	Author *RestChangesetToCommitAuthor `json:"author,omitempty"`
	AuthorTimestamp *int64 `json:"authorTimestamp,omitempty"`
	Committer *RestChangesetToCommitAuthor `json:"committer,omitempty"`
	CommitterTimestamp *int64 `json:"committerTimestamp,omitempty"`
	DisplayId *string `json:"displayId,omitempty"`
	Id *string `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
	Parents []RestMinimalCommit `json:"parents,omitempty"`
}

// NewRestChangesetToCommit instantiates a new RestChangesetToCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestChangesetToCommit() *RestChangesetToCommit {
	this := RestChangesetToCommit{}
	return &this
}

// NewRestChangesetToCommitWithDefaults instantiates a new RestChangesetToCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestChangesetToCommitWithDefaults() *RestChangesetToCommit {
	this := RestChangesetToCommit{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *RestChangesetToCommit) GetAuthor() RestChangesetToCommitAuthor {
	if o == nil || IsNil(o.Author) {
		var ret RestChangesetToCommitAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetToCommit) GetAuthorOk() (*RestChangesetToCommitAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *RestChangesetToCommit) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given RestChangesetToCommitAuthor and assigns it to the Author field.
func (o *RestChangesetToCommit) SetAuthor(v RestChangesetToCommitAuthor) {
	o.Author = &v
}

// GetAuthorTimestamp returns the AuthorTimestamp field value if set, zero value otherwise.
func (o *RestChangesetToCommit) GetAuthorTimestamp() int64 {
	if o == nil || IsNil(o.AuthorTimestamp) {
		var ret int64
		return ret
	}
	return *o.AuthorTimestamp
}

// GetAuthorTimestampOk returns a tuple with the AuthorTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetToCommit) GetAuthorTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthorTimestamp) {
		return nil, false
	}
	return o.AuthorTimestamp, true
}

// HasAuthorTimestamp returns a boolean if a field has been set.
func (o *RestChangesetToCommit) HasAuthorTimestamp() bool {
	if o != nil && !IsNil(o.AuthorTimestamp) {
		return true
	}

	return false
}

// SetAuthorTimestamp gets a reference to the given int64 and assigns it to the AuthorTimestamp field.
func (o *RestChangesetToCommit) SetAuthorTimestamp(v int64) {
	o.AuthorTimestamp = &v
}

// GetCommitter returns the Committer field value if set, zero value otherwise.
func (o *RestChangesetToCommit) GetCommitter() RestChangesetToCommitAuthor {
	if o == nil || IsNil(o.Committer) {
		var ret RestChangesetToCommitAuthor
		return ret
	}
	return *o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetToCommit) GetCommitterOk() (*RestChangesetToCommitAuthor, bool) {
	if o == nil || IsNil(o.Committer) {
		return nil, false
	}
	return o.Committer, true
}

// HasCommitter returns a boolean if a field has been set.
func (o *RestChangesetToCommit) HasCommitter() bool {
	if o != nil && !IsNil(o.Committer) {
		return true
	}

	return false
}

// SetCommitter gets a reference to the given RestChangesetToCommitAuthor and assigns it to the Committer field.
func (o *RestChangesetToCommit) SetCommitter(v RestChangesetToCommitAuthor) {
	o.Committer = &v
}

// GetCommitterTimestamp returns the CommitterTimestamp field value if set, zero value otherwise.
func (o *RestChangesetToCommit) GetCommitterTimestamp() int64 {
	if o == nil || IsNil(o.CommitterTimestamp) {
		var ret int64
		return ret
	}
	return *o.CommitterTimestamp
}

// GetCommitterTimestampOk returns a tuple with the CommitterTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetToCommit) GetCommitterTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CommitterTimestamp) {
		return nil, false
	}
	return o.CommitterTimestamp, true
}

// HasCommitterTimestamp returns a boolean if a field has been set.
func (o *RestChangesetToCommit) HasCommitterTimestamp() bool {
	if o != nil && !IsNil(o.CommitterTimestamp) {
		return true
	}

	return false
}

// SetCommitterTimestamp gets a reference to the given int64 and assigns it to the CommitterTimestamp field.
func (o *RestChangesetToCommit) SetCommitterTimestamp(v int64) {
	o.CommitterTimestamp = &v
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *RestChangesetToCommit) GetDisplayId() string {
	if o == nil || IsNil(o.DisplayId) {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetToCommit) GetDisplayIdOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayId) {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *RestChangesetToCommit) HasDisplayId() bool {
	if o != nil && !IsNil(o.DisplayId) {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *RestChangesetToCommit) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RestChangesetToCommit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetToCommit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RestChangesetToCommit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RestChangesetToCommit) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RestChangesetToCommit) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetToCommit) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RestChangesetToCommit) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RestChangesetToCommit) SetMessage(v string) {
	o.Message = &v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *RestChangesetToCommit) GetParents() []RestMinimalCommit {
	if o == nil || IsNil(o.Parents) {
		var ret []RestMinimalCommit
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetToCommit) GetParentsOk() ([]RestMinimalCommit, bool) {
	if o == nil || IsNil(o.Parents) {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *RestChangesetToCommit) HasParents() bool {
	if o != nil && !IsNil(o.Parents) {
		return true
	}

	return false
}

// SetParents gets a reference to the given []RestMinimalCommit and assigns it to the Parents field.
func (o *RestChangesetToCommit) SetParents(v []RestMinimalCommit) {
	o.Parents = v
}

func (o RestChangesetToCommit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestChangesetToCommit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.AuthorTimestamp) {
		toSerialize["authorTimestamp"] = o.AuthorTimestamp
	}
	if !IsNil(o.Committer) {
		toSerialize["committer"] = o.Committer
	}
	if !IsNil(o.CommitterTimestamp) {
		toSerialize["committerTimestamp"] = o.CommitterTimestamp
	}
	if !IsNil(o.DisplayId) {
		toSerialize["displayId"] = o.DisplayId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Parents) {
		toSerialize["parents"] = o.Parents
	}
	return toSerialize, nil
}

type NullableRestChangesetToCommit struct {
	value *RestChangesetToCommit
	isSet bool
}

func (v NullableRestChangesetToCommit) Get() *RestChangesetToCommit {
	return v.value
}

func (v *NullableRestChangesetToCommit) Set(val *RestChangesetToCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableRestChangesetToCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableRestChangesetToCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestChangesetToCommit(val *RestChangesetToCommit) *NullableRestChangesetToCommit {
	return &NullableRestChangesetToCommit{value: val, isSet: true}
}

func (v NullableRestChangesetToCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestChangesetToCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



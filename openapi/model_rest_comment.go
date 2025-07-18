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

// checks if the RestComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestComment{}

// RestComment struct for RestComment
type RestComment struct {
	Anchor *RestCommentAnchor `json:"anchor,omitempty"`
	Anchored *bool `json:"anchored,omitempty"`
	Author *RestCommentAuthor `json:"author,omitempty"`
	Comments []RestComment `json:"comments,omitempty"`
	CreatedDate *int64 `json:"createdDate,omitempty"`
	Html *string `json:"html,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Parent *RestCommentParent `json:"parent,omitempty"`
	Pending *bool `json:"pending,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Reply *bool `json:"reply,omitempty"`
	ResolvedDate *int64 `json:"resolvedDate,omitempty"`
	Resolver *RestCommentAuthor `json:"resolver,omitempty"`
	Severity *string `json:"severity,omitempty"`
	State *string `json:"state,omitempty"`
	Text *string `json:"text,omitempty"`
	// Indicates if this comment thread has been marked as resolved or not
	ThreadResolved *bool `json:"threadResolved,omitempty"`
	ThreadResolvedDate *int64 `json:"threadResolvedDate,omitempty"`
	ThreadResolver *RestCommentAuthor `json:"threadResolver,omitempty"`
	UpdatedDate *int64 `json:"updatedDate,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewRestComment instantiates a new RestComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestComment() *RestComment {
	this := RestComment{}
	return &this
}

// NewRestCommentWithDefaults instantiates a new RestComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestCommentWithDefaults() *RestComment {
	this := RestComment{}
	return &this
}

// GetAnchor returns the Anchor field value if set, zero value otherwise.
func (o *RestComment) GetAnchor() RestCommentAnchor {
	if o == nil || IsNil(o.Anchor) {
		var ret RestCommentAnchor
		return ret
	}
	return *o.Anchor
}

// GetAnchorOk returns a tuple with the Anchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetAnchorOk() (*RestCommentAnchor, bool) {
	if o == nil || IsNil(o.Anchor) {
		return nil, false
	}
	return o.Anchor, true
}

// HasAnchor returns a boolean if a field has been set.
func (o *RestComment) HasAnchor() bool {
	if o != nil && !IsNil(o.Anchor) {
		return true
	}

	return false
}

// SetAnchor gets a reference to the given RestCommentAnchor and assigns it to the Anchor field.
func (o *RestComment) SetAnchor(v RestCommentAnchor) {
	o.Anchor = &v
}

// GetAnchored returns the Anchored field value if set, zero value otherwise.
func (o *RestComment) GetAnchored() bool {
	if o == nil || IsNil(o.Anchored) {
		var ret bool
		return ret
	}
	return *o.Anchored
}

// GetAnchoredOk returns a tuple with the Anchored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetAnchoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Anchored) {
		return nil, false
	}
	return o.Anchored, true
}

// HasAnchored returns a boolean if a field has been set.
func (o *RestComment) HasAnchored() bool {
	if o != nil && !IsNil(o.Anchored) {
		return true
	}

	return false
}

// SetAnchored gets a reference to the given bool and assigns it to the Anchored field.
func (o *RestComment) SetAnchored(v bool) {
	o.Anchored = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *RestComment) GetAuthor() RestCommentAuthor {
	if o == nil || IsNil(o.Author) {
		var ret RestCommentAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetAuthorOk() (*RestCommentAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *RestComment) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given RestCommentAuthor and assigns it to the Author field.
func (o *RestComment) SetAuthor(v RestCommentAuthor) {
	o.Author = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *RestComment) GetComments() []RestComment {
	if o == nil || IsNil(o.Comments) {
		var ret []RestComment
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetCommentsOk() ([]RestComment, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *RestComment) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given []RestComment and assigns it to the Comments field.
func (o *RestComment) SetComments(v []RestComment) {
	o.Comments = v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *RestComment) GetCreatedDate() int64 {
	if o == nil || IsNil(o.CreatedDate) {
		var ret int64
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetCreatedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *RestComment) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given int64 and assigns it to the CreatedDate field.
func (o *RestComment) SetCreatedDate(v int64) {
	o.CreatedDate = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *RestComment) GetHtml() string {
	if o == nil || IsNil(o.Html) {
		var ret string
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.Html) {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *RestComment) HasHtml() bool {
	if o != nil && !IsNil(o.Html) {
		return true
	}

	return false
}

// SetHtml gets a reference to the given string and assigns it to the Html field.
func (o *RestComment) SetHtml(v string) {
	o.Html = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RestComment) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RestComment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RestComment) SetId(v int64) {
	o.Id = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *RestComment) GetParent() RestCommentParent {
	if o == nil || IsNil(o.Parent) {
		var ret RestCommentParent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetParentOk() (*RestCommentParent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *RestComment) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given RestCommentParent and assigns it to the Parent field.
func (o *RestComment) SetParent(v RestCommentParent) {
	o.Parent = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *RestComment) GetPending() bool {
	if o == nil || IsNil(o.Pending) {
		var ret bool
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetPendingOk() (*bool, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *RestComment) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given bool and assigns it to the Pending field.
func (o *RestComment) SetPending(v bool) {
	o.Pending = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *RestComment) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *RestComment) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *RestComment) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetReply returns the Reply field value if set, zero value otherwise.
func (o *RestComment) GetReply() bool {
	if o == nil || IsNil(o.Reply) {
		var ret bool
		return ret
	}
	return *o.Reply
}

// GetReplyOk returns a tuple with the Reply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetReplyOk() (*bool, bool) {
	if o == nil || IsNil(o.Reply) {
		return nil, false
	}
	return o.Reply, true
}

// HasReply returns a boolean if a field has been set.
func (o *RestComment) HasReply() bool {
	if o != nil && !IsNil(o.Reply) {
		return true
	}

	return false
}

// SetReply gets a reference to the given bool and assigns it to the Reply field.
func (o *RestComment) SetReply(v bool) {
	o.Reply = &v
}

// GetResolvedDate returns the ResolvedDate field value if set, zero value otherwise.
func (o *RestComment) GetResolvedDate() int64 {
	if o == nil || IsNil(o.ResolvedDate) {
		var ret int64
		return ret
	}
	return *o.ResolvedDate
}

// GetResolvedDateOk returns a tuple with the ResolvedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetResolvedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ResolvedDate) {
		return nil, false
	}
	return o.ResolvedDate, true
}

// HasResolvedDate returns a boolean if a field has been set.
func (o *RestComment) HasResolvedDate() bool {
	if o != nil && !IsNil(o.ResolvedDate) {
		return true
	}

	return false
}

// SetResolvedDate gets a reference to the given int64 and assigns it to the ResolvedDate field.
func (o *RestComment) SetResolvedDate(v int64) {
	o.ResolvedDate = &v
}

// GetResolver returns the Resolver field value if set, zero value otherwise.
func (o *RestComment) GetResolver() RestCommentAuthor {
	if o == nil || IsNil(o.Resolver) {
		var ret RestCommentAuthor
		return ret
	}
	return *o.Resolver
}

// GetResolverOk returns a tuple with the Resolver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetResolverOk() (*RestCommentAuthor, bool) {
	if o == nil || IsNil(o.Resolver) {
		return nil, false
	}
	return o.Resolver, true
}

// HasResolver returns a boolean if a field has been set.
func (o *RestComment) HasResolver() bool {
	if o != nil && !IsNil(o.Resolver) {
		return true
	}

	return false
}

// SetResolver gets a reference to the given RestCommentAuthor and assigns it to the Resolver field.
func (o *RestComment) SetResolver(v RestCommentAuthor) {
	o.Resolver = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *RestComment) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *RestComment) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *RestComment) SetSeverity(v string) {
	o.Severity = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RestComment) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RestComment) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RestComment) SetState(v string) {
	o.State = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *RestComment) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *RestComment) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *RestComment) SetText(v string) {
	o.Text = &v
}

// GetThreadResolved returns the ThreadResolved field value if set, zero value otherwise.
func (o *RestComment) GetThreadResolved() bool {
	if o == nil || IsNil(o.ThreadResolved) {
		var ret bool
		return ret
	}
	return *o.ThreadResolved
}

// GetThreadResolvedOk returns a tuple with the ThreadResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetThreadResolvedOk() (*bool, bool) {
	if o == nil || IsNil(o.ThreadResolved) {
		return nil, false
	}
	return o.ThreadResolved, true
}

// HasThreadResolved returns a boolean if a field has been set.
func (o *RestComment) HasThreadResolved() bool {
	if o != nil && !IsNil(o.ThreadResolved) {
		return true
	}

	return false
}

// SetThreadResolved gets a reference to the given bool and assigns it to the ThreadResolved field.
func (o *RestComment) SetThreadResolved(v bool) {
	o.ThreadResolved = &v
}

// GetThreadResolvedDate returns the ThreadResolvedDate field value if set, zero value otherwise.
func (o *RestComment) GetThreadResolvedDate() int64 {
	if o == nil || IsNil(o.ThreadResolvedDate) {
		var ret int64
		return ret
	}
	return *o.ThreadResolvedDate
}

// GetThreadResolvedDateOk returns a tuple with the ThreadResolvedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetThreadResolvedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ThreadResolvedDate) {
		return nil, false
	}
	return o.ThreadResolvedDate, true
}

// HasThreadResolvedDate returns a boolean if a field has been set.
func (o *RestComment) HasThreadResolvedDate() bool {
	if o != nil && !IsNil(o.ThreadResolvedDate) {
		return true
	}

	return false
}

// SetThreadResolvedDate gets a reference to the given int64 and assigns it to the ThreadResolvedDate field.
func (o *RestComment) SetThreadResolvedDate(v int64) {
	o.ThreadResolvedDate = &v
}

// GetThreadResolver returns the ThreadResolver field value if set, zero value otherwise.
func (o *RestComment) GetThreadResolver() RestCommentAuthor {
	if o == nil || IsNil(o.ThreadResolver) {
		var ret RestCommentAuthor
		return ret
	}
	return *o.ThreadResolver
}

// GetThreadResolverOk returns a tuple with the ThreadResolver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetThreadResolverOk() (*RestCommentAuthor, bool) {
	if o == nil || IsNil(o.ThreadResolver) {
		return nil, false
	}
	return o.ThreadResolver, true
}

// HasThreadResolver returns a boolean if a field has been set.
func (o *RestComment) HasThreadResolver() bool {
	if o != nil && !IsNil(o.ThreadResolver) {
		return true
	}

	return false
}

// SetThreadResolver gets a reference to the given RestCommentAuthor and assigns it to the ThreadResolver field.
func (o *RestComment) SetThreadResolver(v RestCommentAuthor) {
	o.ThreadResolver = &v
}

// GetUpdatedDate returns the UpdatedDate field value if set, zero value otherwise.
func (o *RestComment) GetUpdatedDate() int64 {
	if o == nil || IsNil(o.UpdatedDate) {
		var ret int64
		return ret
	}
	return *o.UpdatedDate
}

// GetUpdatedDateOk returns a tuple with the UpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetUpdatedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedDate) {
		return nil, false
	}
	return o.UpdatedDate, true
}

// HasUpdatedDate returns a boolean if a field has been set.
func (o *RestComment) HasUpdatedDate() bool {
	if o != nil && !IsNil(o.UpdatedDate) {
		return true
	}

	return false
}

// SetUpdatedDate gets a reference to the given int64 and assigns it to the UpdatedDate field.
func (o *RestComment) SetUpdatedDate(v int64) {
	o.UpdatedDate = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RestComment) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestComment) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RestComment) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *RestComment) SetVersion(v int32) {
	o.Version = &v
}

func (o RestComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Anchor) {
		toSerialize["anchor"] = o.Anchor
	}
	if !IsNil(o.Anchored) {
		toSerialize["anchored"] = o.Anchored
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.Html) {
		toSerialize["html"] = o.Html
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Reply) {
		toSerialize["reply"] = o.Reply
	}
	if !IsNil(o.ResolvedDate) {
		toSerialize["resolvedDate"] = o.ResolvedDate
	}
	if !IsNil(o.Resolver) {
		toSerialize["resolver"] = o.Resolver
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.ThreadResolved) {
		toSerialize["threadResolved"] = o.ThreadResolved
	}
	if !IsNil(o.ThreadResolvedDate) {
		toSerialize["threadResolvedDate"] = o.ThreadResolvedDate
	}
	if !IsNil(o.ThreadResolver) {
		toSerialize["threadResolver"] = o.ThreadResolver
	}
	if !IsNil(o.UpdatedDate) {
		toSerialize["updatedDate"] = o.UpdatedDate
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableRestComment struct {
	value *RestComment
	isSet bool
}

func (v NullableRestComment) Get() *RestComment {
	return v.value
}

func (v *NullableRestComment) Set(val *RestComment) {
	v.value = val
	v.isSet = true
}

func (v NullableRestComment) IsSet() bool {
	return v.isSet
}

func (v *NullableRestComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestComment(val *RestComment) *NullableRestComment {
	return &NullableRestComment{value: val, isSet: true}
}

func (v NullableRestComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



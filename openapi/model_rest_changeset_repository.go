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

// checks if the RestChangesetRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestChangesetRepository{}

// RestChangesetRepository struct for RestChangesetRepository
type RestChangesetRepository struct {
	Archived *bool `json:"archived,omitempty"`
	DefaultBranch *string `json:"defaultBranch,omitempty"`
	Description *string `json:"description,omitempty"`
	Forkable *bool `json:"forkable,omitempty"`
	HierarchyId *string `json:"hierarchyId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Name *string `json:"name,omitempty"`
	Origin *RestChangesetRepositoryOrigin `json:"origin,omitempty"`
	Partition *int32 `json:"partition,omitempty"`
	Project *RestChangesetRepositoryOriginProject `json:"project,omitempty"`
	Public *bool `json:"public,omitempty"`
	RelatedLinks map[string]interface{} `json:"relatedLinks,omitempty"`
	ScmId *string `json:"scmId,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Slug *string `json:"slug,omitempty"`
	State *string `json:"state,omitempty"`
	StatusMessage *string `json:"statusMessage,omitempty"`
}

// NewRestChangesetRepository instantiates a new RestChangesetRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestChangesetRepository() *RestChangesetRepository {
	this := RestChangesetRepository{}
	return &this
}

// NewRestChangesetRepositoryWithDefaults instantiates a new RestChangesetRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestChangesetRepositoryWithDefaults() *RestChangesetRepository {
	this := RestChangesetRepository{}
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *RestChangesetRepository) SetArchived(v bool) {
	o.Archived = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetDefaultBranch() string {
	if o == nil || IsNil(o.DefaultBranch) {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetDefaultBranchOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBranch) {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasDefaultBranch() bool {
	if o != nil && !IsNil(o.DefaultBranch) {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *RestChangesetRepository) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RestChangesetRepository) SetDescription(v string) {
	o.Description = &v
}

// GetForkable returns the Forkable field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetForkable() bool {
	if o == nil || IsNil(o.Forkable) {
		var ret bool
		return ret
	}
	return *o.Forkable
}

// GetForkableOk returns a tuple with the Forkable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetForkableOk() (*bool, bool) {
	if o == nil || IsNil(o.Forkable) {
		return nil, false
	}
	return o.Forkable, true
}

// HasForkable returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasForkable() bool {
	if o != nil && !IsNil(o.Forkable) {
		return true
	}

	return false
}

// SetForkable gets a reference to the given bool and assigns it to the Forkable field.
func (o *RestChangesetRepository) SetForkable(v bool) {
	o.Forkable = &v
}

// GetHierarchyId returns the HierarchyId field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetHierarchyId() string {
	if o == nil || IsNil(o.HierarchyId) {
		var ret string
		return ret
	}
	return *o.HierarchyId
}

// GetHierarchyIdOk returns a tuple with the HierarchyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetHierarchyIdOk() (*string, bool) {
	if o == nil || IsNil(o.HierarchyId) {
		return nil, false
	}
	return o.HierarchyId, true
}

// HasHierarchyId returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasHierarchyId() bool {
	if o != nil && !IsNil(o.HierarchyId) {
		return true
	}

	return false
}

// SetHierarchyId gets a reference to the given string and assigns it to the HierarchyId field.
func (o *RestChangesetRepository) SetHierarchyId(v string) {
	o.HierarchyId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RestChangesetRepository) SetId(v int32) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetLinks() map[string]interface{} {
	if o == nil || IsNil(o.Links) {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]interface{}{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *RestChangesetRepository) SetLinks(v map[string]interface{}) {
	o.Links = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RestChangesetRepository) SetName(v string) {
	o.Name = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetOrigin() RestChangesetRepositoryOrigin {
	if o == nil || IsNil(o.Origin) {
		var ret RestChangesetRepositoryOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetOriginOk() (*RestChangesetRepositoryOrigin, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given RestChangesetRepositoryOrigin and assigns it to the Origin field.
func (o *RestChangesetRepository) SetOrigin(v RestChangesetRepositoryOrigin) {
	o.Origin = &v
}

// GetPartition returns the Partition field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetPartition() int32 {
	if o == nil || IsNil(o.Partition) {
		var ret int32
		return ret
	}
	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetPartitionOk() (*int32, bool) {
	if o == nil || IsNil(o.Partition) {
		return nil, false
	}
	return o.Partition, true
}

// HasPartition returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasPartition() bool {
	if o != nil && !IsNil(o.Partition) {
		return true
	}

	return false
}

// SetPartition gets a reference to the given int32 and assigns it to the Partition field.
func (o *RestChangesetRepository) SetPartition(v int32) {
	o.Partition = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetProject() RestChangesetRepositoryOriginProject {
	if o == nil || IsNil(o.Project) {
		var ret RestChangesetRepositoryOriginProject
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetProjectOk() (*RestChangesetRepositoryOriginProject, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given RestChangesetRepositoryOriginProject and assigns it to the Project field.
func (o *RestChangesetRepository) SetProject(v RestChangesetRepositoryOriginProject) {
	o.Project = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *RestChangesetRepository) SetPublic(v bool) {
	o.Public = &v
}

// GetRelatedLinks returns the RelatedLinks field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetRelatedLinks() map[string]interface{} {
	if o == nil || IsNil(o.RelatedLinks) {
		var ret map[string]interface{}
		return ret
	}
	return o.RelatedLinks
}

// GetRelatedLinksOk returns a tuple with the RelatedLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetRelatedLinksOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RelatedLinks) {
		return map[string]interface{}{}, false
	}
	return o.RelatedLinks, true
}

// HasRelatedLinks returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasRelatedLinks() bool {
	if o != nil && !IsNil(o.RelatedLinks) {
		return true
	}

	return false
}

// SetRelatedLinks gets a reference to the given map[string]interface{} and assigns it to the RelatedLinks field.
func (o *RestChangesetRepository) SetRelatedLinks(v map[string]interface{}) {
	o.RelatedLinks = v
}

// GetScmId returns the ScmId field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetScmId() string {
	if o == nil || IsNil(o.ScmId) {
		var ret string
		return ret
	}
	return *o.ScmId
}

// GetScmIdOk returns a tuple with the ScmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetScmIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScmId) {
		return nil, false
	}
	return o.ScmId, true
}

// HasScmId returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasScmId() bool {
	if o != nil && !IsNil(o.ScmId) {
		return true
	}

	return false
}

// SetScmId gets a reference to the given string and assigns it to the ScmId field.
func (o *RestChangesetRepository) SetScmId(v string) {
	o.ScmId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *RestChangesetRepository) SetScope(v string) {
	o.Scope = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *RestChangesetRepository) SetSlug(v string) {
	o.Slug = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RestChangesetRepository) SetState(v string) {
	o.State = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *RestChangesetRepository) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestChangesetRepository) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *RestChangesetRepository) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *RestChangesetRepository) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

func (o RestChangesetRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestChangesetRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.DefaultBranch) {
		toSerialize["defaultBranch"] = o.DefaultBranch
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Forkable) {
		toSerialize["forkable"] = o.Forkable
	}
	if !IsNil(o.HierarchyId) {
		toSerialize["hierarchyId"] = o.HierarchyId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Partition) {
		toSerialize["partition"] = o.Partition
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.RelatedLinks) {
		toSerialize["relatedLinks"] = o.RelatedLinks
	}
	if !IsNil(o.ScmId) {
		toSerialize["scmId"] = o.ScmId
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	return toSerialize, nil
}

type NullableRestChangesetRepository struct {
	value *RestChangesetRepository
	isSet bool
}

func (v NullableRestChangesetRepository) Get() *RestChangesetRepository {
	return v.value
}

func (v *NullableRestChangesetRepository) Set(val *RestChangesetRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableRestChangesetRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableRestChangesetRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestChangesetRepository(val *RestChangesetRepository) *NullableRestChangesetRepository {
	return &NullableRestChangesetRepository{value: val, isSet: true}
}

func (v NullableRestChangesetRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestChangesetRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



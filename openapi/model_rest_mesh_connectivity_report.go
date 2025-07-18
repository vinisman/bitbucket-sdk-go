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

// checks if the RestMeshConnectivityReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestMeshConnectivityReport{}

// RestMeshConnectivityReport struct for RestMeshConnectivityReport
type RestMeshConnectivityReport struct {
	Reports []RestNodeConnectivityReport `json:"reports,omitempty"`
}

// NewRestMeshConnectivityReport instantiates a new RestMeshConnectivityReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestMeshConnectivityReport() *RestMeshConnectivityReport {
	this := RestMeshConnectivityReport{}
	return &this
}

// NewRestMeshConnectivityReportWithDefaults instantiates a new RestMeshConnectivityReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestMeshConnectivityReportWithDefaults() *RestMeshConnectivityReport {
	this := RestMeshConnectivityReport{}
	return &this
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *RestMeshConnectivityReport) GetReports() []RestNodeConnectivityReport {
	if o == nil || IsNil(o.Reports) {
		var ret []RestNodeConnectivityReport
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestMeshConnectivityReport) GetReportsOk() ([]RestNodeConnectivityReport, bool) {
	if o == nil || IsNil(o.Reports) {
		return nil, false
	}
	return o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *RestMeshConnectivityReport) HasReports() bool {
	if o != nil && !IsNil(o.Reports) {
		return true
	}

	return false
}

// SetReports gets a reference to the given []RestNodeConnectivityReport and assigns it to the Reports field.
func (o *RestMeshConnectivityReport) SetReports(v []RestNodeConnectivityReport) {
	o.Reports = v
}

func (o RestMeshConnectivityReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestMeshConnectivityReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reports) {
		toSerialize["reports"] = o.Reports
	}
	return toSerialize, nil
}

type NullableRestMeshConnectivityReport struct {
	value *RestMeshConnectivityReport
	isSet bool
}

func (v NullableRestMeshConnectivityReport) Get() *RestMeshConnectivityReport {
	return v.value
}

func (v *NullableRestMeshConnectivityReport) Set(val *RestMeshConnectivityReport) {
	v.value = val
	v.isSet = true
}

func (v NullableRestMeshConnectivityReport) IsSet() bool {
	return v.isSet
}

func (v *NullableRestMeshConnectivityReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestMeshConnectivityReport(val *RestMeshConnectivityReport) *NullableRestMeshConnectivityReport {
	return &NullableRestMeshConnectivityReport{value: val, isSet: true}
}

func (v NullableRestMeshConnectivityReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestMeshConnectivityReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



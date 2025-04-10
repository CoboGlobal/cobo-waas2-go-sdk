/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"encoding/json"
)

// checks if the MPCProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MPCProject{}

// MPCProject The data for project information.
type MPCProject struct {
	// The project ID.
	ProjectId *string `json:"project_id,omitempty"`
	// The [organization](https://manuals.cobo.com/en/portal/organization/introduction) ID.
	OrgId *string `json:"org_id,omitempty"`
	// The project name.
	Name *string `json:"name,omitempty"`
	// The number of key share holders in the project.
	Participants *int32 `json:"participants,omitempty"`
	// The number of key share holders required to sign an operation in the project.
	Threshold *int32 `json:"threshold,omitempty"`
	// The project's creation time in Unix timestamp format, measured in milliseconds.
	CreatedTimestamp *int64 `json:"created_timestamp,omitempty"`
}

// NewMPCProject instantiates a new MPCProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMPCProject() *MPCProject {
	this := MPCProject{}
	return &this
}

// NewMPCProjectWithDefaults instantiates a new MPCProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMPCProjectWithDefaults() *MPCProject {
	this := MPCProject{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *MPCProject) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MPCProject) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *MPCProject) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *MPCProject) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *MPCProject) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MPCProject) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *MPCProject) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *MPCProject) SetOrgId(v string) {
	o.OrgId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MPCProject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MPCProject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MPCProject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MPCProject) SetName(v string) {
	o.Name = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *MPCProject) GetParticipants() int32 {
	if o == nil || IsNil(o.Participants) {
		var ret int32
		return ret
	}
	return *o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MPCProject) GetParticipantsOk() (*int32, bool) {
	if o == nil || IsNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *MPCProject) HasParticipants() bool {
	if o != nil && !IsNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given int32 and assigns it to the Participants field.
func (o *MPCProject) SetParticipants(v int32) {
	o.Participants = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *MPCProject) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MPCProject) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *MPCProject) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *MPCProject) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *MPCProject) GetCreatedTimestamp() int64 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MPCProject) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *MPCProject) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int64 and assigns it to the CreatedTimestamp field.
func (o *MPCProject) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = &v
}

func (o MPCProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MPCProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Participants) {
		toSerialize["participants"] = o.Participants
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["created_timestamp"] = o.CreatedTimestamp
	}
	return toSerialize, nil
}

type NullableMPCProject struct {
	value *MPCProject
	isSet bool
}

func (v NullableMPCProject) Get() *MPCProject {
	return v.value
}

func (v *NullableMPCProject) Set(val *MPCProject) {
	v.value = val
	v.isSet = true
}

func (v NullableMPCProject) IsSet() bool {
	return v.isSet
}

func (v *NullableMPCProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMPCProject(val *MPCProject) *NullableMPCProject {
	return &NullableMPCProject{value: val, isSet: true}
}

func (v NullableMPCProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMPCProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



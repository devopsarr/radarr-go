/*
Radarr

Radarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// Version struct for Version
type Version struct {
	Major *int32 `json:"major,omitempty"`
	Minor *int32 `json:"minor,omitempty"`
	Build *int32 `json:"build,omitempty"`
	Revision *int32 `json:"revision,omitempty"`
	MajorRevision *int32 `json:"majorRevision,omitempty"`
	MinorRevision *int32 `json:"minorRevision,omitempty"`
}

// NewVersion instantiates a new Version object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersion() *Version {
	this := Version{}
	return &this
}

// NewVersionWithDefaults instantiates a new Version object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionWithDefaults() *Version {
	this := Version{}
	return &this
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *Version) GetMajor() int32 {
	if o == nil || isNil(o.Major) {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetMajorOk() (*int32, bool) {
	if o == nil || isNil(o.Major) {
    return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *Version) HasMajor() bool {
	if o != nil && !isNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *Version) SetMajor(v int32) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *Version) GetMinor() int32 {
	if o == nil || isNil(o.Minor) {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetMinorOk() (*int32, bool) {
	if o == nil || isNil(o.Minor) {
    return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *Version) HasMinor() bool {
	if o != nil && !isNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *Version) SetMinor(v int32) {
	o.Minor = &v
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *Version) GetBuild() int32 {
	if o == nil || isNil(o.Build) {
		var ret int32
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetBuildOk() (*int32, bool) {
	if o == nil || isNil(o.Build) {
    return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *Version) HasBuild() bool {
	if o != nil && !isNil(o.Build) {
		return true
	}

	return false
}

// SetBuild gets a reference to the given int32 and assigns it to the Build field.
func (o *Version) SetBuild(v int32) {
	o.Build = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *Version) GetRevision() int32 {
	if o == nil || isNil(o.Revision) {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetRevisionOk() (*int32, bool) {
	if o == nil || isNil(o.Revision) {
    return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *Version) HasRevision() bool {
	if o != nil && !isNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *Version) SetRevision(v int32) {
	o.Revision = &v
}

// GetMajorRevision returns the MajorRevision field value if set, zero value otherwise.
func (o *Version) GetMajorRevision() int32 {
	if o == nil || isNil(o.MajorRevision) {
		var ret int32
		return ret
	}
	return *o.MajorRevision
}

// GetMajorRevisionOk returns a tuple with the MajorRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetMajorRevisionOk() (*int32, bool) {
	if o == nil || isNil(o.MajorRevision) {
    return nil, false
	}
	return o.MajorRevision, true
}

// HasMajorRevision returns a boolean if a field has been set.
func (o *Version) HasMajorRevision() bool {
	if o != nil && !isNil(o.MajorRevision) {
		return true
	}

	return false
}

// SetMajorRevision gets a reference to the given int32 and assigns it to the MajorRevision field.
func (o *Version) SetMajorRevision(v int32) {
	o.MajorRevision = &v
}

// GetMinorRevision returns the MinorRevision field value if set, zero value otherwise.
func (o *Version) GetMinorRevision() int32 {
	if o == nil || isNil(o.MinorRevision) {
		var ret int32
		return ret
	}
	return *o.MinorRevision
}

// GetMinorRevisionOk returns a tuple with the MinorRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Version) GetMinorRevisionOk() (*int32, bool) {
	if o == nil || isNil(o.MinorRevision) {
    return nil, false
	}
	return o.MinorRevision, true
}

// HasMinorRevision returns a boolean if a field has been set.
func (o *Version) HasMinorRevision() bool {
	if o != nil && !isNil(o.MinorRevision) {
		return true
	}

	return false
}

// SetMinorRevision gets a reference to the given int32 and assigns it to the MinorRevision field.
func (o *Version) SetMinorRevision(v int32) {
	o.MinorRevision = &v
}

func (o Version) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !isNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	if !isNil(o.Build) {
		toSerialize["build"] = o.Build
	}
	if !isNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !isNil(o.MajorRevision) {
		toSerialize["majorRevision"] = o.MajorRevision
	}
	if !isNil(o.MinorRevision) {
		toSerialize["minorRevision"] = o.MinorRevision
	}
	return json.Marshal(toSerialize)
}

type NullableVersion struct {
	value *Version
	isSet bool
}

func (v NullableVersion) Get() *Version {
	return v.value
}

func (v *NullableVersion) Set(val *Version) {
	v.value = val
	v.isSet = true
}

func (v NullableVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersion(val *Version) *NullableVersion {
	return &NullableVersion{value: val, isSet: true}
}

func (v NullableVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



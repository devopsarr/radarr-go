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

// MovieStatisticsResource struct for MovieStatisticsResource
type MovieStatisticsResource struct {
	MovieFileCount *int32 `json:"movieFileCount,omitempty"`
	SizeOnDisk *int64 `json:"sizeOnDisk,omitempty"`
	ReleaseGroups []*string `json:"releaseGroups,omitempty"`
}

// NewMovieStatisticsResource instantiates a new MovieStatisticsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieStatisticsResource() *MovieStatisticsResource {
	this := MovieStatisticsResource{}
	return &this
}

// NewMovieStatisticsResourceWithDefaults instantiates a new MovieStatisticsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieStatisticsResourceWithDefaults() *MovieStatisticsResource {
	this := MovieStatisticsResource{}
	return &this
}

// GetMovieFileCount returns the MovieFileCount field value if set, zero value otherwise.
func (o *MovieStatisticsResource) GetMovieFileCount() int32 {
	if o == nil || isNil(o.MovieFileCount) {
		var ret int32
		return ret
	}
	return *o.MovieFileCount
}

// GetMovieFileCountOk returns a tuple with the MovieFileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieStatisticsResource) GetMovieFileCountOk() (*int32, bool) {
	if o == nil || isNil(o.MovieFileCount) {
    return nil, false
	}
	return o.MovieFileCount, true
}

// HasMovieFileCount returns a boolean if a field has been set.
func (o *MovieStatisticsResource) HasMovieFileCount() bool {
	if o != nil && !isNil(o.MovieFileCount) {
		return true
	}

	return false
}

// SetMovieFileCount gets a reference to the given int32 and assigns it to the MovieFileCount field.
func (o *MovieStatisticsResource) SetMovieFileCount(v int32) {
	o.MovieFileCount = &v
}

// GetSizeOnDisk returns the SizeOnDisk field value if set, zero value otherwise.
func (o *MovieStatisticsResource) GetSizeOnDisk() int64 {
	if o == nil || isNil(o.SizeOnDisk) {
		var ret int64
		return ret
	}
	return *o.SizeOnDisk
}

// GetSizeOnDiskOk returns a tuple with the SizeOnDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieStatisticsResource) GetSizeOnDiskOk() (*int64, bool) {
	if o == nil || isNil(o.SizeOnDisk) {
    return nil, false
	}
	return o.SizeOnDisk, true
}

// HasSizeOnDisk returns a boolean if a field has been set.
func (o *MovieStatisticsResource) HasSizeOnDisk() bool {
	if o != nil && !isNil(o.SizeOnDisk) {
		return true
	}

	return false
}

// SetSizeOnDisk gets a reference to the given int64 and assigns it to the SizeOnDisk field.
func (o *MovieStatisticsResource) SetSizeOnDisk(v int64) {
	o.SizeOnDisk = &v
}

// GetReleaseGroups returns the ReleaseGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieStatisticsResource) GetReleaseGroups() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.ReleaseGroups
}

// GetReleaseGroupsOk returns a tuple with the ReleaseGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieStatisticsResource) GetReleaseGroupsOk() ([]*string, bool) {
	if o == nil || isNil(o.ReleaseGroups) {
    return nil, false
	}
	return o.ReleaseGroups, true
}

// HasReleaseGroups returns a boolean if a field has been set.
func (o *MovieStatisticsResource) HasReleaseGroups() bool {
	if o != nil && isNil(o.ReleaseGroups) {
		return true
	}

	return false
}

// SetReleaseGroups gets a reference to the given []string and assigns it to the ReleaseGroups field.
func (o *MovieStatisticsResource) SetReleaseGroups(v []*string) {
	o.ReleaseGroups = v
}

func (o MovieStatisticsResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MovieFileCount) {
		toSerialize["movieFileCount"] = o.MovieFileCount
	}
	if !isNil(o.SizeOnDisk) {
		toSerialize["sizeOnDisk"] = o.SizeOnDisk
	}
	if o.ReleaseGroups != nil {
		toSerialize["releaseGroups"] = o.ReleaseGroups
	}
	return json.Marshal(toSerialize)
}

type NullableMovieStatisticsResource struct {
	value *MovieStatisticsResource
	isSet bool
}

func (v NullableMovieStatisticsResource) Get() *MovieStatisticsResource {
	return v.value
}

func (v *NullableMovieStatisticsResource) Set(val *MovieStatisticsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieStatisticsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieStatisticsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieStatisticsResource(val *MovieStatisticsResource) *NullableMovieStatisticsResource {
	return &NullableMovieStatisticsResource{value: val, isSet: true}
}

func (v NullableMovieStatisticsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieStatisticsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Radarr

Radarr API docs

API version: v5.18.4.9674
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the RenameMovieResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenameMovieResource{}

// RenameMovieResource struct for RenameMovieResource
type RenameMovieResource struct {
	Id *int32 `json:"id,omitempty"`
	MovieId *int32 `json:"movieId,omitempty"`
	MovieFileId *int32 `json:"movieFileId,omitempty"`
	ExistingPath NullableString `json:"existingPath,omitempty"`
	NewPath NullableString `json:"newPath,omitempty"`
}

// NewRenameMovieResource instantiates a new RenameMovieResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenameMovieResource() *RenameMovieResource {
	this := RenameMovieResource{}
	return &this
}

// NewRenameMovieResourceWithDefaults instantiates a new RenameMovieResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenameMovieResourceWithDefaults() *RenameMovieResource {
	this := RenameMovieResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RenameMovieResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameMovieResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RenameMovieResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RenameMovieResource) SetId(v int32) {
	o.Id = &v
}

// GetMovieId returns the MovieId field value if set, zero value otherwise.
func (o *RenameMovieResource) GetMovieId() int32 {
	if o == nil || IsNil(o.MovieId) {
		var ret int32
		return ret
	}
	return *o.MovieId
}

// GetMovieIdOk returns a tuple with the MovieId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameMovieResource) GetMovieIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MovieId) {
		return nil, false
	}
	return o.MovieId, true
}

// HasMovieId returns a boolean if a field has been set.
func (o *RenameMovieResource) HasMovieId() bool {
	if o != nil && !IsNil(o.MovieId) {
		return true
	}

	return false
}

// SetMovieId gets a reference to the given int32 and assigns it to the MovieId field.
func (o *RenameMovieResource) SetMovieId(v int32) {
	o.MovieId = &v
}

// GetMovieFileId returns the MovieFileId field value if set, zero value otherwise.
func (o *RenameMovieResource) GetMovieFileId() int32 {
	if o == nil || IsNil(o.MovieFileId) {
		var ret int32
		return ret
	}
	return *o.MovieFileId
}

// GetMovieFileIdOk returns a tuple with the MovieFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenameMovieResource) GetMovieFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MovieFileId) {
		return nil, false
	}
	return o.MovieFileId, true
}

// HasMovieFileId returns a boolean if a field has been set.
func (o *RenameMovieResource) HasMovieFileId() bool {
	if o != nil && !IsNil(o.MovieFileId) {
		return true
	}

	return false
}

// SetMovieFileId gets a reference to the given int32 and assigns it to the MovieFileId field.
func (o *RenameMovieResource) SetMovieFileId(v int32) {
	o.MovieFileId = &v
}

// GetExistingPath returns the ExistingPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RenameMovieResource) GetExistingPath() string {
	if o == nil || IsNil(o.ExistingPath.Get()) {
		var ret string
		return ret
	}
	return *o.ExistingPath.Get()
}

// GetExistingPathOk returns a tuple with the ExistingPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RenameMovieResource) GetExistingPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExistingPath.Get(), o.ExistingPath.IsSet()
}

// HasExistingPath returns a boolean if a field has been set.
func (o *RenameMovieResource) HasExistingPath() bool {
	if o != nil && o.ExistingPath.IsSet() {
		return true
	}

	return false
}

// SetExistingPath gets a reference to the given NullableString and assigns it to the ExistingPath field.
func (o *RenameMovieResource) SetExistingPath(v string) {
	o.ExistingPath.Set(&v)
}
// SetExistingPathNil sets the value for ExistingPath to be an explicit nil
func (o *RenameMovieResource) SetExistingPathNil() {
	o.ExistingPath.Set(nil)
}

// UnsetExistingPath ensures that no value is present for ExistingPath, not even an explicit nil
func (o *RenameMovieResource) UnsetExistingPath() {
	o.ExistingPath.Unset()
}

// GetNewPath returns the NewPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RenameMovieResource) GetNewPath() string {
	if o == nil || IsNil(o.NewPath.Get()) {
		var ret string
		return ret
	}
	return *o.NewPath.Get()
}

// GetNewPathOk returns a tuple with the NewPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RenameMovieResource) GetNewPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewPath.Get(), o.NewPath.IsSet()
}

// HasNewPath returns a boolean if a field has been set.
func (o *RenameMovieResource) HasNewPath() bool {
	if o != nil && o.NewPath.IsSet() {
		return true
	}

	return false
}

// SetNewPath gets a reference to the given NullableString and assigns it to the NewPath field.
func (o *RenameMovieResource) SetNewPath(v string) {
	o.NewPath.Set(&v)
}
// SetNewPathNil sets the value for NewPath to be an explicit nil
func (o *RenameMovieResource) SetNewPathNil() {
	o.NewPath.Set(nil)
}

// UnsetNewPath ensures that no value is present for NewPath, not even an explicit nil
func (o *RenameMovieResource) UnsetNewPath() {
	o.NewPath.Unset()
}

func (o RenameMovieResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenameMovieResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MovieId) {
		toSerialize["movieId"] = o.MovieId
	}
	if !IsNil(o.MovieFileId) {
		toSerialize["movieFileId"] = o.MovieFileId
	}
	if o.ExistingPath.IsSet() {
		toSerialize["existingPath"] = o.ExistingPath.Get()
	}
	if o.NewPath.IsSet() {
		toSerialize["newPath"] = o.NewPath.Get()
	}
	return toSerialize, nil
}

type NullableRenameMovieResource struct {
	value *RenameMovieResource
	isSet bool
}

func (v NullableRenameMovieResource) Get() *RenameMovieResource {
	return v.value
}

func (v *NullableRenameMovieResource) Set(val *RenameMovieResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRenameMovieResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRenameMovieResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenameMovieResource(val *RenameMovieResource) *NullableRenameMovieResource {
	return &NullableRenameMovieResource{value: val, isSet: true}
}

func (v NullableRenameMovieResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenameMovieResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



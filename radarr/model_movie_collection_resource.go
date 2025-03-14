/*
Radarr

Radarr API docs

API version: v5.19.3.9730
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the MovieCollectionResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MovieCollectionResource{}

// MovieCollectionResource struct for MovieCollectionResource
type MovieCollectionResource struct {
	Title NullableString `json:"title,omitempty"`
	TmdbId *int32 `json:"tmdbId,omitempty"`
}

// NewMovieCollectionResource instantiates a new MovieCollectionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieCollectionResource() *MovieCollectionResource {
	this := MovieCollectionResource{}
	return &this
}

// NewMovieCollectionResourceWithDefaults instantiates a new MovieCollectionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieCollectionResourceWithDefaults() *MovieCollectionResource {
	this := MovieCollectionResource{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieCollectionResource) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieCollectionResource) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MovieCollectionResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MovieCollectionResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *MovieCollectionResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MovieCollectionResource) UnsetTitle() {
	o.Title.Unset()
}

// GetTmdbId returns the TmdbId field value if set, zero value otherwise.
func (o *MovieCollectionResource) GetTmdbId() int32 {
	if o == nil || IsNil(o.TmdbId) {
		var ret int32
		return ret
	}
	return *o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieCollectionResource) GetTmdbIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TmdbId) {
		return nil, false
	}
	return o.TmdbId, true
}

// HasTmdbId returns a boolean if a field has been set.
func (o *MovieCollectionResource) HasTmdbId() bool {
	if o != nil && !IsNil(o.TmdbId) {
		return true
	}

	return false
}

// SetTmdbId gets a reference to the given int32 and assigns it to the TmdbId field.
func (o *MovieCollectionResource) SetTmdbId(v int32) {
	o.TmdbId = &v
}

func (o MovieCollectionResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovieCollectionResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.TmdbId) {
		toSerialize["tmdbId"] = o.TmdbId
	}
	return toSerialize, nil
}

type NullableMovieCollectionResource struct {
	value *MovieCollectionResource
	isSet bool
}

func (v NullableMovieCollectionResource) Get() *MovieCollectionResource {
	return v.value
}

func (v *NullableMovieCollectionResource) Set(val *MovieCollectionResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieCollectionResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieCollectionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieCollectionResource(val *MovieCollectionResource) *NullableMovieCollectionResource {
	return &NullableMovieCollectionResource{value: val, isSet: true}
}

func (v NullableMovieCollectionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieCollectionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



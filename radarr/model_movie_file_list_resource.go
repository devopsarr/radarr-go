/*
Radarr

Radarr API docs

API version: v5.2.6.8376
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// MovieFileListResource struct for MovieFileListResource
type MovieFileListResource struct {
	MovieFileIds []*int32 `json:"movieFileIds,omitempty"`
	Languages []*Language `json:"languages,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	Edition NullableString `json:"edition,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	SceneName NullableString `json:"sceneName,omitempty"`
	IndexerFlags NullableInt32 `json:"indexerFlags,omitempty"`
}

// NewMovieFileListResource instantiates a new MovieFileListResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieFileListResource() *MovieFileListResource {
	this := MovieFileListResource{}
	return &this
}

// NewMovieFileListResourceWithDefaults instantiates a new MovieFileListResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieFileListResourceWithDefaults() *MovieFileListResource {
	this := MovieFileListResource{}
	return &this
}

// GetMovieFileIds returns the MovieFileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileListResource) GetMovieFileIds() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.MovieFileIds
}

// GetMovieFileIdsOk returns a tuple with the MovieFileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileListResource) GetMovieFileIdsOk() ([]*int32, bool) {
	if o == nil || isNil(o.MovieFileIds) {
    return nil, false
	}
	return o.MovieFileIds, true
}

// HasMovieFileIds returns a boolean if a field has been set.
func (o *MovieFileListResource) HasMovieFileIds() bool {
	if o != nil && isNil(o.MovieFileIds) {
		return true
	}

	return false
}

// SetMovieFileIds gets a reference to the given []int32 and assigns it to the MovieFileIds field.
func (o *MovieFileListResource) SetMovieFileIds(v []*int32) {
	o.MovieFileIds = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileListResource) GetLanguages() []*Language {
	if o == nil {
		var ret []*Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileListResource) GetLanguagesOk() ([]*Language, bool) {
	if o == nil || isNil(o.Languages) {
    return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *MovieFileListResource) HasLanguages() bool {
	if o != nil && isNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *MovieFileListResource) SetLanguages(v []*Language) {
	o.Languages = v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *MovieFileListResource) GetQuality() QualityModel {
	if o == nil || isNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieFileListResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *MovieFileListResource) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *MovieFileListResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetEdition returns the Edition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileListResource) GetEdition() string {
	if o == nil || isNil(o.Edition.Get()) {
		var ret string
		return ret
	}
	return *o.Edition.Get()
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileListResource) GetEditionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Edition.Get(), o.Edition.IsSet()
}

// HasEdition returns a boolean if a field has been set.
func (o *MovieFileListResource) HasEdition() bool {
	if o != nil && o.Edition.IsSet() {
		return true
	}

	return false
}

// SetEdition gets a reference to the given NullableString and assigns it to the Edition field.
func (o *MovieFileListResource) SetEdition(v string) {
	o.Edition.Set(&v)
}
// SetEditionNil sets the value for Edition to be an explicit nil
func (o *MovieFileListResource) SetEditionNil() {
	o.Edition.Set(nil)
}

// UnsetEdition ensures that no value is present for Edition, not even an explicit nil
func (o *MovieFileListResource) UnsetEdition() {
	o.Edition.Unset()
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileListResource) GetReleaseGroup() string {
	if o == nil || isNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileListResource) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *MovieFileListResource) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *MovieFileListResource) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *MovieFileListResource) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *MovieFileListResource) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetSceneName returns the SceneName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileListResource) GetSceneName() string {
	if o == nil || isNil(o.SceneName.Get()) {
		var ret string
		return ret
	}
	return *o.SceneName.Get()
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileListResource) GetSceneNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SceneName.Get(), o.SceneName.IsSet()
}

// HasSceneName returns a boolean if a field has been set.
func (o *MovieFileListResource) HasSceneName() bool {
	if o != nil && o.SceneName.IsSet() {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given NullableString and assigns it to the SceneName field.
func (o *MovieFileListResource) SetSceneName(v string) {
	o.SceneName.Set(&v)
}
// SetSceneNameNil sets the value for SceneName to be an explicit nil
func (o *MovieFileListResource) SetSceneNameNil() {
	o.SceneName.Set(nil)
}

// UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
func (o *MovieFileListResource) UnsetSceneName() {
	o.SceneName.Unset()
}

// GetIndexerFlags returns the IndexerFlags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieFileListResource) GetIndexerFlags() int32 {
	if o == nil || isNil(o.IndexerFlags.Get()) {
		var ret int32
		return ret
	}
	return *o.IndexerFlags.Get()
}

// GetIndexerFlagsOk returns a tuple with the IndexerFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieFileListResource) GetIndexerFlagsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.IndexerFlags.Get(), o.IndexerFlags.IsSet()
}

// HasIndexerFlags returns a boolean if a field has been set.
func (o *MovieFileListResource) HasIndexerFlags() bool {
	if o != nil && o.IndexerFlags.IsSet() {
		return true
	}

	return false
}

// SetIndexerFlags gets a reference to the given NullableInt32 and assigns it to the IndexerFlags field.
func (o *MovieFileListResource) SetIndexerFlags(v int32) {
	o.IndexerFlags.Set(&v)
}
// SetIndexerFlagsNil sets the value for IndexerFlags to be an explicit nil
func (o *MovieFileListResource) SetIndexerFlagsNil() {
	o.IndexerFlags.Set(nil)
}

// UnsetIndexerFlags ensures that no value is present for IndexerFlags, not even an explicit nil
func (o *MovieFileListResource) UnsetIndexerFlags() {
	o.IndexerFlags.Unset()
}

func (o MovieFileListResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MovieFileIds != nil {
		toSerialize["movieFileIds"] = o.MovieFileIds
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.Edition.IsSet() {
		toSerialize["edition"] = o.Edition.Get()
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if o.SceneName.IsSet() {
		toSerialize["sceneName"] = o.SceneName.Get()
	}
	if o.IndexerFlags.IsSet() {
		toSerialize["indexerFlags"] = o.IndexerFlags.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMovieFileListResource struct {
	value *MovieFileListResource
	isSet bool
}

func (v NullableMovieFileListResource) Get() *MovieFileListResource {
	return v.value
}

func (v *NullableMovieFileListResource) Set(val *MovieFileListResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieFileListResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieFileListResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieFileListResource(val *MovieFileListResource) *NullableMovieFileListResource {
	return &NullableMovieFileListResource{value: val, isSet: true}
}

func (v NullableMovieFileListResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieFileListResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



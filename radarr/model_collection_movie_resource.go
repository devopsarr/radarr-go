/*
Radarr

Radarr API docs

API version: v5.15.1.9463
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the CollectionMovieResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionMovieResource{}

// CollectionMovieResource struct for CollectionMovieResource
type CollectionMovieResource struct {
	TmdbId *int32 `json:"tmdbId,omitempty"`
	ImdbId NullableString `json:"imdbId,omitempty"`
	Title NullableString `json:"title,omitempty"`
	CleanTitle NullableString `json:"cleanTitle,omitempty"`
	SortTitle NullableString `json:"sortTitle,omitempty"`
	Status *MovieStatusType `json:"status,omitempty"`
	Overview NullableString `json:"overview,omitempty"`
	Runtime *int32 `json:"runtime,omitempty"`
	Images []MediaCover `json:"images,omitempty"`
	Year *int32 `json:"year,omitempty"`
	Ratings *Ratings `json:"ratings,omitempty"`
	Genres []string `json:"genres,omitempty"`
	Folder NullableString `json:"folder,omitempty"`
	IsExisting *bool `json:"isExisting,omitempty"`
	IsExcluded *bool `json:"isExcluded,omitempty"`
}

// NewCollectionMovieResource instantiates a new CollectionMovieResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionMovieResource() *CollectionMovieResource {
	this := CollectionMovieResource{}
	return &this
}

// NewCollectionMovieResourceWithDefaults instantiates a new CollectionMovieResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionMovieResourceWithDefaults() *CollectionMovieResource {
	this := CollectionMovieResource{}
	return &this
}

// GetTmdbId returns the TmdbId field value if set, zero value otherwise.
func (o *CollectionMovieResource) GetTmdbId() int32 {
	if o == nil || IsNil(o.TmdbId) {
		var ret int32
		return ret
	}
	return *o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMovieResource) GetTmdbIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TmdbId) {
		return nil, false
	}
	return o.TmdbId, true
}

// HasTmdbId returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasTmdbId() bool {
	if o != nil && !IsNil(o.TmdbId) {
		return true
	}

	return false
}

// SetTmdbId gets a reference to the given int32 and assigns it to the TmdbId field.
func (o *CollectionMovieResource) SetTmdbId(v int32) {
	o.TmdbId = &v
}

// GetImdbId returns the ImdbId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionMovieResource) GetImdbId() string {
	if o == nil || IsNil(o.ImdbId.Get()) {
		var ret string
		return ret
	}
	return *o.ImdbId.Get()
}

// GetImdbIdOk returns a tuple with the ImdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionMovieResource) GetImdbIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImdbId.Get(), o.ImdbId.IsSet()
}

// HasImdbId returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasImdbId() bool {
	if o != nil && o.ImdbId.IsSet() {
		return true
	}

	return false
}

// SetImdbId gets a reference to the given NullableString and assigns it to the ImdbId field.
func (o *CollectionMovieResource) SetImdbId(v string) {
	o.ImdbId.Set(&v)
}
// SetImdbIdNil sets the value for ImdbId to be an explicit nil
func (o *CollectionMovieResource) SetImdbIdNil() {
	o.ImdbId.Set(nil)
}

// UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
func (o *CollectionMovieResource) UnsetImdbId() {
	o.ImdbId.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionMovieResource) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionMovieResource) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *CollectionMovieResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *CollectionMovieResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *CollectionMovieResource) UnsetTitle() {
	o.Title.Unset()
}

// GetCleanTitle returns the CleanTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionMovieResource) GetCleanTitle() string {
	if o == nil || IsNil(o.CleanTitle.Get()) {
		var ret string
		return ret
	}
	return *o.CleanTitle.Get()
}

// GetCleanTitleOk returns a tuple with the CleanTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionMovieResource) GetCleanTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CleanTitle.Get(), o.CleanTitle.IsSet()
}

// HasCleanTitle returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasCleanTitle() bool {
	if o != nil && o.CleanTitle.IsSet() {
		return true
	}

	return false
}

// SetCleanTitle gets a reference to the given NullableString and assigns it to the CleanTitle field.
func (o *CollectionMovieResource) SetCleanTitle(v string) {
	o.CleanTitle.Set(&v)
}
// SetCleanTitleNil sets the value for CleanTitle to be an explicit nil
func (o *CollectionMovieResource) SetCleanTitleNil() {
	o.CleanTitle.Set(nil)
}

// UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
func (o *CollectionMovieResource) UnsetCleanTitle() {
	o.CleanTitle.Unset()
}

// GetSortTitle returns the SortTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionMovieResource) GetSortTitle() string {
	if o == nil || IsNil(o.SortTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SortTitle.Get()
}

// GetSortTitleOk returns a tuple with the SortTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionMovieResource) GetSortTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortTitle.Get(), o.SortTitle.IsSet()
}

// HasSortTitle returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasSortTitle() bool {
	if o != nil && o.SortTitle.IsSet() {
		return true
	}

	return false
}

// SetSortTitle gets a reference to the given NullableString and assigns it to the SortTitle field.
func (o *CollectionMovieResource) SetSortTitle(v string) {
	o.SortTitle.Set(&v)
}
// SetSortTitleNil sets the value for SortTitle to be an explicit nil
func (o *CollectionMovieResource) SetSortTitleNil() {
	o.SortTitle.Set(nil)
}

// UnsetSortTitle ensures that no value is present for SortTitle, not even an explicit nil
func (o *CollectionMovieResource) UnsetSortTitle() {
	o.SortTitle.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CollectionMovieResource) GetStatus() MovieStatusType {
	if o == nil || IsNil(o.Status) {
		var ret MovieStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMovieResource) GetStatusOk() (*MovieStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MovieStatusType and assigns it to the Status field.
func (o *CollectionMovieResource) SetStatus(v MovieStatusType) {
	o.Status = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionMovieResource) GetOverview() string {
	if o == nil || IsNil(o.Overview.Get()) {
		var ret string
		return ret
	}
	return *o.Overview.Get()
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionMovieResource) GetOverviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Overview.Get(), o.Overview.IsSet()
}

// HasOverview returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasOverview() bool {
	if o != nil && o.Overview.IsSet() {
		return true
	}

	return false
}

// SetOverview gets a reference to the given NullableString and assigns it to the Overview field.
func (o *CollectionMovieResource) SetOverview(v string) {
	o.Overview.Set(&v)
}
// SetOverviewNil sets the value for Overview to be an explicit nil
func (o *CollectionMovieResource) SetOverviewNil() {
	o.Overview.Set(nil)
}

// UnsetOverview ensures that no value is present for Overview, not even an explicit nil
func (o *CollectionMovieResource) UnsetOverview() {
	o.Overview.Unset()
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *CollectionMovieResource) GetRuntime() int32 {
	if o == nil || IsNil(o.Runtime) {
		var ret int32
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMovieResource) GetRuntimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Runtime) {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasRuntime() bool {
	if o != nil && !IsNil(o.Runtime) {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given int32 and assigns it to the Runtime field.
func (o *CollectionMovieResource) SetRuntime(v int32) {
	o.Runtime = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionMovieResource) GetImages() []MediaCover {
	if o == nil {
		var ret []MediaCover
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionMovieResource) GetImagesOk() ([]MediaCover, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []MediaCover and assigns it to the Images field.
func (o *CollectionMovieResource) SetImages(v []MediaCover) {
	o.Images = v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *CollectionMovieResource) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMovieResource) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *CollectionMovieResource) SetYear(v int32) {
	o.Year = &v
}

// GetRatings returns the Ratings field value if set, zero value otherwise.
func (o *CollectionMovieResource) GetRatings() Ratings {
	if o == nil || IsNil(o.Ratings) {
		var ret Ratings
		return ret
	}
	return *o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMovieResource) GetRatingsOk() (*Ratings, bool) {
	if o == nil || IsNil(o.Ratings) {
		return nil, false
	}
	return o.Ratings, true
}

// HasRatings returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasRatings() bool {
	if o != nil && !IsNil(o.Ratings) {
		return true
	}

	return false
}

// SetRatings gets a reference to the given Ratings and assigns it to the Ratings field.
func (o *CollectionMovieResource) SetRatings(v Ratings) {
	o.Ratings = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionMovieResource) GetGenres() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionMovieResource) GetGenresOk() ([]string, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *CollectionMovieResource) SetGenres(v []string) {
	o.Genres = v
}

// GetFolder returns the Folder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionMovieResource) GetFolder() string {
	if o == nil || IsNil(o.Folder.Get()) {
		var ret string
		return ret
	}
	return *o.Folder.Get()
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionMovieResource) GetFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Folder.Get(), o.Folder.IsSet()
}

// HasFolder returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasFolder() bool {
	if o != nil && o.Folder.IsSet() {
		return true
	}

	return false
}

// SetFolder gets a reference to the given NullableString and assigns it to the Folder field.
func (o *CollectionMovieResource) SetFolder(v string) {
	o.Folder.Set(&v)
}
// SetFolderNil sets the value for Folder to be an explicit nil
func (o *CollectionMovieResource) SetFolderNil() {
	o.Folder.Set(nil)
}

// UnsetFolder ensures that no value is present for Folder, not even an explicit nil
func (o *CollectionMovieResource) UnsetFolder() {
	o.Folder.Unset()
}

// GetIsExisting returns the IsExisting field value if set, zero value otherwise.
func (o *CollectionMovieResource) GetIsExisting() bool {
	if o == nil || IsNil(o.IsExisting) {
		var ret bool
		return ret
	}
	return *o.IsExisting
}

// GetIsExistingOk returns a tuple with the IsExisting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMovieResource) GetIsExistingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExisting) {
		return nil, false
	}
	return o.IsExisting, true
}

// HasIsExisting returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasIsExisting() bool {
	if o != nil && !IsNil(o.IsExisting) {
		return true
	}

	return false
}

// SetIsExisting gets a reference to the given bool and assigns it to the IsExisting field.
func (o *CollectionMovieResource) SetIsExisting(v bool) {
	o.IsExisting = &v
}

// GetIsExcluded returns the IsExcluded field value if set, zero value otherwise.
func (o *CollectionMovieResource) GetIsExcluded() bool {
	if o == nil || IsNil(o.IsExcluded) {
		var ret bool
		return ret
	}
	return *o.IsExcluded
}

// GetIsExcludedOk returns a tuple with the IsExcluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMovieResource) GetIsExcludedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExcluded) {
		return nil, false
	}
	return o.IsExcluded, true
}

// HasIsExcluded returns a boolean if a field has been set.
func (o *CollectionMovieResource) HasIsExcluded() bool {
	if o != nil && !IsNil(o.IsExcluded) {
		return true
	}

	return false
}

// SetIsExcluded gets a reference to the given bool and assigns it to the IsExcluded field.
func (o *CollectionMovieResource) SetIsExcluded(v bool) {
	o.IsExcluded = &v
}

func (o CollectionMovieResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionMovieResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TmdbId) {
		toSerialize["tmdbId"] = o.TmdbId
	}
	if o.ImdbId.IsSet() {
		toSerialize["imdbId"] = o.ImdbId.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.CleanTitle.IsSet() {
		toSerialize["cleanTitle"] = o.CleanTitle.Get()
	}
	if o.SortTitle.IsSet() {
		toSerialize["sortTitle"] = o.SortTitle.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Overview.IsSet() {
		toSerialize["overview"] = o.Overview.Get()
	}
	if !IsNil(o.Runtime) {
		toSerialize["runtime"] = o.Runtime
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Ratings) {
		toSerialize["ratings"] = o.Ratings
	}
	if o.Genres != nil {
		toSerialize["genres"] = o.Genres
	}
	if o.Folder.IsSet() {
		toSerialize["folder"] = o.Folder.Get()
	}
	if !IsNil(o.IsExisting) {
		toSerialize["isExisting"] = o.IsExisting
	}
	if !IsNil(o.IsExcluded) {
		toSerialize["isExcluded"] = o.IsExcluded
	}
	return toSerialize, nil
}

type NullableCollectionMovieResource struct {
	value *CollectionMovieResource
	isSet bool
}

func (v NullableCollectionMovieResource) Get() *CollectionMovieResource {
	return v.value
}

func (v *NullableCollectionMovieResource) Set(val *CollectionMovieResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionMovieResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionMovieResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionMovieResource(val *CollectionMovieResource) *NullableCollectionMovieResource {
	return &NullableCollectionMovieResource{value: val, isSet: true}
}

func (v NullableCollectionMovieResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionMovieResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



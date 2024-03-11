/*
Radarr

Radarr API docs

API version: v5.3.6.8612
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the CollectionResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResource{}

// CollectionResource struct for CollectionResource
type CollectionResource struct {
	Id *int32 `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	SortTitle NullableString `json:"sortTitle,omitempty"`
	TmdbId *int32 `json:"tmdbId,omitempty"`
	Images []MediaCover `json:"images,omitempty"`
	Overview NullableString `json:"overview,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
	RootFolderPath NullableString `json:"rootFolderPath,omitempty"`
	QualityProfileId *int32 `json:"qualityProfileId,omitempty"`
	SearchOnAdd *bool `json:"searchOnAdd,omitempty"`
	MinimumAvailability *MovieStatusType `json:"minimumAvailability,omitempty"`
	Movies []CollectionMovieResource `json:"movies,omitempty"`
	MissingMovies *int32 `json:"missingMovies,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
}

// NewCollectionResource instantiates a new CollectionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResource() *CollectionResource {
	this := CollectionResource{}
	return &this
}

// NewCollectionResourceWithDefaults instantiates a new CollectionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResourceWithDefaults() *CollectionResource {
	this := CollectionResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CollectionResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CollectionResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CollectionResource) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionResource) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionResource) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *CollectionResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *CollectionResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *CollectionResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *CollectionResource) UnsetTitle() {
	o.Title.Unset()
}

// GetSortTitle returns the SortTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionResource) GetSortTitle() string {
	if o == nil || IsNil(o.SortTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SortTitle.Get()
}

// GetSortTitleOk returns a tuple with the SortTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionResource) GetSortTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortTitle.Get(), o.SortTitle.IsSet()
}

// HasSortTitle returns a boolean if a field has been set.
func (o *CollectionResource) HasSortTitle() bool {
	if o != nil && o.SortTitle.IsSet() {
		return true
	}

	return false
}

// SetSortTitle gets a reference to the given NullableString and assigns it to the SortTitle field.
func (o *CollectionResource) SetSortTitle(v string) {
	o.SortTitle.Set(&v)
}
// SetSortTitleNil sets the value for SortTitle to be an explicit nil
func (o *CollectionResource) SetSortTitleNil() {
	o.SortTitle.Set(nil)
}

// UnsetSortTitle ensures that no value is present for SortTitle, not even an explicit nil
func (o *CollectionResource) UnsetSortTitle() {
	o.SortTitle.Unset()
}

// GetTmdbId returns the TmdbId field value if set, zero value otherwise.
func (o *CollectionResource) GetTmdbId() int32 {
	if o == nil || IsNil(o.TmdbId) {
		var ret int32
		return ret
	}
	return *o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResource) GetTmdbIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TmdbId) {
		return nil, false
	}
	return o.TmdbId, true
}

// HasTmdbId returns a boolean if a field has been set.
func (o *CollectionResource) HasTmdbId() bool {
	if o != nil && !IsNil(o.TmdbId) {
		return true
	}

	return false
}

// SetTmdbId gets a reference to the given int32 and assigns it to the TmdbId field.
func (o *CollectionResource) SetTmdbId(v int32) {
	o.TmdbId = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionResource) GetImages() []MediaCover {
	if o == nil {
		var ret []MediaCover
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionResource) GetImagesOk() ([]MediaCover, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *CollectionResource) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []MediaCover and assigns it to the Images field.
func (o *CollectionResource) SetImages(v []MediaCover) {
	o.Images = v
}

// GetOverview returns the Overview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionResource) GetOverview() string {
	if o == nil || IsNil(o.Overview.Get()) {
		var ret string
		return ret
	}
	return *o.Overview.Get()
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionResource) GetOverviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Overview.Get(), o.Overview.IsSet()
}

// HasOverview returns a boolean if a field has been set.
func (o *CollectionResource) HasOverview() bool {
	if o != nil && o.Overview.IsSet() {
		return true
	}

	return false
}

// SetOverview gets a reference to the given NullableString and assigns it to the Overview field.
func (o *CollectionResource) SetOverview(v string) {
	o.Overview.Set(&v)
}
// SetOverviewNil sets the value for Overview to be an explicit nil
func (o *CollectionResource) SetOverviewNil() {
	o.Overview.Set(nil)
}

// UnsetOverview ensures that no value is present for Overview, not even an explicit nil
func (o *CollectionResource) UnsetOverview() {
	o.Overview.Unset()
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *CollectionResource) GetMonitored() bool {
	if o == nil || IsNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResource) GetMonitoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Monitored) {
		return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *CollectionResource) HasMonitored() bool {
	if o != nil && !IsNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *CollectionResource) SetMonitored(v bool) {
	o.Monitored = &v
}

// GetRootFolderPath returns the RootFolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionResource) GetRootFolderPath() string {
	if o == nil || IsNil(o.RootFolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.RootFolderPath.Get()
}

// GetRootFolderPathOk returns a tuple with the RootFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionResource) GetRootFolderPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootFolderPath.Get(), o.RootFolderPath.IsSet()
}

// HasRootFolderPath returns a boolean if a field has been set.
func (o *CollectionResource) HasRootFolderPath() bool {
	if o != nil && o.RootFolderPath.IsSet() {
		return true
	}

	return false
}

// SetRootFolderPath gets a reference to the given NullableString and assigns it to the RootFolderPath field.
func (o *CollectionResource) SetRootFolderPath(v string) {
	o.RootFolderPath.Set(&v)
}
// SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil
func (o *CollectionResource) SetRootFolderPathNil() {
	o.RootFolderPath.Set(nil)
}

// UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
func (o *CollectionResource) UnsetRootFolderPath() {
	o.RootFolderPath.Unset()
}

// GetQualityProfileId returns the QualityProfileId field value if set, zero value otherwise.
func (o *CollectionResource) GetQualityProfileId() int32 {
	if o == nil || IsNil(o.QualityProfileId) {
		var ret int32
		return ret
	}
	return *o.QualityProfileId
}

// GetQualityProfileIdOk returns a tuple with the QualityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResource) GetQualityProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.QualityProfileId) {
		return nil, false
	}
	return o.QualityProfileId, true
}

// HasQualityProfileId returns a boolean if a field has been set.
func (o *CollectionResource) HasQualityProfileId() bool {
	if o != nil && !IsNil(o.QualityProfileId) {
		return true
	}

	return false
}

// SetQualityProfileId gets a reference to the given int32 and assigns it to the QualityProfileId field.
func (o *CollectionResource) SetQualityProfileId(v int32) {
	o.QualityProfileId = &v
}

// GetSearchOnAdd returns the SearchOnAdd field value if set, zero value otherwise.
func (o *CollectionResource) GetSearchOnAdd() bool {
	if o == nil || IsNil(o.SearchOnAdd) {
		var ret bool
		return ret
	}
	return *o.SearchOnAdd
}

// GetSearchOnAddOk returns a tuple with the SearchOnAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResource) GetSearchOnAddOk() (*bool, bool) {
	if o == nil || IsNil(o.SearchOnAdd) {
		return nil, false
	}
	return o.SearchOnAdd, true
}

// HasSearchOnAdd returns a boolean if a field has been set.
func (o *CollectionResource) HasSearchOnAdd() bool {
	if o != nil && !IsNil(o.SearchOnAdd) {
		return true
	}

	return false
}

// SetSearchOnAdd gets a reference to the given bool and assigns it to the SearchOnAdd field.
func (o *CollectionResource) SetSearchOnAdd(v bool) {
	o.SearchOnAdd = &v
}

// GetMinimumAvailability returns the MinimumAvailability field value if set, zero value otherwise.
func (o *CollectionResource) GetMinimumAvailability() MovieStatusType {
	if o == nil || IsNil(o.MinimumAvailability) {
		var ret MovieStatusType
		return ret
	}
	return *o.MinimumAvailability
}

// GetMinimumAvailabilityOk returns a tuple with the MinimumAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResource) GetMinimumAvailabilityOk() (*MovieStatusType, bool) {
	if o == nil || IsNil(o.MinimumAvailability) {
		return nil, false
	}
	return o.MinimumAvailability, true
}

// HasMinimumAvailability returns a boolean if a field has been set.
func (o *CollectionResource) HasMinimumAvailability() bool {
	if o != nil && !IsNil(o.MinimumAvailability) {
		return true
	}

	return false
}

// SetMinimumAvailability gets a reference to the given MovieStatusType and assigns it to the MinimumAvailability field.
func (o *CollectionResource) SetMinimumAvailability(v MovieStatusType) {
	o.MinimumAvailability = &v
}

// GetMovies returns the Movies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionResource) GetMovies() []CollectionMovieResource {
	if o == nil {
		var ret []CollectionMovieResource
		return ret
	}
	return o.Movies
}

// GetMoviesOk returns a tuple with the Movies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionResource) GetMoviesOk() ([]CollectionMovieResource, bool) {
	if o == nil || IsNil(o.Movies) {
		return nil, false
	}
	return o.Movies, true
}

// HasMovies returns a boolean if a field has been set.
func (o *CollectionResource) HasMovies() bool {
	if o != nil && !IsNil(o.Movies) {
		return true
	}

	return false
}

// SetMovies gets a reference to the given []CollectionMovieResource and assigns it to the Movies field.
func (o *CollectionResource) SetMovies(v []CollectionMovieResource) {
	o.Movies = v
}

// GetMissingMovies returns the MissingMovies field value if set, zero value otherwise.
func (o *CollectionResource) GetMissingMovies() int32 {
	if o == nil || IsNil(o.MissingMovies) {
		var ret int32
		return ret
	}
	return *o.MissingMovies
}

// GetMissingMoviesOk returns a tuple with the MissingMovies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResource) GetMissingMoviesOk() (*int32, bool) {
	if o == nil || IsNil(o.MissingMovies) {
		return nil, false
	}
	return o.MissingMovies, true
}

// HasMissingMovies returns a boolean if a field has been set.
func (o *CollectionResource) HasMissingMovies() bool {
	if o != nil && !IsNil(o.MissingMovies) {
		return true
	}

	return false
}

// SetMissingMovies gets a reference to the given int32 and assigns it to the MissingMovies field.
func (o *CollectionResource) SetMissingMovies(v int32) {
	o.MissingMovies = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectionResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CollectionResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *CollectionResource) SetTags(v []int32) {
	o.Tags = v
}

func (o CollectionResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.SortTitle.IsSet() {
		toSerialize["sortTitle"] = o.SortTitle.Get()
	}
	if !IsNil(o.TmdbId) {
		toSerialize["tmdbId"] = o.TmdbId
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.Overview.IsSet() {
		toSerialize["overview"] = o.Overview.Get()
	}
	if !IsNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	if o.RootFolderPath.IsSet() {
		toSerialize["rootFolderPath"] = o.RootFolderPath.Get()
	}
	if !IsNil(o.QualityProfileId) {
		toSerialize["qualityProfileId"] = o.QualityProfileId
	}
	if !IsNil(o.SearchOnAdd) {
		toSerialize["searchOnAdd"] = o.SearchOnAdd
	}
	if !IsNil(o.MinimumAvailability) {
		toSerialize["minimumAvailability"] = o.MinimumAvailability
	}
	if o.Movies != nil {
		toSerialize["movies"] = o.Movies
	}
	if !IsNil(o.MissingMovies) {
		toSerialize["missingMovies"] = o.MissingMovies
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableCollectionResource struct {
	value *CollectionResource
	isSet bool
}

func (v NullableCollectionResource) Get() *CollectionResource {
	return v.value
}

func (v *NullableCollectionResource) Set(val *CollectionResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResource(val *CollectionResource) *NullableCollectionResource {
	return &NullableCollectionResource{value: val, isSet: true}
}

func (v NullableCollectionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ManualImportResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualImportResource{}

// ManualImportResource struct for ManualImportResource
type ManualImportResource struct {
	Id *int32 `json:"id,omitempty"`
	Path NullableString `json:"path,omitempty"`
	RelativePath NullableString `json:"relativePath,omitempty"`
	FolderName NullableString `json:"folderName,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Movie *MovieResource `json:"movie,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	Languages []Language `json:"languages,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	QualityWeight *int32 `json:"qualityWeight,omitempty"`
	DownloadId NullableString `json:"downloadId,omitempty"`
	CustomFormats []CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore *int32 `json:"customFormatScore,omitempty"`
	Rejections []Rejection `json:"rejections,omitempty"`
}

// NewManualImportResource instantiates a new ManualImportResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualImportResource() *ManualImportResource {
	this := ManualImportResource{}
	return &this
}

// NewManualImportResourceWithDefaults instantiates a new ManualImportResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualImportResourceWithDefaults() *ManualImportResource {
	this := ManualImportResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManualImportResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManualImportResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManualImportResource) SetId(v int32) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *ManualImportResource) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *ManualImportResource) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *ManualImportResource) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *ManualImportResource) UnsetPath() {
	o.Path.Unset()
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath.Get()) {
		var ret string
		return ret
	}
	return *o.RelativePath.Get()
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelativePath.Get(), o.RelativePath.IsSet()
}

// HasRelativePath returns a boolean if a field has been set.
func (o *ManualImportResource) HasRelativePath() bool {
	if o != nil && o.RelativePath.IsSet() {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given NullableString and assigns it to the RelativePath field.
func (o *ManualImportResource) SetRelativePath(v string) {
	o.RelativePath.Set(&v)
}
// SetRelativePathNil sets the value for RelativePath to be an explicit nil
func (o *ManualImportResource) SetRelativePathNil() {
	o.RelativePath.Set(nil)
}

// UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
func (o *ManualImportResource) UnsetRelativePath() {
	o.RelativePath.Unset()
}

// GetFolderName returns the FolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetFolderName() string {
	if o == nil || IsNil(o.FolderName.Get()) {
		var ret string
		return ret
	}
	return *o.FolderName.Get()
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetFolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderName.Get(), o.FolderName.IsSet()
}

// HasFolderName returns a boolean if a field has been set.
func (o *ManualImportResource) HasFolderName() bool {
	if o != nil && o.FolderName.IsSet() {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given NullableString and assigns it to the FolderName field.
func (o *ManualImportResource) SetFolderName(v string) {
	o.FolderName.Set(&v)
}
// SetFolderNameNil sets the value for FolderName to be an explicit nil
func (o *ManualImportResource) SetFolderNameNil() {
	o.FolderName.Set(nil)
}

// UnsetFolderName ensures that no value is present for FolderName, not even an explicit nil
func (o *ManualImportResource) UnsetFolderName() {
	o.FolderName.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ManualImportResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ManualImportResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ManualImportResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ManualImportResource) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ManualImportResource) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ManualImportResource) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ManualImportResource) SetSize(v int64) {
	o.Size = &v
}

// GetMovie returns the Movie field value if set, zero value otherwise.
func (o *ManualImportResource) GetMovie() MovieResource {
	if o == nil || IsNil(o.Movie) {
		var ret MovieResource
		return ret
	}
	return *o.Movie
}

// GetMovieOk returns a tuple with the Movie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetMovieOk() (*MovieResource, bool) {
	if o == nil || IsNil(o.Movie) {
		return nil, false
	}
	return o.Movie, true
}

// HasMovie returns a boolean if a field has been set.
func (o *ManualImportResource) HasMovie() bool {
	if o != nil && !IsNil(o.Movie) {
		return true
	}

	return false
}

// SetMovie gets a reference to the given MovieResource and assigns it to the Movie field.
func (o *ManualImportResource) SetMovie(v MovieResource) {
	o.Movie = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ManualImportResource) GetQuality() QualityModel {
	if o == nil || IsNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetQualityOk() (*QualityModel, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ManualImportResource) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *ManualImportResource) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetLanguages() []Language {
	if o == nil {
		var ret []Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetLanguagesOk() ([]Language, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ManualImportResource) HasLanguages() bool {
	if o != nil && IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *ManualImportResource) SetLanguages(v []Language) {
	o.Languages = v
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetReleaseGroup() string {
	if o == nil || IsNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *ManualImportResource) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *ManualImportResource) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *ManualImportResource) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *ManualImportResource) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetQualityWeight returns the QualityWeight field value if set, zero value otherwise.
func (o *ManualImportResource) GetQualityWeight() int32 {
	if o == nil || IsNil(o.QualityWeight) {
		var ret int32
		return ret
	}
	return *o.QualityWeight
}

// GetQualityWeightOk returns a tuple with the QualityWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetQualityWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.QualityWeight) {
		return nil, false
	}
	return o.QualityWeight, true
}

// HasQualityWeight returns a boolean if a field has been set.
func (o *ManualImportResource) HasQualityWeight() bool {
	if o != nil && !IsNil(o.QualityWeight) {
		return true
	}

	return false
}

// SetQualityWeight gets a reference to the given int32 and assigns it to the QualityWeight field.
func (o *ManualImportResource) SetQualityWeight(v int32) {
	o.QualityWeight = &v
}

// GetDownloadId returns the DownloadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetDownloadId() string {
	if o == nil || IsNil(o.DownloadId.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadId.Get()
}

// GetDownloadIdOk returns a tuple with the DownloadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetDownloadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadId.Get(), o.DownloadId.IsSet()
}

// HasDownloadId returns a boolean if a field has been set.
func (o *ManualImportResource) HasDownloadId() bool {
	if o != nil && o.DownloadId.IsSet() {
		return true
	}

	return false
}

// SetDownloadId gets a reference to the given NullableString and assigns it to the DownloadId field.
func (o *ManualImportResource) SetDownloadId(v string) {
	o.DownloadId.Set(&v)
}
// SetDownloadIdNil sets the value for DownloadId to be an explicit nil
func (o *ManualImportResource) SetDownloadIdNil() {
	o.DownloadId.Set(nil)
}

// UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
func (o *ManualImportResource) UnsetDownloadId() {
	o.DownloadId.Unset()
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetCustomFormats() []CustomFormatResource {
	if o == nil {
		var ret []CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetCustomFormatsOk() ([]CustomFormatResource, bool) {
	if o == nil || IsNil(o.CustomFormats) {
		return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *ManualImportResource) HasCustomFormats() bool {
	if o != nil && IsNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *ManualImportResource) SetCustomFormats(v []CustomFormatResource) {
	o.CustomFormats = v
}

// GetCustomFormatScore returns the CustomFormatScore field value if set, zero value otherwise.
func (o *ManualImportResource) GetCustomFormatScore() int32 {
	if o == nil || IsNil(o.CustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.CustomFormatScore
}

// GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualImportResource) GetCustomFormatScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomFormatScore) {
		return nil, false
	}
	return o.CustomFormatScore, true
}

// HasCustomFormatScore returns a boolean if a field has been set.
func (o *ManualImportResource) HasCustomFormatScore() bool {
	if o != nil && !IsNil(o.CustomFormatScore) {
		return true
	}

	return false
}

// SetCustomFormatScore gets a reference to the given int32 and assigns it to the CustomFormatScore field.
func (o *ManualImportResource) SetCustomFormatScore(v int32) {
	o.CustomFormatScore = &v
}

// GetRejections returns the Rejections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualImportResource) GetRejections() []Rejection {
	if o == nil {
		var ret []Rejection
		return ret
	}
	return o.Rejections
}

// GetRejectionsOk returns a tuple with the Rejections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualImportResource) GetRejectionsOk() ([]Rejection, bool) {
	if o == nil || IsNil(o.Rejections) {
		return nil, false
	}
	return o.Rejections, true
}

// HasRejections returns a boolean if a field has been set.
func (o *ManualImportResource) HasRejections() bool {
	if o != nil && IsNil(o.Rejections) {
		return true
	}

	return false
}

// SetRejections gets a reference to the given []Rejection and assigns it to the Rejections field.
func (o *ManualImportResource) SetRejections(v []Rejection) {
	o.Rejections = v
}

func (o ManualImportResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualImportResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.RelativePath.IsSet() {
		toSerialize["relativePath"] = o.RelativePath.Get()
	}
	if o.FolderName.IsSet() {
		toSerialize["folderName"] = o.FolderName.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Movie) {
		toSerialize["movie"] = o.Movie
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.ReleaseGroup.IsSet() {
		toSerialize["releaseGroup"] = o.ReleaseGroup.Get()
	}
	if !IsNil(o.QualityWeight) {
		toSerialize["qualityWeight"] = o.QualityWeight
	}
	if o.DownloadId.IsSet() {
		toSerialize["downloadId"] = o.DownloadId.Get()
	}
	if o.CustomFormats != nil {
		toSerialize["customFormats"] = o.CustomFormats
	}
	if !IsNil(o.CustomFormatScore) {
		toSerialize["customFormatScore"] = o.CustomFormatScore
	}
	if o.Rejections != nil {
		toSerialize["rejections"] = o.Rejections
	}
	return toSerialize, nil
}

type NullableManualImportResource struct {
	value *ManualImportResource
	isSet bool
}

func (v NullableManualImportResource) Get() *ManualImportResource {
	return v.value
}

func (v *NullableManualImportResource) Set(val *ManualImportResource) {
	v.value = val
	v.isSet = true
}

func (v NullableManualImportResource) IsSet() bool {
	return v.isSet
}

func (v *NullableManualImportResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualImportResource(val *ManualImportResource) *NullableManualImportResource {
	return &NullableManualImportResource{value: val, isSet: true}
}

func (v NullableManualImportResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualImportResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



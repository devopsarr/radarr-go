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

// checks if the ParsedMovieInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParsedMovieInfo{}

// ParsedMovieInfo struct for ParsedMovieInfo
type ParsedMovieInfo struct {
	MovieTitles []string `json:"movieTitles,omitempty"`
	OriginalTitle NullableString `json:"originalTitle,omitempty"`
	ReleaseTitle NullableString `json:"releaseTitle,omitempty"`
	SimpleReleaseTitle NullableString `json:"simpleReleaseTitle,omitempty"`
	Quality *QualityModel `json:"quality,omitempty"`
	Languages []Language `json:"languages,omitempty"`
	ReleaseGroup NullableString `json:"releaseGroup,omitempty"`
	ReleaseHash NullableString `json:"releaseHash,omitempty"`
	Edition NullableString `json:"edition,omitempty"`
	Year *int32 `json:"year,omitempty"`
	ImdbId NullableString `json:"imdbId,omitempty"`
	TmdbId *int32 `json:"tmdbId,omitempty"`
	HardcodedSubs NullableString `json:"hardcodedSubs,omitempty"`
	MovieTitle NullableString `json:"movieTitle,omitempty"`
	PrimaryMovieTitle NullableString `json:"primaryMovieTitle,omitempty"`
}

// NewParsedMovieInfo instantiates a new ParsedMovieInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParsedMovieInfo() *ParsedMovieInfo {
	this := ParsedMovieInfo{}
	return &this
}

// NewParsedMovieInfoWithDefaults instantiates a new ParsedMovieInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParsedMovieInfoWithDefaults() *ParsedMovieInfo {
	this := ParsedMovieInfo{}
	return &this
}

// GetMovieTitles returns the MovieTitles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetMovieTitles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MovieTitles
}

// GetMovieTitlesOk returns a tuple with the MovieTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetMovieTitlesOk() ([]string, bool) {
	if o == nil || IsNil(o.MovieTitles) {
		return nil, false
	}
	return o.MovieTitles, true
}

// HasMovieTitles returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasMovieTitles() bool {
	if o != nil && !IsNil(o.MovieTitles) {
		return true
	}

	return false
}

// SetMovieTitles gets a reference to the given []string and assigns it to the MovieTitles field.
func (o *ParsedMovieInfo) SetMovieTitles(v []string) {
	o.MovieTitles = v
}

// GetOriginalTitle returns the OriginalTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetOriginalTitle() string {
	if o == nil || IsNil(o.OriginalTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalTitle.Get()
}

// GetOriginalTitleOk returns a tuple with the OriginalTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetOriginalTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalTitle.Get(), o.OriginalTitle.IsSet()
}

// HasOriginalTitle returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasOriginalTitle() bool {
	if o != nil && o.OriginalTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginalTitle gets a reference to the given NullableString and assigns it to the OriginalTitle field.
func (o *ParsedMovieInfo) SetOriginalTitle(v string) {
	o.OriginalTitle.Set(&v)
}
// SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil
func (o *ParsedMovieInfo) SetOriginalTitleNil() {
	o.OriginalTitle.Set(nil)
}

// UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
func (o *ParsedMovieInfo) UnsetOriginalTitle() {
	o.OriginalTitle.Unset()
}

// GetReleaseTitle returns the ReleaseTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetReleaseTitle() string {
	if o == nil || IsNil(o.ReleaseTitle.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseTitle.Get()
}

// GetReleaseTitleOk returns a tuple with the ReleaseTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetReleaseTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseTitle.Get(), o.ReleaseTitle.IsSet()
}

// HasReleaseTitle returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasReleaseTitle() bool {
	if o != nil && o.ReleaseTitle.IsSet() {
		return true
	}

	return false
}

// SetReleaseTitle gets a reference to the given NullableString and assigns it to the ReleaseTitle field.
func (o *ParsedMovieInfo) SetReleaseTitle(v string) {
	o.ReleaseTitle.Set(&v)
}
// SetReleaseTitleNil sets the value for ReleaseTitle to be an explicit nil
func (o *ParsedMovieInfo) SetReleaseTitleNil() {
	o.ReleaseTitle.Set(nil)
}

// UnsetReleaseTitle ensures that no value is present for ReleaseTitle, not even an explicit nil
func (o *ParsedMovieInfo) UnsetReleaseTitle() {
	o.ReleaseTitle.Unset()
}

// GetSimpleReleaseTitle returns the SimpleReleaseTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetSimpleReleaseTitle() string {
	if o == nil || IsNil(o.SimpleReleaseTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SimpleReleaseTitle.Get()
}

// GetSimpleReleaseTitleOk returns a tuple with the SimpleReleaseTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetSimpleReleaseTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SimpleReleaseTitle.Get(), o.SimpleReleaseTitle.IsSet()
}

// HasSimpleReleaseTitle returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasSimpleReleaseTitle() bool {
	if o != nil && o.SimpleReleaseTitle.IsSet() {
		return true
	}

	return false
}

// SetSimpleReleaseTitle gets a reference to the given NullableString and assigns it to the SimpleReleaseTitle field.
func (o *ParsedMovieInfo) SetSimpleReleaseTitle(v string) {
	o.SimpleReleaseTitle.Set(&v)
}
// SetSimpleReleaseTitleNil sets the value for SimpleReleaseTitle to be an explicit nil
func (o *ParsedMovieInfo) SetSimpleReleaseTitleNil() {
	o.SimpleReleaseTitle.Set(nil)
}

// UnsetSimpleReleaseTitle ensures that no value is present for SimpleReleaseTitle, not even an explicit nil
func (o *ParsedMovieInfo) UnsetSimpleReleaseTitle() {
	o.SimpleReleaseTitle.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ParsedMovieInfo) GetQuality() QualityModel {
	if o == nil || IsNil(o.Quality) {
		var ret QualityModel
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedMovieInfo) GetQualityOk() (*QualityModel, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given QualityModel and assigns it to the Quality field.
func (o *ParsedMovieInfo) SetQuality(v QualityModel) {
	o.Quality = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetLanguages() []Language {
	if o == nil {
		var ret []Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetLanguagesOk() ([]Language, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *ParsedMovieInfo) SetLanguages(v []Language) {
	o.Languages = v
}

// GetReleaseGroup returns the ReleaseGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetReleaseGroup() string {
	if o == nil || IsNil(o.ReleaseGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseGroup.Get()
}

// GetReleaseGroupOk returns a tuple with the ReleaseGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetReleaseGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseGroup.Get(), o.ReleaseGroup.IsSet()
}

// HasReleaseGroup returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasReleaseGroup() bool {
	if o != nil && o.ReleaseGroup.IsSet() {
		return true
	}

	return false
}

// SetReleaseGroup gets a reference to the given NullableString and assigns it to the ReleaseGroup field.
func (o *ParsedMovieInfo) SetReleaseGroup(v string) {
	o.ReleaseGroup.Set(&v)
}
// SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil
func (o *ParsedMovieInfo) SetReleaseGroupNil() {
	o.ReleaseGroup.Set(nil)
}

// UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
func (o *ParsedMovieInfo) UnsetReleaseGroup() {
	o.ReleaseGroup.Unset()
}

// GetReleaseHash returns the ReleaseHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetReleaseHash() string {
	if o == nil || IsNil(o.ReleaseHash.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseHash.Get()
}

// GetReleaseHashOk returns a tuple with the ReleaseHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetReleaseHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseHash.Get(), o.ReleaseHash.IsSet()
}

// HasReleaseHash returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasReleaseHash() bool {
	if o != nil && o.ReleaseHash.IsSet() {
		return true
	}

	return false
}

// SetReleaseHash gets a reference to the given NullableString and assigns it to the ReleaseHash field.
func (o *ParsedMovieInfo) SetReleaseHash(v string) {
	o.ReleaseHash.Set(&v)
}
// SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil
func (o *ParsedMovieInfo) SetReleaseHashNil() {
	o.ReleaseHash.Set(nil)
}

// UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
func (o *ParsedMovieInfo) UnsetReleaseHash() {
	o.ReleaseHash.Unset()
}

// GetEdition returns the Edition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetEdition() string {
	if o == nil || IsNil(o.Edition.Get()) {
		var ret string
		return ret
	}
	return *o.Edition.Get()
}

// GetEditionOk returns a tuple with the Edition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Edition.Get(), o.Edition.IsSet()
}

// HasEdition returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasEdition() bool {
	if o != nil && o.Edition.IsSet() {
		return true
	}

	return false
}

// SetEdition gets a reference to the given NullableString and assigns it to the Edition field.
func (o *ParsedMovieInfo) SetEdition(v string) {
	o.Edition.Set(&v)
}
// SetEditionNil sets the value for Edition to be an explicit nil
func (o *ParsedMovieInfo) SetEditionNil() {
	o.Edition.Set(nil)
}

// UnsetEdition ensures that no value is present for Edition, not even an explicit nil
func (o *ParsedMovieInfo) UnsetEdition() {
	o.Edition.Unset()
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *ParsedMovieInfo) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedMovieInfo) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *ParsedMovieInfo) SetYear(v int32) {
	o.Year = &v
}

// GetImdbId returns the ImdbId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetImdbId() string {
	if o == nil || IsNil(o.ImdbId.Get()) {
		var ret string
		return ret
	}
	return *o.ImdbId.Get()
}

// GetImdbIdOk returns a tuple with the ImdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetImdbIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImdbId.Get(), o.ImdbId.IsSet()
}

// HasImdbId returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasImdbId() bool {
	if o != nil && o.ImdbId.IsSet() {
		return true
	}

	return false
}

// SetImdbId gets a reference to the given NullableString and assigns it to the ImdbId field.
func (o *ParsedMovieInfo) SetImdbId(v string) {
	o.ImdbId.Set(&v)
}
// SetImdbIdNil sets the value for ImdbId to be an explicit nil
func (o *ParsedMovieInfo) SetImdbIdNil() {
	o.ImdbId.Set(nil)
}

// UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
func (o *ParsedMovieInfo) UnsetImdbId() {
	o.ImdbId.Unset()
}

// GetTmdbId returns the TmdbId field value if set, zero value otherwise.
func (o *ParsedMovieInfo) GetTmdbId() int32 {
	if o == nil || IsNil(o.TmdbId) {
		var ret int32
		return ret
	}
	return *o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParsedMovieInfo) GetTmdbIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TmdbId) {
		return nil, false
	}
	return o.TmdbId, true
}

// HasTmdbId returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasTmdbId() bool {
	if o != nil && !IsNil(o.TmdbId) {
		return true
	}

	return false
}

// SetTmdbId gets a reference to the given int32 and assigns it to the TmdbId field.
func (o *ParsedMovieInfo) SetTmdbId(v int32) {
	o.TmdbId = &v
}

// GetHardcodedSubs returns the HardcodedSubs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetHardcodedSubs() string {
	if o == nil || IsNil(o.HardcodedSubs.Get()) {
		var ret string
		return ret
	}
	return *o.HardcodedSubs.Get()
}

// GetHardcodedSubsOk returns a tuple with the HardcodedSubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetHardcodedSubsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HardcodedSubs.Get(), o.HardcodedSubs.IsSet()
}

// HasHardcodedSubs returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasHardcodedSubs() bool {
	if o != nil && o.HardcodedSubs.IsSet() {
		return true
	}

	return false
}

// SetHardcodedSubs gets a reference to the given NullableString and assigns it to the HardcodedSubs field.
func (o *ParsedMovieInfo) SetHardcodedSubs(v string) {
	o.HardcodedSubs.Set(&v)
}
// SetHardcodedSubsNil sets the value for HardcodedSubs to be an explicit nil
func (o *ParsedMovieInfo) SetHardcodedSubsNil() {
	o.HardcodedSubs.Set(nil)
}

// UnsetHardcodedSubs ensures that no value is present for HardcodedSubs, not even an explicit nil
func (o *ParsedMovieInfo) UnsetHardcodedSubs() {
	o.HardcodedSubs.Unset()
}

// GetMovieTitle returns the MovieTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetMovieTitle() string {
	if o == nil || IsNil(o.MovieTitle.Get()) {
		var ret string
		return ret
	}
	return *o.MovieTitle.Get()
}

// GetMovieTitleOk returns a tuple with the MovieTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetMovieTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MovieTitle.Get(), o.MovieTitle.IsSet()
}

// HasMovieTitle returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasMovieTitle() bool {
	if o != nil && o.MovieTitle.IsSet() {
		return true
	}

	return false
}

// SetMovieTitle gets a reference to the given NullableString and assigns it to the MovieTitle field.
func (o *ParsedMovieInfo) SetMovieTitle(v string) {
	o.MovieTitle.Set(&v)
}
// SetMovieTitleNil sets the value for MovieTitle to be an explicit nil
func (o *ParsedMovieInfo) SetMovieTitleNil() {
	o.MovieTitle.Set(nil)
}

// UnsetMovieTitle ensures that no value is present for MovieTitle, not even an explicit nil
func (o *ParsedMovieInfo) UnsetMovieTitle() {
	o.MovieTitle.Unset()
}

// GetPrimaryMovieTitle returns the PrimaryMovieTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParsedMovieInfo) GetPrimaryMovieTitle() string {
	if o == nil || IsNil(o.PrimaryMovieTitle.Get()) {
		var ret string
		return ret
	}
	return *o.PrimaryMovieTitle.Get()
}

// GetPrimaryMovieTitleOk returns a tuple with the PrimaryMovieTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParsedMovieInfo) GetPrimaryMovieTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryMovieTitle.Get(), o.PrimaryMovieTitle.IsSet()
}

// HasPrimaryMovieTitle returns a boolean if a field has been set.
func (o *ParsedMovieInfo) HasPrimaryMovieTitle() bool {
	if o != nil && o.PrimaryMovieTitle.IsSet() {
		return true
	}

	return false
}

// SetPrimaryMovieTitle gets a reference to the given NullableString and assigns it to the PrimaryMovieTitle field.
func (o *ParsedMovieInfo) SetPrimaryMovieTitle(v string) {
	o.PrimaryMovieTitle.Set(&v)
}
// SetPrimaryMovieTitleNil sets the value for PrimaryMovieTitle to be an explicit nil
func (o *ParsedMovieInfo) SetPrimaryMovieTitleNil() {
	o.PrimaryMovieTitle.Set(nil)
}

// UnsetPrimaryMovieTitle ensures that no value is present for PrimaryMovieTitle, not even an explicit nil
func (o *ParsedMovieInfo) UnsetPrimaryMovieTitle() {
	o.PrimaryMovieTitle.Unset()
}

func (o ParsedMovieInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParsedMovieInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MovieTitles != nil {
		toSerialize["movieTitles"] = o.MovieTitles
	}
	if o.OriginalTitle.IsSet() {
		toSerialize["originalTitle"] = o.OriginalTitle.Get()
	}
	if o.ReleaseTitle.IsSet() {
		toSerialize["releaseTitle"] = o.ReleaseTitle.Get()
	}
	if o.SimpleReleaseTitle.IsSet() {
		toSerialize["simpleReleaseTitle"] = o.SimpleReleaseTitle.Get()
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
	if o.ReleaseHash.IsSet() {
		toSerialize["releaseHash"] = o.ReleaseHash.Get()
	}
	if o.Edition.IsSet() {
		toSerialize["edition"] = o.Edition.Get()
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if o.ImdbId.IsSet() {
		toSerialize["imdbId"] = o.ImdbId.Get()
	}
	if !IsNil(o.TmdbId) {
		toSerialize["tmdbId"] = o.TmdbId
	}
	if o.HardcodedSubs.IsSet() {
		toSerialize["hardcodedSubs"] = o.HardcodedSubs.Get()
	}
	if o.MovieTitle.IsSet() {
		toSerialize["movieTitle"] = o.MovieTitle.Get()
	}
	if o.PrimaryMovieTitle.IsSet() {
		toSerialize["primaryMovieTitle"] = o.PrimaryMovieTitle.Get()
	}
	return toSerialize, nil
}

type NullableParsedMovieInfo struct {
	value *ParsedMovieInfo
	isSet bool
}

func (v NullableParsedMovieInfo) Get() *ParsedMovieInfo {
	return v.value
}

func (v *NullableParsedMovieInfo) Set(val *ParsedMovieInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableParsedMovieInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableParsedMovieInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParsedMovieInfo(val *ParsedMovieInfo) *NullableParsedMovieInfo {
	return &NullableParsedMovieInfo{value: val, isSet: true}
}

func (v NullableParsedMovieInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParsedMovieInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



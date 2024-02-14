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

// ParseResource struct for ParseResource
type ParseResource struct {
	Id *int32 `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	ParsedMovieInfo *ParsedMovieInfo `json:"parsedMovieInfo,omitempty"`
	Movie *MovieResource `json:"movie,omitempty"`
	Languages []*Language `json:"languages,omitempty"`
	CustomFormats []*CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore *int32 `json:"customFormatScore,omitempty"`
}

// NewParseResource instantiates a new ParseResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParseResource() *ParseResource {
	this := ParseResource{}
	return &this
}

// NewParseResourceWithDefaults instantiates a new ParseResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParseResourceWithDefaults() *ParseResource {
	this := ParseResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ParseResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ParseResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ParseResource) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseResource) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseResource) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ParseResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ParseResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ParseResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ParseResource) UnsetTitle() {
	o.Title.Unset()
}

// GetParsedMovieInfo returns the ParsedMovieInfo field value if set, zero value otherwise.
func (o *ParseResource) GetParsedMovieInfo() ParsedMovieInfo {
	if o == nil || isNil(o.ParsedMovieInfo) {
		var ret ParsedMovieInfo
		return ret
	}
	return *o.ParsedMovieInfo
}

// GetParsedMovieInfoOk returns a tuple with the ParsedMovieInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseResource) GetParsedMovieInfoOk() (*ParsedMovieInfo, bool) {
	if o == nil || isNil(o.ParsedMovieInfo) {
    return nil, false
	}
	return o.ParsedMovieInfo, true
}

// HasParsedMovieInfo returns a boolean if a field has been set.
func (o *ParseResource) HasParsedMovieInfo() bool {
	if o != nil && !isNil(o.ParsedMovieInfo) {
		return true
	}

	return false
}

// SetParsedMovieInfo gets a reference to the given ParsedMovieInfo and assigns it to the ParsedMovieInfo field.
func (o *ParseResource) SetParsedMovieInfo(v ParsedMovieInfo) {
	o.ParsedMovieInfo = &v
}

// GetMovie returns the Movie field value if set, zero value otherwise.
func (o *ParseResource) GetMovie() MovieResource {
	if o == nil || isNil(o.Movie) {
		var ret MovieResource
		return ret
	}
	return *o.Movie
}

// GetMovieOk returns a tuple with the Movie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseResource) GetMovieOk() (*MovieResource, bool) {
	if o == nil || isNil(o.Movie) {
    return nil, false
	}
	return o.Movie, true
}

// HasMovie returns a boolean if a field has been set.
func (o *ParseResource) HasMovie() bool {
	if o != nil && !isNil(o.Movie) {
		return true
	}

	return false
}

// SetMovie gets a reference to the given MovieResource and assigns it to the Movie field.
func (o *ParseResource) SetMovie(v MovieResource) {
	o.Movie = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseResource) GetLanguages() []*Language {
	if o == nil {
		var ret []*Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseResource) GetLanguagesOk() ([]*Language, bool) {
	if o == nil || isNil(o.Languages) {
    return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ParseResource) HasLanguages() bool {
	if o != nil && isNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *ParseResource) SetLanguages(v []*Language) {
	o.Languages = v
}

// GetCustomFormats returns the CustomFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseResource) GetCustomFormats() []*CustomFormatResource {
	if o == nil {
		var ret []*CustomFormatResource
		return ret
	}
	return o.CustomFormats
}

// GetCustomFormatsOk returns a tuple with the CustomFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseResource) GetCustomFormatsOk() ([]*CustomFormatResource, bool) {
	if o == nil || isNil(o.CustomFormats) {
    return nil, false
	}
	return o.CustomFormats, true
}

// HasCustomFormats returns a boolean if a field has been set.
func (o *ParseResource) HasCustomFormats() bool {
	if o != nil && isNil(o.CustomFormats) {
		return true
	}

	return false
}

// SetCustomFormats gets a reference to the given []CustomFormatResource and assigns it to the CustomFormats field.
func (o *ParseResource) SetCustomFormats(v []*CustomFormatResource) {
	o.CustomFormats = v
}

// GetCustomFormatScore returns the CustomFormatScore field value if set, zero value otherwise.
func (o *ParseResource) GetCustomFormatScore() int32 {
	if o == nil || isNil(o.CustomFormatScore) {
		var ret int32
		return ret
	}
	return *o.CustomFormatScore
}

// GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParseResource) GetCustomFormatScoreOk() (*int32, bool) {
	if o == nil || isNil(o.CustomFormatScore) {
    return nil, false
	}
	return o.CustomFormatScore, true
}

// HasCustomFormatScore returns a boolean if a field has been set.
func (o *ParseResource) HasCustomFormatScore() bool {
	if o != nil && !isNil(o.CustomFormatScore) {
		return true
	}

	return false
}

// SetCustomFormatScore gets a reference to the given int32 and assigns it to the CustomFormatScore field.
func (o *ParseResource) SetCustomFormatScore(v int32) {
	o.CustomFormatScore = &v
}

func (o ParseResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !isNil(o.ParsedMovieInfo) {
		toSerialize["parsedMovieInfo"] = o.ParsedMovieInfo
	}
	if !isNil(o.Movie) {
		toSerialize["movie"] = o.Movie
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.CustomFormats != nil {
		toSerialize["customFormats"] = o.CustomFormats
	}
	if !isNil(o.CustomFormatScore) {
		toSerialize["customFormatScore"] = o.CustomFormatScore
	}
	return json.Marshal(toSerialize)
}

type NullableParseResource struct {
	value *ParseResource
	isSet bool
}

func (v NullableParseResource) Get() *ParseResource {
	return v.value
}

func (v *NullableParseResource) Set(val *ParseResource) {
	v.value = val
	v.isSet = true
}

func (v NullableParseResource) IsSet() bool {
	return v.isSet
}

func (v *NullableParseResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParseResource(val *ParseResource) *NullableParseResource {
	return &NullableParseResource{value: val, isSet: true}
}

func (v NullableParseResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParseResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



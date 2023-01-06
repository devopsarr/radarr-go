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

// AlternativeTitle struct for AlternativeTitle
type AlternativeTitle struct {
	Id *int32 `json:"id,omitempty"`
	SourceType *SourceType `json:"sourceType,omitempty"`
	MovieMetadataId *int32 `json:"movieMetadataId,omitempty"`
	Title NullableString `json:"title,omitempty"`
	CleanTitle NullableString `json:"cleanTitle,omitempty"`
	SourceId *int32 `json:"sourceId,omitempty"`
	Votes *int32 `json:"votes,omitempty"`
	VoteCount *int32 `json:"voteCount,omitempty"`
	Language *Language `json:"language,omitempty"`
}

// NewAlternativeTitle instantiates a new AlternativeTitle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeTitle() *AlternativeTitle {
	this := AlternativeTitle{}
	return &this
}

// NewAlternativeTitleWithDefaults instantiates a new AlternativeTitle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeTitleWithDefaults() *AlternativeTitle {
	this := AlternativeTitle{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlternativeTitle) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitle) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlternativeTitle) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlternativeTitle) SetId(v int32) {
	o.Id = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *AlternativeTitle) GetSourceType() SourceType {
	if o == nil || isNil(o.SourceType) {
		var ret SourceType
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitle) GetSourceTypeOk() (*SourceType, bool) {
	if o == nil || isNil(o.SourceType) {
    return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *AlternativeTitle) HasSourceType() bool {
	if o != nil && !isNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given SourceType and assigns it to the SourceType field.
func (o *AlternativeTitle) SetSourceType(v SourceType) {
	o.SourceType = &v
}

// GetMovieMetadataId returns the MovieMetadataId field value if set, zero value otherwise.
func (o *AlternativeTitle) GetMovieMetadataId() int32 {
	if o == nil || isNil(o.MovieMetadataId) {
		var ret int32
		return ret
	}
	return *o.MovieMetadataId
}

// GetMovieMetadataIdOk returns a tuple with the MovieMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitle) GetMovieMetadataIdOk() (*int32, bool) {
	if o == nil || isNil(o.MovieMetadataId) {
    return nil, false
	}
	return o.MovieMetadataId, true
}

// HasMovieMetadataId returns a boolean if a field has been set.
func (o *AlternativeTitle) HasMovieMetadataId() bool {
	if o != nil && !isNil(o.MovieMetadataId) {
		return true
	}

	return false
}

// SetMovieMetadataId gets a reference to the given int32 and assigns it to the MovieMetadataId field.
func (o *AlternativeTitle) SetMovieMetadataId(v int32) {
	o.MovieMetadataId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlternativeTitle) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlternativeTitle) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AlternativeTitle) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AlternativeTitle) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AlternativeTitle) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AlternativeTitle) UnsetTitle() {
	o.Title.Unset()
}

// GetCleanTitle returns the CleanTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlternativeTitle) GetCleanTitle() string {
	if o == nil || isNil(o.CleanTitle.Get()) {
		var ret string
		return ret
	}
	return *o.CleanTitle.Get()
}

// GetCleanTitleOk returns a tuple with the CleanTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlternativeTitle) GetCleanTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CleanTitle.Get(), o.CleanTitle.IsSet()
}

// HasCleanTitle returns a boolean if a field has been set.
func (o *AlternativeTitle) HasCleanTitle() bool {
	if o != nil && o.CleanTitle.IsSet() {
		return true
	}

	return false
}

// SetCleanTitle gets a reference to the given NullableString and assigns it to the CleanTitle field.
func (o *AlternativeTitle) SetCleanTitle(v string) {
	o.CleanTitle.Set(&v)
}
// SetCleanTitleNil sets the value for CleanTitle to be an explicit nil
func (o *AlternativeTitle) SetCleanTitleNil() {
	o.CleanTitle.Set(nil)
}

// UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
func (o *AlternativeTitle) UnsetCleanTitle() {
	o.CleanTitle.Unset()
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AlternativeTitle) GetSourceId() int32 {
	if o == nil || isNil(o.SourceId) {
		var ret int32
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitle) GetSourceIdOk() (*int32, bool) {
	if o == nil || isNil(o.SourceId) {
    return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AlternativeTitle) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given int32 and assigns it to the SourceId field.
func (o *AlternativeTitle) SetSourceId(v int32) {
	o.SourceId = &v
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *AlternativeTitle) GetVotes() int32 {
	if o == nil || isNil(o.Votes) {
		var ret int32
		return ret
	}
	return *o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitle) GetVotesOk() (*int32, bool) {
	if o == nil || isNil(o.Votes) {
    return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *AlternativeTitle) HasVotes() bool {
	if o != nil && !isNil(o.Votes) {
		return true
	}

	return false
}

// SetVotes gets a reference to the given int32 and assigns it to the Votes field.
func (o *AlternativeTitle) SetVotes(v int32) {
	o.Votes = &v
}

// GetVoteCount returns the VoteCount field value if set, zero value otherwise.
func (o *AlternativeTitle) GetVoteCount() int32 {
	if o == nil || isNil(o.VoteCount) {
		var ret int32
		return ret
	}
	return *o.VoteCount
}

// GetVoteCountOk returns a tuple with the VoteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitle) GetVoteCountOk() (*int32, bool) {
	if o == nil || isNil(o.VoteCount) {
    return nil, false
	}
	return o.VoteCount, true
}

// HasVoteCount returns a boolean if a field has been set.
func (o *AlternativeTitle) HasVoteCount() bool {
	if o != nil && !isNil(o.VoteCount) {
		return true
	}

	return false
}

// SetVoteCount gets a reference to the given int32 and assigns it to the VoteCount field.
func (o *AlternativeTitle) SetVoteCount(v int32) {
	o.VoteCount = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AlternativeTitle) GetLanguage() Language {
	if o == nil || isNil(o.Language) {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeTitle) GetLanguageOk() (*Language, bool) {
	if o == nil || isNil(o.Language) {
    return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AlternativeTitle) HasLanguage() bool {
	if o != nil && !isNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *AlternativeTitle) SetLanguage(v Language) {
	o.Language = &v
}

func (o AlternativeTitle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !isNil(o.MovieMetadataId) {
		toSerialize["movieMetadataId"] = o.MovieMetadataId
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.CleanTitle.IsSet() {
		toSerialize["cleanTitle"] = o.CleanTitle.Get()
	}
	if !isNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !isNil(o.Votes) {
		toSerialize["votes"] = o.Votes
	}
	if !isNil(o.VoteCount) {
		toSerialize["voteCount"] = o.VoteCount
	}
	if !isNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return json.Marshal(toSerialize)
}

type NullableAlternativeTitle struct {
	value *AlternativeTitle
	isSet bool
}

func (v NullableAlternativeTitle) Get() *AlternativeTitle {
	return v.value
}

func (v *NullableAlternativeTitle) Set(val *AlternativeTitle) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeTitle) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeTitle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeTitle(val *AlternativeTitle) *NullableAlternativeTitle {
	return &NullableAlternativeTitle{value: val, isSet: true}
}

func (v NullableAlternativeTitle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeTitle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



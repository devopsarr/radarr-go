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

// ImportExclusionsResource struct for ImportExclusionsResource
type ImportExclusionsResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Fields []*Field `json:"fields,omitempty"`
	ImplementationName NullableString `json:"implementationName,omitempty"`
	Implementation NullableString `json:"implementation,omitempty"`
	ConfigContract NullableString `json:"configContract,omitempty"`
	InfoLink NullableString `json:"infoLink,omitempty"`
	Message *ProviderMessage `json:"message,omitempty"`
	Tags []*int32 `json:"tags,omitempty"`
	Presets []*ImportExclusionsResource `json:"presets,omitempty"`
	TmdbId *int32 `json:"tmdbId,omitempty"`
	MovieTitle NullableString `json:"movieTitle,omitempty"`
	MovieYear *int32 `json:"movieYear,omitempty"`
}

// NewImportExclusionsResource instantiates a new ImportExclusionsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportExclusionsResource() *ImportExclusionsResource {
	this := ImportExclusionsResource{}
	return &this
}

// NewImportExclusionsResourceWithDefaults instantiates a new ImportExclusionsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportExclusionsResourceWithDefaults() *ImportExclusionsResource {
	this := ImportExclusionsResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportExclusionsResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportExclusionsResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ImportExclusionsResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportExclusionsResource) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportExclusionsResource) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ImportExclusionsResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ImportExclusionsResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ImportExclusionsResource) UnsetName() {
	o.Name.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportExclusionsResource) GetFields() []*Field {
	if o == nil {
		var ret []*Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportExclusionsResource) GetFieldsOk() ([]*Field, bool) {
	if o == nil || isNil(o.Fields) {
    return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasFields() bool {
	if o != nil && isNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *ImportExclusionsResource) SetFields(v []*Field) {
	o.Fields = v
}

// GetImplementationName returns the ImplementationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportExclusionsResource) GetImplementationName() string {
	if o == nil || isNil(o.ImplementationName.Get()) {
		var ret string
		return ret
	}
	return *o.ImplementationName.Get()
}

// GetImplementationNameOk returns a tuple with the ImplementationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportExclusionsResource) GetImplementationNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ImplementationName.Get(), o.ImplementationName.IsSet()
}

// HasImplementationName returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasImplementationName() bool {
	if o != nil && o.ImplementationName.IsSet() {
		return true
	}

	return false
}

// SetImplementationName gets a reference to the given NullableString and assigns it to the ImplementationName field.
func (o *ImportExclusionsResource) SetImplementationName(v string) {
	o.ImplementationName.Set(&v)
}
// SetImplementationNameNil sets the value for ImplementationName to be an explicit nil
func (o *ImportExclusionsResource) SetImplementationNameNil() {
	o.ImplementationName.Set(nil)
}

// UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
func (o *ImportExclusionsResource) UnsetImplementationName() {
	o.ImplementationName.Unset()
}

// GetImplementation returns the Implementation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportExclusionsResource) GetImplementation() string {
	if o == nil || isNil(o.Implementation.Get()) {
		var ret string
		return ret
	}
	return *o.Implementation.Get()
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportExclusionsResource) GetImplementationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Implementation.Get(), o.Implementation.IsSet()
}

// HasImplementation returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasImplementation() bool {
	if o != nil && o.Implementation.IsSet() {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given NullableString and assigns it to the Implementation field.
func (o *ImportExclusionsResource) SetImplementation(v string) {
	o.Implementation.Set(&v)
}
// SetImplementationNil sets the value for Implementation to be an explicit nil
func (o *ImportExclusionsResource) SetImplementationNil() {
	o.Implementation.Set(nil)
}

// UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
func (o *ImportExclusionsResource) UnsetImplementation() {
	o.Implementation.Unset()
}

// GetConfigContract returns the ConfigContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportExclusionsResource) GetConfigContract() string {
	if o == nil || isNil(o.ConfigContract.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigContract.Get()
}

// GetConfigContractOk returns a tuple with the ConfigContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportExclusionsResource) GetConfigContractOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ConfigContract.Get(), o.ConfigContract.IsSet()
}

// HasConfigContract returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasConfigContract() bool {
	if o != nil && o.ConfigContract.IsSet() {
		return true
	}

	return false
}

// SetConfigContract gets a reference to the given NullableString and assigns it to the ConfigContract field.
func (o *ImportExclusionsResource) SetConfigContract(v string) {
	o.ConfigContract.Set(&v)
}
// SetConfigContractNil sets the value for ConfigContract to be an explicit nil
func (o *ImportExclusionsResource) SetConfigContractNil() {
	o.ConfigContract.Set(nil)
}

// UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
func (o *ImportExclusionsResource) UnsetConfigContract() {
	o.ConfigContract.Unset()
}

// GetInfoLink returns the InfoLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportExclusionsResource) GetInfoLink() string {
	if o == nil || isNil(o.InfoLink.Get()) {
		var ret string
		return ret
	}
	return *o.InfoLink.Get()
}

// GetInfoLinkOk returns a tuple with the InfoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportExclusionsResource) GetInfoLinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.InfoLink.Get(), o.InfoLink.IsSet()
}

// HasInfoLink returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasInfoLink() bool {
	if o != nil && o.InfoLink.IsSet() {
		return true
	}

	return false
}

// SetInfoLink gets a reference to the given NullableString and assigns it to the InfoLink field.
func (o *ImportExclusionsResource) SetInfoLink(v string) {
	o.InfoLink.Set(&v)
}
// SetInfoLinkNil sets the value for InfoLink to be an explicit nil
func (o *ImportExclusionsResource) SetInfoLinkNil() {
	o.InfoLink.Set(nil)
}

// UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
func (o *ImportExclusionsResource) UnsetInfoLink() {
	o.InfoLink.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ImportExclusionsResource) GetMessage() ProviderMessage {
	if o == nil || isNil(o.Message) {
		var ret ProviderMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportExclusionsResource) GetMessageOk() (*ProviderMessage, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ProviderMessage and assigns it to the Message field.
func (o *ImportExclusionsResource) SetMessage(v ProviderMessage) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportExclusionsResource) GetTags() []*int32 {
	if o == nil {
		var ret []*int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportExclusionsResource) GetTagsOk() ([]*int32, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasTags() bool {
	if o != nil && isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *ImportExclusionsResource) SetTags(v []*int32) {
	o.Tags = v
}

// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportExclusionsResource) GetPresets() []*ImportExclusionsResource {
	if o == nil {
		var ret []*ImportExclusionsResource
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportExclusionsResource) GetPresetsOk() ([]*ImportExclusionsResource, bool) {
	if o == nil || isNil(o.Presets) {
    return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasPresets() bool {
	if o != nil && isNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []ImportExclusionsResource and assigns it to the Presets field.
func (o *ImportExclusionsResource) SetPresets(v []*ImportExclusionsResource) {
	o.Presets = v
}

// GetTmdbId returns the TmdbId field value if set, zero value otherwise.
func (o *ImportExclusionsResource) GetTmdbId() int32 {
	if o == nil || isNil(o.TmdbId) {
		var ret int32
		return ret
	}
	return *o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportExclusionsResource) GetTmdbIdOk() (*int32, bool) {
	if o == nil || isNil(o.TmdbId) {
    return nil, false
	}
	return o.TmdbId, true
}

// HasTmdbId returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasTmdbId() bool {
	if o != nil && !isNil(o.TmdbId) {
		return true
	}

	return false
}

// SetTmdbId gets a reference to the given int32 and assigns it to the TmdbId field.
func (o *ImportExclusionsResource) SetTmdbId(v int32) {
	o.TmdbId = &v
}

// GetMovieTitle returns the MovieTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportExclusionsResource) GetMovieTitle() string {
	if o == nil || isNil(o.MovieTitle.Get()) {
		var ret string
		return ret
	}
	return *o.MovieTitle.Get()
}

// GetMovieTitleOk returns a tuple with the MovieTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportExclusionsResource) GetMovieTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MovieTitle.Get(), o.MovieTitle.IsSet()
}

// HasMovieTitle returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasMovieTitle() bool {
	if o != nil && o.MovieTitle.IsSet() {
		return true
	}

	return false
}

// SetMovieTitle gets a reference to the given NullableString and assigns it to the MovieTitle field.
func (o *ImportExclusionsResource) SetMovieTitle(v string) {
	o.MovieTitle.Set(&v)
}
// SetMovieTitleNil sets the value for MovieTitle to be an explicit nil
func (o *ImportExclusionsResource) SetMovieTitleNil() {
	o.MovieTitle.Set(nil)
}

// UnsetMovieTitle ensures that no value is present for MovieTitle, not even an explicit nil
func (o *ImportExclusionsResource) UnsetMovieTitle() {
	o.MovieTitle.Unset()
}

// GetMovieYear returns the MovieYear field value if set, zero value otherwise.
func (o *ImportExclusionsResource) GetMovieYear() int32 {
	if o == nil || isNil(o.MovieYear) {
		var ret int32
		return ret
	}
	return *o.MovieYear
}

// GetMovieYearOk returns a tuple with the MovieYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportExclusionsResource) GetMovieYearOk() (*int32, bool) {
	if o == nil || isNil(o.MovieYear) {
    return nil, false
	}
	return o.MovieYear, true
}

// HasMovieYear returns a boolean if a field has been set.
func (o *ImportExclusionsResource) HasMovieYear() bool {
	if o != nil && !isNil(o.MovieYear) {
		return true
	}

	return false
}

// SetMovieYear gets a reference to the given int32 and assigns it to the MovieYear field.
func (o *ImportExclusionsResource) SetMovieYear(v int32) {
	o.MovieYear = &v
}

func (o ImportExclusionsResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.ImplementationName.IsSet() {
		toSerialize["implementationName"] = o.ImplementationName.Get()
	}
	if o.Implementation.IsSet() {
		toSerialize["implementation"] = o.Implementation.Get()
	}
	if o.ConfigContract.IsSet() {
		toSerialize["configContract"] = o.ConfigContract.Get()
	}
	if o.InfoLink.IsSet() {
		toSerialize["infoLink"] = o.InfoLink.Get()
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Presets != nil {
		toSerialize["presets"] = o.Presets
	}
	if !isNil(o.TmdbId) {
		toSerialize["tmdbId"] = o.TmdbId
	}
	if o.MovieTitle.IsSet() {
		toSerialize["movieTitle"] = o.MovieTitle.Get()
	}
	if !isNil(o.MovieYear) {
		toSerialize["movieYear"] = o.MovieYear
	}
	return json.Marshal(toSerialize)
}

type NullableImportExclusionsResource struct {
	value *ImportExclusionsResource
	isSet bool
}

func (v NullableImportExclusionsResource) Get() *ImportExclusionsResource {
	return v.value
}

func (v *NullableImportExclusionsResource) Set(val *ImportExclusionsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportExclusionsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportExclusionsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportExclusionsResource(val *ImportExclusionsResource) *NullableImportExclusionsResource {
	return &NullableImportExclusionsResource{value: val, isSet: true}
}

func (v NullableImportExclusionsResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportExclusionsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



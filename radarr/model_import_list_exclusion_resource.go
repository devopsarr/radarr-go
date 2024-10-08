/*
Radarr

Radarr API docs

API version: v5.10.4.9218
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the ImportListExclusionResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportListExclusionResource{}

// ImportListExclusionResource struct for ImportListExclusionResource
type ImportListExclusionResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	ImplementationName NullableString `json:"implementationName,omitempty"`
	Implementation NullableString `json:"implementation,omitempty"`
	ConfigContract NullableString `json:"configContract,omitempty"`
	InfoLink NullableString `json:"infoLink,omitempty"`
	Message *ProviderMessage `json:"message,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	Presets []ImportListExclusionResource `json:"presets,omitempty"`
	TmdbId *int32 `json:"tmdbId,omitempty"`
	MovieTitle NullableString `json:"movieTitle,omitempty"`
	MovieYear *int32 `json:"movieYear,omitempty"`
}

// NewImportListExclusionResource instantiates a new ImportListExclusionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportListExclusionResource() *ImportListExclusionResource {
	this := ImportListExclusionResource{}
	return &this
}

// NewImportListExclusionResourceWithDefaults instantiates a new ImportListExclusionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportListExclusionResourceWithDefaults() *ImportListExclusionResource {
	this := ImportListExclusionResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportListExclusionResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ImportListExclusionResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ImportListExclusionResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ImportListExclusionResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ImportListExclusionResource) UnsetName() {
	o.Name.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetFields() []Field {
	if o == nil {
		var ret []Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetFieldsOk() ([]Field, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *ImportListExclusionResource) SetFields(v []Field) {
	o.Fields = v
}

// GetImplementationName returns the ImplementationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetImplementationName() string {
	if o == nil || IsNil(o.ImplementationName.Get()) {
		var ret string
		return ret
	}
	return *o.ImplementationName.Get()
}

// GetImplementationNameOk returns a tuple with the ImplementationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetImplementationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImplementationName.Get(), o.ImplementationName.IsSet()
}

// HasImplementationName returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasImplementationName() bool {
	if o != nil && o.ImplementationName.IsSet() {
		return true
	}

	return false
}

// SetImplementationName gets a reference to the given NullableString and assigns it to the ImplementationName field.
func (o *ImportListExclusionResource) SetImplementationName(v string) {
	o.ImplementationName.Set(&v)
}
// SetImplementationNameNil sets the value for ImplementationName to be an explicit nil
func (o *ImportListExclusionResource) SetImplementationNameNil() {
	o.ImplementationName.Set(nil)
}

// UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
func (o *ImportListExclusionResource) UnsetImplementationName() {
	o.ImplementationName.Unset()
}

// GetImplementation returns the Implementation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetImplementation() string {
	if o == nil || IsNil(o.Implementation.Get()) {
		var ret string
		return ret
	}
	return *o.Implementation.Get()
}

// GetImplementationOk returns a tuple with the Implementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetImplementationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Implementation.Get(), o.Implementation.IsSet()
}

// HasImplementation returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasImplementation() bool {
	if o != nil && o.Implementation.IsSet() {
		return true
	}

	return false
}

// SetImplementation gets a reference to the given NullableString and assigns it to the Implementation field.
func (o *ImportListExclusionResource) SetImplementation(v string) {
	o.Implementation.Set(&v)
}
// SetImplementationNil sets the value for Implementation to be an explicit nil
func (o *ImportListExclusionResource) SetImplementationNil() {
	o.Implementation.Set(nil)
}

// UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
func (o *ImportListExclusionResource) UnsetImplementation() {
	o.Implementation.Unset()
}

// GetConfigContract returns the ConfigContract field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetConfigContract() string {
	if o == nil || IsNil(o.ConfigContract.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigContract.Get()
}

// GetConfigContractOk returns a tuple with the ConfigContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetConfigContractOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigContract.Get(), o.ConfigContract.IsSet()
}

// HasConfigContract returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasConfigContract() bool {
	if o != nil && o.ConfigContract.IsSet() {
		return true
	}

	return false
}

// SetConfigContract gets a reference to the given NullableString and assigns it to the ConfigContract field.
func (o *ImportListExclusionResource) SetConfigContract(v string) {
	o.ConfigContract.Set(&v)
}
// SetConfigContractNil sets the value for ConfigContract to be an explicit nil
func (o *ImportListExclusionResource) SetConfigContractNil() {
	o.ConfigContract.Set(nil)
}

// UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
func (o *ImportListExclusionResource) UnsetConfigContract() {
	o.ConfigContract.Unset()
}

// GetInfoLink returns the InfoLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetInfoLink() string {
	if o == nil || IsNil(o.InfoLink.Get()) {
		var ret string
		return ret
	}
	return *o.InfoLink.Get()
}

// GetInfoLinkOk returns a tuple with the InfoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetInfoLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfoLink.Get(), o.InfoLink.IsSet()
}

// HasInfoLink returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasInfoLink() bool {
	if o != nil && o.InfoLink.IsSet() {
		return true
	}

	return false
}

// SetInfoLink gets a reference to the given NullableString and assigns it to the InfoLink field.
func (o *ImportListExclusionResource) SetInfoLink(v string) {
	o.InfoLink.Set(&v)
}
// SetInfoLinkNil sets the value for InfoLink to be an explicit nil
func (o *ImportListExclusionResource) SetInfoLinkNil() {
	o.InfoLink.Set(nil)
}

// UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
func (o *ImportListExclusionResource) UnsetInfoLink() {
	o.InfoLink.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ImportListExclusionResource) GetMessage() ProviderMessage {
	if o == nil || IsNil(o.Message) {
		var ret ProviderMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResource) GetMessageOk() (*ProviderMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ProviderMessage and assigns it to the Message field.
func (o *ImportListExclusionResource) SetMessage(v ProviderMessage) {
	o.Message = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *ImportListExclusionResource) SetTags(v []int32) {
	o.Tags = v
}

// GetPresets returns the Presets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetPresets() []ImportListExclusionResource {
	if o == nil {
		var ret []ImportListExclusionResource
		return ret
	}
	return o.Presets
}

// GetPresetsOk returns a tuple with the Presets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetPresetsOk() ([]ImportListExclusionResource, bool) {
	if o == nil || IsNil(o.Presets) {
		return nil, false
	}
	return o.Presets, true
}

// HasPresets returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasPresets() bool {
	if o != nil && !IsNil(o.Presets) {
		return true
	}

	return false
}

// SetPresets gets a reference to the given []ImportListExclusionResource and assigns it to the Presets field.
func (o *ImportListExclusionResource) SetPresets(v []ImportListExclusionResource) {
	o.Presets = v
}

// GetTmdbId returns the TmdbId field value if set, zero value otherwise.
func (o *ImportListExclusionResource) GetTmdbId() int32 {
	if o == nil || IsNil(o.TmdbId) {
		var ret int32
		return ret
	}
	return *o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResource) GetTmdbIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TmdbId) {
		return nil, false
	}
	return o.TmdbId, true
}

// HasTmdbId returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasTmdbId() bool {
	if o != nil && !IsNil(o.TmdbId) {
		return true
	}

	return false
}

// SetTmdbId gets a reference to the given int32 and assigns it to the TmdbId field.
func (o *ImportListExclusionResource) SetTmdbId(v int32) {
	o.TmdbId = &v
}

// GetMovieTitle returns the MovieTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResource) GetMovieTitle() string {
	if o == nil || IsNil(o.MovieTitle.Get()) {
		var ret string
		return ret
	}
	return *o.MovieTitle.Get()
}

// GetMovieTitleOk returns a tuple with the MovieTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResource) GetMovieTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MovieTitle.Get(), o.MovieTitle.IsSet()
}

// HasMovieTitle returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasMovieTitle() bool {
	if o != nil && o.MovieTitle.IsSet() {
		return true
	}

	return false
}

// SetMovieTitle gets a reference to the given NullableString and assigns it to the MovieTitle field.
func (o *ImportListExclusionResource) SetMovieTitle(v string) {
	o.MovieTitle.Set(&v)
}
// SetMovieTitleNil sets the value for MovieTitle to be an explicit nil
func (o *ImportListExclusionResource) SetMovieTitleNil() {
	o.MovieTitle.Set(nil)
}

// UnsetMovieTitle ensures that no value is present for MovieTitle, not even an explicit nil
func (o *ImportListExclusionResource) UnsetMovieTitle() {
	o.MovieTitle.Unset()
}

// GetMovieYear returns the MovieYear field value if set, zero value otherwise.
func (o *ImportListExclusionResource) GetMovieYear() int32 {
	if o == nil || IsNil(o.MovieYear) {
		var ret int32
		return ret
	}
	return *o.MovieYear
}

// GetMovieYearOk returns a tuple with the MovieYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResource) GetMovieYearOk() (*int32, bool) {
	if o == nil || IsNil(o.MovieYear) {
		return nil, false
	}
	return o.MovieYear, true
}

// HasMovieYear returns a boolean if a field has been set.
func (o *ImportListExclusionResource) HasMovieYear() bool {
	if o != nil && !IsNil(o.MovieYear) {
		return true
	}

	return false
}

// SetMovieYear gets a reference to the given int32 and assigns it to the MovieYear field.
func (o *ImportListExclusionResource) SetMovieYear(v int32) {
	o.MovieYear = &v
}

func (o ImportListExclusionResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportListExclusionResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
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
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Presets != nil {
		toSerialize["presets"] = o.Presets
	}
	if !IsNil(o.TmdbId) {
		toSerialize["tmdbId"] = o.TmdbId
	}
	if o.MovieTitle.IsSet() {
		toSerialize["movieTitle"] = o.MovieTitle.Get()
	}
	if !IsNil(o.MovieYear) {
		toSerialize["movieYear"] = o.MovieYear
	}
	return toSerialize, nil
}

type NullableImportListExclusionResource struct {
	value *ImportListExclusionResource
	isSet bool
}

func (v NullableImportListExclusionResource) Get() *ImportListExclusionResource {
	return v.value
}

func (v *NullableImportListExclusionResource) Set(val *ImportListExclusionResource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListExclusionResource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListExclusionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListExclusionResource(val *ImportListExclusionResource) *NullableImportListExclusionResource {
	return &NullableImportListExclusionResource{value: val, isSet: true}
}

func (v NullableImportListExclusionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListExclusionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


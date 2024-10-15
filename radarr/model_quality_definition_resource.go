/*
Radarr

Radarr API docs

API version: v5.12.2.9335
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the QualityDefinitionResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QualityDefinitionResource{}

// QualityDefinitionResource struct for QualityDefinitionResource
type QualityDefinitionResource struct {
	Id *int32 `json:"id,omitempty"`
	Quality *Quality `json:"quality,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	MinSize NullableFloat64 `json:"minSize,omitempty"`
	MaxSize NullableFloat64 `json:"maxSize,omitempty"`
	PreferredSize NullableFloat64 `json:"preferredSize,omitempty"`
}

// NewQualityDefinitionResource instantiates a new QualityDefinitionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualityDefinitionResource() *QualityDefinitionResource {
	this := QualityDefinitionResource{}
	return &this
}

// NewQualityDefinitionResourceWithDefaults instantiates a new QualityDefinitionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityDefinitionResourceWithDefaults() *QualityDefinitionResource {
	this := QualityDefinitionResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QualityDefinitionResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityDefinitionResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QualityDefinitionResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *QualityDefinitionResource) SetId(v int32) {
	o.Id = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *QualityDefinitionResource) GetQuality() Quality {
	if o == nil || IsNil(o.Quality) {
		var ret Quality
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityDefinitionResource) GetQualityOk() (*Quality, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *QualityDefinitionResource) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given Quality and assigns it to the Quality field.
func (o *QualityDefinitionResource) SetQuality(v Quality) {
	o.Quality = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityDefinitionResource) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityDefinitionResource) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *QualityDefinitionResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *QualityDefinitionResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *QualityDefinitionResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *QualityDefinitionResource) UnsetTitle() {
	o.Title.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *QualityDefinitionResource) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityDefinitionResource) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *QualityDefinitionResource) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *QualityDefinitionResource) SetWeight(v int32) {
	o.Weight = &v
}

// GetMinSize returns the MinSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityDefinitionResource) GetMinSize() float64 {
	if o == nil || IsNil(o.MinSize.Get()) {
		var ret float64
		return ret
	}
	return *o.MinSize.Get()
}

// GetMinSizeOk returns a tuple with the MinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityDefinitionResource) GetMinSizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinSize.Get(), o.MinSize.IsSet()
}

// HasMinSize returns a boolean if a field has been set.
func (o *QualityDefinitionResource) HasMinSize() bool {
	if o != nil && o.MinSize.IsSet() {
		return true
	}

	return false
}

// SetMinSize gets a reference to the given NullableFloat64 and assigns it to the MinSize field.
func (o *QualityDefinitionResource) SetMinSize(v float64) {
	o.MinSize.Set(&v)
}
// SetMinSizeNil sets the value for MinSize to be an explicit nil
func (o *QualityDefinitionResource) SetMinSizeNil() {
	o.MinSize.Set(nil)
}

// UnsetMinSize ensures that no value is present for MinSize, not even an explicit nil
func (o *QualityDefinitionResource) UnsetMinSize() {
	o.MinSize.Unset()
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityDefinitionResource) GetMaxSize() float64 {
	if o == nil || IsNil(o.MaxSize.Get()) {
		var ret float64
		return ret
	}
	return *o.MaxSize.Get()
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityDefinitionResource) GetMaxSizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxSize.Get(), o.MaxSize.IsSet()
}

// HasMaxSize returns a boolean if a field has been set.
func (o *QualityDefinitionResource) HasMaxSize() bool {
	if o != nil && o.MaxSize.IsSet() {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given NullableFloat64 and assigns it to the MaxSize field.
func (o *QualityDefinitionResource) SetMaxSize(v float64) {
	o.MaxSize.Set(&v)
}
// SetMaxSizeNil sets the value for MaxSize to be an explicit nil
func (o *QualityDefinitionResource) SetMaxSizeNil() {
	o.MaxSize.Set(nil)
}

// UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
func (o *QualityDefinitionResource) UnsetMaxSize() {
	o.MaxSize.Unset()
}

// GetPreferredSize returns the PreferredSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityDefinitionResource) GetPreferredSize() float64 {
	if o == nil || IsNil(o.PreferredSize.Get()) {
		var ret float64
		return ret
	}
	return *o.PreferredSize.Get()
}

// GetPreferredSizeOk returns a tuple with the PreferredSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityDefinitionResource) GetPreferredSizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredSize.Get(), o.PreferredSize.IsSet()
}

// HasPreferredSize returns a boolean if a field has been set.
func (o *QualityDefinitionResource) HasPreferredSize() bool {
	if o != nil && o.PreferredSize.IsSet() {
		return true
	}

	return false
}

// SetPreferredSize gets a reference to the given NullableFloat64 and assigns it to the PreferredSize field.
func (o *QualityDefinitionResource) SetPreferredSize(v float64) {
	o.PreferredSize.Set(&v)
}
// SetPreferredSizeNil sets the value for PreferredSize to be an explicit nil
func (o *QualityDefinitionResource) SetPreferredSizeNil() {
	o.PreferredSize.Set(nil)
}

// UnsetPreferredSize ensures that no value is present for PreferredSize, not even an explicit nil
func (o *QualityDefinitionResource) UnsetPreferredSize() {
	o.PreferredSize.Unset()
}

func (o QualityDefinitionResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QualityDefinitionResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if o.MinSize.IsSet() {
		toSerialize["minSize"] = o.MinSize.Get()
	}
	if o.MaxSize.IsSet() {
		toSerialize["maxSize"] = o.MaxSize.Get()
	}
	if o.PreferredSize.IsSet() {
		toSerialize["preferredSize"] = o.PreferredSize.Get()
	}
	return toSerialize, nil
}

type NullableQualityDefinitionResource struct {
	value *QualityDefinitionResource
	isSet bool
}

func (v NullableQualityDefinitionResource) Get() *QualityDefinitionResource {
	return v.value
}

func (v *NullableQualityDefinitionResource) Set(val *QualityDefinitionResource) {
	v.value = val
	v.isSet = true
}

func (v NullableQualityDefinitionResource) IsSet() bool {
	return v.isSet
}

func (v *NullableQualityDefinitionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualityDefinitionResource(val *QualityDefinitionResource) *NullableQualityDefinitionResource {
	return &NullableQualityDefinitionResource{value: val, isSet: true}
}

func (v NullableQualityDefinitionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualityDefinitionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



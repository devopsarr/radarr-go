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

// ProfileFormatItemResource struct for ProfileFormatItemResource
type ProfileFormatItemResource struct {
	Id *int32 `json:"id,omitempty"`
	Format *int32 `json:"format,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Score *int32 `json:"score,omitempty"`
}

// NewProfileFormatItemResource instantiates a new ProfileFormatItemResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileFormatItemResource() *ProfileFormatItemResource {
	this := ProfileFormatItemResource{}
	return &this
}

// NewProfileFormatItemResourceWithDefaults instantiates a new ProfileFormatItemResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileFormatItemResourceWithDefaults() *ProfileFormatItemResource {
	this := ProfileFormatItemResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProfileFormatItemResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileFormatItemResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProfileFormatItemResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProfileFormatItemResource) SetId(v int32) {
	o.Id = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ProfileFormatItemResource) GetFormat() int32 {
	if o == nil || isNil(o.Format) {
		var ret int32
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileFormatItemResource) GetFormatOk() (*int32, bool) {
	if o == nil || isNil(o.Format) {
    return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ProfileFormatItemResource) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given int32 and assigns it to the Format field.
func (o *ProfileFormatItemResource) SetFormat(v int32) {
	o.Format = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileFormatItemResource) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileFormatItemResource) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProfileFormatItemResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProfileFormatItemResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProfileFormatItemResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProfileFormatItemResource) UnsetName() {
	o.Name.Unset()
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ProfileFormatItemResource) GetScore() int32 {
	if o == nil || isNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileFormatItemResource) GetScoreOk() (*int32, bool) {
	if o == nil || isNil(o.Score) {
    return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ProfileFormatItemResource) HasScore() bool {
	if o != nil && !isNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *ProfileFormatItemResource) SetScore(v int32) {
	o.Score = &v
}

func (o ProfileFormatItemResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableProfileFormatItemResource struct {
	value *ProfileFormatItemResource
	isSet bool
}

func (v NullableProfileFormatItemResource) Get() *ProfileFormatItemResource {
	return v.value
}

func (v *NullableProfileFormatItemResource) Set(val *ProfileFormatItemResource) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileFormatItemResource) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileFormatItemResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileFormatItemResource(val *ProfileFormatItemResource) *NullableProfileFormatItemResource {
	return &NullableProfileFormatItemResource{value: val, isSet: true}
}

func (v NullableProfileFormatItemResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileFormatItemResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



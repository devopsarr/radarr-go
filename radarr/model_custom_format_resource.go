/*
Radarr

Radarr API docs

API version: v5.16.3.9541
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the CustomFormatResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFormatResource{}

// CustomFormatResource struct for CustomFormatResource
type CustomFormatResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	IncludeCustomFormatWhenRenaming NullableBool `json:"includeCustomFormatWhenRenaming,omitempty"`
	Specifications []CustomFormatSpecificationSchema `json:"specifications,omitempty"`
}

// NewCustomFormatResource instantiates a new CustomFormatResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFormatResource() *CustomFormatResource {
	this := CustomFormatResource{}
	return &this
}

// NewCustomFormatResourceWithDefaults instantiates a new CustomFormatResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFormatResourceWithDefaults() *CustomFormatResource {
	this := CustomFormatResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomFormatResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFormatResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomFormatResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomFormatResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormatResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormatResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CustomFormatResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CustomFormatResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CustomFormatResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CustomFormatResource) UnsetName() {
	o.Name.Unset()
}

// GetIncludeCustomFormatWhenRenaming returns the IncludeCustomFormatWhenRenaming field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormatResource) GetIncludeCustomFormatWhenRenaming() bool {
	if o == nil || IsNil(o.IncludeCustomFormatWhenRenaming.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeCustomFormatWhenRenaming.Get()
}

// GetIncludeCustomFormatWhenRenamingOk returns a tuple with the IncludeCustomFormatWhenRenaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormatResource) GetIncludeCustomFormatWhenRenamingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeCustomFormatWhenRenaming.Get(), o.IncludeCustomFormatWhenRenaming.IsSet()
}

// HasIncludeCustomFormatWhenRenaming returns a boolean if a field has been set.
func (o *CustomFormatResource) HasIncludeCustomFormatWhenRenaming() bool {
	if o != nil && o.IncludeCustomFormatWhenRenaming.IsSet() {
		return true
	}

	return false
}

// SetIncludeCustomFormatWhenRenaming gets a reference to the given NullableBool and assigns it to the IncludeCustomFormatWhenRenaming field.
func (o *CustomFormatResource) SetIncludeCustomFormatWhenRenaming(v bool) {
	o.IncludeCustomFormatWhenRenaming.Set(&v)
}
// SetIncludeCustomFormatWhenRenamingNil sets the value for IncludeCustomFormatWhenRenaming to be an explicit nil
func (o *CustomFormatResource) SetIncludeCustomFormatWhenRenamingNil() {
	o.IncludeCustomFormatWhenRenaming.Set(nil)
}

// UnsetIncludeCustomFormatWhenRenaming ensures that no value is present for IncludeCustomFormatWhenRenaming, not even an explicit nil
func (o *CustomFormatResource) UnsetIncludeCustomFormatWhenRenaming() {
	o.IncludeCustomFormatWhenRenaming.Unset()
}

// GetSpecifications returns the Specifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormatResource) GetSpecifications() []CustomFormatSpecificationSchema {
	if o == nil {
		var ret []CustomFormatSpecificationSchema
		return ret
	}
	return o.Specifications
}

// GetSpecificationsOk returns a tuple with the Specifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormatResource) GetSpecificationsOk() ([]CustomFormatSpecificationSchema, bool) {
	if o == nil || IsNil(o.Specifications) {
		return nil, false
	}
	return o.Specifications, true
}

// HasSpecifications returns a boolean if a field has been set.
func (o *CustomFormatResource) HasSpecifications() bool {
	if o != nil && !IsNil(o.Specifications) {
		return true
	}

	return false
}

// SetSpecifications gets a reference to the given []CustomFormatSpecificationSchema and assigns it to the Specifications field.
func (o *CustomFormatResource) SetSpecifications(v []CustomFormatSpecificationSchema) {
	o.Specifications = v
}

func (o CustomFormatResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFormatResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IncludeCustomFormatWhenRenaming.IsSet() {
		toSerialize["includeCustomFormatWhenRenaming"] = o.IncludeCustomFormatWhenRenaming.Get()
	}
	if o.Specifications != nil {
		toSerialize["specifications"] = o.Specifications
	}
	return toSerialize, nil
}

type NullableCustomFormatResource struct {
	value *CustomFormatResource
	isSet bool
}

func (v NullableCustomFormatResource) Get() *CustomFormatResource {
	return v.value
}

func (v *NullableCustomFormatResource) Set(val *CustomFormatResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFormatResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFormatResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFormatResource(val *CustomFormatResource) *NullableCustomFormatResource {
	return &NullableCustomFormatResource{value: val, isSet: true}
}

func (v NullableCustomFormatResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFormatResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



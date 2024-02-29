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

// checks if the SelectOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectOption{}

// SelectOption struct for SelectOption
type SelectOption struct {
	Value *int32 `json:"value,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Order *int32 `json:"order,omitempty"`
	Hint NullableString `json:"hint,omitempty"`
	DividerAfter *bool `json:"dividerAfter,omitempty"`
}

// NewSelectOption instantiates a new SelectOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectOption() *SelectOption {
	this := SelectOption{}
	return &this
}

// NewSelectOptionWithDefaults instantiates a new SelectOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectOptionWithDefaults() *SelectOption {
	this := SelectOption{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SelectOption) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectOption) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SelectOption) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *SelectOption) SetValue(v int32) {
	o.Value = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectOption) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelectOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SelectOption) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SelectOption) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SelectOption) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SelectOption) UnsetName() {
	o.Name.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *SelectOption) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectOption) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *SelectOption) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *SelectOption) SetOrder(v int32) {
	o.Order = &v
}

// GetHint returns the Hint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SelectOption) GetHint() string {
	if o == nil || IsNil(o.Hint.Get()) {
		var ret string
		return ret
	}
	return *o.Hint.Get()
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SelectOption) GetHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hint.Get(), o.Hint.IsSet()
}

// HasHint returns a boolean if a field has been set.
func (o *SelectOption) HasHint() bool {
	if o != nil && o.Hint.IsSet() {
		return true
	}

	return false
}

// SetHint gets a reference to the given NullableString and assigns it to the Hint field.
func (o *SelectOption) SetHint(v string) {
	o.Hint.Set(&v)
}
// SetHintNil sets the value for Hint to be an explicit nil
func (o *SelectOption) SetHintNil() {
	o.Hint.Set(nil)
}

// UnsetHint ensures that no value is present for Hint, not even an explicit nil
func (o *SelectOption) UnsetHint() {
	o.Hint.Unset()
}

// GetDividerAfter returns the DividerAfter field value if set, zero value otherwise.
func (o *SelectOption) GetDividerAfter() bool {
	if o == nil || IsNil(o.DividerAfter) {
		var ret bool
		return ret
	}
	return *o.DividerAfter
}

// GetDividerAfterOk returns a tuple with the DividerAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectOption) GetDividerAfterOk() (*bool, bool) {
	if o == nil || IsNil(o.DividerAfter) {
		return nil, false
	}
	return o.DividerAfter, true
}

// HasDividerAfter returns a boolean if a field has been set.
func (o *SelectOption) HasDividerAfter() bool {
	if o != nil && !IsNil(o.DividerAfter) {
		return true
	}

	return false
}

// SetDividerAfter gets a reference to the given bool and assigns it to the DividerAfter field.
func (o *SelectOption) SetDividerAfter(v bool) {
	o.DividerAfter = &v
}

func (o SelectOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.Hint.IsSet() {
		toSerialize["hint"] = o.Hint.Get()
	}
	if !IsNil(o.DividerAfter) {
		toSerialize["dividerAfter"] = o.DividerAfter
	}
	return toSerialize, nil
}

type NullableSelectOption struct {
	value *SelectOption
	isSet bool
}

func (v NullableSelectOption) Get() *SelectOption {
	return v.value
}

func (v *NullableSelectOption) Set(val *SelectOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectOption(val *SelectOption) *NullableSelectOption {
	return &NullableSelectOption{value: val, isSet: true}
}

func (v NullableSelectOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



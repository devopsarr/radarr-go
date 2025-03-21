/*
Radarr

Radarr API docs

API version: v5.19.3.9730
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// Modifier the model 'Modifier'
type Modifier string

// List of Modifier
const (
	MODIFIER_NONE Modifier = "none"
	MODIFIER_REGIONAL Modifier = "regional"
	MODIFIER_SCREENER Modifier = "screener"
	MODIFIER_RAWHD Modifier = "rawhd"
	MODIFIER_BRDISK Modifier = "brdisk"
	MODIFIER_REMUX Modifier = "remux"
)

// All allowed values of Modifier enum
var AllowedModifierEnumValues = []Modifier{
	"none",
	"regional",
	"screener",
	"rawhd",
	"brdisk",
	"remux",
}

func (v *Modifier) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Modifier(value)
	for _, existing := range AllowedModifierEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Modifier", value)
}

// NewModifierFromValue returns a pointer to a valid Modifier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModifierFromValue(v string) (*Modifier, error) {
	ev := Modifier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Modifier: valid values are %v", v, AllowedModifierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Modifier) IsValid() bool {
	for _, existing := range AllowedModifierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Modifier value
func (v Modifier) Ptr() *Modifier {
	return &v
}

type NullableModifier struct {
	value *Modifier
	isSet bool
}

func (v NullableModifier) Get() *Modifier {
	return v.value
}

func (v *NullableModifier) Set(val *Modifier) {
	v.value = val
	v.isSet = true
}

func (v NullableModifier) IsSet() bool {
	return v.isSet
}

func (v *NullableModifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifier(val *Modifier) *NullableModifier {
	return &NullableModifier{value: val, isSet: true}
}

func (v NullableModifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
Radarr

Radarr API docs

API version: v5.12.2.9335
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// ExtraFileType the model 'ExtraFileType'
type ExtraFileType string

// List of ExtraFileType
const (
	EXTRAFILETYPE_SUBTITLE ExtraFileType = "subtitle"
	EXTRAFILETYPE_METADATA ExtraFileType = "metadata"
	EXTRAFILETYPE_OTHER ExtraFileType = "other"
)

// All allowed values of ExtraFileType enum
var AllowedExtraFileTypeEnumValues = []ExtraFileType{
	"subtitle",
	"metadata",
	"other",
}

func (v *ExtraFileType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExtraFileType(value)
	for _, existing := range AllowedExtraFileTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExtraFileType", value)
}

// NewExtraFileTypeFromValue returns a pointer to a valid ExtraFileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExtraFileTypeFromValue(v string) (*ExtraFileType, error) {
	ev := ExtraFileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExtraFileType: valid values are %v", v, AllowedExtraFileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExtraFileType) IsValid() bool {
	for _, existing := range AllowedExtraFileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExtraFileType value
func (v ExtraFileType) Ptr() *ExtraFileType {
	return &v
}

type NullableExtraFileType struct {
	value *ExtraFileType
	isSet bool
}

func (v NullableExtraFileType) Get() *ExtraFileType {
	return v.value
}

func (v *NullableExtraFileType) Set(val *ExtraFileType) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraFileType) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraFileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraFileType(val *ExtraFileType) *NullableExtraFileType {
	return &NullableExtraFileType{value: val, isSet: true}
}

func (v NullableExtraFileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraFileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


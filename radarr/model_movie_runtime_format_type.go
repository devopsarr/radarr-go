/*
Radarr

Radarr API docs

API version: v5.16.3.9541
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// MovieRuntimeFormatType the model 'MovieRuntimeFormatType'
type MovieRuntimeFormatType string

// List of MovieRuntimeFormatType
const (
	MOVIERUNTIMEFORMATTYPE_HOURS_MINUTES MovieRuntimeFormatType = "hoursMinutes"
	MOVIERUNTIMEFORMATTYPE_MINUTES MovieRuntimeFormatType = "minutes"
)

// All allowed values of MovieRuntimeFormatType enum
var AllowedMovieRuntimeFormatTypeEnumValues = []MovieRuntimeFormatType{
	"hoursMinutes",
	"minutes",
}

func (v *MovieRuntimeFormatType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MovieRuntimeFormatType(value)
	for _, existing := range AllowedMovieRuntimeFormatTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MovieRuntimeFormatType", value)
}

// NewMovieRuntimeFormatTypeFromValue returns a pointer to a valid MovieRuntimeFormatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMovieRuntimeFormatTypeFromValue(v string) (*MovieRuntimeFormatType, error) {
	ev := MovieRuntimeFormatType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MovieRuntimeFormatType: valid values are %v", v, AllowedMovieRuntimeFormatTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MovieRuntimeFormatType) IsValid() bool {
	for _, existing := range AllowedMovieRuntimeFormatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MovieRuntimeFormatType value
func (v MovieRuntimeFormatType) Ptr() *MovieRuntimeFormatType {
	return &v
}

type NullableMovieRuntimeFormatType struct {
	value *MovieRuntimeFormatType
	isSet bool
}

func (v NullableMovieRuntimeFormatType) Get() *MovieRuntimeFormatType {
	return v.value
}

func (v *NullableMovieRuntimeFormatType) Set(val *MovieRuntimeFormatType) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieRuntimeFormatType) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieRuntimeFormatType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieRuntimeFormatType(val *MovieRuntimeFormatType) *NullableMovieRuntimeFormatType {
	return &NullableMovieRuntimeFormatType{value: val, isSet: true}
}

func (v NullableMovieRuntimeFormatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieRuntimeFormatType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


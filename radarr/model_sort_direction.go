/*
Radarr

Radarr API docs

API version: v5.10.4.9218
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// SortDirection the model 'SortDirection'
type SortDirection string

// List of SortDirection
const (
	SORTDIRECTION_DEFAULT SortDirection = "default"
	SORTDIRECTION_ASCENDING SortDirection = "ascending"
	SORTDIRECTION_DESCENDING SortDirection = "descending"
)

// All allowed values of SortDirection enum
var AllowedSortDirectionEnumValues = []SortDirection{
	"default",
	"ascending",
	"descending",
}

func (v *SortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SortDirection(value)
	for _, existing := range AllowedSortDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortDirection", value)
}

// NewSortDirectionFromValue returns a pointer to a valid SortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortDirectionFromValue(v string) (*SortDirection, error) {
	ev := SortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortDirection: valid values are %v", v, AllowedSortDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortDirection) IsValid() bool {
	for _, existing := range AllowedSortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortDirection value
func (v SortDirection) Ptr() *SortDirection {
	return &v
}

type NullableSortDirection struct {
	value *SortDirection
	isSet bool
}

func (v NullableSortDirection) Get() *SortDirection {
	return v.value
}

func (v *NullableSortDirection) Set(val *SortDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableSortDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableSortDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortDirection(val *SortDirection) *NullableSortDirection {
	return &NullableSortDirection{value: val, isSet: true}
}

func (v NullableSortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


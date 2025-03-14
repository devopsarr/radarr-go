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

// SourceType the model 'SourceType'
type SourceType string

// List of SourceType
const (
	SOURCETYPE_TMDB SourceType = "tmdb"
	SOURCETYPE_MAPPINGS SourceType = "mappings"
	SOURCETYPE_USER SourceType = "user"
	SOURCETYPE_INDEXER SourceType = "indexer"
)

// All allowed values of SourceType enum
var AllowedSourceTypeEnumValues = []SourceType{
	"tmdb",
	"mappings",
	"user",
	"indexer",
}

func (v *SourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SourceType(value)
	for _, existing := range AllowedSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SourceType", value)
}

// NewSourceTypeFromValue returns a pointer to a valid SourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSourceTypeFromValue(v string) (*SourceType, error) {
	ev := SourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SourceType: valid values are %v", v, AllowedSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SourceType) IsValid() bool {
	for _, existing := range AllowedSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SourceType value
func (v SourceType) Ptr() *SourceType {
	return &v
}

type NullableSourceType struct {
	value *SourceType
	isSet bool
}

func (v NullableSourceType) Get() *SourceType {
	return v.value
}

func (v *NullableSourceType) Set(val *SourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceType(val *SourceType) *NullableSourceType {
	return &NullableSourceType{value: val, isSet: true}
}

func (v NullableSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


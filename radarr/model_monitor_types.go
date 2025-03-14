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

// MonitorTypes the model 'MonitorTypes'
type MonitorTypes string

// List of MonitorTypes
const (
	MONITORTYPES_MOVIE_ONLY MonitorTypes = "movieOnly"
	MONITORTYPES_MOVIE_AND_COLLECTION MonitorTypes = "movieAndCollection"
	MONITORTYPES_NONE MonitorTypes = "none"
)

// All allowed values of MonitorTypes enum
var AllowedMonitorTypesEnumValues = []MonitorTypes{
	"movieOnly",
	"movieAndCollection",
	"none",
}

func (v *MonitorTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MonitorTypes(value)
	for _, existing := range AllowedMonitorTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MonitorTypes", value)
}

// NewMonitorTypesFromValue returns a pointer to a valid MonitorTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitorTypesFromValue(v string) (*MonitorTypes, error) {
	ev := MonitorTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitorTypes: valid values are %v", v, AllowedMonitorTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitorTypes) IsValid() bool {
	for _, existing := range AllowedMonitorTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitorTypes value
func (v MonitorTypes) Ptr() *MonitorTypes {
	return &v
}

type NullableMonitorTypes struct {
	value *MonitorTypes
	isSet bool
}

func (v NullableMonitorTypes) Get() *MonitorTypes {
	return v.value
}

func (v *NullableMonitorTypes) Set(val *MonitorTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorTypes(val *MonitorTypes) *NullableMonitorTypes {
	return &NullableMonitorTypes{value: val, isSet: true}
}

func (v NullableMonitorTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


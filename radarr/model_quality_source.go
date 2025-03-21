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

// QualitySource the model 'QualitySource'
type QualitySource string

// List of QualitySource
const (
	QUALITYSOURCE_UNKNOWN QualitySource = "unknown"
	QUALITYSOURCE_CAM QualitySource = "cam"
	QUALITYSOURCE_TELESYNC QualitySource = "telesync"
	QUALITYSOURCE_TELECINE QualitySource = "telecine"
	QUALITYSOURCE_WORKPRINT QualitySource = "workprint"
	QUALITYSOURCE_DVD QualitySource = "dvd"
	QUALITYSOURCE_TV QualitySource = "tv"
	QUALITYSOURCE_WEBDL QualitySource = "webdl"
	QUALITYSOURCE_WEBRIP QualitySource = "webrip"
	QUALITYSOURCE_BLURAY QualitySource = "bluray"
)

// All allowed values of QualitySource enum
var AllowedQualitySourceEnumValues = []QualitySource{
	"unknown",
	"cam",
	"telesync",
	"telecine",
	"workprint",
	"dvd",
	"tv",
	"webdl",
	"webrip",
	"bluray",
}

func (v *QualitySource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QualitySource(value)
	for _, existing := range AllowedQualitySourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QualitySource", value)
}

// NewQualitySourceFromValue returns a pointer to a valid QualitySource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQualitySourceFromValue(v string) (*QualitySource, error) {
	ev := QualitySource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QualitySource: valid values are %v", v, AllowedQualitySourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QualitySource) IsValid() bool {
	for _, existing := range AllowedQualitySourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QualitySource value
func (v QualitySource) Ptr() *QualitySource {
	return &v
}

type NullableQualitySource struct {
	value *QualitySource
	isSet bool
}

func (v NullableQualitySource) Get() *QualitySource {
	return v.value
}

func (v *NullableQualitySource) Set(val *QualitySource) {
	v.value = val
	v.isSet = true
}

func (v NullableQualitySource) IsSet() bool {
	return v.isSet
}

func (v *NullableQualitySource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualitySource(val *QualitySource) *NullableQualitySource {
	return &NullableQualitySource{value: val, isSet: true}
}

func (v NullableQualitySource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualitySource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
Radarr

Radarr API docs

API version: v5.3.6.8612
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// MediaCoverTypes the model 'MediaCoverTypes'
type MediaCoverTypes string

// List of MediaCoverTypes
const (
	MEDIACOVERTYPES_UNKNOWN MediaCoverTypes = "unknown"
	MEDIACOVERTYPES_POSTER MediaCoverTypes = "poster"
	MEDIACOVERTYPES_BANNER MediaCoverTypes = "banner"
	MEDIACOVERTYPES_FANART MediaCoverTypes = "fanart"
	MEDIACOVERTYPES_SCREENSHOT MediaCoverTypes = "screenshot"
	MEDIACOVERTYPES_HEADSHOT MediaCoverTypes = "headshot"
	MEDIACOVERTYPES_CLEARLOGO MediaCoverTypes = "clearlogo"
)

// All allowed values of MediaCoverTypes enum
var AllowedMediaCoverTypesEnumValues = []MediaCoverTypes{
	"unknown",
	"poster",
	"banner",
	"fanart",
	"screenshot",
	"headshot",
	"clearlogo",
}

func (v *MediaCoverTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MediaCoverTypes(value)
	for _, existing := range AllowedMediaCoverTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MediaCoverTypes", value)
}

// NewMediaCoverTypesFromValue returns a pointer to a valid MediaCoverTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaCoverTypesFromValue(v string) (*MediaCoverTypes, error) {
	ev := MediaCoverTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaCoverTypes: valid values are %v", v, AllowedMediaCoverTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaCoverTypes) IsValid() bool {
	for _, existing := range AllowedMediaCoverTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaCoverTypes value
func (v MediaCoverTypes) Ptr() *MediaCoverTypes {
	return &v
}

type NullableMediaCoverTypes struct {
	value *MediaCoverTypes
	isSet bool
}

func (v NullableMediaCoverTypes) Get() *MediaCoverTypes {
	return v.value
}

func (v *NullableMediaCoverTypes) Set(val *MediaCoverTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaCoverTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaCoverTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaCoverTypes(val *MediaCoverTypes) *NullableMediaCoverTypes {
	return &NullableMediaCoverTypes{value: val, isSet: true}
}

func (v NullableMediaCoverTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaCoverTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


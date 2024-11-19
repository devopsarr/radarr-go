/*
Radarr

Radarr API docs

API version: v5.15.1.9463
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// PrivacyLevel the model 'PrivacyLevel'
type PrivacyLevel string

// List of PrivacyLevel
const (
	PRIVACYLEVEL_NORMAL PrivacyLevel = "normal"
	PRIVACYLEVEL_PASSWORD PrivacyLevel = "password"
	PRIVACYLEVEL_API_KEY PrivacyLevel = "apiKey"
	PRIVACYLEVEL_USER_NAME PrivacyLevel = "userName"
)

// All allowed values of PrivacyLevel enum
var AllowedPrivacyLevelEnumValues = []PrivacyLevel{
	"normal",
	"password",
	"apiKey",
	"userName",
}

func (v *PrivacyLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrivacyLevel(value)
	for _, existing := range AllowedPrivacyLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrivacyLevel", value)
}

// NewPrivacyLevelFromValue returns a pointer to a valid PrivacyLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrivacyLevelFromValue(v string) (*PrivacyLevel, error) {
	ev := PrivacyLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrivacyLevel: valid values are %v", v, AllowedPrivacyLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrivacyLevel) IsValid() bool {
	for _, existing := range AllowedPrivacyLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrivacyLevel value
func (v PrivacyLevel) Ptr() *PrivacyLevel {
	return &v
}

type NullablePrivacyLevel struct {
	value *PrivacyLevel
	isSet bool
}

func (v NullablePrivacyLevel) Get() *PrivacyLevel {
	return v.value
}

func (v *NullablePrivacyLevel) Set(val *PrivacyLevel) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivacyLevel) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivacyLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivacyLevel(val *PrivacyLevel) *NullablePrivacyLevel {
	return &NullablePrivacyLevel{value: val, isSet: true}
}

func (v NullablePrivacyLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivacyLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


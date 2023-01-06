/*
Radarr

Radarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// AuthenticationType the model 'AuthenticationType'
type AuthenticationType string

// List of AuthenticationType
const (
	AUTHENTICATIONTYPE_NONE AuthenticationType = "none"
	AUTHENTICATIONTYPE_BASIC AuthenticationType = "basic"
	AUTHENTICATIONTYPE_FORMS AuthenticationType = "forms"
)

// All allowed values of AuthenticationType enum
var AllowedAuthenticationTypeEnumValues = []AuthenticationType{
	"none",
	"basic",
	"forms",
}

func (v *AuthenticationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticationType(value)
	for _, existing := range AllowedAuthenticationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticationType", value)
}

// NewAuthenticationTypeFromValue returns a pointer to a valid AuthenticationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticationTypeFromValue(v string) (*AuthenticationType, error) {
	ev := AuthenticationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticationType: valid values are %v", v, AllowedAuthenticationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticationType) IsValid() bool {
	for _, existing := range AllowedAuthenticationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthenticationType value
func (v AuthenticationType) Ptr() *AuthenticationType {
	return &v
}

type NullableAuthenticationType struct {
	value *AuthenticationType
	isSet bool
}

func (v NullableAuthenticationType) Get() *AuthenticationType {
	return v.value
}

func (v *NullableAuthenticationType) Set(val *AuthenticationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationType(val *AuthenticationType) *NullableAuthenticationType {
	return &NullableAuthenticationType{value: val, isSet: true}
}

func (v NullableAuthenticationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


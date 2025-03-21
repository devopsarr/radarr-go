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

// TMDbCountryCode the model 'TMDbCountryCode'
type TMDbCountryCode string

// List of TMDbCountryCode
const (
	TMDBCOUNTRYCODE_AU TMDbCountryCode = "au"
	TMDBCOUNTRYCODE_BR TMDbCountryCode = "br"
	TMDBCOUNTRYCODE_CA TMDbCountryCode = "ca"
	TMDBCOUNTRYCODE_FR TMDbCountryCode = "fr"
	TMDBCOUNTRYCODE_DE TMDbCountryCode = "de"
	TMDBCOUNTRYCODE_GB TMDbCountryCode = "gb"
	TMDBCOUNTRYCODE_IN TMDbCountryCode = "in"
	TMDBCOUNTRYCODE_IE TMDbCountryCode = "ie"
	TMDBCOUNTRYCODE_IT TMDbCountryCode = "it"
	TMDBCOUNTRYCODE_NZ TMDbCountryCode = "nz"
	TMDBCOUNTRYCODE_RO TMDbCountryCode = "ro"
	TMDBCOUNTRYCODE_ES TMDbCountryCode = "es"
	TMDBCOUNTRYCODE_US TMDbCountryCode = "us"
)

// All allowed values of TMDbCountryCode enum
var AllowedTMDbCountryCodeEnumValues = []TMDbCountryCode{
	"au",
	"br",
	"ca",
	"fr",
	"de",
	"gb",
	"in",
	"ie",
	"it",
	"nz",
	"ro",
	"es",
	"us",
}

func (v *TMDbCountryCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TMDbCountryCode(value)
	for _, existing := range AllowedTMDbCountryCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TMDbCountryCode", value)
}

// NewTMDbCountryCodeFromValue returns a pointer to a valid TMDbCountryCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTMDbCountryCodeFromValue(v string) (*TMDbCountryCode, error) {
	ev := TMDbCountryCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TMDbCountryCode: valid values are %v", v, AllowedTMDbCountryCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TMDbCountryCode) IsValid() bool {
	for _, existing := range AllowedTMDbCountryCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TMDbCountryCode value
func (v TMDbCountryCode) Ptr() *TMDbCountryCode {
	return &v
}

type NullableTMDbCountryCode struct {
	value *TMDbCountryCode
	isSet bool
}

func (v NullableTMDbCountryCode) Get() *TMDbCountryCode {
	return v.value
}

func (v *NullableTMDbCountryCode) Set(val *TMDbCountryCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTMDbCountryCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTMDbCountryCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTMDbCountryCode(val *TMDbCountryCode) *NullableTMDbCountryCode {
	return &NullableTMDbCountryCode{value: val, isSet: true}
}

func (v NullableTMDbCountryCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTMDbCountryCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
Radarr

Radarr API docs

API version: v5.17.2.9580
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// CreditType the model 'CreditType'
type CreditType string

// List of CreditType
const (
	CREDITTYPE_CAST CreditType = "cast"
	CREDITTYPE_CREW CreditType = "crew"
)

// All allowed values of CreditType enum
var AllowedCreditTypeEnumValues = []CreditType{
	"cast",
	"crew",
}

func (v *CreditType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreditType(value)
	for _, existing := range AllowedCreditTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreditType", value)
}

// NewCreditTypeFromValue returns a pointer to a valid CreditType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditTypeFromValue(v string) (*CreditType, error) {
	ev := CreditType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreditType: valid values are %v", v, AllowedCreditTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditType) IsValid() bool {
	for _, existing := range AllowedCreditTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditType value
func (v CreditType) Ptr() *CreditType {
	return &v
}

type NullableCreditType struct {
	value *CreditType
	isSet bool
}

func (v NullableCreditType) Get() *CreditType {
	return v.value
}

func (v *NullableCreditType) Set(val *CreditType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditType(val *CreditType) *NullableCreditType {
	return &NullableCreditType{value: val, isSet: true}
}

func (v NullableCreditType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


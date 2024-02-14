/*
Radarr

Radarr API docs

API version: v5.2.6.8376
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// CertificateValidationType the model 'CertificateValidationType'
type CertificateValidationType string

// List of CertificateValidationType
const (
	CERTIFICATEVALIDATIONTYPE_ENABLED CertificateValidationType = "enabled"
	CERTIFICATEVALIDATIONTYPE_DISABLED_FOR_LOCAL_ADDRESSES CertificateValidationType = "disabledForLocalAddresses"
	CERTIFICATEVALIDATIONTYPE_DISABLED CertificateValidationType = "disabled"
)

// All allowed values of CertificateValidationType enum
var AllowedCertificateValidationTypeEnumValues = []CertificateValidationType{
	"enabled",
	"disabledForLocalAddresses",
	"disabled",
}

func (v *CertificateValidationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CertificateValidationType(value)
	for _, existing := range AllowedCertificateValidationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CertificateValidationType", value)
}

// NewCertificateValidationTypeFromValue returns a pointer to a valid CertificateValidationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCertificateValidationTypeFromValue(v string) (*CertificateValidationType, error) {
	ev := CertificateValidationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CertificateValidationType: valid values are %v", v, AllowedCertificateValidationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CertificateValidationType) IsValid() bool {
	for _, existing := range AllowedCertificateValidationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CertificateValidationType value
func (v CertificateValidationType) Ptr() *CertificateValidationType {
	return &v
}

type NullableCertificateValidationType struct {
	value *CertificateValidationType
	isSet bool
}

func (v NullableCertificateValidationType) Get() *CertificateValidationType {
	return v.value
}

func (v *NullableCertificateValidationType) Set(val *CertificateValidationType) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateValidationType) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateValidationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateValidationType(val *CertificateValidationType) *NullableCertificateValidationType {
	return &NullableCertificateValidationType{value: val, isSet: true}
}

func (v NullableCertificateValidationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateValidationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


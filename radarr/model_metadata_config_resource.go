/*
Radarr

Radarr API docs

API version: v5.2.6.8376
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the MetadataConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataConfigResource{}

// MetadataConfigResource struct for MetadataConfigResource
type MetadataConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	CertificationCountry *TMDbCountryCode `json:"certificationCountry,omitempty"`
}

// NewMetadataConfigResource instantiates a new MetadataConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataConfigResource() *MetadataConfigResource {
	this := MetadataConfigResource{}
	return &this
}

// NewMetadataConfigResourceWithDefaults instantiates a new MetadataConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataConfigResourceWithDefaults() *MetadataConfigResource {
	this := MetadataConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetadataConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetadataConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MetadataConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetCertificationCountry returns the CertificationCountry field value if set, zero value otherwise.
func (o *MetadataConfigResource) GetCertificationCountry() TMDbCountryCode {
	if o == nil || IsNil(o.CertificationCountry) {
		var ret TMDbCountryCode
		return ret
	}
	return *o.CertificationCountry
}

// GetCertificationCountryOk returns a tuple with the CertificationCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataConfigResource) GetCertificationCountryOk() (*TMDbCountryCode, bool) {
	if o == nil || IsNil(o.CertificationCountry) {
		return nil, false
	}
	return o.CertificationCountry, true
}

// HasCertificationCountry returns a boolean if a field has been set.
func (o *MetadataConfigResource) HasCertificationCountry() bool {
	if o != nil && !IsNil(o.CertificationCountry) {
		return true
	}

	return false
}

// SetCertificationCountry gets a reference to the given TMDbCountryCode and assigns it to the CertificationCountry field.
func (o *MetadataConfigResource) SetCertificationCountry(v TMDbCountryCode) {
	o.CertificationCountry = &v
}

func (o MetadataConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CertificationCountry) {
		toSerialize["certificationCountry"] = o.CertificationCountry
	}
	return toSerialize, nil
}

type NullableMetadataConfigResource struct {
	value *MetadataConfigResource
	isSet bool
}

func (v NullableMetadataConfigResource) Get() *MetadataConfigResource {
	return v.value
}

func (v *NullableMetadataConfigResource) Set(val *MetadataConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataConfigResource(val *MetadataConfigResource) *NullableMetadataConfigResource {
	return &NullableMetadataConfigResource{value: val, isSet: true}
}

func (v NullableMetadataConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



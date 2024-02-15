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

// checks if the PingResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingResource{}

// PingResource struct for PingResource
type PingResource struct {
	Status NullableString `json:"status,omitempty"`
}

// NewPingResource instantiates a new PingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingResource() *PingResource {
	this := PingResource{}
	return &this
}

// NewPingResourceWithDefaults instantiates a new PingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingResourceWithDefaults() *PingResource {
	this := PingResource{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PingResource) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PingResource) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PingResource) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *PingResource) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *PingResource) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PingResource) UnsetStatus() {
	o.Status.Unset()
}

func (o PingResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return toSerialize, nil
}

type NullablePingResource struct {
	value *PingResource
	isSet bool
}

func (v NullablePingResource) Get() *PingResource {
	return v.value
}

func (v *NullablePingResource) Set(val *PingResource) {
	v.value = val
	v.isSet = true
}

func (v NullablePingResource) IsSet() bool {
	return v.isSet
}

func (v *NullablePingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingResource(val *PingResource) *NullablePingResource {
	return &NullablePingResource{value: val, isSet: true}
}

func (v NullablePingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



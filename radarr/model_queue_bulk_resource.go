/*
Radarr

Radarr API docs

API version: v5.12.2.9335
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the QueueBulkResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueBulkResource{}

// QueueBulkResource struct for QueueBulkResource
type QueueBulkResource struct {
	Ids []int32 `json:"ids,omitempty"`
}

// NewQueueBulkResource instantiates a new QueueBulkResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueBulkResource() *QueueBulkResource {
	this := QueueBulkResource{}
	return &this
}

// NewQueueBulkResourceWithDefaults instantiates a new QueueBulkResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueBulkResourceWithDefaults() *QueueBulkResource {
	this := QueueBulkResource{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueueBulkResource) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueueBulkResource) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *QueueBulkResource) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *QueueBulkResource) SetIds(v []int32) {
	o.Ids = v
}

func (o QueueBulkResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueBulkResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableQueueBulkResource struct {
	value *QueueBulkResource
	isSet bool
}

func (v NullableQueueBulkResource) Get() *QueueBulkResource {
	return v.value
}

func (v *NullableQueueBulkResource) Set(val *QueueBulkResource) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueBulkResource) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueBulkResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueBulkResource(val *QueueBulkResource) *NullableQueueBulkResource {
	return &NullableQueueBulkResource{value: val, isSet: true}
}

func (v NullableQueueBulkResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueBulkResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



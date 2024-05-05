/*
Radarr

Radarr API docs

API version: v5.4.6.8723
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the RatingChild type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RatingChild{}

// RatingChild struct for RatingChild
type RatingChild struct {
	Votes *int32 `json:"votes,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Type *RatingType `json:"type,omitempty"`
}

// NewRatingChild instantiates a new RatingChild object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatingChild() *RatingChild {
	this := RatingChild{}
	return &this
}

// NewRatingChildWithDefaults instantiates a new RatingChild object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatingChildWithDefaults() *RatingChild {
	this := RatingChild{}
	return &this
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *RatingChild) GetVotes() int32 {
	if o == nil || IsNil(o.Votes) {
		var ret int32
		return ret
	}
	return *o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingChild) GetVotesOk() (*int32, bool) {
	if o == nil || IsNil(o.Votes) {
		return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *RatingChild) HasVotes() bool {
	if o != nil && !IsNil(o.Votes) {
		return true
	}

	return false
}

// SetVotes gets a reference to the given int32 and assigns it to the Votes field.
func (o *RatingChild) SetVotes(v int32) {
	o.Votes = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RatingChild) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingChild) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RatingChild) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *RatingChild) SetValue(v float64) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RatingChild) GetType() RatingType {
	if o == nil || IsNil(o.Type) {
		var ret RatingType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingChild) GetTypeOk() (*RatingType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RatingChild) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RatingType and assigns it to the Type field.
func (o *RatingChild) SetType(v RatingType) {
	o.Type = &v
}

func (o RatingChild) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RatingChild) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Votes) {
		toSerialize["votes"] = o.Votes
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableRatingChild struct {
	value *RatingChild
	isSet bool
}

func (v NullableRatingChild) Get() *RatingChild {
	return v.value
}

func (v *NullableRatingChild) Set(val *RatingChild) {
	v.value = val
	v.isSet = true
}

func (v NullableRatingChild) IsSet() bool {
	return v.isSet
}

func (v *NullableRatingChild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatingChild(val *RatingChild) *NullableRatingChild {
	return &NullableRatingChild{value: val, isSet: true}
}

func (v NullableRatingChild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatingChild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



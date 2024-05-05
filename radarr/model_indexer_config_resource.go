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

// checks if the IndexerConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndexerConfigResource{}

// IndexerConfigResource struct for IndexerConfigResource
type IndexerConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	MinimumAge *int32 `json:"minimumAge,omitempty"`
	MaximumSize *int32 `json:"maximumSize,omitempty"`
	Retention *int32 `json:"retention,omitempty"`
	RssSyncInterval *int32 `json:"rssSyncInterval,omitempty"`
	PreferIndexerFlags *bool `json:"preferIndexerFlags,omitempty"`
	AvailabilityDelay *int32 `json:"availabilityDelay,omitempty"`
	AllowHardcodedSubs *bool `json:"allowHardcodedSubs,omitempty"`
	WhitelistedHardcodedSubs NullableString `json:"whitelistedHardcodedSubs,omitempty"`
}

// NewIndexerConfigResource instantiates a new IndexerConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexerConfigResource() *IndexerConfigResource {
	this := IndexerConfigResource{}
	return &this
}

// NewIndexerConfigResourceWithDefaults instantiates a new IndexerConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexerConfigResourceWithDefaults() *IndexerConfigResource {
	this := IndexerConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IndexerConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IndexerConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IndexerConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetMinimumAge returns the MinimumAge field value if set, zero value otherwise.
func (o *IndexerConfigResource) GetMinimumAge() int32 {
	if o == nil || IsNil(o.MinimumAge) {
		var ret int32
		return ret
	}
	return *o.MinimumAge
}

// GetMinimumAgeOk returns a tuple with the MinimumAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerConfigResource) GetMinimumAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumAge) {
		return nil, false
	}
	return o.MinimumAge, true
}

// HasMinimumAge returns a boolean if a field has been set.
func (o *IndexerConfigResource) HasMinimumAge() bool {
	if o != nil && !IsNil(o.MinimumAge) {
		return true
	}

	return false
}

// SetMinimumAge gets a reference to the given int32 and assigns it to the MinimumAge field.
func (o *IndexerConfigResource) SetMinimumAge(v int32) {
	o.MinimumAge = &v
}

// GetMaximumSize returns the MaximumSize field value if set, zero value otherwise.
func (o *IndexerConfigResource) GetMaximumSize() int32 {
	if o == nil || IsNil(o.MaximumSize) {
		var ret int32
		return ret
	}
	return *o.MaximumSize
}

// GetMaximumSizeOk returns a tuple with the MaximumSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerConfigResource) GetMaximumSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumSize) {
		return nil, false
	}
	return o.MaximumSize, true
}

// HasMaximumSize returns a boolean if a field has been set.
func (o *IndexerConfigResource) HasMaximumSize() bool {
	if o != nil && !IsNil(o.MaximumSize) {
		return true
	}

	return false
}

// SetMaximumSize gets a reference to the given int32 and assigns it to the MaximumSize field.
func (o *IndexerConfigResource) SetMaximumSize(v int32) {
	o.MaximumSize = &v
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *IndexerConfigResource) GetRetention() int32 {
	if o == nil || IsNil(o.Retention) {
		var ret int32
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerConfigResource) GetRetentionOk() (*int32, bool) {
	if o == nil || IsNil(o.Retention) {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *IndexerConfigResource) HasRetention() bool {
	if o != nil && !IsNil(o.Retention) {
		return true
	}

	return false
}

// SetRetention gets a reference to the given int32 and assigns it to the Retention field.
func (o *IndexerConfigResource) SetRetention(v int32) {
	o.Retention = &v
}

// GetRssSyncInterval returns the RssSyncInterval field value if set, zero value otherwise.
func (o *IndexerConfigResource) GetRssSyncInterval() int32 {
	if o == nil || IsNil(o.RssSyncInterval) {
		var ret int32
		return ret
	}
	return *o.RssSyncInterval
}

// GetRssSyncIntervalOk returns a tuple with the RssSyncInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerConfigResource) GetRssSyncIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.RssSyncInterval) {
		return nil, false
	}
	return o.RssSyncInterval, true
}

// HasRssSyncInterval returns a boolean if a field has been set.
func (o *IndexerConfigResource) HasRssSyncInterval() bool {
	if o != nil && !IsNil(o.RssSyncInterval) {
		return true
	}

	return false
}

// SetRssSyncInterval gets a reference to the given int32 and assigns it to the RssSyncInterval field.
func (o *IndexerConfigResource) SetRssSyncInterval(v int32) {
	o.RssSyncInterval = &v
}

// GetPreferIndexerFlags returns the PreferIndexerFlags field value if set, zero value otherwise.
func (o *IndexerConfigResource) GetPreferIndexerFlags() bool {
	if o == nil || IsNil(o.PreferIndexerFlags) {
		var ret bool
		return ret
	}
	return *o.PreferIndexerFlags
}

// GetPreferIndexerFlagsOk returns a tuple with the PreferIndexerFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerConfigResource) GetPreferIndexerFlagsOk() (*bool, bool) {
	if o == nil || IsNil(o.PreferIndexerFlags) {
		return nil, false
	}
	return o.PreferIndexerFlags, true
}

// HasPreferIndexerFlags returns a boolean if a field has been set.
func (o *IndexerConfigResource) HasPreferIndexerFlags() bool {
	if o != nil && !IsNil(o.PreferIndexerFlags) {
		return true
	}

	return false
}

// SetPreferIndexerFlags gets a reference to the given bool and assigns it to the PreferIndexerFlags field.
func (o *IndexerConfigResource) SetPreferIndexerFlags(v bool) {
	o.PreferIndexerFlags = &v
}

// GetAvailabilityDelay returns the AvailabilityDelay field value if set, zero value otherwise.
func (o *IndexerConfigResource) GetAvailabilityDelay() int32 {
	if o == nil || IsNil(o.AvailabilityDelay) {
		var ret int32
		return ret
	}
	return *o.AvailabilityDelay
}

// GetAvailabilityDelayOk returns a tuple with the AvailabilityDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerConfigResource) GetAvailabilityDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailabilityDelay) {
		return nil, false
	}
	return o.AvailabilityDelay, true
}

// HasAvailabilityDelay returns a boolean if a field has been set.
func (o *IndexerConfigResource) HasAvailabilityDelay() bool {
	if o != nil && !IsNil(o.AvailabilityDelay) {
		return true
	}

	return false
}

// SetAvailabilityDelay gets a reference to the given int32 and assigns it to the AvailabilityDelay field.
func (o *IndexerConfigResource) SetAvailabilityDelay(v int32) {
	o.AvailabilityDelay = &v
}

// GetAllowHardcodedSubs returns the AllowHardcodedSubs field value if set, zero value otherwise.
func (o *IndexerConfigResource) GetAllowHardcodedSubs() bool {
	if o == nil || IsNil(o.AllowHardcodedSubs) {
		var ret bool
		return ret
	}
	return *o.AllowHardcodedSubs
}

// GetAllowHardcodedSubsOk returns a tuple with the AllowHardcodedSubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexerConfigResource) GetAllowHardcodedSubsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowHardcodedSubs) {
		return nil, false
	}
	return o.AllowHardcodedSubs, true
}

// HasAllowHardcodedSubs returns a boolean if a field has been set.
func (o *IndexerConfigResource) HasAllowHardcodedSubs() bool {
	if o != nil && !IsNil(o.AllowHardcodedSubs) {
		return true
	}

	return false
}

// SetAllowHardcodedSubs gets a reference to the given bool and assigns it to the AllowHardcodedSubs field.
func (o *IndexerConfigResource) SetAllowHardcodedSubs(v bool) {
	o.AllowHardcodedSubs = &v
}

// GetWhitelistedHardcodedSubs returns the WhitelistedHardcodedSubs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexerConfigResource) GetWhitelistedHardcodedSubs() string {
	if o == nil || IsNil(o.WhitelistedHardcodedSubs.Get()) {
		var ret string
		return ret
	}
	return *o.WhitelistedHardcodedSubs.Get()
}

// GetWhitelistedHardcodedSubsOk returns a tuple with the WhitelistedHardcodedSubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexerConfigResource) GetWhitelistedHardcodedSubsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WhitelistedHardcodedSubs.Get(), o.WhitelistedHardcodedSubs.IsSet()
}

// HasWhitelistedHardcodedSubs returns a boolean if a field has been set.
func (o *IndexerConfigResource) HasWhitelistedHardcodedSubs() bool {
	if o != nil && o.WhitelistedHardcodedSubs.IsSet() {
		return true
	}

	return false
}

// SetWhitelistedHardcodedSubs gets a reference to the given NullableString and assigns it to the WhitelistedHardcodedSubs field.
func (o *IndexerConfigResource) SetWhitelistedHardcodedSubs(v string) {
	o.WhitelistedHardcodedSubs.Set(&v)
}
// SetWhitelistedHardcodedSubsNil sets the value for WhitelistedHardcodedSubs to be an explicit nil
func (o *IndexerConfigResource) SetWhitelistedHardcodedSubsNil() {
	o.WhitelistedHardcodedSubs.Set(nil)
}

// UnsetWhitelistedHardcodedSubs ensures that no value is present for WhitelistedHardcodedSubs, not even an explicit nil
func (o *IndexerConfigResource) UnsetWhitelistedHardcodedSubs() {
	o.WhitelistedHardcodedSubs.Unset()
}

func (o IndexerConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexerConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MinimumAge) {
		toSerialize["minimumAge"] = o.MinimumAge
	}
	if !IsNil(o.MaximumSize) {
		toSerialize["maximumSize"] = o.MaximumSize
	}
	if !IsNil(o.Retention) {
		toSerialize["retention"] = o.Retention
	}
	if !IsNil(o.RssSyncInterval) {
		toSerialize["rssSyncInterval"] = o.RssSyncInterval
	}
	if !IsNil(o.PreferIndexerFlags) {
		toSerialize["preferIndexerFlags"] = o.PreferIndexerFlags
	}
	if !IsNil(o.AvailabilityDelay) {
		toSerialize["availabilityDelay"] = o.AvailabilityDelay
	}
	if !IsNil(o.AllowHardcodedSubs) {
		toSerialize["allowHardcodedSubs"] = o.AllowHardcodedSubs
	}
	if o.WhitelistedHardcodedSubs.IsSet() {
		toSerialize["whitelistedHardcodedSubs"] = o.WhitelistedHardcodedSubs.Get()
	}
	return toSerialize, nil
}

type NullableIndexerConfigResource struct {
	value *IndexerConfigResource
	isSet bool
}

func (v NullableIndexerConfigResource) Get() *IndexerConfigResource {
	return v.value
}

func (v *NullableIndexerConfigResource) Set(val *IndexerConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexerConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexerConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexerConfigResource(val *IndexerConfigResource) *NullableIndexerConfigResource {
	return &NullableIndexerConfigResource{value: val, isSet: true}
}

func (v NullableIndexerConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexerConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



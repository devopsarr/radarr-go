/*
Radarr

Radarr API docs

API version: v5.19.3.9730
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the DownloadClientBulkResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadClientBulkResource{}

// DownloadClientBulkResource struct for DownloadClientBulkResource
type DownloadClientBulkResource struct {
	Ids []int32 `json:"ids,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	ApplyTags *ApplyTags `json:"applyTags,omitempty"`
	Enable NullableBool `json:"enable,omitempty"`
	Priority NullableInt32 `json:"priority,omitempty"`
	RemoveCompletedDownloads NullableBool `json:"removeCompletedDownloads,omitempty"`
	RemoveFailedDownloads NullableBool `json:"removeFailedDownloads,omitempty"`
}

// NewDownloadClientBulkResource instantiates a new DownloadClientBulkResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadClientBulkResource() *DownloadClientBulkResource {
	this := DownloadClientBulkResource{}
	return &this
}

// NewDownloadClientBulkResourceWithDefaults instantiates a new DownloadClientBulkResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadClientBulkResourceWithDefaults() *DownloadClientBulkResource {
	this := DownloadClientBulkResource{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientBulkResource) GetIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientBulkResource) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *DownloadClientBulkResource) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *DownloadClientBulkResource) SetIds(v []int32) {
	o.Ids = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientBulkResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientBulkResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DownloadClientBulkResource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *DownloadClientBulkResource) SetTags(v []int32) {
	o.Tags = v
}

// GetApplyTags returns the ApplyTags field value if set, zero value otherwise.
func (o *DownloadClientBulkResource) GetApplyTags() ApplyTags {
	if o == nil || IsNil(o.ApplyTags) {
		var ret ApplyTags
		return ret
	}
	return *o.ApplyTags
}

// GetApplyTagsOk returns a tuple with the ApplyTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientBulkResource) GetApplyTagsOk() (*ApplyTags, bool) {
	if o == nil || IsNil(o.ApplyTags) {
		return nil, false
	}
	return o.ApplyTags, true
}

// HasApplyTags returns a boolean if a field has been set.
func (o *DownloadClientBulkResource) HasApplyTags() bool {
	if o != nil && !IsNil(o.ApplyTags) {
		return true
	}

	return false
}

// SetApplyTags gets a reference to the given ApplyTags and assigns it to the ApplyTags field.
func (o *DownloadClientBulkResource) SetApplyTags(v ApplyTags) {
	o.ApplyTags = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientBulkResource) GetEnable() bool {
	if o == nil || IsNil(o.Enable.Get()) {
		var ret bool
		return ret
	}
	return *o.Enable.Get()
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientBulkResource) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enable.Get(), o.Enable.IsSet()
}

// HasEnable returns a boolean if a field has been set.
func (o *DownloadClientBulkResource) HasEnable() bool {
	if o != nil && o.Enable.IsSet() {
		return true
	}

	return false
}

// SetEnable gets a reference to the given NullableBool and assigns it to the Enable field.
func (o *DownloadClientBulkResource) SetEnable(v bool) {
	o.Enable.Set(&v)
}
// SetEnableNil sets the value for Enable to be an explicit nil
func (o *DownloadClientBulkResource) SetEnableNil() {
	o.Enable.Set(nil)
}

// UnsetEnable ensures that no value is present for Enable, not even an explicit nil
func (o *DownloadClientBulkResource) UnsetEnable() {
	o.Enable.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientBulkResource) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientBulkResource) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *DownloadClientBulkResource) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *DownloadClientBulkResource) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *DownloadClientBulkResource) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *DownloadClientBulkResource) UnsetPriority() {
	o.Priority.Unset()
}

// GetRemoveCompletedDownloads returns the RemoveCompletedDownloads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientBulkResource) GetRemoveCompletedDownloads() bool {
	if o == nil || IsNil(o.RemoveCompletedDownloads.Get()) {
		var ret bool
		return ret
	}
	return *o.RemoveCompletedDownloads.Get()
}

// GetRemoveCompletedDownloadsOk returns a tuple with the RemoveCompletedDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientBulkResource) GetRemoveCompletedDownloadsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveCompletedDownloads.Get(), o.RemoveCompletedDownloads.IsSet()
}

// HasRemoveCompletedDownloads returns a boolean if a field has been set.
func (o *DownloadClientBulkResource) HasRemoveCompletedDownloads() bool {
	if o != nil && o.RemoveCompletedDownloads.IsSet() {
		return true
	}

	return false
}

// SetRemoveCompletedDownloads gets a reference to the given NullableBool and assigns it to the RemoveCompletedDownloads field.
func (o *DownloadClientBulkResource) SetRemoveCompletedDownloads(v bool) {
	o.RemoveCompletedDownloads.Set(&v)
}
// SetRemoveCompletedDownloadsNil sets the value for RemoveCompletedDownloads to be an explicit nil
func (o *DownloadClientBulkResource) SetRemoveCompletedDownloadsNil() {
	o.RemoveCompletedDownloads.Set(nil)
}

// UnsetRemoveCompletedDownloads ensures that no value is present for RemoveCompletedDownloads, not even an explicit nil
func (o *DownloadClientBulkResource) UnsetRemoveCompletedDownloads() {
	o.RemoveCompletedDownloads.Unset()
}

// GetRemoveFailedDownloads returns the RemoveFailedDownloads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientBulkResource) GetRemoveFailedDownloads() bool {
	if o == nil || IsNil(o.RemoveFailedDownloads.Get()) {
		var ret bool
		return ret
	}
	return *o.RemoveFailedDownloads.Get()
}

// GetRemoveFailedDownloadsOk returns a tuple with the RemoveFailedDownloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientBulkResource) GetRemoveFailedDownloadsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveFailedDownloads.Get(), o.RemoveFailedDownloads.IsSet()
}

// HasRemoveFailedDownloads returns a boolean if a field has been set.
func (o *DownloadClientBulkResource) HasRemoveFailedDownloads() bool {
	if o != nil && o.RemoveFailedDownloads.IsSet() {
		return true
	}

	return false
}

// SetRemoveFailedDownloads gets a reference to the given NullableBool and assigns it to the RemoveFailedDownloads field.
func (o *DownloadClientBulkResource) SetRemoveFailedDownloads(v bool) {
	o.RemoveFailedDownloads.Set(&v)
}
// SetRemoveFailedDownloadsNil sets the value for RemoveFailedDownloads to be an explicit nil
func (o *DownloadClientBulkResource) SetRemoveFailedDownloadsNil() {
	o.RemoveFailedDownloads.Set(nil)
}

// UnsetRemoveFailedDownloads ensures that no value is present for RemoveFailedDownloads, not even an explicit nil
func (o *DownloadClientBulkResource) UnsetRemoveFailedDownloads() {
	o.RemoveFailedDownloads.Unset()
}

func (o DownloadClientBulkResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadClientBulkResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ApplyTags) {
		toSerialize["applyTags"] = o.ApplyTags
	}
	if o.Enable.IsSet() {
		toSerialize["enable"] = o.Enable.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.RemoveCompletedDownloads.IsSet() {
		toSerialize["removeCompletedDownloads"] = o.RemoveCompletedDownloads.Get()
	}
	if o.RemoveFailedDownloads.IsSet() {
		toSerialize["removeFailedDownloads"] = o.RemoveFailedDownloads.Get()
	}
	return toSerialize, nil
}

type NullableDownloadClientBulkResource struct {
	value *DownloadClientBulkResource
	isSet bool
}

func (v NullableDownloadClientBulkResource) Get() *DownloadClientBulkResource {
	return v.value
}

func (v *NullableDownloadClientBulkResource) Set(val *DownloadClientBulkResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadClientBulkResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadClientBulkResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadClientBulkResource(val *DownloadClientBulkResource) *NullableDownloadClientBulkResource {
	return &NullableDownloadClientBulkResource{value: val, isSet: true}
}

func (v NullableDownloadClientBulkResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadClientBulkResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



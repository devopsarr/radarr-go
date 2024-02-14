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

// ImportListConfigResource struct for ImportListConfigResource
type ImportListConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	ListSyncLevel NullableString `json:"listSyncLevel,omitempty"`
	ImportExclusions NullableString `json:"importExclusions,omitempty"`
}

// NewImportListConfigResource instantiates a new ImportListConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportListConfigResource() *ImportListConfigResource {
	this := ImportListConfigResource{}
	return &this
}

// NewImportListConfigResourceWithDefaults instantiates a new ImportListConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportListConfigResourceWithDefaults() *ImportListConfigResource {
	this := ImportListConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImportListConfigResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImportListConfigResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ImportListConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetListSyncLevel returns the ListSyncLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListConfigResource) GetListSyncLevel() string {
	if o == nil || isNil(o.ListSyncLevel.Get()) {
		var ret string
		return ret
	}
	return *o.ListSyncLevel.Get()
}

// GetListSyncLevelOk returns a tuple with the ListSyncLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListConfigResource) GetListSyncLevelOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ListSyncLevel.Get(), o.ListSyncLevel.IsSet()
}

// HasListSyncLevel returns a boolean if a field has been set.
func (o *ImportListConfigResource) HasListSyncLevel() bool {
	if o != nil && o.ListSyncLevel.IsSet() {
		return true
	}

	return false
}

// SetListSyncLevel gets a reference to the given NullableString and assigns it to the ListSyncLevel field.
func (o *ImportListConfigResource) SetListSyncLevel(v string) {
	o.ListSyncLevel.Set(&v)
}
// SetListSyncLevelNil sets the value for ListSyncLevel to be an explicit nil
func (o *ImportListConfigResource) SetListSyncLevelNil() {
	o.ListSyncLevel.Set(nil)
}

// UnsetListSyncLevel ensures that no value is present for ListSyncLevel, not even an explicit nil
func (o *ImportListConfigResource) UnsetListSyncLevel() {
	o.ListSyncLevel.Unset()
}

// GetImportExclusions returns the ImportExclusions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListConfigResource) GetImportExclusions() string {
	if o == nil || isNil(o.ImportExclusions.Get()) {
		var ret string
		return ret
	}
	return *o.ImportExclusions.Get()
}

// GetImportExclusionsOk returns a tuple with the ImportExclusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListConfigResource) GetImportExclusionsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ImportExclusions.Get(), o.ImportExclusions.IsSet()
}

// HasImportExclusions returns a boolean if a field has been set.
func (o *ImportListConfigResource) HasImportExclusions() bool {
	if o != nil && o.ImportExclusions.IsSet() {
		return true
	}

	return false
}

// SetImportExclusions gets a reference to the given NullableString and assigns it to the ImportExclusions field.
func (o *ImportListConfigResource) SetImportExclusions(v string) {
	o.ImportExclusions.Set(&v)
}
// SetImportExclusionsNil sets the value for ImportExclusions to be an explicit nil
func (o *ImportListConfigResource) SetImportExclusionsNil() {
	o.ImportExclusions.Set(nil)
}

// UnsetImportExclusions ensures that no value is present for ImportExclusions, not even an explicit nil
func (o *ImportListConfigResource) UnsetImportExclusions() {
	o.ImportExclusions.Unset()
}

func (o ImportListConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ListSyncLevel.IsSet() {
		toSerialize["listSyncLevel"] = o.ListSyncLevel.Get()
	}
	if o.ImportExclusions.IsSet() {
		toSerialize["importExclusions"] = o.ImportExclusions.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableImportListConfigResource struct {
	value *ImportListConfigResource
	isSet bool
}

func (v NullableImportListConfigResource) Get() *ImportListConfigResource {
	return v.value
}

func (v *NullableImportListConfigResource) Set(val *ImportListConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListConfigResource(val *ImportListConfigResource) *NullableImportListConfigResource {
	return &NullableImportListConfigResource{value: val, isSet: true}
}

func (v NullableImportListConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Radarr

Radarr API docs

API version: v5.17.2.9580
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the DownloadClientConfigResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadClientConfigResource{}

// DownloadClientConfigResource struct for DownloadClientConfigResource
type DownloadClientConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	DownloadClientWorkingFolders NullableString `json:"downloadClientWorkingFolders,omitempty"`
	EnableCompletedDownloadHandling *bool `json:"enableCompletedDownloadHandling,omitempty"`
	CheckForFinishedDownloadInterval *int32 `json:"checkForFinishedDownloadInterval,omitempty"`
	AutoRedownloadFailed *bool `json:"autoRedownloadFailed,omitempty"`
	AutoRedownloadFailedFromInteractiveSearch *bool `json:"autoRedownloadFailedFromInteractiveSearch,omitempty"`
}

// NewDownloadClientConfigResource instantiates a new DownloadClientConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadClientConfigResource() *DownloadClientConfigResource {
	this := DownloadClientConfigResource{}
	return &this
}

// NewDownloadClientConfigResourceWithDefaults instantiates a new DownloadClientConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadClientConfigResourceWithDefaults() *DownloadClientConfigResource {
	this := DownloadClientConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DownloadClientConfigResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DownloadClientConfigResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DownloadClientConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetDownloadClientWorkingFolders returns the DownloadClientWorkingFolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadClientConfigResource) GetDownloadClientWorkingFolders() string {
	if o == nil || IsNil(o.DownloadClientWorkingFolders.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadClientWorkingFolders.Get()
}

// GetDownloadClientWorkingFoldersOk returns a tuple with the DownloadClientWorkingFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadClientConfigResource) GetDownloadClientWorkingFoldersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadClientWorkingFolders.Get(), o.DownloadClientWorkingFolders.IsSet()
}

// HasDownloadClientWorkingFolders returns a boolean if a field has been set.
func (o *DownloadClientConfigResource) HasDownloadClientWorkingFolders() bool {
	if o != nil && o.DownloadClientWorkingFolders.IsSet() {
		return true
	}

	return false
}

// SetDownloadClientWorkingFolders gets a reference to the given NullableString and assigns it to the DownloadClientWorkingFolders field.
func (o *DownloadClientConfigResource) SetDownloadClientWorkingFolders(v string) {
	o.DownloadClientWorkingFolders.Set(&v)
}
// SetDownloadClientWorkingFoldersNil sets the value for DownloadClientWorkingFolders to be an explicit nil
func (o *DownloadClientConfigResource) SetDownloadClientWorkingFoldersNil() {
	o.DownloadClientWorkingFolders.Set(nil)
}

// UnsetDownloadClientWorkingFolders ensures that no value is present for DownloadClientWorkingFolders, not even an explicit nil
func (o *DownloadClientConfigResource) UnsetDownloadClientWorkingFolders() {
	o.DownloadClientWorkingFolders.Unset()
}

// GetEnableCompletedDownloadHandling returns the EnableCompletedDownloadHandling field value if set, zero value otherwise.
func (o *DownloadClientConfigResource) GetEnableCompletedDownloadHandling() bool {
	if o == nil || IsNil(o.EnableCompletedDownloadHandling) {
		var ret bool
		return ret
	}
	return *o.EnableCompletedDownloadHandling
}

// GetEnableCompletedDownloadHandlingOk returns a tuple with the EnableCompletedDownloadHandling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientConfigResource) GetEnableCompletedDownloadHandlingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCompletedDownloadHandling) {
		return nil, false
	}
	return o.EnableCompletedDownloadHandling, true
}

// HasEnableCompletedDownloadHandling returns a boolean if a field has been set.
func (o *DownloadClientConfigResource) HasEnableCompletedDownloadHandling() bool {
	if o != nil && !IsNil(o.EnableCompletedDownloadHandling) {
		return true
	}

	return false
}

// SetEnableCompletedDownloadHandling gets a reference to the given bool and assigns it to the EnableCompletedDownloadHandling field.
func (o *DownloadClientConfigResource) SetEnableCompletedDownloadHandling(v bool) {
	o.EnableCompletedDownloadHandling = &v
}

// GetCheckForFinishedDownloadInterval returns the CheckForFinishedDownloadInterval field value if set, zero value otherwise.
func (o *DownloadClientConfigResource) GetCheckForFinishedDownloadInterval() int32 {
	if o == nil || IsNil(o.CheckForFinishedDownloadInterval) {
		var ret int32
		return ret
	}
	return *o.CheckForFinishedDownloadInterval
}

// GetCheckForFinishedDownloadIntervalOk returns a tuple with the CheckForFinishedDownloadInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientConfigResource) GetCheckForFinishedDownloadIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckForFinishedDownloadInterval) {
		return nil, false
	}
	return o.CheckForFinishedDownloadInterval, true
}

// HasCheckForFinishedDownloadInterval returns a boolean if a field has been set.
func (o *DownloadClientConfigResource) HasCheckForFinishedDownloadInterval() bool {
	if o != nil && !IsNil(o.CheckForFinishedDownloadInterval) {
		return true
	}

	return false
}

// SetCheckForFinishedDownloadInterval gets a reference to the given int32 and assigns it to the CheckForFinishedDownloadInterval field.
func (o *DownloadClientConfigResource) SetCheckForFinishedDownloadInterval(v int32) {
	o.CheckForFinishedDownloadInterval = &v
}

// GetAutoRedownloadFailed returns the AutoRedownloadFailed field value if set, zero value otherwise.
func (o *DownloadClientConfigResource) GetAutoRedownloadFailed() bool {
	if o == nil || IsNil(o.AutoRedownloadFailed) {
		var ret bool
		return ret
	}
	return *o.AutoRedownloadFailed
}

// GetAutoRedownloadFailedOk returns a tuple with the AutoRedownloadFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientConfigResource) GetAutoRedownloadFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRedownloadFailed) {
		return nil, false
	}
	return o.AutoRedownloadFailed, true
}

// HasAutoRedownloadFailed returns a boolean if a field has been set.
func (o *DownloadClientConfigResource) HasAutoRedownloadFailed() bool {
	if o != nil && !IsNil(o.AutoRedownloadFailed) {
		return true
	}

	return false
}

// SetAutoRedownloadFailed gets a reference to the given bool and assigns it to the AutoRedownloadFailed field.
func (o *DownloadClientConfigResource) SetAutoRedownloadFailed(v bool) {
	o.AutoRedownloadFailed = &v
}

// GetAutoRedownloadFailedFromInteractiveSearch returns the AutoRedownloadFailedFromInteractiveSearch field value if set, zero value otherwise.
func (o *DownloadClientConfigResource) GetAutoRedownloadFailedFromInteractiveSearch() bool {
	if o == nil || IsNil(o.AutoRedownloadFailedFromInteractiveSearch) {
		var ret bool
		return ret
	}
	return *o.AutoRedownloadFailedFromInteractiveSearch
}

// GetAutoRedownloadFailedFromInteractiveSearchOk returns a tuple with the AutoRedownloadFailedFromInteractiveSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadClientConfigResource) GetAutoRedownloadFailedFromInteractiveSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRedownloadFailedFromInteractiveSearch) {
		return nil, false
	}
	return o.AutoRedownloadFailedFromInteractiveSearch, true
}

// HasAutoRedownloadFailedFromInteractiveSearch returns a boolean if a field has been set.
func (o *DownloadClientConfigResource) HasAutoRedownloadFailedFromInteractiveSearch() bool {
	if o != nil && !IsNil(o.AutoRedownloadFailedFromInteractiveSearch) {
		return true
	}

	return false
}

// SetAutoRedownloadFailedFromInteractiveSearch gets a reference to the given bool and assigns it to the AutoRedownloadFailedFromInteractiveSearch field.
func (o *DownloadClientConfigResource) SetAutoRedownloadFailedFromInteractiveSearch(v bool) {
	o.AutoRedownloadFailedFromInteractiveSearch = &v
}

func (o DownloadClientConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadClientConfigResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.DownloadClientWorkingFolders.IsSet() {
		toSerialize["downloadClientWorkingFolders"] = o.DownloadClientWorkingFolders.Get()
	}
	if !IsNil(o.EnableCompletedDownloadHandling) {
		toSerialize["enableCompletedDownloadHandling"] = o.EnableCompletedDownloadHandling
	}
	if !IsNil(o.CheckForFinishedDownloadInterval) {
		toSerialize["checkForFinishedDownloadInterval"] = o.CheckForFinishedDownloadInterval
	}
	if !IsNil(o.AutoRedownloadFailed) {
		toSerialize["autoRedownloadFailed"] = o.AutoRedownloadFailed
	}
	if !IsNil(o.AutoRedownloadFailedFromInteractiveSearch) {
		toSerialize["autoRedownloadFailedFromInteractiveSearch"] = o.AutoRedownloadFailedFromInteractiveSearch
	}
	return toSerialize, nil
}

type NullableDownloadClientConfigResource struct {
	value *DownloadClientConfigResource
	isSet bool
}

func (v NullableDownloadClientConfigResource) Get() *DownloadClientConfigResource {
	return v.value
}

func (v *NullableDownloadClientConfigResource) Set(val *DownloadClientConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadClientConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadClientConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadClientConfigResource(val *DownloadClientConfigResource) *NullableDownloadClientConfigResource {
	return &NullableDownloadClientConfigResource{value: val, isSet: true}
}

func (v NullableDownloadClientConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadClientConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



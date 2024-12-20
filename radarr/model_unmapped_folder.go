/*
Radarr

Radarr API docs

API version: v5.16.3.9541
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the UnmappedFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnmappedFolder{}

// UnmappedFolder struct for UnmappedFolder
type UnmappedFolder struct {
	Name NullableString `json:"name,omitempty"`
	Path NullableString `json:"path,omitempty"`
	RelativePath NullableString `json:"relativePath,omitempty"`
}

// NewUnmappedFolder instantiates a new UnmappedFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnmappedFolder() *UnmappedFolder {
	this := UnmappedFolder{}
	return &this
}

// NewUnmappedFolderWithDefaults instantiates a new UnmappedFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnmappedFolderWithDefaults() *UnmappedFolder {
	this := UnmappedFolder{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnmappedFolder) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnmappedFolder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UnmappedFolder) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UnmappedFolder) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UnmappedFolder) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UnmappedFolder) UnsetName() {
	o.Name.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnmappedFolder) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnmappedFolder) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *UnmappedFolder) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *UnmappedFolder) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *UnmappedFolder) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *UnmappedFolder) UnsetPath() {
	o.Path.Unset()
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnmappedFolder) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath.Get()) {
		var ret string
		return ret
	}
	return *o.RelativePath.Get()
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnmappedFolder) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelativePath.Get(), o.RelativePath.IsSet()
}

// HasRelativePath returns a boolean if a field has been set.
func (o *UnmappedFolder) HasRelativePath() bool {
	if o != nil && o.RelativePath.IsSet() {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given NullableString and assigns it to the RelativePath field.
func (o *UnmappedFolder) SetRelativePath(v string) {
	o.RelativePath.Set(&v)
}
// SetRelativePathNil sets the value for RelativePath to be an explicit nil
func (o *UnmappedFolder) SetRelativePathNil() {
	o.RelativePath.Set(nil)
}

// UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
func (o *UnmappedFolder) UnsetRelativePath() {
	o.RelativePath.Unset()
}

func (o UnmappedFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnmappedFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.RelativePath.IsSet() {
		toSerialize["relativePath"] = o.RelativePath.Get()
	}
	return toSerialize, nil
}

type NullableUnmappedFolder struct {
	value *UnmappedFolder
	isSet bool
}

func (v NullableUnmappedFolder) Get() *UnmappedFolder {
	return v.value
}

func (v *NullableUnmappedFolder) Set(val *UnmappedFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableUnmappedFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableUnmappedFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnmappedFolder(val *UnmappedFolder) *NullableUnmappedFolder {
	return &NullableUnmappedFolder{value: val, isSet: true}
}

func (v NullableUnmappedFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnmappedFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



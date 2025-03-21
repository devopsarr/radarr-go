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

// checks if the RemotePathMappingResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemotePathMappingResource{}

// RemotePathMappingResource struct for RemotePathMappingResource
type RemotePathMappingResource struct {
	Id *int32 `json:"id,omitempty"`
	Host NullableString `json:"host,omitempty"`
	RemotePath NullableString `json:"remotePath,omitempty"`
	LocalPath NullableString `json:"localPath,omitempty"`
}

// NewRemotePathMappingResource instantiates a new RemotePathMappingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemotePathMappingResource() *RemotePathMappingResource {
	this := RemotePathMappingResource{}
	return &this
}

// NewRemotePathMappingResourceWithDefaults instantiates a new RemotePathMappingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemotePathMappingResourceWithDefaults() *RemotePathMappingResource {
	this := RemotePathMappingResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemotePathMappingResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemotePathMappingResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemotePathMappingResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RemotePathMappingResource) SetId(v int32) {
	o.Id = &v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemotePathMappingResource) GetHost() string {
	if o == nil || IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemotePathMappingResource) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *RemotePathMappingResource) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *RemotePathMappingResource) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *RemotePathMappingResource) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *RemotePathMappingResource) UnsetHost() {
	o.Host.Unset()
}

// GetRemotePath returns the RemotePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemotePathMappingResource) GetRemotePath() string {
	if o == nil || IsNil(o.RemotePath.Get()) {
		var ret string
		return ret
	}
	return *o.RemotePath.Get()
}

// GetRemotePathOk returns a tuple with the RemotePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemotePathMappingResource) GetRemotePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemotePath.Get(), o.RemotePath.IsSet()
}

// HasRemotePath returns a boolean if a field has been set.
func (o *RemotePathMappingResource) HasRemotePath() bool {
	if o != nil && o.RemotePath.IsSet() {
		return true
	}

	return false
}

// SetRemotePath gets a reference to the given NullableString and assigns it to the RemotePath field.
func (o *RemotePathMappingResource) SetRemotePath(v string) {
	o.RemotePath.Set(&v)
}
// SetRemotePathNil sets the value for RemotePath to be an explicit nil
func (o *RemotePathMappingResource) SetRemotePathNil() {
	o.RemotePath.Set(nil)
}

// UnsetRemotePath ensures that no value is present for RemotePath, not even an explicit nil
func (o *RemotePathMappingResource) UnsetRemotePath() {
	o.RemotePath.Unset()
}

// GetLocalPath returns the LocalPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemotePathMappingResource) GetLocalPath() string {
	if o == nil || IsNil(o.LocalPath.Get()) {
		var ret string
		return ret
	}
	return *o.LocalPath.Get()
}

// GetLocalPathOk returns a tuple with the LocalPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemotePathMappingResource) GetLocalPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalPath.Get(), o.LocalPath.IsSet()
}

// HasLocalPath returns a boolean if a field has been set.
func (o *RemotePathMappingResource) HasLocalPath() bool {
	if o != nil && o.LocalPath.IsSet() {
		return true
	}

	return false
}

// SetLocalPath gets a reference to the given NullableString and assigns it to the LocalPath field.
func (o *RemotePathMappingResource) SetLocalPath(v string) {
	o.LocalPath.Set(&v)
}
// SetLocalPathNil sets the value for LocalPath to be an explicit nil
func (o *RemotePathMappingResource) SetLocalPathNil() {
	o.LocalPath.Set(nil)
}

// UnsetLocalPath ensures that no value is present for LocalPath, not even an explicit nil
func (o *RemotePathMappingResource) UnsetLocalPath() {
	o.LocalPath.Unset()
}

func (o RemotePathMappingResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemotePathMappingResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.RemotePath.IsSet() {
		toSerialize["remotePath"] = o.RemotePath.Get()
	}
	if o.LocalPath.IsSet() {
		toSerialize["localPath"] = o.LocalPath.Get()
	}
	return toSerialize, nil
}

type NullableRemotePathMappingResource struct {
	value *RemotePathMappingResource
	isSet bool
}

func (v NullableRemotePathMappingResource) Get() *RemotePathMappingResource {
	return v.value
}

func (v *NullableRemotePathMappingResource) Set(val *RemotePathMappingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableRemotePathMappingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableRemotePathMappingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemotePathMappingResource(val *RemotePathMappingResource) *NullableRemotePathMappingResource {
	return &NullableRemotePathMappingResource{value: val, isSet: true}
}

func (v NullableRemotePathMappingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemotePathMappingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



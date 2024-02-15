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

// checks if the AutoTaggingResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoTaggingResource{}

// AutoTaggingResource struct for AutoTaggingResource
type AutoTaggingResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	RemoveTagsAutomatically *bool `json:"removeTagsAutomatically,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	Specifications []AutoTaggingSpecificationSchema `json:"specifications,omitempty"`
}

// NewAutoTaggingResource instantiates a new AutoTaggingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoTaggingResource() *AutoTaggingResource {
	this := AutoTaggingResource{}
	return &this
}

// NewAutoTaggingResourceWithDefaults instantiates a new AutoTaggingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoTaggingResourceWithDefaults() *AutoTaggingResource {
	this := AutoTaggingResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoTaggingResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTaggingResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoTaggingResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AutoTaggingResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTaggingResource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTaggingResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AutoTaggingResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AutoTaggingResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AutoTaggingResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AutoTaggingResource) UnsetName() {
	o.Name.Unset()
}

// GetRemoveTagsAutomatically returns the RemoveTagsAutomatically field value if set, zero value otherwise.
func (o *AutoTaggingResource) GetRemoveTagsAutomatically() bool {
	if o == nil || IsNil(o.RemoveTagsAutomatically) {
		var ret bool
		return ret
	}
	return *o.RemoveTagsAutomatically
}

// GetRemoveTagsAutomaticallyOk returns a tuple with the RemoveTagsAutomatically field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoTaggingResource) GetRemoveTagsAutomaticallyOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveTagsAutomatically) {
		return nil, false
	}
	return o.RemoveTagsAutomatically, true
}

// HasRemoveTagsAutomatically returns a boolean if a field has been set.
func (o *AutoTaggingResource) HasRemoveTagsAutomatically() bool {
	if o != nil && !IsNil(o.RemoveTagsAutomatically) {
		return true
	}

	return false
}

// SetRemoveTagsAutomatically gets a reference to the given bool and assigns it to the RemoveTagsAutomatically field.
func (o *AutoTaggingResource) SetRemoveTagsAutomatically(v bool) {
	o.RemoveTagsAutomatically = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTaggingResource) GetTags() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTaggingResource) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AutoTaggingResource) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *AutoTaggingResource) SetTags(v []int32) {
	o.Tags = v
}

// GetSpecifications returns the Specifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutoTaggingResource) GetSpecifications() []AutoTaggingSpecificationSchema {
	if o == nil {
		var ret []AutoTaggingSpecificationSchema
		return ret
	}
	return o.Specifications
}

// GetSpecificationsOk returns a tuple with the Specifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutoTaggingResource) GetSpecificationsOk() ([]AutoTaggingSpecificationSchema, bool) {
	if o == nil || IsNil(o.Specifications) {
		return nil, false
	}
	return o.Specifications, true
}

// HasSpecifications returns a boolean if a field has been set.
func (o *AutoTaggingResource) HasSpecifications() bool {
	if o != nil && IsNil(o.Specifications) {
		return true
	}

	return false
}

// SetSpecifications gets a reference to the given []AutoTaggingSpecificationSchema and assigns it to the Specifications field.
func (o *AutoTaggingResource) SetSpecifications(v []AutoTaggingSpecificationSchema) {
	o.Specifications = v
}

func (o AutoTaggingResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoTaggingResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.RemoveTagsAutomatically) {
		toSerialize["removeTagsAutomatically"] = o.RemoveTagsAutomatically
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Specifications != nil {
		toSerialize["specifications"] = o.Specifications
	}
	return toSerialize, nil
}

type NullableAutoTaggingResource struct {
	value *AutoTaggingResource
	isSet bool
}

func (v NullableAutoTaggingResource) Get() *AutoTaggingResource {
	return v.value
}

func (v *NullableAutoTaggingResource) Set(val *AutoTaggingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoTaggingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoTaggingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoTaggingResource(val *AutoTaggingResource) *NullableAutoTaggingResource {
	return &NullableAutoTaggingResource{value: val, isSet: true}
}

func (v NullableAutoTaggingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoTaggingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ImportListExclusionResourcePagingResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportListExclusionResourcePagingResource{}

// ImportListExclusionResourcePagingResource struct for ImportListExclusionResourcePagingResource
type ImportListExclusionResourcePagingResource struct {
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	SortKey NullableString `json:"sortKey,omitempty"`
	SortDirection *SortDirection `json:"sortDirection,omitempty"`
	TotalRecords *int32 `json:"totalRecords,omitempty"`
	Records []ImportListExclusionResource `json:"records,omitempty"`
}

// NewImportListExclusionResourcePagingResource instantiates a new ImportListExclusionResourcePagingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportListExclusionResourcePagingResource() *ImportListExclusionResourcePagingResource {
	this := ImportListExclusionResourcePagingResource{}
	return &this
}

// NewImportListExclusionResourcePagingResourceWithDefaults instantiates a new ImportListExclusionResourcePagingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportListExclusionResourcePagingResourceWithDefaults() *ImportListExclusionResourcePagingResource {
	this := ImportListExclusionResourcePagingResource{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ImportListExclusionResourcePagingResource) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResourcePagingResource) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ImportListExclusionResourcePagingResource) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *ImportListExclusionResourcePagingResource) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ImportListExclusionResourcePagingResource) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResourcePagingResource) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ImportListExclusionResourcePagingResource) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ImportListExclusionResourcePagingResource) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResourcePagingResource) GetSortKey() string {
	if o == nil || IsNil(o.SortKey.Get()) {
		var ret string
		return ret
	}
	return *o.SortKey.Get()
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResourcePagingResource) GetSortKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortKey.Get(), o.SortKey.IsSet()
}

// HasSortKey returns a boolean if a field has been set.
func (o *ImportListExclusionResourcePagingResource) HasSortKey() bool {
	if o != nil && o.SortKey.IsSet() {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given NullableString and assigns it to the SortKey field.
func (o *ImportListExclusionResourcePagingResource) SetSortKey(v string) {
	o.SortKey.Set(&v)
}
// SetSortKeyNil sets the value for SortKey to be an explicit nil
func (o *ImportListExclusionResourcePagingResource) SetSortKeyNil() {
	o.SortKey.Set(nil)
}

// UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
func (o *ImportListExclusionResourcePagingResource) UnsetSortKey() {
	o.SortKey.Unset()
}

// GetSortDirection returns the SortDirection field value if set, zero value otherwise.
func (o *ImportListExclusionResourcePagingResource) GetSortDirection() SortDirection {
	if o == nil || IsNil(o.SortDirection) {
		var ret SortDirection
		return ret
	}
	return *o.SortDirection
}

// GetSortDirectionOk returns a tuple with the SortDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool) {
	if o == nil || IsNil(o.SortDirection) {
		return nil, false
	}
	return o.SortDirection, true
}

// HasSortDirection returns a boolean if a field has been set.
func (o *ImportListExclusionResourcePagingResource) HasSortDirection() bool {
	if o != nil && !IsNil(o.SortDirection) {
		return true
	}

	return false
}

// SetSortDirection gets a reference to the given SortDirection and assigns it to the SortDirection field.
func (o *ImportListExclusionResourcePagingResource) SetSortDirection(v SortDirection) {
	o.SortDirection = &v
}

// GetTotalRecords returns the TotalRecords field value if set, zero value otherwise.
func (o *ImportListExclusionResourcePagingResource) GetTotalRecords() int32 {
	if o == nil || IsNil(o.TotalRecords) {
		var ret int32
		return ret
	}
	return *o.TotalRecords
}

// GetTotalRecordsOk returns a tuple with the TotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportListExclusionResourcePagingResource) GetTotalRecordsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecords) {
		return nil, false
	}
	return o.TotalRecords, true
}

// HasTotalRecords returns a boolean if a field has been set.
func (o *ImportListExclusionResourcePagingResource) HasTotalRecords() bool {
	if o != nil && !IsNil(o.TotalRecords) {
		return true
	}

	return false
}

// SetTotalRecords gets a reference to the given int32 and assigns it to the TotalRecords field.
func (o *ImportListExclusionResourcePagingResource) SetTotalRecords(v int32) {
	o.TotalRecords = &v
}

// GetRecords returns the Records field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportListExclusionResourcePagingResource) GetRecords() []ImportListExclusionResource {
	if o == nil {
		var ret []ImportListExclusionResource
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportListExclusionResourcePagingResource) GetRecordsOk() ([]ImportListExclusionResource, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *ImportListExclusionResourcePagingResource) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []ImportListExclusionResource and assigns it to the Records field.
func (o *ImportListExclusionResourcePagingResource) SetRecords(v []ImportListExclusionResource) {
	o.Records = v
}

func (o ImportListExclusionResourcePagingResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportListExclusionResourcePagingResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.SortKey.IsSet() {
		toSerialize["sortKey"] = o.SortKey.Get()
	}
	if !IsNil(o.SortDirection) {
		toSerialize["sortDirection"] = o.SortDirection
	}
	if !IsNil(o.TotalRecords) {
		toSerialize["totalRecords"] = o.TotalRecords
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return toSerialize, nil
}

type NullableImportListExclusionResourcePagingResource struct {
	value *ImportListExclusionResourcePagingResource
	isSet bool
}

func (v NullableImportListExclusionResourcePagingResource) Get() *ImportListExclusionResourcePagingResource {
	return v.value
}

func (v *NullableImportListExclusionResourcePagingResource) Set(val *ImportListExclusionResourcePagingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportListExclusionResourcePagingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportListExclusionResourcePagingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportListExclusionResourcePagingResource(val *ImportListExclusionResourcePagingResource) *NullableImportListExclusionResourcePagingResource {
	return &NullableImportListExclusionResourcePagingResource{value: val, isSet: true}
}

func (v NullableImportListExclusionResourcePagingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportListExclusionResourcePagingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



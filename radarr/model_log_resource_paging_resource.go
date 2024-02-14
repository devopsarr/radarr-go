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

// LogResourcePagingResource struct for LogResourcePagingResource
type LogResourcePagingResource struct {
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	SortKey NullableString `json:"sortKey,omitempty"`
	SortDirection *SortDirection `json:"sortDirection,omitempty"`
	TotalRecords *int32 `json:"totalRecords,omitempty"`
	Records []*LogResource `json:"records,omitempty"`
}

// NewLogResourcePagingResource instantiates a new LogResourcePagingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogResourcePagingResource() *LogResourcePagingResource {
	this := LogResourcePagingResource{}
	return &this
}

// NewLogResourcePagingResourceWithDefaults instantiates a new LogResourcePagingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogResourcePagingResourceWithDefaults() *LogResourcePagingResource {
	this := LogResourcePagingResource{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *LogResourcePagingResource) GetPage() int32 {
	if o == nil || isNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResourcePagingResource) GetPageOk() (*int32, bool) {
	if o == nil || isNil(o.Page) {
    return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *LogResourcePagingResource) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *LogResourcePagingResource) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *LogResourcePagingResource) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResourcePagingResource) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *LogResourcePagingResource) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *LogResourcePagingResource) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResourcePagingResource) GetSortKey() string {
	if o == nil || isNil(o.SortKey.Get()) {
		var ret string
		return ret
	}
	return *o.SortKey.Get()
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResourcePagingResource) GetSortKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SortKey.Get(), o.SortKey.IsSet()
}

// HasSortKey returns a boolean if a field has been set.
func (o *LogResourcePagingResource) HasSortKey() bool {
	if o != nil && o.SortKey.IsSet() {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given NullableString and assigns it to the SortKey field.
func (o *LogResourcePagingResource) SetSortKey(v string) {
	o.SortKey.Set(&v)
}
// SetSortKeyNil sets the value for SortKey to be an explicit nil
func (o *LogResourcePagingResource) SetSortKeyNil() {
	o.SortKey.Set(nil)
}

// UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
func (o *LogResourcePagingResource) UnsetSortKey() {
	o.SortKey.Unset()
}

// GetSortDirection returns the SortDirection field value if set, zero value otherwise.
func (o *LogResourcePagingResource) GetSortDirection() SortDirection {
	if o == nil || isNil(o.SortDirection) {
		var ret SortDirection
		return ret
	}
	return *o.SortDirection
}

// GetSortDirectionOk returns a tuple with the SortDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool) {
	if o == nil || isNil(o.SortDirection) {
    return nil, false
	}
	return o.SortDirection, true
}

// HasSortDirection returns a boolean if a field has been set.
func (o *LogResourcePagingResource) HasSortDirection() bool {
	if o != nil && !isNil(o.SortDirection) {
		return true
	}

	return false
}

// SetSortDirection gets a reference to the given SortDirection and assigns it to the SortDirection field.
func (o *LogResourcePagingResource) SetSortDirection(v SortDirection) {
	o.SortDirection = &v
}

// GetTotalRecords returns the TotalRecords field value if set, zero value otherwise.
func (o *LogResourcePagingResource) GetTotalRecords() int32 {
	if o == nil || isNil(o.TotalRecords) {
		var ret int32
		return ret
	}
	return *o.TotalRecords
}

// GetTotalRecordsOk returns a tuple with the TotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogResourcePagingResource) GetTotalRecordsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalRecords) {
    return nil, false
	}
	return o.TotalRecords, true
}

// HasTotalRecords returns a boolean if a field has been set.
func (o *LogResourcePagingResource) HasTotalRecords() bool {
	if o != nil && !isNil(o.TotalRecords) {
		return true
	}

	return false
}

// SetTotalRecords gets a reference to the given int32 and assigns it to the TotalRecords field.
func (o *LogResourcePagingResource) SetTotalRecords(v int32) {
	o.TotalRecords = &v
}

// GetRecords returns the Records field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogResourcePagingResource) GetRecords() []*LogResource {
	if o == nil {
		var ret []*LogResource
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogResourcePagingResource) GetRecordsOk() ([]*LogResource, bool) {
	if o == nil || isNil(o.Records) {
    return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *LogResourcePagingResource) HasRecords() bool {
	if o != nil && isNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []LogResource and assigns it to the Records field.
func (o *LogResourcePagingResource) SetRecords(v []*LogResource) {
	o.Records = v
}

func (o LogResourcePagingResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.SortKey.IsSet() {
		toSerialize["sortKey"] = o.SortKey.Get()
	}
	if !isNil(o.SortDirection) {
		toSerialize["sortDirection"] = o.SortDirection
	}
	if !isNil(o.TotalRecords) {
		toSerialize["totalRecords"] = o.TotalRecords
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableLogResourcePagingResource struct {
	value *LogResourcePagingResource
	isSet bool
}

func (v NullableLogResourcePagingResource) Get() *LogResourcePagingResource {
	return v.value
}

func (v *NullableLogResourcePagingResource) Set(val *LogResourcePagingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableLogResourcePagingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableLogResourcePagingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogResourcePagingResource(val *LogResourcePagingResource) *NullableLogResourcePagingResource {
	return &NullableLogResourcePagingResource{value: val, isSet: true}
}

func (v NullableLogResourcePagingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogResourcePagingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



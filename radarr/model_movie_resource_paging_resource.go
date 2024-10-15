/*
Radarr

Radarr API docs

API version: v5.12.2.9335
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the MovieResourcePagingResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MovieResourcePagingResource{}

// MovieResourcePagingResource struct for MovieResourcePagingResource
type MovieResourcePagingResource struct {
	Page *int32 `json:"page,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	SortKey NullableString `json:"sortKey,omitempty"`
	SortDirection *SortDirection `json:"sortDirection,omitempty"`
	TotalRecords *int32 `json:"totalRecords,omitempty"`
	Records []MovieResource `json:"records,omitempty"`
}

// NewMovieResourcePagingResource instantiates a new MovieResourcePagingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieResourcePagingResource() *MovieResourcePagingResource {
	this := MovieResourcePagingResource{}
	return &this
}

// NewMovieResourcePagingResourceWithDefaults instantiates a new MovieResourcePagingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieResourcePagingResourceWithDefaults() *MovieResourcePagingResource {
	this := MovieResourcePagingResource{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *MovieResourcePagingResource) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieResourcePagingResource) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *MovieResourcePagingResource) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *MovieResourcePagingResource) SetPage(v int32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *MovieResourcePagingResource) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieResourcePagingResource) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *MovieResourcePagingResource) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *MovieResourcePagingResource) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieResourcePagingResource) GetSortKey() string {
	if o == nil || IsNil(o.SortKey.Get()) {
		var ret string
		return ret
	}
	return *o.SortKey.Get()
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieResourcePagingResource) GetSortKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortKey.Get(), o.SortKey.IsSet()
}

// HasSortKey returns a boolean if a field has been set.
func (o *MovieResourcePagingResource) HasSortKey() bool {
	if o != nil && o.SortKey.IsSet() {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given NullableString and assigns it to the SortKey field.
func (o *MovieResourcePagingResource) SetSortKey(v string) {
	o.SortKey.Set(&v)
}
// SetSortKeyNil sets the value for SortKey to be an explicit nil
func (o *MovieResourcePagingResource) SetSortKeyNil() {
	o.SortKey.Set(nil)
}

// UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
func (o *MovieResourcePagingResource) UnsetSortKey() {
	o.SortKey.Unset()
}

// GetSortDirection returns the SortDirection field value if set, zero value otherwise.
func (o *MovieResourcePagingResource) GetSortDirection() SortDirection {
	if o == nil || IsNil(o.SortDirection) {
		var ret SortDirection
		return ret
	}
	return *o.SortDirection
}

// GetSortDirectionOk returns a tuple with the SortDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool) {
	if o == nil || IsNil(o.SortDirection) {
		return nil, false
	}
	return o.SortDirection, true
}

// HasSortDirection returns a boolean if a field has been set.
func (o *MovieResourcePagingResource) HasSortDirection() bool {
	if o != nil && !IsNil(o.SortDirection) {
		return true
	}

	return false
}

// SetSortDirection gets a reference to the given SortDirection and assigns it to the SortDirection field.
func (o *MovieResourcePagingResource) SetSortDirection(v SortDirection) {
	o.SortDirection = &v
}

// GetTotalRecords returns the TotalRecords field value if set, zero value otherwise.
func (o *MovieResourcePagingResource) GetTotalRecords() int32 {
	if o == nil || IsNil(o.TotalRecords) {
		var ret int32
		return ret
	}
	return *o.TotalRecords
}

// GetTotalRecordsOk returns a tuple with the TotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieResourcePagingResource) GetTotalRecordsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecords) {
		return nil, false
	}
	return o.TotalRecords, true
}

// HasTotalRecords returns a boolean if a field has been set.
func (o *MovieResourcePagingResource) HasTotalRecords() bool {
	if o != nil && !IsNil(o.TotalRecords) {
		return true
	}

	return false
}

// SetTotalRecords gets a reference to the given int32 and assigns it to the TotalRecords field.
func (o *MovieResourcePagingResource) SetTotalRecords(v int32) {
	o.TotalRecords = &v
}

// GetRecords returns the Records field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieResourcePagingResource) GetRecords() []MovieResource {
	if o == nil {
		var ret []MovieResource
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieResourcePagingResource) GetRecordsOk() ([]MovieResource, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *MovieResourcePagingResource) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []MovieResource and assigns it to the Records field.
func (o *MovieResourcePagingResource) SetRecords(v []MovieResource) {
	o.Records = v
}

func (o MovieResourcePagingResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovieResourcePagingResource) ToMap() (map[string]interface{}, error) {
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

type NullableMovieResourcePagingResource struct {
	value *MovieResourcePagingResource
	isSet bool
}

func (v NullableMovieResourcePagingResource) Get() *MovieResourcePagingResource {
	return v.value
}

func (v *NullableMovieResourcePagingResource) Set(val *MovieResourcePagingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieResourcePagingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieResourcePagingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieResourcePagingResource(val *MovieResourcePagingResource) *NullableMovieResourcePagingResource {
	return &NullableMovieResourcePagingResource{value: val, isSet: true}
}

func (v NullableMovieResourcePagingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieResourcePagingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



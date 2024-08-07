# MovieResourcePagingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**SortKey** | Pointer to **NullableString** |  | [optional] 
**SortDirection** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 
**Records** | Pointer to [**[]MovieResource**](MovieResource.md) |  | [optional] 

## Methods

### NewMovieResourcePagingResource

`func NewMovieResourcePagingResource() *MovieResourcePagingResource`

NewMovieResourcePagingResource instantiates a new MovieResourcePagingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieResourcePagingResourceWithDefaults

`func NewMovieResourcePagingResourceWithDefaults() *MovieResourcePagingResource`

NewMovieResourcePagingResourceWithDefaults instantiates a new MovieResourcePagingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *MovieResourcePagingResource) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *MovieResourcePagingResource) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *MovieResourcePagingResource) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *MovieResourcePagingResource) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *MovieResourcePagingResource) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *MovieResourcePagingResource) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *MovieResourcePagingResource) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *MovieResourcePagingResource) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSortKey

`func (o *MovieResourcePagingResource) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *MovieResourcePagingResource) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *MovieResourcePagingResource) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *MovieResourcePagingResource) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### SetSortKeyNil

`func (o *MovieResourcePagingResource) SetSortKeyNil(b bool)`

 SetSortKeyNil sets the value for SortKey to be an explicit nil

### UnsetSortKey
`func (o *MovieResourcePagingResource) UnsetSortKey()`

UnsetSortKey ensures that no value is present for SortKey, not even an explicit nil
### GetSortDirection

`func (o *MovieResourcePagingResource) GetSortDirection() SortDirection`

GetSortDirection returns the SortDirection field if non-nil, zero value otherwise.

### GetSortDirectionOk

`func (o *MovieResourcePagingResource) GetSortDirectionOk() (*SortDirection, bool)`

GetSortDirectionOk returns a tuple with the SortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDirection

`func (o *MovieResourcePagingResource) SetSortDirection(v SortDirection)`

SetSortDirection sets SortDirection field to given value.

### HasSortDirection

`func (o *MovieResourcePagingResource) HasSortDirection() bool`

HasSortDirection returns a boolean if a field has been set.

### GetTotalRecords

`func (o *MovieResourcePagingResource) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *MovieResourcePagingResource) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *MovieResourcePagingResource) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *MovieResourcePagingResource) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetRecords

`func (o *MovieResourcePagingResource) GetRecords() []MovieResource`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *MovieResourcePagingResource) GetRecordsOk() (*[]MovieResource, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *MovieResourcePagingResource) SetRecords(v []MovieResource)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *MovieResourcePagingResource) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### SetRecordsNil

`func (o *MovieResourcePagingResource) SetRecordsNil(b bool)`

 SetRecordsNil sets the value for Records to be an explicit nil

### UnsetRecords
`func (o *MovieResourcePagingResource) UnsetRecords()`

UnsetRecords ensures that no value is present for Records, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



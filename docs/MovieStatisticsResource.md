# MovieStatisticsResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MovieFileCount** | Pointer to **int32** |  | [optional] 
**SizeOnDisk** | Pointer to **int64** |  | [optional] 
**ReleaseGroups** | Pointer to **[]string** |  | [optional] 
**MovieFileQualities** | Pointer to [**[]Quality**](Quality.md) |  | [optional] 

## Methods

### NewMovieStatisticsResource

`func NewMovieStatisticsResource() *MovieStatisticsResource`

NewMovieStatisticsResource instantiates a new MovieStatisticsResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieStatisticsResourceWithDefaults

`func NewMovieStatisticsResourceWithDefaults() *MovieStatisticsResource`

NewMovieStatisticsResourceWithDefaults instantiates a new MovieStatisticsResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovieFileCount

`func (o *MovieStatisticsResource) GetMovieFileCount() int32`

GetMovieFileCount returns the MovieFileCount field if non-nil, zero value otherwise.

### GetMovieFileCountOk

`func (o *MovieStatisticsResource) GetMovieFileCountOk() (*int32, bool)`

GetMovieFileCountOk returns a tuple with the MovieFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieFileCount

`func (o *MovieStatisticsResource) SetMovieFileCount(v int32)`

SetMovieFileCount sets MovieFileCount field to given value.

### HasMovieFileCount

`func (o *MovieStatisticsResource) HasMovieFileCount() bool`

HasMovieFileCount returns a boolean if a field has been set.

### GetSizeOnDisk

`func (o *MovieStatisticsResource) GetSizeOnDisk() int64`

GetSizeOnDisk returns the SizeOnDisk field if non-nil, zero value otherwise.

### GetSizeOnDiskOk

`func (o *MovieStatisticsResource) GetSizeOnDiskOk() (*int64, bool)`

GetSizeOnDiskOk returns a tuple with the SizeOnDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeOnDisk

`func (o *MovieStatisticsResource) SetSizeOnDisk(v int64)`

SetSizeOnDisk sets SizeOnDisk field to given value.

### HasSizeOnDisk

`func (o *MovieStatisticsResource) HasSizeOnDisk() bool`

HasSizeOnDisk returns a boolean if a field has been set.

### GetReleaseGroups

`func (o *MovieStatisticsResource) GetReleaseGroups() []string`

GetReleaseGroups returns the ReleaseGroups field if non-nil, zero value otherwise.

### GetReleaseGroupsOk

`func (o *MovieStatisticsResource) GetReleaseGroupsOk() (*[]string, bool)`

GetReleaseGroupsOk returns a tuple with the ReleaseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroups

`func (o *MovieStatisticsResource) SetReleaseGroups(v []string)`

SetReleaseGroups sets ReleaseGroups field to given value.

### HasReleaseGroups

`func (o *MovieStatisticsResource) HasReleaseGroups() bool`

HasReleaseGroups returns a boolean if a field has been set.

### SetReleaseGroupsNil

`func (o *MovieStatisticsResource) SetReleaseGroupsNil(b bool)`

 SetReleaseGroupsNil sets the value for ReleaseGroups to be an explicit nil

### UnsetReleaseGroups
`func (o *MovieStatisticsResource) UnsetReleaseGroups()`

UnsetReleaseGroups ensures that no value is present for ReleaseGroups, not even an explicit nil
### GetMovieFileQualities

`func (o *MovieStatisticsResource) GetMovieFileQualities() []Quality`

GetMovieFileQualities returns the MovieFileQualities field if non-nil, zero value otherwise.

### GetMovieFileQualitiesOk

`func (o *MovieStatisticsResource) GetMovieFileQualitiesOk() (*[]Quality, bool)`

GetMovieFileQualitiesOk returns a tuple with the MovieFileQualities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieFileQualities

`func (o *MovieStatisticsResource) SetMovieFileQualities(v []Quality)`

SetMovieFileQualities sets MovieFileQualities field to given value.

### HasMovieFileQualities

`func (o *MovieStatisticsResource) HasMovieFileQualities() bool`

HasMovieFileQualities returns a boolean if a field has been set.

### SetMovieFileQualitiesNil

`func (o *MovieStatisticsResource) SetMovieFileQualitiesNil(b bool)`

 SetMovieFileQualitiesNil sets the value for MovieFileQualities to be an explicit nil

### UnsetMovieFileQualities
`func (o *MovieStatisticsResource) UnsetMovieFileQualities()`

UnsetMovieFileQualities ensures that no value is present for MovieFileQualities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



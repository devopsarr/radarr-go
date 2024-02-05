# CollectionUpdateResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionIds** | Pointer to **[]int32** |  | [optional] 
**Monitored** | Pointer to **NullableBool** |  | [optional] 
**MonitorMovies** | Pointer to **NullableBool** |  | [optional] 
**QualityProfileId** | Pointer to **NullableInt32** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**MinimumAvailability** | Pointer to [**MovieStatusType**](MovieStatusType.md) |  | [optional] 

## Methods

### NewCollectionUpdateResource

`func NewCollectionUpdateResource() *CollectionUpdateResource`

NewCollectionUpdateResource instantiates a new CollectionUpdateResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionUpdateResourceWithDefaults

`func NewCollectionUpdateResourceWithDefaults() *CollectionUpdateResource`

NewCollectionUpdateResourceWithDefaults instantiates a new CollectionUpdateResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionIds

`func (o *CollectionUpdateResource) GetCollectionIds() []int32`

GetCollectionIds returns the CollectionIds field if non-nil, zero value otherwise.

### GetCollectionIdsOk

`func (o *CollectionUpdateResource) GetCollectionIdsOk() (*[]int32, bool)`

GetCollectionIdsOk returns a tuple with the CollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionIds

`func (o *CollectionUpdateResource) SetCollectionIds(v []int32)`

SetCollectionIds sets CollectionIds field to given value.

### HasCollectionIds

`func (o *CollectionUpdateResource) HasCollectionIds() bool`

HasCollectionIds returns a boolean if a field has been set.

### SetCollectionIdsNil

`func (o *CollectionUpdateResource) SetCollectionIdsNil(b bool)`

 SetCollectionIdsNil sets the value for CollectionIds to be an explicit nil

### UnsetCollectionIds
`func (o *CollectionUpdateResource) UnsetCollectionIds()`

UnsetCollectionIds ensures that no value is present for CollectionIds, not even an explicit nil
### GetMonitored

`func (o *CollectionUpdateResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *CollectionUpdateResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *CollectionUpdateResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *CollectionUpdateResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### SetMonitoredNil

`func (o *CollectionUpdateResource) SetMonitoredNil(b bool)`

 SetMonitoredNil sets the value for Monitored to be an explicit nil

### UnsetMonitored
`func (o *CollectionUpdateResource) UnsetMonitored()`

UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
### GetMonitorMovies

`func (o *CollectionUpdateResource) GetMonitorMovies() bool`

GetMonitorMovies returns the MonitorMovies field if non-nil, zero value otherwise.

### GetMonitorMoviesOk

`func (o *CollectionUpdateResource) GetMonitorMoviesOk() (*bool, bool)`

GetMonitorMoviesOk returns a tuple with the MonitorMovies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorMovies

`func (o *CollectionUpdateResource) SetMonitorMovies(v bool)`

SetMonitorMovies sets MonitorMovies field to given value.

### HasMonitorMovies

`func (o *CollectionUpdateResource) HasMonitorMovies() bool`

HasMonitorMovies returns a boolean if a field has been set.

### SetMonitorMoviesNil

`func (o *CollectionUpdateResource) SetMonitorMoviesNil(b bool)`

 SetMonitorMoviesNil sets the value for MonitorMovies to be an explicit nil

### UnsetMonitorMovies
`func (o *CollectionUpdateResource) UnsetMonitorMovies()`

UnsetMonitorMovies ensures that no value is present for MonitorMovies, not even an explicit nil
### GetQualityProfileId

`func (o *CollectionUpdateResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *CollectionUpdateResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *CollectionUpdateResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *CollectionUpdateResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### SetQualityProfileIdNil

`func (o *CollectionUpdateResource) SetQualityProfileIdNil(b bool)`

 SetQualityProfileIdNil sets the value for QualityProfileId to be an explicit nil

### UnsetQualityProfileId
`func (o *CollectionUpdateResource) UnsetQualityProfileId()`

UnsetQualityProfileId ensures that no value is present for QualityProfileId, not even an explicit nil
### GetRootFolderPath

`func (o *CollectionUpdateResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *CollectionUpdateResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *CollectionUpdateResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *CollectionUpdateResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *CollectionUpdateResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *CollectionUpdateResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetMinimumAvailability

`func (o *CollectionUpdateResource) GetMinimumAvailability() MovieStatusType`

GetMinimumAvailability returns the MinimumAvailability field if non-nil, zero value otherwise.

### GetMinimumAvailabilityOk

`func (o *CollectionUpdateResource) GetMinimumAvailabilityOk() (*MovieStatusType, bool)`

GetMinimumAvailabilityOk returns a tuple with the MinimumAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAvailability

`func (o *CollectionUpdateResource) SetMinimumAvailability(v MovieStatusType)`

SetMinimumAvailability sets MinimumAvailability field to given value.

### HasMinimumAvailability

`func (o *CollectionUpdateResource) HasMinimumAvailability() bool`

HasMinimumAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



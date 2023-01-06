# MovieCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**SortTitle** | Pointer to **NullableString** |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**QualityProfileId** | Pointer to **int32** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**SearchOnAdd** | Pointer to **bool** |  | [optional] 
**MinimumAvailability** | Pointer to [**MovieStatusType**](MovieStatusType.md) |  | [optional] 
**LastInfoSync** | Pointer to **NullableTime** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**Movies** | Pointer to [**[]MovieMetadata**](MovieMetadata.md) |  | [optional] 

## Methods

### NewMovieCollection

`func NewMovieCollection() *MovieCollection`

NewMovieCollection instantiates a new MovieCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieCollectionWithDefaults

`func NewMovieCollectionWithDefaults() *MovieCollection`

NewMovieCollectionWithDefaults instantiates a new MovieCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MovieCollection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MovieCollection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MovieCollection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MovieCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *MovieCollection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MovieCollection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MovieCollection) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MovieCollection) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MovieCollection) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MovieCollection) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCleanTitle

`func (o *MovieCollection) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *MovieCollection) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *MovieCollection) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *MovieCollection) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *MovieCollection) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *MovieCollection) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetSortTitle

`func (o *MovieCollection) GetSortTitle() string`

GetSortTitle returns the SortTitle field if non-nil, zero value otherwise.

### GetSortTitleOk

`func (o *MovieCollection) GetSortTitleOk() (*string, bool)`

GetSortTitleOk returns a tuple with the SortTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortTitle

`func (o *MovieCollection) SetSortTitle(v string)`

SetSortTitle sets SortTitle field to given value.

### HasSortTitle

`func (o *MovieCollection) HasSortTitle() bool`

HasSortTitle returns a boolean if a field has been set.

### SetSortTitleNil

`func (o *MovieCollection) SetSortTitleNil(b bool)`

 SetSortTitleNil sets the value for SortTitle to be an explicit nil

### UnsetSortTitle
`func (o *MovieCollection) UnsetSortTitle()`

UnsetSortTitle ensures that no value is present for SortTitle, not even an explicit nil
### GetTmdbId

`func (o *MovieCollection) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *MovieCollection) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *MovieCollection) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *MovieCollection) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetOverview

`func (o *MovieCollection) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *MovieCollection) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *MovieCollection) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *MovieCollection) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *MovieCollection) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *MovieCollection) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetMonitored

`func (o *MovieCollection) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *MovieCollection) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *MovieCollection) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *MovieCollection) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetQualityProfileId

`func (o *MovieCollection) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *MovieCollection) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *MovieCollection) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *MovieCollection) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetRootFolderPath

`func (o *MovieCollection) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *MovieCollection) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *MovieCollection) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *MovieCollection) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *MovieCollection) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *MovieCollection) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetSearchOnAdd

`func (o *MovieCollection) GetSearchOnAdd() bool`

GetSearchOnAdd returns the SearchOnAdd field if non-nil, zero value otherwise.

### GetSearchOnAddOk

`func (o *MovieCollection) GetSearchOnAddOk() (*bool, bool)`

GetSearchOnAddOk returns a tuple with the SearchOnAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchOnAdd

`func (o *MovieCollection) SetSearchOnAdd(v bool)`

SetSearchOnAdd sets SearchOnAdd field to given value.

### HasSearchOnAdd

`func (o *MovieCollection) HasSearchOnAdd() bool`

HasSearchOnAdd returns a boolean if a field has been set.

### GetMinimumAvailability

`func (o *MovieCollection) GetMinimumAvailability() MovieStatusType`

GetMinimumAvailability returns the MinimumAvailability field if non-nil, zero value otherwise.

### GetMinimumAvailabilityOk

`func (o *MovieCollection) GetMinimumAvailabilityOk() (*MovieStatusType, bool)`

GetMinimumAvailabilityOk returns a tuple with the MinimumAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAvailability

`func (o *MovieCollection) SetMinimumAvailability(v MovieStatusType)`

SetMinimumAvailability sets MinimumAvailability field to given value.

### HasMinimumAvailability

`func (o *MovieCollection) HasMinimumAvailability() bool`

HasMinimumAvailability returns a boolean if a field has been set.

### GetLastInfoSync

`func (o *MovieCollection) GetLastInfoSync() time.Time`

GetLastInfoSync returns the LastInfoSync field if non-nil, zero value otherwise.

### GetLastInfoSyncOk

`func (o *MovieCollection) GetLastInfoSyncOk() (*time.Time, bool)`

GetLastInfoSyncOk returns a tuple with the LastInfoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInfoSync

`func (o *MovieCollection) SetLastInfoSync(v time.Time)`

SetLastInfoSync sets LastInfoSync field to given value.

### HasLastInfoSync

`func (o *MovieCollection) HasLastInfoSync() bool`

HasLastInfoSync returns a boolean if a field has been set.

### SetLastInfoSyncNil

`func (o *MovieCollection) SetLastInfoSyncNil(b bool)`

 SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil

### UnsetLastInfoSync
`func (o *MovieCollection) UnsetLastInfoSync()`

UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
### GetImages

`func (o *MovieCollection) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *MovieCollection) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *MovieCollection) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *MovieCollection) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *MovieCollection) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *MovieCollection) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetAdded

`func (o *MovieCollection) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *MovieCollection) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *MovieCollection) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *MovieCollection) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetMovies

`func (o *MovieCollection) GetMovies() []MovieMetadata`

GetMovies returns the Movies field if non-nil, zero value otherwise.

### GetMoviesOk

`func (o *MovieCollection) GetMoviesOk() (*[]MovieMetadata, bool)`

GetMoviesOk returns a tuple with the Movies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovies

`func (o *MovieCollection) SetMovies(v []MovieMetadata)`

SetMovies sets Movies field to given value.

### HasMovies

`func (o *MovieCollection) HasMovies() bool`

HasMovies returns a boolean if a field has been set.

### SetMoviesNil

`func (o *MovieCollection) SetMoviesNil(b bool)`

 SetMoviesNil sets the value for Movies to be an explicit nil

### UnsetMovies
`func (o *MovieCollection) UnsetMovies()`

UnsetMovies ensures that no value is present for Movies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



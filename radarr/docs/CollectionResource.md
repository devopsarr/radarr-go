# CollectionResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**SortTitle** | Pointer to **NullableString** |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**QualityProfileId** | Pointer to **int32** |  | [optional] 
**SearchOnAdd** | Pointer to **bool** |  | [optional] 
**MinimumAvailability** | Pointer to [**MovieStatusType**](MovieStatusType.md) |  | [optional] 
**Movies** | Pointer to [**[]CollectionMovieResource**](CollectionMovieResource.md) |  | [optional] 
**MissingMovies** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCollectionResource

`func NewCollectionResource() *CollectionResource`

NewCollectionResource instantiates a new CollectionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourceWithDefaults

`func NewCollectionResourceWithDefaults() *CollectionResource`

NewCollectionResourceWithDefaults instantiates a new CollectionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CollectionResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *CollectionResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CollectionResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CollectionResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CollectionResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CollectionResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CollectionResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSortTitle

`func (o *CollectionResource) GetSortTitle() string`

GetSortTitle returns the SortTitle field if non-nil, zero value otherwise.

### GetSortTitleOk

`func (o *CollectionResource) GetSortTitleOk() (*string, bool)`

GetSortTitleOk returns a tuple with the SortTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortTitle

`func (o *CollectionResource) SetSortTitle(v string)`

SetSortTitle sets SortTitle field to given value.

### HasSortTitle

`func (o *CollectionResource) HasSortTitle() bool`

HasSortTitle returns a boolean if a field has been set.

### SetSortTitleNil

`func (o *CollectionResource) SetSortTitleNil(b bool)`

 SetSortTitleNil sets the value for SortTitle to be an explicit nil

### UnsetSortTitle
`func (o *CollectionResource) UnsetSortTitle()`

UnsetSortTitle ensures that no value is present for SortTitle, not even an explicit nil
### GetTmdbId

`func (o *CollectionResource) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *CollectionResource) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *CollectionResource) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *CollectionResource) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetImages

`func (o *CollectionResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CollectionResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CollectionResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *CollectionResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CollectionResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CollectionResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetOverview

`func (o *CollectionResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *CollectionResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *CollectionResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *CollectionResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *CollectionResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *CollectionResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetMonitored

`func (o *CollectionResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *CollectionResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *CollectionResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *CollectionResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetRootFolderPath

`func (o *CollectionResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *CollectionResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *CollectionResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *CollectionResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *CollectionResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *CollectionResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetQualityProfileId

`func (o *CollectionResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *CollectionResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *CollectionResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *CollectionResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetSearchOnAdd

`func (o *CollectionResource) GetSearchOnAdd() bool`

GetSearchOnAdd returns the SearchOnAdd field if non-nil, zero value otherwise.

### GetSearchOnAddOk

`func (o *CollectionResource) GetSearchOnAddOk() (*bool, bool)`

GetSearchOnAddOk returns a tuple with the SearchOnAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchOnAdd

`func (o *CollectionResource) SetSearchOnAdd(v bool)`

SetSearchOnAdd sets SearchOnAdd field to given value.

### HasSearchOnAdd

`func (o *CollectionResource) HasSearchOnAdd() bool`

HasSearchOnAdd returns a boolean if a field has been set.

### GetMinimumAvailability

`func (o *CollectionResource) GetMinimumAvailability() MovieStatusType`

GetMinimumAvailability returns the MinimumAvailability field if non-nil, zero value otherwise.

### GetMinimumAvailabilityOk

`func (o *CollectionResource) GetMinimumAvailabilityOk() (*MovieStatusType, bool)`

GetMinimumAvailabilityOk returns a tuple with the MinimumAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAvailability

`func (o *CollectionResource) SetMinimumAvailability(v MovieStatusType)`

SetMinimumAvailability sets MinimumAvailability field to given value.

### HasMinimumAvailability

`func (o *CollectionResource) HasMinimumAvailability() bool`

HasMinimumAvailability returns a boolean if a field has been set.

### GetMovies

`func (o *CollectionResource) GetMovies() []CollectionMovieResource`

GetMovies returns the Movies field if non-nil, zero value otherwise.

### GetMoviesOk

`func (o *CollectionResource) GetMoviesOk() (*[]CollectionMovieResource, bool)`

GetMoviesOk returns a tuple with the Movies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovies

`func (o *CollectionResource) SetMovies(v []CollectionMovieResource)`

SetMovies sets Movies field to given value.

### HasMovies

`func (o *CollectionResource) HasMovies() bool`

HasMovies returns a boolean if a field has been set.

### SetMoviesNil

`func (o *CollectionResource) SetMoviesNil(b bool)`

 SetMoviesNil sets the value for Movies to be an explicit nil

### UnsetMovies
`func (o *CollectionResource) UnsetMovies()`

UnsetMovies ensures that no value is present for Movies, not even an explicit nil
### GetMissingMovies

`func (o *CollectionResource) GetMissingMovies() int32`

GetMissingMovies returns the MissingMovies field if non-nil, zero value otherwise.

### GetMissingMoviesOk

`func (o *CollectionResource) GetMissingMoviesOk() (*int32, bool)`

GetMissingMoviesOk returns a tuple with the MissingMovies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingMovies

`func (o *CollectionResource) SetMissingMovies(v int32)`

SetMissingMovies sets MissingMovies field to given value.

### HasMissingMovies

`func (o *CollectionResource) HasMissingMovies() bool`

HasMissingMovies returns a boolean if a field has been set.

### GetTags

`func (o *CollectionResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CollectionResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CollectionResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CollectionResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CollectionResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CollectionResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



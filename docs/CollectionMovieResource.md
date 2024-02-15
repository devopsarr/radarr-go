# CollectionMovieResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TmdbId** | Pointer to **int32** |  | [optional] 
**ImdbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**SortTitle** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**MovieStatusType**](MovieStatusType.md) |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**Runtime** | Pointer to **int32** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Folder** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCollectionMovieResource

`func NewCollectionMovieResource() *CollectionMovieResource`

NewCollectionMovieResource instantiates a new CollectionMovieResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionMovieResourceWithDefaults

`func NewCollectionMovieResourceWithDefaults() *CollectionMovieResource`

NewCollectionMovieResourceWithDefaults instantiates a new CollectionMovieResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTmdbId

`func (o *CollectionMovieResource) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *CollectionMovieResource) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *CollectionMovieResource) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *CollectionMovieResource) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetImdbId

`func (o *CollectionMovieResource) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *CollectionMovieResource) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *CollectionMovieResource) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *CollectionMovieResource) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### SetImdbIdNil

`func (o *CollectionMovieResource) SetImdbIdNil(b bool)`

 SetImdbIdNil sets the value for ImdbId to be an explicit nil

### UnsetImdbId
`func (o *CollectionMovieResource) UnsetImdbId()`

UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
### GetTitle

`func (o *CollectionMovieResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CollectionMovieResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CollectionMovieResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CollectionMovieResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CollectionMovieResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CollectionMovieResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCleanTitle

`func (o *CollectionMovieResource) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *CollectionMovieResource) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *CollectionMovieResource) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *CollectionMovieResource) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *CollectionMovieResource) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *CollectionMovieResource) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetSortTitle

`func (o *CollectionMovieResource) GetSortTitle() string`

GetSortTitle returns the SortTitle field if non-nil, zero value otherwise.

### GetSortTitleOk

`func (o *CollectionMovieResource) GetSortTitleOk() (*string, bool)`

GetSortTitleOk returns a tuple with the SortTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortTitle

`func (o *CollectionMovieResource) SetSortTitle(v string)`

SetSortTitle sets SortTitle field to given value.

### HasSortTitle

`func (o *CollectionMovieResource) HasSortTitle() bool`

HasSortTitle returns a boolean if a field has been set.

### SetSortTitleNil

`func (o *CollectionMovieResource) SetSortTitleNil(b bool)`

 SetSortTitleNil sets the value for SortTitle to be an explicit nil

### UnsetSortTitle
`func (o *CollectionMovieResource) UnsetSortTitle()`

UnsetSortTitle ensures that no value is present for SortTitle, not even an explicit nil
### GetStatus

`func (o *CollectionMovieResource) GetStatus() MovieStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CollectionMovieResource) GetStatusOk() (*MovieStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CollectionMovieResource) SetStatus(v MovieStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CollectionMovieResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOverview

`func (o *CollectionMovieResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *CollectionMovieResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *CollectionMovieResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *CollectionMovieResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *CollectionMovieResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *CollectionMovieResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetRuntime

`func (o *CollectionMovieResource) GetRuntime() int32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *CollectionMovieResource) GetRuntimeOk() (*int32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *CollectionMovieResource) SetRuntime(v int32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *CollectionMovieResource) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetImages

`func (o *CollectionMovieResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CollectionMovieResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CollectionMovieResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *CollectionMovieResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CollectionMovieResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CollectionMovieResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetYear

`func (o *CollectionMovieResource) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *CollectionMovieResource) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *CollectionMovieResource) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *CollectionMovieResource) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetRatings

`func (o *CollectionMovieResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *CollectionMovieResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *CollectionMovieResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *CollectionMovieResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetGenres

`func (o *CollectionMovieResource) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *CollectionMovieResource) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *CollectionMovieResource) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *CollectionMovieResource) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *CollectionMovieResource) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *CollectionMovieResource) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetFolder

`func (o *CollectionMovieResource) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *CollectionMovieResource) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *CollectionMovieResource) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *CollectionMovieResource) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *CollectionMovieResource) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *CollectionMovieResource) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



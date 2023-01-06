# MovieFileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MovieId** | Pointer to **int32** |  | [optional] 
**RelativePath** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**DateAdded** | Pointer to **time.Time** |  | [optional] 
**SceneName** | Pointer to **NullableString** |  | [optional] 
**IndexerFlags** | Pointer to **int32** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfoResource**](MediaInfoResource.md) |  | [optional] 
**OriginalFilePath** | Pointer to **NullableString** |  | [optional] 
**QualityCutoffNotMet** | Pointer to **bool** |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**Edition** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMovieFileResource

`func NewMovieFileResource() *MovieFileResource`

NewMovieFileResource instantiates a new MovieFileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieFileResourceWithDefaults

`func NewMovieFileResourceWithDefaults() *MovieFileResource`

NewMovieFileResourceWithDefaults instantiates a new MovieFileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MovieFileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MovieFileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MovieFileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MovieFileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMovieId

`func (o *MovieFileResource) GetMovieId() int32`

GetMovieId returns the MovieId field if non-nil, zero value otherwise.

### GetMovieIdOk

`func (o *MovieFileResource) GetMovieIdOk() (*int32, bool)`

GetMovieIdOk returns a tuple with the MovieId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieId

`func (o *MovieFileResource) SetMovieId(v int32)`

SetMovieId sets MovieId field to given value.

### HasMovieId

`func (o *MovieFileResource) HasMovieId() bool`

HasMovieId returns a boolean if a field has been set.

### GetRelativePath

`func (o *MovieFileResource) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *MovieFileResource) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *MovieFileResource) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *MovieFileResource) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *MovieFileResource) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *MovieFileResource) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
### GetPath

`func (o *MovieFileResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MovieFileResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MovieFileResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MovieFileResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MovieFileResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MovieFileResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSize

`func (o *MovieFileResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MovieFileResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MovieFileResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MovieFileResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDateAdded

`func (o *MovieFileResource) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *MovieFileResource) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *MovieFileResource) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *MovieFileResource) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetSceneName

`func (o *MovieFileResource) GetSceneName() string`

GetSceneName returns the SceneName field if non-nil, zero value otherwise.

### GetSceneNameOk

`func (o *MovieFileResource) GetSceneNameOk() (*string, bool)`

GetSceneNameOk returns a tuple with the SceneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneName

`func (o *MovieFileResource) SetSceneName(v string)`

SetSceneName sets SceneName field to given value.

### HasSceneName

`func (o *MovieFileResource) HasSceneName() bool`

HasSceneName returns a boolean if a field has been set.

### SetSceneNameNil

`func (o *MovieFileResource) SetSceneNameNil(b bool)`

 SetSceneNameNil sets the value for SceneName to be an explicit nil

### UnsetSceneName
`func (o *MovieFileResource) UnsetSceneName()`

UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
### GetIndexerFlags

`func (o *MovieFileResource) GetIndexerFlags() int32`

GetIndexerFlags returns the IndexerFlags field if non-nil, zero value otherwise.

### GetIndexerFlagsOk

`func (o *MovieFileResource) GetIndexerFlagsOk() (*int32, bool)`

GetIndexerFlagsOk returns a tuple with the IndexerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerFlags

`func (o *MovieFileResource) SetIndexerFlags(v int32)`

SetIndexerFlags sets IndexerFlags field to given value.

### HasIndexerFlags

`func (o *MovieFileResource) HasIndexerFlags() bool`

HasIndexerFlags returns a boolean if a field has been set.

### GetQuality

`func (o *MovieFileResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *MovieFileResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *MovieFileResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *MovieFileResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetCustomFormats

`func (o *MovieFileResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *MovieFileResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *MovieFileResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *MovieFileResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *MovieFileResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *MovieFileResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetMediaInfo

`func (o *MovieFileResource) GetMediaInfo() MediaInfoResource`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *MovieFileResource) GetMediaInfoOk() (*MediaInfoResource, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *MovieFileResource) SetMediaInfo(v MediaInfoResource)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *MovieFileResource) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetOriginalFilePath

`func (o *MovieFileResource) GetOriginalFilePath() string`

GetOriginalFilePath returns the OriginalFilePath field if non-nil, zero value otherwise.

### GetOriginalFilePathOk

`func (o *MovieFileResource) GetOriginalFilePathOk() (*string, bool)`

GetOriginalFilePathOk returns a tuple with the OriginalFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFilePath

`func (o *MovieFileResource) SetOriginalFilePath(v string)`

SetOriginalFilePath sets OriginalFilePath field to given value.

### HasOriginalFilePath

`func (o *MovieFileResource) HasOriginalFilePath() bool`

HasOriginalFilePath returns a boolean if a field has been set.

### SetOriginalFilePathNil

`func (o *MovieFileResource) SetOriginalFilePathNil(b bool)`

 SetOriginalFilePathNil sets the value for OriginalFilePath to be an explicit nil

### UnsetOriginalFilePath
`func (o *MovieFileResource) UnsetOriginalFilePath()`

UnsetOriginalFilePath ensures that no value is present for OriginalFilePath, not even an explicit nil
### GetQualityCutoffNotMet

`func (o *MovieFileResource) GetQualityCutoffNotMet() bool`

GetQualityCutoffNotMet returns the QualityCutoffNotMet field if non-nil, zero value otherwise.

### GetQualityCutoffNotMetOk

`func (o *MovieFileResource) GetQualityCutoffNotMetOk() (*bool, bool)`

GetQualityCutoffNotMetOk returns a tuple with the QualityCutoffNotMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityCutoffNotMet

`func (o *MovieFileResource) SetQualityCutoffNotMet(v bool)`

SetQualityCutoffNotMet sets QualityCutoffNotMet field to given value.

### HasQualityCutoffNotMet

`func (o *MovieFileResource) HasQualityCutoffNotMet() bool`

HasQualityCutoffNotMet returns a boolean if a field has been set.

### GetLanguages

`func (o *MovieFileResource) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *MovieFileResource) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *MovieFileResource) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *MovieFileResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *MovieFileResource) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *MovieFileResource) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetReleaseGroup

`func (o *MovieFileResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *MovieFileResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *MovieFileResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *MovieFileResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *MovieFileResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *MovieFileResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetEdition

`func (o *MovieFileResource) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *MovieFileResource) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *MovieFileResource) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *MovieFileResource) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### SetEditionNil

`func (o *MovieFileResource) SetEditionNil(b bool)`

 SetEditionNil sets the value for Edition to be an explicit nil

### UnsetEdition
`func (o *MovieFileResource) UnsetEdition()`

UnsetEdition ensures that no value is present for Edition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



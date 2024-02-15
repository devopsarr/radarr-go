# MovieFileListResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MovieFileIds** | Pointer to **[]int32** |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**Edition** | Pointer to **NullableString** |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**SceneName** | Pointer to **NullableString** |  | [optional] 
**IndexerFlags** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewMovieFileListResource

`func NewMovieFileListResource() *MovieFileListResource`

NewMovieFileListResource instantiates a new MovieFileListResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieFileListResourceWithDefaults

`func NewMovieFileListResourceWithDefaults() *MovieFileListResource`

NewMovieFileListResourceWithDefaults instantiates a new MovieFileListResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovieFileIds

`func (o *MovieFileListResource) GetMovieFileIds() []int32`

GetMovieFileIds returns the MovieFileIds field if non-nil, zero value otherwise.

### GetMovieFileIdsOk

`func (o *MovieFileListResource) GetMovieFileIdsOk() (*[]int32, bool)`

GetMovieFileIdsOk returns a tuple with the MovieFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieFileIds

`func (o *MovieFileListResource) SetMovieFileIds(v []int32)`

SetMovieFileIds sets MovieFileIds field to given value.

### HasMovieFileIds

`func (o *MovieFileListResource) HasMovieFileIds() bool`

HasMovieFileIds returns a boolean if a field has been set.

### SetMovieFileIdsNil

`func (o *MovieFileListResource) SetMovieFileIdsNil(b bool)`

 SetMovieFileIdsNil sets the value for MovieFileIds to be an explicit nil

### UnsetMovieFileIds
`func (o *MovieFileListResource) UnsetMovieFileIds()`

UnsetMovieFileIds ensures that no value is present for MovieFileIds, not even an explicit nil
### GetLanguages

`func (o *MovieFileListResource) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *MovieFileListResource) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *MovieFileListResource) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *MovieFileListResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *MovieFileListResource) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *MovieFileListResource) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetQuality

`func (o *MovieFileListResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *MovieFileListResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *MovieFileListResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *MovieFileListResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetEdition

`func (o *MovieFileListResource) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *MovieFileListResource) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *MovieFileListResource) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *MovieFileListResource) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### SetEditionNil

`func (o *MovieFileListResource) SetEditionNil(b bool)`

 SetEditionNil sets the value for Edition to be an explicit nil

### UnsetEdition
`func (o *MovieFileListResource) UnsetEdition()`

UnsetEdition ensures that no value is present for Edition, not even an explicit nil
### GetReleaseGroup

`func (o *MovieFileListResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *MovieFileListResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *MovieFileListResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *MovieFileListResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *MovieFileListResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *MovieFileListResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetSceneName

`func (o *MovieFileListResource) GetSceneName() string`

GetSceneName returns the SceneName field if non-nil, zero value otherwise.

### GetSceneNameOk

`func (o *MovieFileListResource) GetSceneNameOk() (*string, bool)`

GetSceneNameOk returns a tuple with the SceneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSceneName

`func (o *MovieFileListResource) SetSceneName(v string)`

SetSceneName sets SceneName field to given value.

### HasSceneName

`func (o *MovieFileListResource) HasSceneName() bool`

HasSceneName returns a boolean if a field has been set.

### SetSceneNameNil

`func (o *MovieFileListResource) SetSceneNameNil(b bool)`

 SetSceneNameNil sets the value for SceneName to be an explicit nil

### UnsetSceneName
`func (o *MovieFileListResource) UnsetSceneName()`

UnsetSceneName ensures that no value is present for SceneName, not even an explicit nil
### GetIndexerFlags

`func (o *MovieFileListResource) GetIndexerFlags() int32`

GetIndexerFlags returns the IndexerFlags field if non-nil, zero value otherwise.

### GetIndexerFlagsOk

`func (o *MovieFileListResource) GetIndexerFlagsOk() (*int32, bool)`

GetIndexerFlagsOk returns a tuple with the IndexerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerFlags

`func (o *MovieFileListResource) SetIndexerFlags(v int32)`

SetIndexerFlags sets IndexerFlags field to given value.

### HasIndexerFlags

`func (o *MovieFileListResource) HasIndexerFlags() bool`

HasIndexerFlags returns a boolean if a field has been set.

### SetIndexerFlagsNil

`func (o *MovieFileListResource) SetIndexerFlagsNil(b bool)`

 SetIndexerFlagsNil sets the value for IndexerFlags to be an explicit nil

### UnsetIndexerFlags
`func (o *MovieFileListResource) UnsetIndexerFlags()`

UnsetIndexerFlags ensures that no value is present for IndexerFlags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



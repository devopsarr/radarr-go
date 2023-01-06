# ManualImportResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**RelativePath** | Pointer to **NullableString** |  | [optional] 
**FolderName** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Movie** | Pointer to [**MovieResource**](MovieResource.md) |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**QualityWeight** | Pointer to **int32** |  | [optional] 
**DownloadId** | Pointer to **NullableString** |  | [optional] 
**Rejections** | Pointer to [**[]Rejection**](Rejection.md) |  | [optional] 

## Methods

### NewManualImportResource

`func NewManualImportResource() *ManualImportResource`

NewManualImportResource instantiates a new ManualImportResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualImportResourceWithDefaults

`func NewManualImportResourceWithDefaults() *ManualImportResource`

NewManualImportResourceWithDefaults instantiates a new ManualImportResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualImportResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualImportResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualImportResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManualImportResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *ManualImportResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManualImportResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManualImportResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManualImportResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ManualImportResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ManualImportResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRelativePath

`func (o *ManualImportResource) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *ManualImportResource) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *ManualImportResource) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *ManualImportResource) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *ManualImportResource) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *ManualImportResource) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
### GetFolderName

`func (o *ManualImportResource) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *ManualImportResource) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *ManualImportResource) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *ManualImportResource) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### SetFolderNameNil

`func (o *ManualImportResource) SetFolderNameNil(b bool)`

 SetFolderNameNil sets the value for FolderName to be an explicit nil

### UnsetFolderName
`func (o *ManualImportResource) UnsetFolderName()`

UnsetFolderName ensures that no value is present for FolderName, not even an explicit nil
### GetName

`func (o *ManualImportResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManualImportResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManualImportResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManualImportResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ManualImportResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ManualImportResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *ManualImportResource) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ManualImportResource) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ManualImportResource) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ManualImportResource) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMovie

`func (o *ManualImportResource) GetMovie() MovieResource`

GetMovie returns the Movie field if non-nil, zero value otherwise.

### GetMovieOk

`func (o *ManualImportResource) GetMovieOk() (*MovieResource, bool)`

GetMovieOk returns a tuple with the Movie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovie

`func (o *ManualImportResource) SetMovie(v MovieResource)`

SetMovie sets Movie field to given value.

### HasMovie

`func (o *ManualImportResource) HasMovie() bool`

HasMovie returns a boolean if a field has been set.

### GetQuality

`func (o *ManualImportResource) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ManualImportResource) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ManualImportResource) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ManualImportResource) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetLanguages

`func (o *ManualImportResource) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ManualImportResource) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ManualImportResource) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ManualImportResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *ManualImportResource) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *ManualImportResource) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetReleaseGroup

`func (o *ManualImportResource) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ManualImportResource) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ManualImportResource) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ManualImportResource) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ManualImportResource) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ManualImportResource) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetQualityWeight

`func (o *ManualImportResource) GetQualityWeight() int32`

GetQualityWeight returns the QualityWeight field if non-nil, zero value otherwise.

### GetQualityWeightOk

`func (o *ManualImportResource) GetQualityWeightOk() (*int32, bool)`

GetQualityWeightOk returns a tuple with the QualityWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityWeight

`func (o *ManualImportResource) SetQualityWeight(v int32)`

SetQualityWeight sets QualityWeight field to given value.

### HasQualityWeight

`func (o *ManualImportResource) HasQualityWeight() bool`

HasQualityWeight returns a boolean if a field has been set.

### GetDownloadId

`func (o *ManualImportResource) GetDownloadId() string`

GetDownloadId returns the DownloadId field if non-nil, zero value otherwise.

### GetDownloadIdOk

`func (o *ManualImportResource) GetDownloadIdOk() (*string, bool)`

GetDownloadIdOk returns a tuple with the DownloadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadId

`func (o *ManualImportResource) SetDownloadId(v string)`

SetDownloadId sets DownloadId field to given value.

### HasDownloadId

`func (o *ManualImportResource) HasDownloadId() bool`

HasDownloadId returns a boolean if a field has been set.

### SetDownloadIdNil

`func (o *ManualImportResource) SetDownloadIdNil(b bool)`

 SetDownloadIdNil sets the value for DownloadId to be an explicit nil

### UnsetDownloadId
`func (o *ManualImportResource) UnsetDownloadId()`

UnsetDownloadId ensures that no value is present for DownloadId, not even an explicit nil
### GetRejections

`func (o *ManualImportResource) GetRejections() []Rejection`

GetRejections returns the Rejections field if non-nil, zero value otherwise.

### GetRejectionsOk

`func (o *ManualImportResource) GetRejectionsOk() (*[]Rejection, bool)`

GetRejectionsOk returns a tuple with the Rejections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejections

`func (o *ManualImportResource) SetRejections(v []Rejection)`

SetRejections sets Rejections field to given value.

### HasRejections

`func (o *ManualImportResource) HasRejections() bool`

HasRejections returns a boolean if a field has been set.

### SetRejectionsNil

`func (o *ManualImportResource) SetRejectionsNil(b bool)`

 SetRejectionsNil sets the value for Rejections to be an explicit nil

### UnsetRejections
`func (o *ManualImportResource) UnsetRejections()`

UnsetRejections ensures that no value is present for Rejections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



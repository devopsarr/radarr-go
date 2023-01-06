# MovieEditorResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MovieIds** | Pointer to **[]int32** |  | [optional] 
**Monitored** | Pointer to **NullableBool** |  | [optional] 
**QualityProfileId** | Pointer to **NullableInt32** |  | [optional] 
**MinimumAvailability** | Pointer to [**MovieStatusType**](MovieStatusType.md) |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**ApplyTags** | Pointer to [**ApplyTags**](ApplyTags.md) |  | [optional] 
**MoveFiles** | Pointer to **bool** |  | [optional] 
**DeleteFiles** | Pointer to **bool** |  | [optional] 
**AddImportExclusion** | Pointer to **bool** |  | [optional] 

## Methods

### NewMovieEditorResource

`func NewMovieEditorResource() *MovieEditorResource`

NewMovieEditorResource instantiates a new MovieEditorResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieEditorResourceWithDefaults

`func NewMovieEditorResourceWithDefaults() *MovieEditorResource`

NewMovieEditorResourceWithDefaults instantiates a new MovieEditorResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovieIds

`func (o *MovieEditorResource) GetMovieIds() []int32`

GetMovieIds returns the MovieIds field if non-nil, zero value otherwise.

### GetMovieIdsOk

`func (o *MovieEditorResource) GetMovieIdsOk() (*[]int32, bool)`

GetMovieIdsOk returns a tuple with the MovieIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieIds

`func (o *MovieEditorResource) SetMovieIds(v []int32)`

SetMovieIds sets MovieIds field to given value.

### HasMovieIds

`func (o *MovieEditorResource) HasMovieIds() bool`

HasMovieIds returns a boolean if a field has been set.

### SetMovieIdsNil

`func (o *MovieEditorResource) SetMovieIdsNil(b bool)`

 SetMovieIdsNil sets the value for MovieIds to be an explicit nil

### UnsetMovieIds
`func (o *MovieEditorResource) UnsetMovieIds()`

UnsetMovieIds ensures that no value is present for MovieIds, not even an explicit nil
### GetMonitored

`func (o *MovieEditorResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *MovieEditorResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *MovieEditorResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *MovieEditorResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### SetMonitoredNil

`func (o *MovieEditorResource) SetMonitoredNil(b bool)`

 SetMonitoredNil sets the value for Monitored to be an explicit nil

### UnsetMonitored
`func (o *MovieEditorResource) UnsetMonitored()`

UnsetMonitored ensures that no value is present for Monitored, not even an explicit nil
### GetQualityProfileId

`func (o *MovieEditorResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *MovieEditorResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *MovieEditorResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *MovieEditorResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### SetQualityProfileIdNil

`func (o *MovieEditorResource) SetQualityProfileIdNil(b bool)`

 SetQualityProfileIdNil sets the value for QualityProfileId to be an explicit nil

### UnsetQualityProfileId
`func (o *MovieEditorResource) UnsetQualityProfileId()`

UnsetQualityProfileId ensures that no value is present for QualityProfileId, not even an explicit nil
### GetMinimumAvailability

`func (o *MovieEditorResource) GetMinimumAvailability() MovieStatusType`

GetMinimumAvailability returns the MinimumAvailability field if non-nil, zero value otherwise.

### GetMinimumAvailabilityOk

`func (o *MovieEditorResource) GetMinimumAvailabilityOk() (*MovieStatusType, bool)`

GetMinimumAvailabilityOk returns a tuple with the MinimumAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAvailability

`func (o *MovieEditorResource) SetMinimumAvailability(v MovieStatusType)`

SetMinimumAvailability sets MinimumAvailability field to given value.

### HasMinimumAvailability

`func (o *MovieEditorResource) HasMinimumAvailability() bool`

HasMinimumAvailability returns a boolean if a field has been set.

### GetRootFolderPath

`func (o *MovieEditorResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *MovieEditorResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *MovieEditorResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *MovieEditorResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *MovieEditorResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *MovieEditorResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetTags

`func (o *MovieEditorResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MovieEditorResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MovieEditorResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MovieEditorResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MovieEditorResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MovieEditorResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetApplyTags

`func (o *MovieEditorResource) GetApplyTags() ApplyTags`

GetApplyTags returns the ApplyTags field if non-nil, zero value otherwise.

### GetApplyTagsOk

`func (o *MovieEditorResource) GetApplyTagsOk() (*ApplyTags, bool)`

GetApplyTagsOk returns a tuple with the ApplyTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTags

`func (o *MovieEditorResource) SetApplyTags(v ApplyTags)`

SetApplyTags sets ApplyTags field to given value.

### HasApplyTags

`func (o *MovieEditorResource) HasApplyTags() bool`

HasApplyTags returns a boolean if a field has been set.

### GetMoveFiles

`func (o *MovieEditorResource) GetMoveFiles() bool`

GetMoveFiles returns the MoveFiles field if non-nil, zero value otherwise.

### GetMoveFilesOk

`func (o *MovieEditorResource) GetMoveFilesOk() (*bool, bool)`

GetMoveFilesOk returns a tuple with the MoveFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveFiles

`func (o *MovieEditorResource) SetMoveFiles(v bool)`

SetMoveFiles sets MoveFiles field to given value.

### HasMoveFiles

`func (o *MovieEditorResource) HasMoveFiles() bool`

HasMoveFiles returns a boolean if a field has been set.

### GetDeleteFiles

`func (o *MovieEditorResource) GetDeleteFiles() bool`

GetDeleteFiles returns the DeleteFiles field if non-nil, zero value otherwise.

### GetDeleteFilesOk

`func (o *MovieEditorResource) GetDeleteFilesOk() (*bool, bool)`

GetDeleteFilesOk returns a tuple with the DeleteFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFiles

`func (o *MovieEditorResource) SetDeleteFiles(v bool)`

SetDeleteFiles sets DeleteFiles field to given value.

### HasDeleteFiles

`func (o *MovieEditorResource) HasDeleteFiles() bool`

HasDeleteFiles returns a boolean if a field has been set.

### GetAddImportExclusion

`func (o *MovieEditorResource) GetAddImportExclusion() bool`

GetAddImportExclusion returns the AddImportExclusion field if non-nil, zero value otherwise.

### GetAddImportExclusionOk

`func (o *MovieEditorResource) GetAddImportExclusionOk() (*bool, bool)`

GetAddImportExclusionOk returns a tuple with the AddImportExclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddImportExclusion

`func (o *MovieEditorResource) SetAddImportExclusion(v bool)`

SetAddImportExclusion sets AddImportExclusion field to given value.

### HasAddImportExclusion

`func (o *MovieEditorResource) HasAddImportExclusion() bool`

HasAddImportExclusion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



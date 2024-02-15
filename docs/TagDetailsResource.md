# TagDetailsResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**DelayProfileIds** | Pointer to **[]int32** |  | [optional] 
**ImportListIds** | Pointer to **[]int32** |  | [optional] 
**NotificationIds** | Pointer to **[]int32** |  | [optional] 
**ReleaseProfileIds** | Pointer to **[]int32** |  | [optional] 
**IndexerIds** | Pointer to **[]int32** |  | [optional] 
**DownloadClientIds** | Pointer to **[]int32** |  | [optional] 
**AutoTagIds** | Pointer to **[]int32** |  | [optional] 
**MovieIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewTagDetailsResource

`func NewTagDetailsResource() *TagDetailsResource`

NewTagDetailsResource instantiates a new TagDetailsResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagDetailsResourceWithDefaults

`func NewTagDetailsResourceWithDefaults() *TagDetailsResource`

NewTagDetailsResourceWithDefaults instantiates a new TagDetailsResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagDetailsResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagDetailsResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagDetailsResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TagDetailsResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *TagDetailsResource) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TagDetailsResource) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TagDetailsResource) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TagDetailsResource) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *TagDetailsResource) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *TagDetailsResource) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetDelayProfileIds

`func (o *TagDetailsResource) GetDelayProfileIds() []int32`

GetDelayProfileIds returns the DelayProfileIds field if non-nil, zero value otherwise.

### GetDelayProfileIdsOk

`func (o *TagDetailsResource) GetDelayProfileIdsOk() (*[]int32, bool)`

GetDelayProfileIdsOk returns a tuple with the DelayProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayProfileIds

`func (o *TagDetailsResource) SetDelayProfileIds(v []int32)`

SetDelayProfileIds sets DelayProfileIds field to given value.

### HasDelayProfileIds

`func (o *TagDetailsResource) HasDelayProfileIds() bool`

HasDelayProfileIds returns a boolean if a field has been set.

### SetDelayProfileIdsNil

`func (o *TagDetailsResource) SetDelayProfileIdsNil(b bool)`

 SetDelayProfileIdsNil sets the value for DelayProfileIds to be an explicit nil

### UnsetDelayProfileIds
`func (o *TagDetailsResource) UnsetDelayProfileIds()`

UnsetDelayProfileIds ensures that no value is present for DelayProfileIds, not even an explicit nil
### GetImportListIds

`func (o *TagDetailsResource) GetImportListIds() []int32`

GetImportListIds returns the ImportListIds field if non-nil, zero value otherwise.

### GetImportListIdsOk

`func (o *TagDetailsResource) GetImportListIdsOk() (*[]int32, bool)`

GetImportListIdsOk returns a tuple with the ImportListIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportListIds

`func (o *TagDetailsResource) SetImportListIds(v []int32)`

SetImportListIds sets ImportListIds field to given value.

### HasImportListIds

`func (o *TagDetailsResource) HasImportListIds() bool`

HasImportListIds returns a boolean if a field has been set.

### SetImportListIdsNil

`func (o *TagDetailsResource) SetImportListIdsNil(b bool)`

 SetImportListIdsNil sets the value for ImportListIds to be an explicit nil

### UnsetImportListIds
`func (o *TagDetailsResource) UnsetImportListIds()`

UnsetImportListIds ensures that no value is present for ImportListIds, not even an explicit nil
### GetNotificationIds

`func (o *TagDetailsResource) GetNotificationIds() []int32`

GetNotificationIds returns the NotificationIds field if non-nil, zero value otherwise.

### GetNotificationIdsOk

`func (o *TagDetailsResource) GetNotificationIdsOk() (*[]int32, bool)`

GetNotificationIdsOk returns a tuple with the NotificationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationIds

`func (o *TagDetailsResource) SetNotificationIds(v []int32)`

SetNotificationIds sets NotificationIds field to given value.

### HasNotificationIds

`func (o *TagDetailsResource) HasNotificationIds() bool`

HasNotificationIds returns a boolean if a field has been set.

### SetNotificationIdsNil

`func (o *TagDetailsResource) SetNotificationIdsNil(b bool)`

 SetNotificationIdsNil sets the value for NotificationIds to be an explicit nil

### UnsetNotificationIds
`func (o *TagDetailsResource) UnsetNotificationIds()`

UnsetNotificationIds ensures that no value is present for NotificationIds, not even an explicit nil
### GetReleaseProfileIds

`func (o *TagDetailsResource) GetReleaseProfileIds() []int32`

GetReleaseProfileIds returns the ReleaseProfileIds field if non-nil, zero value otherwise.

### GetReleaseProfileIdsOk

`func (o *TagDetailsResource) GetReleaseProfileIdsOk() (*[]int32, bool)`

GetReleaseProfileIdsOk returns a tuple with the ReleaseProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseProfileIds

`func (o *TagDetailsResource) SetReleaseProfileIds(v []int32)`

SetReleaseProfileIds sets ReleaseProfileIds field to given value.

### HasReleaseProfileIds

`func (o *TagDetailsResource) HasReleaseProfileIds() bool`

HasReleaseProfileIds returns a boolean if a field has been set.

### SetReleaseProfileIdsNil

`func (o *TagDetailsResource) SetReleaseProfileIdsNil(b bool)`

 SetReleaseProfileIdsNil sets the value for ReleaseProfileIds to be an explicit nil

### UnsetReleaseProfileIds
`func (o *TagDetailsResource) UnsetReleaseProfileIds()`

UnsetReleaseProfileIds ensures that no value is present for ReleaseProfileIds, not even an explicit nil
### GetIndexerIds

`func (o *TagDetailsResource) GetIndexerIds() []int32`

GetIndexerIds returns the IndexerIds field if non-nil, zero value otherwise.

### GetIndexerIdsOk

`func (o *TagDetailsResource) GetIndexerIdsOk() (*[]int32, bool)`

GetIndexerIdsOk returns a tuple with the IndexerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerIds

`func (o *TagDetailsResource) SetIndexerIds(v []int32)`

SetIndexerIds sets IndexerIds field to given value.

### HasIndexerIds

`func (o *TagDetailsResource) HasIndexerIds() bool`

HasIndexerIds returns a boolean if a field has been set.

### SetIndexerIdsNil

`func (o *TagDetailsResource) SetIndexerIdsNil(b bool)`

 SetIndexerIdsNil sets the value for IndexerIds to be an explicit nil

### UnsetIndexerIds
`func (o *TagDetailsResource) UnsetIndexerIds()`

UnsetIndexerIds ensures that no value is present for IndexerIds, not even an explicit nil
### GetDownloadClientIds

`func (o *TagDetailsResource) GetDownloadClientIds() []int32`

GetDownloadClientIds returns the DownloadClientIds field if non-nil, zero value otherwise.

### GetDownloadClientIdsOk

`func (o *TagDetailsResource) GetDownloadClientIdsOk() (*[]int32, bool)`

GetDownloadClientIdsOk returns a tuple with the DownloadClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadClientIds

`func (o *TagDetailsResource) SetDownloadClientIds(v []int32)`

SetDownloadClientIds sets DownloadClientIds field to given value.

### HasDownloadClientIds

`func (o *TagDetailsResource) HasDownloadClientIds() bool`

HasDownloadClientIds returns a boolean if a field has been set.

### SetDownloadClientIdsNil

`func (o *TagDetailsResource) SetDownloadClientIdsNil(b bool)`

 SetDownloadClientIdsNil sets the value for DownloadClientIds to be an explicit nil

### UnsetDownloadClientIds
`func (o *TagDetailsResource) UnsetDownloadClientIds()`

UnsetDownloadClientIds ensures that no value is present for DownloadClientIds, not even an explicit nil
### GetAutoTagIds

`func (o *TagDetailsResource) GetAutoTagIds() []int32`

GetAutoTagIds returns the AutoTagIds field if non-nil, zero value otherwise.

### GetAutoTagIdsOk

`func (o *TagDetailsResource) GetAutoTagIdsOk() (*[]int32, bool)`

GetAutoTagIdsOk returns a tuple with the AutoTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTagIds

`func (o *TagDetailsResource) SetAutoTagIds(v []int32)`

SetAutoTagIds sets AutoTagIds field to given value.

### HasAutoTagIds

`func (o *TagDetailsResource) HasAutoTagIds() bool`

HasAutoTagIds returns a boolean if a field has been set.

### SetAutoTagIdsNil

`func (o *TagDetailsResource) SetAutoTagIdsNil(b bool)`

 SetAutoTagIdsNil sets the value for AutoTagIds to be an explicit nil

### UnsetAutoTagIds
`func (o *TagDetailsResource) UnsetAutoTagIds()`

UnsetAutoTagIds ensures that no value is present for AutoTagIds, not even an explicit nil
### GetMovieIds

`func (o *TagDetailsResource) GetMovieIds() []int32`

GetMovieIds returns the MovieIds field if non-nil, zero value otherwise.

### GetMovieIdsOk

`func (o *TagDetailsResource) GetMovieIdsOk() (*[]int32, bool)`

GetMovieIdsOk returns a tuple with the MovieIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieIds

`func (o *TagDetailsResource) SetMovieIds(v []int32)`

SetMovieIds sets MovieIds field to given value.

### HasMovieIds

`func (o *TagDetailsResource) HasMovieIds() bool`

HasMovieIds returns a boolean if a field has been set.

### SetMovieIdsNil

`func (o *TagDetailsResource) SetMovieIdsNil(b bool)`

 SetMovieIdsNil sets the value for MovieIds to be an explicit nil

### UnsetMovieIds
`func (o *TagDetailsResource) UnsetMovieIds()`

UnsetMovieIds ensures that no value is present for MovieIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



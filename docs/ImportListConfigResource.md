# ImportListConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ListSyncLevel** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewImportListConfigResource

`func NewImportListConfigResource() *ImportListConfigResource`

NewImportListConfigResource instantiates a new ImportListConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportListConfigResourceWithDefaults

`func NewImportListConfigResourceWithDefaults() *ImportListConfigResource`

NewImportListConfigResourceWithDefaults instantiates a new ImportListConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportListConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportListConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportListConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ImportListConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListSyncLevel

`func (o *ImportListConfigResource) GetListSyncLevel() string`

GetListSyncLevel returns the ListSyncLevel field if non-nil, zero value otherwise.

### GetListSyncLevelOk

`func (o *ImportListConfigResource) GetListSyncLevelOk() (*string, bool)`

GetListSyncLevelOk returns a tuple with the ListSyncLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListSyncLevel

`func (o *ImportListConfigResource) SetListSyncLevel(v string)`

SetListSyncLevel sets ListSyncLevel field to given value.

### HasListSyncLevel

`func (o *ImportListConfigResource) HasListSyncLevel() bool`

HasListSyncLevel returns a boolean if a field has been set.

### SetListSyncLevelNil

`func (o *ImportListConfigResource) SetListSyncLevelNil(b bool)`

 SetListSyncLevelNil sets the value for ListSyncLevel to be an explicit nil

### UnsetListSyncLevel
`func (o *ImportListConfigResource) UnsetListSyncLevel()`

UnsetListSyncLevel ensures that no value is present for ListSyncLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



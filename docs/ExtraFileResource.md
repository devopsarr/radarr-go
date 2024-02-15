# ExtraFileResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MovieId** | Pointer to **int32** |  | [optional] 
**MovieFileId** | Pointer to **NullableInt32** |  | [optional] 
**RelativePath** | Pointer to **NullableString** |  | [optional] 
**Extension** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**ExtraFileType**](ExtraFileType.md) |  | [optional] 

## Methods

### NewExtraFileResource

`func NewExtraFileResource() *ExtraFileResource`

NewExtraFileResource instantiates a new ExtraFileResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtraFileResourceWithDefaults

`func NewExtraFileResourceWithDefaults() *ExtraFileResource`

NewExtraFileResourceWithDefaults instantiates a new ExtraFileResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtraFileResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtraFileResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtraFileResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExtraFileResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMovieId

`func (o *ExtraFileResource) GetMovieId() int32`

GetMovieId returns the MovieId field if non-nil, zero value otherwise.

### GetMovieIdOk

`func (o *ExtraFileResource) GetMovieIdOk() (*int32, bool)`

GetMovieIdOk returns a tuple with the MovieId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieId

`func (o *ExtraFileResource) SetMovieId(v int32)`

SetMovieId sets MovieId field to given value.

### HasMovieId

`func (o *ExtraFileResource) HasMovieId() bool`

HasMovieId returns a boolean if a field has been set.

### GetMovieFileId

`func (o *ExtraFileResource) GetMovieFileId() int32`

GetMovieFileId returns the MovieFileId field if non-nil, zero value otherwise.

### GetMovieFileIdOk

`func (o *ExtraFileResource) GetMovieFileIdOk() (*int32, bool)`

GetMovieFileIdOk returns a tuple with the MovieFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieFileId

`func (o *ExtraFileResource) SetMovieFileId(v int32)`

SetMovieFileId sets MovieFileId field to given value.

### HasMovieFileId

`func (o *ExtraFileResource) HasMovieFileId() bool`

HasMovieFileId returns a boolean if a field has been set.

### SetMovieFileIdNil

`func (o *ExtraFileResource) SetMovieFileIdNil(b bool)`

 SetMovieFileIdNil sets the value for MovieFileId to be an explicit nil

### UnsetMovieFileId
`func (o *ExtraFileResource) UnsetMovieFileId()`

UnsetMovieFileId ensures that no value is present for MovieFileId, not even an explicit nil
### GetRelativePath

`func (o *ExtraFileResource) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *ExtraFileResource) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *ExtraFileResource) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *ExtraFileResource) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### SetRelativePathNil

`func (o *ExtraFileResource) SetRelativePathNil(b bool)`

 SetRelativePathNil sets the value for RelativePath to be an explicit nil

### UnsetRelativePath
`func (o *ExtraFileResource) UnsetRelativePath()`

UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
### GetExtension

`func (o *ExtraFileResource) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *ExtraFileResource) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *ExtraFileResource) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *ExtraFileResource) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### SetExtensionNil

`func (o *ExtraFileResource) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *ExtraFileResource) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil
### GetType

`func (o *ExtraFileResource) GetType() ExtraFileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtraFileResource) GetTypeOk() (*ExtraFileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtraFileResource) SetType(v ExtraFileType)`

SetType sets Type field to given value.

### HasType

`func (o *ExtraFileResource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FileSystemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**FileSystemEntityType**](FileSystemEntityType.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Extension** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**LastModified** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewFileSystemModel

`func NewFileSystemModel() *FileSystemModel`

NewFileSystemModel instantiates a new FileSystemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSystemModelWithDefaults

`func NewFileSystemModelWithDefaults() *FileSystemModel`

NewFileSystemModelWithDefaults instantiates a new FileSystemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FileSystemModel) GetType() FileSystemEntityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileSystemModel) GetTypeOk() (*FileSystemEntityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileSystemModel) SetType(v FileSystemEntityType)`

SetType sets Type field to given value.

### HasType

`func (o *FileSystemModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *FileSystemModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileSystemModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileSystemModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileSystemModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FileSystemModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FileSystemModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *FileSystemModel) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileSystemModel) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileSystemModel) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileSystemModel) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *FileSystemModel) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *FileSystemModel) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetExtension

`func (o *FileSystemModel) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *FileSystemModel) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *FileSystemModel) SetExtension(v string)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *FileSystemModel) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### SetExtensionNil

`func (o *FileSystemModel) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *FileSystemModel) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil
### GetSize

`func (o *FileSystemModel) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileSystemModel) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileSystemModel) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileSystemModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLastModified

`func (o *FileSystemModel) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *FileSystemModel) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *FileSystemModel) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *FileSystemModel) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *FileSystemModel) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *FileSystemModel) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



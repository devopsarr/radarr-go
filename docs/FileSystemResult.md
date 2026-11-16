# FileSystemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | Pointer to **NullableString** |  | [optional] 
**Directories** | Pointer to [**[]FileSystemModel**](FileSystemModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileSystemModel**](FileSystemModel.md) |  | [optional] 

## Methods

### NewFileSystemResult

`func NewFileSystemResult() *FileSystemResult`

NewFileSystemResult instantiates a new FileSystemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSystemResultWithDefaults

`func NewFileSystemResultWithDefaults() *FileSystemResult`

NewFileSystemResultWithDefaults instantiates a new FileSystemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *FileSystemResult) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FileSystemResult) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FileSystemResult) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FileSystemResult) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *FileSystemResult) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *FileSystemResult) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetDirectories

`func (o *FileSystemResult) GetDirectories() []FileSystemModel`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *FileSystemResult) GetDirectoriesOk() (*[]FileSystemModel, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *FileSystemResult) SetDirectories(v []FileSystemModel)`

SetDirectories sets Directories field to given value.

### HasDirectories

`func (o *FileSystemResult) HasDirectories() bool`

HasDirectories returns a boolean if a field has been set.

### SetDirectoriesNil

`func (o *FileSystemResult) SetDirectoriesNil(b bool)`

 SetDirectoriesNil sets the value for Directories to be an explicit nil

### UnsetDirectories
`func (o *FileSystemResult) UnsetDirectories()`

UnsetDirectories ensures that no value is present for Directories, not even an explicit nil
### GetFiles

`func (o *FileSystemResult) GetFiles() []FileSystemModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FileSystemResult) GetFilesOk() (*[]FileSystemModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FileSystemResult) SetFiles(v []FileSystemModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FileSystemResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *FileSystemResult) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *FileSystemResult) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



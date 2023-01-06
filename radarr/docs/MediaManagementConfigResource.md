# MediaManagementConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AutoUnmonitorPreviouslyDownloadedMovies** | Pointer to **bool** |  | [optional] 
**RecycleBin** | Pointer to **NullableString** |  | [optional] 
**RecycleBinCleanupDays** | Pointer to **int32** |  | [optional] 
**DownloadPropersAndRepacks** | Pointer to [**ProperDownloadTypes**](ProperDownloadTypes.md) |  | [optional] 
**CreateEmptyMovieFolders** | Pointer to **bool** |  | [optional] 
**DeleteEmptyFolders** | Pointer to **bool** |  | [optional] 
**FileDate** | Pointer to [**FileDateType**](FileDateType.md) |  | [optional] 
**RescanAfterRefresh** | Pointer to [**RescanAfterRefreshType**](RescanAfterRefreshType.md) |  | [optional] 
**AutoRenameFolders** | Pointer to **bool** |  | [optional] 
**PathsDefaultStatic** | Pointer to **bool** |  | [optional] 
**SetPermissionsLinux** | Pointer to **bool** |  | [optional] 
**ChmodFolder** | Pointer to **NullableString** |  | [optional] 
**ChownGroup** | Pointer to **NullableString** |  | [optional] 
**SkipFreeSpaceCheckWhenImporting** | Pointer to **bool** |  | [optional] 
**MinimumFreeSpaceWhenImporting** | Pointer to **int32** |  | [optional] 
**CopyUsingHardlinks** | Pointer to **bool** |  | [optional] 
**ImportExtraFiles** | Pointer to **bool** |  | [optional] 
**ExtraFileExtensions** | Pointer to **NullableString** |  | [optional] 
**EnableMediaInfo** | Pointer to **bool** |  | [optional] 

## Methods

### NewMediaManagementConfigResource

`func NewMediaManagementConfigResource() *MediaManagementConfigResource`

NewMediaManagementConfigResource instantiates a new MediaManagementConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaManagementConfigResourceWithDefaults

`func NewMediaManagementConfigResourceWithDefaults() *MediaManagementConfigResource`

NewMediaManagementConfigResourceWithDefaults instantiates a new MediaManagementConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaManagementConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaManagementConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaManagementConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MediaManagementConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAutoUnmonitorPreviouslyDownloadedMovies

`func (o *MediaManagementConfigResource) GetAutoUnmonitorPreviouslyDownloadedMovies() bool`

GetAutoUnmonitorPreviouslyDownloadedMovies returns the AutoUnmonitorPreviouslyDownloadedMovies field if non-nil, zero value otherwise.

### GetAutoUnmonitorPreviouslyDownloadedMoviesOk

`func (o *MediaManagementConfigResource) GetAutoUnmonitorPreviouslyDownloadedMoviesOk() (*bool, bool)`

GetAutoUnmonitorPreviouslyDownloadedMoviesOk returns a tuple with the AutoUnmonitorPreviouslyDownloadedMovies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUnmonitorPreviouslyDownloadedMovies

`func (o *MediaManagementConfigResource) SetAutoUnmonitorPreviouslyDownloadedMovies(v bool)`

SetAutoUnmonitorPreviouslyDownloadedMovies sets AutoUnmonitorPreviouslyDownloadedMovies field to given value.

### HasAutoUnmonitorPreviouslyDownloadedMovies

`func (o *MediaManagementConfigResource) HasAutoUnmonitorPreviouslyDownloadedMovies() bool`

HasAutoUnmonitorPreviouslyDownloadedMovies returns a boolean if a field has been set.

### GetRecycleBin

`func (o *MediaManagementConfigResource) GetRecycleBin() string`

GetRecycleBin returns the RecycleBin field if non-nil, zero value otherwise.

### GetRecycleBinOk

`func (o *MediaManagementConfigResource) GetRecycleBinOk() (*string, bool)`

GetRecycleBinOk returns a tuple with the RecycleBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleBin

`func (o *MediaManagementConfigResource) SetRecycleBin(v string)`

SetRecycleBin sets RecycleBin field to given value.

### HasRecycleBin

`func (o *MediaManagementConfigResource) HasRecycleBin() bool`

HasRecycleBin returns a boolean if a field has been set.

### SetRecycleBinNil

`func (o *MediaManagementConfigResource) SetRecycleBinNil(b bool)`

 SetRecycleBinNil sets the value for RecycleBin to be an explicit nil

### UnsetRecycleBin
`func (o *MediaManagementConfigResource) UnsetRecycleBin()`

UnsetRecycleBin ensures that no value is present for RecycleBin, not even an explicit nil
### GetRecycleBinCleanupDays

`func (o *MediaManagementConfigResource) GetRecycleBinCleanupDays() int32`

GetRecycleBinCleanupDays returns the RecycleBinCleanupDays field if non-nil, zero value otherwise.

### GetRecycleBinCleanupDaysOk

`func (o *MediaManagementConfigResource) GetRecycleBinCleanupDaysOk() (*int32, bool)`

GetRecycleBinCleanupDaysOk returns a tuple with the RecycleBinCleanupDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleBinCleanupDays

`func (o *MediaManagementConfigResource) SetRecycleBinCleanupDays(v int32)`

SetRecycleBinCleanupDays sets RecycleBinCleanupDays field to given value.

### HasRecycleBinCleanupDays

`func (o *MediaManagementConfigResource) HasRecycleBinCleanupDays() bool`

HasRecycleBinCleanupDays returns a boolean if a field has been set.

### GetDownloadPropersAndRepacks

`func (o *MediaManagementConfigResource) GetDownloadPropersAndRepacks() ProperDownloadTypes`

GetDownloadPropersAndRepacks returns the DownloadPropersAndRepacks field if non-nil, zero value otherwise.

### GetDownloadPropersAndRepacksOk

`func (o *MediaManagementConfigResource) GetDownloadPropersAndRepacksOk() (*ProperDownloadTypes, bool)`

GetDownloadPropersAndRepacksOk returns a tuple with the DownloadPropersAndRepacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPropersAndRepacks

`func (o *MediaManagementConfigResource) SetDownloadPropersAndRepacks(v ProperDownloadTypes)`

SetDownloadPropersAndRepacks sets DownloadPropersAndRepacks field to given value.

### HasDownloadPropersAndRepacks

`func (o *MediaManagementConfigResource) HasDownloadPropersAndRepacks() bool`

HasDownloadPropersAndRepacks returns a boolean if a field has been set.

### GetCreateEmptyMovieFolders

`func (o *MediaManagementConfigResource) GetCreateEmptyMovieFolders() bool`

GetCreateEmptyMovieFolders returns the CreateEmptyMovieFolders field if non-nil, zero value otherwise.

### GetCreateEmptyMovieFoldersOk

`func (o *MediaManagementConfigResource) GetCreateEmptyMovieFoldersOk() (*bool, bool)`

GetCreateEmptyMovieFoldersOk returns a tuple with the CreateEmptyMovieFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEmptyMovieFolders

`func (o *MediaManagementConfigResource) SetCreateEmptyMovieFolders(v bool)`

SetCreateEmptyMovieFolders sets CreateEmptyMovieFolders field to given value.

### HasCreateEmptyMovieFolders

`func (o *MediaManagementConfigResource) HasCreateEmptyMovieFolders() bool`

HasCreateEmptyMovieFolders returns a boolean if a field has been set.

### GetDeleteEmptyFolders

`func (o *MediaManagementConfigResource) GetDeleteEmptyFolders() bool`

GetDeleteEmptyFolders returns the DeleteEmptyFolders field if non-nil, zero value otherwise.

### GetDeleteEmptyFoldersOk

`func (o *MediaManagementConfigResource) GetDeleteEmptyFoldersOk() (*bool, bool)`

GetDeleteEmptyFoldersOk returns a tuple with the DeleteEmptyFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteEmptyFolders

`func (o *MediaManagementConfigResource) SetDeleteEmptyFolders(v bool)`

SetDeleteEmptyFolders sets DeleteEmptyFolders field to given value.

### HasDeleteEmptyFolders

`func (o *MediaManagementConfigResource) HasDeleteEmptyFolders() bool`

HasDeleteEmptyFolders returns a boolean if a field has been set.

### GetFileDate

`func (o *MediaManagementConfigResource) GetFileDate() FileDateType`

GetFileDate returns the FileDate field if non-nil, zero value otherwise.

### GetFileDateOk

`func (o *MediaManagementConfigResource) GetFileDateOk() (*FileDateType, bool)`

GetFileDateOk returns a tuple with the FileDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDate

`func (o *MediaManagementConfigResource) SetFileDate(v FileDateType)`

SetFileDate sets FileDate field to given value.

### HasFileDate

`func (o *MediaManagementConfigResource) HasFileDate() bool`

HasFileDate returns a boolean if a field has been set.

### GetRescanAfterRefresh

`func (o *MediaManagementConfigResource) GetRescanAfterRefresh() RescanAfterRefreshType`

GetRescanAfterRefresh returns the RescanAfterRefresh field if non-nil, zero value otherwise.

### GetRescanAfterRefreshOk

`func (o *MediaManagementConfigResource) GetRescanAfterRefreshOk() (*RescanAfterRefreshType, bool)`

GetRescanAfterRefreshOk returns a tuple with the RescanAfterRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescanAfterRefresh

`func (o *MediaManagementConfigResource) SetRescanAfterRefresh(v RescanAfterRefreshType)`

SetRescanAfterRefresh sets RescanAfterRefresh field to given value.

### HasRescanAfterRefresh

`func (o *MediaManagementConfigResource) HasRescanAfterRefresh() bool`

HasRescanAfterRefresh returns a boolean if a field has been set.

### GetAutoRenameFolders

`func (o *MediaManagementConfigResource) GetAutoRenameFolders() bool`

GetAutoRenameFolders returns the AutoRenameFolders field if non-nil, zero value otherwise.

### GetAutoRenameFoldersOk

`func (o *MediaManagementConfigResource) GetAutoRenameFoldersOk() (*bool, bool)`

GetAutoRenameFoldersOk returns a tuple with the AutoRenameFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenameFolders

`func (o *MediaManagementConfigResource) SetAutoRenameFolders(v bool)`

SetAutoRenameFolders sets AutoRenameFolders field to given value.

### HasAutoRenameFolders

`func (o *MediaManagementConfigResource) HasAutoRenameFolders() bool`

HasAutoRenameFolders returns a boolean if a field has been set.

### GetPathsDefaultStatic

`func (o *MediaManagementConfigResource) GetPathsDefaultStatic() bool`

GetPathsDefaultStatic returns the PathsDefaultStatic field if non-nil, zero value otherwise.

### GetPathsDefaultStaticOk

`func (o *MediaManagementConfigResource) GetPathsDefaultStaticOk() (*bool, bool)`

GetPathsDefaultStaticOk returns a tuple with the PathsDefaultStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathsDefaultStatic

`func (o *MediaManagementConfigResource) SetPathsDefaultStatic(v bool)`

SetPathsDefaultStatic sets PathsDefaultStatic field to given value.

### HasPathsDefaultStatic

`func (o *MediaManagementConfigResource) HasPathsDefaultStatic() bool`

HasPathsDefaultStatic returns a boolean if a field has been set.

### GetSetPermissionsLinux

`func (o *MediaManagementConfigResource) GetSetPermissionsLinux() bool`

GetSetPermissionsLinux returns the SetPermissionsLinux field if non-nil, zero value otherwise.

### GetSetPermissionsLinuxOk

`func (o *MediaManagementConfigResource) GetSetPermissionsLinuxOk() (*bool, bool)`

GetSetPermissionsLinuxOk returns a tuple with the SetPermissionsLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPermissionsLinux

`func (o *MediaManagementConfigResource) SetSetPermissionsLinux(v bool)`

SetSetPermissionsLinux sets SetPermissionsLinux field to given value.

### HasSetPermissionsLinux

`func (o *MediaManagementConfigResource) HasSetPermissionsLinux() bool`

HasSetPermissionsLinux returns a boolean if a field has been set.

### GetChmodFolder

`func (o *MediaManagementConfigResource) GetChmodFolder() string`

GetChmodFolder returns the ChmodFolder field if non-nil, zero value otherwise.

### GetChmodFolderOk

`func (o *MediaManagementConfigResource) GetChmodFolderOk() (*string, bool)`

GetChmodFolderOk returns a tuple with the ChmodFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChmodFolder

`func (o *MediaManagementConfigResource) SetChmodFolder(v string)`

SetChmodFolder sets ChmodFolder field to given value.

### HasChmodFolder

`func (o *MediaManagementConfigResource) HasChmodFolder() bool`

HasChmodFolder returns a boolean if a field has been set.

### SetChmodFolderNil

`func (o *MediaManagementConfigResource) SetChmodFolderNil(b bool)`

 SetChmodFolderNil sets the value for ChmodFolder to be an explicit nil

### UnsetChmodFolder
`func (o *MediaManagementConfigResource) UnsetChmodFolder()`

UnsetChmodFolder ensures that no value is present for ChmodFolder, not even an explicit nil
### GetChownGroup

`func (o *MediaManagementConfigResource) GetChownGroup() string`

GetChownGroup returns the ChownGroup field if non-nil, zero value otherwise.

### GetChownGroupOk

`func (o *MediaManagementConfigResource) GetChownGroupOk() (*string, bool)`

GetChownGroupOk returns a tuple with the ChownGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChownGroup

`func (o *MediaManagementConfigResource) SetChownGroup(v string)`

SetChownGroup sets ChownGroup field to given value.

### HasChownGroup

`func (o *MediaManagementConfigResource) HasChownGroup() bool`

HasChownGroup returns a boolean if a field has been set.

### SetChownGroupNil

`func (o *MediaManagementConfigResource) SetChownGroupNil(b bool)`

 SetChownGroupNil sets the value for ChownGroup to be an explicit nil

### UnsetChownGroup
`func (o *MediaManagementConfigResource) UnsetChownGroup()`

UnsetChownGroup ensures that no value is present for ChownGroup, not even an explicit nil
### GetSkipFreeSpaceCheckWhenImporting

`func (o *MediaManagementConfigResource) GetSkipFreeSpaceCheckWhenImporting() bool`

GetSkipFreeSpaceCheckWhenImporting returns the SkipFreeSpaceCheckWhenImporting field if non-nil, zero value otherwise.

### GetSkipFreeSpaceCheckWhenImportingOk

`func (o *MediaManagementConfigResource) GetSkipFreeSpaceCheckWhenImportingOk() (*bool, bool)`

GetSkipFreeSpaceCheckWhenImportingOk returns a tuple with the SkipFreeSpaceCheckWhenImporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipFreeSpaceCheckWhenImporting

`func (o *MediaManagementConfigResource) SetSkipFreeSpaceCheckWhenImporting(v bool)`

SetSkipFreeSpaceCheckWhenImporting sets SkipFreeSpaceCheckWhenImporting field to given value.

### HasSkipFreeSpaceCheckWhenImporting

`func (o *MediaManagementConfigResource) HasSkipFreeSpaceCheckWhenImporting() bool`

HasSkipFreeSpaceCheckWhenImporting returns a boolean if a field has been set.

### GetMinimumFreeSpaceWhenImporting

`func (o *MediaManagementConfigResource) GetMinimumFreeSpaceWhenImporting() int32`

GetMinimumFreeSpaceWhenImporting returns the MinimumFreeSpaceWhenImporting field if non-nil, zero value otherwise.

### GetMinimumFreeSpaceWhenImportingOk

`func (o *MediaManagementConfigResource) GetMinimumFreeSpaceWhenImportingOk() (*int32, bool)`

GetMinimumFreeSpaceWhenImportingOk returns a tuple with the MinimumFreeSpaceWhenImporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFreeSpaceWhenImporting

`func (o *MediaManagementConfigResource) SetMinimumFreeSpaceWhenImporting(v int32)`

SetMinimumFreeSpaceWhenImporting sets MinimumFreeSpaceWhenImporting field to given value.

### HasMinimumFreeSpaceWhenImporting

`func (o *MediaManagementConfigResource) HasMinimumFreeSpaceWhenImporting() bool`

HasMinimumFreeSpaceWhenImporting returns a boolean if a field has been set.

### GetCopyUsingHardlinks

`func (o *MediaManagementConfigResource) GetCopyUsingHardlinks() bool`

GetCopyUsingHardlinks returns the CopyUsingHardlinks field if non-nil, zero value otherwise.

### GetCopyUsingHardlinksOk

`func (o *MediaManagementConfigResource) GetCopyUsingHardlinksOk() (*bool, bool)`

GetCopyUsingHardlinksOk returns a tuple with the CopyUsingHardlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyUsingHardlinks

`func (o *MediaManagementConfigResource) SetCopyUsingHardlinks(v bool)`

SetCopyUsingHardlinks sets CopyUsingHardlinks field to given value.

### HasCopyUsingHardlinks

`func (o *MediaManagementConfigResource) HasCopyUsingHardlinks() bool`

HasCopyUsingHardlinks returns a boolean if a field has been set.

### GetImportExtraFiles

`func (o *MediaManagementConfigResource) GetImportExtraFiles() bool`

GetImportExtraFiles returns the ImportExtraFiles field if non-nil, zero value otherwise.

### GetImportExtraFilesOk

`func (o *MediaManagementConfigResource) GetImportExtraFilesOk() (*bool, bool)`

GetImportExtraFilesOk returns a tuple with the ImportExtraFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExtraFiles

`func (o *MediaManagementConfigResource) SetImportExtraFiles(v bool)`

SetImportExtraFiles sets ImportExtraFiles field to given value.

### HasImportExtraFiles

`func (o *MediaManagementConfigResource) HasImportExtraFiles() bool`

HasImportExtraFiles returns a boolean if a field has been set.

### GetExtraFileExtensions

`func (o *MediaManagementConfigResource) GetExtraFileExtensions() string`

GetExtraFileExtensions returns the ExtraFileExtensions field if non-nil, zero value otherwise.

### GetExtraFileExtensionsOk

`func (o *MediaManagementConfigResource) GetExtraFileExtensionsOk() (*string, bool)`

GetExtraFileExtensionsOk returns a tuple with the ExtraFileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFileExtensions

`func (o *MediaManagementConfigResource) SetExtraFileExtensions(v string)`

SetExtraFileExtensions sets ExtraFileExtensions field to given value.

### HasExtraFileExtensions

`func (o *MediaManagementConfigResource) HasExtraFileExtensions() bool`

HasExtraFileExtensions returns a boolean if a field has been set.

### SetExtraFileExtensionsNil

`func (o *MediaManagementConfigResource) SetExtraFileExtensionsNil(b bool)`

 SetExtraFileExtensionsNil sets the value for ExtraFileExtensions to be an explicit nil

### UnsetExtraFileExtensions
`func (o *MediaManagementConfigResource) UnsetExtraFileExtensions()`

UnsetExtraFileExtensions ensures that no value is present for ExtraFileExtensions, not even an explicit nil
### GetEnableMediaInfo

`func (o *MediaManagementConfigResource) GetEnableMediaInfo() bool`

GetEnableMediaInfo returns the EnableMediaInfo field if non-nil, zero value otherwise.

### GetEnableMediaInfoOk

`func (o *MediaManagementConfigResource) GetEnableMediaInfoOk() (*bool, bool)`

GetEnableMediaInfoOk returns a tuple with the EnableMediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaInfo

`func (o *MediaManagementConfigResource) SetEnableMediaInfo(v bool)`

SetEnableMediaInfo sets EnableMediaInfo field to given value.

### HasEnableMediaInfo

`func (o *MediaManagementConfigResource) HasEnableMediaInfo() bool`

HasEnableMediaInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



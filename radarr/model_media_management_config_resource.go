/*
Radarr

Radarr API docs

API version: v5.2.6.8376
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// MediaManagementConfigResource struct for MediaManagementConfigResource
type MediaManagementConfigResource struct {
	Id *int32 `json:"id,omitempty"`
	AutoUnmonitorPreviouslyDownloadedMovies *bool `json:"autoUnmonitorPreviouslyDownloadedMovies,omitempty"`
	RecycleBin NullableString `json:"recycleBin,omitempty"`
	RecycleBinCleanupDays *int32 `json:"recycleBinCleanupDays,omitempty"`
	DownloadPropersAndRepacks *ProperDownloadTypes `json:"downloadPropersAndRepacks,omitempty"`
	CreateEmptyMovieFolders *bool `json:"createEmptyMovieFolders,omitempty"`
	DeleteEmptyFolders *bool `json:"deleteEmptyFolders,omitempty"`
	FileDate *FileDateType `json:"fileDate,omitempty"`
	RescanAfterRefresh *RescanAfterRefreshType `json:"rescanAfterRefresh,omitempty"`
	AutoRenameFolders *bool `json:"autoRenameFolders,omitempty"`
	PathsDefaultStatic *bool `json:"pathsDefaultStatic,omitempty"`
	SetPermissionsLinux *bool `json:"setPermissionsLinux,omitempty"`
	ChmodFolder NullableString `json:"chmodFolder,omitempty"`
	ChownGroup NullableString `json:"chownGroup,omitempty"`
	SkipFreeSpaceCheckWhenImporting *bool `json:"skipFreeSpaceCheckWhenImporting,omitempty"`
	MinimumFreeSpaceWhenImporting *int32 `json:"minimumFreeSpaceWhenImporting,omitempty"`
	CopyUsingHardlinks *bool `json:"copyUsingHardlinks,omitempty"`
	UseScriptImport *bool `json:"useScriptImport,omitempty"`
	ScriptImportPath NullableString `json:"scriptImportPath,omitempty"`
	ImportExtraFiles *bool `json:"importExtraFiles,omitempty"`
	ExtraFileExtensions NullableString `json:"extraFileExtensions,omitempty"`
	EnableMediaInfo *bool `json:"enableMediaInfo,omitempty"`
}

// NewMediaManagementConfigResource instantiates a new MediaManagementConfigResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaManagementConfigResource() *MediaManagementConfigResource {
	this := MediaManagementConfigResource{}
	return &this
}

// NewMediaManagementConfigResourceWithDefaults instantiates a new MediaManagementConfigResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaManagementConfigResourceWithDefaults() *MediaManagementConfigResource {
	this := MediaManagementConfigResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MediaManagementConfigResource) SetId(v int32) {
	o.Id = &v
}

// GetAutoUnmonitorPreviouslyDownloadedMovies returns the AutoUnmonitorPreviouslyDownloadedMovies field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetAutoUnmonitorPreviouslyDownloadedMovies() bool {
	if o == nil || isNil(o.AutoUnmonitorPreviouslyDownloadedMovies) {
		var ret bool
		return ret
	}
	return *o.AutoUnmonitorPreviouslyDownloadedMovies
}

// GetAutoUnmonitorPreviouslyDownloadedMoviesOk returns a tuple with the AutoUnmonitorPreviouslyDownloadedMovies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetAutoUnmonitorPreviouslyDownloadedMoviesOk() (*bool, bool) {
	if o == nil || isNil(o.AutoUnmonitorPreviouslyDownloadedMovies) {
    return nil, false
	}
	return o.AutoUnmonitorPreviouslyDownloadedMovies, true
}

// HasAutoUnmonitorPreviouslyDownloadedMovies returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasAutoUnmonitorPreviouslyDownloadedMovies() bool {
	if o != nil && !isNil(o.AutoUnmonitorPreviouslyDownloadedMovies) {
		return true
	}

	return false
}

// SetAutoUnmonitorPreviouslyDownloadedMovies gets a reference to the given bool and assigns it to the AutoUnmonitorPreviouslyDownloadedMovies field.
func (o *MediaManagementConfigResource) SetAutoUnmonitorPreviouslyDownloadedMovies(v bool) {
	o.AutoUnmonitorPreviouslyDownloadedMovies = &v
}

// GetRecycleBin returns the RecycleBin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaManagementConfigResource) GetRecycleBin() string {
	if o == nil || isNil(o.RecycleBin.Get()) {
		var ret string
		return ret
	}
	return *o.RecycleBin.Get()
}

// GetRecycleBinOk returns a tuple with the RecycleBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaManagementConfigResource) GetRecycleBinOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RecycleBin.Get(), o.RecycleBin.IsSet()
}

// HasRecycleBin returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasRecycleBin() bool {
	if o != nil && o.RecycleBin.IsSet() {
		return true
	}

	return false
}

// SetRecycleBin gets a reference to the given NullableString and assigns it to the RecycleBin field.
func (o *MediaManagementConfigResource) SetRecycleBin(v string) {
	o.RecycleBin.Set(&v)
}
// SetRecycleBinNil sets the value for RecycleBin to be an explicit nil
func (o *MediaManagementConfigResource) SetRecycleBinNil() {
	o.RecycleBin.Set(nil)
}

// UnsetRecycleBin ensures that no value is present for RecycleBin, not even an explicit nil
func (o *MediaManagementConfigResource) UnsetRecycleBin() {
	o.RecycleBin.Unset()
}

// GetRecycleBinCleanupDays returns the RecycleBinCleanupDays field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetRecycleBinCleanupDays() int32 {
	if o == nil || isNil(o.RecycleBinCleanupDays) {
		var ret int32
		return ret
	}
	return *o.RecycleBinCleanupDays
}

// GetRecycleBinCleanupDaysOk returns a tuple with the RecycleBinCleanupDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetRecycleBinCleanupDaysOk() (*int32, bool) {
	if o == nil || isNil(o.RecycleBinCleanupDays) {
    return nil, false
	}
	return o.RecycleBinCleanupDays, true
}

// HasRecycleBinCleanupDays returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasRecycleBinCleanupDays() bool {
	if o != nil && !isNil(o.RecycleBinCleanupDays) {
		return true
	}

	return false
}

// SetRecycleBinCleanupDays gets a reference to the given int32 and assigns it to the RecycleBinCleanupDays field.
func (o *MediaManagementConfigResource) SetRecycleBinCleanupDays(v int32) {
	o.RecycleBinCleanupDays = &v
}

// GetDownloadPropersAndRepacks returns the DownloadPropersAndRepacks field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetDownloadPropersAndRepacks() ProperDownloadTypes {
	if o == nil || isNil(o.DownloadPropersAndRepacks) {
		var ret ProperDownloadTypes
		return ret
	}
	return *o.DownloadPropersAndRepacks
}

// GetDownloadPropersAndRepacksOk returns a tuple with the DownloadPropersAndRepacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetDownloadPropersAndRepacksOk() (*ProperDownloadTypes, bool) {
	if o == nil || isNil(o.DownloadPropersAndRepacks) {
    return nil, false
	}
	return o.DownloadPropersAndRepacks, true
}

// HasDownloadPropersAndRepacks returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasDownloadPropersAndRepacks() bool {
	if o != nil && !isNil(o.DownloadPropersAndRepacks) {
		return true
	}

	return false
}

// SetDownloadPropersAndRepacks gets a reference to the given ProperDownloadTypes and assigns it to the DownloadPropersAndRepacks field.
func (o *MediaManagementConfigResource) SetDownloadPropersAndRepacks(v ProperDownloadTypes) {
	o.DownloadPropersAndRepacks = &v
}

// GetCreateEmptyMovieFolders returns the CreateEmptyMovieFolders field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetCreateEmptyMovieFolders() bool {
	if o == nil || isNil(o.CreateEmptyMovieFolders) {
		var ret bool
		return ret
	}
	return *o.CreateEmptyMovieFolders
}

// GetCreateEmptyMovieFoldersOk returns a tuple with the CreateEmptyMovieFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetCreateEmptyMovieFoldersOk() (*bool, bool) {
	if o == nil || isNil(o.CreateEmptyMovieFolders) {
    return nil, false
	}
	return o.CreateEmptyMovieFolders, true
}

// HasCreateEmptyMovieFolders returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasCreateEmptyMovieFolders() bool {
	if o != nil && !isNil(o.CreateEmptyMovieFolders) {
		return true
	}

	return false
}

// SetCreateEmptyMovieFolders gets a reference to the given bool and assigns it to the CreateEmptyMovieFolders field.
func (o *MediaManagementConfigResource) SetCreateEmptyMovieFolders(v bool) {
	o.CreateEmptyMovieFolders = &v
}

// GetDeleteEmptyFolders returns the DeleteEmptyFolders field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetDeleteEmptyFolders() bool {
	if o == nil || isNil(o.DeleteEmptyFolders) {
		var ret bool
		return ret
	}
	return *o.DeleteEmptyFolders
}

// GetDeleteEmptyFoldersOk returns a tuple with the DeleteEmptyFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetDeleteEmptyFoldersOk() (*bool, bool) {
	if o == nil || isNil(o.DeleteEmptyFolders) {
    return nil, false
	}
	return o.DeleteEmptyFolders, true
}

// HasDeleteEmptyFolders returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasDeleteEmptyFolders() bool {
	if o != nil && !isNil(o.DeleteEmptyFolders) {
		return true
	}

	return false
}

// SetDeleteEmptyFolders gets a reference to the given bool and assigns it to the DeleteEmptyFolders field.
func (o *MediaManagementConfigResource) SetDeleteEmptyFolders(v bool) {
	o.DeleteEmptyFolders = &v
}

// GetFileDate returns the FileDate field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetFileDate() FileDateType {
	if o == nil || isNil(o.FileDate) {
		var ret FileDateType
		return ret
	}
	return *o.FileDate
}

// GetFileDateOk returns a tuple with the FileDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetFileDateOk() (*FileDateType, bool) {
	if o == nil || isNil(o.FileDate) {
    return nil, false
	}
	return o.FileDate, true
}

// HasFileDate returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasFileDate() bool {
	if o != nil && !isNil(o.FileDate) {
		return true
	}

	return false
}

// SetFileDate gets a reference to the given FileDateType and assigns it to the FileDate field.
func (o *MediaManagementConfigResource) SetFileDate(v FileDateType) {
	o.FileDate = &v
}

// GetRescanAfterRefresh returns the RescanAfterRefresh field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetRescanAfterRefresh() RescanAfterRefreshType {
	if o == nil || isNil(o.RescanAfterRefresh) {
		var ret RescanAfterRefreshType
		return ret
	}
	return *o.RescanAfterRefresh
}

// GetRescanAfterRefreshOk returns a tuple with the RescanAfterRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetRescanAfterRefreshOk() (*RescanAfterRefreshType, bool) {
	if o == nil || isNil(o.RescanAfterRefresh) {
    return nil, false
	}
	return o.RescanAfterRefresh, true
}

// HasRescanAfterRefresh returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasRescanAfterRefresh() bool {
	if o != nil && !isNil(o.RescanAfterRefresh) {
		return true
	}

	return false
}

// SetRescanAfterRefresh gets a reference to the given RescanAfterRefreshType and assigns it to the RescanAfterRefresh field.
func (o *MediaManagementConfigResource) SetRescanAfterRefresh(v RescanAfterRefreshType) {
	o.RescanAfterRefresh = &v
}

// GetAutoRenameFolders returns the AutoRenameFolders field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetAutoRenameFolders() bool {
	if o == nil || isNil(o.AutoRenameFolders) {
		var ret bool
		return ret
	}
	return *o.AutoRenameFolders
}

// GetAutoRenameFoldersOk returns a tuple with the AutoRenameFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetAutoRenameFoldersOk() (*bool, bool) {
	if o == nil || isNil(o.AutoRenameFolders) {
    return nil, false
	}
	return o.AutoRenameFolders, true
}

// HasAutoRenameFolders returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasAutoRenameFolders() bool {
	if o != nil && !isNil(o.AutoRenameFolders) {
		return true
	}

	return false
}

// SetAutoRenameFolders gets a reference to the given bool and assigns it to the AutoRenameFolders field.
func (o *MediaManagementConfigResource) SetAutoRenameFolders(v bool) {
	o.AutoRenameFolders = &v
}

// GetPathsDefaultStatic returns the PathsDefaultStatic field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetPathsDefaultStatic() bool {
	if o == nil || isNil(o.PathsDefaultStatic) {
		var ret bool
		return ret
	}
	return *o.PathsDefaultStatic
}

// GetPathsDefaultStaticOk returns a tuple with the PathsDefaultStatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetPathsDefaultStaticOk() (*bool, bool) {
	if o == nil || isNil(o.PathsDefaultStatic) {
    return nil, false
	}
	return o.PathsDefaultStatic, true
}

// HasPathsDefaultStatic returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasPathsDefaultStatic() bool {
	if o != nil && !isNil(o.PathsDefaultStatic) {
		return true
	}

	return false
}

// SetPathsDefaultStatic gets a reference to the given bool and assigns it to the PathsDefaultStatic field.
func (o *MediaManagementConfigResource) SetPathsDefaultStatic(v bool) {
	o.PathsDefaultStatic = &v
}

// GetSetPermissionsLinux returns the SetPermissionsLinux field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetSetPermissionsLinux() bool {
	if o == nil || isNil(o.SetPermissionsLinux) {
		var ret bool
		return ret
	}
	return *o.SetPermissionsLinux
}

// GetSetPermissionsLinuxOk returns a tuple with the SetPermissionsLinux field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetSetPermissionsLinuxOk() (*bool, bool) {
	if o == nil || isNil(o.SetPermissionsLinux) {
    return nil, false
	}
	return o.SetPermissionsLinux, true
}

// HasSetPermissionsLinux returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasSetPermissionsLinux() bool {
	if o != nil && !isNil(o.SetPermissionsLinux) {
		return true
	}

	return false
}

// SetSetPermissionsLinux gets a reference to the given bool and assigns it to the SetPermissionsLinux field.
func (o *MediaManagementConfigResource) SetSetPermissionsLinux(v bool) {
	o.SetPermissionsLinux = &v
}

// GetChmodFolder returns the ChmodFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaManagementConfigResource) GetChmodFolder() string {
	if o == nil || isNil(o.ChmodFolder.Get()) {
		var ret string
		return ret
	}
	return *o.ChmodFolder.Get()
}

// GetChmodFolderOk returns a tuple with the ChmodFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaManagementConfigResource) GetChmodFolderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChmodFolder.Get(), o.ChmodFolder.IsSet()
}

// HasChmodFolder returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasChmodFolder() bool {
	if o != nil && o.ChmodFolder.IsSet() {
		return true
	}

	return false
}

// SetChmodFolder gets a reference to the given NullableString and assigns it to the ChmodFolder field.
func (o *MediaManagementConfigResource) SetChmodFolder(v string) {
	o.ChmodFolder.Set(&v)
}
// SetChmodFolderNil sets the value for ChmodFolder to be an explicit nil
func (o *MediaManagementConfigResource) SetChmodFolderNil() {
	o.ChmodFolder.Set(nil)
}

// UnsetChmodFolder ensures that no value is present for ChmodFolder, not even an explicit nil
func (o *MediaManagementConfigResource) UnsetChmodFolder() {
	o.ChmodFolder.Unset()
}

// GetChownGroup returns the ChownGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaManagementConfigResource) GetChownGroup() string {
	if o == nil || isNil(o.ChownGroup.Get()) {
		var ret string
		return ret
	}
	return *o.ChownGroup.Get()
}

// GetChownGroupOk returns a tuple with the ChownGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaManagementConfigResource) GetChownGroupOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ChownGroup.Get(), o.ChownGroup.IsSet()
}

// HasChownGroup returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasChownGroup() bool {
	if o != nil && o.ChownGroup.IsSet() {
		return true
	}

	return false
}

// SetChownGroup gets a reference to the given NullableString and assigns it to the ChownGroup field.
func (o *MediaManagementConfigResource) SetChownGroup(v string) {
	o.ChownGroup.Set(&v)
}
// SetChownGroupNil sets the value for ChownGroup to be an explicit nil
func (o *MediaManagementConfigResource) SetChownGroupNil() {
	o.ChownGroup.Set(nil)
}

// UnsetChownGroup ensures that no value is present for ChownGroup, not even an explicit nil
func (o *MediaManagementConfigResource) UnsetChownGroup() {
	o.ChownGroup.Unset()
}

// GetSkipFreeSpaceCheckWhenImporting returns the SkipFreeSpaceCheckWhenImporting field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetSkipFreeSpaceCheckWhenImporting() bool {
	if o == nil || isNil(o.SkipFreeSpaceCheckWhenImporting) {
		var ret bool
		return ret
	}
	return *o.SkipFreeSpaceCheckWhenImporting
}

// GetSkipFreeSpaceCheckWhenImportingOk returns a tuple with the SkipFreeSpaceCheckWhenImporting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetSkipFreeSpaceCheckWhenImportingOk() (*bool, bool) {
	if o == nil || isNil(o.SkipFreeSpaceCheckWhenImporting) {
    return nil, false
	}
	return o.SkipFreeSpaceCheckWhenImporting, true
}

// HasSkipFreeSpaceCheckWhenImporting returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasSkipFreeSpaceCheckWhenImporting() bool {
	if o != nil && !isNil(o.SkipFreeSpaceCheckWhenImporting) {
		return true
	}

	return false
}

// SetSkipFreeSpaceCheckWhenImporting gets a reference to the given bool and assigns it to the SkipFreeSpaceCheckWhenImporting field.
func (o *MediaManagementConfigResource) SetSkipFreeSpaceCheckWhenImporting(v bool) {
	o.SkipFreeSpaceCheckWhenImporting = &v
}

// GetMinimumFreeSpaceWhenImporting returns the MinimumFreeSpaceWhenImporting field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetMinimumFreeSpaceWhenImporting() int32 {
	if o == nil || isNil(o.MinimumFreeSpaceWhenImporting) {
		var ret int32
		return ret
	}
	return *o.MinimumFreeSpaceWhenImporting
}

// GetMinimumFreeSpaceWhenImportingOk returns a tuple with the MinimumFreeSpaceWhenImporting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetMinimumFreeSpaceWhenImportingOk() (*int32, bool) {
	if o == nil || isNil(o.MinimumFreeSpaceWhenImporting) {
    return nil, false
	}
	return o.MinimumFreeSpaceWhenImporting, true
}

// HasMinimumFreeSpaceWhenImporting returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasMinimumFreeSpaceWhenImporting() bool {
	if o != nil && !isNil(o.MinimumFreeSpaceWhenImporting) {
		return true
	}

	return false
}

// SetMinimumFreeSpaceWhenImporting gets a reference to the given int32 and assigns it to the MinimumFreeSpaceWhenImporting field.
func (o *MediaManagementConfigResource) SetMinimumFreeSpaceWhenImporting(v int32) {
	o.MinimumFreeSpaceWhenImporting = &v
}

// GetCopyUsingHardlinks returns the CopyUsingHardlinks field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetCopyUsingHardlinks() bool {
	if o == nil || isNil(o.CopyUsingHardlinks) {
		var ret bool
		return ret
	}
	return *o.CopyUsingHardlinks
}

// GetCopyUsingHardlinksOk returns a tuple with the CopyUsingHardlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetCopyUsingHardlinksOk() (*bool, bool) {
	if o == nil || isNil(o.CopyUsingHardlinks) {
    return nil, false
	}
	return o.CopyUsingHardlinks, true
}

// HasCopyUsingHardlinks returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasCopyUsingHardlinks() bool {
	if o != nil && !isNil(o.CopyUsingHardlinks) {
		return true
	}

	return false
}

// SetCopyUsingHardlinks gets a reference to the given bool and assigns it to the CopyUsingHardlinks field.
func (o *MediaManagementConfigResource) SetCopyUsingHardlinks(v bool) {
	o.CopyUsingHardlinks = &v
}

// GetUseScriptImport returns the UseScriptImport field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetUseScriptImport() bool {
	if o == nil || isNil(o.UseScriptImport) {
		var ret bool
		return ret
	}
	return *o.UseScriptImport
}

// GetUseScriptImportOk returns a tuple with the UseScriptImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetUseScriptImportOk() (*bool, bool) {
	if o == nil || isNil(o.UseScriptImport) {
    return nil, false
	}
	return o.UseScriptImport, true
}

// HasUseScriptImport returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasUseScriptImport() bool {
	if o != nil && !isNil(o.UseScriptImport) {
		return true
	}

	return false
}

// SetUseScriptImport gets a reference to the given bool and assigns it to the UseScriptImport field.
func (o *MediaManagementConfigResource) SetUseScriptImport(v bool) {
	o.UseScriptImport = &v
}

// GetScriptImportPath returns the ScriptImportPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaManagementConfigResource) GetScriptImportPath() string {
	if o == nil || isNil(o.ScriptImportPath.Get()) {
		var ret string
		return ret
	}
	return *o.ScriptImportPath.Get()
}

// GetScriptImportPathOk returns a tuple with the ScriptImportPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaManagementConfigResource) GetScriptImportPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ScriptImportPath.Get(), o.ScriptImportPath.IsSet()
}

// HasScriptImportPath returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasScriptImportPath() bool {
	if o != nil && o.ScriptImportPath.IsSet() {
		return true
	}

	return false
}

// SetScriptImportPath gets a reference to the given NullableString and assigns it to the ScriptImportPath field.
func (o *MediaManagementConfigResource) SetScriptImportPath(v string) {
	o.ScriptImportPath.Set(&v)
}
// SetScriptImportPathNil sets the value for ScriptImportPath to be an explicit nil
func (o *MediaManagementConfigResource) SetScriptImportPathNil() {
	o.ScriptImportPath.Set(nil)
}

// UnsetScriptImportPath ensures that no value is present for ScriptImportPath, not even an explicit nil
func (o *MediaManagementConfigResource) UnsetScriptImportPath() {
	o.ScriptImportPath.Unset()
}

// GetImportExtraFiles returns the ImportExtraFiles field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetImportExtraFiles() bool {
	if o == nil || isNil(o.ImportExtraFiles) {
		var ret bool
		return ret
	}
	return *o.ImportExtraFiles
}

// GetImportExtraFilesOk returns a tuple with the ImportExtraFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetImportExtraFilesOk() (*bool, bool) {
	if o == nil || isNil(o.ImportExtraFiles) {
    return nil, false
	}
	return o.ImportExtraFiles, true
}

// HasImportExtraFiles returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasImportExtraFiles() bool {
	if o != nil && !isNil(o.ImportExtraFiles) {
		return true
	}

	return false
}

// SetImportExtraFiles gets a reference to the given bool and assigns it to the ImportExtraFiles field.
func (o *MediaManagementConfigResource) SetImportExtraFiles(v bool) {
	o.ImportExtraFiles = &v
}

// GetExtraFileExtensions returns the ExtraFileExtensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaManagementConfigResource) GetExtraFileExtensions() string {
	if o == nil || isNil(o.ExtraFileExtensions.Get()) {
		var ret string
		return ret
	}
	return *o.ExtraFileExtensions.Get()
}

// GetExtraFileExtensionsOk returns a tuple with the ExtraFileExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaManagementConfigResource) GetExtraFileExtensionsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ExtraFileExtensions.Get(), o.ExtraFileExtensions.IsSet()
}

// HasExtraFileExtensions returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasExtraFileExtensions() bool {
	if o != nil && o.ExtraFileExtensions.IsSet() {
		return true
	}

	return false
}

// SetExtraFileExtensions gets a reference to the given NullableString and assigns it to the ExtraFileExtensions field.
func (o *MediaManagementConfigResource) SetExtraFileExtensions(v string) {
	o.ExtraFileExtensions.Set(&v)
}
// SetExtraFileExtensionsNil sets the value for ExtraFileExtensions to be an explicit nil
func (o *MediaManagementConfigResource) SetExtraFileExtensionsNil() {
	o.ExtraFileExtensions.Set(nil)
}

// UnsetExtraFileExtensions ensures that no value is present for ExtraFileExtensions, not even an explicit nil
func (o *MediaManagementConfigResource) UnsetExtraFileExtensions() {
	o.ExtraFileExtensions.Unset()
}

// GetEnableMediaInfo returns the EnableMediaInfo field value if set, zero value otherwise.
func (o *MediaManagementConfigResource) GetEnableMediaInfo() bool {
	if o == nil || isNil(o.EnableMediaInfo) {
		var ret bool
		return ret
	}
	return *o.EnableMediaInfo
}

// GetEnableMediaInfoOk returns a tuple with the EnableMediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaManagementConfigResource) GetEnableMediaInfoOk() (*bool, bool) {
	if o == nil || isNil(o.EnableMediaInfo) {
    return nil, false
	}
	return o.EnableMediaInfo, true
}

// HasEnableMediaInfo returns a boolean if a field has been set.
func (o *MediaManagementConfigResource) HasEnableMediaInfo() bool {
	if o != nil && !isNil(o.EnableMediaInfo) {
		return true
	}

	return false
}

// SetEnableMediaInfo gets a reference to the given bool and assigns it to the EnableMediaInfo field.
func (o *MediaManagementConfigResource) SetEnableMediaInfo(v bool) {
	o.EnableMediaInfo = &v
}

func (o MediaManagementConfigResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.AutoUnmonitorPreviouslyDownloadedMovies) {
		toSerialize["autoUnmonitorPreviouslyDownloadedMovies"] = o.AutoUnmonitorPreviouslyDownloadedMovies
	}
	if o.RecycleBin.IsSet() {
		toSerialize["recycleBin"] = o.RecycleBin.Get()
	}
	if !isNil(o.RecycleBinCleanupDays) {
		toSerialize["recycleBinCleanupDays"] = o.RecycleBinCleanupDays
	}
	if !isNil(o.DownloadPropersAndRepacks) {
		toSerialize["downloadPropersAndRepacks"] = o.DownloadPropersAndRepacks
	}
	if !isNil(o.CreateEmptyMovieFolders) {
		toSerialize["createEmptyMovieFolders"] = o.CreateEmptyMovieFolders
	}
	if !isNil(o.DeleteEmptyFolders) {
		toSerialize["deleteEmptyFolders"] = o.DeleteEmptyFolders
	}
	if !isNil(o.FileDate) {
		toSerialize["fileDate"] = o.FileDate
	}
	if !isNil(o.RescanAfterRefresh) {
		toSerialize["rescanAfterRefresh"] = o.RescanAfterRefresh
	}
	if !isNil(o.AutoRenameFolders) {
		toSerialize["autoRenameFolders"] = o.AutoRenameFolders
	}
	if !isNil(o.PathsDefaultStatic) {
		toSerialize["pathsDefaultStatic"] = o.PathsDefaultStatic
	}
	if !isNil(o.SetPermissionsLinux) {
		toSerialize["setPermissionsLinux"] = o.SetPermissionsLinux
	}
	if o.ChmodFolder.IsSet() {
		toSerialize["chmodFolder"] = o.ChmodFolder.Get()
	}
	if o.ChownGroup.IsSet() {
		toSerialize["chownGroup"] = o.ChownGroup.Get()
	}
	if !isNil(o.SkipFreeSpaceCheckWhenImporting) {
		toSerialize["skipFreeSpaceCheckWhenImporting"] = o.SkipFreeSpaceCheckWhenImporting
	}
	if !isNil(o.MinimumFreeSpaceWhenImporting) {
		toSerialize["minimumFreeSpaceWhenImporting"] = o.MinimumFreeSpaceWhenImporting
	}
	if !isNil(o.CopyUsingHardlinks) {
		toSerialize["copyUsingHardlinks"] = o.CopyUsingHardlinks
	}
	if !isNil(o.UseScriptImport) {
		toSerialize["useScriptImport"] = o.UseScriptImport
	}
	if o.ScriptImportPath.IsSet() {
		toSerialize["scriptImportPath"] = o.ScriptImportPath.Get()
	}
	if !isNil(o.ImportExtraFiles) {
		toSerialize["importExtraFiles"] = o.ImportExtraFiles
	}
	if o.ExtraFileExtensions.IsSet() {
		toSerialize["extraFileExtensions"] = o.ExtraFileExtensions.Get()
	}
	if !isNil(o.EnableMediaInfo) {
		toSerialize["enableMediaInfo"] = o.EnableMediaInfo
	}
	return json.Marshal(toSerialize)
}

type NullableMediaManagementConfigResource struct {
	value *MediaManagementConfigResource
	isSet bool
}

func (v NullableMediaManagementConfigResource) Get() *MediaManagementConfigResource {
	return v.value
}

func (v *NullableMediaManagementConfigResource) Set(val *MediaManagementConfigResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaManagementConfigResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaManagementConfigResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaManagementConfigResource(val *MediaManagementConfigResource) *NullableMediaManagementConfigResource {
	return &NullableMediaManagementConfigResource{value: val, isSet: true}
}

func (v NullableMediaManagementConfigResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaManagementConfigResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



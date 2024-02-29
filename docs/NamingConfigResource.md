# NamingConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RenameMovies** | Pointer to **bool** |  | [optional] 
**ReplaceIllegalCharacters** | Pointer to **bool** |  | [optional] 
**ColonReplacementFormat** | Pointer to [**ColonReplacementFormat**](ColonReplacementFormat.md) |  | [optional] 
**StandardMovieFormat** | Pointer to **NullableString** |  | [optional] 
**MovieFolderFormat** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNamingConfigResource

`func NewNamingConfigResource() *NamingConfigResource`

NewNamingConfigResource instantiates a new NamingConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingConfigResourceWithDefaults

`func NewNamingConfigResourceWithDefaults() *NamingConfigResource`

NewNamingConfigResourceWithDefaults instantiates a new NamingConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NamingConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamingConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamingConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NamingConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRenameMovies

`func (o *NamingConfigResource) GetRenameMovies() bool`

GetRenameMovies returns the RenameMovies field if non-nil, zero value otherwise.

### GetRenameMoviesOk

`func (o *NamingConfigResource) GetRenameMoviesOk() (*bool, bool)`

GetRenameMoviesOk returns a tuple with the RenameMovies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameMovies

`func (o *NamingConfigResource) SetRenameMovies(v bool)`

SetRenameMovies sets RenameMovies field to given value.

### HasRenameMovies

`func (o *NamingConfigResource) HasRenameMovies() bool`

HasRenameMovies returns a boolean if a field has been set.

### GetReplaceIllegalCharacters

`func (o *NamingConfigResource) GetReplaceIllegalCharacters() bool`

GetReplaceIllegalCharacters returns the ReplaceIllegalCharacters field if non-nil, zero value otherwise.

### GetReplaceIllegalCharactersOk

`func (o *NamingConfigResource) GetReplaceIllegalCharactersOk() (*bool, bool)`

GetReplaceIllegalCharactersOk returns a tuple with the ReplaceIllegalCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceIllegalCharacters

`func (o *NamingConfigResource) SetReplaceIllegalCharacters(v bool)`

SetReplaceIllegalCharacters sets ReplaceIllegalCharacters field to given value.

### HasReplaceIllegalCharacters

`func (o *NamingConfigResource) HasReplaceIllegalCharacters() bool`

HasReplaceIllegalCharacters returns a boolean if a field has been set.

### GetColonReplacementFormat

`func (o *NamingConfigResource) GetColonReplacementFormat() ColonReplacementFormat`

GetColonReplacementFormat returns the ColonReplacementFormat field if non-nil, zero value otherwise.

### GetColonReplacementFormatOk

`func (o *NamingConfigResource) GetColonReplacementFormatOk() (*ColonReplacementFormat, bool)`

GetColonReplacementFormatOk returns a tuple with the ColonReplacementFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColonReplacementFormat

`func (o *NamingConfigResource) SetColonReplacementFormat(v ColonReplacementFormat)`

SetColonReplacementFormat sets ColonReplacementFormat field to given value.

### HasColonReplacementFormat

`func (o *NamingConfigResource) HasColonReplacementFormat() bool`

HasColonReplacementFormat returns a boolean if a field has been set.

### GetStandardMovieFormat

`func (o *NamingConfigResource) GetStandardMovieFormat() string`

GetStandardMovieFormat returns the StandardMovieFormat field if non-nil, zero value otherwise.

### GetStandardMovieFormatOk

`func (o *NamingConfigResource) GetStandardMovieFormatOk() (*string, bool)`

GetStandardMovieFormatOk returns a tuple with the StandardMovieFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardMovieFormat

`func (o *NamingConfigResource) SetStandardMovieFormat(v string)`

SetStandardMovieFormat sets StandardMovieFormat field to given value.

### HasStandardMovieFormat

`func (o *NamingConfigResource) HasStandardMovieFormat() bool`

HasStandardMovieFormat returns a boolean if a field has been set.

### SetStandardMovieFormatNil

`func (o *NamingConfigResource) SetStandardMovieFormatNil(b bool)`

 SetStandardMovieFormatNil sets the value for StandardMovieFormat to be an explicit nil

### UnsetStandardMovieFormat
`func (o *NamingConfigResource) UnsetStandardMovieFormat()`

UnsetStandardMovieFormat ensures that no value is present for StandardMovieFormat, not even an explicit nil
### GetMovieFolderFormat

`func (o *NamingConfigResource) GetMovieFolderFormat() string`

GetMovieFolderFormat returns the MovieFolderFormat field if non-nil, zero value otherwise.

### GetMovieFolderFormatOk

`func (o *NamingConfigResource) GetMovieFolderFormatOk() (*string, bool)`

GetMovieFolderFormatOk returns a tuple with the MovieFolderFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieFolderFormat

`func (o *NamingConfigResource) SetMovieFolderFormat(v string)`

SetMovieFolderFormat sets MovieFolderFormat field to given value.

### HasMovieFolderFormat

`func (o *NamingConfigResource) HasMovieFolderFormat() bool`

HasMovieFolderFormat returns a boolean if a field has been set.

### SetMovieFolderFormatNil

`func (o *NamingConfigResource) SetMovieFolderFormatNil(b bool)`

 SetMovieFolderFormatNil sets the value for MovieFolderFormat to be an explicit nil

### UnsetMovieFolderFormat
`func (o *NamingConfigResource) UnsetMovieFolderFormat()`

UnsetMovieFolderFormat ensures that no value is present for MovieFolderFormat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



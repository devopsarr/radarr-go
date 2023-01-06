# ParsedMovieInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MovieTitles** | Pointer to **[]string** |  | [optional] 
**OriginalTitle** | Pointer to **NullableString** |  | [optional] 
**ReleaseTitle** | Pointer to **NullableString** |  | [optional] 
**SimpleReleaseTitle** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to [**QualityModel**](QualityModel.md) |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**ReleaseGroup** | Pointer to **NullableString** |  | [optional] 
**ReleaseHash** | Pointer to **NullableString** |  | [optional] 
**Edition** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**ImdbId** | Pointer to **NullableString** |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 
**ExtraInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**MovieTitle** | Pointer to **NullableString** |  | [optional] [readonly] 
**PrimaryMovieTitle** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewParsedMovieInfo

`func NewParsedMovieInfo() *ParsedMovieInfo`

NewParsedMovieInfo instantiates a new ParsedMovieInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsedMovieInfoWithDefaults

`func NewParsedMovieInfoWithDefaults() *ParsedMovieInfo`

NewParsedMovieInfoWithDefaults instantiates a new ParsedMovieInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovieTitles

`func (o *ParsedMovieInfo) GetMovieTitles() []string`

GetMovieTitles returns the MovieTitles field if non-nil, zero value otherwise.

### GetMovieTitlesOk

`func (o *ParsedMovieInfo) GetMovieTitlesOk() (*[]string, bool)`

GetMovieTitlesOk returns a tuple with the MovieTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieTitles

`func (o *ParsedMovieInfo) SetMovieTitles(v []string)`

SetMovieTitles sets MovieTitles field to given value.

### HasMovieTitles

`func (o *ParsedMovieInfo) HasMovieTitles() bool`

HasMovieTitles returns a boolean if a field has been set.

### SetMovieTitlesNil

`func (o *ParsedMovieInfo) SetMovieTitlesNil(b bool)`

 SetMovieTitlesNil sets the value for MovieTitles to be an explicit nil

### UnsetMovieTitles
`func (o *ParsedMovieInfo) UnsetMovieTitles()`

UnsetMovieTitles ensures that no value is present for MovieTitles, not even an explicit nil
### GetOriginalTitle

`func (o *ParsedMovieInfo) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *ParsedMovieInfo) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *ParsedMovieInfo) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *ParsedMovieInfo) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *ParsedMovieInfo) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *ParsedMovieInfo) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetReleaseTitle

`func (o *ParsedMovieInfo) GetReleaseTitle() string`

GetReleaseTitle returns the ReleaseTitle field if non-nil, zero value otherwise.

### GetReleaseTitleOk

`func (o *ParsedMovieInfo) GetReleaseTitleOk() (*string, bool)`

GetReleaseTitleOk returns a tuple with the ReleaseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTitle

`func (o *ParsedMovieInfo) SetReleaseTitle(v string)`

SetReleaseTitle sets ReleaseTitle field to given value.

### HasReleaseTitle

`func (o *ParsedMovieInfo) HasReleaseTitle() bool`

HasReleaseTitle returns a boolean if a field has been set.

### SetReleaseTitleNil

`func (o *ParsedMovieInfo) SetReleaseTitleNil(b bool)`

 SetReleaseTitleNil sets the value for ReleaseTitle to be an explicit nil

### UnsetReleaseTitle
`func (o *ParsedMovieInfo) UnsetReleaseTitle()`

UnsetReleaseTitle ensures that no value is present for ReleaseTitle, not even an explicit nil
### GetSimpleReleaseTitle

`func (o *ParsedMovieInfo) GetSimpleReleaseTitle() string`

GetSimpleReleaseTitle returns the SimpleReleaseTitle field if non-nil, zero value otherwise.

### GetSimpleReleaseTitleOk

`func (o *ParsedMovieInfo) GetSimpleReleaseTitleOk() (*string, bool)`

GetSimpleReleaseTitleOk returns a tuple with the SimpleReleaseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleReleaseTitle

`func (o *ParsedMovieInfo) SetSimpleReleaseTitle(v string)`

SetSimpleReleaseTitle sets SimpleReleaseTitle field to given value.

### HasSimpleReleaseTitle

`func (o *ParsedMovieInfo) HasSimpleReleaseTitle() bool`

HasSimpleReleaseTitle returns a boolean if a field has been set.

### SetSimpleReleaseTitleNil

`func (o *ParsedMovieInfo) SetSimpleReleaseTitleNil(b bool)`

 SetSimpleReleaseTitleNil sets the value for SimpleReleaseTitle to be an explicit nil

### UnsetSimpleReleaseTitle
`func (o *ParsedMovieInfo) UnsetSimpleReleaseTitle()`

UnsetSimpleReleaseTitle ensures that no value is present for SimpleReleaseTitle, not even an explicit nil
### GetQuality

`func (o *ParsedMovieInfo) GetQuality() QualityModel`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ParsedMovieInfo) GetQualityOk() (*QualityModel, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ParsedMovieInfo) SetQuality(v QualityModel)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ParsedMovieInfo) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetLanguages

`func (o *ParsedMovieInfo) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ParsedMovieInfo) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ParsedMovieInfo) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ParsedMovieInfo) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *ParsedMovieInfo) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *ParsedMovieInfo) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetReleaseGroup

`func (o *ParsedMovieInfo) GetReleaseGroup() string`

GetReleaseGroup returns the ReleaseGroup field if non-nil, zero value otherwise.

### GetReleaseGroupOk

`func (o *ParsedMovieInfo) GetReleaseGroupOk() (*string, bool)`

GetReleaseGroupOk returns a tuple with the ReleaseGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroup

`func (o *ParsedMovieInfo) SetReleaseGroup(v string)`

SetReleaseGroup sets ReleaseGroup field to given value.

### HasReleaseGroup

`func (o *ParsedMovieInfo) HasReleaseGroup() bool`

HasReleaseGroup returns a boolean if a field has been set.

### SetReleaseGroupNil

`func (o *ParsedMovieInfo) SetReleaseGroupNil(b bool)`

 SetReleaseGroupNil sets the value for ReleaseGroup to be an explicit nil

### UnsetReleaseGroup
`func (o *ParsedMovieInfo) UnsetReleaseGroup()`

UnsetReleaseGroup ensures that no value is present for ReleaseGroup, not even an explicit nil
### GetReleaseHash

`func (o *ParsedMovieInfo) GetReleaseHash() string`

GetReleaseHash returns the ReleaseHash field if non-nil, zero value otherwise.

### GetReleaseHashOk

`func (o *ParsedMovieInfo) GetReleaseHashOk() (*string, bool)`

GetReleaseHashOk returns a tuple with the ReleaseHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseHash

`func (o *ParsedMovieInfo) SetReleaseHash(v string)`

SetReleaseHash sets ReleaseHash field to given value.

### HasReleaseHash

`func (o *ParsedMovieInfo) HasReleaseHash() bool`

HasReleaseHash returns a boolean if a field has been set.

### SetReleaseHashNil

`func (o *ParsedMovieInfo) SetReleaseHashNil(b bool)`

 SetReleaseHashNil sets the value for ReleaseHash to be an explicit nil

### UnsetReleaseHash
`func (o *ParsedMovieInfo) UnsetReleaseHash()`

UnsetReleaseHash ensures that no value is present for ReleaseHash, not even an explicit nil
### GetEdition

`func (o *ParsedMovieInfo) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *ParsedMovieInfo) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *ParsedMovieInfo) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *ParsedMovieInfo) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### SetEditionNil

`func (o *ParsedMovieInfo) SetEditionNil(b bool)`

 SetEditionNil sets the value for Edition to be an explicit nil

### UnsetEdition
`func (o *ParsedMovieInfo) UnsetEdition()`

UnsetEdition ensures that no value is present for Edition, not even an explicit nil
### GetYear

`func (o *ParsedMovieInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ParsedMovieInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ParsedMovieInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ParsedMovieInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetImdbId

`func (o *ParsedMovieInfo) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *ParsedMovieInfo) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *ParsedMovieInfo) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *ParsedMovieInfo) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### SetImdbIdNil

`func (o *ParsedMovieInfo) SetImdbIdNil(b bool)`

 SetImdbIdNil sets the value for ImdbId to be an explicit nil

### UnsetImdbId
`func (o *ParsedMovieInfo) UnsetImdbId()`

UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
### GetTmdbId

`func (o *ParsedMovieInfo) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *ParsedMovieInfo) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *ParsedMovieInfo) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *ParsedMovieInfo) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetExtraInfo

`func (o *ParsedMovieInfo) GetExtraInfo() map[string]interface{}`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *ParsedMovieInfo) GetExtraInfoOk() (*map[string]interface{}, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *ParsedMovieInfo) SetExtraInfo(v map[string]interface{})`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *ParsedMovieInfo) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### SetExtraInfoNil

`func (o *ParsedMovieInfo) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *ParsedMovieInfo) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil
### GetMovieTitle

`func (o *ParsedMovieInfo) GetMovieTitle() string`

GetMovieTitle returns the MovieTitle field if non-nil, zero value otherwise.

### GetMovieTitleOk

`func (o *ParsedMovieInfo) GetMovieTitleOk() (*string, bool)`

GetMovieTitleOk returns a tuple with the MovieTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieTitle

`func (o *ParsedMovieInfo) SetMovieTitle(v string)`

SetMovieTitle sets MovieTitle field to given value.

### HasMovieTitle

`func (o *ParsedMovieInfo) HasMovieTitle() bool`

HasMovieTitle returns a boolean if a field has been set.

### SetMovieTitleNil

`func (o *ParsedMovieInfo) SetMovieTitleNil(b bool)`

 SetMovieTitleNil sets the value for MovieTitle to be an explicit nil

### UnsetMovieTitle
`func (o *ParsedMovieInfo) UnsetMovieTitle()`

UnsetMovieTitle ensures that no value is present for MovieTitle, not even an explicit nil
### GetPrimaryMovieTitle

`func (o *ParsedMovieInfo) GetPrimaryMovieTitle() string`

GetPrimaryMovieTitle returns the PrimaryMovieTitle field if non-nil, zero value otherwise.

### GetPrimaryMovieTitleOk

`func (o *ParsedMovieInfo) GetPrimaryMovieTitleOk() (*string, bool)`

GetPrimaryMovieTitleOk returns a tuple with the PrimaryMovieTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMovieTitle

`func (o *ParsedMovieInfo) SetPrimaryMovieTitle(v string)`

SetPrimaryMovieTitle sets PrimaryMovieTitle field to given value.

### HasPrimaryMovieTitle

`func (o *ParsedMovieInfo) HasPrimaryMovieTitle() bool`

HasPrimaryMovieTitle returns a boolean if a field has been set.

### SetPrimaryMovieTitleNil

`func (o *ParsedMovieInfo) SetPrimaryMovieTitleNil(b bool)`

 SetPrimaryMovieTitleNil sets the value for PrimaryMovieTitle to be an explicit nil

### UnsetPrimaryMovieTitle
`func (o *ParsedMovieInfo) UnsetPrimaryMovieTitle()`

UnsetPrimaryMovieTitle ensures that no value is present for PrimaryMovieTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



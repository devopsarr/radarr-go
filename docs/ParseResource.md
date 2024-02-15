# ParseResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ParsedMovieInfo** | Pointer to [**ParsedMovieInfo**](ParsedMovieInfo.md) |  | [optional] 
**Movie** | Pointer to [**MovieResource**](MovieResource.md) |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 
**CustomFormats** | Pointer to [**[]CustomFormatResource**](CustomFormatResource.md) |  | [optional] 
**CustomFormatScore** | Pointer to **int32** |  | [optional] 

## Methods

### NewParseResource

`func NewParseResource() *ParseResource`

NewParseResource instantiates a new ParseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseResourceWithDefaults

`func NewParseResourceWithDefaults() *ParseResource`

NewParseResourceWithDefaults instantiates a new ParseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParseResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParseResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParseResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ParseResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ParseResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ParseResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ParseResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ParseResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ParseResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ParseResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetParsedMovieInfo

`func (o *ParseResource) GetParsedMovieInfo() ParsedMovieInfo`

GetParsedMovieInfo returns the ParsedMovieInfo field if non-nil, zero value otherwise.

### GetParsedMovieInfoOk

`func (o *ParseResource) GetParsedMovieInfoOk() (*ParsedMovieInfo, bool)`

GetParsedMovieInfoOk returns a tuple with the ParsedMovieInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedMovieInfo

`func (o *ParseResource) SetParsedMovieInfo(v ParsedMovieInfo)`

SetParsedMovieInfo sets ParsedMovieInfo field to given value.

### HasParsedMovieInfo

`func (o *ParseResource) HasParsedMovieInfo() bool`

HasParsedMovieInfo returns a boolean if a field has been set.

### GetMovie

`func (o *ParseResource) GetMovie() MovieResource`

GetMovie returns the Movie field if non-nil, zero value otherwise.

### GetMovieOk

`func (o *ParseResource) GetMovieOk() (*MovieResource, bool)`

GetMovieOk returns a tuple with the Movie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovie

`func (o *ParseResource) SetMovie(v MovieResource)`

SetMovie sets Movie field to given value.

### HasMovie

`func (o *ParseResource) HasMovie() bool`

HasMovie returns a boolean if a field has been set.

### GetLanguages

`func (o *ParseResource) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ParseResource) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ParseResource) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ParseResource) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *ParseResource) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *ParseResource) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetCustomFormats

`func (o *ParseResource) GetCustomFormats() []CustomFormatResource`

GetCustomFormats returns the CustomFormats field if non-nil, zero value otherwise.

### GetCustomFormatsOk

`func (o *ParseResource) GetCustomFormatsOk() (*[]CustomFormatResource, bool)`

GetCustomFormatsOk returns a tuple with the CustomFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormats

`func (o *ParseResource) SetCustomFormats(v []CustomFormatResource)`

SetCustomFormats sets CustomFormats field to given value.

### HasCustomFormats

`func (o *ParseResource) HasCustomFormats() bool`

HasCustomFormats returns a boolean if a field has been set.

### SetCustomFormatsNil

`func (o *ParseResource) SetCustomFormatsNil(b bool)`

 SetCustomFormatsNil sets the value for CustomFormats to be an explicit nil

### UnsetCustomFormats
`func (o *ParseResource) UnsetCustomFormats()`

UnsetCustomFormats ensures that no value is present for CustomFormats, not even an explicit nil
### GetCustomFormatScore

`func (o *ParseResource) GetCustomFormatScore() int32`

GetCustomFormatScore returns the CustomFormatScore field if non-nil, zero value otherwise.

### GetCustomFormatScoreOk

`func (o *ParseResource) GetCustomFormatScoreOk() (*int32, bool)`

GetCustomFormatScoreOk returns a tuple with the CustomFormatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFormatScore

`func (o *ParseResource) SetCustomFormatScore(v int32)`

SetCustomFormatScore sets CustomFormatScore field to given value.

### HasCustomFormatScore

`func (o *ParseResource) HasCustomFormatScore() bool`

HasCustomFormatScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AlternativeTitle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SourceType** | Pointer to [**SourceType**](SourceType.md) |  | [optional] 
**MovieMetadataId** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**SourceId** | Pointer to **int32** |  | [optional] 
**Votes** | Pointer to **int32** |  | [optional] 
**VoteCount** | Pointer to **int32** |  | [optional] 
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 

## Methods

### NewAlternativeTitle

`func NewAlternativeTitle() *AlternativeTitle`

NewAlternativeTitle instantiates a new AlternativeTitle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeTitleWithDefaults

`func NewAlternativeTitleWithDefaults() *AlternativeTitle`

NewAlternativeTitleWithDefaults instantiates a new AlternativeTitle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlternativeTitle) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlternativeTitle) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlternativeTitle) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlternativeTitle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSourceType

`func (o *AlternativeTitle) GetSourceType() SourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AlternativeTitle) GetSourceTypeOk() (*SourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AlternativeTitle) SetSourceType(v SourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AlternativeTitle) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetMovieMetadataId

`func (o *AlternativeTitle) GetMovieMetadataId() int32`

GetMovieMetadataId returns the MovieMetadataId field if non-nil, zero value otherwise.

### GetMovieMetadataIdOk

`func (o *AlternativeTitle) GetMovieMetadataIdOk() (*int32, bool)`

GetMovieMetadataIdOk returns a tuple with the MovieMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieMetadataId

`func (o *AlternativeTitle) SetMovieMetadataId(v int32)`

SetMovieMetadataId sets MovieMetadataId field to given value.

### HasMovieMetadataId

`func (o *AlternativeTitle) HasMovieMetadataId() bool`

HasMovieMetadataId returns a boolean if a field has been set.

### GetTitle

`func (o *AlternativeTitle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlternativeTitle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlternativeTitle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlternativeTitle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AlternativeTitle) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AlternativeTitle) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCleanTitle

`func (o *AlternativeTitle) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *AlternativeTitle) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *AlternativeTitle) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *AlternativeTitle) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *AlternativeTitle) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *AlternativeTitle) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetSourceId

`func (o *AlternativeTitle) GetSourceId() int32`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *AlternativeTitle) GetSourceIdOk() (*int32, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *AlternativeTitle) SetSourceId(v int32)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *AlternativeTitle) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetVotes

`func (o *AlternativeTitle) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *AlternativeTitle) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *AlternativeTitle) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *AlternativeTitle) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetVoteCount

`func (o *AlternativeTitle) GetVoteCount() int32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *AlternativeTitle) GetVoteCountOk() (*int32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *AlternativeTitle) SetVoteCount(v int32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *AlternativeTitle) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.

### GetLanguage

`func (o *AlternativeTitle) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AlternativeTitle) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AlternativeTitle) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AlternativeTitle) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ImportExclusionsResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**ImplementationName** | Pointer to **NullableString** |  | [optional] 
**Implementation** | Pointer to **NullableString** |  | [optional] 
**ConfigContract** | Pointer to **NullableString** |  | [optional] 
**InfoLink** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to [**ProviderMessage**](ProviderMessage.md) |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Presets** | Pointer to [**[]ImportExclusionsResource**](ImportExclusionsResource.md) |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 
**MovieTitle** | Pointer to **NullableString** |  | [optional] 
**MovieYear** | Pointer to **int32** |  | [optional] 

## Methods

### NewImportExclusionsResource

`func NewImportExclusionsResource() *ImportExclusionsResource`

NewImportExclusionsResource instantiates a new ImportExclusionsResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportExclusionsResourceWithDefaults

`func NewImportExclusionsResourceWithDefaults() *ImportExclusionsResource`

NewImportExclusionsResourceWithDefaults instantiates a new ImportExclusionsResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportExclusionsResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportExclusionsResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportExclusionsResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ImportExclusionsResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ImportExclusionsResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportExclusionsResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportExclusionsResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportExclusionsResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImportExclusionsResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImportExclusionsResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFields

`func (o *ImportExclusionsResource) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ImportExclusionsResource) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ImportExclusionsResource) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ImportExclusionsResource) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *ImportExclusionsResource) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *ImportExclusionsResource) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetImplementationName

`func (o *ImportExclusionsResource) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *ImportExclusionsResource) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *ImportExclusionsResource) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *ImportExclusionsResource) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *ImportExclusionsResource) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *ImportExclusionsResource) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetImplementation

`func (o *ImportExclusionsResource) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *ImportExclusionsResource) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *ImportExclusionsResource) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *ImportExclusionsResource) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *ImportExclusionsResource) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *ImportExclusionsResource) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetConfigContract

`func (o *ImportExclusionsResource) GetConfigContract() string`

GetConfigContract returns the ConfigContract field if non-nil, zero value otherwise.

### GetConfigContractOk

`func (o *ImportExclusionsResource) GetConfigContractOk() (*string, bool)`

GetConfigContractOk returns a tuple with the ConfigContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContract

`func (o *ImportExclusionsResource) SetConfigContract(v string)`

SetConfigContract sets ConfigContract field to given value.

### HasConfigContract

`func (o *ImportExclusionsResource) HasConfigContract() bool`

HasConfigContract returns a boolean if a field has been set.

### SetConfigContractNil

`func (o *ImportExclusionsResource) SetConfigContractNil(b bool)`

 SetConfigContractNil sets the value for ConfigContract to be an explicit nil

### UnsetConfigContract
`func (o *ImportExclusionsResource) UnsetConfigContract()`

UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
### GetInfoLink

`func (o *ImportExclusionsResource) GetInfoLink() string`

GetInfoLink returns the InfoLink field if non-nil, zero value otherwise.

### GetInfoLinkOk

`func (o *ImportExclusionsResource) GetInfoLinkOk() (*string, bool)`

GetInfoLinkOk returns a tuple with the InfoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLink

`func (o *ImportExclusionsResource) SetInfoLink(v string)`

SetInfoLink sets InfoLink field to given value.

### HasInfoLink

`func (o *ImportExclusionsResource) HasInfoLink() bool`

HasInfoLink returns a boolean if a field has been set.

### SetInfoLinkNil

`func (o *ImportExclusionsResource) SetInfoLinkNil(b bool)`

 SetInfoLinkNil sets the value for InfoLink to be an explicit nil

### UnsetInfoLink
`func (o *ImportExclusionsResource) UnsetInfoLink()`

UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
### GetMessage

`func (o *ImportExclusionsResource) GetMessage() ProviderMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImportExclusionsResource) GetMessageOk() (*ProviderMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImportExclusionsResource) SetMessage(v ProviderMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ImportExclusionsResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *ImportExclusionsResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportExclusionsResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportExclusionsResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportExclusionsResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ImportExclusionsResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ImportExclusionsResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPresets

`func (o *ImportExclusionsResource) GetPresets() []ImportExclusionsResource`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *ImportExclusionsResource) GetPresetsOk() (*[]ImportExclusionsResource, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *ImportExclusionsResource) SetPresets(v []ImportExclusionsResource)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *ImportExclusionsResource) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresetsNil

`func (o *ImportExclusionsResource) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *ImportExclusionsResource) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil
### GetTmdbId

`func (o *ImportExclusionsResource) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *ImportExclusionsResource) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *ImportExclusionsResource) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *ImportExclusionsResource) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetMovieTitle

`func (o *ImportExclusionsResource) GetMovieTitle() string`

GetMovieTitle returns the MovieTitle field if non-nil, zero value otherwise.

### GetMovieTitleOk

`func (o *ImportExclusionsResource) GetMovieTitleOk() (*string, bool)`

GetMovieTitleOk returns a tuple with the MovieTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieTitle

`func (o *ImportExclusionsResource) SetMovieTitle(v string)`

SetMovieTitle sets MovieTitle field to given value.

### HasMovieTitle

`func (o *ImportExclusionsResource) HasMovieTitle() bool`

HasMovieTitle returns a boolean if a field has been set.

### SetMovieTitleNil

`func (o *ImportExclusionsResource) SetMovieTitleNil(b bool)`

 SetMovieTitleNil sets the value for MovieTitle to be an explicit nil

### UnsetMovieTitle
`func (o *ImportExclusionsResource) UnsetMovieTitle()`

UnsetMovieTitle ensures that no value is present for MovieTitle, not even an explicit nil
### GetMovieYear

`func (o *ImportExclusionsResource) GetMovieYear() int32`

GetMovieYear returns the MovieYear field if non-nil, zero value otherwise.

### GetMovieYearOk

`func (o *ImportExclusionsResource) GetMovieYearOk() (*int32, bool)`

GetMovieYearOk returns a tuple with the MovieYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieYear

`func (o *ImportExclusionsResource) SetMovieYear(v int32)`

SetMovieYear sets MovieYear field to given value.

### HasMovieYear

`func (o *ImportExclusionsResource) HasMovieYear() bool`

HasMovieYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



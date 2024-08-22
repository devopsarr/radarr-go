# ImportListExclusionResource

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
**Presets** | Pointer to [**[]ImportListExclusionResource**](ImportListExclusionResource.md) |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 
**MovieTitle** | Pointer to **NullableString** |  | [optional] 
**MovieYear** | Pointer to **int32** |  | [optional] 

## Methods

### NewImportListExclusionResource

`func NewImportListExclusionResource() *ImportListExclusionResource`

NewImportListExclusionResource instantiates a new ImportListExclusionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportListExclusionResourceWithDefaults

`func NewImportListExclusionResourceWithDefaults() *ImportListExclusionResource`

NewImportListExclusionResourceWithDefaults instantiates a new ImportListExclusionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportListExclusionResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportListExclusionResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportListExclusionResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ImportListExclusionResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ImportListExclusionResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportListExclusionResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportListExclusionResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportListExclusionResource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ImportListExclusionResource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ImportListExclusionResource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFields

`func (o *ImportListExclusionResource) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ImportListExclusionResource) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ImportListExclusionResource) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ImportListExclusionResource) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *ImportListExclusionResource) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *ImportListExclusionResource) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetImplementationName

`func (o *ImportListExclusionResource) GetImplementationName() string`

GetImplementationName returns the ImplementationName field if non-nil, zero value otherwise.

### GetImplementationNameOk

`func (o *ImportListExclusionResource) GetImplementationNameOk() (*string, bool)`

GetImplementationNameOk returns a tuple with the ImplementationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationName

`func (o *ImportListExclusionResource) SetImplementationName(v string)`

SetImplementationName sets ImplementationName field to given value.

### HasImplementationName

`func (o *ImportListExclusionResource) HasImplementationName() bool`

HasImplementationName returns a boolean if a field has been set.

### SetImplementationNameNil

`func (o *ImportListExclusionResource) SetImplementationNameNil(b bool)`

 SetImplementationNameNil sets the value for ImplementationName to be an explicit nil

### UnsetImplementationName
`func (o *ImportListExclusionResource) UnsetImplementationName()`

UnsetImplementationName ensures that no value is present for ImplementationName, not even an explicit nil
### GetImplementation

`func (o *ImportListExclusionResource) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *ImportListExclusionResource) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *ImportListExclusionResource) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.

### HasImplementation

`func (o *ImportListExclusionResource) HasImplementation() bool`

HasImplementation returns a boolean if a field has been set.

### SetImplementationNil

`func (o *ImportListExclusionResource) SetImplementationNil(b bool)`

 SetImplementationNil sets the value for Implementation to be an explicit nil

### UnsetImplementation
`func (o *ImportListExclusionResource) UnsetImplementation()`

UnsetImplementation ensures that no value is present for Implementation, not even an explicit nil
### GetConfigContract

`func (o *ImportListExclusionResource) GetConfigContract() string`

GetConfigContract returns the ConfigContract field if non-nil, zero value otherwise.

### GetConfigContractOk

`func (o *ImportListExclusionResource) GetConfigContractOk() (*string, bool)`

GetConfigContractOk returns a tuple with the ConfigContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContract

`func (o *ImportListExclusionResource) SetConfigContract(v string)`

SetConfigContract sets ConfigContract field to given value.

### HasConfigContract

`func (o *ImportListExclusionResource) HasConfigContract() bool`

HasConfigContract returns a boolean if a field has been set.

### SetConfigContractNil

`func (o *ImportListExclusionResource) SetConfigContractNil(b bool)`

 SetConfigContractNil sets the value for ConfigContract to be an explicit nil

### UnsetConfigContract
`func (o *ImportListExclusionResource) UnsetConfigContract()`

UnsetConfigContract ensures that no value is present for ConfigContract, not even an explicit nil
### GetInfoLink

`func (o *ImportListExclusionResource) GetInfoLink() string`

GetInfoLink returns the InfoLink field if non-nil, zero value otherwise.

### GetInfoLinkOk

`func (o *ImportListExclusionResource) GetInfoLinkOk() (*string, bool)`

GetInfoLinkOk returns a tuple with the InfoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoLink

`func (o *ImportListExclusionResource) SetInfoLink(v string)`

SetInfoLink sets InfoLink field to given value.

### HasInfoLink

`func (o *ImportListExclusionResource) HasInfoLink() bool`

HasInfoLink returns a boolean if a field has been set.

### SetInfoLinkNil

`func (o *ImportListExclusionResource) SetInfoLinkNil(b bool)`

 SetInfoLinkNil sets the value for InfoLink to be an explicit nil

### UnsetInfoLink
`func (o *ImportListExclusionResource) UnsetInfoLink()`

UnsetInfoLink ensures that no value is present for InfoLink, not even an explicit nil
### GetMessage

`func (o *ImportListExclusionResource) GetMessage() ProviderMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImportListExclusionResource) GetMessageOk() (*ProviderMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImportListExclusionResource) SetMessage(v ProviderMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ImportListExclusionResource) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTags

`func (o *ImportListExclusionResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImportListExclusionResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImportListExclusionResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImportListExclusionResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ImportListExclusionResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ImportListExclusionResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPresets

`func (o *ImportListExclusionResource) GetPresets() []ImportListExclusionResource`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *ImportListExclusionResource) GetPresetsOk() (*[]ImportListExclusionResource, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *ImportListExclusionResource) SetPresets(v []ImportListExclusionResource)`

SetPresets sets Presets field to given value.

### HasPresets

`func (o *ImportListExclusionResource) HasPresets() bool`

HasPresets returns a boolean if a field has been set.

### SetPresetsNil

`func (o *ImportListExclusionResource) SetPresetsNil(b bool)`

 SetPresetsNil sets the value for Presets to be an explicit nil

### UnsetPresets
`func (o *ImportListExclusionResource) UnsetPresets()`

UnsetPresets ensures that no value is present for Presets, not even an explicit nil
### GetTmdbId

`func (o *ImportListExclusionResource) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *ImportListExclusionResource) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *ImportListExclusionResource) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *ImportListExclusionResource) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetMovieTitle

`func (o *ImportListExclusionResource) GetMovieTitle() string`

GetMovieTitle returns the MovieTitle field if non-nil, zero value otherwise.

### GetMovieTitleOk

`func (o *ImportListExclusionResource) GetMovieTitleOk() (*string, bool)`

GetMovieTitleOk returns a tuple with the MovieTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieTitle

`func (o *ImportListExclusionResource) SetMovieTitle(v string)`

SetMovieTitle sets MovieTitle field to given value.

### HasMovieTitle

`func (o *ImportListExclusionResource) HasMovieTitle() bool`

HasMovieTitle returns a boolean if a field has been set.

### SetMovieTitleNil

`func (o *ImportListExclusionResource) SetMovieTitleNil(b bool)`

 SetMovieTitleNil sets the value for MovieTitle to be an explicit nil

### UnsetMovieTitle
`func (o *ImportListExclusionResource) UnsetMovieTitle()`

UnsetMovieTitle ensures that no value is present for MovieTitle, not even an explicit nil
### GetMovieYear

`func (o *ImportListExclusionResource) GetMovieYear() int32`

GetMovieYear returns the MovieYear field if non-nil, zero value otherwise.

### GetMovieYearOk

`func (o *ImportListExclusionResource) GetMovieYearOk() (*int32, bool)`

GetMovieYearOk returns a tuple with the MovieYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieYear

`func (o *ImportListExclusionResource) SetMovieYear(v int32)`

SetMovieYear sets MovieYear field to given value.

### HasMovieYear

`func (o *ImportListExclusionResource) HasMovieYear() bool`

HasMovieYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



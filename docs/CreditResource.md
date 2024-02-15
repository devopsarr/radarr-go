# CreditResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PersonName** | Pointer to **NullableString** |  | [optional] 
**CreditTmdbId** | Pointer to **NullableString** |  | [optional] 
**PersonTmdbId** | Pointer to **int32** |  | [optional] 
**MovieMetadataId** | Pointer to **int32** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Department** | Pointer to **NullableString** |  | [optional] 
**Job** | Pointer to **NullableString** |  | [optional] 
**Character** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**CreditType**](CreditType.md) |  | [optional] 

## Methods

### NewCreditResource

`func NewCreditResource() *CreditResource`

NewCreditResource instantiates a new CreditResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditResourceWithDefaults

`func NewCreditResourceWithDefaults() *CreditResource`

NewCreditResourceWithDefaults instantiates a new CreditResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreditResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPersonName

`func (o *CreditResource) GetPersonName() string`

GetPersonName returns the PersonName field if non-nil, zero value otherwise.

### GetPersonNameOk

`func (o *CreditResource) GetPersonNameOk() (*string, bool)`

GetPersonNameOk returns a tuple with the PersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonName

`func (o *CreditResource) SetPersonName(v string)`

SetPersonName sets PersonName field to given value.

### HasPersonName

`func (o *CreditResource) HasPersonName() bool`

HasPersonName returns a boolean if a field has been set.

### SetPersonNameNil

`func (o *CreditResource) SetPersonNameNil(b bool)`

 SetPersonNameNil sets the value for PersonName to be an explicit nil

### UnsetPersonName
`func (o *CreditResource) UnsetPersonName()`

UnsetPersonName ensures that no value is present for PersonName, not even an explicit nil
### GetCreditTmdbId

`func (o *CreditResource) GetCreditTmdbId() string`

GetCreditTmdbId returns the CreditTmdbId field if non-nil, zero value otherwise.

### GetCreditTmdbIdOk

`func (o *CreditResource) GetCreditTmdbIdOk() (*string, bool)`

GetCreditTmdbIdOk returns a tuple with the CreditTmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditTmdbId

`func (o *CreditResource) SetCreditTmdbId(v string)`

SetCreditTmdbId sets CreditTmdbId field to given value.

### HasCreditTmdbId

`func (o *CreditResource) HasCreditTmdbId() bool`

HasCreditTmdbId returns a boolean if a field has been set.

### SetCreditTmdbIdNil

`func (o *CreditResource) SetCreditTmdbIdNil(b bool)`

 SetCreditTmdbIdNil sets the value for CreditTmdbId to be an explicit nil

### UnsetCreditTmdbId
`func (o *CreditResource) UnsetCreditTmdbId()`

UnsetCreditTmdbId ensures that no value is present for CreditTmdbId, not even an explicit nil
### GetPersonTmdbId

`func (o *CreditResource) GetPersonTmdbId() int32`

GetPersonTmdbId returns the PersonTmdbId field if non-nil, zero value otherwise.

### GetPersonTmdbIdOk

`func (o *CreditResource) GetPersonTmdbIdOk() (*int32, bool)`

GetPersonTmdbIdOk returns a tuple with the PersonTmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonTmdbId

`func (o *CreditResource) SetPersonTmdbId(v int32)`

SetPersonTmdbId sets PersonTmdbId field to given value.

### HasPersonTmdbId

`func (o *CreditResource) HasPersonTmdbId() bool`

HasPersonTmdbId returns a boolean if a field has been set.

### GetMovieMetadataId

`func (o *CreditResource) GetMovieMetadataId() int32`

GetMovieMetadataId returns the MovieMetadataId field if non-nil, zero value otherwise.

### GetMovieMetadataIdOk

`func (o *CreditResource) GetMovieMetadataIdOk() (*int32, bool)`

GetMovieMetadataIdOk returns a tuple with the MovieMetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieMetadataId

`func (o *CreditResource) SetMovieMetadataId(v int32)`

SetMovieMetadataId sets MovieMetadataId field to given value.

### HasMovieMetadataId

`func (o *CreditResource) HasMovieMetadataId() bool`

HasMovieMetadataId returns a boolean if a field has been set.

### GetImages

`func (o *CreditResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CreditResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CreditResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *CreditResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CreditResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CreditResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetDepartment

`func (o *CreditResource) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CreditResource) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CreditResource) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *CreditResource) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *CreditResource) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *CreditResource) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetJob

`func (o *CreditResource) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CreditResource) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CreditResource) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *CreditResource) HasJob() bool`

HasJob returns a boolean if a field has been set.

### SetJobNil

`func (o *CreditResource) SetJobNil(b bool)`

 SetJobNil sets the value for Job to be an explicit nil

### UnsetJob
`func (o *CreditResource) UnsetJob()`

UnsetJob ensures that no value is present for Job, not even an explicit nil
### GetCharacter

`func (o *CreditResource) GetCharacter() string`

GetCharacter returns the Character field if non-nil, zero value otherwise.

### GetCharacterOk

`func (o *CreditResource) GetCharacterOk() (*string, bool)`

GetCharacterOk returns a tuple with the Character field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacter

`func (o *CreditResource) SetCharacter(v string)`

SetCharacter sets Character field to given value.

### HasCharacter

`func (o *CreditResource) HasCharacter() bool`

HasCharacter returns a boolean if a field has been set.

### SetCharacterNil

`func (o *CreditResource) SetCharacterNil(b bool)`

 SetCharacterNil sets the value for Character to be an explicit nil

### UnsetCharacter
`func (o *CreditResource) UnsetCharacter()`

UnsetCharacter ensures that no value is present for Character, not even an explicit nil
### GetOrder

`func (o *CreditResource) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreditResource) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreditResource) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreditResource) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *CreditResource) GetType() CreditType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditResource) GetTypeOk() (*CreditType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditResource) SetType(v CreditType)`

SetType sets Type field to given value.

### HasType

`func (o *CreditResource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



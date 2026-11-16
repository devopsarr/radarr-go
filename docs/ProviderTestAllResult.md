# ProviderTestAllResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] [readonly] 
**ValidationFailures** | Pointer to [**[]ValidationFailure**](ValidationFailure.md) |  | [optional] 

## Methods

### NewProviderTestAllResult

`func NewProviderTestAllResult() *ProviderTestAllResult`

NewProviderTestAllResult instantiates a new ProviderTestAllResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderTestAllResultWithDefaults

`func NewProviderTestAllResultWithDefaults() *ProviderTestAllResult`

NewProviderTestAllResultWithDefaults instantiates a new ProviderTestAllResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderTestAllResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderTestAllResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderTestAllResult) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderTestAllResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsValid

`func (o *ProviderTestAllResult) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ProviderTestAllResult) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ProviderTestAllResult) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *ProviderTestAllResult) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetValidationFailures

`func (o *ProviderTestAllResult) GetValidationFailures() []ValidationFailure`

GetValidationFailures returns the ValidationFailures field if non-nil, zero value otherwise.

### GetValidationFailuresOk

`func (o *ProviderTestAllResult) GetValidationFailuresOk() (*[]ValidationFailure, bool)`

GetValidationFailuresOk returns a tuple with the ValidationFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationFailures

`func (o *ProviderTestAllResult) SetValidationFailures(v []ValidationFailure)`

SetValidationFailures sets ValidationFailures field to given value.

### HasValidationFailures

`func (o *ProviderTestAllResult) HasValidationFailures() bool`

HasValidationFailures returns a boolean if a field has been set.

### SetValidationFailuresNil

`func (o *ProviderTestAllResult) SetValidationFailuresNil(b bool)`

 SetValidationFailuresNil sets the value for ValidationFailures to be an explicit nil

### UnsetValidationFailures
`func (o *ProviderTestAllResult) UnsetValidationFailures()`

UnsetValidationFailures ensures that no value is present for ValidationFailures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



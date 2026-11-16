# ValidationFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**AttemptedValue** | Pointer to **interface{}** |  | [optional] 
**CustomState** | Pointer to **interface{}** |  | [optional] 
**Severity** | Pointer to [**Severity**](Severity.md) |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**FormattedMessageArguments** | Pointer to **[]interface{}** |  | [optional] 
**FormattedMessagePlaceholderValues** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewValidationFailure

`func NewValidationFailure() *ValidationFailure`

NewValidationFailure instantiates a new ValidationFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationFailureWithDefaults

`func NewValidationFailureWithDefaults() *ValidationFailure`

NewValidationFailureWithDefaults instantiates a new ValidationFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *ValidationFailure) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *ValidationFailure) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *ValidationFailure) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *ValidationFailure) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### SetPropertyNameNil

`func (o *ValidationFailure) SetPropertyNameNil(b bool)`

 SetPropertyNameNil sets the value for PropertyName to be an explicit nil

### UnsetPropertyName
`func (o *ValidationFailure) UnsetPropertyName()`

UnsetPropertyName ensures that no value is present for PropertyName, not even an explicit nil
### GetErrorMessage

`func (o *ValidationFailure) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ValidationFailure) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ValidationFailure) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ValidationFailure) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ValidationFailure) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ValidationFailure) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetAttemptedValue

`func (o *ValidationFailure) GetAttemptedValue() interface{}`

GetAttemptedValue returns the AttemptedValue field if non-nil, zero value otherwise.

### GetAttemptedValueOk

`func (o *ValidationFailure) GetAttemptedValueOk() (*interface{}, bool)`

GetAttemptedValueOk returns a tuple with the AttemptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedValue

`func (o *ValidationFailure) SetAttemptedValue(v interface{})`

SetAttemptedValue sets AttemptedValue field to given value.

### HasAttemptedValue

`func (o *ValidationFailure) HasAttemptedValue() bool`

HasAttemptedValue returns a boolean if a field has been set.

### SetAttemptedValueNil

`func (o *ValidationFailure) SetAttemptedValueNil(b bool)`

 SetAttemptedValueNil sets the value for AttemptedValue to be an explicit nil

### UnsetAttemptedValue
`func (o *ValidationFailure) UnsetAttemptedValue()`

UnsetAttemptedValue ensures that no value is present for AttemptedValue, not even an explicit nil
### GetCustomState

`func (o *ValidationFailure) GetCustomState() interface{}`

GetCustomState returns the CustomState field if non-nil, zero value otherwise.

### GetCustomStateOk

`func (o *ValidationFailure) GetCustomStateOk() (*interface{}, bool)`

GetCustomStateOk returns a tuple with the CustomState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomState

`func (o *ValidationFailure) SetCustomState(v interface{})`

SetCustomState sets CustomState field to given value.

### HasCustomState

`func (o *ValidationFailure) HasCustomState() bool`

HasCustomState returns a boolean if a field has been set.

### SetCustomStateNil

`func (o *ValidationFailure) SetCustomStateNil(b bool)`

 SetCustomStateNil sets the value for CustomState to be an explicit nil

### UnsetCustomState
`func (o *ValidationFailure) UnsetCustomState()`

UnsetCustomState ensures that no value is present for CustomState, not even an explicit nil
### GetSeverity

`func (o *ValidationFailure) GetSeverity() Severity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ValidationFailure) GetSeverityOk() (*Severity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ValidationFailure) SetSeverity(v Severity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ValidationFailure) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetErrorCode

`func (o *ValidationFailure) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ValidationFailure) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ValidationFailure) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ValidationFailure) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ValidationFailure) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ValidationFailure) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetFormattedMessageArguments

`func (o *ValidationFailure) GetFormattedMessageArguments() []interface{}`

GetFormattedMessageArguments returns the FormattedMessageArguments field if non-nil, zero value otherwise.

### GetFormattedMessageArgumentsOk

`func (o *ValidationFailure) GetFormattedMessageArgumentsOk() (*[]interface{}, bool)`

GetFormattedMessageArgumentsOk returns a tuple with the FormattedMessageArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedMessageArguments

`func (o *ValidationFailure) SetFormattedMessageArguments(v []interface{})`

SetFormattedMessageArguments sets FormattedMessageArguments field to given value.

### HasFormattedMessageArguments

`func (o *ValidationFailure) HasFormattedMessageArguments() bool`

HasFormattedMessageArguments returns a boolean if a field has been set.

### SetFormattedMessageArgumentsNil

`func (o *ValidationFailure) SetFormattedMessageArgumentsNil(b bool)`

 SetFormattedMessageArgumentsNil sets the value for FormattedMessageArguments to be an explicit nil

### UnsetFormattedMessageArguments
`func (o *ValidationFailure) UnsetFormattedMessageArguments()`

UnsetFormattedMessageArguments ensures that no value is present for FormattedMessageArguments, not even an explicit nil
### GetFormattedMessagePlaceholderValues

`func (o *ValidationFailure) GetFormattedMessagePlaceholderValues() map[string]interface{}`

GetFormattedMessagePlaceholderValues returns the FormattedMessagePlaceholderValues field if non-nil, zero value otherwise.

### GetFormattedMessagePlaceholderValuesOk

`func (o *ValidationFailure) GetFormattedMessagePlaceholderValuesOk() (*map[string]interface{}, bool)`

GetFormattedMessagePlaceholderValuesOk returns a tuple with the FormattedMessagePlaceholderValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedMessagePlaceholderValues

`func (o *ValidationFailure) SetFormattedMessagePlaceholderValues(v map[string]interface{})`

SetFormattedMessagePlaceholderValues sets FormattedMessagePlaceholderValues field to given value.

### HasFormattedMessagePlaceholderValues

`func (o *ValidationFailure) HasFormattedMessagePlaceholderValues() bool`

HasFormattedMessagePlaceholderValues returns a boolean if a field has been set.

### SetFormattedMessagePlaceholderValuesNil

`func (o *ValidationFailure) SetFormattedMessagePlaceholderValuesNil(b bool)`

 SetFormattedMessagePlaceholderValuesNil sets the value for FormattedMessagePlaceholderValues to be an explicit nil

### UnsetFormattedMessagePlaceholderValues
`func (o *ValidationFailure) UnsetFormattedMessagePlaceholderValues()`

UnsetFormattedMessagePlaceholderValues ensures that no value is present for FormattedMessagePlaceholderValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



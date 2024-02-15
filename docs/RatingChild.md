# RatingChild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Votes** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to [**RatingType**](RatingType.md) |  | [optional] 

## Methods

### NewRatingChild

`func NewRatingChild() *RatingChild`

NewRatingChild instantiates a new RatingChild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingChildWithDefaults

`func NewRatingChildWithDefaults() *RatingChild`

NewRatingChildWithDefaults instantiates a new RatingChild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVotes

`func (o *RatingChild) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *RatingChild) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *RatingChild) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *RatingChild) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetValue

`func (o *RatingChild) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RatingChild) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RatingChild) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *RatingChild) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *RatingChild) GetType() RatingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RatingChild) GetTypeOk() (*RatingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RatingChild) SetType(v RatingType)`

SetType sets Type field to given value.

### HasType

`func (o *RatingChild) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



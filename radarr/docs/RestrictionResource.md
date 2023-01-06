# RestrictionResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Required** | Pointer to **NullableString** |  | [optional] 
**Preferred** | Pointer to **NullableString** |  | [optional] 
**Ignored** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewRestrictionResource

`func NewRestrictionResource() *RestrictionResource`

NewRestrictionResource instantiates a new RestrictionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictionResourceWithDefaults

`func NewRestrictionResourceWithDefaults() *RestrictionResource`

NewRestrictionResourceWithDefaults instantiates a new RestrictionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestrictionResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestrictionResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestrictionResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RestrictionResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequired

`func (o *RestrictionResource) GetRequired() string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RestrictionResource) GetRequiredOk() (*string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RestrictionResource) SetRequired(v string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *RestrictionResource) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *RestrictionResource) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *RestrictionResource) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetPreferred

`func (o *RestrictionResource) GetPreferred() string`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *RestrictionResource) GetPreferredOk() (*string, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *RestrictionResource) SetPreferred(v string)`

SetPreferred sets Preferred field to given value.

### HasPreferred

`func (o *RestrictionResource) HasPreferred() bool`

HasPreferred returns a boolean if a field has been set.

### SetPreferredNil

`func (o *RestrictionResource) SetPreferredNil(b bool)`

 SetPreferredNil sets the value for Preferred to be an explicit nil

### UnsetPreferred
`func (o *RestrictionResource) UnsetPreferred()`

UnsetPreferred ensures that no value is present for Preferred, not even an explicit nil
### GetIgnored

`func (o *RestrictionResource) GetIgnored() string`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *RestrictionResource) GetIgnoredOk() (*string, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *RestrictionResource) SetIgnored(v string)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *RestrictionResource) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### SetIgnoredNil

`func (o *RestrictionResource) SetIgnoredNil(b bool)`

 SetIgnoredNil sets the value for Ignored to be an explicit nil

### UnsetIgnored
`func (o *RestrictionResource) UnsetIgnored()`

UnsetIgnored ensures that no value is present for Ignored, not even an explicit nil
### GetTags

`func (o *RestrictionResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RestrictionResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RestrictionResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RestrictionResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *RestrictionResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *RestrictionResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



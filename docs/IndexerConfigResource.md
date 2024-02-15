# IndexerConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MinimumAge** | Pointer to **int32** |  | [optional] 
**MaximumSize** | Pointer to **int32** |  | [optional] 
**Retention** | Pointer to **int32** |  | [optional] 
**RssSyncInterval** | Pointer to **int32** |  | [optional] 
**PreferIndexerFlags** | Pointer to **bool** |  | [optional] 
**AvailabilityDelay** | Pointer to **int32** |  | [optional] 
**AllowHardcodedSubs** | Pointer to **bool** |  | [optional] 
**WhitelistedHardcodedSubs** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIndexerConfigResource

`func NewIndexerConfigResource() *IndexerConfigResource`

NewIndexerConfigResource instantiates a new IndexerConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexerConfigResourceWithDefaults

`func NewIndexerConfigResourceWithDefaults() *IndexerConfigResource`

NewIndexerConfigResourceWithDefaults instantiates a new IndexerConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexerConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexerConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexerConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IndexerConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMinimumAge

`func (o *IndexerConfigResource) GetMinimumAge() int32`

GetMinimumAge returns the MinimumAge field if non-nil, zero value otherwise.

### GetMinimumAgeOk

`func (o *IndexerConfigResource) GetMinimumAgeOk() (*int32, bool)`

GetMinimumAgeOk returns a tuple with the MinimumAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAge

`func (o *IndexerConfigResource) SetMinimumAge(v int32)`

SetMinimumAge sets MinimumAge field to given value.

### HasMinimumAge

`func (o *IndexerConfigResource) HasMinimumAge() bool`

HasMinimumAge returns a boolean if a field has been set.

### GetMaximumSize

`func (o *IndexerConfigResource) GetMaximumSize() int32`

GetMaximumSize returns the MaximumSize field if non-nil, zero value otherwise.

### GetMaximumSizeOk

`func (o *IndexerConfigResource) GetMaximumSizeOk() (*int32, bool)`

GetMaximumSizeOk returns a tuple with the MaximumSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSize

`func (o *IndexerConfigResource) SetMaximumSize(v int32)`

SetMaximumSize sets MaximumSize field to given value.

### HasMaximumSize

`func (o *IndexerConfigResource) HasMaximumSize() bool`

HasMaximumSize returns a boolean if a field has been set.

### GetRetention

`func (o *IndexerConfigResource) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *IndexerConfigResource) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *IndexerConfigResource) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *IndexerConfigResource) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetRssSyncInterval

`func (o *IndexerConfigResource) GetRssSyncInterval() int32`

GetRssSyncInterval returns the RssSyncInterval field if non-nil, zero value otherwise.

### GetRssSyncIntervalOk

`func (o *IndexerConfigResource) GetRssSyncIntervalOk() (*int32, bool)`

GetRssSyncIntervalOk returns a tuple with the RssSyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssSyncInterval

`func (o *IndexerConfigResource) SetRssSyncInterval(v int32)`

SetRssSyncInterval sets RssSyncInterval field to given value.

### HasRssSyncInterval

`func (o *IndexerConfigResource) HasRssSyncInterval() bool`

HasRssSyncInterval returns a boolean if a field has been set.

### GetPreferIndexerFlags

`func (o *IndexerConfigResource) GetPreferIndexerFlags() bool`

GetPreferIndexerFlags returns the PreferIndexerFlags field if non-nil, zero value otherwise.

### GetPreferIndexerFlagsOk

`func (o *IndexerConfigResource) GetPreferIndexerFlagsOk() (*bool, bool)`

GetPreferIndexerFlagsOk returns a tuple with the PreferIndexerFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferIndexerFlags

`func (o *IndexerConfigResource) SetPreferIndexerFlags(v bool)`

SetPreferIndexerFlags sets PreferIndexerFlags field to given value.

### HasPreferIndexerFlags

`func (o *IndexerConfigResource) HasPreferIndexerFlags() bool`

HasPreferIndexerFlags returns a boolean if a field has been set.

### GetAvailabilityDelay

`func (o *IndexerConfigResource) GetAvailabilityDelay() int32`

GetAvailabilityDelay returns the AvailabilityDelay field if non-nil, zero value otherwise.

### GetAvailabilityDelayOk

`func (o *IndexerConfigResource) GetAvailabilityDelayOk() (*int32, bool)`

GetAvailabilityDelayOk returns a tuple with the AvailabilityDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityDelay

`func (o *IndexerConfigResource) SetAvailabilityDelay(v int32)`

SetAvailabilityDelay sets AvailabilityDelay field to given value.

### HasAvailabilityDelay

`func (o *IndexerConfigResource) HasAvailabilityDelay() bool`

HasAvailabilityDelay returns a boolean if a field has been set.

### GetAllowHardcodedSubs

`func (o *IndexerConfigResource) GetAllowHardcodedSubs() bool`

GetAllowHardcodedSubs returns the AllowHardcodedSubs field if non-nil, zero value otherwise.

### GetAllowHardcodedSubsOk

`func (o *IndexerConfigResource) GetAllowHardcodedSubsOk() (*bool, bool)`

GetAllowHardcodedSubsOk returns a tuple with the AllowHardcodedSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHardcodedSubs

`func (o *IndexerConfigResource) SetAllowHardcodedSubs(v bool)`

SetAllowHardcodedSubs sets AllowHardcodedSubs field to given value.

### HasAllowHardcodedSubs

`func (o *IndexerConfigResource) HasAllowHardcodedSubs() bool`

HasAllowHardcodedSubs returns a boolean if a field has been set.

### GetWhitelistedHardcodedSubs

`func (o *IndexerConfigResource) GetWhitelistedHardcodedSubs() string`

GetWhitelistedHardcodedSubs returns the WhitelistedHardcodedSubs field if non-nil, zero value otherwise.

### GetWhitelistedHardcodedSubsOk

`func (o *IndexerConfigResource) GetWhitelistedHardcodedSubsOk() (*string, bool)`

GetWhitelistedHardcodedSubsOk returns a tuple with the WhitelistedHardcodedSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistedHardcodedSubs

`func (o *IndexerConfigResource) SetWhitelistedHardcodedSubs(v string)`

SetWhitelistedHardcodedSubs sets WhitelistedHardcodedSubs field to given value.

### HasWhitelistedHardcodedSubs

`func (o *IndexerConfigResource) HasWhitelistedHardcodedSubs() bool`

HasWhitelistedHardcodedSubs returns a boolean if a field has been set.

### SetWhitelistedHardcodedSubsNil

`func (o *IndexerConfigResource) SetWhitelistedHardcodedSubsNil(b bool)`

 SetWhitelistedHardcodedSubsNil sets the value for WhitelistedHardcodedSubs to be an explicit nil

### UnsetWhitelistedHardcodedSubs
`func (o *IndexerConfigResource) UnsetWhitelistedHardcodedSubs()`

UnsetWhitelistedHardcodedSubs ensures that no value is present for WhitelistedHardcodedSubs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



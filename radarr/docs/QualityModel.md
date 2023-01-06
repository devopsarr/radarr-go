# QualityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quality** | Pointer to [**Quality**](Quality.md) |  | [optional] 
**Revision** | Pointer to [**Revision**](Revision.md) |  | [optional] 
**HardcodedSubs** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQualityModel

`func NewQualityModel() *QualityModel`

NewQualityModel instantiates a new QualityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQualityModelWithDefaults

`func NewQualityModelWithDefaults() *QualityModel`

NewQualityModelWithDefaults instantiates a new QualityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuality

`func (o *QualityModel) GetQuality() Quality`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *QualityModel) GetQualityOk() (*Quality, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *QualityModel) SetQuality(v Quality)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *QualityModel) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetRevision

`func (o *QualityModel) GetRevision() Revision`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *QualityModel) GetRevisionOk() (*Revision, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *QualityModel) SetRevision(v Revision)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *QualityModel) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetHardcodedSubs

`func (o *QualityModel) GetHardcodedSubs() string`

GetHardcodedSubs returns the HardcodedSubs field if non-nil, zero value otherwise.

### GetHardcodedSubsOk

`func (o *QualityModel) GetHardcodedSubsOk() (*string, bool)`

GetHardcodedSubsOk returns a tuple with the HardcodedSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardcodedSubs

`func (o *QualityModel) SetHardcodedSubs(v string)`

SetHardcodedSubs sets HardcodedSubs field to given value.

### HasHardcodedSubs

`func (o *QualityModel) HasHardcodedSubs() bool`

HasHardcodedSubs returns a boolean if a field has been set.

### SetHardcodedSubsNil

`func (o *QualityModel) SetHardcodedSubsNil(b bool)`

 SetHardcodedSubsNil sets the value for HardcodedSubs to be an explicit nil

### UnsetHardcodedSubs
`func (o *QualityModel) UnsetHardcodedSubs()`

UnsetHardcodedSubs ensures that no value is present for HardcodedSubs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



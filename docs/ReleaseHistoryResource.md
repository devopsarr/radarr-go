# ReleaseHistoryResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grabbed** | Pointer to **NullableTime** |  | [optional] 
**Failed** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewReleaseHistoryResource

`func NewReleaseHistoryResource() *ReleaseHistoryResource`

NewReleaseHistoryResource instantiates a new ReleaseHistoryResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseHistoryResourceWithDefaults

`func NewReleaseHistoryResourceWithDefaults() *ReleaseHistoryResource`

NewReleaseHistoryResourceWithDefaults instantiates a new ReleaseHistoryResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrabbed

`func (o *ReleaseHistoryResource) GetGrabbed() time.Time`

GetGrabbed returns the Grabbed field if non-nil, zero value otherwise.

### GetGrabbedOk

`func (o *ReleaseHistoryResource) GetGrabbedOk() (*time.Time, bool)`

GetGrabbedOk returns a tuple with the Grabbed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrabbed

`func (o *ReleaseHistoryResource) SetGrabbed(v time.Time)`

SetGrabbed sets Grabbed field to given value.

### HasGrabbed

`func (o *ReleaseHistoryResource) HasGrabbed() bool`

HasGrabbed returns a boolean if a field has been set.

### SetGrabbedNil

`func (o *ReleaseHistoryResource) SetGrabbedNil(b bool)`

 SetGrabbedNil sets the value for Grabbed to be an explicit nil

### UnsetGrabbed
`func (o *ReleaseHistoryResource) UnsetGrabbed()`

UnsetGrabbed ensures that no value is present for Grabbed, not even an explicit nil
### GetFailed

`func (o *ReleaseHistoryResource) GetFailed() time.Time`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ReleaseHistoryResource) GetFailedOk() (*time.Time, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ReleaseHistoryResource) SetFailed(v time.Time)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ReleaseHistoryResource) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *ReleaseHistoryResource) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *ReleaseHistoryResource) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



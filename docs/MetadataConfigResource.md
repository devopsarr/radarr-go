# MetadataConfigResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CertificationCountry** | Pointer to [**TMDbCountryCode**](TMDbCountryCode.md) |  | [optional] 

## Methods

### NewMetadataConfigResource

`func NewMetadataConfigResource() *MetadataConfigResource`

NewMetadataConfigResource instantiates a new MetadataConfigResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataConfigResourceWithDefaults

`func NewMetadataConfigResourceWithDefaults() *MetadataConfigResource`

NewMetadataConfigResourceWithDefaults instantiates a new MetadataConfigResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataConfigResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataConfigResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataConfigResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataConfigResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCertificationCountry

`func (o *MetadataConfigResource) GetCertificationCountry() TMDbCountryCode`

GetCertificationCountry returns the CertificationCountry field if non-nil, zero value otherwise.

### GetCertificationCountryOk

`func (o *MetadataConfigResource) GetCertificationCountryOk() (*TMDbCountryCode, bool)`

GetCertificationCountryOk returns a tuple with the CertificationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationCountry

`func (o *MetadataConfigResource) SetCertificationCountry(v TMDbCountryCode)`

SetCertificationCountry sets CertificationCountry field to given value.

### HasCertificationCountry

`func (o *MetadataConfigResource) HasCertificationCountry() bool`

HasCertificationCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MovieCollectionResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 

## Methods

### NewMovieCollectionResource

`func NewMovieCollectionResource() *MovieCollectionResource`

NewMovieCollectionResource instantiates a new MovieCollectionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieCollectionResourceWithDefaults

`func NewMovieCollectionResourceWithDefaults() *MovieCollectionResource`

NewMovieCollectionResourceWithDefaults instantiates a new MovieCollectionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *MovieCollectionResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MovieCollectionResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MovieCollectionResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MovieCollectionResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MovieCollectionResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MovieCollectionResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTmdbId

`func (o *MovieCollectionResource) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *MovieCollectionResource) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *MovieCollectionResource) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *MovieCollectionResource) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



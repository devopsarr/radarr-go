# AddMovieOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreEpisodesWithFiles** | Pointer to **bool** |  | [optional] 
**IgnoreEpisodesWithoutFiles** | Pointer to **bool** |  | [optional] 
**Monitor** | Pointer to [**MonitorTypes**](MonitorTypes.md) |  | [optional] 
**SearchForMovie** | Pointer to **bool** |  | [optional] 
**AddMethod** | Pointer to [**AddMovieMethod**](AddMovieMethod.md) |  | [optional] 

## Methods

### NewAddMovieOptions

`func NewAddMovieOptions() *AddMovieOptions`

NewAddMovieOptions instantiates a new AddMovieOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMovieOptionsWithDefaults

`func NewAddMovieOptionsWithDefaults() *AddMovieOptions`

NewAddMovieOptionsWithDefaults instantiates a new AddMovieOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreEpisodesWithFiles

`func (o *AddMovieOptions) GetIgnoreEpisodesWithFiles() bool`

GetIgnoreEpisodesWithFiles returns the IgnoreEpisodesWithFiles field if non-nil, zero value otherwise.

### GetIgnoreEpisodesWithFilesOk

`func (o *AddMovieOptions) GetIgnoreEpisodesWithFilesOk() (*bool, bool)`

GetIgnoreEpisodesWithFilesOk returns a tuple with the IgnoreEpisodesWithFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEpisodesWithFiles

`func (o *AddMovieOptions) SetIgnoreEpisodesWithFiles(v bool)`

SetIgnoreEpisodesWithFiles sets IgnoreEpisodesWithFiles field to given value.

### HasIgnoreEpisodesWithFiles

`func (o *AddMovieOptions) HasIgnoreEpisodesWithFiles() bool`

HasIgnoreEpisodesWithFiles returns a boolean if a field has been set.

### GetIgnoreEpisodesWithoutFiles

`func (o *AddMovieOptions) GetIgnoreEpisodesWithoutFiles() bool`

GetIgnoreEpisodesWithoutFiles returns the IgnoreEpisodesWithoutFiles field if non-nil, zero value otherwise.

### GetIgnoreEpisodesWithoutFilesOk

`func (o *AddMovieOptions) GetIgnoreEpisodesWithoutFilesOk() (*bool, bool)`

GetIgnoreEpisodesWithoutFilesOk returns a tuple with the IgnoreEpisodesWithoutFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEpisodesWithoutFiles

`func (o *AddMovieOptions) SetIgnoreEpisodesWithoutFiles(v bool)`

SetIgnoreEpisodesWithoutFiles sets IgnoreEpisodesWithoutFiles field to given value.

### HasIgnoreEpisodesWithoutFiles

`func (o *AddMovieOptions) HasIgnoreEpisodesWithoutFiles() bool`

HasIgnoreEpisodesWithoutFiles returns a boolean if a field has been set.

### GetMonitor

`func (o *AddMovieOptions) GetMonitor() MonitorTypes`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *AddMovieOptions) GetMonitorOk() (*MonitorTypes, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *AddMovieOptions) SetMonitor(v MonitorTypes)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *AddMovieOptions) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetSearchForMovie

`func (o *AddMovieOptions) GetSearchForMovie() bool`

GetSearchForMovie returns the SearchForMovie field if non-nil, zero value otherwise.

### GetSearchForMovieOk

`func (o *AddMovieOptions) GetSearchForMovieOk() (*bool, bool)`

GetSearchForMovieOk returns a tuple with the SearchForMovie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchForMovie

`func (o *AddMovieOptions) SetSearchForMovie(v bool)`

SetSearchForMovie sets SearchForMovie field to given value.

### HasSearchForMovie

`func (o *AddMovieOptions) HasSearchForMovie() bool`

HasSearchForMovie returns a boolean if a field has been set.

### GetAddMethod

`func (o *AddMovieOptions) GetAddMethod() AddMovieMethod`

GetAddMethod returns the AddMethod field if non-nil, zero value otherwise.

### GetAddMethodOk

`func (o *AddMovieOptions) GetAddMethodOk() (*AddMovieMethod, bool)`

GetAddMethodOk returns a tuple with the AddMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddMethod

`func (o *AddMovieOptions) SetAddMethod(v AddMovieMethod)`

SetAddMethod sets AddMethod field to given value.

### HasAddMethod

`func (o *AddMovieOptions) HasAddMethod() bool`

HasAddMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Ratings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imdb** | Pointer to [**RatingChild**](RatingChild.md) |  | [optional] 
**Tmdb** | Pointer to [**RatingChild**](RatingChild.md) |  | [optional] 
**Metacritic** | Pointer to [**RatingChild**](RatingChild.md) |  | [optional] 
**RottenTomatoes** | Pointer to [**RatingChild**](RatingChild.md) |  | [optional] 

## Methods

### NewRatings

`func NewRatings() *Ratings`

NewRatings instantiates a new Ratings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingsWithDefaults

`func NewRatingsWithDefaults() *Ratings`

NewRatingsWithDefaults instantiates a new Ratings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImdb

`func (o *Ratings) GetImdb() RatingChild`

GetImdb returns the Imdb field if non-nil, zero value otherwise.

### GetImdbOk

`func (o *Ratings) GetImdbOk() (*RatingChild, bool)`

GetImdbOk returns a tuple with the Imdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdb

`func (o *Ratings) SetImdb(v RatingChild)`

SetImdb sets Imdb field to given value.

### HasImdb

`func (o *Ratings) HasImdb() bool`

HasImdb returns a boolean if a field has been set.

### GetTmdb

`func (o *Ratings) GetTmdb() RatingChild`

GetTmdb returns the Tmdb field if non-nil, zero value otherwise.

### GetTmdbOk

`func (o *Ratings) GetTmdbOk() (*RatingChild, bool)`

GetTmdbOk returns a tuple with the Tmdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdb

`func (o *Ratings) SetTmdb(v RatingChild)`

SetTmdb sets Tmdb field to given value.

### HasTmdb

`func (o *Ratings) HasTmdb() bool`

HasTmdb returns a boolean if a field has been set.

### GetMetacritic

`func (o *Ratings) GetMetacritic() RatingChild`

GetMetacritic returns the Metacritic field if non-nil, zero value otherwise.

### GetMetacriticOk

`func (o *Ratings) GetMetacriticOk() (*RatingChild, bool)`

GetMetacriticOk returns a tuple with the Metacritic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetacritic

`func (o *Ratings) SetMetacritic(v RatingChild)`

SetMetacritic sets Metacritic field to given value.

### HasMetacritic

`func (o *Ratings) HasMetacritic() bool`

HasMetacritic returns a boolean if a field has been set.

### GetRottenTomatoes

`func (o *Ratings) GetRottenTomatoes() RatingChild`

GetRottenTomatoes returns the RottenTomatoes field if non-nil, zero value otherwise.

### GetRottenTomatoesOk

`func (o *Ratings) GetRottenTomatoesOk() (*RatingChild, bool)`

GetRottenTomatoesOk returns a tuple with the RottenTomatoes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRottenTomatoes

`func (o *Ratings) SetRottenTomatoes(v RatingChild)`

SetRottenTomatoes sets RottenTomatoes field to given value.

### HasRottenTomatoes

`func (o *Ratings) HasRottenTomatoes() bool`

HasRottenTomatoes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



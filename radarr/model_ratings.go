/*
Radarr

Radarr API docs

API version: v5.16.3.9541
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the Ratings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ratings{}

// Ratings struct for Ratings
type Ratings struct {
	Imdb *RatingChild `json:"imdb,omitempty"`
	Tmdb *RatingChild `json:"tmdb,omitempty"`
	Metacritic *RatingChild `json:"metacritic,omitempty"`
	RottenTomatoes *RatingChild `json:"rottenTomatoes,omitempty"`
	Trakt *RatingChild `json:"trakt,omitempty"`
}

// NewRatings instantiates a new Ratings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatings() *Ratings {
	this := Ratings{}
	return &this
}

// NewRatingsWithDefaults instantiates a new Ratings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatingsWithDefaults() *Ratings {
	this := Ratings{}
	return &this
}

// GetImdb returns the Imdb field value if set, zero value otherwise.
func (o *Ratings) GetImdb() RatingChild {
	if o == nil || IsNil(o.Imdb) {
		var ret RatingChild
		return ret
	}
	return *o.Imdb
}

// GetImdbOk returns a tuple with the Imdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ratings) GetImdbOk() (*RatingChild, bool) {
	if o == nil || IsNil(o.Imdb) {
		return nil, false
	}
	return o.Imdb, true
}

// HasImdb returns a boolean if a field has been set.
func (o *Ratings) HasImdb() bool {
	if o != nil && !IsNil(o.Imdb) {
		return true
	}

	return false
}

// SetImdb gets a reference to the given RatingChild and assigns it to the Imdb field.
func (o *Ratings) SetImdb(v RatingChild) {
	o.Imdb = &v
}

// GetTmdb returns the Tmdb field value if set, zero value otherwise.
func (o *Ratings) GetTmdb() RatingChild {
	if o == nil || IsNil(o.Tmdb) {
		var ret RatingChild
		return ret
	}
	return *o.Tmdb
}

// GetTmdbOk returns a tuple with the Tmdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ratings) GetTmdbOk() (*RatingChild, bool) {
	if o == nil || IsNil(o.Tmdb) {
		return nil, false
	}
	return o.Tmdb, true
}

// HasTmdb returns a boolean if a field has been set.
func (o *Ratings) HasTmdb() bool {
	if o != nil && !IsNil(o.Tmdb) {
		return true
	}

	return false
}

// SetTmdb gets a reference to the given RatingChild and assigns it to the Tmdb field.
func (o *Ratings) SetTmdb(v RatingChild) {
	o.Tmdb = &v
}

// GetMetacritic returns the Metacritic field value if set, zero value otherwise.
func (o *Ratings) GetMetacritic() RatingChild {
	if o == nil || IsNil(o.Metacritic) {
		var ret RatingChild
		return ret
	}
	return *o.Metacritic
}

// GetMetacriticOk returns a tuple with the Metacritic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ratings) GetMetacriticOk() (*RatingChild, bool) {
	if o == nil || IsNil(o.Metacritic) {
		return nil, false
	}
	return o.Metacritic, true
}

// HasMetacritic returns a boolean if a field has been set.
func (o *Ratings) HasMetacritic() bool {
	if o != nil && !IsNil(o.Metacritic) {
		return true
	}

	return false
}

// SetMetacritic gets a reference to the given RatingChild and assigns it to the Metacritic field.
func (o *Ratings) SetMetacritic(v RatingChild) {
	o.Metacritic = &v
}

// GetRottenTomatoes returns the RottenTomatoes field value if set, zero value otherwise.
func (o *Ratings) GetRottenTomatoes() RatingChild {
	if o == nil || IsNil(o.RottenTomatoes) {
		var ret RatingChild
		return ret
	}
	return *o.RottenTomatoes
}

// GetRottenTomatoesOk returns a tuple with the RottenTomatoes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ratings) GetRottenTomatoesOk() (*RatingChild, bool) {
	if o == nil || IsNil(o.RottenTomatoes) {
		return nil, false
	}
	return o.RottenTomatoes, true
}

// HasRottenTomatoes returns a boolean if a field has been set.
func (o *Ratings) HasRottenTomatoes() bool {
	if o != nil && !IsNil(o.RottenTomatoes) {
		return true
	}

	return false
}

// SetRottenTomatoes gets a reference to the given RatingChild and assigns it to the RottenTomatoes field.
func (o *Ratings) SetRottenTomatoes(v RatingChild) {
	o.RottenTomatoes = &v
}

// GetTrakt returns the Trakt field value if set, zero value otherwise.
func (o *Ratings) GetTrakt() RatingChild {
	if o == nil || IsNil(o.Trakt) {
		var ret RatingChild
		return ret
	}
	return *o.Trakt
}

// GetTraktOk returns a tuple with the Trakt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ratings) GetTraktOk() (*RatingChild, bool) {
	if o == nil || IsNil(o.Trakt) {
		return nil, false
	}
	return o.Trakt, true
}

// HasTrakt returns a boolean if a field has been set.
func (o *Ratings) HasTrakt() bool {
	if o != nil && !IsNil(o.Trakt) {
		return true
	}

	return false
}

// SetTrakt gets a reference to the given RatingChild and assigns it to the Trakt field.
func (o *Ratings) SetTrakt(v RatingChild) {
	o.Trakt = &v
}

func (o Ratings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ratings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Imdb) {
		toSerialize["imdb"] = o.Imdb
	}
	if !IsNil(o.Tmdb) {
		toSerialize["tmdb"] = o.Tmdb
	}
	if !IsNil(o.Metacritic) {
		toSerialize["metacritic"] = o.Metacritic
	}
	if !IsNil(o.RottenTomatoes) {
		toSerialize["rottenTomatoes"] = o.RottenTomatoes
	}
	if !IsNil(o.Trakt) {
		toSerialize["trakt"] = o.Trakt
	}
	return toSerialize, nil
}

type NullableRatings struct {
	value *Ratings
	isSet bool
}

func (v NullableRatings) Get() *Ratings {
	return v.value
}

func (v *NullableRatings) Set(val *Ratings) {
	v.value = val
	v.isSet = true
}

func (v NullableRatings) IsSet() bool {
	return v.isSet
}

func (v *NullableRatings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatings(val *Ratings) *NullableRatings {
	return &NullableRatings{value: val, isSet: true}
}

func (v NullableRatings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



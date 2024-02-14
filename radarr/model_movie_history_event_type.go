/*
Radarr

Radarr API docs

API version: v5.2.6.8376
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// MovieHistoryEventType the model 'MovieHistoryEventType'
type MovieHistoryEventType string

// List of MovieHistoryEventType
const (
	MOVIEHISTORYEVENTTYPE_UNKNOWN MovieHistoryEventType = "unknown"
	MOVIEHISTORYEVENTTYPE_GRABBED MovieHistoryEventType = "grabbed"
	MOVIEHISTORYEVENTTYPE_DOWNLOAD_FOLDER_IMPORTED MovieHistoryEventType = "downloadFolderImported"
	MOVIEHISTORYEVENTTYPE_DOWNLOAD_FAILED MovieHistoryEventType = "downloadFailed"
	MOVIEHISTORYEVENTTYPE_MOVIE_FILE_DELETED MovieHistoryEventType = "movieFileDeleted"
	MOVIEHISTORYEVENTTYPE_MOVIE_FOLDER_IMPORTED MovieHistoryEventType = "movieFolderImported"
	MOVIEHISTORYEVENTTYPE_MOVIE_FILE_RENAMED MovieHistoryEventType = "movieFileRenamed"
	MOVIEHISTORYEVENTTYPE_DOWNLOAD_IGNORED MovieHistoryEventType = "downloadIgnored"
)

// All allowed values of MovieHistoryEventType enum
var AllowedMovieHistoryEventTypeEnumValues = []MovieHistoryEventType{
	"unknown",
	"grabbed",
	"downloadFolderImported",
	"downloadFailed",
	"movieFileDeleted",
	"movieFolderImported",
	"movieFileRenamed",
	"downloadIgnored",
}

func (v *MovieHistoryEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MovieHistoryEventType(value)
	for _, existing := range AllowedMovieHistoryEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MovieHistoryEventType", value)
}

// NewMovieHistoryEventTypeFromValue returns a pointer to a valid MovieHistoryEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMovieHistoryEventTypeFromValue(v string) (*MovieHistoryEventType, error) {
	ev := MovieHistoryEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MovieHistoryEventType: valid values are %v", v, AllowedMovieHistoryEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MovieHistoryEventType) IsValid() bool {
	for _, existing := range AllowedMovieHistoryEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MovieHistoryEventType value
func (v MovieHistoryEventType) Ptr() *MovieHistoryEventType {
	return &v
}

type NullableMovieHistoryEventType struct {
	value *MovieHistoryEventType
	isSet bool
}

func (v NullableMovieHistoryEventType) Get() *MovieHistoryEventType {
	return v.value
}

func (v *NullableMovieHistoryEventType) Set(val *MovieHistoryEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieHistoryEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieHistoryEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieHistoryEventType(val *MovieHistoryEventType) *NullableMovieHistoryEventType {
	return &NullableMovieHistoryEventType{value: val, isSet: true}
}

func (v NullableMovieHistoryEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieHistoryEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


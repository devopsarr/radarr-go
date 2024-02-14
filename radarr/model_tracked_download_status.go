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

// TrackedDownloadStatus the model 'TrackedDownloadStatus'
type TrackedDownloadStatus string

// List of TrackedDownloadStatus
const (
	TRACKEDDOWNLOADSTATUS_OK TrackedDownloadStatus = "ok"
	TRACKEDDOWNLOADSTATUS_WARNING TrackedDownloadStatus = "warning"
	TRACKEDDOWNLOADSTATUS_ERROR TrackedDownloadStatus = "error"
)

// All allowed values of TrackedDownloadStatus enum
var AllowedTrackedDownloadStatusEnumValues = []TrackedDownloadStatus{
	"ok",
	"warning",
	"error",
}

func (v *TrackedDownloadStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrackedDownloadStatus(value)
	for _, existing := range AllowedTrackedDownloadStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrackedDownloadStatus", value)
}

// NewTrackedDownloadStatusFromValue returns a pointer to a valid TrackedDownloadStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrackedDownloadStatusFromValue(v string) (*TrackedDownloadStatus, error) {
	ev := TrackedDownloadStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrackedDownloadStatus: valid values are %v", v, AllowedTrackedDownloadStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrackedDownloadStatus) IsValid() bool {
	for _, existing := range AllowedTrackedDownloadStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrackedDownloadStatus value
func (v TrackedDownloadStatus) Ptr() *TrackedDownloadStatus {
	return &v
}

type NullableTrackedDownloadStatus struct {
	value *TrackedDownloadStatus
	isSet bool
}

func (v NullableTrackedDownloadStatus) Get() *TrackedDownloadStatus {
	return v.value
}

func (v *NullableTrackedDownloadStatus) Set(val *TrackedDownloadStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackedDownloadStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackedDownloadStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackedDownloadStatus(val *TrackedDownloadStatus) *NullableTrackedDownloadStatus {
	return &NullableTrackedDownloadStatus{value: val, isSet: true}
}

func (v NullableTrackedDownloadStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackedDownloadStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


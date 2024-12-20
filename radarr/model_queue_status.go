/*
Radarr

Radarr API docs

API version: v5.16.3.9541
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
	"fmt"
)

// QueueStatus the model 'QueueStatus'
type QueueStatus string

// List of QueueStatus
const (
	QUEUESTATUS_UNKNOWN QueueStatus = "unknown"
	QUEUESTATUS_QUEUED QueueStatus = "queued"
	QUEUESTATUS_PAUSED QueueStatus = "paused"
	QUEUESTATUS_DOWNLOADING QueueStatus = "downloading"
	QUEUESTATUS_COMPLETED QueueStatus = "completed"
	QUEUESTATUS_FAILED QueueStatus = "failed"
	QUEUESTATUS_WARNING QueueStatus = "warning"
	QUEUESTATUS_DELAY QueueStatus = "delay"
	QUEUESTATUS_DOWNLOAD_CLIENT_UNAVAILABLE QueueStatus = "downloadClientUnavailable"
	QUEUESTATUS_FALLBACK QueueStatus = "fallback"
)

// All allowed values of QueueStatus enum
var AllowedQueueStatusEnumValues = []QueueStatus{
	"unknown",
	"queued",
	"paused",
	"downloading",
	"completed",
	"failed",
	"warning",
	"delay",
	"downloadClientUnavailable",
	"fallback",
}

func (v *QueueStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueueStatus(value)
	for _, existing := range AllowedQueueStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueueStatus", value)
}

// NewQueueStatusFromValue returns a pointer to a valid QueueStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueueStatusFromValue(v string) (*QueueStatus, error) {
	ev := QueueStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueueStatus: valid values are %v", v, AllowedQueueStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueueStatus) IsValid() bool {
	for _, existing := range AllowedQueueStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueueStatus value
func (v QueueStatus) Ptr() *QueueStatus {
	return &v
}

type NullableQueueStatus struct {
	value *QueueStatus
	isSet bool
}

func (v NullableQueueStatus) Get() *QueueStatus {
	return v.value
}

func (v *NullableQueueStatus) Set(val *QueueStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueStatus(val *QueueStatus) *NullableQueueStatus {
	return &NullableQueueStatus{value: val, isSet: true}
}

func (v NullableQueueStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


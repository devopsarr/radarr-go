/*
Radarr

Radarr API docs

API version: v5.8.3.8933
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"encoding/json"
)

// checks if the MediaInfoResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaInfoResource{}

// MediaInfoResource struct for MediaInfoResource
type MediaInfoResource struct {
	Id *int32 `json:"id,omitempty"`
	AudioBitrate *int64 `json:"audioBitrate,omitempty"`
	AudioChannels *float64 `json:"audioChannels,omitempty"`
	AudioCodec NullableString `json:"audioCodec,omitempty"`
	AudioLanguages NullableString `json:"audioLanguages,omitempty"`
	AudioStreamCount *int32 `json:"audioStreamCount,omitempty"`
	VideoBitDepth *int32 `json:"videoBitDepth,omitempty"`
	VideoBitrate *int64 `json:"videoBitrate,omitempty"`
	VideoCodec NullableString `json:"videoCodec,omitempty"`
	VideoFps *float64 `json:"videoFps,omitempty"`
	VideoDynamicRange NullableString `json:"videoDynamicRange,omitempty"`
	VideoDynamicRangeType NullableString `json:"videoDynamicRangeType,omitempty"`
	Resolution NullableString `json:"resolution,omitempty"`
	RunTime NullableString `json:"runTime,omitempty"`
	ScanType NullableString `json:"scanType,omitempty"`
	Subtitles NullableString `json:"subtitles,omitempty"`
}

// NewMediaInfoResource instantiates a new MediaInfoResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaInfoResource() *MediaInfoResource {
	this := MediaInfoResource{}
	return &this
}

// NewMediaInfoResourceWithDefaults instantiates a new MediaInfoResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaInfoResourceWithDefaults() *MediaInfoResource {
	this := MediaInfoResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MediaInfoResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MediaInfoResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MediaInfoResource) SetId(v int32) {
	o.Id = &v
}

// GetAudioBitrate returns the AudioBitrate field value if set, zero value otherwise.
func (o *MediaInfoResource) GetAudioBitrate() int64 {
	if o == nil || IsNil(o.AudioBitrate) {
		var ret int64
		return ret
	}
	return *o.AudioBitrate
}

// GetAudioBitrateOk returns a tuple with the AudioBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoResource) GetAudioBitrateOk() (*int64, bool) {
	if o == nil || IsNil(o.AudioBitrate) {
		return nil, false
	}
	return o.AudioBitrate, true
}

// HasAudioBitrate returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioBitrate() bool {
	if o != nil && !IsNil(o.AudioBitrate) {
		return true
	}

	return false
}

// SetAudioBitrate gets a reference to the given int64 and assigns it to the AudioBitrate field.
func (o *MediaInfoResource) SetAudioBitrate(v int64) {
	o.AudioBitrate = &v
}

// GetAudioChannels returns the AudioChannels field value if set, zero value otherwise.
func (o *MediaInfoResource) GetAudioChannels() float64 {
	if o == nil || IsNil(o.AudioChannels) {
		var ret float64
		return ret
	}
	return *o.AudioChannels
}

// GetAudioChannelsOk returns a tuple with the AudioChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoResource) GetAudioChannelsOk() (*float64, bool) {
	if o == nil || IsNil(o.AudioChannels) {
		return nil, false
	}
	return o.AudioChannels, true
}

// HasAudioChannels returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioChannels() bool {
	if o != nil && !IsNil(o.AudioChannels) {
		return true
	}

	return false
}

// SetAudioChannels gets a reference to the given float64 and assigns it to the AudioChannels field.
func (o *MediaInfoResource) SetAudioChannels(v float64) {
	o.AudioChannels = &v
}

// GetAudioCodec returns the AudioCodec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetAudioCodec() string {
	if o == nil || IsNil(o.AudioCodec.Get()) {
		var ret string
		return ret
	}
	return *o.AudioCodec.Get()
}

// GetAudioCodecOk returns a tuple with the AudioCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetAudioCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioCodec.Get(), o.AudioCodec.IsSet()
}

// HasAudioCodec returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioCodec() bool {
	if o != nil && o.AudioCodec.IsSet() {
		return true
	}

	return false
}

// SetAudioCodec gets a reference to the given NullableString and assigns it to the AudioCodec field.
func (o *MediaInfoResource) SetAudioCodec(v string) {
	o.AudioCodec.Set(&v)
}
// SetAudioCodecNil sets the value for AudioCodec to be an explicit nil
func (o *MediaInfoResource) SetAudioCodecNil() {
	o.AudioCodec.Set(nil)
}

// UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
func (o *MediaInfoResource) UnsetAudioCodec() {
	o.AudioCodec.Unset()
}

// GetAudioLanguages returns the AudioLanguages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetAudioLanguages() string {
	if o == nil || IsNil(o.AudioLanguages.Get()) {
		var ret string
		return ret
	}
	return *o.AudioLanguages.Get()
}

// GetAudioLanguagesOk returns a tuple with the AudioLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetAudioLanguagesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioLanguages.Get(), o.AudioLanguages.IsSet()
}

// HasAudioLanguages returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioLanguages() bool {
	if o != nil && o.AudioLanguages.IsSet() {
		return true
	}

	return false
}

// SetAudioLanguages gets a reference to the given NullableString and assigns it to the AudioLanguages field.
func (o *MediaInfoResource) SetAudioLanguages(v string) {
	o.AudioLanguages.Set(&v)
}
// SetAudioLanguagesNil sets the value for AudioLanguages to be an explicit nil
func (o *MediaInfoResource) SetAudioLanguagesNil() {
	o.AudioLanguages.Set(nil)
}

// UnsetAudioLanguages ensures that no value is present for AudioLanguages, not even an explicit nil
func (o *MediaInfoResource) UnsetAudioLanguages() {
	o.AudioLanguages.Unset()
}

// GetAudioStreamCount returns the AudioStreamCount field value if set, zero value otherwise.
func (o *MediaInfoResource) GetAudioStreamCount() int32 {
	if o == nil || IsNil(o.AudioStreamCount) {
		var ret int32
		return ret
	}
	return *o.AudioStreamCount
}

// GetAudioStreamCountOk returns a tuple with the AudioStreamCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoResource) GetAudioStreamCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AudioStreamCount) {
		return nil, false
	}
	return o.AudioStreamCount, true
}

// HasAudioStreamCount returns a boolean if a field has been set.
func (o *MediaInfoResource) HasAudioStreamCount() bool {
	if o != nil && !IsNil(o.AudioStreamCount) {
		return true
	}

	return false
}

// SetAudioStreamCount gets a reference to the given int32 and assigns it to the AudioStreamCount field.
func (o *MediaInfoResource) SetAudioStreamCount(v int32) {
	o.AudioStreamCount = &v
}

// GetVideoBitDepth returns the VideoBitDepth field value if set, zero value otherwise.
func (o *MediaInfoResource) GetVideoBitDepth() int32 {
	if o == nil || IsNil(o.VideoBitDepth) {
		var ret int32
		return ret
	}
	return *o.VideoBitDepth
}

// GetVideoBitDepthOk returns a tuple with the VideoBitDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoResource) GetVideoBitDepthOk() (*int32, bool) {
	if o == nil || IsNil(o.VideoBitDepth) {
		return nil, false
	}
	return o.VideoBitDepth, true
}

// HasVideoBitDepth returns a boolean if a field has been set.
func (o *MediaInfoResource) HasVideoBitDepth() bool {
	if o != nil && !IsNil(o.VideoBitDepth) {
		return true
	}

	return false
}

// SetVideoBitDepth gets a reference to the given int32 and assigns it to the VideoBitDepth field.
func (o *MediaInfoResource) SetVideoBitDepth(v int32) {
	o.VideoBitDepth = &v
}

// GetVideoBitrate returns the VideoBitrate field value if set, zero value otherwise.
func (o *MediaInfoResource) GetVideoBitrate() int64 {
	if o == nil || IsNil(o.VideoBitrate) {
		var ret int64
		return ret
	}
	return *o.VideoBitrate
}

// GetVideoBitrateOk returns a tuple with the VideoBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoResource) GetVideoBitrateOk() (*int64, bool) {
	if o == nil || IsNil(o.VideoBitrate) {
		return nil, false
	}
	return o.VideoBitrate, true
}

// HasVideoBitrate returns a boolean if a field has been set.
func (o *MediaInfoResource) HasVideoBitrate() bool {
	if o != nil && !IsNil(o.VideoBitrate) {
		return true
	}

	return false
}

// SetVideoBitrate gets a reference to the given int64 and assigns it to the VideoBitrate field.
func (o *MediaInfoResource) SetVideoBitrate(v int64) {
	o.VideoBitrate = &v
}

// GetVideoCodec returns the VideoCodec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetVideoCodec() string {
	if o == nil || IsNil(o.VideoCodec.Get()) {
		var ret string
		return ret
	}
	return *o.VideoCodec.Get()
}

// GetVideoCodecOk returns a tuple with the VideoCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetVideoCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoCodec.Get(), o.VideoCodec.IsSet()
}

// HasVideoCodec returns a boolean if a field has been set.
func (o *MediaInfoResource) HasVideoCodec() bool {
	if o != nil && o.VideoCodec.IsSet() {
		return true
	}

	return false
}

// SetVideoCodec gets a reference to the given NullableString and assigns it to the VideoCodec field.
func (o *MediaInfoResource) SetVideoCodec(v string) {
	o.VideoCodec.Set(&v)
}
// SetVideoCodecNil sets the value for VideoCodec to be an explicit nil
func (o *MediaInfoResource) SetVideoCodecNil() {
	o.VideoCodec.Set(nil)
}

// UnsetVideoCodec ensures that no value is present for VideoCodec, not even an explicit nil
func (o *MediaInfoResource) UnsetVideoCodec() {
	o.VideoCodec.Unset()
}

// GetVideoFps returns the VideoFps field value if set, zero value otherwise.
func (o *MediaInfoResource) GetVideoFps() float64 {
	if o == nil || IsNil(o.VideoFps) {
		var ret float64
		return ret
	}
	return *o.VideoFps
}

// GetVideoFpsOk returns a tuple with the VideoFps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaInfoResource) GetVideoFpsOk() (*float64, bool) {
	if o == nil || IsNil(o.VideoFps) {
		return nil, false
	}
	return o.VideoFps, true
}

// HasVideoFps returns a boolean if a field has been set.
func (o *MediaInfoResource) HasVideoFps() bool {
	if o != nil && !IsNil(o.VideoFps) {
		return true
	}

	return false
}

// SetVideoFps gets a reference to the given float64 and assigns it to the VideoFps field.
func (o *MediaInfoResource) SetVideoFps(v float64) {
	o.VideoFps = &v
}

// GetVideoDynamicRange returns the VideoDynamicRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetVideoDynamicRange() string {
	if o == nil || IsNil(o.VideoDynamicRange.Get()) {
		var ret string
		return ret
	}
	return *o.VideoDynamicRange.Get()
}

// GetVideoDynamicRangeOk returns a tuple with the VideoDynamicRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetVideoDynamicRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoDynamicRange.Get(), o.VideoDynamicRange.IsSet()
}

// HasVideoDynamicRange returns a boolean if a field has been set.
func (o *MediaInfoResource) HasVideoDynamicRange() bool {
	if o != nil && o.VideoDynamicRange.IsSet() {
		return true
	}

	return false
}

// SetVideoDynamicRange gets a reference to the given NullableString and assigns it to the VideoDynamicRange field.
func (o *MediaInfoResource) SetVideoDynamicRange(v string) {
	o.VideoDynamicRange.Set(&v)
}
// SetVideoDynamicRangeNil sets the value for VideoDynamicRange to be an explicit nil
func (o *MediaInfoResource) SetVideoDynamicRangeNil() {
	o.VideoDynamicRange.Set(nil)
}

// UnsetVideoDynamicRange ensures that no value is present for VideoDynamicRange, not even an explicit nil
func (o *MediaInfoResource) UnsetVideoDynamicRange() {
	o.VideoDynamicRange.Unset()
}

// GetVideoDynamicRangeType returns the VideoDynamicRangeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetVideoDynamicRangeType() string {
	if o == nil || IsNil(o.VideoDynamicRangeType.Get()) {
		var ret string
		return ret
	}
	return *o.VideoDynamicRangeType.Get()
}

// GetVideoDynamicRangeTypeOk returns a tuple with the VideoDynamicRangeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetVideoDynamicRangeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoDynamicRangeType.Get(), o.VideoDynamicRangeType.IsSet()
}

// HasVideoDynamicRangeType returns a boolean if a field has been set.
func (o *MediaInfoResource) HasVideoDynamicRangeType() bool {
	if o != nil && o.VideoDynamicRangeType.IsSet() {
		return true
	}

	return false
}

// SetVideoDynamicRangeType gets a reference to the given NullableString and assigns it to the VideoDynamicRangeType field.
func (o *MediaInfoResource) SetVideoDynamicRangeType(v string) {
	o.VideoDynamicRangeType.Set(&v)
}
// SetVideoDynamicRangeTypeNil sets the value for VideoDynamicRangeType to be an explicit nil
func (o *MediaInfoResource) SetVideoDynamicRangeTypeNil() {
	o.VideoDynamicRangeType.Set(nil)
}

// UnsetVideoDynamicRangeType ensures that no value is present for VideoDynamicRangeType, not even an explicit nil
func (o *MediaInfoResource) UnsetVideoDynamicRangeType() {
	o.VideoDynamicRangeType.Unset()
}

// GetResolution returns the Resolution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetResolution() string {
	if o == nil || IsNil(o.Resolution.Get()) {
		var ret string
		return ret
	}
	return *o.Resolution.Get()
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetResolutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolution.Get(), o.Resolution.IsSet()
}

// HasResolution returns a boolean if a field has been set.
func (o *MediaInfoResource) HasResolution() bool {
	if o != nil && o.Resolution.IsSet() {
		return true
	}

	return false
}

// SetResolution gets a reference to the given NullableString and assigns it to the Resolution field.
func (o *MediaInfoResource) SetResolution(v string) {
	o.Resolution.Set(&v)
}
// SetResolutionNil sets the value for Resolution to be an explicit nil
func (o *MediaInfoResource) SetResolutionNil() {
	o.Resolution.Set(nil)
}

// UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
func (o *MediaInfoResource) UnsetResolution() {
	o.Resolution.Unset()
}

// GetRunTime returns the RunTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetRunTime() string {
	if o == nil || IsNil(o.RunTime.Get()) {
		var ret string
		return ret
	}
	return *o.RunTime.Get()
}

// GetRunTimeOk returns a tuple with the RunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetRunTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunTime.Get(), o.RunTime.IsSet()
}

// HasRunTime returns a boolean if a field has been set.
func (o *MediaInfoResource) HasRunTime() bool {
	if o != nil && o.RunTime.IsSet() {
		return true
	}

	return false
}

// SetRunTime gets a reference to the given NullableString and assigns it to the RunTime field.
func (o *MediaInfoResource) SetRunTime(v string) {
	o.RunTime.Set(&v)
}
// SetRunTimeNil sets the value for RunTime to be an explicit nil
func (o *MediaInfoResource) SetRunTimeNil() {
	o.RunTime.Set(nil)
}

// UnsetRunTime ensures that no value is present for RunTime, not even an explicit nil
func (o *MediaInfoResource) UnsetRunTime() {
	o.RunTime.Unset()
}

// GetScanType returns the ScanType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetScanType() string {
	if o == nil || IsNil(o.ScanType.Get()) {
		var ret string
		return ret
	}
	return *o.ScanType.Get()
}

// GetScanTypeOk returns a tuple with the ScanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScanType.Get(), o.ScanType.IsSet()
}

// HasScanType returns a boolean if a field has been set.
func (o *MediaInfoResource) HasScanType() bool {
	if o != nil && o.ScanType.IsSet() {
		return true
	}

	return false
}

// SetScanType gets a reference to the given NullableString and assigns it to the ScanType field.
func (o *MediaInfoResource) SetScanType(v string) {
	o.ScanType.Set(&v)
}
// SetScanTypeNil sets the value for ScanType to be an explicit nil
func (o *MediaInfoResource) SetScanTypeNil() {
	o.ScanType.Set(nil)
}

// UnsetScanType ensures that no value is present for ScanType, not even an explicit nil
func (o *MediaInfoResource) UnsetScanType() {
	o.ScanType.Unset()
}

// GetSubtitles returns the Subtitles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaInfoResource) GetSubtitles() string {
	if o == nil || IsNil(o.Subtitles.Get()) {
		var ret string
		return ret
	}
	return *o.Subtitles.Get()
}

// GetSubtitlesOk returns a tuple with the Subtitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaInfoResource) GetSubtitlesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subtitles.Get(), o.Subtitles.IsSet()
}

// HasSubtitles returns a boolean if a field has been set.
func (o *MediaInfoResource) HasSubtitles() bool {
	if o != nil && o.Subtitles.IsSet() {
		return true
	}

	return false
}

// SetSubtitles gets a reference to the given NullableString and assigns it to the Subtitles field.
func (o *MediaInfoResource) SetSubtitles(v string) {
	o.Subtitles.Set(&v)
}
// SetSubtitlesNil sets the value for Subtitles to be an explicit nil
func (o *MediaInfoResource) SetSubtitlesNil() {
	o.Subtitles.Set(nil)
}

// UnsetSubtitles ensures that no value is present for Subtitles, not even an explicit nil
func (o *MediaInfoResource) UnsetSubtitles() {
	o.Subtitles.Unset()
}

func (o MediaInfoResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaInfoResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AudioBitrate) {
		toSerialize["audioBitrate"] = o.AudioBitrate
	}
	if !IsNil(o.AudioChannels) {
		toSerialize["audioChannels"] = o.AudioChannels
	}
	if o.AudioCodec.IsSet() {
		toSerialize["audioCodec"] = o.AudioCodec.Get()
	}
	if o.AudioLanguages.IsSet() {
		toSerialize["audioLanguages"] = o.AudioLanguages.Get()
	}
	if !IsNil(o.AudioStreamCount) {
		toSerialize["audioStreamCount"] = o.AudioStreamCount
	}
	if !IsNil(o.VideoBitDepth) {
		toSerialize["videoBitDepth"] = o.VideoBitDepth
	}
	if !IsNil(o.VideoBitrate) {
		toSerialize["videoBitrate"] = o.VideoBitrate
	}
	if o.VideoCodec.IsSet() {
		toSerialize["videoCodec"] = o.VideoCodec.Get()
	}
	if !IsNil(o.VideoFps) {
		toSerialize["videoFps"] = o.VideoFps
	}
	if o.VideoDynamicRange.IsSet() {
		toSerialize["videoDynamicRange"] = o.VideoDynamicRange.Get()
	}
	if o.VideoDynamicRangeType.IsSet() {
		toSerialize["videoDynamicRangeType"] = o.VideoDynamicRangeType.Get()
	}
	if o.Resolution.IsSet() {
		toSerialize["resolution"] = o.Resolution.Get()
	}
	if o.RunTime.IsSet() {
		toSerialize["runTime"] = o.RunTime.Get()
	}
	if o.ScanType.IsSet() {
		toSerialize["scanType"] = o.ScanType.Get()
	}
	if o.Subtitles.IsSet() {
		toSerialize["subtitles"] = o.Subtitles.Get()
	}
	return toSerialize, nil
}

type NullableMediaInfoResource struct {
	value *MediaInfoResource
	isSet bool
}

func (v NullableMediaInfoResource) Get() *MediaInfoResource {
	return v.value
}

func (v *NullableMediaInfoResource) Set(val *MediaInfoResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaInfoResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaInfoResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaInfoResource(val *MediaInfoResource) *NullableMediaInfoResource {
	return &NullableMediaInfoResource{value: val, isSet: true}
}

func (v NullableMediaInfoResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaInfoResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



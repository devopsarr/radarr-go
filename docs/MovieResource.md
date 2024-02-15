# MovieResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**OriginalTitle** | Pointer to **NullableString** |  | [optional] 
**OriginalLanguage** | Pointer to [**Language**](Language.md) |  | [optional] 
**AlternateTitles** | Pointer to [**[]AlternativeTitleResource**](AlternativeTitleResource.md) |  | [optional] 
**SecondaryYear** | Pointer to **NullableInt32** |  | [optional] 
**SecondaryYearSourceId** | Pointer to **int32** |  | [optional] 
**SortTitle** | Pointer to **NullableString** |  | [optional] 
**SizeOnDisk** | Pointer to **NullableInt64** |  | [optional] 
**Status** | Pointer to [**MovieStatusType**](MovieStatusType.md) |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**InCinemas** | Pointer to **NullableTime** |  | [optional] 
**PhysicalRelease** | Pointer to **NullableTime** |  | [optional] 
**DigitalRelease** | Pointer to **NullableTime** |  | [optional] 
**PhysicalReleaseNote** | Pointer to **NullableString** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**RemotePoster** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**HasFile** | Pointer to **bool** |  | [optional] 
**YouTubeTrailerId** | Pointer to **NullableString** |  | [optional] 
**Studio** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**QualityProfileId** | Pointer to **int32** |  | [optional] 
**Monitored** | Pointer to **bool** |  | [optional] 
**MinimumAvailability** | Pointer to [**MovieStatusType**](MovieStatusType.md) |  | [optional] 
**IsAvailable** | Pointer to **bool** |  | [optional] 
**FolderName** | Pointer to **NullableString** |  | [optional] 
**Runtime** | Pointer to **int32** |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**ImdbId** | Pointer to **NullableString** |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 
**TitleSlug** | Pointer to **NullableString** |  | [optional] 
**RootFolderPath** | Pointer to **NullableString** |  | [optional] 
**Folder** | Pointer to **NullableString** |  | [optional] 
**Certification** | Pointer to **NullableString** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Added** | Pointer to **time.Time** |  | [optional] 
**AddOptions** | Pointer to [**AddMovieOptions**](AddMovieOptions.md) |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**MovieFile** | Pointer to [**MovieFileResource**](MovieFileResource.md) |  | [optional] 
**Collection** | Pointer to [**MovieCollection**](MovieCollection.md) |  | [optional] 
**Popularity** | Pointer to **float32** |  | [optional] 

## Methods

### NewMovieResource

`func NewMovieResource() *MovieResource`

NewMovieResource instantiates a new MovieResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieResourceWithDefaults

`func NewMovieResourceWithDefaults() *MovieResource`

NewMovieResourceWithDefaults instantiates a new MovieResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MovieResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MovieResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MovieResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MovieResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *MovieResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MovieResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MovieResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MovieResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MovieResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MovieResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetOriginalTitle

`func (o *MovieResource) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *MovieResource) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *MovieResource) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *MovieResource) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *MovieResource) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *MovieResource) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetOriginalLanguage

`func (o *MovieResource) GetOriginalLanguage() Language`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *MovieResource) GetOriginalLanguageOk() (*Language, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *MovieResource) SetOriginalLanguage(v Language)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *MovieResource) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetAlternateTitles

`func (o *MovieResource) GetAlternateTitles() []AlternativeTitleResource`

GetAlternateTitles returns the AlternateTitles field if non-nil, zero value otherwise.

### GetAlternateTitlesOk

`func (o *MovieResource) GetAlternateTitlesOk() (*[]AlternativeTitleResource, bool)`

GetAlternateTitlesOk returns a tuple with the AlternateTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateTitles

`func (o *MovieResource) SetAlternateTitles(v []AlternativeTitleResource)`

SetAlternateTitles sets AlternateTitles field to given value.

### HasAlternateTitles

`func (o *MovieResource) HasAlternateTitles() bool`

HasAlternateTitles returns a boolean if a field has been set.

### SetAlternateTitlesNil

`func (o *MovieResource) SetAlternateTitlesNil(b bool)`

 SetAlternateTitlesNil sets the value for AlternateTitles to be an explicit nil

### UnsetAlternateTitles
`func (o *MovieResource) UnsetAlternateTitles()`

UnsetAlternateTitles ensures that no value is present for AlternateTitles, not even an explicit nil
### GetSecondaryYear

`func (o *MovieResource) GetSecondaryYear() int32`

GetSecondaryYear returns the SecondaryYear field if non-nil, zero value otherwise.

### GetSecondaryYearOk

`func (o *MovieResource) GetSecondaryYearOk() (*int32, bool)`

GetSecondaryYearOk returns a tuple with the SecondaryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryYear

`func (o *MovieResource) SetSecondaryYear(v int32)`

SetSecondaryYear sets SecondaryYear field to given value.

### HasSecondaryYear

`func (o *MovieResource) HasSecondaryYear() bool`

HasSecondaryYear returns a boolean if a field has been set.

### SetSecondaryYearNil

`func (o *MovieResource) SetSecondaryYearNil(b bool)`

 SetSecondaryYearNil sets the value for SecondaryYear to be an explicit nil

### UnsetSecondaryYear
`func (o *MovieResource) UnsetSecondaryYear()`

UnsetSecondaryYear ensures that no value is present for SecondaryYear, not even an explicit nil
### GetSecondaryYearSourceId

`func (o *MovieResource) GetSecondaryYearSourceId() int32`

GetSecondaryYearSourceId returns the SecondaryYearSourceId field if non-nil, zero value otherwise.

### GetSecondaryYearSourceIdOk

`func (o *MovieResource) GetSecondaryYearSourceIdOk() (*int32, bool)`

GetSecondaryYearSourceIdOk returns a tuple with the SecondaryYearSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryYearSourceId

`func (o *MovieResource) SetSecondaryYearSourceId(v int32)`

SetSecondaryYearSourceId sets SecondaryYearSourceId field to given value.

### HasSecondaryYearSourceId

`func (o *MovieResource) HasSecondaryYearSourceId() bool`

HasSecondaryYearSourceId returns a boolean if a field has been set.

### GetSortTitle

`func (o *MovieResource) GetSortTitle() string`

GetSortTitle returns the SortTitle field if non-nil, zero value otherwise.

### GetSortTitleOk

`func (o *MovieResource) GetSortTitleOk() (*string, bool)`

GetSortTitleOk returns a tuple with the SortTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortTitle

`func (o *MovieResource) SetSortTitle(v string)`

SetSortTitle sets SortTitle field to given value.

### HasSortTitle

`func (o *MovieResource) HasSortTitle() bool`

HasSortTitle returns a boolean if a field has been set.

### SetSortTitleNil

`func (o *MovieResource) SetSortTitleNil(b bool)`

 SetSortTitleNil sets the value for SortTitle to be an explicit nil

### UnsetSortTitle
`func (o *MovieResource) UnsetSortTitle()`

UnsetSortTitle ensures that no value is present for SortTitle, not even an explicit nil
### GetSizeOnDisk

`func (o *MovieResource) GetSizeOnDisk() int64`

GetSizeOnDisk returns the SizeOnDisk field if non-nil, zero value otherwise.

### GetSizeOnDiskOk

`func (o *MovieResource) GetSizeOnDiskOk() (*int64, bool)`

GetSizeOnDiskOk returns a tuple with the SizeOnDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeOnDisk

`func (o *MovieResource) SetSizeOnDisk(v int64)`

SetSizeOnDisk sets SizeOnDisk field to given value.

### HasSizeOnDisk

`func (o *MovieResource) HasSizeOnDisk() bool`

HasSizeOnDisk returns a boolean if a field has been set.

### SetSizeOnDiskNil

`func (o *MovieResource) SetSizeOnDiskNil(b bool)`

 SetSizeOnDiskNil sets the value for SizeOnDisk to be an explicit nil

### UnsetSizeOnDisk
`func (o *MovieResource) UnsetSizeOnDisk()`

UnsetSizeOnDisk ensures that no value is present for SizeOnDisk, not even an explicit nil
### GetStatus

`func (o *MovieResource) GetStatus() MovieStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MovieResource) GetStatusOk() (*MovieStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MovieResource) SetStatus(v MovieStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MovieResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOverview

`func (o *MovieResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *MovieResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *MovieResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *MovieResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *MovieResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *MovieResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetInCinemas

`func (o *MovieResource) GetInCinemas() time.Time`

GetInCinemas returns the InCinemas field if non-nil, zero value otherwise.

### GetInCinemasOk

`func (o *MovieResource) GetInCinemasOk() (*time.Time, bool)`

GetInCinemasOk returns a tuple with the InCinemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInCinemas

`func (o *MovieResource) SetInCinemas(v time.Time)`

SetInCinemas sets InCinemas field to given value.

### HasInCinemas

`func (o *MovieResource) HasInCinemas() bool`

HasInCinemas returns a boolean if a field has been set.

### SetInCinemasNil

`func (o *MovieResource) SetInCinemasNil(b bool)`

 SetInCinemasNil sets the value for InCinemas to be an explicit nil

### UnsetInCinemas
`func (o *MovieResource) UnsetInCinemas()`

UnsetInCinemas ensures that no value is present for InCinemas, not even an explicit nil
### GetPhysicalRelease

`func (o *MovieResource) GetPhysicalRelease() time.Time`

GetPhysicalRelease returns the PhysicalRelease field if non-nil, zero value otherwise.

### GetPhysicalReleaseOk

`func (o *MovieResource) GetPhysicalReleaseOk() (*time.Time, bool)`

GetPhysicalReleaseOk returns a tuple with the PhysicalRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalRelease

`func (o *MovieResource) SetPhysicalRelease(v time.Time)`

SetPhysicalRelease sets PhysicalRelease field to given value.

### HasPhysicalRelease

`func (o *MovieResource) HasPhysicalRelease() bool`

HasPhysicalRelease returns a boolean if a field has been set.

### SetPhysicalReleaseNil

`func (o *MovieResource) SetPhysicalReleaseNil(b bool)`

 SetPhysicalReleaseNil sets the value for PhysicalRelease to be an explicit nil

### UnsetPhysicalRelease
`func (o *MovieResource) UnsetPhysicalRelease()`

UnsetPhysicalRelease ensures that no value is present for PhysicalRelease, not even an explicit nil
### GetDigitalRelease

`func (o *MovieResource) GetDigitalRelease() time.Time`

GetDigitalRelease returns the DigitalRelease field if non-nil, zero value otherwise.

### GetDigitalReleaseOk

`func (o *MovieResource) GetDigitalReleaseOk() (*time.Time, bool)`

GetDigitalReleaseOk returns a tuple with the DigitalRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalRelease

`func (o *MovieResource) SetDigitalRelease(v time.Time)`

SetDigitalRelease sets DigitalRelease field to given value.

### HasDigitalRelease

`func (o *MovieResource) HasDigitalRelease() bool`

HasDigitalRelease returns a boolean if a field has been set.

### SetDigitalReleaseNil

`func (o *MovieResource) SetDigitalReleaseNil(b bool)`

 SetDigitalReleaseNil sets the value for DigitalRelease to be an explicit nil

### UnsetDigitalRelease
`func (o *MovieResource) UnsetDigitalRelease()`

UnsetDigitalRelease ensures that no value is present for DigitalRelease, not even an explicit nil
### GetPhysicalReleaseNote

`func (o *MovieResource) GetPhysicalReleaseNote() string`

GetPhysicalReleaseNote returns the PhysicalReleaseNote field if non-nil, zero value otherwise.

### GetPhysicalReleaseNoteOk

`func (o *MovieResource) GetPhysicalReleaseNoteOk() (*string, bool)`

GetPhysicalReleaseNoteOk returns a tuple with the PhysicalReleaseNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalReleaseNote

`func (o *MovieResource) SetPhysicalReleaseNote(v string)`

SetPhysicalReleaseNote sets PhysicalReleaseNote field to given value.

### HasPhysicalReleaseNote

`func (o *MovieResource) HasPhysicalReleaseNote() bool`

HasPhysicalReleaseNote returns a boolean if a field has been set.

### SetPhysicalReleaseNoteNil

`func (o *MovieResource) SetPhysicalReleaseNoteNil(b bool)`

 SetPhysicalReleaseNoteNil sets the value for PhysicalReleaseNote to be an explicit nil

### UnsetPhysicalReleaseNote
`func (o *MovieResource) UnsetPhysicalReleaseNote()`

UnsetPhysicalReleaseNote ensures that no value is present for PhysicalReleaseNote, not even an explicit nil
### GetImages

`func (o *MovieResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *MovieResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *MovieResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *MovieResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *MovieResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *MovieResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetWebsite

`func (o *MovieResource) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *MovieResource) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *MovieResource) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *MovieResource) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *MovieResource) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *MovieResource) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetRemotePoster

`func (o *MovieResource) GetRemotePoster() string`

GetRemotePoster returns the RemotePoster field if non-nil, zero value otherwise.

### GetRemotePosterOk

`func (o *MovieResource) GetRemotePosterOk() (*string, bool)`

GetRemotePosterOk returns a tuple with the RemotePoster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePoster

`func (o *MovieResource) SetRemotePoster(v string)`

SetRemotePoster sets RemotePoster field to given value.

### HasRemotePoster

`func (o *MovieResource) HasRemotePoster() bool`

HasRemotePoster returns a boolean if a field has been set.

### SetRemotePosterNil

`func (o *MovieResource) SetRemotePosterNil(b bool)`

 SetRemotePosterNil sets the value for RemotePoster to be an explicit nil

### UnsetRemotePoster
`func (o *MovieResource) UnsetRemotePoster()`

UnsetRemotePoster ensures that no value is present for RemotePoster, not even an explicit nil
### GetYear

`func (o *MovieResource) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MovieResource) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MovieResource) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *MovieResource) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetHasFile

`func (o *MovieResource) GetHasFile() bool`

GetHasFile returns the HasFile field if non-nil, zero value otherwise.

### GetHasFileOk

`func (o *MovieResource) GetHasFileOk() (*bool, bool)`

GetHasFileOk returns a tuple with the HasFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFile

`func (o *MovieResource) SetHasFile(v bool)`

SetHasFile sets HasFile field to given value.

### HasHasFile

`func (o *MovieResource) HasHasFile() bool`

HasHasFile returns a boolean if a field has been set.

### GetYouTubeTrailerId

`func (o *MovieResource) GetYouTubeTrailerId() string`

GetYouTubeTrailerId returns the YouTubeTrailerId field if non-nil, zero value otherwise.

### GetYouTubeTrailerIdOk

`func (o *MovieResource) GetYouTubeTrailerIdOk() (*string, bool)`

GetYouTubeTrailerIdOk returns a tuple with the YouTubeTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeTrailerId

`func (o *MovieResource) SetYouTubeTrailerId(v string)`

SetYouTubeTrailerId sets YouTubeTrailerId field to given value.

### HasYouTubeTrailerId

`func (o *MovieResource) HasYouTubeTrailerId() bool`

HasYouTubeTrailerId returns a boolean if a field has been set.

### SetYouTubeTrailerIdNil

`func (o *MovieResource) SetYouTubeTrailerIdNil(b bool)`

 SetYouTubeTrailerIdNil sets the value for YouTubeTrailerId to be an explicit nil

### UnsetYouTubeTrailerId
`func (o *MovieResource) UnsetYouTubeTrailerId()`

UnsetYouTubeTrailerId ensures that no value is present for YouTubeTrailerId, not even an explicit nil
### GetStudio

`func (o *MovieResource) GetStudio() string`

GetStudio returns the Studio field if non-nil, zero value otherwise.

### GetStudioOk

`func (o *MovieResource) GetStudioOk() (*string, bool)`

GetStudioOk returns a tuple with the Studio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudio

`func (o *MovieResource) SetStudio(v string)`

SetStudio sets Studio field to given value.

### HasStudio

`func (o *MovieResource) HasStudio() bool`

HasStudio returns a boolean if a field has been set.

### SetStudioNil

`func (o *MovieResource) SetStudioNil(b bool)`

 SetStudioNil sets the value for Studio to be an explicit nil

### UnsetStudio
`func (o *MovieResource) UnsetStudio()`

UnsetStudio ensures that no value is present for Studio, not even an explicit nil
### GetPath

`func (o *MovieResource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MovieResource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MovieResource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MovieResource) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MovieResource) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MovieResource) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetQualityProfileId

`func (o *MovieResource) GetQualityProfileId() int32`

GetQualityProfileId returns the QualityProfileId field if non-nil, zero value otherwise.

### GetQualityProfileIdOk

`func (o *MovieResource) GetQualityProfileIdOk() (*int32, bool)`

GetQualityProfileIdOk returns a tuple with the QualityProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityProfileId

`func (o *MovieResource) SetQualityProfileId(v int32)`

SetQualityProfileId sets QualityProfileId field to given value.

### HasQualityProfileId

`func (o *MovieResource) HasQualityProfileId() bool`

HasQualityProfileId returns a boolean if a field has been set.

### GetMonitored

`func (o *MovieResource) GetMonitored() bool`

GetMonitored returns the Monitored field if non-nil, zero value otherwise.

### GetMonitoredOk

`func (o *MovieResource) GetMonitoredOk() (*bool, bool)`

GetMonitoredOk returns a tuple with the Monitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitored

`func (o *MovieResource) SetMonitored(v bool)`

SetMonitored sets Monitored field to given value.

### HasMonitored

`func (o *MovieResource) HasMonitored() bool`

HasMonitored returns a boolean if a field has been set.

### GetMinimumAvailability

`func (o *MovieResource) GetMinimumAvailability() MovieStatusType`

GetMinimumAvailability returns the MinimumAvailability field if non-nil, zero value otherwise.

### GetMinimumAvailabilityOk

`func (o *MovieResource) GetMinimumAvailabilityOk() (*MovieStatusType, bool)`

GetMinimumAvailabilityOk returns a tuple with the MinimumAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAvailability

`func (o *MovieResource) SetMinimumAvailability(v MovieStatusType)`

SetMinimumAvailability sets MinimumAvailability field to given value.

### HasMinimumAvailability

`func (o *MovieResource) HasMinimumAvailability() bool`

HasMinimumAvailability returns a boolean if a field has been set.

### GetIsAvailable

`func (o *MovieResource) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *MovieResource) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *MovieResource) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.

### HasIsAvailable

`func (o *MovieResource) HasIsAvailable() bool`

HasIsAvailable returns a boolean if a field has been set.

### GetFolderName

`func (o *MovieResource) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *MovieResource) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *MovieResource) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *MovieResource) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### SetFolderNameNil

`func (o *MovieResource) SetFolderNameNil(b bool)`

 SetFolderNameNil sets the value for FolderName to be an explicit nil

### UnsetFolderName
`func (o *MovieResource) UnsetFolderName()`

UnsetFolderName ensures that no value is present for FolderName, not even an explicit nil
### GetRuntime

`func (o *MovieResource) GetRuntime() int32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *MovieResource) GetRuntimeOk() (*int32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *MovieResource) SetRuntime(v int32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *MovieResource) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetCleanTitle

`func (o *MovieResource) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *MovieResource) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *MovieResource) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *MovieResource) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *MovieResource) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *MovieResource) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetImdbId

`func (o *MovieResource) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *MovieResource) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *MovieResource) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *MovieResource) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### SetImdbIdNil

`func (o *MovieResource) SetImdbIdNil(b bool)`

 SetImdbIdNil sets the value for ImdbId to be an explicit nil

### UnsetImdbId
`func (o *MovieResource) UnsetImdbId()`

UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
### GetTmdbId

`func (o *MovieResource) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *MovieResource) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *MovieResource) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *MovieResource) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetTitleSlug

`func (o *MovieResource) GetTitleSlug() string`

GetTitleSlug returns the TitleSlug field if non-nil, zero value otherwise.

### GetTitleSlugOk

`func (o *MovieResource) GetTitleSlugOk() (*string, bool)`

GetTitleSlugOk returns a tuple with the TitleSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSlug

`func (o *MovieResource) SetTitleSlug(v string)`

SetTitleSlug sets TitleSlug field to given value.

### HasTitleSlug

`func (o *MovieResource) HasTitleSlug() bool`

HasTitleSlug returns a boolean if a field has been set.

### SetTitleSlugNil

`func (o *MovieResource) SetTitleSlugNil(b bool)`

 SetTitleSlugNil sets the value for TitleSlug to be an explicit nil

### UnsetTitleSlug
`func (o *MovieResource) UnsetTitleSlug()`

UnsetTitleSlug ensures that no value is present for TitleSlug, not even an explicit nil
### GetRootFolderPath

`func (o *MovieResource) GetRootFolderPath() string`

GetRootFolderPath returns the RootFolderPath field if non-nil, zero value otherwise.

### GetRootFolderPathOk

`func (o *MovieResource) GetRootFolderPathOk() (*string, bool)`

GetRootFolderPathOk returns a tuple with the RootFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolderPath

`func (o *MovieResource) SetRootFolderPath(v string)`

SetRootFolderPath sets RootFolderPath field to given value.

### HasRootFolderPath

`func (o *MovieResource) HasRootFolderPath() bool`

HasRootFolderPath returns a boolean if a field has been set.

### SetRootFolderPathNil

`func (o *MovieResource) SetRootFolderPathNil(b bool)`

 SetRootFolderPathNil sets the value for RootFolderPath to be an explicit nil

### UnsetRootFolderPath
`func (o *MovieResource) UnsetRootFolderPath()`

UnsetRootFolderPath ensures that no value is present for RootFolderPath, not even an explicit nil
### GetFolder

`func (o *MovieResource) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *MovieResource) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *MovieResource) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *MovieResource) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *MovieResource) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *MovieResource) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetCertification

`func (o *MovieResource) GetCertification() string`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *MovieResource) GetCertificationOk() (*string, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *MovieResource) SetCertification(v string)`

SetCertification sets Certification field to given value.

### HasCertification

`func (o *MovieResource) HasCertification() bool`

HasCertification returns a boolean if a field has been set.

### SetCertificationNil

`func (o *MovieResource) SetCertificationNil(b bool)`

 SetCertificationNil sets the value for Certification to be an explicit nil

### UnsetCertification
`func (o *MovieResource) UnsetCertification()`

UnsetCertification ensures that no value is present for Certification, not even an explicit nil
### GetGenres

`func (o *MovieResource) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *MovieResource) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *MovieResource) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *MovieResource) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *MovieResource) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *MovieResource) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetTags

`func (o *MovieResource) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MovieResource) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MovieResource) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MovieResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MovieResource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MovieResource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAdded

`func (o *MovieResource) GetAdded() time.Time`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *MovieResource) GetAddedOk() (*time.Time, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *MovieResource) SetAdded(v time.Time)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *MovieResource) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetAddOptions

`func (o *MovieResource) GetAddOptions() AddMovieOptions`

GetAddOptions returns the AddOptions field if non-nil, zero value otherwise.

### GetAddOptionsOk

`func (o *MovieResource) GetAddOptionsOk() (*AddMovieOptions, bool)`

GetAddOptionsOk returns a tuple with the AddOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOptions

`func (o *MovieResource) SetAddOptions(v AddMovieOptions)`

SetAddOptions sets AddOptions field to given value.

### HasAddOptions

`func (o *MovieResource) HasAddOptions() bool`

HasAddOptions returns a boolean if a field has been set.

### GetRatings

`func (o *MovieResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *MovieResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *MovieResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *MovieResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetMovieFile

`func (o *MovieResource) GetMovieFile() MovieFileResource`

GetMovieFile returns the MovieFile field if non-nil, zero value otherwise.

### GetMovieFileOk

`func (o *MovieResource) GetMovieFileOk() (*MovieFileResource, bool)`

GetMovieFileOk returns a tuple with the MovieFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieFile

`func (o *MovieResource) SetMovieFile(v MovieFileResource)`

SetMovieFile sets MovieFile field to given value.

### HasMovieFile

`func (o *MovieResource) HasMovieFile() bool`

HasMovieFile returns a boolean if a field has been set.

### GetCollection

`func (o *MovieResource) GetCollection() MovieCollection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *MovieResource) GetCollectionOk() (*MovieCollection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *MovieResource) SetCollection(v MovieCollection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *MovieResource) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetPopularity

`func (o *MovieResource) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *MovieResource) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *MovieResource) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *MovieResource) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



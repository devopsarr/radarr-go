# MovieMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**InCinemas** | Pointer to **NullableTime** |  | [optional] 
**PhysicalRelease** | Pointer to **NullableTime** |  | [optional] 
**DigitalRelease** | Pointer to **NullableTime** |  | [optional] 
**Certification** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**CollectionTmdbId** | Pointer to **int32** |  | [optional] 
**CollectionTitle** | Pointer to **NullableString** |  | [optional] 
**LastInfoSync** | Pointer to **NullableTime** |  | [optional] 
**Runtime** | Pointer to **int32** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**ImdbId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**CleanTitle** | Pointer to **NullableString** |  | [optional] 
**SortTitle** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**MovieStatusType**](MovieStatusType.md) |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**AlternativeTitles** | Pointer to [**[]AlternativeTitle**](AlternativeTitle.md) |  | [optional] 
**Translations** | Pointer to [**[]MovieTranslation**](MovieTranslation.md) |  | [optional] 
**SecondaryYear** | Pointer to **NullableInt32** |  | [optional] 
**YouTubeTrailerId** | Pointer to **NullableString** |  | [optional] 
**Studio** | Pointer to **NullableString** |  | [optional] 
**OriginalTitle** | Pointer to **NullableString** |  | [optional] 
**CleanOriginalTitle** | Pointer to **NullableString** |  | [optional] 
**OriginalLanguage** | Pointer to [**Language**](Language.md) |  | [optional] 
**Recommendations** | Pointer to **[]int32** |  | [optional] 
**Popularity** | Pointer to **float32** |  | [optional] 
**IsRecentMovie** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewMovieMetadata

`func NewMovieMetadata() *MovieMetadata`

NewMovieMetadata instantiates a new MovieMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieMetadataWithDefaults

`func NewMovieMetadataWithDefaults() *MovieMetadata`

NewMovieMetadataWithDefaults instantiates a new MovieMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MovieMetadata) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MovieMetadata) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MovieMetadata) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MovieMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTmdbId

`func (o *MovieMetadata) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *MovieMetadata) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *MovieMetadata) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *MovieMetadata) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetImages

`func (o *MovieMetadata) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *MovieMetadata) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *MovieMetadata) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *MovieMetadata) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *MovieMetadata) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *MovieMetadata) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetGenres

`func (o *MovieMetadata) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *MovieMetadata) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *MovieMetadata) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *MovieMetadata) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *MovieMetadata) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *MovieMetadata) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetInCinemas

`func (o *MovieMetadata) GetInCinemas() time.Time`

GetInCinemas returns the InCinemas field if non-nil, zero value otherwise.

### GetInCinemasOk

`func (o *MovieMetadata) GetInCinemasOk() (*time.Time, bool)`

GetInCinemasOk returns a tuple with the InCinemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInCinemas

`func (o *MovieMetadata) SetInCinemas(v time.Time)`

SetInCinemas sets InCinemas field to given value.

### HasInCinemas

`func (o *MovieMetadata) HasInCinemas() bool`

HasInCinemas returns a boolean if a field has been set.

### SetInCinemasNil

`func (o *MovieMetadata) SetInCinemasNil(b bool)`

 SetInCinemasNil sets the value for InCinemas to be an explicit nil

### UnsetInCinemas
`func (o *MovieMetadata) UnsetInCinemas()`

UnsetInCinemas ensures that no value is present for InCinemas, not even an explicit nil
### GetPhysicalRelease

`func (o *MovieMetadata) GetPhysicalRelease() time.Time`

GetPhysicalRelease returns the PhysicalRelease field if non-nil, zero value otherwise.

### GetPhysicalReleaseOk

`func (o *MovieMetadata) GetPhysicalReleaseOk() (*time.Time, bool)`

GetPhysicalReleaseOk returns a tuple with the PhysicalRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalRelease

`func (o *MovieMetadata) SetPhysicalRelease(v time.Time)`

SetPhysicalRelease sets PhysicalRelease field to given value.

### HasPhysicalRelease

`func (o *MovieMetadata) HasPhysicalRelease() bool`

HasPhysicalRelease returns a boolean if a field has been set.

### SetPhysicalReleaseNil

`func (o *MovieMetadata) SetPhysicalReleaseNil(b bool)`

 SetPhysicalReleaseNil sets the value for PhysicalRelease to be an explicit nil

### UnsetPhysicalRelease
`func (o *MovieMetadata) UnsetPhysicalRelease()`

UnsetPhysicalRelease ensures that no value is present for PhysicalRelease, not even an explicit nil
### GetDigitalRelease

`func (o *MovieMetadata) GetDigitalRelease() time.Time`

GetDigitalRelease returns the DigitalRelease field if non-nil, zero value otherwise.

### GetDigitalReleaseOk

`func (o *MovieMetadata) GetDigitalReleaseOk() (*time.Time, bool)`

GetDigitalReleaseOk returns a tuple with the DigitalRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalRelease

`func (o *MovieMetadata) SetDigitalRelease(v time.Time)`

SetDigitalRelease sets DigitalRelease field to given value.

### HasDigitalRelease

`func (o *MovieMetadata) HasDigitalRelease() bool`

HasDigitalRelease returns a boolean if a field has been set.

### SetDigitalReleaseNil

`func (o *MovieMetadata) SetDigitalReleaseNil(b bool)`

 SetDigitalReleaseNil sets the value for DigitalRelease to be an explicit nil

### UnsetDigitalRelease
`func (o *MovieMetadata) UnsetDigitalRelease()`

UnsetDigitalRelease ensures that no value is present for DigitalRelease, not even an explicit nil
### GetCertification

`func (o *MovieMetadata) GetCertification() string`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *MovieMetadata) GetCertificationOk() (*string, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *MovieMetadata) SetCertification(v string)`

SetCertification sets Certification field to given value.

### HasCertification

`func (o *MovieMetadata) HasCertification() bool`

HasCertification returns a boolean if a field has been set.

### SetCertificationNil

`func (o *MovieMetadata) SetCertificationNil(b bool)`

 SetCertificationNil sets the value for Certification to be an explicit nil

### UnsetCertification
`func (o *MovieMetadata) UnsetCertification()`

UnsetCertification ensures that no value is present for Certification, not even an explicit nil
### GetYear

`func (o *MovieMetadata) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MovieMetadata) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MovieMetadata) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *MovieMetadata) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetRatings

`func (o *MovieMetadata) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *MovieMetadata) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *MovieMetadata) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *MovieMetadata) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetCollectionTmdbId

`func (o *MovieMetadata) GetCollectionTmdbId() int32`

GetCollectionTmdbId returns the CollectionTmdbId field if non-nil, zero value otherwise.

### GetCollectionTmdbIdOk

`func (o *MovieMetadata) GetCollectionTmdbIdOk() (*int32, bool)`

GetCollectionTmdbIdOk returns a tuple with the CollectionTmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTmdbId

`func (o *MovieMetadata) SetCollectionTmdbId(v int32)`

SetCollectionTmdbId sets CollectionTmdbId field to given value.

### HasCollectionTmdbId

`func (o *MovieMetadata) HasCollectionTmdbId() bool`

HasCollectionTmdbId returns a boolean if a field has been set.

### GetCollectionTitle

`func (o *MovieMetadata) GetCollectionTitle() string`

GetCollectionTitle returns the CollectionTitle field if non-nil, zero value otherwise.

### GetCollectionTitleOk

`func (o *MovieMetadata) GetCollectionTitleOk() (*string, bool)`

GetCollectionTitleOk returns a tuple with the CollectionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTitle

`func (o *MovieMetadata) SetCollectionTitle(v string)`

SetCollectionTitle sets CollectionTitle field to given value.

### HasCollectionTitle

`func (o *MovieMetadata) HasCollectionTitle() bool`

HasCollectionTitle returns a boolean if a field has been set.

### SetCollectionTitleNil

`func (o *MovieMetadata) SetCollectionTitleNil(b bool)`

 SetCollectionTitleNil sets the value for CollectionTitle to be an explicit nil

### UnsetCollectionTitle
`func (o *MovieMetadata) UnsetCollectionTitle()`

UnsetCollectionTitle ensures that no value is present for CollectionTitle, not even an explicit nil
### GetLastInfoSync

`func (o *MovieMetadata) GetLastInfoSync() time.Time`

GetLastInfoSync returns the LastInfoSync field if non-nil, zero value otherwise.

### GetLastInfoSyncOk

`func (o *MovieMetadata) GetLastInfoSyncOk() (*time.Time, bool)`

GetLastInfoSyncOk returns a tuple with the LastInfoSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInfoSync

`func (o *MovieMetadata) SetLastInfoSync(v time.Time)`

SetLastInfoSync sets LastInfoSync field to given value.

### HasLastInfoSync

`func (o *MovieMetadata) HasLastInfoSync() bool`

HasLastInfoSync returns a boolean if a field has been set.

### SetLastInfoSyncNil

`func (o *MovieMetadata) SetLastInfoSyncNil(b bool)`

 SetLastInfoSyncNil sets the value for LastInfoSync to be an explicit nil

### UnsetLastInfoSync
`func (o *MovieMetadata) UnsetLastInfoSync()`

UnsetLastInfoSync ensures that no value is present for LastInfoSync, not even an explicit nil
### GetRuntime

`func (o *MovieMetadata) GetRuntime() int32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *MovieMetadata) GetRuntimeOk() (*int32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *MovieMetadata) SetRuntime(v int32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *MovieMetadata) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetWebsite

`func (o *MovieMetadata) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *MovieMetadata) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *MovieMetadata) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *MovieMetadata) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *MovieMetadata) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *MovieMetadata) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetImdbId

`func (o *MovieMetadata) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *MovieMetadata) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *MovieMetadata) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *MovieMetadata) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### SetImdbIdNil

`func (o *MovieMetadata) SetImdbIdNil(b bool)`

 SetImdbIdNil sets the value for ImdbId to be an explicit nil

### UnsetImdbId
`func (o *MovieMetadata) UnsetImdbId()`

UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
### GetTitle

`func (o *MovieMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MovieMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MovieMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MovieMetadata) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MovieMetadata) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MovieMetadata) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCleanTitle

`func (o *MovieMetadata) GetCleanTitle() string`

GetCleanTitle returns the CleanTitle field if non-nil, zero value otherwise.

### GetCleanTitleOk

`func (o *MovieMetadata) GetCleanTitleOk() (*string, bool)`

GetCleanTitleOk returns a tuple with the CleanTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanTitle

`func (o *MovieMetadata) SetCleanTitle(v string)`

SetCleanTitle sets CleanTitle field to given value.

### HasCleanTitle

`func (o *MovieMetadata) HasCleanTitle() bool`

HasCleanTitle returns a boolean if a field has been set.

### SetCleanTitleNil

`func (o *MovieMetadata) SetCleanTitleNil(b bool)`

 SetCleanTitleNil sets the value for CleanTitle to be an explicit nil

### UnsetCleanTitle
`func (o *MovieMetadata) UnsetCleanTitle()`

UnsetCleanTitle ensures that no value is present for CleanTitle, not even an explicit nil
### GetSortTitle

`func (o *MovieMetadata) GetSortTitle() string`

GetSortTitle returns the SortTitle field if non-nil, zero value otherwise.

### GetSortTitleOk

`func (o *MovieMetadata) GetSortTitleOk() (*string, bool)`

GetSortTitleOk returns a tuple with the SortTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortTitle

`func (o *MovieMetadata) SetSortTitle(v string)`

SetSortTitle sets SortTitle field to given value.

### HasSortTitle

`func (o *MovieMetadata) HasSortTitle() bool`

HasSortTitle returns a boolean if a field has been set.

### SetSortTitleNil

`func (o *MovieMetadata) SetSortTitleNil(b bool)`

 SetSortTitleNil sets the value for SortTitle to be an explicit nil

### UnsetSortTitle
`func (o *MovieMetadata) UnsetSortTitle()`

UnsetSortTitle ensures that no value is present for SortTitle, not even an explicit nil
### GetStatus

`func (o *MovieMetadata) GetStatus() MovieStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MovieMetadata) GetStatusOk() (*MovieStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MovieMetadata) SetStatus(v MovieStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MovieMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOverview

`func (o *MovieMetadata) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *MovieMetadata) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *MovieMetadata) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *MovieMetadata) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *MovieMetadata) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *MovieMetadata) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetAlternativeTitles

`func (o *MovieMetadata) GetAlternativeTitles() []AlternativeTitle`

GetAlternativeTitles returns the AlternativeTitles field if non-nil, zero value otherwise.

### GetAlternativeTitlesOk

`func (o *MovieMetadata) GetAlternativeTitlesOk() (*[]AlternativeTitle, bool)`

GetAlternativeTitlesOk returns a tuple with the AlternativeTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeTitles

`func (o *MovieMetadata) SetAlternativeTitles(v []AlternativeTitle)`

SetAlternativeTitles sets AlternativeTitles field to given value.

### HasAlternativeTitles

`func (o *MovieMetadata) HasAlternativeTitles() bool`

HasAlternativeTitles returns a boolean if a field has been set.

### SetAlternativeTitlesNil

`func (o *MovieMetadata) SetAlternativeTitlesNil(b bool)`

 SetAlternativeTitlesNil sets the value for AlternativeTitles to be an explicit nil

### UnsetAlternativeTitles
`func (o *MovieMetadata) UnsetAlternativeTitles()`

UnsetAlternativeTitles ensures that no value is present for AlternativeTitles, not even an explicit nil
### GetTranslations

`func (o *MovieMetadata) GetTranslations() []MovieTranslation`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *MovieMetadata) GetTranslationsOk() (*[]MovieTranslation, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *MovieMetadata) SetTranslations(v []MovieTranslation)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *MovieMetadata) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### SetTranslationsNil

`func (o *MovieMetadata) SetTranslationsNil(b bool)`

 SetTranslationsNil sets the value for Translations to be an explicit nil

### UnsetTranslations
`func (o *MovieMetadata) UnsetTranslations()`

UnsetTranslations ensures that no value is present for Translations, not even an explicit nil
### GetSecondaryYear

`func (o *MovieMetadata) GetSecondaryYear() int32`

GetSecondaryYear returns the SecondaryYear field if non-nil, zero value otherwise.

### GetSecondaryYearOk

`func (o *MovieMetadata) GetSecondaryYearOk() (*int32, bool)`

GetSecondaryYearOk returns a tuple with the SecondaryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryYear

`func (o *MovieMetadata) SetSecondaryYear(v int32)`

SetSecondaryYear sets SecondaryYear field to given value.

### HasSecondaryYear

`func (o *MovieMetadata) HasSecondaryYear() bool`

HasSecondaryYear returns a boolean if a field has been set.

### SetSecondaryYearNil

`func (o *MovieMetadata) SetSecondaryYearNil(b bool)`

 SetSecondaryYearNil sets the value for SecondaryYear to be an explicit nil

### UnsetSecondaryYear
`func (o *MovieMetadata) UnsetSecondaryYear()`

UnsetSecondaryYear ensures that no value is present for SecondaryYear, not even an explicit nil
### GetYouTubeTrailerId

`func (o *MovieMetadata) GetYouTubeTrailerId() string`

GetYouTubeTrailerId returns the YouTubeTrailerId field if non-nil, zero value otherwise.

### GetYouTubeTrailerIdOk

`func (o *MovieMetadata) GetYouTubeTrailerIdOk() (*string, bool)`

GetYouTubeTrailerIdOk returns a tuple with the YouTubeTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeTrailerId

`func (o *MovieMetadata) SetYouTubeTrailerId(v string)`

SetYouTubeTrailerId sets YouTubeTrailerId field to given value.

### HasYouTubeTrailerId

`func (o *MovieMetadata) HasYouTubeTrailerId() bool`

HasYouTubeTrailerId returns a boolean if a field has been set.

### SetYouTubeTrailerIdNil

`func (o *MovieMetadata) SetYouTubeTrailerIdNil(b bool)`

 SetYouTubeTrailerIdNil sets the value for YouTubeTrailerId to be an explicit nil

### UnsetYouTubeTrailerId
`func (o *MovieMetadata) UnsetYouTubeTrailerId()`

UnsetYouTubeTrailerId ensures that no value is present for YouTubeTrailerId, not even an explicit nil
### GetStudio

`func (o *MovieMetadata) GetStudio() string`

GetStudio returns the Studio field if non-nil, zero value otherwise.

### GetStudioOk

`func (o *MovieMetadata) GetStudioOk() (*string, bool)`

GetStudioOk returns a tuple with the Studio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudio

`func (o *MovieMetadata) SetStudio(v string)`

SetStudio sets Studio field to given value.

### HasStudio

`func (o *MovieMetadata) HasStudio() bool`

HasStudio returns a boolean if a field has been set.

### SetStudioNil

`func (o *MovieMetadata) SetStudioNil(b bool)`

 SetStudioNil sets the value for Studio to be an explicit nil

### UnsetStudio
`func (o *MovieMetadata) UnsetStudio()`

UnsetStudio ensures that no value is present for Studio, not even an explicit nil
### GetOriginalTitle

`func (o *MovieMetadata) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *MovieMetadata) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *MovieMetadata) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *MovieMetadata) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *MovieMetadata) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *MovieMetadata) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetCleanOriginalTitle

`func (o *MovieMetadata) GetCleanOriginalTitle() string`

GetCleanOriginalTitle returns the CleanOriginalTitle field if non-nil, zero value otherwise.

### GetCleanOriginalTitleOk

`func (o *MovieMetadata) GetCleanOriginalTitleOk() (*string, bool)`

GetCleanOriginalTitleOk returns a tuple with the CleanOriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanOriginalTitle

`func (o *MovieMetadata) SetCleanOriginalTitle(v string)`

SetCleanOriginalTitle sets CleanOriginalTitle field to given value.

### HasCleanOriginalTitle

`func (o *MovieMetadata) HasCleanOriginalTitle() bool`

HasCleanOriginalTitle returns a boolean if a field has been set.

### SetCleanOriginalTitleNil

`func (o *MovieMetadata) SetCleanOriginalTitleNil(b bool)`

 SetCleanOriginalTitleNil sets the value for CleanOriginalTitle to be an explicit nil

### UnsetCleanOriginalTitle
`func (o *MovieMetadata) UnsetCleanOriginalTitle()`

UnsetCleanOriginalTitle ensures that no value is present for CleanOriginalTitle, not even an explicit nil
### GetOriginalLanguage

`func (o *MovieMetadata) GetOriginalLanguage() Language`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *MovieMetadata) GetOriginalLanguageOk() (*Language, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *MovieMetadata) SetOriginalLanguage(v Language)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *MovieMetadata) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetRecommendations

`func (o *MovieMetadata) GetRecommendations() []int32`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *MovieMetadata) GetRecommendationsOk() (*[]int32, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *MovieMetadata) SetRecommendations(v []int32)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *MovieMetadata) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### SetRecommendationsNil

`func (o *MovieMetadata) SetRecommendationsNil(b bool)`

 SetRecommendationsNil sets the value for Recommendations to be an explicit nil

### UnsetRecommendations
`func (o *MovieMetadata) UnsetRecommendations()`

UnsetRecommendations ensures that no value is present for Recommendations, not even an explicit nil
### GetPopularity

`func (o *MovieMetadata) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *MovieMetadata) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *MovieMetadata) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *MovieMetadata) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetIsRecentMovie

`func (o *MovieMetadata) GetIsRecentMovie() bool`

GetIsRecentMovie returns the IsRecentMovie field if non-nil, zero value otherwise.

### GetIsRecentMovieOk

`func (o *MovieMetadata) GetIsRecentMovieOk() (*bool, bool)`

GetIsRecentMovieOk returns a tuple with the IsRecentMovie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecentMovie

`func (o *MovieMetadata) SetIsRecentMovie(v bool)`

SetIsRecentMovie sets IsRecentMovie field to given value.

### HasIsRecentMovie

`func (o *MovieMetadata) HasIsRecentMovie() bool`

HasIsRecentMovie returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



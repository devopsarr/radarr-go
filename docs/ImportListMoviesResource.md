# ImportListMoviesResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**SortTitle** | Pointer to **NullableString** |  | [optional] 
**OriginalLanguage** | Pointer to [**Language**](Language.md) |  | [optional] 
**Status** | Pointer to [**MovieStatusType**](MovieStatusType.md) |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**InCinemas** | Pointer to **NullableTime** |  | [optional] 
**PhysicalRelease** | Pointer to **NullableTime** |  | [optional] 
**DigitalRelease** | Pointer to **NullableTime** |  | [optional] 
**Images** | Pointer to [**[]MediaCover**](MediaCover.md) |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**RemotePoster** | Pointer to **NullableString** |  | [optional] 
**Year** | Pointer to **int32** |  | [optional] 
**YouTubeTrailerId** | Pointer to **NullableString** |  | [optional] 
**Studio** | Pointer to **NullableString** |  | [optional] 
**Runtime** | Pointer to **int32** |  | [optional] 
**ImdbId** | Pointer to **NullableString** |  | [optional] 
**TmdbId** | Pointer to **int32** |  | [optional] 
**Folder** | Pointer to **NullableString** |  | [optional] 
**Certification** | Pointer to **NullableString** |  | [optional] 
**Popularity** | Pointer to **float32** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**Ratings** | Pointer to [**Ratings**](Ratings.md) |  | [optional] 
**Collection** | Pointer to [**MovieCollection**](MovieCollection.md) |  | [optional] 
**IsExcluded** | Pointer to **bool** |  | [optional] 
**IsExisting** | Pointer to **bool** |  | [optional] 
**IsTrending** | Pointer to **bool** |  | [optional] 
**IsPopular** | Pointer to **bool** |  | [optional] 
**IsRecommendation** | Pointer to **bool** |  | [optional] 
**Lists** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewImportListMoviesResource

`func NewImportListMoviesResource() *ImportListMoviesResource`

NewImportListMoviesResource instantiates a new ImportListMoviesResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportListMoviesResourceWithDefaults

`func NewImportListMoviesResourceWithDefaults() *ImportListMoviesResource`

NewImportListMoviesResourceWithDefaults instantiates a new ImportListMoviesResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportListMoviesResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportListMoviesResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportListMoviesResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ImportListMoviesResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ImportListMoviesResource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ImportListMoviesResource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ImportListMoviesResource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ImportListMoviesResource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ImportListMoviesResource) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ImportListMoviesResource) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSortTitle

`func (o *ImportListMoviesResource) GetSortTitle() string`

GetSortTitle returns the SortTitle field if non-nil, zero value otherwise.

### GetSortTitleOk

`func (o *ImportListMoviesResource) GetSortTitleOk() (*string, bool)`

GetSortTitleOk returns a tuple with the SortTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortTitle

`func (o *ImportListMoviesResource) SetSortTitle(v string)`

SetSortTitle sets SortTitle field to given value.

### HasSortTitle

`func (o *ImportListMoviesResource) HasSortTitle() bool`

HasSortTitle returns a boolean if a field has been set.

### SetSortTitleNil

`func (o *ImportListMoviesResource) SetSortTitleNil(b bool)`

 SetSortTitleNil sets the value for SortTitle to be an explicit nil

### UnsetSortTitle
`func (o *ImportListMoviesResource) UnsetSortTitle()`

UnsetSortTitle ensures that no value is present for SortTitle, not even an explicit nil
### GetOriginalLanguage

`func (o *ImportListMoviesResource) GetOriginalLanguage() Language`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *ImportListMoviesResource) GetOriginalLanguageOk() (*Language, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *ImportListMoviesResource) SetOriginalLanguage(v Language)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *ImportListMoviesResource) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetStatus

`func (o *ImportListMoviesResource) GetStatus() MovieStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportListMoviesResource) GetStatusOk() (*MovieStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportListMoviesResource) SetStatus(v MovieStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImportListMoviesResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOverview

`func (o *ImportListMoviesResource) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ImportListMoviesResource) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ImportListMoviesResource) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ImportListMoviesResource) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *ImportListMoviesResource) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *ImportListMoviesResource) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetInCinemas

`func (o *ImportListMoviesResource) GetInCinemas() time.Time`

GetInCinemas returns the InCinemas field if non-nil, zero value otherwise.

### GetInCinemasOk

`func (o *ImportListMoviesResource) GetInCinemasOk() (*time.Time, bool)`

GetInCinemasOk returns a tuple with the InCinemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInCinemas

`func (o *ImportListMoviesResource) SetInCinemas(v time.Time)`

SetInCinemas sets InCinemas field to given value.

### HasInCinemas

`func (o *ImportListMoviesResource) HasInCinemas() bool`

HasInCinemas returns a boolean if a field has been set.

### SetInCinemasNil

`func (o *ImportListMoviesResource) SetInCinemasNil(b bool)`

 SetInCinemasNil sets the value for InCinemas to be an explicit nil

### UnsetInCinemas
`func (o *ImportListMoviesResource) UnsetInCinemas()`

UnsetInCinemas ensures that no value is present for InCinemas, not even an explicit nil
### GetPhysicalRelease

`func (o *ImportListMoviesResource) GetPhysicalRelease() time.Time`

GetPhysicalRelease returns the PhysicalRelease field if non-nil, zero value otherwise.

### GetPhysicalReleaseOk

`func (o *ImportListMoviesResource) GetPhysicalReleaseOk() (*time.Time, bool)`

GetPhysicalReleaseOk returns a tuple with the PhysicalRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalRelease

`func (o *ImportListMoviesResource) SetPhysicalRelease(v time.Time)`

SetPhysicalRelease sets PhysicalRelease field to given value.

### HasPhysicalRelease

`func (o *ImportListMoviesResource) HasPhysicalRelease() bool`

HasPhysicalRelease returns a boolean if a field has been set.

### SetPhysicalReleaseNil

`func (o *ImportListMoviesResource) SetPhysicalReleaseNil(b bool)`

 SetPhysicalReleaseNil sets the value for PhysicalRelease to be an explicit nil

### UnsetPhysicalRelease
`func (o *ImportListMoviesResource) UnsetPhysicalRelease()`

UnsetPhysicalRelease ensures that no value is present for PhysicalRelease, not even an explicit nil
### GetDigitalRelease

`func (o *ImportListMoviesResource) GetDigitalRelease() time.Time`

GetDigitalRelease returns the DigitalRelease field if non-nil, zero value otherwise.

### GetDigitalReleaseOk

`func (o *ImportListMoviesResource) GetDigitalReleaseOk() (*time.Time, bool)`

GetDigitalReleaseOk returns a tuple with the DigitalRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalRelease

`func (o *ImportListMoviesResource) SetDigitalRelease(v time.Time)`

SetDigitalRelease sets DigitalRelease field to given value.

### HasDigitalRelease

`func (o *ImportListMoviesResource) HasDigitalRelease() bool`

HasDigitalRelease returns a boolean if a field has been set.

### SetDigitalReleaseNil

`func (o *ImportListMoviesResource) SetDigitalReleaseNil(b bool)`

 SetDigitalReleaseNil sets the value for DigitalRelease to be an explicit nil

### UnsetDigitalRelease
`func (o *ImportListMoviesResource) UnsetDigitalRelease()`

UnsetDigitalRelease ensures that no value is present for DigitalRelease, not even an explicit nil
### GetImages

`func (o *ImportListMoviesResource) GetImages() []MediaCover`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ImportListMoviesResource) GetImagesOk() (*[]MediaCover, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ImportListMoviesResource) SetImages(v []MediaCover)`

SetImages sets Images field to given value.

### HasImages

`func (o *ImportListMoviesResource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *ImportListMoviesResource) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *ImportListMoviesResource) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetWebsite

`func (o *ImportListMoviesResource) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ImportListMoviesResource) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ImportListMoviesResource) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ImportListMoviesResource) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *ImportListMoviesResource) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *ImportListMoviesResource) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetRemotePoster

`func (o *ImportListMoviesResource) GetRemotePoster() string`

GetRemotePoster returns the RemotePoster field if non-nil, zero value otherwise.

### GetRemotePosterOk

`func (o *ImportListMoviesResource) GetRemotePosterOk() (*string, bool)`

GetRemotePosterOk returns a tuple with the RemotePoster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePoster

`func (o *ImportListMoviesResource) SetRemotePoster(v string)`

SetRemotePoster sets RemotePoster field to given value.

### HasRemotePoster

`func (o *ImportListMoviesResource) HasRemotePoster() bool`

HasRemotePoster returns a boolean if a field has been set.

### SetRemotePosterNil

`func (o *ImportListMoviesResource) SetRemotePosterNil(b bool)`

 SetRemotePosterNil sets the value for RemotePoster to be an explicit nil

### UnsetRemotePoster
`func (o *ImportListMoviesResource) UnsetRemotePoster()`

UnsetRemotePoster ensures that no value is present for RemotePoster, not even an explicit nil
### GetYear

`func (o *ImportListMoviesResource) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ImportListMoviesResource) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ImportListMoviesResource) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ImportListMoviesResource) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetYouTubeTrailerId

`func (o *ImportListMoviesResource) GetYouTubeTrailerId() string`

GetYouTubeTrailerId returns the YouTubeTrailerId field if non-nil, zero value otherwise.

### GetYouTubeTrailerIdOk

`func (o *ImportListMoviesResource) GetYouTubeTrailerIdOk() (*string, bool)`

GetYouTubeTrailerIdOk returns a tuple with the YouTubeTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeTrailerId

`func (o *ImportListMoviesResource) SetYouTubeTrailerId(v string)`

SetYouTubeTrailerId sets YouTubeTrailerId field to given value.

### HasYouTubeTrailerId

`func (o *ImportListMoviesResource) HasYouTubeTrailerId() bool`

HasYouTubeTrailerId returns a boolean if a field has been set.

### SetYouTubeTrailerIdNil

`func (o *ImportListMoviesResource) SetYouTubeTrailerIdNil(b bool)`

 SetYouTubeTrailerIdNil sets the value for YouTubeTrailerId to be an explicit nil

### UnsetYouTubeTrailerId
`func (o *ImportListMoviesResource) UnsetYouTubeTrailerId()`

UnsetYouTubeTrailerId ensures that no value is present for YouTubeTrailerId, not even an explicit nil
### GetStudio

`func (o *ImportListMoviesResource) GetStudio() string`

GetStudio returns the Studio field if non-nil, zero value otherwise.

### GetStudioOk

`func (o *ImportListMoviesResource) GetStudioOk() (*string, bool)`

GetStudioOk returns a tuple with the Studio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudio

`func (o *ImportListMoviesResource) SetStudio(v string)`

SetStudio sets Studio field to given value.

### HasStudio

`func (o *ImportListMoviesResource) HasStudio() bool`

HasStudio returns a boolean if a field has been set.

### SetStudioNil

`func (o *ImportListMoviesResource) SetStudioNil(b bool)`

 SetStudioNil sets the value for Studio to be an explicit nil

### UnsetStudio
`func (o *ImportListMoviesResource) UnsetStudio()`

UnsetStudio ensures that no value is present for Studio, not even an explicit nil
### GetRuntime

`func (o *ImportListMoviesResource) GetRuntime() int32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ImportListMoviesResource) GetRuntimeOk() (*int32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ImportListMoviesResource) SetRuntime(v int32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *ImportListMoviesResource) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetImdbId

`func (o *ImportListMoviesResource) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *ImportListMoviesResource) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *ImportListMoviesResource) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *ImportListMoviesResource) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### SetImdbIdNil

`func (o *ImportListMoviesResource) SetImdbIdNil(b bool)`

 SetImdbIdNil sets the value for ImdbId to be an explicit nil

### UnsetImdbId
`func (o *ImportListMoviesResource) UnsetImdbId()`

UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
### GetTmdbId

`func (o *ImportListMoviesResource) GetTmdbId() int32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *ImportListMoviesResource) GetTmdbIdOk() (*int32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *ImportListMoviesResource) SetTmdbId(v int32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *ImportListMoviesResource) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetFolder

`func (o *ImportListMoviesResource) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *ImportListMoviesResource) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *ImportListMoviesResource) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *ImportListMoviesResource) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *ImportListMoviesResource) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *ImportListMoviesResource) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetCertification

`func (o *ImportListMoviesResource) GetCertification() string`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *ImportListMoviesResource) GetCertificationOk() (*string, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *ImportListMoviesResource) SetCertification(v string)`

SetCertification sets Certification field to given value.

### HasCertification

`func (o *ImportListMoviesResource) HasCertification() bool`

HasCertification returns a boolean if a field has been set.

### SetCertificationNil

`func (o *ImportListMoviesResource) SetCertificationNil(b bool)`

 SetCertificationNil sets the value for Certification to be an explicit nil

### UnsetCertification
`func (o *ImportListMoviesResource) UnsetCertification()`

UnsetCertification ensures that no value is present for Certification, not even an explicit nil
### GetPopularity

`func (o *ImportListMoviesResource) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *ImportListMoviesResource) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *ImportListMoviesResource) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *ImportListMoviesResource) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetGenres

`func (o *ImportListMoviesResource) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *ImportListMoviesResource) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *ImportListMoviesResource) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *ImportListMoviesResource) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *ImportListMoviesResource) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *ImportListMoviesResource) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetRatings

`func (o *ImportListMoviesResource) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *ImportListMoviesResource) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *ImportListMoviesResource) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *ImportListMoviesResource) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### GetCollection

`func (o *ImportListMoviesResource) GetCollection() MovieCollection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ImportListMoviesResource) GetCollectionOk() (*MovieCollection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ImportListMoviesResource) SetCollection(v MovieCollection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *ImportListMoviesResource) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetIsExcluded

`func (o *ImportListMoviesResource) GetIsExcluded() bool`

GetIsExcluded returns the IsExcluded field if non-nil, zero value otherwise.

### GetIsExcludedOk

`func (o *ImportListMoviesResource) GetIsExcludedOk() (*bool, bool)`

GetIsExcludedOk returns a tuple with the IsExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExcluded

`func (o *ImportListMoviesResource) SetIsExcluded(v bool)`

SetIsExcluded sets IsExcluded field to given value.

### HasIsExcluded

`func (o *ImportListMoviesResource) HasIsExcluded() bool`

HasIsExcluded returns a boolean if a field has been set.

### GetIsExisting

`func (o *ImportListMoviesResource) GetIsExisting() bool`

GetIsExisting returns the IsExisting field if non-nil, zero value otherwise.

### GetIsExistingOk

`func (o *ImportListMoviesResource) GetIsExistingOk() (*bool, bool)`

GetIsExistingOk returns a tuple with the IsExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExisting

`func (o *ImportListMoviesResource) SetIsExisting(v bool)`

SetIsExisting sets IsExisting field to given value.

### HasIsExisting

`func (o *ImportListMoviesResource) HasIsExisting() bool`

HasIsExisting returns a boolean if a field has been set.

### GetIsTrending

`func (o *ImportListMoviesResource) GetIsTrending() bool`

GetIsTrending returns the IsTrending field if non-nil, zero value otherwise.

### GetIsTrendingOk

`func (o *ImportListMoviesResource) GetIsTrendingOk() (*bool, bool)`

GetIsTrendingOk returns a tuple with the IsTrending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrending

`func (o *ImportListMoviesResource) SetIsTrending(v bool)`

SetIsTrending sets IsTrending field to given value.

### HasIsTrending

`func (o *ImportListMoviesResource) HasIsTrending() bool`

HasIsTrending returns a boolean if a field has been set.

### GetIsPopular

`func (o *ImportListMoviesResource) GetIsPopular() bool`

GetIsPopular returns the IsPopular field if non-nil, zero value otherwise.

### GetIsPopularOk

`func (o *ImportListMoviesResource) GetIsPopularOk() (*bool, bool)`

GetIsPopularOk returns a tuple with the IsPopular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPopular

`func (o *ImportListMoviesResource) SetIsPopular(v bool)`

SetIsPopular sets IsPopular field to given value.

### HasIsPopular

`func (o *ImportListMoviesResource) HasIsPopular() bool`

HasIsPopular returns a boolean if a field has been set.

### GetIsRecommendation

`func (o *ImportListMoviesResource) GetIsRecommendation() bool`

GetIsRecommendation returns the IsRecommendation field if non-nil, zero value otherwise.

### GetIsRecommendationOk

`func (o *ImportListMoviesResource) GetIsRecommendationOk() (*bool, bool)`

GetIsRecommendationOk returns a tuple with the IsRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecommendation

`func (o *ImportListMoviesResource) SetIsRecommendation(v bool)`

SetIsRecommendation sets IsRecommendation field to given value.

### HasIsRecommendation

`func (o *ImportListMoviesResource) HasIsRecommendation() bool`

HasIsRecommendation returns a boolean if a field has been set.

### GetLists

`func (o *ImportListMoviesResource) GetLists() []int32`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *ImportListMoviesResource) GetListsOk() (*[]int32, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *ImportListMoviesResource) SetLists(v []int32)`

SetLists sets Lists field to given value.

### HasLists

`func (o *ImportListMoviesResource) HasLists() bool`

HasLists returns a boolean if a field has been set.

### SetListsNil

`func (o *ImportListMoviesResource) SetListsNil(b bool)`

 SetListsNil sets the value for Lists to be an explicit nil

### UnsetLists
`func (o *ImportListMoviesResource) UnsetLists()`

UnsetLists ensures that no value is present for Lists, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



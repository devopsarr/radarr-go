# Go API client for radarr

Radarr API docs

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

[comment]: # (x-release-please-start-version)
- Package version: 0.2.1

[comment]: # (x-release-please-end)
- API version: 3.0.0

- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import radarr "github.com/devopsarr/radarr-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), radarr.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), radarr.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), radarr.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), radarr.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:7878*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AlternativeTitleApi* | [**GetAlttitleById**](radarr/docs/AlternativeTitleApi.md#getalttitlebyid) | **Get** /api/v3/alttitle/{id} | 
*AlternativeTitleApi* | [**ListAlttitle**](radarr/docs/AlternativeTitleApi.md#listalttitle) | **Get** /api/v3/alttitle | 
*ApiInfoApi* | [**GetApi**](radarr/docs/ApiInfoApi.md#getapi) | **Get** /api | 
*AuthenticationApi* | [**CreateLogin**](radarr/docs/AuthenticationApi.md#createlogin) | **Post** /login | 
*AuthenticationApi* | [**GetLogout**](radarr/docs/AuthenticationApi.md#getlogout) | **Get** /logout | 
*BackupApi* | [**CreateSystemBackupRestoreById**](radarr/docs/BackupApi.md#createsystembackuprestorebyid) | **Post** /api/v3/system/backup/restore/{id} | 
*BackupApi* | [**CreateSystemBackupRestoreUpload**](radarr/docs/BackupApi.md#createsystembackuprestoreupload) | **Post** /api/v3/system/backup/restore/upload | 
*BackupApi* | [**DeleteSystemBackup**](radarr/docs/BackupApi.md#deletesystembackup) | **Delete** /api/v3/system/backup/{id} | 
*BackupApi* | [**ListSystemBackup**](radarr/docs/BackupApi.md#listsystembackup) | **Get** /api/v3/system/backup | 
*BlocklistApi* | [**DeleteBlocklist**](radarr/docs/BlocklistApi.md#deleteblocklist) | **Delete** /api/v3/blocklist/{id} | 
*BlocklistApi* | [**DeleteBlocklistBulk**](radarr/docs/BlocklistApi.md#deleteblocklistbulk) | **Delete** /api/v3/blocklist/bulk | 
*BlocklistApi* | [**GetBlocklist**](radarr/docs/BlocklistApi.md#getblocklist) | **Get** /api/v3/blocklist | 
*BlocklistApi* | [**ListBlocklistMovie**](radarr/docs/BlocklistApi.md#listblocklistmovie) | **Get** /api/v3/blocklist/movie | 
*CalendarApi* | [**GetCalendarById**](radarr/docs/CalendarApi.md#getcalendarbyid) | **Get** /api/v3/calendar/{id} | 
*CalendarApi* | [**ListCalendar**](radarr/docs/CalendarApi.md#listcalendar) | **Get** /api/v3/calendar | 
*CalendarFeedApi* | [**GetFeedV3CalendarRadarrIcs**](radarr/docs/CalendarFeedApi.md#getfeedv3calendarradarrics) | **Get** /feed/v3/calendar/radarr.ics | 
*CollectionApi* | [**GetCollectionById**](radarr/docs/CollectionApi.md#getcollectionbyid) | **Get** /api/v3/collection/{id} | 
*CollectionApi* | [**ListCollection**](radarr/docs/CollectionApi.md#listcollection) | **Get** /api/v3/collection | 
*CollectionApi* | [**PutCollection**](radarr/docs/CollectionApi.md#putcollection) | **Put** /api/v3/collection | 
*CollectionApi* | [**UpdateCollection**](radarr/docs/CollectionApi.md#updatecollection) | **Put** /api/v3/collection/{id} | 
*CommandApi* | [**CreateCommand**](radarr/docs/CommandApi.md#createcommand) | **Post** /api/v3/command | 
*CommandApi* | [**DeleteCommand**](radarr/docs/CommandApi.md#deletecommand) | **Delete** /api/v3/command/{id} | 
*CommandApi* | [**GetCommandById**](radarr/docs/CommandApi.md#getcommandbyid) | **Get** /api/v3/command/{id} | 
*CommandApi* | [**ListCommand**](radarr/docs/CommandApi.md#listcommand) | **Get** /api/v3/command | 
*CreditApi* | [**GetCreditById**](radarr/docs/CreditApi.md#getcreditbyid) | **Get** /api/v3/credit/{id} | 
*CreditApi* | [**ListCredit**](radarr/docs/CreditApi.md#listcredit) | **Get** /api/v3/credit | 
*CustomFilterApi* | [**CreateCustomFilter**](radarr/docs/CustomFilterApi.md#createcustomfilter) | **Post** /api/v3/customfilter | 
*CustomFilterApi* | [**DeleteCustomFilter**](radarr/docs/CustomFilterApi.md#deletecustomfilter) | **Delete** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**GetCustomFilterById**](radarr/docs/CustomFilterApi.md#getcustomfilterbyid) | **Get** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ListCustomFilter**](radarr/docs/CustomFilterApi.md#listcustomfilter) | **Get** /api/v3/customfilter | 
*CustomFilterApi* | [**UpdateCustomFilter**](radarr/docs/CustomFilterApi.md#updatecustomfilter) | **Put** /api/v3/customfilter/{id} | 
*CustomFormatApi* | [**CreateCustomFormat**](radarr/docs/CustomFormatApi.md#createcustomformat) | **Post** /api/v3/customformat | 
*CustomFormatApi* | [**DeleteCustomFormat**](radarr/docs/CustomFormatApi.md#deletecustomformat) | **Delete** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**GetCustomFormatById**](radarr/docs/CustomFormatApi.md#getcustomformatbyid) | **Get** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**GetCustomFormatSchema**](radarr/docs/CustomFormatApi.md#getcustomformatschema) | **Get** /api/v3/customformat/schema | 
*CustomFormatApi* | [**ListCustomFormat**](radarr/docs/CustomFormatApi.md#listcustomformat) | **Get** /api/v3/customformat | 
*CustomFormatApi* | [**UpdateCustomFormat**](radarr/docs/CustomFormatApi.md#updatecustomformat) | **Put** /api/v3/customformat/{id} | 
*DelayProfileApi* | [**CreateDelayProfile**](radarr/docs/DelayProfileApi.md#createdelayprofile) | **Post** /api/v3/delayprofile | 
*DelayProfileApi* | [**DeleteDelayProfile**](radarr/docs/DelayProfileApi.md#deletedelayprofile) | **Delete** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**GetDelayProfileById**](radarr/docs/DelayProfileApi.md#getdelayprofilebyid) | **Get** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ListDelayProfile**](radarr/docs/DelayProfileApi.md#listdelayprofile) | **Get** /api/v3/delayprofile | 
*DelayProfileApi* | [**UpdateDelayProfile**](radarr/docs/DelayProfileApi.md#updatedelayprofile) | **Put** /api/v3/delayprofile/{id} | 
*DiskSpaceApi* | [**ListDiskSpace**](radarr/docs/DiskSpaceApi.md#listdiskspace) | **Get** /api/v3/diskspace | 
*DownloadClientApi* | [**CreateDownloadClient**](radarr/docs/DownloadClientApi.md#createdownloadclient) | **Post** /api/v3/downloadclient | 
*DownloadClientApi* | [**CreateDownloadClientActionByName**](radarr/docs/DownloadClientApi.md#createdownloadclientactionbyname) | **Post** /api/v3/downloadclient/action/{name} | 
*DownloadClientApi* | [**DeleteDownloadClient**](radarr/docs/DownloadClientApi.md#deletedownloadclient) | **Delete** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**GetDownloadClientById**](radarr/docs/DownloadClientApi.md#getdownloadclientbyid) | **Get** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ListDownloadClient**](radarr/docs/DownloadClientApi.md#listdownloadclient) | **Get** /api/v3/downloadclient | 
*DownloadClientApi* | [**ListDownloadClientSchema**](radarr/docs/DownloadClientApi.md#listdownloadclientschema) | **Get** /api/v3/downloadclient/schema | 
*DownloadClientApi* | [**TestDownloadClient**](radarr/docs/DownloadClientApi.md#testdownloadclient) | **Post** /api/v3/downloadclient/test | 
*DownloadClientApi* | [**TestallDownloadClient**](radarr/docs/DownloadClientApi.md#testalldownloadclient) | **Post** /api/v3/downloadclient/testall | 
*DownloadClientApi* | [**UpdateDownloadClient**](radarr/docs/DownloadClientApi.md#updatedownloadclient) | **Put** /api/v3/downloadclient/{id} | 
*DownloadClientConfigApi* | [**GetDownloadClientConfig**](radarr/docs/DownloadClientConfigApi.md#getdownloadclientconfig) | **Get** /api/v3/config/downloadclient | 
*DownloadClientConfigApi* | [**GetDownloadClientConfigById**](radarr/docs/DownloadClientConfigApi.md#getdownloadclientconfigbyid) | **Get** /api/v3/config/downloadclient/{id} | 
*DownloadClientConfigApi* | [**UpdateDownloadClientConfig**](radarr/docs/DownloadClientConfigApi.md#updatedownloadclientconfig) | **Put** /api/v3/config/downloadclient/{id} | 
*ExtraFileApi* | [**ListExtraFile**](radarr/docs/ExtraFileApi.md#listextrafile) | **Get** /api/v3/extrafile | 
*FileSystemApi* | [**GetFileSystem**](radarr/docs/FileSystemApi.md#getfilesystem) | **Get** /api/v3/filesystem | 
*FileSystemApi* | [**GetFileSystemMediafiles**](radarr/docs/FileSystemApi.md#getfilesystemmediafiles) | **Get** /api/v3/filesystem/mediafiles | 
*FileSystemApi* | [**GetFileSystemType**](radarr/docs/FileSystemApi.md#getfilesystemtype) | **Get** /api/v3/filesystem/type | 
*HealthApi* | [**GetHealthById**](radarr/docs/HealthApi.md#gethealthbyid) | **Get** /api/v3/health/{id} | 
*HealthApi* | [**ListHealth**](radarr/docs/HealthApi.md#listhealth) | **Get** /api/v3/health | 
*HistoryApi* | [**CreateHistoryFailedById**](radarr/docs/HistoryApi.md#createhistoryfailedbyid) | **Post** /api/v3/history/failed/{id} | 
*HistoryApi* | [**GetHistory**](radarr/docs/HistoryApi.md#gethistory) | **Get** /api/v3/history | 
*HistoryApi* | [**ListHistoryMovie**](radarr/docs/HistoryApi.md#listhistorymovie) | **Get** /api/v3/history/movie | 
*HistoryApi* | [**ListHistorySince**](radarr/docs/HistoryApi.md#listhistorysince) | **Get** /api/v3/history/since | 
*HostConfigApi* | [**GetHostConfig**](radarr/docs/HostConfigApi.md#gethostconfig) | **Get** /api/v3/config/host | 
*HostConfigApi* | [**GetHostConfigById**](radarr/docs/HostConfigApi.md#gethostconfigbyid) | **Get** /api/v3/config/host/{id} | 
*HostConfigApi* | [**UpdateHostConfig**](radarr/docs/HostConfigApi.md#updatehostconfig) | **Put** /api/v3/config/host/{id} | 
*ImportExclusionsApi* | [**CreateExclusions**](radarr/docs/ImportExclusionsApi.md#createexclusions) | **Post** /api/v3/exclusions | 
*ImportExclusionsApi* | [**CreateExclusionsBulk**](radarr/docs/ImportExclusionsApi.md#createexclusionsbulk) | **Post** /api/v3/exclusions/bulk | 
*ImportExclusionsApi* | [**DeleteExclusions**](radarr/docs/ImportExclusionsApi.md#deleteexclusions) | **Delete** /api/v3/exclusions/{id} | 
*ImportExclusionsApi* | [**GetExclusionsById**](radarr/docs/ImportExclusionsApi.md#getexclusionsbyid) | **Get** /api/v3/exclusions/{id} | 
*ImportExclusionsApi* | [**ListExclusions**](radarr/docs/ImportExclusionsApi.md#listexclusions) | **Get** /api/v3/exclusions | 
*ImportExclusionsApi* | [**UpdateExclusions**](radarr/docs/ImportExclusionsApi.md#updateexclusions) | **Put** /api/v3/exclusions/{id} | 
*ImportListApi* | [**CreateImportList**](radarr/docs/ImportListApi.md#createimportlist) | **Post** /api/v3/importlist | 
*ImportListApi* | [**CreateImportListActionByName**](radarr/docs/ImportListApi.md#createimportlistactionbyname) | **Post** /api/v3/importlist/action/{name} | 
*ImportListApi* | [**DeleteImportList**](radarr/docs/ImportListApi.md#deleteimportlist) | **Delete** /api/v3/importlist/{id} | 
*ImportListApi* | [**GetImportListById**](radarr/docs/ImportListApi.md#getimportlistbyid) | **Get** /api/v3/importlist/{id} | 
*ImportListApi* | [**ListImportList**](radarr/docs/ImportListApi.md#listimportlist) | **Get** /api/v3/importlist | 
*ImportListApi* | [**ListImportListSchema**](radarr/docs/ImportListApi.md#listimportlistschema) | **Get** /api/v3/importlist/schema | 
*ImportListApi* | [**TestImportList**](radarr/docs/ImportListApi.md#testimportlist) | **Post** /api/v3/importlist/test | 
*ImportListApi* | [**TestallImportList**](radarr/docs/ImportListApi.md#testallimportlist) | **Post** /api/v3/importlist/testall | 
*ImportListApi* | [**UpdateImportList**](radarr/docs/ImportListApi.md#updateimportlist) | **Put** /api/v3/importlist/{id} | 
*ImportListConfigApi* | [**GetImportListConfig**](radarr/docs/ImportListConfigApi.md#getimportlistconfig) | **Get** /api/v3/config/importlist | 
*ImportListConfigApi* | [**GetImportListConfigById**](radarr/docs/ImportListConfigApi.md#getimportlistconfigbyid) | **Get** /api/v3/config/importlist/{id} | 
*ImportListConfigApi* | [**UpdateImportListConfig**](radarr/docs/ImportListConfigApi.md#updateimportlistconfig) | **Put** /api/v3/config/importlist/{id} | 
*ImportListMoviesApi* | [**CreateImportlistMovie**](radarr/docs/ImportListMoviesApi.md#createimportlistmovie) | **Post** /api/v3/importlist/movie | 
*ImportListMoviesApi* | [**GetImportlistMovie**](radarr/docs/ImportListMoviesApi.md#getimportlistmovie) | **Get** /api/v3/importlist/movie | 
*IndexerApi* | [**CreateIndexer**](radarr/docs/IndexerApi.md#createindexer) | **Post** /api/v3/indexer | 
*IndexerApi* | [**CreateIndexerActionByName**](radarr/docs/IndexerApi.md#createindexeractionbyname) | **Post** /api/v3/indexer/action/{name} | 
*IndexerApi* | [**DeleteIndexer**](radarr/docs/IndexerApi.md#deleteindexer) | **Delete** /api/v3/indexer/{id} | 
*IndexerApi* | [**GetIndexerById**](radarr/docs/IndexerApi.md#getindexerbyid) | **Get** /api/v3/indexer/{id} | 
*IndexerApi* | [**ListIndexer**](radarr/docs/IndexerApi.md#listindexer) | **Get** /api/v3/indexer | 
*IndexerApi* | [**ListIndexerSchema**](radarr/docs/IndexerApi.md#listindexerschema) | **Get** /api/v3/indexer/schema | 
*IndexerApi* | [**TestIndexer**](radarr/docs/IndexerApi.md#testindexer) | **Post** /api/v3/indexer/test | 
*IndexerApi* | [**TestallIndexer**](radarr/docs/IndexerApi.md#testallindexer) | **Post** /api/v3/indexer/testall | 
*IndexerApi* | [**UpdateIndexer**](radarr/docs/IndexerApi.md#updateindexer) | **Put** /api/v3/indexer/{id} | 
*IndexerConfigApi* | [**GetIndexerConfig**](radarr/docs/IndexerConfigApi.md#getindexerconfig) | **Get** /api/v3/config/indexer | 
*IndexerConfigApi* | [**GetIndexerConfigById**](radarr/docs/IndexerConfigApi.md#getindexerconfigbyid) | **Get** /api/v3/config/indexer/{id} | 
*IndexerConfigApi* | [**UpdateIndexerConfig**](radarr/docs/IndexerConfigApi.md#updateindexerconfig) | **Put** /api/v3/config/indexer/{id} | 
*IndexerFlagApi* | [**ListIndexerFlag**](radarr/docs/IndexerFlagApi.md#listindexerflag) | **Get** /api/v3/indexerflag | 
*InitializeJsApi* | [**GetInitializeJs**](radarr/docs/InitializeJsApi.md#getinitializejs) | **Get** /initialize.js | 
*LanguageApi* | [**GetLanguageById**](radarr/docs/LanguageApi.md#getlanguagebyid) | **Get** /api/v3/language/{id} | 
*LanguageApi* | [**ListLanguage**](radarr/docs/LanguageApi.md#listlanguage) | **Get** /api/v3/language | 
*LocalizationApi* | [**GetLocalization**](radarr/docs/LocalizationApi.md#getlocalization) | **Get** /api/v3/localization | 
*LogApi* | [**GetLog**](radarr/docs/LogApi.md#getlog) | **Get** /api/v3/log | 
*LogFileApi* | [**GetLogFileByFilename**](radarr/docs/LogFileApi.md#getlogfilebyfilename) | **Get** /api/v3/log/file/{filename} | 
*LogFileApi* | [**ListLogFile**](radarr/docs/LogFileApi.md#listlogfile) | **Get** /api/v3/log/file | 
*ManualImportApi* | [**CreateManualImport**](radarr/docs/ManualImportApi.md#createmanualimport) | **Post** /api/v3/manualimport | 
*ManualImportApi* | [**ListManualImport**](radarr/docs/ManualImportApi.md#listmanualimport) | **Get** /api/v3/manualimport | 
*MediaCoverApi* | [**GetMediaCovermovieIdByFilename**](radarr/docs/MediaCoverApi.md#getmediacovermovieidbyfilename) | **Get** /api/v3/mediacover/{movieId}/{filename} | 
*MediaManagementConfigApi* | [**GetMediaManagementConfig**](radarr/docs/MediaManagementConfigApi.md#getmediamanagementconfig) | **Get** /api/v3/config/mediamanagement | 
*MediaManagementConfigApi* | [**GetMediaManagementConfigById**](radarr/docs/MediaManagementConfigApi.md#getmediamanagementconfigbyid) | **Get** /api/v3/config/mediamanagement/{id} | 
*MediaManagementConfigApi* | [**UpdateMediaManagementConfig**](radarr/docs/MediaManagementConfigApi.md#updatemediamanagementconfig) | **Put** /api/v3/config/mediamanagement/{id} | 
*MetadataApi* | [**CreateMetadata**](radarr/docs/MetadataApi.md#createmetadata) | **Post** /api/v3/metadata | 
*MetadataApi* | [**CreateMetadataActionByName**](radarr/docs/MetadataApi.md#createmetadataactionbyname) | **Post** /api/v3/metadata/action/{name} | 
*MetadataApi* | [**DeleteMetadata**](radarr/docs/MetadataApi.md#deletemetadata) | **Delete** /api/v3/metadata/{id} | 
*MetadataApi* | [**GetMetadataById**](radarr/docs/MetadataApi.md#getmetadatabyid) | **Get** /api/v3/metadata/{id} | 
*MetadataApi* | [**ListMetadata**](radarr/docs/MetadataApi.md#listmetadata) | **Get** /api/v3/metadata | 
*MetadataApi* | [**ListMetadataSchema**](radarr/docs/MetadataApi.md#listmetadataschema) | **Get** /api/v3/metadata/schema | 
*MetadataApi* | [**TestMetadata**](radarr/docs/MetadataApi.md#testmetadata) | **Post** /api/v3/metadata/test | 
*MetadataApi* | [**TestallMetadata**](radarr/docs/MetadataApi.md#testallmetadata) | **Post** /api/v3/metadata/testall | 
*MetadataApi* | [**UpdateMetadata**](radarr/docs/MetadataApi.md#updatemetadata) | **Put** /api/v3/metadata/{id} | 
*MetadataConfigApi* | [**GetMetadataConfig**](radarr/docs/MetadataConfigApi.md#getmetadataconfig) | **Get** /api/v3/config/metadata | 
*MetadataConfigApi* | [**GetMetadataConfigById**](radarr/docs/MetadataConfigApi.md#getmetadataconfigbyid) | **Get** /api/v3/config/metadata/{id} | 
*MetadataConfigApi* | [**UpdateMetadataConfig**](radarr/docs/MetadataConfigApi.md#updatemetadataconfig) | **Put** /api/v3/config/metadata/{id} | 
*MovieApi* | [**CreateMovie**](radarr/docs/MovieApi.md#createmovie) | **Post** /api/v3/movie | 
*MovieApi* | [**DeleteMovie**](radarr/docs/MovieApi.md#deletemovie) | **Delete** /api/v3/movie/{id} | 
*MovieApi* | [**GetMovieById**](radarr/docs/MovieApi.md#getmoviebyid) | **Get** /api/v3/movie/{id} | 
*MovieApi* | [**ListMovie**](radarr/docs/MovieApi.md#listmovie) | **Get** /api/v3/movie | 
*MovieApi* | [**UpdateMovie**](radarr/docs/MovieApi.md#updatemovie) | **Put** /api/v3/movie/{id} | 
*MovieEditorApi* | [**DeleteMovieEditor**](radarr/docs/MovieEditorApi.md#deletemovieeditor) | **Delete** /api/v3/movie/editor | 
*MovieEditorApi* | [**PutMovieEditor**](radarr/docs/MovieEditorApi.md#putmovieeditor) | **Put** /api/v3/movie/editor | 
*MovieFileApi* | [**DeleteMovieFile**](radarr/docs/MovieFileApi.md#deletemoviefile) | **Delete** /api/v3/moviefile/{id} | 
*MovieFileApi* | [**DeleteMovieFileBulk**](radarr/docs/MovieFileApi.md#deletemoviefilebulk) | **Delete** /api/v3/moviefile/bulk | 
*MovieFileApi* | [**GetMovieFileById**](radarr/docs/MovieFileApi.md#getmoviefilebyid) | **Get** /api/v3/moviefile/{id} | 
*MovieFileApi* | [**ListMovieFile**](radarr/docs/MovieFileApi.md#listmoviefile) | **Get** /api/v3/moviefile | 
*MovieFileApi* | [**PutMovieFileEditor**](radarr/docs/MovieFileApi.md#putmoviefileeditor) | **Put** /api/v3/moviefile/editor | 
*MovieFileApi* | [**UpdateMovieFile**](radarr/docs/MovieFileApi.md#updatemoviefile) | **Put** /api/v3/moviefile/{id} | 
*MovieImportApi* | [**CreateMovieImport**](radarr/docs/MovieImportApi.md#createmovieimport) | **Post** /api/v3/movie/import | 
*MovieImportApi* | [**GetMovieImportById**](radarr/docs/MovieImportApi.md#getmovieimportbyid) | **Get** /api/v3/movie/import/{id} | 
*MovieLookupApi* | [**GetMovieLookup**](radarr/docs/MovieLookupApi.md#getmovielookup) | **Get** /api/v3/movie/lookup | 
*MovieLookupApi* | [**GetMovieLookupById**](radarr/docs/MovieLookupApi.md#getmovielookupbyid) | **Get** /api/v3/movie/lookup/{id} | 
*MovieLookupApi* | [**GetMovieLookupImdb**](radarr/docs/MovieLookupApi.md#getmovielookupimdb) | **Get** /api/v3/movie/lookup/imdb | 
*MovieLookupApi* | [**GetMovieLookupTmdb**](radarr/docs/MovieLookupApi.md#getmovielookuptmdb) | **Get** /api/v3/movie/lookup/tmdb | 
*NamingConfigApi* | [**GetNamingConfig**](radarr/docs/NamingConfigApi.md#getnamingconfig) | **Get** /api/v3/config/naming | 
*NamingConfigApi* | [**GetNamingConfigById**](radarr/docs/NamingConfigApi.md#getnamingconfigbyid) | **Get** /api/v3/config/naming/{id} | 
*NamingConfigApi* | [**GetNamingConfigExamples**](radarr/docs/NamingConfigApi.md#getnamingconfigexamples) | **Get** /api/v3/config/naming/examples | 
*NamingConfigApi* | [**UpdateNamingConfig**](radarr/docs/NamingConfigApi.md#updatenamingconfig) | **Put** /api/v3/config/naming/{id} | 
*NotificationApi* | [**CreateNotification**](radarr/docs/NotificationApi.md#createnotification) | **Post** /api/v3/notification | 
*NotificationApi* | [**CreateNotificationActionByName**](radarr/docs/NotificationApi.md#createnotificationactionbyname) | **Post** /api/v3/notification/action/{name} | 
*NotificationApi* | [**DeleteNotification**](radarr/docs/NotificationApi.md#deletenotification) | **Delete** /api/v3/notification/{id} | 
*NotificationApi* | [**GetNotificationById**](radarr/docs/NotificationApi.md#getnotificationbyid) | **Get** /api/v3/notification/{id} | 
*NotificationApi* | [**ListNotification**](radarr/docs/NotificationApi.md#listnotification) | **Get** /api/v3/notification | 
*NotificationApi* | [**ListNotificationSchema**](radarr/docs/NotificationApi.md#listnotificationschema) | **Get** /api/v3/notification/schema | 
*NotificationApi* | [**TestNotification**](radarr/docs/NotificationApi.md#testnotification) | **Post** /api/v3/notification/test | 
*NotificationApi* | [**TestallNotification**](radarr/docs/NotificationApi.md#testallnotification) | **Post** /api/v3/notification/testall | 
*NotificationApi* | [**UpdateNotification**](radarr/docs/NotificationApi.md#updatenotification) | **Put** /api/v3/notification/{id} | 
*ParseApi* | [**GetParse**](radarr/docs/ParseApi.md#getparse) | **Get** /api/v3/parse | 
*PingApi* | [**GetPing**](radarr/docs/PingApi.md#getping) | **Get** /ping | 
*QualityDefinitionApi* | [**GetQualityDefinitionById**](radarr/docs/QualityDefinitionApi.md#getqualitydefinitionbyid) | **Get** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionApi* | [**ListQualityDefinition**](radarr/docs/QualityDefinitionApi.md#listqualitydefinition) | **Get** /api/v3/qualitydefinition | 
*QualityDefinitionApi* | [**PutQualityDefinitionUpdate**](radarr/docs/QualityDefinitionApi.md#putqualitydefinitionupdate) | **Put** /api/v3/qualitydefinition/update | 
*QualityDefinitionApi* | [**UpdateQualityDefinition**](radarr/docs/QualityDefinitionApi.md#updatequalitydefinition) | **Put** /api/v3/qualitydefinition/{id} | 
*QualityProfileApi* | [**CreateQualityProfile**](radarr/docs/QualityProfileApi.md#createqualityprofile) | **Post** /api/v3/qualityprofile | 
*QualityProfileApi* | [**DeleteQualityProfile**](radarr/docs/QualityProfileApi.md#deletequalityprofile) | **Delete** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**GetQualityProfileById**](radarr/docs/QualityProfileApi.md#getqualityprofilebyid) | **Get** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ListQualityProfile**](radarr/docs/QualityProfileApi.md#listqualityprofile) | **Get** /api/v3/qualityprofile | 
*QualityProfileApi* | [**UpdateQualityProfile**](radarr/docs/QualityProfileApi.md#updatequalityprofile) | **Put** /api/v3/qualityprofile/{id} | 
*QualityProfileSchemaApi* | [**GetQualityprofileSchema**](radarr/docs/QualityProfileSchemaApi.md#getqualityprofileschema) | **Get** /api/v3/qualityprofile/schema | 
*QueueApi* | [**DeleteQueue**](radarr/docs/QueueApi.md#deletequeue) | **Delete** /api/v3/queue/{id} | 
*QueueApi* | [**DeleteQueueBulk**](radarr/docs/QueueApi.md#deletequeuebulk) | **Delete** /api/v3/queue/bulk | 
*QueueApi* | [**GetQueue**](radarr/docs/QueueApi.md#getqueue) | **Get** /api/v3/queue | 
*QueueApi* | [**GetQueueById**](radarr/docs/QueueApi.md#getqueuebyid) | **Get** /api/v3/queue/{id} | 
*QueueActionApi* | [**CreateQueueGrabBulk**](radarr/docs/QueueActionApi.md#createqueuegrabbulk) | **Post** /api/v3/queue/grab/bulk | 
*QueueActionApi* | [**CreateQueueGrabById**](radarr/docs/QueueActionApi.md#createqueuegrabbyid) | **Post** /api/v3/queue/grab/{id} | 
*QueueDetailsApi* | [**GetQueueDetailsById**](radarr/docs/QueueDetailsApi.md#getqueuedetailsbyid) | **Get** /api/v3/queue/details/{id} | 
*QueueDetailsApi* | [**ListQueueDetails**](radarr/docs/QueueDetailsApi.md#listqueuedetails) | **Get** /api/v3/queue/details | 
*QueueStatusApi* | [**GetQueueStatus**](radarr/docs/QueueStatusApi.md#getqueuestatus) | **Get** /api/v3/queue/status | 
*QueueStatusApi* | [**GetQueueStatusById**](radarr/docs/QueueStatusApi.md#getqueuestatusbyid) | **Get** /api/v3/queue/status/{id} | 
*ReleaseApi* | [**CreateRelease**](radarr/docs/ReleaseApi.md#createrelease) | **Post** /api/v3/release | 
*ReleaseApi* | [**GetReleaseById**](radarr/docs/ReleaseApi.md#getreleasebyid) | **Get** /api/v3/release/{id} | 
*ReleaseApi* | [**ListRelease**](radarr/docs/ReleaseApi.md#listrelease) | **Get** /api/v3/release | 
*ReleasePushApi* | [**CreateReleasePush**](radarr/docs/ReleasePushApi.md#createreleasepush) | **Post** /api/v3/release/push | 
*ReleasePushApi* | [**GetReleasePushById**](radarr/docs/ReleasePushApi.md#getreleasepushbyid) | **Get** /api/v3/release/push/{id} | 
*RemotePathMappingApi* | [**CreateRemotePathMapping**](radarr/docs/RemotePathMappingApi.md#createremotepathmapping) | **Post** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**DeleteRemotePathMapping**](radarr/docs/RemotePathMappingApi.md#deleteremotepathmapping) | **Delete** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**GetRemotePathMappingById**](radarr/docs/RemotePathMappingApi.md#getremotepathmappingbyid) | **Get** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ListRemotePathMapping**](radarr/docs/RemotePathMappingApi.md#listremotepathmapping) | **Get** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**UpdateRemotePathMapping**](radarr/docs/RemotePathMappingApi.md#updateremotepathmapping) | **Put** /api/v3/remotepathmapping/{id} | 
*RenameMovieApi* | [**ListRename**](radarr/docs/RenameMovieApi.md#listrename) | **Get** /api/v3/rename | 
*RestrictionApi* | [**CreateRestriction**](radarr/docs/RestrictionApi.md#createrestriction) | **Post** /api/v3/restriction | 
*RestrictionApi* | [**DeleteRestriction**](radarr/docs/RestrictionApi.md#deleterestriction) | **Delete** /api/v3/restriction/{id} | 
*RestrictionApi* | [**GetRestrictionById**](radarr/docs/RestrictionApi.md#getrestrictionbyid) | **Get** /api/v3/restriction/{id} | 
*RestrictionApi* | [**ListRestriction**](radarr/docs/RestrictionApi.md#listrestriction) | **Get** /api/v3/restriction | 
*RestrictionApi* | [**UpdateRestriction**](radarr/docs/RestrictionApi.md#updaterestriction) | **Put** /api/v3/restriction/{id} | 
*RootFolderApi* | [**CreateRootFolder**](radarr/docs/RootFolderApi.md#createrootfolder) | **Post** /api/v3/rootfolder | 
*RootFolderApi* | [**DeleteRootFolder**](radarr/docs/RootFolderApi.md#deleterootfolder) | **Delete** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**GetRootFolderById**](radarr/docs/RootFolderApi.md#getrootfolderbyid) | **Get** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**ListRootFolder**](radarr/docs/RootFolderApi.md#listrootfolder) | **Get** /api/v3/rootfolder | 
*StaticResourceApi* | [**Get**](radarr/docs/StaticResourceApi.md#get) | **Get** / | 
*StaticResourceApi* | [**GetByPath**](radarr/docs/StaticResourceApi.md#getbypath) | **Get** /{path} | 
*StaticResourceApi* | [**GetContentByPath**](radarr/docs/StaticResourceApi.md#getcontentbypath) | **Get** /content/{path} | 
*StaticResourceApi* | [**GetLogin**](radarr/docs/StaticResourceApi.md#getlogin) | **Get** /login | 
*SystemApi* | [**CreateSystemRestart**](radarr/docs/SystemApi.md#createsystemrestart) | **Post** /api/v3/system/restart | 
*SystemApi* | [**CreateSystemShutdown**](radarr/docs/SystemApi.md#createsystemshutdown) | **Post** /api/v3/system/shutdown | 
*SystemApi* | [**GetSystemRoutes**](radarr/docs/SystemApi.md#getsystemroutes) | **Get** /api/v3/system/routes | 
*SystemApi* | [**GetSystemRoutesDuplicate**](radarr/docs/SystemApi.md#getsystemroutesduplicate) | **Get** /api/v3/system/routes/duplicate | 
*SystemApi* | [**GetSystemStatus**](radarr/docs/SystemApi.md#getsystemstatus) | **Get** /api/v3/system/status | 
*TagApi* | [**CreateTag**](radarr/docs/TagApi.md#createtag) | **Post** /api/v3/tag | 
*TagApi* | [**DeleteTag**](radarr/docs/TagApi.md#deletetag) | **Delete** /api/v3/tag/{id} | 
*TagApi* | [**GetTagById**](radarr/docs/TagApi.md#gettagbyid) | **Get** /api/v3/tag/{id} | 
*TagApi* | [**ListTag**](radarr/docs/TagApi.md#listtag) | **Get** /api/v3/tag | 
*TagApi* | [**UpdateTag**](radarr/docs/TagApi.md#updatetag) | **Put** /api/v3/tag/{id} | 
*TagDetailsApi* | [**GetTagDetailById**](radarr/docs/TagDetailsApi.md#gettagdetailbyid) | **Get** /api/v3/tag/detail/{id} | 
*TagDetailsApi* | [**ListTagDetail**](radarr/docs/TagDetailsApi.md#listtagdetail) | **Get** /api/v3/tag/detail | 
*TaskApi* | [**GetSystemTaskById**](radarr/docs/TaskApi.md#getsystemtaskbyid) | **Get** /api/v3/system/task/{id} | 
*TaskApi* | [**ListSystemTask**](radarr/docs/TaskApi.md#listsystemtask) | **Get** /api/v3/system/task | 
*UiConfigApi* | [**GetUiConfig**](radarr/docs/UiConfigApi.md#getuiconfig) | **Get** /api/v3/config/ui | 
*UiConfigApi* | [**GetUiConfigById**](radarr/docs/UiConfigApi.md#getuiconfigbyid) | **Get** /api/v3/config/ui/{id} | 
*UiConfigApi* | [**UpdateUiConfig**](radarr/docs/UiConfigApi.md#updateuiconfig) | **Put** /api/v3/config/ui/{id} | 
*UpdateApi* | [**ListUpdate**](radarr/docs/UpdateApi.md#listupdate) | **Get** /api/v3/update | 
*UpdateLogFileApi* | [**GetLogFileUpdateByFilename**](radarr/docs/UpdateLogFileApi.md#getlogfileupdatebyfilename) | **Get** /api/v3/log/file/update/{filename} | 
*UpdateLogFileApi* | [**ListLogFileUpdate**](radarr/docs/UpdateLogFileApi.md#listlogfileupdate) | **Get** /api/v3/log/file/update | 


## Documentation For Models

 - [AddMovieMethod](docs/AddMovieMethod.md)
 - [AddMovieOptions](docs/AddMovieOptions.md)
 - [AlternativeTitle](docs/AlternativeTitle.md)
 - [AlternativeTitleResource](docs/AlternativeTitleResource.md)
 - [ApiInfoResource](docs/ApiInfoResource.md)
 - [ApplyTags](docs/ApplyTags.md)
 - [AuthenticationType](docs/AuthenticationType.md)
 - [BackupResource](docs/BackupResource.md)
 - [BackupType](docs/BackupType.md)
 - [BlocklistBulkResource](docs/BlocklistBulkResource.md)
 - [BlocklistResource](docs/BlocklistResource.md)
 - [BlocklistResourcePagingResource](docs/BlocklistResourcePagingResource.md)
 - [CertificateValidationType](docs/CertificateValidationType.md)
 - [CollectionMovieResource](docs/CollectionMovieResource.md)
 - [CollectionResource](docs/CollectionResource.md)
 - [CollectionUpdateResource](docs/CollectionUpdateResource.md)
 - [ColonReplacementFormat](docs/ColonReplacementFormat.md)
 - [Command](docs/Command.md)
 - [CommandPriority](docs/CommandPriority.md)
 - [CommandResource](docs/CommandResource.md)
 - [CommandStatus](docs/CommandStatus.md)
 - [CommandTrigger](docs/CommandTrigger.md)
 - [CreditResource](docs/CreditResource.md)
 - [CreditType](docs/CreditType.md)
 - [CustomFilterResource](docs/CustomFilterResource.md)
 - [CustomFormatResource](docs/CustomFormatResource.md)
 - [CustomFormatSpecificationSchema](docs/CustomFormatSpecificationSchema.md)
 - [DatabaseType](docs/DatabaseType.md)
 - [DelayProfileResource](docs/DelayProfileResource.md)
 - [DiskSpaceResource](docs/DiskSpaceResource.md)
 - [DownloadClientConfigResource](docs/DownloadClientConfigResource.md)
 - [DownloadClientResource](docs/DownloadClientResource.md)
 - [DownloadProtocol](docs/DownloadProtocol.md)
 - [ExtraFileResource](docs/ExtraFileResource.md)
 - [ExtraFileType](docs/ExtraFileType.md)
 - [Field](docs/Field.md)
 - [FileDateType](docs/FileDateType.md)
 - [HealthCheckResult](docs/HealthCheckResult.md)
 - [HealthResource](docs/HealthResource.md)
 - [HistoryResource](docs/HistoryResource.md)
 - [HistoryResourcePagingResource](docs/HistoryResourcePagingResource.md)
 - [HostConfigResource](docs/HostConfigResource.md)
 - [ImportExclusionsResource](docs/ImportExclusionsResource.md)
 - [ImportListConfigResource](docs/ImportListConfigResource.md)
 - [ImportListResource](docs/ImportListResource.md)
 - [ImportListType](docs/ImportListType.md)
 - [IndexerConfigResource](docs/IndexerConfigResource.md)
 - [IndexerFlagResource](docs/IndexerFlagResource.md)
 - [IndexerResource](docs/IndexerResource.md)
 - [Language](docs/Language.md)
 - [LanguageResource](docs/LanguageResource.md)
 - [LogFileResource](docs/LogFileResource.md)
 - [LogResource](docs/LogResource.md)
 - [LogResourcePagingResource](docs/LogResourcePagingResource.md)
 - [ManualImportReprocessResource](docs/ManualImportReprocessResource.md)
 - [ManualImportResource](docs/ManualImportResource.md)
 - [MediaCover](docs/MediaCover.md)
 - [MediaCoverTypes](docs/MediaCoverTypes.md)
 - [MediaInfoResource](docs/MediaInfoResource.md)
 - [MediaManagementConfigResource](docs/MediaManagementConfigResource.md)
 - [MetadataConfigResource](docs/MetadataConfigResource.md)
 - [MetadataResource](docs/MetadataResource.md)
 - [Modifier](docs/Modifier.md)
 - [MonitorTypes](docs/MonitorTypes.md)
 - [MovieCollection](docs/MovieCollection.md)
 - [MovieEditorResource](docs/MovieEditorResource.md)
 - [MovieFileListResource](docs/MovieFileListResource.md)
 - [MovieFileResource](docs/MovieFileResource.md)
 - [MovieHistoryEventType](docs/MovieHistoryEventType.md)
 - [MovieMetadata](docs/MovieMetadata.md)
 - [MovieResource](docs/MovieResource.md)
 - [MovieRuntimeFormatType](docs/MovieRuntimeFormatType.md)
 - [MovieStatusType](docs/MovieStatusType.md)
 - [MovieTranslation](docs/MovieTranslation.md)
 - [NamingConfigResource](docs/NamingConfigResource.md)
 - [NotificationResource](docs/NotificationResource.md)
 - [PagingResourceFilter](docs/PagingResourceFilter.md)
 - [ParseResource](docs/ParseResource.md)
 - [ParsedMovieInfo](docs/ParsedMovieInfo.md)
 - [PingResource](docs/PingResource.md)
 - [ProfileFormatItemResource](docs/ProfileFormatItemResource.md)
 - [ProperDownloadTypes](docs/ProperDownloadTypes.md)
 - [ProviderMessage](docs/ProviderMessage.md)
 - [ProviderMessageType](docs/ProviderMessageType.md)
 - [ProxyType](docs/ProxyType.md)
 - [Quality](docs/Quality.md)
 - [QualityDefinitionResource](docs/QualityDefinitionResource.md)
 - [QualityModel](docs/QualityModel.md)
 - [QualityProfileQualityItemResource](docs/QualityProfileQualityItemResource.md)
 - [QualityProfileResource](docs/QualityProfileResource.md)
 - [QueueBulkResource](docs/QueueBulkResource.md)
 - [QueueResource](docs/QueueResource.md)
 - [QueueResourcePagingResource](docs/QueueResourcePagingResource.md)
 - [QueueStatusResource](docs/QueueStatusResource.md)
 - [RatingChild](docs/RatingChild.md)
 - [RatingType](docs/RatingType.md)
 - [Ratings](docs/Ratings.md)
 - [Rejection](docs/Rejection.md)
 - [RejectionType](docs/RejectionType.md)
 - [ReleaseResource](docs/ReleaseResource.md)
 - [RemotePathMappingResource](docs/RemotePathMappingResource.md)
 - [RenameMovieResource](docs/RenameMovieResource.md)
 - [RescanAfterRefreshType](docs/RescanAfterRefreshType.md)
 - [RestrictionResource](docs/RestrictionResource.md)
 - [Revision](docs/Revision.md)
 - [RootFolderResource](docs/RootFolderResource.md)
 - [RuntimeMode](docs/RuntimeMode.md)
 - [SelectOption](docs/SelectOption.md)
 - [SortDirection](docs/SortDirection.md)
 - [Source](docs/Source.md)
 - [SourceType](docs/SourceType.md)
 - [SystemResource](docs/SystemResource.md)
 - [TMDbCountryCode](docs/TMDbCountryCode.md)
 - [TagDetailsResource](docs/TagDetailsResource.md)
 - [TagResource](docs/TagResource.md)
 - [TaskResource](docs/TaskResource.md)
 - [TrackedDownloadState](docs/TrackedDownloadState.md)
 - [TrackedDownloadStatus](docs/TrackedDownloadStatus.md)
 - [TrackedDownloadStatusMessage](docs/TrackedDownloadStatusMessage.md)
 - [UiConfigResource](docs/UiConfigResource.md)
 - [UnmappedFolder](docs/UnmappedFolder.md)
 - [UpdateChanges](docs/UpdateChanges.md)
 - [UpdateMechanism](docs/UpdateMechanism.md)
 - [UpdateResource](docs/UpdateResource.md)


## Documentation For Authorization



### X-Api-Key

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Api-Key and passed in as the auth context for each request.


### apikey

- **Type**: API key
- **API key parameter name**: apikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apikey and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




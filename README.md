# Go API client for radarr

Radarr API docs

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 0.0.1
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
*AlternativeTitleApi* | [**GetAlttitleById**](docs/AlternativeTitleApi.md#getalttitlebyid) | **Get** /api/v3/alttitle/{id} | 
*AlternativeTitleApi* | [**ListAlttitle**](docs/AlternativeTitleApi.md#listalttitle) | **Get** /api/v3/alttitle | 
*ApiInfoApi* | [**GetApi**](docs/ApiInfoApi.md#getapi) | **Get** /api | 
*AuthenticationApi* | [**CreateLogin**](docs/AuthenticationApi.md#createlogin) | **Post** /login | 
*AuthenticationApi* | [**GetLogout**](docs/AuthenticationApi.md#getlogout) | **Get** /logout | 
*BackupApi* | [**CreateSystemBackupRestoreById**](docs/BackupApi.md#createsystembackuprestorebyid) | **Post** /api/v3/system/backup/restore/{id} | 
*BackupApi* | [**CreateSystemBackupRestoreUpload**](docs/BackupApi.md#createsystembackuprestoreupload) | **Post** /api/v3/system/backup/restore/upload | 
*BackupApi* | [**DeleteSystemBackup**](docs/BackupApi.md#deletesystembackup) | **Delete** /api/v3/system/backup/{id} | 
*BackupApi* | [**ListSystemBackup**](docs/BackupApi.md#listsystembackup) | **Get** /api/v3/system/backup | 
*BlocklistApi* | [**DeleteBlocklist**](docs/BlocklistApi.md#deleteblocklist) | **Delete** /api/v3/blocklist/{id} | 
*BlocklistApi* | [**DeleteBlocklistBulk**](docs/BlocklistApi.md#deleteblocklistbulk) | **Delete** /api/v3/blocklist/bulk | 
*BlocklistApi* | [**GetBlocklist**](docs/BlocklistApi.md#getblocklist) | **Get** /api/v3/blocklist | 
*BlocklistApi* | [**ListBlocklistMovie**](docs/BlocklistApi.md#listblocklistmovie) | **Get** /api/v3/blocklist/movie | 
*CalendarApi* | [**GetCalendarById**](docs/CalendarApi.md#getcalendarbyid) | **Get** /api/v3/calendar/{id} | 
*CalendarApi* | [**ListCalendar**](docs/CalendarApi.md#listcalendar) | **Get** /api/v3/calendar | 
*CalendarFeedApi* | [**GetFeedV3CalendarRadarrIcs**](docs/CalendarFeedApi.md#getfeedv3calendarradarrics) | **Get** /feed/v3/calendar/radarr.ics | 
*CollectionApi* | [**GetCollectionById**](docs/CollectionApi.md#getcollectionbyid) | **Get** /api/v3/collection/{id} | 
*CollectionApi* | [**ListCollection**](docs/CollectionApi.md#listcollection) | **Get** /api/v3/collection | 
*CollectionApi* | [**PutCollection**](docs/CollectionApi.md#putcollection) | **Put** /api/v3/collection | 
*CollectionApi* | [**UpdateCollection**](docs/CollectionApi.md#updatecollection) | **Put** /api/v3/collection/{id} | 
*CommandApi* | [**CreateCommand**](docs/CommandApi.md#createcommand) | **Post** /api/v3/command | 
*CommandApi* | [**DeleteCommand**](docs/CommandApi.md#deletecommand) | **Delete** /api/v3/command/{id} | 
*CommandApi* | [**GetCommandById**](docs/CommandApi.md#getcommandbyid) | **Get** /api/v3/command/{id} | 
*CommandApi* | [**ListCommand**](docs/CommandApi.md#listcommand) | **Get** /api/v3/command | 
*CreditApi* | [**GetCreditById**](docs/CreditApi.md#getcreditbyid) | **Get** /api/v3/credit/{id} | 
*CreditApi* | [**ListCredit**](docs/CreditApi.md#listcredit) | **Get** /api/v3/credit | 
*CustomFilterApi* | [**CreateCustomFilter**](docs/CustomFilterApi.md#createcustomfilter) | **Post** /api/v3/customfilter | 
*CustomFilterApi* | [**DeleteCustomFilter**](docs/CustomFilterApi.md#deletecustomfilter) | **Delete** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**GetCustomFilterById**](docs/CustomFilterApi.md#getcustomfilterbyid) | **Get** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ListCustomFilter**](docs/CustomFilterApi.md#listcustomfilter) | **Get** /api/v3/customfilter | 
*CustomFilterApi* | [**UpdateCustomFilter**](docs/CustomFilterApi.md#updatecustomfilter) | **Put** /api/v3/customfilter/{id} | 
*CustomFormatApi* | [**CreateCustomFormat**](docs/CustomFormatApi.md#createcustomformat) | **Post** /api/v3/customformat | 
*CustomFormatApi* | [**DeleteCustomFormat**](docs/CustomFormatApi.md#deletecustomformat) | **Delete** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**GetCustomFormatById**](docs/CustomFormatApi.md#getcustomformatbyid) | **Get** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**GetCustomFormatSchema**](docs/CustomFormatApi.md#getcustomformatschema) | **Get** /api/v3/customformat/schema | 
*CustomFormatApi* | [**ListCustomFormat**](docs/CustomFormatApi.md#listcustomformat) | **Get** /api/v3/customformat | 
*CustomFormatApi* | [**UpdateCustomFormat**](docs/CustomFormatApi.md#updatecustomformat) | **Put** /api/v3/customformat/{id} | 
*DelayProfileApi* | [**CreateDelayProfile**](docs/DelayProfileApi.md#createdelayprofile) | **Post** /api/v3/delayprofile | 
*DelayProfileApi* | [**DeleteDelayProfile**](docs/DelayProfileApi.md#deletedelayprofile) | **Delete** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**GetDelayProfileById**](docs/DelayProfileApi.md#getdelayprofilebyid) | **Get** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ListDelayProfile**](docs/DelayProfileApi.md#listdelayprofile) | **Get** /api/v3/delayprofile | 
*DelayProfileApi* | [**UpdateDelayProfile**](docs/DelayProfileApi.md#updatedelayprofile) | **Put** /api/v3/delayprofile/{id} | 
*DiskSpaceApi* | [**ListDiskSpace**](docs/DiskSpaceApi.md#listdiskspace) | **Get** /api/v3/diskspace | 
*DownloadClientApi* | [**CreateDownloadClient**](docs/DownloadClientApi.md#createdownloadclient) | **Post** /api/v3/downloadclient | 
*DownloadClientApi* | [**CreateDownloadClientActionByName**](docs/DownloadClientApi.md#createdownloadclientactionbyname) | **Post** /api/v3/downloadclient/action/{name} | 
*DownloadClientApi* | [**DeleteDownloadClient**](docs/DownloadClientApi.md#deletedownloadclient) | **Delete** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**GetDownloadClientById**](docs/DownloadClientApi.md#getdownloadclientbyid) | **Get** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ListDownloadClient**](docs/DownloadClientApi.md#listdownloadclient) | **Get** /api/v3/downloadclient | 
*DownloadClientApi* | [**ListDownloadClientSchema**](docs/DownloadClientApi.md#listdownloadclientschema) | **Get** /api/v3/downloadclient/schema | 
*DownloadClientApi* | [**TestDownloadClient**](docs/DownloadClientApi.md#testdownloadclient) | **Post** /api/v3/downloadclient/test | 
*DownloadClientApi* | [**TestallDownloadClient**](docs/DownloadClientApi.md#testalldownloadclient) | **Post** /api/v3/downloadclient/testall | 
*DownloadClientApi* | [**UpdateDownloadClient**](docs/DownloadClientApi.md#updatedownloadclient) | **Put** /api/v3/downloadclient/{id} | 
*DownloadClientConfigApi* | [**GetDownloadClientConfig**](docs/DownloadClientConfigApi.md#getdownloadclientconfig) | **Get** /api/v3/config/downloadclient | 
*DownloadClientConfigApi* | [**GetDownloadClientConfigById**](docs/DownloadClientConfigApi.md#getdownloadclientconfigbyid) | **Get** /api/v3/config/downloadclient/{id} | 
*DownloadClientConfigApi* | [**UpdateDownloadClientConfig**](docs/DownloadClientConfigApi.md#updatedownloadclientconfig) | **Put** /api/v3/config/downloadclient/{id} | 
*ExtraFileApi* | [**ListExtraFile**](docs/ExtraFileApi.md#listextrafile) | **Get** /api/v3/extrafile | 
*FileSystemApi* | [**GetFileSystem**](docs/FileSystemApi.md#getfilesystem) | **Get** /api/v3/filesystem | 
*FileSystemApi* | [**GetFileSystemMediafiles**](docs/FileSystemApi.md#getfilesystemmediafiles) | **Get** /api/v3/filesystem/mediafiles | 
*FileSystemApi* | [**GetFileSystemType**](docs/FileSystemApi.md#getfilesystemtype) | **Get** /api/v3/filesystem/type | 
*HealthApi* | [**GetHealthById**](docs/HealthApi.md#gethealthbyid) | **Get** /api/v3/health/{id} | 
*HealthApi* | [**ListHealth**](docs/HealthApi.md#listhealth) | **Get** /api/v3/health | 
*HistoryApi* | [**CreateHistoryFailedById**](docs/HistoryApi.md#createhistoryfailedbyid) | **Post** /api/v3/history/failed/{id} | 
*HistoryApi* | [**GetHistory**](docs/HistoryApi.md#gethistory) | **Get** /api/v3/history | 
*HistoryApi* | [**ListHistoryMovie**](docs/HistoryApi.md#listhistorymovie) | **Get** /api/v3/history/movie | 
*HistoryApi* | [**ListHistorySince**](docs/HistoryApi.md#listhistorysince) | **Get** /api/v3/history/since | 
*HostConfigApi* | [**GetHostConfig**](docs/HostConfigApi.md#gethostconfig) | **Get** /api/v3/config/host | 
*HostConfigApi* | [**GetHostConfigById**](docs/HostConfigApi.md#gethostconfigbyid) | **Get** /api/v3/config/host/{id} | 
*HostConfigApi* | [**UpdateHostConfig**](docs/HostConfigApi.md#updatehostconfig) | **Put** /api/v3/config/host/{id} | 
*ImportExclusionsApi* | [**CreateExclusions**](docs/ImportExclusionsApi.md#createexclusions) | **Post** /api/v3/exclusions | 
*ImportExclusionsApi* | [**CreateExclusionsBulk**](docs/ImportExclusionsApi.md#createexclusionsbulk) | **Post** /api/v3/exclusions/bulk | 
*ImportExclusionsApi* | [**DeleteExclusions**](docs/ImportExclusionsApi.md#deleteexclusions) | **Delete** /api/v3/exclusions/{id} | 
*ImportExclusionsApi* | [**GetExclusionsById**](docs/ImportExclusionsApi.md#getexclusionsbyid) | **Get** /api/v3/exclusions/{id} | 
*ImportExclusionsApi* | [**ListExclusions**](docs/ImportExclusionsApi.md#listexclusions) | **Get** /api/v3/exclusions | 
*ImportExclusionsApi* | [**UpdateExclusions**](docs/ImportExclusionsApi.md#updateexclusions) | **Put** /api/v3/exclusions/{id} | 
*ImportListApi* | [**CreateImportList**](docs/ImportListApi.md#createimportlist) | **Post** /api/v3/importlist | 
*ImportListApi* | [**CreateImportListActionByName**](docs/ImportListApi.md#createimportlistactionbyname) | **Post** /api/v3/importlist/action/{name} | 
*ImportListApi* | [**DeleteImportList**](docs/ImportListApi.md#deleteimportlist) | **Delete** /api/v3/importlist/{id} | 
*ImportListApi* | [**GetImportListById**](docs/ImportListApi.md#getimportlistbyid) | **Get** /api/v3/importlist/{id} | 
*ImportListApi* | [**ListImportList**](docs/ImportListApi.md#listimportlist) | **Get** /api/v3/importlist | 
*ImportListApi* | [**ListImportListSchema**](docs/ImportListApi.md#listimportlistschema) | **Get** /api/v3/importlist/schema | 
*ImportListApi* | [**TestImportList**](docs/ImportListApi.md#testimportlist) | **Post** /api/v3/importlist/test | 
*ImportListApi* | [**TestallImportList**](docs/ImportListApi.md#testallimportlist) | **Post** /api/v3/importlist/testall | 
*ImportListApi* | [**UpdateImportList**](docs/ImportListApi.md#updateimportlist) | **Put** /api/v3/importlist/{id} | 
*ImportListConfigApi* | [**GetImportListConfig**](docs/ImportListConfigApi.md#getimportlistconfig) | **Get** /api/v3/config/importlist | 
*ImportListConfigApi* | [**GetImportListConfigById**](docs/ImportListConfigApi.md#getimportlistconfigbyid) | **Get** /api/v3/config/importlist/{id} | 
*ImportListConfigApi* | [**UpdateImportListConfig**](docs/ImportListConfigApi.md#updateimportlistconfig) | **Put** /api/v3/config/importlist/{id} | 
*ImportListMoviesApi* | [**CreateImportlistMovie**](docs/ImportListMoviesApi.md#createimportlistmovie) | **Post** /api/v3/importlist/movie | 
*ImportListMoviesApi* | [**GetImportlistMovie**](docs/ImportListMoviesApi.md#getimportlistmovie) | **Get** /api/v3/importlist/movie | 
*IndexerApi* | [**CreateIndexer**](docs/IndexerApi.md#createindexer) | **Post** /api/v3/indexer | 
*IndexerApi* | [**CreateIndexerActionByName**](docs/IndexerApi.md#createindexeractionbyname) | **Post** /api/v3/indexer/action/{name} | 
*IndexerApi* | [**DeleteIndexer**](docs/IndexerApi.md#deleteindexer) | **Delete** /api/v3/indexer/{id} | 
*IndexerApi* | [**GetIndexerById**](docs/IndexerApi.md#getindexerbyid) | **Get** /api/v3/indexer/{id} | 
*IndexerApi* | [**ListIndexer**](docs/IndexerApi.md#listindexer) | **Get** /api/v3/indexer | 
*IndexerApi* | [**ListIndexerSchema**](docs/IndexerApi.md#listindexerschema) | **Get** /api/v3/indexer/schema | 
*IndexerApi* | [**TestIndexer**](docs/IndexerApi.md#testindexer) | **Post** /api/v3/indexer/test | 
*IndexerApi* | [**TestallIndexer**](docs/IndexerApi.md#testallindexer) | **Post** /api/v3/indexer/testall | 
*IndexerApi* | [**UpdateIndexer**](docs/IndexerApi.md#updateindexer) | **Put** /api/v3/indexer/{id} | 
*IndexerConfigApi* | [**GetIndexerConfig**](docs/IndexerConfigApi.md#getindexerconfig) | **Get** /api/v3/config/indexer | 
*IndexerConfigApi* | [**GetIndexerConfigById**](docs/IndexerConfigApi.md#getindexerconfigbyid) | **Get** /api/v3/config/indexer/{id} | 
*IndexerConfigApi* | [**UpdateIndexerConfig**](docs/IndexerConfigApi.md#updateindexerconfig) | **Put** /api/v3/config/indexer/{id} | 
*IndexerFlagApi* | [**ListIndexerFlag**](docs/IndexerFlagApi.md#listindexerflag) | **Get** /api/v3/indexerflag | 
*InitializeJsApi* | [**GetInitializeJs**](docs/InitializeJsApi.md#getinitializejs) | **Get** /initialize.js | 
*LanguageApi* | [**GetLanguageById**](docs/LanguageApi.md#getlanguagebyid) | **Get** /api/v3/language/{id} | 
*LanguageApi* | [**ListLanguage**](docs/LanguageApi.md#listlanguage) | **Get** /api/v3/language | 
*LocalizationApi* | [**GetLocalization**](docs/LocalizationApi.md#getlocalization) | **Get** /api/v3/localization | 
*LogApi* | [**GetLog**](docs/LogApi.md#getlog) | **Get** /api/v3/log | 
*LogFileApi* | [**GetLogFileByFilename**](docs/LogFileApi.md#getlogfilebyfilename) | **Get** /api/v3/log/file/{filename} | 
*LogFileApi* | [**ListLogFile**](docs/LogFileApi.md#listlogfile) | **Get** /api/v3/log/file | 
*ManualImportApi* | [**CreateManualImport**](docs/ManualImportApi.md#createmanualimport) | **Post** /api/v3/manualimport | 
*ManualImportApi* | [**ListManualImport**](docs/ManualImportApi.md#listmanualimport) | **Get** /api/v3/manualimport | 
*MediaCoverApi* | [**GetMediaCovermovieIdByFilename**](docs/MediaCoverApi.md#getmediacovermovieidbyfilename) | **Get** /api/v3/mediacover/{movieId}/{filename} | 
*MediaManagementConfigApi* | [**GetMediaManagementConfig**](docs/MediaManagementConfigApi.md#getmediamanagementconfig) | **Get** /api/v3/config/mediamanagement | 
*MediaManagementConfigApi* | [**GetMediaManagementConfigById**](docs/MediaManagementConfigApi.md#getmediamanagementconfigbyid) | **Get** /api/v3/config/mediamanagement/{id} | 
*MediaManagementConfigApi* | [**UpdateMediaManagementConfig**](docs/MediaManagementConfigApi.md#updatemediamanagementconfig) | **Put** /api/v3/config/mediamanagement/{id} | 
*MetadataApi* | [**CreateMetadata**](docs/MetadataApi.md#createmetadata) | **Post** /api/v3/metadata | 
*MetadataApi* | [**CreateMetadataActionByName**](docs/MetadataApi.md#createmetadataactionbyname) | **Post** /api/v3/metadata/action/{name} | 
*MetadataApi* | [**DeleteMetadata**](docs/MetadataApi.md#deletemetadata) | **Delete** /api/v3/metadata/{id} | 
*MetadataApi* | [**GetMetadataById**](docs/MetadataApi.md#getmetadatabyid) | **Get** /api/v3/metadata/{id} | 
*MetadataApi* | [**ListMetadata**](docs/MetadataApi.md#listmetadata) | **Get** /api/v3/metadata | 
*MetadataApi* | [**ListMetadataSchema**](docs/MetadataApi.md#listmetadataschema) | **Get** /api/v3/metadata/schema | 
*MetadataApi* | [**TestMetadata**](docs/MetadataApi.md#testmetadata) | **Post** /api/v3/metadata/test | 
*MetadataApi* | [**TestallMetadata**](docs/MetadataApi.md#testallmetadata) | **Post** /api/v3/metadata/testall | 
*MetadataApi* | [**UpdateMetadata**](docs/MetadataApi.md#updatemetadata) | **Put** /api/v3/metadata/{id} | 
*MetadataConfigApi* | [**GetMetadataConfig**](docs/MetadataConfigApi.md#getmetadataconfig) | **Get** /api/v3/config/metadata | 
*MetadataConfigApi* | [**GetMetadataConfigById**](docs/MetadataConfigApi.md#getmetadataconfigbyid) | **Get** /api/v3/config/metadata/{id} | 
*MetadataConfigApi* | [**UpdateMetadataConfig**](docs/MetadataConfigApi.md#updatemetadataconfig) | **Put** /api/v3/config/metadata/{id} | 
*MovieApi* | [**CreateMovie**](docs/MovieApi.md#createmovie) | **Post** /api/v3/movie | 
*MovieApi* | [**DeleteMovie**](docs/MovieApi.md#deletemovie) | **Delete** /api/v3/movie/{id} | 
*MovieApi* | [**GetMovieById**](docs/MovieApi.md#getmoviebyid) | **Get** /api/v3/movie/{id} | 
*MovieApi* | [**ListMovie**](docs/MovieApi.md#listmovie) | **Get** /api/v3/movie | 
*MovieApi* | [**UpdateMovie**](docs/MovieApi.md#updatemovie) | **Put** /api/v3/movie/{id} | 
*MovieEditorApi* | [**DeleteMovieEditor**](docs/MovieEditorApi.md#deletemovieeditor) | **Delete** /api/v3/movie/editor | 
*MovieEditorApi* | [**PutMovieEditor**](docs/MovieEditorApi.md#putmovieeditor) | **Put** /api/v3/movie/editor | 
*MovieFileApi* | [**DeleteMovieFile**](docs/MovieFileApi.md#deletemoviefile) | **Delete** /api/v3/moviefile/{id} | 
*MovieFileApi* | [**DeleteMovieFileBulk**](docs/MovieFileApi.md#deletemoviefilebulk) | **Delete** /api/v3/moviefile/bulk | 
*MovieFileApi* | [**GetMovieFileById**](docs/MovieFileApi.md#getmoviefilebyid) | **Get** /api/v3/moviefile/{id} | 
*MovieFileApi* | [**ListMovieFile**](docs/MovieFileApi.md#listmoviefile) | **Get** /api/v3/moviefile | 
*MovieFileApi* | [**PutMovieFileEditor**](docs/MovieFileApi.md#putmoviefileeditor) | **Put** /api/v3/moviefile/editor | 
*MovieFileApi* | [**UpdateMovieFile**](docs/MovieFileApi.md#updatemoviefile) | **Put** /api/v3/moviefile/{id} | 
*MovieImportApi* | [**CreateMovieImport**](docs/MovieImportApi.md#createmovieimport) | **Post** /api/v3/movie/import | 
*MovieImportApi* | [**GetMovieImportById**](docs/MovieImportApi.md#getmovieimportbyid) | **Get** /api/v3/movie/import/{id} | 
*MovieLookupApi* | [**GetMovieLookup**](docs/MovieLookupApi.md#getmovielookup) | **Get** /api/v3/movie/lookup | 
*MovieLookupApi* | [**GetMovieLookupById**](docs/MovieLookupApi.md#getmovielookupbyid) | **Get** /api/v3/movie/lookup/{id} | 
*MovieLookupApi* | [**GetMovieLookupImdb**](docs/MovieLookupApi.md#getmovielookupimdb) | **Get** /api/v3/movie/lookup/imdb | 
*MovieLookupApi* | [**GetMovieLookupTmdb**](docs/MovieLookupApi.md#getmovielookuptmdb) | **Get** /api/v3/movie/lookup/tmdb | 
*NamingConfigApi* | [**GetNamingConfig**](docs/NamingConfigApi.md#getnamingconfig) | **Get** /api/v3/config/naming | 
*NamingConfigApi* | [**GetNamingConfigById**](docs/NamingConfigApi.md#getnamingconfigbyid) | **Get** /api/v3/config/naming/{id} | 
*NamingConfigApi* | [**GetNamingConfigExamples**](docs/NamingConfigApi.md#getnamingconfigexamples) | **Get** /api/v3/config/naming/examples | 
*NamingConfigApi* | [**UpdateNamingConfig**](docs/NamingConfigApi.md#updatenamingconfig) | **Put** /api/v3/config/naming/{id} | 
*NotificationApi* | [**CreateNotification**](docs/NotificationApi.md#createnotification) | **Post** /api/v3/notification | 
*NotificationApi* | [**CreateNotificationActionByName**](docs/NotificationApi.md#createnotificationactionbyname) | **Post** /api/v3/notification/action/{name} | 
*NotificationApi* | [**DeleteNotification**](docs/NotificationApi.md#deletenotification) | **Delete** /api/v3/notification/{id} | 
*NotificationApi* | [**GetNotificationById**](docs/NotificationApi.md#getnotificationbyid) | **Get** /api/v3/notification/{id} | 
*NotificationApi* | [**ListNotification**](docs/NotificationApi.md#listnotification) | **Get** /api/v3/notification | 
*NotificationApi* | [**ListNotificationSchema**](docs/NotificationApi.md#listnotificationschema) | **Get** /api/v3/notification/schema | 
*NotificationApi* | [**TestNotification**](docs/NotificationApi.md#testnotification) | **Post** /api/v3/notification/test | 
*NotificationApi* | [**TestallNotification**](docs/NotificationApi.md#testallnotification) | **Post** /api/v3/notification/testall | 
*NotificationApi* | [**UpdateNotification**](docs/NotificationApi.md#updatenotification) | **Put** /api/v3/notification/{id} | 
*ParseApi* | [**GetParse**](docs/ParseApi.md#getparse) | **Get** /api/v3/parse | 
*PingApi* | [**GetPing**](docs/PingApi.md#getping) | **Get** /ping | 
*QualityDefinitionApi* | [**GetQualityDefinitionById**](docs/QualityDefinitionApi.md#getqualitydefinitionbyid) | **Get** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionApi* | [**ListQualityDefinition**](docs/QualityDefinitionApi.md#listqualitydefinition) | **Get** /api/v3/qualitydefinition | 
*QualityDefinitionApi* | [**PutQualityDefinitionUpdate**](docs/QualityDefinitionApi.md#putqualitydefinitionupdate) | **Put** /api/v3/qualitydefinition/update | 
*QualityDefinitionApi* | [**UpdateQualityDefinition**](docs/QualityDefinitionApi.md#updatequalitydefinition) | **Put** /api/v3/qualitydefinition/{id} | 
*QualityProfileApi* | [**CreateQualityProfile**](docs/QualityProfileApi.md#createqualityprofile) | **Post** /api/v3/qualityprofile | 
*QualityProfileApi* | [**DeleteQualityProfile**](docs/QualityProfileApi.md#deletequalityprofile) | **Delete** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**GetQualityProfileById**](docs/QualityProfileApi.md#getqualityprofilebyid) | **Get** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ListQualityProfile**](docs/QualityProfileApi.md#listqualityprofile) | **Get** /api/v3/qualityprofile | 
*QualityProfileApi* | [**UpdateQualityProfile**](docs/QualityProfileApi.md#updatequalityprofile) | **Put** /api/v3/qualityprofile/{id} | 
*QualityProfileSchemaApi* | [**GetQualityprofileSchema**](docs/QualityProfileSchemaApi.md#getqualityprofileschema) | **Get** /api/v3/qualityprofile/schema | 
*QueueApi* | [**DeleteQueue**](docs/QueueApi.md#deletequeue) | **Delete** /api/v3/queue/{id} | 
*QueueApi* | [**DeleteQueueBulk**](docs/QueueApi.md#deletequeuebulk) | **Delete** /api/v3/queue/bulk | 
*QueueApi* | [**GetQueue**](docs/QueueApi.md#getqueue) | **Get** /api/v3/queue | 
*QueueApi* | [**GetQueueById**](docs/QueueApi.md#getqueuebyid) | **Get** /api/v3/queue/{id} | 
*QueueActionApi* | [**CreateQueueGrabBulk**](docs/QueueActionApi.md#createqueuegrabbulk) | **Post** /api/v3/queue/grab/bulk | 
*QueueActionApi* | [**CreateQueueGrabById**](docs/QueueActionApi.md#createqueuegrabbyid) | **Post** /api/v3/queue/grab/{id} | 
*QueueDetailsApi* | [**GetQueueDetailsById**](docs/QueueDetailsApi.md#getqueuedetailsbyid) | **Get** /api/v3/queue/details/{id} | 
*QueueDetailsApi* | [**ListQueueDetails**](docs/QueueDetailsApi.md#listqueuedetails) | **Get** /api/v3/queue/details | 
*QueueStatusApi* | [**GetQueueStatus**](docs/QueueStatusApi.md#getqueuestatus) | **Get** /api/v3/queue/status | 
*QueueStatusApi* | [**GetQueueStatusById**](docs/QueueStatusApi.md#getqueuestatusbyid) | **Get** /api/v3/queue/status/{id} | 
*ReleaseApi* | [**CreateRelease**](docs/ReleaseApi.md#createrelease) | **Post** /api/v3/release | 
*ReleaseApi* | [**GetReleaseById**](docs/ReleaseApi.md#getreleasebyid) | **Get** /api/v3/release/{id} | 
*ReleaseApi* | [**ListRelease**](docs/ReleaseApi.md#listrelease) | **Get** /api/v3/release | 
*ReleasePushApi* | [**CreateReleasePush**](docs/ReleasePushApi.md#createreleasepush) | **Post** /api/v3/release/push | 
*ReleasePushApi* | [**GetReleasePushById**](docs/ReleasePushApi.md#getreleasepushbyid) | **Get** /api/v3/release/push/{id} | 
*RemotePathMappingApi* | [**CreateRemotePathMapping**](docs/RemotePathMappingApi.md#createremotepathmapping) | **Post** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**DeleteRemotePathMapping**](docs/RemotePathMappingApi.md#deleteremotepathmapping) | **Delete** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**GetRemotePathMappingById**](docs/RemotePathMappingApi.md#getremotepathmappingbyid) | **Get** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ListRemotePathMapping**](docs/RemotePathMappingApi.md#listremotepathmapping) | **Get** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**UpdateRemotePathMapping**](docs/RemotePathMappingApi.md#updateremotepathmapping) | **Put** /api/v3/remotepathmapping/{id} | 
*RenameMovieApi* | [**ListRename**](docs/RenameMovieApi.md#listrename) | **Get** /api/v3/rename | 
*RestrictionApi* | [**CreateRestriction**](docs/RestrictionApi.md#createrestriction) | **Post** /api/v3/restriction | 
*RestrictionApi* | [**DeleteRestriction**](docs/RestrictionApi.md#deleterestriction) | **Delete** /api/v3/restriction/{id} | 
*RestrictionApi* | [**GetRestrictionById**](docs/RestrictionApi.md#getrestrictionbyid) | **Get** /api/v3/restriction/{id} | 
*RestrictionApi* | [**ListRestriction**](docs/RestrictionApi.md#listrestriction) | **Get** /api/v3/restriction | 
*RestrictionApi* | [**UpdateRestriction**](docs/RestrictionApi.md#updaterestriction) | **Put** /api/v3/restriction/{id} | 
*RootFolderApi* | [**CreateRootFolder**](docs/RootFolderApi.md#createrootfolder) | **Post** /api/v3/rootfolder | 
*RootFolderApi* | [**DeleteRootFolder**](docs/RootFolderApi.md#deleterootfolder) | **Delete** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**GetRootFolderById**](docs/RootFolderApi.md#getrootfolderbyid) | **Get** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**ListRootFolder**](docs/RootFolderApi.md#listrootfolder) | **Get** /api/v3/rootfolder | 
*StaticResourceApi* | [**Get**](docs/StaticResourceApi.md#get) | **Get** / | 
*StaticResourceApi* | [**GetByPath**](docs/StaticResourceApi.md#getbypath) | **Get** /{path} | 
*StaticResourceApi* | [**GetContentByPath**](docs/StaticResourceApi.md#getcontentbypath) | **Get** /content/{path} | 
*StaticResourceApi* | [**GetLogin**](docs/StaticResourceApi.md#getlogin) | **Get** /login | 
*SystemApi* | [**CreateSystemRestart**](docs/SystemApi.md#createsystemrestart) | **Post** /api/v3/system/restart | 
*SystemApi* | [**CreateSystemShutdown**](docs/SystemApi.md#createsystemshutdown) | **Post** /api/v3/system/shutdown | 
*SystemApi* | [**GetSystemRoutes**](docs/SystemApi.md#getsystemroutes) | **Get** /api/v3/system/routes | 
*SystemApi* | [**GetSystemRoutesDuplicate**](docs/SystemApi.md#getsystemroutesduplicate) | **Get** /api/v3/system/routes/duplicate | 
*SystemApi* | [**GetSystemStatus**](docs/SystemApi.md#getsystemstatus) | **Get** /api/v3/system/status | 
*TagApi* | [**CreateTag**](docs/TagApi.md#createtag) | **Post** /api/v3/tag | 
*TagApi* | [**DeleteTag**](docs/TagApi.md#deletetag) | **Delete** /api/v3/tag/{id} | 
*TagApi* | [**GetTagById**](docs/TagApi.md#gettagbyid) | **Get** /api/v3/tag/{id} | 
*TagApi* | [**ListTag**](docs/TagApi.md#listtag) | **Get** /api/v3/tag | 
*TagApi* | [**UpdateTag**](docs/TagApi.md#updatetag) | **Put** /api/v3/tag/{id} | 
*TagDetailsApi* | [**GetTagDetailById**](docs/TagDetailsApi.md#gettagdetailbyid) | **Get** /api/v3/tag/detail/{id} | 
*TagDetailsApi* | [**ListTagDetail**](docs/TagDetailsApi.md#listtagdetail) | **Get** /api/v3/tag/detail | 
*TaskApi* | [**GetSystemTaskById**](docs/TaskApi.md#getsystemtaskbyid) | **Get** /api/v3/system/task/{id} | 
*TaskApi* | [**ListSystemTask**](docs/TaskApi.md#listsystemtask) | **Get** /api/v3/system/task | 
*UiConfigApi* | [**GetUiConfig**](docs/UiConfigApi.md#getuiconfig) | **Get** /api/v3/config/ui | 
*UiConfigApi* | [**GetUiConfigById**](docs/UiConfigApi.md#getuiconfigbyid) | **Get** /api/v3/config/ui/{id} | 
*UiConfigApi* | [**UpdateUiConfig**](docs/UiConfigApi.md#updateuiconfig) | **Put** /api/v3/config/ui/{id} | 
*UpdateApi* | [**ListUpdate**](docs/UpdateApi.md#listupdate) | **Get** /api/v3/update | 
*UpdateLogFileApi* | [**GetLogFileUpdateByFilename**](docs/UpdateLogFileApi.md#getlogfileupdatebyfilename) | **Get** /api/v3/log/file/update/{filename} | 
*UpdateLogFileApi* | [**ListLogFileUpdate**](docs/UpdateLogFileApi.md#listlogfileupdate) | **Get** /api/v3/log/file/update | 


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
 - [HttpUri](docs/HttpUri.md)
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
 - [TimeSpan](docs/TimeSpan.md)
 - [TrackedDownloadState](docs/TrackedDownloadState.md)
 - [TrackedDownloadStatus](docs/TrackedDownloadStatus.md)
 - [TrackedDownloadStatusMessage](docs/TrackedDownloadStatusMessage.md)
 - [UiConfigResource](docs/UiConfigResource.md)
 - [UnmappedFolder](docs/UnmappedFolder.md)
 - [UpdateChanges](docs/UpdateChanges.md)
 - [UpdateMechanism](docs/UpdateMechanism.md)
 - [UpdateResource](docs/UpdateResource.md)
 - [Version](docs/Version.md)


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




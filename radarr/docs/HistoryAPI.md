# \HistoryAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHistoryFailedById**](HistoryAPI.md#CreateHistoryFailedById) | **Post** /api/v3/history/failed/{id} | 
[**GetHistory**](HistoryAPI.md#GetHistory) | **Get** /api/v3/history | 
[**ListHistoryMovie**](HistoryAPI.md#ListHistoryMovie) | **Get** /api/v3/history/movie | 
[**ListHistorySince**](HistoryAPI.md#ListHistorySince) | **Get** /api/v3/history/since | 



## CreateHistoryFailedById

> CreateHistoryFailedById(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    radarrClient "./openapi"
)

func main() {
    id := int32(56) // int32 | 

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryAPI.CreateHistoryFailedById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.CreateHistoryFailedById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHistoryFailedByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistory

> HistoryResourcePagingResource GetHistory(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeMovie(includeMovie).EventType(eventType).DownloadId(downloadId).MovieIds(movieIds).Languages(languages).Quality(quality).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    radarrClient "./openapi"
)

func main() {
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 10)
    sortKey := "sortKey_example" // string |  (optional)
    sortDirection := radarrClient.SortDirection("default") // SortDirection |  (optional)
    includeMovie := true // bool |  (optional)
    eventType := int32(56) // int32 |  (optional)
    downloadId := "downloadId_example" // string |  (optional)
    movieIds := []int32{int32(123)} // []int32 |  (optional)
    languages := []int32{int32(123)} // []int32 |  (optional)
    quality := []int32{int32(123)} // []int32 |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryAPI.GetHistory(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).IncludeMovie(includeMovie).EventType(eventType).DownloadId(downloadId).MovieIds(movieIds).Languages(languages).Quality(quality).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.GetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistory`: HistoryResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.GetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 
 **includeMovie** | **bool** |  | 
 **eventType** | **int32** |  | 
 **downloadId** | **string** |  | 
 **movieIds** | **[]int32** |  | 
 **languages** | **[]int32** |  | 
 **quality** | **[]int32** |  | 

### Return type

[**HistoryResourcePagingResource**](HistoryResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistoryMovie

> []HistoryResource ListHistoryMovie(ctx).MovieId(movieId).EventType(eventType).IncludeMovie(includeMovie).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    radarrClient "./openapi"
)

func main() {
    movieId := int32(56) // int32 |  (optional)
    eventType := radarrClient.MovieHistoryEventType("unknown") // MovieHistoryEventType |  (optional)
    includeMovie := true // bool |  (optional) (default to false)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryAPI.ListHistoryMovie(context.Background()).MovieId(movieId).EventType(eventType).IncludeMovie(includeMovie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.ListHistoryMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoryMovie`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.ListHistoryMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistoryMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int32** |  | 
 **eventType** | [**MovieHistoryEventType**](MovieHistoryEventType.md) |  | 
 **includeMovie** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHistorySince

> []HistoryResource ListHistorySince(ctx).Date(date).EventType(eventType).IncludeMovie(includeMovie).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    radarrClient "./openapi"
)

func main() {
    date := time.Now() // time.Time |  (optional)
    eventType := radarrClient.MovieHistoryEventType("unknown") // MovieHistoryEventType |  (optional)
    includeMovie := true // bool |  (optional) (default to false)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryAPI.ListHistorySince(context.Background()).Date(date).EventType(eventType).IncludeMovie(includeMovie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryAPI.ListHistorySince``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistorySince`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryAPI.ListHistorySince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHistorySinceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** |  | 
 **eventType** | [**MovieHistoryEventType**](MovieHistoryEventType.md) |  | 
 **includeMovie** | **bool** |  | [default to false]

### Return type

[**[]HistoryResource**](HistoryResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


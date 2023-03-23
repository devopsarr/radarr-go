# \HistoryApi

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHistoryFailedById**](HistoryApi.md#CreateHistoryFailedById) | **Post** /api/v3/history/failed/{id} | 
[**GetHistory**](HistoryApi.md#GetHistory) | **Get** /api/v3/history | 
[**ListHistoryMovie**](HistoryApi.md#ListHistoryMovie) | **Get** /api/v3/history/movie | 
[**ListHistorySince**](HistoryApi.md#ListHistorySince) | **Get** /api/v3/history/since | 



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
    resp, r, err := apiClient.HistoryApi.CreateHistoryFailedById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.CreateHistoryFailedById``: %v\n", err)
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

> HistoryResourcePagingResource GetHistory(ctx).IncludeMovie(includeMovie).Execute()



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
    includeMovie := true // bool |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.HistoryApi.GetHistory(context.Background()).IncludeMovie(includeMovie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.GetHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHistory`: HistoryResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.GetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeMovie** | **bool** |  | 

### Return type

[**HistoryResourcePagingResource**](HistoryResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

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
    resp, r, err := apiClient.HistoryApi.ListHistoryMovie(context.Background()).MovieId(movieId).EventType(eventType).IncludeMovie(includeMovie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListHistoryMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistoryMovie`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListHistoryMovie`: %v\n", resp)
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
- **Accept**: text/plain, application/json, text/json

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
    resp, r, err := apiClient.HistoryApi.ListHistorySince(context.Background()).Date(date).EventType(eventType).IncludeMovie(includeMovie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HistoryApi.ListHistorySince``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHistorySince`: []HistoryResource
    fmt.Fprintf(os.Stdout, "Response from `HistoryApi.ListHistorySince`: %v\n", resp)
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
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


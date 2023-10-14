# \MovieAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMovie**](MovieAPI.md#CreateMovie) | **Post** /api/v3/movie | 
[**DeleteMovie**](MovieAPI.md#DeleteMovie) | **Delete** /api/v3/movie/{id} | 
[**GetMovieById**](MovieAPI.md#GetMovieById) | **Get** /api/v3/movie/{id} | 
[**ListMovie**](MovieAPI.md#ListMovie) | **Get** /api/v3/movie | 
[**UpdateMovie**](MovieAPI.md#UpdateMovie) | **Put** /api/v3/movie/{id} | 



## CreateMovie

> MovieResource CreateMovie(ctx).MovieResource(movieResource).Execute()



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
    movieResource := *radarrClient.NewMovieResource() // MovieResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieAPI.CreateMovie(context.Background()).MovieResource(movieResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieAPI.CreateMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMovie`: MovieResource
    fmt.Fprintf(os.Stdout, "Response from `MovieAPI.CreateMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieResource** | [**MovieResource**](MovieResource.md) |  | 

### Return type

[**MovieResource**](MovieResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMovie

> DeleteMovie(ctx, id).DeleteFiles(deleteFiles).AddImportExclusion(addImportExclusion).Execute()



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
    deleteFiles := true // bool |  (optional) (default to false)
    addImportExclusion := true // bool |  (optional) (default to false)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieAPI.DeleteMovie(context.Background(), id).DeleteFiles(deleteFiles).AddImportExclusion(addImportExclusion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieAPI.DeleteMovie``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteFiles** | **bool** |  | [default to false]
 **addImportExclusion** | **bool** |  | [default to false]

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


## GetMovieById

> MovieResource GetMovieById(ctx, id).Execute()



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
    resp, r, err := apiClient.MovieAPI.GetMovieById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieAPI.GetMovieById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMovieById`: MovieResource
    fmt.Fprintf(os.Stdout, "Response from `MovieAPI.GetMovieById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MovieResource**](MovieResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMovie

> []MovieResource ListMovie(ctx).TmdbId(tmdbId).ExcludeLocalCovers(excludeLocalCovers).Execute()



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
    tmdbId := int32(56) // int32 |  (optional)
    excludeLocalCovers := true // bool |  (optional) (default to false)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieAPI.ListMovie(context.Background()).TmdbId(tmdbId).ExcludeLocalCovers(excludeLocalCovers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieAPI.ListMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMovie`: []MovieResource
    fmt.Fprintf(os.Stdout, "Response from `MovieAPI.ListMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmdbId** | **int32** |  | 
 **excludeLocalCovers** | **bool** |  | [default to false]

### Return type

[**[]MovieResource**](MovieResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMovie

> MovieResource UpdateMovie(ctx, id).MoveFiles(moveFiles).MovieResource(movieResource).Execute()



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
    id := "id_example" // string | 
    moveFiles := true // bool |  (optional) (default to false)
    movieResource := *radarrClient.NewMovieResource() // MovieResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieAPI.UpdateMovie(context.Background(), id).MoveFiles(moveFiles).MovieResource(movieResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieAPI.UpdateMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMovie`: MovieResource
    fmt.Fprintf(os.Stdout, "Response from `MovieAPI.UpdateMovie`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveFiles** | **bool** |  | [default to false]
 **movieResource** | [**MovieResource**](MovieResource.md) |  | 

### Return type

[**MovieResource**](MovieResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


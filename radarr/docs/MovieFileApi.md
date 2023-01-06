# \MovieFileApi

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMoviefile**](MovieFileApi.md#DeleteMoviefile) | **Delete** /api/v3/moviefile/{id} | 
[**DeleteMoviefileBulk**](MovieFileApi.md#DeleteMoviefileBulk) | **Delete** /api/v3/moviefile/bulk | 
[**GetMoviefileById**](MovieFileApi.md#GetMoviefileById) | **Get** /api/v3/moviefile/{id} | 
[**ListMoviefile**](MovieFileApi.md#ListMoviefile) | **Get** /api/v3/moviefile | 
[**PutMoviefileEditor**](MovieFileApi.md#PutMoviefileEditor) | **Put** /api/v3/moviefile/editor | 
[**UpdateMoviefile**](MovieFileApi.md#UpdateMoviefile) | **Put** /api/v3/moviefile/{id} | 



## DeleteMoviefile

> DeleteMoviefile(ctx, id).Execute()



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
    resp, r, err := apiClient.MovieFileApi.DeleteMoviefile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieFileApi.DeleteMoviefile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMoviefileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMoviefileBulk

> DeleteMoviefileBulk(ctx).MovieFileListResource(movieFileListResource).Execute()



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
    movieFileListResource := *radarrClient.NewMovieFileListResource() // MovieFileListResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieFileApi.DeleteMoviefileBulk(context.Background()).MovieFileListResource(movieFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieFileApi.DeleteMoviefileBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMoviefileBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieFileListResource** | [**MovieFileListResource**](MovieFileListResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMoviefileById

> MovieFileResource GetMoviefileById(ctx, id).Execute()



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
    resp, r, err := apiClient.MovieFileApi.GetMoviefileById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieFileApi.GetMoviefileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMoviefileById`: MovieFileResource
    fmt.Fprintf(os.Stdout, "Response from `MovieFileApi.GetMoviefileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoviefileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MovieFileResource**](MovieFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMoviefile

> []MovieFileResource ListMoviefile(ctx).MovieId(movieId).MovieFileIds(movieFileIds).Execute()



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
    movieFileIds := []int32{int32(123)} // []int32 |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieFileApi.ListMoviefile(context.Background()).MovieId(movieId).MovieFileIds(movieFileIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieFileApi.ListMoviefile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMoviefile`: []MovieFileResource
    fmt.Fprintf(os.Stdout, "Response from `MovieFileApi.ListMoviefile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMoviefileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int32** |  | 
 **movieFileIds** | **[]int32** |  | 

### Return type

[**[]MovieFileResource**](MovieFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMoviefileEditor

> PutMoviefileEditor(ctx).MovieFileListResource(movieFileListResource).Execute()



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
    movieFileListResource := *radarrClient.NewMovieFileListResource() // MovieFileListResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieFileApi.PutMoviefileEditor(context.Background()).MovieFileListResource(movieFileListResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieFileApi.PutMoviefileEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutMoviefileEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieFileListResource** | [**MovieFileListResource**](MovieFileListResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMoviefile

> MovieFileResource UpdateMoviefile(ctx, id).MovieFileResource(movieFileResource).Execute()



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
    movieFileResource := *radarrClient.NewMovieFileResource() // MovieFileResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieFileApi.UpdateMoviefile(context.Background(), id).MovieFileResource(movieFileResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieFileApi.UpdateMoviefile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMoviefile`: MovieFileResource
    fmt.Fprintf(os.Stdout, "Response from `MovieFileApi.UpdateMoviefile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMoviefileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **movieFileResource** | [**MovieFileResource**](MovieFileResource.md) |  | 

### Return type

[**MovieFileResource**](MovieFileResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \BlocklistApi

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBlocklist**](BlocklistApi.md#DeleteBlocklist) | **Delete** /api/v3/blocklist/{id} | 
[**DeleteBlocklistBulk**](BlocklistApi.md#DeleteBlocklistBulk) | **Delete** /api/v3/blocklist/bulk | 
[**GetBlocklist**](BlocklistApi.md#GetBlocklist) | **Get** /api/v3/blocklist | 
[**ListBlocklistMovie**](BlocklistApi.md#ListBlocklistMovie) | **Get** /api/v3/blocklist/movie | 



## DeleteBlocklist

> DeleteBlocklist(ctx, id).Execute()



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
    resp, r, err := apiClient.BlocklistApi.DeleteBlocklist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocklistApi.DeleteBlocklist``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteBlocklistRequest struct via the builder pattern


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


## DeleteBlocklistBulk

> DeleteBlocklistBulk(ctx).BlocklistBulkResource(blocklistBulkResource).Execute()



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
    blocklistBulkResource := *radarrClient.NewBlocklistBulkResource() // BlocklistBulkResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocklistApi.DeleteBlocklistBulk(context.Background()).BlocklistBulkResource(blocklistBulkResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocklistApi.DeleteBlocklistBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlocklistBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blocklistBulkResource** | [**BlocklistBulkResource**](BlocklistBulkResource.md) |  | 

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


## GetBlocklist

> BlocklistResourcePagingResource GetBlocklist(ctx).Execute()



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

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocklistApi.GetBlocklist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocklistApi.GetBlocklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlocklist`: BlocklistResourcePagingResource
    fmt.Fprintf(os.Stdout, "Response from `BlocklistApi.GetBlocklist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlocklistRequest struct via the builder pattern


### Return type

[**BlocklistResourcePagingResource**](BlocklistResourcePagingResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlocklistMovie

> []BlocklistResource ListBlocklistMovie(ctx).MovieId(movieId).Execute()



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

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocklistApi.ListBlocklistMovie(context.Background()).MovieId(movieId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocklistApi.ListBlocklistMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBlocklistMovie`: []BlocklistResource
    fmt.Fprintf(os.Stdout, "Response from `BlocklistApi.ListBlocklistMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlocklistMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int32** |  | 

### Return type

[**[]BlocklistResource**](BlocklistResource.md)

### Authorization

[X-Api-Key](../README.md#X-Api-Key), [apikey](../README.md#apikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


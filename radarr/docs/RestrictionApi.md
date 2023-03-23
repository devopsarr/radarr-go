# \RestrictionApi

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRestriction**](RestrictionApi.md#CreateRestriction) | **Post** /api/v3/restriction | 
[**DeleteRestriction**](RestrictionApi.md#DeleteRestriction) | **Delete** /api/v3/restriction/{id} | 
[**GetRestrictionById**](RestrictionApi.md#GetRestrictionById) | **Get** /api/v3/restriction/{id} | 
[**ListRestriction**](RestrictionApi.md#ListRestriction) | **Get** /api/v3/restriction | 
[**UpdateRestriction**](RestrictionApi.md#UpdateRestriction) | **Put** /api/v3/restriction/{id} | 



## CreateRestriction

> RestrictionResource CreateRestriction(ctx).RestrictionResource(restrictionResource).Execute()



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
    restrictionResource := *radarrClient.NewRestrictionResource() // RestrictionResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestrictionApi.CreateRestriction(context.Background()).RestrictionResource(restrictionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestrictionApi.CreateRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRestriction`: RestrictionResource
    fmt.Fprintf(os.Stdout, "Response from `RestrictionApi.CreateRestriction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restrictionResource** | [**RestrictionResource**](RestrictionResource.md) |  | 

### Return type

[**RestrictionResource**](RestrictionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRestriction

> DeleteRestriction(ctx, id).Execute()



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
    resp, r, err := apiClient.RestrictionApi.DeleteRestriction(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestrictionApi.DeleteRestriction``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRestrictionRequest struct via the builder pattern


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


## GetRestrictionById

> RestrictionResource GetRestrictionById(ctx, id).Execute()



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
    resp, r, err := apiClient.RestrictionApi.GetRestrictionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestrictionApi.GetRestrictionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestrictionById`: RestrictionResource
    fmt.Fprintf(os.Stdout, "Response from `RestrictionApi.GetRestrictionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestrictionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RestrictionResource**](RestrictionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRestriction

> []RestrictionResource ListRestriction(ctx).Execute()



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
    resp, r, err := apiClient.RestrictionApi.ListRestriction(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestrictionApi.ListRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRestriction`: []RestrictionResource
    fmt.Fprintf(os.Stdout, "Response from `RestrictionApi.ListRestriction`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRestrictionRequest struct via the builder pattern


### Return type

[**[]RestrictionResource**](RestrictionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRestriction

> RestrictionResource UpdateRestriction(ctx, id).RestrictionResource(restrictionResource).Execute()



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
    restrictionResource := *radarrClient.NewRestrictionResource() // RestrictionResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RestrictionApi.UpdateRestriction(context.Background(), id).RestrictionResource(restrictionResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestrictionApi.UpdateRestriction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRestriction`: RestrictionResource
    fmt.Fprintf(os.Stdout, "Response from `RestrictionApi.UpdateRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restrictionResource** | [**RestrictionResource**](RestrictionResource.md) |  | 

### Return type

[**RestrictionResource**](RestrictionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


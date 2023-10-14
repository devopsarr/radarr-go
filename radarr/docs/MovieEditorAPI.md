# \MovieEditorAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMovieEditor**](MovieEditorAPI.md#DeleteMovieEditor) | **Delete** /api/v3/movie/editor | 
[**PutMovieEditor**](MovieEditorAPI.md#PutMovieEditor) | **Put** /api/v3/movie/editor | 



## DeleteMovieEditor

> DeleteMovieEditor(ctx).MovieEditorResource(movieEditorResource).Execute()



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
    movieEditorResource := *radarrClient.NewMovieEditorResource() // MovieEditorResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieEditorAPI.DeleteMovieEditor(context.Background()).MovieEditorResource(movieEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieEditorAPI.DeleteMovieEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMovieEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieEditorResource** | [**MovieEditorResource**](MovieEditorResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMovieEditor

> PutMovieEditor(ctx).MovieEditorResource(movieEditorResource).Execute()



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
    movieEditorResource := *radarrClient.NewMovieEditorResource() // MovieEditorResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MovieEditorAPI.PutMovieEditor(context.Background()).MovieEditorResource(movieEditorResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MovieEditorAPI.PutMovieEditor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutMovieEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieEditorResource** | [**MovieEditorResource**](MovieEditorResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


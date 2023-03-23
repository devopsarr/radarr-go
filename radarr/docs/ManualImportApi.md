# \ManualImportApi

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManualImport**](ManualImportApi.md#CreateManualImport) | **Post** /api/v3/manualimport | 
[**ListManualImport**](ManualImportApi.md#ListManualImport) | **Get** /api/v3/manualimport | 



## CreateManualImport

> CreateManualImport(ctx).ManualImportReprocessResource(manualImportReprocessResource).Execute()



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
    manualImportReprocessResource := []radarrClient.ManualImportReprocessResource{*radarrClient.NewManualImportReprocessResource()} // []ManualImportReprocessResource |  (optional)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualImportApi.CreateManualImport(context.Background()).ManualImportReprocessResource(manualImportReprocessResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualImportApi.CreateManualImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManualImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manualImportReprocessResource** | [**[]ManualImportReprocessResource**](ManualImportReprocessResource.md) |  | 

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


## ListManualImport

> []ManualImportResource ListManualImport(ctx).Folder(folder).DownloadId(downloadId).MovieId(movieId).FilterExistingFiles(filterExistingFiles).Execute()



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
    folder := "folder_example" // string |  (optional)
    downloadId := "downloadId_example" // string |  (optional)
    movieId := int32(56) // int32 |  (optional)
    filterExistingFiles := true // bool |  (optional) (default to true)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualImportApi.ListManualImport(context.Background()).Folder(folder).DownloadId(downloadId).MovieId(movieId).FilterExistingFiles(filterExistingFiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualImportApi.ListManualImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListManualImport`: []ManualImportResource
    fmt.Fprintf(os.Stdout, "Response from `ManualImportApi.ListManualImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListManualImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folder** | **string** |  | 
 **downloadId** | **string** |  | 
 **movieId** | **int32** |  | 
 **filterExistingFiles** | **bool** |  | [default to true]

### Return type

[**[]ManualImportResource**](ManualImportResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


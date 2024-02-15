# \ImportExclusionsAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExclusions**](ImportExclusionsAPI.md#CreateExclusions) | **Post** /api/v3/exclusions | 
[**CreateExclusionsBulk**](ImportExclusionsAPI.md#CreateExclusionsBulk) | **Post** /api/v3/exclusions/bulk | 
[**DeleteExclusions**](ImportExclusionsAPI.md#DeleteExclusions) | **Delete** /api/v3/exclusions/{id} | 
[**GetExclusionsById**](ImportExclusionsAPI.md#GetExclusionsById) | **Get** /api/v3/exclusions/{id} | 
[**ListExclusions**](ImportExclusionsAPI.md#ListExclusions) | **Get** /api/v3/exclusions | 
[**UpdateExclusions**](ImportExclusionsAPI.md#UpdateExclusions) | **Put** /api/v3/exclusions/{id} | 



## CreateExclusions

> ImportExclusionsResource CreateExclusions(ctx).ImportExclusionsResource(importExclusionsResource).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	radarrClient "github.com/devopsarr/radarr-go/radarr"
)

func main() {
	importExclusionsResource := *radarrClient.NewImportExclusionsResource() // ImportExclusionsResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportExclusionsAPI.CreateExclusions(context.Background()).ImportExclusionsResource(importExclusionsResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportExclusionsAPI.CreateExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExclusions`: ImportExclusionsResource
	fmt.Fprintf(os.Stdout, "Response from `ImportExclusionsAPI.CreateExclusions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importExclusionsResource** | [**ImportExclusionsResource**](ImportExclusionsResource.md) |  | 

### Return type

[**ImportExclusionsResource**](ImportExclusionsResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExclusionsBulk

> CreateExclusionsBulk(ctx).ImportExclusionsResource(importExclusionsResource).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	radarrClient "github.com/devopsarr/radarr-go/radarr"
)

func main() {
	importExclusionsResource := []radarrClient.ImportExclusionsResource{*radarrClient.NewImportExclusionsResource()} // []ImportExclusionsResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.ImportExclusionsAPI.CreateExclusionsBulk(context.Background()).ImportExclusionsResource(importExclusionsResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportExclusionsAPI.CreateExclusionsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExclusionsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importExclusionsResource** | [**[]ImportExclusionsResource**](ImportExclusionsResource.md) |  | 

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


## DeleteExclusions

> DeleteExclusions(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	radarrClient "github.com/devopsarr/radarr-go/radarr"
)

func main() {
	id := int32(56) // int32 | 

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.ImportExclusionsAPI.DeleteExclusions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportExclusionsAPI.DeleteExclusions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteExclusionsRequest struct via the builder pattern


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


## GetExclusionsById

> ImportExclusionsResource GetExclusionsById(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	radarrClient "github.com/devopsarr/radarr-go/radarr"
)

func main() {
	id := int32(56) // int32 | 

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportExclusionsAPI.GetExclusionsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportExclusionsAPI.GetExclusionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExclusionsById`: ImportExclusionsResource
	fmt.Fprintf(os.Stdout, "Response from `ImportExclusionsAPI.GetExclusionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExclusionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportExclusionsResource**](ImportExclusionsResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExclusions

> []ImportExclusionsResource ListExclusions(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	radarrClient "github.com/devopsarr/radarr-go/radarr"
)

func main() {

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportExclusionsAPI.ListExclusions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportExclusionsAPI.ListExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExclusions`: []ImportExclusionsResource
	fmt.Fprintf(os.Stdout, "Response from `ImportExclusionsAPI.ListExclusions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListExclusionsRequest struct via the builder pattern


### Return type

[**[]ImportExclusionsResource**](ImportExclusionsResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExclusions

> ImportExclusionsResource UpdateExclusions(ctx, id).ImportExclusionsResource(importExclusionsResource).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	radarrClient "github.com/devopsarr/radarr-go/radarr"
)

func main() {
	id := "id_example" // string | 
	importExclusionsResource := *radarrClient.NewImportExclusionsResource() // ImportExclusionsResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportExclusionsAPI.UpdateExclusions(context.Background(), id).ImportExclusionsResource(importExclusionsResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportExclusionsAPI.UpdateExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExclusions`: ImportExclusionsResource
	fmt.Fprintf(os.Stdout, "Response from `ImportExclusionsAPI.UpdateExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importExclusionsResource** | [**ImportExclusionsResource**](ImportExclusionsResource.md) |  | 

### Return type

[**ImportExclusionsResource**](ImportExclusionsResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ImportListExclusionAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExclusions**](ImportListExclusionAPI.md#CreateExclusions) | **Post** /api/v3/exclusions | 
[**CreateExclusionsBulk**](ImportListExclusionAPI.md#CreateExclusionsBulk) | **Post** /api/v3/exclusions/bulk | 
[**DeleteExclusions**](ImportListExclusionAPI.md#DeleteExclusions) | **Delete** /api/v3/exclusions/{id} | 
[**DeleteExclusionsBulk**](ImportListExclusionAPI.md#DeleteExclusionsBulk) | **Delete** /api/v3/exclusions/bulk | 
[**GetExclusionsById**](ImportListExclusionAPI.md#GetExclusionsById) | **Get** /api/v3/exclusions/{id} | 
[**GetExclusionsPaged**](ImportListExclusionAPI.md#GetExclusionsPaged) | **Get** /api/v3/exclusions/paged | 
[**ListExclusions**](ImportListExclusionAPI.md#ListExclusions) | **Get** /api/v3/exclusions | 
[**UpdateExclusions**](ImportListExclusionAPI.md#UpdateExclusions) | **Put** /api/v3/exclusions/{id} | 



## CreateExclusions

> ImportListExclusionResource CreateExclusions(ctx).ImportListExclusionResource(importListExclusionResource).Execute()



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
	importListExclusionResource := *radarrClient.NewImportListExclusionResource() // ImportListExclusionResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListExclusionAPI.CreateExclusions(context.Background()).ImportListExclusionResource(importListExclusionResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.CreateExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExclusions`: ImportListExclusionResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionAPI.CreateExclusions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListExclusionResource** | [**ImportListExclusionResource**](ImportListExclusionResource.md) |  | 

### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExclusionsBulk

> CreateExclusionsBulk(ctx).ImportListExclusionResource(importListExclusionResource).Execute()



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
	importListExclusionResource := []radarrClient.ImportListExclusionResource{*radarrClient.NewImportListExclusionResource()} // []ImportListExclusionResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.ImportListExclusionAPI.CreateExclusionsBulk(context.Background()).ImportListExclusionResource(importListExclusionResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.CreateExclusionsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExclusionsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListExclusionResource** | [**[]ImportListExclusionResource**](ImportListExclusionResource.md) |  | 

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
	r, err := apiClient.ImportListExclusionAPI.DeleteExclusions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.DeleteExclusions``: %v\n", err)
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


## DeleteExclusionsBulk

> DeleteExclusionsBulk(ctx).ImportListExclusionBulkResource(importListExclusionBulkResource).Execute()



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
	importListExclusionBulkResource := *radarrClient.NewImportListExclusionBulkResource() // ImportListExclusionBulkResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.ImportListExclusionAPI.DeleteExclusionsBulk(context.Background()).ImportListExclusionBulkResource(importListExclusionBulkResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.DeleteExclusionsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExclusionsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importListExclusionBulkResource** | [**ImportListExclusionBulkResource**](ImportListExclusionBulkResource.md) |  | 

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


## GetExclusionsById

> ImportListExclusionResource GetExclusionsById(ctx, id).Execute()



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
	resp, r, err := apiClient.ImportListExclusionAPI.GetExclusionsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.GetExclusionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExclusionsById`: ImportListExclusionResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionAPI.GetExclusionsById`: %v\n", resp)
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

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExclusionsPaged

> ImportListExclusionResourcePagingResource GetExclusionsPaged(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).Execute()



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	sortKey := "sortKey_example" // string |  (optional)
	sortDirection := radarrClient.SortDirection("default") // SortDirection |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListExclusionAPI.GetExclusionsPaged(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.GetExclusionsPaged``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExclusionsPaged`: ImportListExclusionResourcePagingResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionAPI.GetExclusionsPaged`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExclusionsPagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 

### Return type

[**ImportListExclusionResourcePagingResource**](ImportListExclusionResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExclusions

> []ImportListExclusionResource ListExclusions(ctx).Execute()



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
	resp, r, err := apiClient.ImportListExclusionAPI.ListExclusions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.ListExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExclusions`: []ImportListExclusionResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionAPI.ListExclusions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListExclusionsRequest struct via the builder pattern


### Return type

[**[]ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExclusions

> ImportListExclusionResource UpdateExclusions(ctx, id).ImportListExclusionResource(importListExclusionResource).Execute()



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
	importListExclusionResource := *radarrClient.NewImportListExclusionResource() // ImportListExclusionResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportListExclusionAPI.UpdateExclusions(context.Background(), id).ImportListExclusionResource(importListExclusionResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportListExclusionAPI.UpdateExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExclusions`: ImportListExclusionResource
	fmt.Fprintf(os.Stdout, "Response from `ImportListExclusionAPI.UpdateExclusions`: %v\n", resp)
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

 **importListExclusionResource** | [**ImportListExclusionResource**](ImportListExclusionResource.md) |  | 

### Return type

[**ImportListExclusionResource**](ImportListExclusionResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


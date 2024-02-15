# \AlternativeTitleAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlttitleById**](AlternativeTitleAPI.md#GetAlttitleById) | **Get** /api/v3/alttitle/{id} | 
[**ListAlttitle**](AlternativeTitleAPI.md#ListAlttitle) | **Get** /api/v3/alttitle | 



## GetAlttitleById

> AlternativeTitleResource GetAlttitleById(ctx, id).Execute()



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
	resp, r, err := apiClient.AlternativeTitleAPI.GetAlttitleById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeTitleAPI.GetAlttitleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlttitleById`: AlternativeTitleResource
	fmt.Fprintf(os.Stdout, "Response from `AlternativeTitleAPI.GetAlttitleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlttitleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlternativeTitleResource**](AlternativeTitleResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlttitle

> []AlternativeTitleResource ListAlttitle(ctx).MovieId(movieId).MovieMetadataId(movieMetadataId).Execute()



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
	movieId := int32(56) // int32 |  (optional)
	movieMetadataId := int32(56) // int32 |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternativeTitleAPI.ListAlttitle(context.Background()).MovieId(movieId).MovieMetadataId(movieMetadataId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternativeTitleAPI.ListAlttitle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlttitle`: []AlternativeTitleResource
	fmt.Fprintf(os.Stdout, "Response from `AlternativeTitleAPI.ListAlttitle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlttitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int32** |  | 
 **movieMetadataId** | **int32** |  | 

### Return type

[**[]AlternativeTitleResource**](AlternativeTitleResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


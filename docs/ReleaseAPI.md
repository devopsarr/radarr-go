# \ReleaseAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelease**](ReleaseAPI.md#CreateRelease) | **Post** /api/v3/release | 
[**ListRelease**](ReleaseAPI.md#ListRelease) | **Get** /api/v3/release | 



## CreateRelease

> CreateRelease(ctx).ReleaseResource(releaseResource).Execute()



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
	releaseResource := *radarrClient.NewReleaseResource() // ReleaseResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.ReleaseAPI.CreateRelease(context.Background()).ReleaseResource(releaseResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseAPI.CreateRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **releaseResource** | [**ReleaseResource**](ReleaseResource.md) |  | 

### Return type

 (empty response body)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRelease

> []ReleaseResource ListRelease(ctx).MovieId(movieId).Execute()



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

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseAPI.ListRelease(context.Background()).MovieId(movieId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseAPI.ListRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRelease`: []ReleaseResource
	fmt.Fprintf(os.Stdout, "Response from `ReleaseAPI.ListRelease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int32** |  | 

### Return type

[**[]ReleaseResource**](ReleaseResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


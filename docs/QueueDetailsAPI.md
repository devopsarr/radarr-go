# \QueueDetailsAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListQueueDetails**](QueueDetailsAPI.md#ListQueueDetails) | **Get** /api/v3/queue/details | 



## ListQueueDetails

> []QueueResource ListQueueDetails(ctx).MovieId(movieId).IncludeMovie(includeMovie).Execute()



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
	includeMovie := true // bool |  (optional) (default to false)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueueDetailsAPI.ListQueueDetails(context.Background()).MovieId(movieId).IncludeMovie(includeMovie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueueDetailsAPI.ListQueueDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueueDetails`: []QueueResource
	fmt.Fprintf(os.Stdout, "Response from `QueueDetailsAPI.ListQueueDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListQueueDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int32** |  | 
 **includeMovie** | **bool** |  | [default to false]

### Return type

[**[]QueueResource**](QueueResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


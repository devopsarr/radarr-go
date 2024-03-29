# \ExtraFileAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExtraFile**](ExtraFileAPI.md#ListExtraFile) | **Get** /api/v3/extrafile | 



## ListExtraFile

> []ExtraFileResource ListExtraFile(ctx).MovieId(movieId).Execute()



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
	resp, r, err := apiClient.ExtraFileAPI.ListExtraFile(context.Background()).MovieId(movieId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtraFileAPI.ListExtraFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExtraFile`: []ExtraFileResource
	fmt.Fprintf(os.Stdout, "Response from `ExtraFileAPI.ListExtraFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtraFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **int32** |  | 

### Return type

[**[]ExtraFileResource**](ExtraFileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


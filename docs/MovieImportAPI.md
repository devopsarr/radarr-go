# \MovieImportAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMovieImport**](MovieImportAPI.md#CreateMovieImport) | **Post** /api/v3/movie/import | 



## CreateMovieImport

> []MovieResource CreateMovieImport(ctx).MovieResource(movieResource).Execute()



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
	movieResource := []radarrClient.MovieResource{*radarrClient.NewMovieResource()} // []MovieResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MovieImportAPI.CreateMovieImport(context.Background()).MovieResource(movieResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieImportAPI.CreateMovieImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMovieImport`: []MovieResource
	fmt.Fprintf(os.Stdout, "Response from `MovieImportAPI.CreateMovieImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMovieImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieResource** | [**[]MovieResource**](MovieResource.md) |  | 

### Return type

[**[]MovieResource**](MovieResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


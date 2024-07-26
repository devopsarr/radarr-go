# \MissingAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWantedMissing**](MissingAPI.md#GetWantedMissing) | **Get** /api/v3/wanted/missing | 



## GetWantedMissing

> MovieResourcePagingResource GetWantedMissing(ctx).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).Monitored(monitored).Execute()



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
	monitored := true // bool |  (optional) (default to true)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MissingAPI.GetWantedMissing(context.Background()).Page(page).PageSize(pageSize).SortKey(sortKey).SortDirection(sortDirection).Monitored(monitored).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MissingAPI.GetWantedMissing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWantedMissing`: MovieResourcePagingResource
	fmt.Fprintf(os.Stdout, "Response from `MissingAPI.GetWantedMissing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWantedMissingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sortKey** | **string** |  | 
 **sortDirection** | [**SortDirection**](SortDirection.md) |  | 
 **monitored** | **bool** |  | [default to true]

### Return type

[**MovieResourcePagingResource**](MovieResourcePagingResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \CalendarAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCalendar**](CalendarAPI.md#ListCalendar) | **Get** /api/v3/calendar | 



## ListCalendar

> []MovieResource ListCalendar(ctx).Start(start).End(end).Unmonitored(unmonitored).Tags(tags).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	radarrClient "github.com/devopsarr/radarr-go/radarr"
)

func main() {
	start := time.Now() // time.Time |  (optional)
	end := time.Now() // time.Time |  (optional)
	unmonitored := true // bool |  (optional) (default to false)
	tags := "tags_example" // string |  (optional) (default to "")

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.ListCalendar(context.Background()).Start(start).End(end).Unmonitored(unmonitored).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.ListCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCalendar`: []MovieResource
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.ListCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 
 **unmonitored** | **bool** |  | [default to false]
 **tags** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]MovieResource**](MovieResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


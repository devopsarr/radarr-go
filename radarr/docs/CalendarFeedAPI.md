# \CalendarFeedAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedV3CalendarRadarrIcs**](CalendarFeedAPI.md#GetFeedV3CalendarRadarrIcs) | **Get** /feed/v3/calendar/radarr.ics | 



## GetFeedV3CalendarRadarrIcs

> GetFeedV3CalendarRadarrIcs(ctx).PastDays(pastDays).FutureDays(futureDays).TagList(tagList).Unmonitored(unmonitored).Execute()



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
    pastDays := int32(56) // int32 |  (optional) (default to 7)
    futureDays := int32(56) // int32 |  (optional) (default to 28)
    tagList := "tagList_example" // string |  (optional) (default to "")
    unmonitored := true // bool |  (optional) (default to false)

    configuration := radarrClient.NewConfiguration()
    apiClient := radarrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalendarFeedAPI.GetFeedV3CalendarRadarrIcs(context.Background()).PastDays(pastDays).FutureDays(futureDays).TagList(tagList).Unmonitored(unmonitored).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalendarFeedAPI.GetFeedV3CalendarRadarrIcs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedV3CalendarRadarrIcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pastDays** | **int32** |  | [default to 7]
 **futureDays** | **int32** |  | [default to 28]
 **tagList** | **string** |  | [default to &quot;&quot;]
 **unmonitored** | **bool** |  | [default to false]

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


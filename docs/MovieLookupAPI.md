# \MovieLookupAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMovieLookup**](MovieLookupAPI.md#GetMovieLookup) | **Get** /api/v3/movie/lookup | 
[**GetMovieLookupImdb**](MovieLookupAPI.md#GetMovieLookupImdb) | **Get** /api/v3/movie/lookup/imdb | 
[**GetMovieLookupTmdb**](MovieLookupAPI.md#GetMovieLookupTmdb) | **Get** /api/v3/movie/lookup/tmdb | 



## GetMovieLookup

> GetMovieLookup(ctx).Term(term).Execute()



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
	term := "term_example" // string |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.MovieLookupAPI.GetMovieLookup(context.Background()).Term(term).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieLookupAPI.GetMovieLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** |  | 

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


## GetMovieLookupImdb

> GetMovieLookupImdb(ctx).ImdbId(imdbId).Execute()



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
	imdbId := "imdbId_example" // string |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.MovieLookupAPI.GetMovieLookupImdb(context.Background()).ImdbId(imdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieLookupAPI.GetMovieLookupImdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieLookupImdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imdbId** | **string** |  | 

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


## GetMovieLookupTmdb

> GetMovieLookupTmdb(ctx).TmdbId(tmdbId).Execute()



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
	tmdbId := int32(56) // int32 |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.MovieLookupAPI.GetMovieLookupTmdb(context.Background()).TmdbId(tmdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieLookupAPI.GetMovieLookupTmdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieLookupTmdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmdbId** | **int32** |  | 

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


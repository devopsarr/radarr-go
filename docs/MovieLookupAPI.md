# \MovieLookupAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMovieLookup**](MovieLookupAPI.md#ListMovieLookup) | **Get** /api/v3/movie/lookup | 
[**ListMovieLookupImdb**](MovieLookupAPI.md#ListMovieLookupImdb) | **Get** /api/v3/movie/lookup/imdb | 
[**ListMovieLookupTmdb**](MovieLookupAPI.md#ListMovieLookupTmdb) | **Get** /api/v3/movie/lookup/tmdb | 



## ListMovieLookup

> []MovieResource ListMovieLookup(ctx).Term(term).Execute()



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
	resp, r, err := apiClient.MovieLookupAPI.ListMovieLookup(context.Background()).Term(term).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieLookupAPI.ListMovieLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMovieLookup`: []MovieResource
	fmt.Fprintf(os.Stdout, "Response from `MovieLookupAPI.ListMovieLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMovieLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string** |  | 

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


## ListMovieLookupImdb

> []MovieResource ListMovieLookupImdb(ctx).ImdbId(imdbId).Execute()



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
	resp, r, err := apiClient.MovieLookupAPI.ListMovieLookupImdb(context.Background()).ImdbId(imdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieLookupAPI.ListMovieLookupImdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMovieLookupImdb`: []MovieResource
	fmt.Fprintf(os.Stdout, "Response from `MovieLookupAPI.ListMovieLookupImdb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMovieLookupImdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imdbId** | **string** |  | 

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


## ListMovieLookupTmdb

> []MovieResource ListMovieLookupTmdb(ctx).TmdbId(tmdbId).Execute()



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
	resp, r, err := apiClient.MovieLookupAPI.ListMovieLookupTmdb(context.Background()).TmdbId(tmdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieLookupAPI.ListMovieLookupTmdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMovieLookupTmdb`: []MovieResource
	fmt.Fprintf(os.Stdout, "Response from `MovieLookupAPI.ListMovieLookupTmdb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMovieLookupTmdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmdbId** | **int32** |  | 

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


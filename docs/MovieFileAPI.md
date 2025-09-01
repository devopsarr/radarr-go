# \MovieFileAPI

All URIs are relative to *http://localhost:7878*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMovieFile**](MovieFileAPI.md#DeleteMovieFile) | **Delete** /api/v3/moviefile/{id} | 
[**DeleteMovieFileBulk**](MovieFileAPI.md#DeleteMovieFileBulk) | **Delete** /api/v3/moviefile/bulk | 
[**GetMovieFileById**](MovieFileAPI.md#GetMovieFileById) | **Get** /api/v3/moviefile/{id} | 
[**ListMovieFile**](MovieFileAPI.md#ListMovieFile) | **Get** /api/v3/moviefile | 
[**PutMovieFileBulk**](MovieFileAPI.md#PutMovieFileBulk) | **Put** /api/v3/moviefile/bulk | 
[**PutMovieFileEditor**](MovieFileAPI.md#PutMovieFileEditor) | **Put** /api/v3/moviefile/editor | 
[**UpdateMovieFile**](MovieFileAPI.md#UpdateMovieFile) | **Put** /api/v3/moviefile/{id} | 



## DeleteMovieFile

> DeleteMovieFile(ctx, id).Execute()



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
	r, err := apiClient.MovieFileAPI.DeleteMovieFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieFileAPI.DeleteMovieFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMovieFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteMovieFileBulk

> DeleteMovieFileBulk(ctx).MovieFileListResource(movieFileListResource).Execute()



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
	movieFileListResource := *radarrClient.NewMovieFileListResource() // MovieFileListResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.MovieFileAPI.DeleteMovieFileBulk(context.Background()).MovieFileListResource(movieFileListResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieFileAPI.DeleteMovieFileBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMovieFileBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieFileListResource** | [**MovieFileListResource**](MovieFileListResource.md) |  | 

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


## GetMovieFileById

> MovieFileResource GetMovieFileById(ctx, id).Execute()



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
	resp, r, err := apiClient.MovieFileAPI.GetMovieFileById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieFileAPI.GetMovieFileById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovieFileById`: MovieFileResource
	fmt.Fprintf(os.Stdout, "Response from `MovieFileAPI.GetMovieFileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieFileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MovieFileResource**](MovieFileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMovieFile

> []MovieFileResource ListMovieFile(ctx).MovieId(movieId).MovieFileIds(movieFileIds).Execute()



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
	movieId := []int32{int32(123)} // []int32 |  (optional)
	movieFileIds := []int32{int32(123)} // []int32 |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MovieFileAPI.ListMovieFile(context.Background()).MovieId(movieId).MovieFileIds(movieFileIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieFileAPI.ListMovieFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMovieFile`: []MovieFileResource
	fmt.Fprintf(os.Stdout, "Response from `MovieFileAPI.ListMovieFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMovieFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieId** | **[]int32** |  | 
 **movieFileIds** | **[]int32** |  | 

### Return type

[**[]MovieFileResource**](MovieFileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMovieFileBulk

> PutMovieFileBulk(ctx).MovieFileResource(movieFileResource).Execute()



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
	movieFileResource := []radarrClient.MovieFileResource{*radarrClient.NewMovieFileResource()} // []MovieFileResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.MovieFileAPI.PutMovieFileBulk(context.Background()).MovieFileResource(movieFileResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieFileAPI.PutMovieFileBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutMovieFileBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieFileResource** | [**[]MovieFileResource**](MovieFileResource.md) |  | 

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


## PutMovieFileEditor

> PutMovieFileEditor(ctx).MovieFileListResource(movieFileListResource).Execute()



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
	movieFileListResource := *radarrClient.NewMovieFileListResource() // MovieFileListResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	r, err := apiClient.MovieFileAPI.PutMovieFileEditor(context.Background()).MovieFileListResource(movieFileListResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieFileAPI.PutMovieFileEditor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutMovieFileEditorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieFileListResource** | [**MovieFileListResource**](MovieFileListResource.md) |  | 

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


## UpdateMovieFile

> MovieFileResource UpdateMovieFile(ctx, id).MovieFileResource(movieFileResource).Execute()



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
	id := "id_example" // string | 
	movieFileResource := *radarrClient.NewMovieFileResource() // MovieFileResource |  (optional)

	configuration := radarrClient.NewConfiguration()
	apiClient := radarrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MovieFileAPI.UpdateMovieFile(context.Background(), id).MovieFileResource(movieFileResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MovieFileAPI.UpdateMovieFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMovieFile`: MovieFileResource
	fmt.Fprintf(os.Stdout, "Response from `MovieFileAPI.UpdateMovieFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMovieFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **movieFileResource** | [**MovieFileResource**](MovieFileResource.md) |  | 

### Return type

[**MovieFileResource**](MovieFileResource.md)

### Authorization

[apikey](../README.md#apikey), [X-Api-Key](../README.md#X-Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


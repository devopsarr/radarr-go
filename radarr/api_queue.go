/*
Radarr

Radarr API docs

API version: v5.2.6.8376
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// QueueAPIService QueueAPI service
type QueueAPIService service
type ApiDeleteQueueRequest struct {
	ctx context.Context
	ApiService *QueueAPIService
	id int32
	removeFromClient *bool
	blocklist *bool
	skipRedownload *bool
}

func (r ApiDeleteQueueRequest) RemoveFromClient(removeFromClient bool) ApiDeleteQueueRequest {
	r.removeFromClient = &removeFromClient
	return r
}

func (r ApiDeleteQueueRequest) Blocklist(blocklist bool) ApiDeleteQueueRequest {
	r.blocklist = &blocklist
	return r
}

func (r ApiDeleteQueueRequest) SkipRedownload(skipRedownload bool) ApiDeleteQueueRequest {
	r.skipRedownload = &skipRedownload
	return r
}

func (r ApiDeleteQueueRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteQueueExecute(r)
}

/*
DeleteQueue Method for DeleteQueue

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiDeleteQueueRequest
*/
func (a *QueueAPIService) DeleteQueue(ctx context.Context, id int32) ApiDeleteQueueRequest {
	return ApiDeleteQueueRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *QueueAPIService) DeleteQueueExecute(r ApiDeleteQueueRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueueAPIService.DeleteQueue")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/queue/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.removeFromClient != nil {
		localVarQueryParams.Add("removeFromClient", parameterToString(*r.removeFromClient, ""))
	}
	if r.blocklist != nil {
		localVarQueryParams.Add("blocklist", parameterToString(*r.blocklist, ""))
	}
	if r.skipRedownload != nil {
		localVarQueryParams.Add("skipRedownload", parameterToString(*r.skipRedownload, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type ApiDeleteQueueBulkRequest struct {
	ctx context.Context
	ApiService *QueueAPIService
	removeFromClient *bool
	blocklist *bool
	skipRedownload *bool
	queueBulkResource *QueueBulkResource
}

func (r ApiDeleteQueueBulkRequest) RemoveFromClient(removeFromClient bool) ApiDeleteQueueBulkRequest {
	r.removeFromClient = &removeFromClient
	return r
}

func (r ApiDeleteQueueBulkRequest) Blocklist(blocklist bool) ApiDeleteQueueBulkRequest {
	r.blocklist = &blocklist
	return r
}

func (r ApiDeleteQueueBulkRequest) SkipRedownload(skipRedownload bool) ApiDeleteQueueBulkRequest {
	r.skipRedownload = &skipRedownload
	return r
}

func (r ApiDeleteQueueBulkRequest) QueueBulkResource(queueBulkResource QueueBulkResource) ApiDeleteQueueBulkRequest {
	r.queueBulkResource = &queueBulkResource
	return r
}

func (r ApiDeleteQueueBulkRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteQueueBulkExecute(r)
}

/*
DeleteQueueBulk Method for DeleteQueueBulk

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteQueueBulkRequest
*/
func (a *QueueAPIService) DeleteQueueBulk(ctx context.Context) ApiDeleteQueueBulkRequest {
	return ApiDeleteQueueBulkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *QueueAPIService) DeleteQueueBulkExecute(r ApiDeleteQueueBulkRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueueAPIService.DeleteQueueBulk")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/queue/bulk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.removeFromClient != nil {
		localVarQueryParams.Add("removeFromClient", parameterToString(*r.removeFromClient, ""))
	}
	if r.blocklist != nil {
		localVarQueryParams.Add("blocklist", parameterToString(*r.blocklist, ""))
	}
	if r.skipRedownload != nil {
		localVarQueryParams.Add("skipRedownload", parameterToString(*r.skipRedownload, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.queueBulkResource
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type ApiGetQueueRequest struct {
	ctx context.Context
	ApiService *QueueAPIService
	page *int32
	pageSize *int32
	sortKey *string
	sortDirection *SortDirection
	includeUnknownMovieItems *bool
	includeMovie *bool
	movieIds *[]int32
	protocol *DownloadProtocol
	languages *[]int32
	quality *int32
}

func (r ApiGetQueueRequest) Page(page int32) ApiGetQueueRequest {
	r.page = &page
	return r
}

func (r ApiGetQueueRequest) PageSize(pageSize int32) ApiGetQueueRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetQueueRequest) SortKey(sortKey string) ApiGetQueueRequest {
	r.sortKey = &sortKey
	return r
}

func (r ApiGetQueueRequest) SortDirection(sortDirection SortDirection) ApiGetQueueRequest {
	r.sortDirection = &sortDirection
	return r
}

func (r ApiGetQueueRequest) IncludeUnknownMovieItems(includeUnknownMovieItems bool) ApiGetQueueRequest {
	r.includeUnknownMovieItems = &includeUnknownMovieItems
	return r
}

func (r ApiGetQueueRequest) IncludeMovie(includeMovie bool) ApiGetQueueRequest {
	r.includeMovie = &includeMovie
	return r
}

func (r ApiGetQueueRequest) MovieIds(movieIds []int32) ApiGetQueueRequest {
	r.movieIds = &movieIds
	return r
}

func (r ApiGetQueueRequest) Protocol(protocol DownloadProtocol) ApiGetQueueRequest {
	r.protocol = &protocol
	return r
}

func (r ApiGetQueueRequest) Languages(languages []int32) ApiGetQueueRequest {
	r.languages = &languages
	return r
}

func (r ApiGetQueueRequest) Quality(quality int32) ApiGetQueueRequest {
	r.quality = &quality
	return r
}

func (r ApiGetQueueRequest) Execute() (*QueueResourcePagingResource, *http.Response, error) {
	return r.ApiService.GetQueueExecute(r)
}

/*
GetQueue Method for GetQueue

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetQueueRequest
*/
func (a *QueueAPIService) GetQueue(ctx context.Context) ApiGetQueueRequest {
	return ApiGetQueueRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QueueResourcePagingResource
func (a *QueueAPIService) GetQueueExecute(r ApiGetQueueRequest) (*QueueResourcePagingResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueueResourcePagingResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueueAPIService.GetQueue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/queue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	if r.sortKey != nil {
		localVarQueryParams.Add("sortKey", parameterToString(*r.sortKey, ""))
	}
	if r.sortDirection != nil {
		localVarQueryParams.Add("sortDirection", parameterToString(*r.sortDirection, ""))
	}
	if r.includeUnknownMovieItems != nil {
		localVarQueryParams.Add("includeUnknownMovieItems", parameterToString(*r.includeUnknownMovieItems, ""))
	}
	if r.includeMovie != nil {
		localVarQueryParams.Add("includeMovie", parameterToString(*r.includeMovie, ""))
	}
	if r.movieIds != nil {
		t := *r.movieIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("movieIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("movieIds", parameterToString(t, "multi"))
		}
	}
	if r.protocol != nil {
		localVarQueryParams.Add("protocol", parameterToString(*r.protocol, ""))
	}
	if r.languages != nil {
		t := *r.languages
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("languages", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("languages", parameterToString(t, "multi"))
		}
	}
	if r.quality != nil {
		localVarQueryParams.Add("quality", parameterToString(*r.quality, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type ApiGetQueueByIdRequest struct {
	ctx context.Context
	ApiService *QueueAPIService
	id int32
}

func (r ApiGetQueueByIdRequest) Execute() (*QueueResource, *http.Response, error) {
	return r.ApiService.GetQueueByIdExecute(r)
}

/*
GetQueueById Method for GetQueueById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetQueueByIdRequest
*/
func (a *QueueAPIService) GetQueueById(ctx context.Context, id int32) ApiGetQueueByIdRequest {
	return ApiGetQueueByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return QueueResource
func (a *QueueAPIService) GetQueueByIdExecute(r ApiGetQueueByIdRequest) (*QueueResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueueResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueueAPIService.GetQueueById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/queue/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("apikey", key)
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["X-Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

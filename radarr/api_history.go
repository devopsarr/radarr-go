/*
Radarr

Radarr API docs

API version: v5.8.3.8933
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
	"time"
)


// HistoryAPIService HistoryAPI service
type HistoryAPIService service

type ApiCreateHistoryFailedByIdRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	id int32
}

func (r ApiCreateHistoryFailedByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateHistoryFailedByIdExecute(r)
}

/*
CreateHistoryFailedById Method for CreateHistoryFailedById

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCreateHistoryFailedByIdRequest
*/
func (a *HistoryAPIService) CreateHistoryFailedById(ctx context.Context, id int32) ApiCreateHistoryFailedByIdRequest {
	return ApiCreateHistoryFailedByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *HistoryAPIService) CreateHistoryFailedByIdExecute(r ApiCreateHistoryFailedByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.CreateHistoryFailedById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/history/failed/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiGetHistoryRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	page *int32
	pageSize *int32
	sortKey *string
	sortDirection *SortDirection
	includeMovie *bool
	eventType *[]int32
	downloadId *string
	movieIds *[]int32
	languages *[]int32
	quality *[]int32
}

func (r ApiGetHistoryRequest) Page(page int32) ApiGetHistoryRequest {
	r.page = &page
	return r
}

func (r ApiGetHistoryRequest) PageSize(pageSize int32) ApiGetHistoryRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetHistoryRequest) SortKey(sortKey string) ApiGetHistoryRequest {
	r.sortKey = &sortKey
	return r
}

func (r ApiGetHistoryRequest) SortDirection(sortDirection SortDirection) ApiGetHistoryRequest {
	r.sortDirection = &sortDirection
	return r
}

func (r ApiGetHistoryRequest) IncludeMovie(includeMovie bool) ApiGetHistoryRequest {
	r.includeMovie = &includeMovie
	return r
}

func (r ApiGetHistoryRequest) EventType(eventType []int32) ApiGetHistoryRequest {
	r.eventType = &eventType
	return r
}

func (r ApiGetHistoryRequest) DownloadId(downloadId string) ApiGetHistoryRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiGetHistoryRequest) MovieIds(movieIds []int32) ApiGetHistoryRequest {
	r.movieIds = &movieIds
	return r
}

func (r ApiGetHistoryRequest) Languages(languages []int32) ApiGetHistoryRequest {
	r.languages = &languages
	return r
}

func (r ApiGetHistoryRequest) Quality(quality []int32) ApiGetHistoryRequest {
	r.quality = &quality
	return r
}

func (r ApiGetHistoryRequest) Execute() (*HistoryResourcePagingResource, *http.Response, error) {
	return r.ApiService.GetHistoryExecute(r)
}

/*
GetHistory Method for GetHistory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHistoryRequest
*/
func (a *HistoryAPIService) GetHistory(ctx context.Context) ApiGetHistoryRequest {
	return ApiGetHistoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistoryResourcePagingResource
func (a *HistoryAPIService) GetHistoryExecute(r ApiGetHistoryRequest) (*HistoryResourcePagingResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistoryResourcePagingResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.GetHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue int32 = 10
		r.pageSize = &defaultValue
	}
	if r.sortKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortKey", r.sortKey, "")
	}
	if r.sortDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortDirection", r.sortDirection, "")
	}
	if r.includeMovie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeMovie", r.includeMovie, "")
	}
	if r.eventType != nil {
		t := *r.eventType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", t, "multi")
		}
	}
	if r.downloadId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "")
	}
	if r.movieIds != nil {
		t := *r.movieIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "movieIds", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "movieIds", t, "multi")
		}
	}
	if r.languages != nil {
		t := *r.languages
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "languages", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "languages", t, "multi")
		}
	}
	if r.quality != nil {
		t := *r.quality
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "quality", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "quality", t, "multi")
		}
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

type ApiListHistoryMovieRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	movieId *int32
	eventType *MovieHistoryEventType
	includeMovie *bool
}

func (r ApiListHistoryMovieRequest) MovieId(movieId int32) ApiListHistoryMovieRequest {
	r.movieId = &movieId
	return r
}

func (r ApiListHistoryMovieRequest) EventType(eventType MovieHistoryEventType) ApiListHistoryMovieRequest {
	r.eventType = &eventType
	return r
}

func (r ApiListHistoryMovieRequest) IncludeMovie(includeMovie bool) ApiListHistoryMovieRequest {
	r.includeMovie = &includeMovie
	return r
}

func (r ApiListHistoryMovieRequest) Execute() ([]HistoryResource, *http.Response, error) {
	return r.ApiService.ListHistoryMovieExecute(r)
}

/*
ListHistoryMovie Method for ListHistoryMovie

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHistoryMovieRequest
*/
func (a *HistoryAPIService) ListHistoryMovie(ctx context.Context) ApiListHistoryMovieRequest {
	return ApiListHistoryMovieRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []HistoryResource
func (a *HistoryAPIService) ListHistoryMovieExecute(r ApiListHistoryMovieRequest) ([]HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []HistoryResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.ListHistoryMovie")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/history/movie"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.movieId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "movieId", r.movieId, "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "")
	}
	if r.includeMovie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeMovie", r.includeMovie, "")
	} else {
		var defaultValue bool = false
		r.includeMovie = &defaultValue
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

type ApiListHistorySinceRequest struct {
	ctx context.Context
	ApiService *HistoryAPIService
	date *time.Time
	eventType *MovieHistoryEventType
	includeMovie *bool
}

func (r ApiListHistorySinceRequest) Date(date time.Time) ApiListHistorySinceRequest {
	r.date = &date
	return r
}

func (r ApiListHistorySinceRequest) EventType(eventType MovieHistoryEventType) ApiListHistorySinceRequest {
	r.eventType = &eventType
	return r
}

func (r ApiListHistorySinceRequest) IncludeMovie(includeMovie bool) ApiListHistorySinceRequest {
	r.includeMovie = &includeMovie
	return r
}

func (r ApiListHistorySinceRequest) Execute() ([]HistoryResource, *http.Response, error) {
	return r.ApiService.ListHistorySinceExecute(r)
}

/*
ListHistorySince Method for ListHistorySince

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListHistorySinceRequest
*/
func (a *HistoryAPIService) ListHistorySince(ctx context.Context) ApiListHistorySinceRequest {
	return ApiListHistorySinceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []HistoryResource
func (a *HistoryAPIService) ListHistorySinceExecute(r ApiListHistorySinceRequest) ([]HistoryResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []HistoryResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HistoryAPIService.ListHistorySince")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/history/since"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "")
	}
	if r.includeMovie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeMovie", r.includeMovie, "")
	} else {
		var defaultValue bool = false
		r.includeMovie = &defaultValue
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

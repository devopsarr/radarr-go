/*
Radarr

Radarr API docs

API version: v5.17.2.9580
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package radarr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// QueueDetailsAPIService QueueDetailsAPI service
type QueueDetailsAPIService service

type ApiListQueueDetailsRequest struct {
	ctx context.Context
	ApiService *QueueDetailsAPIService
	movieId *int32
	includeMovie *bool
}

func (r ApiListQueueDetailsRequest) MovieId(movieId int32) ApiListQueueDetailsRequest {
	r.movieId = &movieId
	return r
}

func (r ApiListQueueDetailsRequest) IncludeMovie(includeMovie bool) ApiListQueueDetailsRequest {
	r.includeMovie = &includeMovie
	return r
}

func (r ApiListQueueDetailsRequest) Execute() ([]QueueResource, *http.Response, error) {
	return r.ApiService.ListQueueDetailsExecute(r)
}

/*
ListQueueDetails Method for ListQueueDetails

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListQueueDetailsRequest
*/
func (a *QueueDetailsAPIService) ListQueueDetails(ctx context.Context) ApiListQueueDetailsRequest {
	return ApiListQueueDetailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []QueueResource
func (a *QueueDetailsAPIService) ListQueueDetailsExecute(r ApiListQueueDetailsRequest) ([]QueueResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []QueueResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueueDetailsAPIService.ListQueueDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/queue/details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.movieId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "movieId", r.movieId, "form", "")
	}
	if r.includeMovie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeMovie", r.includeMovie, "form", "")
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

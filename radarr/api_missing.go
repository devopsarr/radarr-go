/*
Radarr

Radarr API docs

API version: v5.18.4.9674
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


// MissingAPIService MissingAPI service
type MissingAPIService service

type ApiGetWantedMissingRequest struct {
	ctx context.Context
	ApiService *MissingAPIService
	page *int32
	pageSize *int32
	sortKey *string
	sortDirection *SortDirection
	monitored *bool
}

func (r ApiGetWantedMissingRequest) Page(page int32) ApiGetWantedMissingRequest {
	r.page = &page
	return r
}

func (r ApiGetWantedMissingRequest) PageSize(pageSize int32) ApiGetWantedMissingRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetWantedMissingRequest) SortKey(sortKey string) ApiGetWantedMissingRequest {
	r.sortKey = &sortKey
	return r
}

func (r ApiGetWantedMissingRequest) SortDirection(sortDirection SortDirection) ApiGetWantedMissingRequest {
	r.sortDirection = &sortDirection
	return r
}

func (r ApiGetWantedMissingRequest) Monitored(monitored bool) ApiGetWantedMissingRequest {
	r.monitored = &monitored
	return r
}

func (r ApiGetWantedMissingRequest) Execute() (*MovieResourcePagingResource, *http.Response, error) {
	return r.ApiService.GetWantedMissingExecute(r)
}

/*
GetWantedMissing Method for GetWantedMissing

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetWantedMissingRequest
*/
func (a *MissingAPIService) GetWantedMissing(ctx context.Context) ApiGetWantedMissingRequest {
	return ApiGetWantedMissingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MovieResourcePagingResource
func (a *MissingAPIService) GetWantedMissingExecute(r ApiGetWantedMissingRequest) (*MovieResourcePagingResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MovieResourcePagingResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MissingAPIService.GetWantedMissing")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/wanted/missing"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 10
		r.pageSize = &defaultValue
	}
	if r.sortKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortKey", r.sortKey, "form", "")
	}
	if r.sortDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortDirection", r.sortDirection, "form", "")
	}
	if r.monitored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "monitored", r.monitored, "form", "")
	} else {
		var defaultValue bool = true
		r.monitored = &defaultValue
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

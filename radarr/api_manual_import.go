/*
Radarr

Radarr API docs

API version: v5.19.3.9730
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


// ManualImportAPIService ManualImportAPI service
type ManualImportAPIService service

type ApiCreateManualImportRequest struct {
	ctx context.Context
	ApiService *ManualImportAPIService
	manualImportReprocessResource *[]ManualImportReprocessResource
}

func (r ApiCreateManualImportRequest) ManualImportReprocessResource(manualImportReprocessResource []ManualImportReprocessResource) ApiCreateManualImportRequest {
	r.manualImportReprocessResource = &manualImportReprocessResource
	return r
}

func (r ApiCreateManualImportRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateManualImportExecute(r)
}

/*
CreateManualImport Method for CreateManualImport

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateManualImportRequest
*/
func (a *ManualImportAPIService) CreateManualImport(ctx context.Context) ApiCreateManualImportRequest {
	return ApiCreateManualImportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ManualImportAPIService) CreateManualImportExecute(r ApiCreateManualImportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualImportAPIService.CreateManualImport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/manualimport"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.manualImportReprocessResource
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

type ApiListManualImportRequest struct {
	ctx context.Context
	ApiService *ManualImportAPIService
	folder *string
	downloadId *string
	movieId *int32
	filterExistingFiles *bool
}

func (r ApiListManualImportRequest) Folder(folder string) ApiListManualImportRequest {
	r.folder = &folder
	return r
}

func (r ApiListManualImportRequest) DownloadId(downloadId string) ApiListManualImportRequest {
	r.downloadId = &downloadId
	return r
}

func (r ApiListManualImportRequest) MovieId(movieId int32) ApiListManualImportRequest {
	r.movieId = &movieId
	return r
}

func (r ApiListManualImportRequest) FilterExistingFiles(filterExistingFiles bool) ApiListManualImportRequest {
	r.filterExistingFiles = &filterExistingFiles
	return r
}

func (r ApiListManualImportRequest) Execute() ([]ManualImportResource, *http.Response, error) {
	return r.ApiService.ListManualImportExecute(r)
}

/*
ListManualImport Method for ListManualImport

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListManualImportRequest
*/
func (a *ManualImportAPIService) ListManualImport(ctx context.Context) ApiListManualImportRequest {
	return ApiListManualImportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ManualImportResource
func (a *ManualImportAPIService) ListManualImportExecute(r ApiListManualImportRequest) ([]ManualImportResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ManualImportResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManualImportAPIService.ListManualImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/manualimport"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.folder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "folder", r.folder, "form", "")
	}
	if r.downloadId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "downloadId", r.downloadId, "form", "")
	}
	if r.movieId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "movieId", r.movieId, "form", "")
	}
	if r.filterExistingFiles != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterExistingFiles", r.filterExistingFiles, "form", "")
	} else {
		var defaultValue bool = true
		r.filterExistingFiles = &defaultValue
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

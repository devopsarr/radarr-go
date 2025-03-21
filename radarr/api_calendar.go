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
	"time"
)


// CalendarAPIService CalendarAPI service
type CalendarAPIService service

type ApiListCalendarRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	start *time.Time
	end *time.Time
	unmonitored *bool
	tags *string
}

func (r ApiListCalendarRequest) Start(start time.Time) ApiListCalendarRequest {
	r.start = &start
	return r
}

func (r ApiListCalendarRequest) End(end time.Time) ApiListCalendarRequest {
	r.end = &end
	return r
}

func (r ApiListCalendarRequest) Unmonitored(unmonitored bool) ApiListCalendarRequest {
	r.unmonitored = &unmonitored
	return r
}

func (r ApiListCalendarRequest) Tags(tags string) ApiListCalendarRequest {
	r.tags = &tags
	return r
}

func (r ApiListCalendarRequest) Execute() ([]MovieResource, *http.Response, error) {
	return r.ApiService.ListCalendarExecute(r)
}

/*
ListCalendar Method for ListCalendar

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCalendarRequest
*/
func (a *CalendarAPIService) ListCalendar(ctx context.Context) ApiListCalendarRequest {
	return ApiListCalendarRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []MovieResource
func (a *CalendarAPIService) ListCalendarExecute(r ApiListCalendarRequest) ([]MovieResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MovieResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.ListCalendar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/calendar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "form", "")
	}
	if r.unmonitored != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unmonitored", r.unmonitored, "form", "")
	} else {
		var defaultValue bool = false
		r.unmonitored = &defaultValue
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "form", "")
	} else {
		var defaultValue string = ""
		r.tags = &defaultValue
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

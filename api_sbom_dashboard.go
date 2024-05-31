/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.176.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SbomDashboardAPIService SbomDashboardAPI service
type SbomDashboardAPIService service

type ApiGetSbomsAnalyzedMetricsRequest struct {
	ctx context.Context
	ApiService *SbomDashboardAPIService
}

func (r ApiGetSbomsAnalyzedMetricsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetSbomsAnalyzedMetricsExecute(r)
}

/*
GetSbomsAnalyzedMetrics Gets total of SBOMs analyzed and the threshold in the product license

Queries how many SBOMs have been analyzed and the threshold in the product license

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSbomsAnalyzedMetricsRequest
*/
func (a *SbomDashboardAPIService) GetSbomsAnalyzedMetrics(ctx context.Context) ApiGetSbomsAnalyzedMetricsRequest {
	return ApiGetSbomsAnalyzedMetricsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SbomDashboardAPIService) GetSbomsAnalyzedMetricsExecute(r ApiGetSbomsAnalyzedMetricsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SbomDashboardAPIService.GetSbomsAnalyzedMetrics")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sbom/dashboard/sbomsAnalyzed"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

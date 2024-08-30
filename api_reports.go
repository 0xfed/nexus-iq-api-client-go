/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.181.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ReportsAPIService ReportsAPI service
type ReportsAPIService service

type ApiGetAll1Request struct {
	ctx context.Context
	ApiService *ReportsAPIService
}

func (r ApiGetAll1Request) Execute() ([]ApiApplicationReportDTOV2, *http.Response, error) {
	return r.ApiService.GetAll1Execute(r)
}

/*
GetAll1 Method for GetAll1

Use this method to view all application reports for applications to which  you have access. 

Permissions required: View IQ Elements 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAll1Request
*/
func (a *ReportsAPIService) GetAll1(ctx context.Context) ApiGetAll1Request {
	return ApiGetAll1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiApplicationReportDTOV2
func (a *ReportsAPIService) GetAll1Execute(r ApiGetAll1Request) ([]ApiApplicationReportDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiApplicationReportDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetAll1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/reports/applications"

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

type ApiGetByApplicationIdRequest struct {
	ctx context.Context
	ApiService *ReportsAPIService
	applicationId string
}

func (r ApiGetByApplicationIdRequest) Execute() ([]ApiApplicationReportDTOV2, *http.Response, error) {
	return r.ApiService.GetByApplicationIdExecute(r)
}

/*
GetByApplicationId Method for GetByApplicationId

Use this method to retrieve the application reports for the specified application Id. You can view application reports only for applications to which you have access. 

Permissions required: View IQ Elements 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Enter the internal application Id. You can use the Applications REST API to get the internal application Id. 
 @return ApiGetByApplicationIdRequest
*/
func (a *ReportsAPIService) GetByApplicationId(ctx context.Context, applicationId string) ApiGetByApplicationIdRequest {
	return ApiGetByApplicationIdRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return []ApiApplicationReportDTOV2
func (a *ReportsAPIService) GetByApplicationIdExecute(r ApiGetByApplicationIdRequest) ([]ApiApplicationReportDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiApplicationReportDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetByApplicationId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/reports/applications/{applicationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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

type ApiGetComponentsInQuarantineRequest struct {
	ctx context.Context
	ApiService *ReportsAPIService
}

func (r ApiGetComponentsInQuarantineRequest) Execute() (*ApiComponentsInQuarantineDTO, *http.Response, error) {
	return r.ApiService.GetComponentsInQuarantineExecute(r)
}

/*
GetComponentsInQuarantine Method for GetComponentsInQuarantine

Use this method to retrieve all repository components that are quarantined. The response contains violation details and the quarantine Id of the component. Use the quarantine Id,  to release the component from quarantine, using the Release from Quarantine REST API. 

Permissions required: View IQ Elements and access to the specific repository.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetComponentsInQuarantineRequest
*/
func (a *ReportsAPIService) GetComponentsInQuarantine(ctx context.Context) ApiGetComponentsInQuarantineRequest {
	return ApiGetComponentsInQuarantineRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiComponentsInQuarantineDTO
func (a *ReportsAPIService) GetComponentsInQuarantineExecute(r ApiGetComponentsInQuarantineRequest) (*ApiComponentsInQuarantineDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiComponentsInQuarantineDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetComponentsInQuarantine")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/reports/components/quarantined"

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

type ApiGetComponentsWithWaiversRequest struct {
	ctx context.Context
	ApiService *ReportsAPIService
	format *string
}

// Enter the format/ecosystem of the component
func (r ApiGetComponentsWithWaiversRequest) Format(format string) ApiGetComponentsWithWaiversRequest {
	r.format = &format
	return r
}

func (r ApiGetComponentsWithWaiversRequest) Execute() (*ApiComponentWaiversDTO, *http.Response, error) {
	return r.ApiService.GetComponentsWithWaiversExecute(r)
}

/*
GetComponentsWithWaivers Method for GetComponentsWithWaivers

Use this method to retrieve existing policy waivers by components. For an up-to-date response, ensure that all application and repository reports are current and contain the most recent re-evaluation data.<p>You can specify the format/ecosystem of the component for a filtered result. <p>Permissions required: View IQ Elements and access to the specific applications and repositories 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetComponentsWithWaiversRequest
*/
func (a *ReportsAPIService) GetComponentsWithWaivers(ctx context.Context) ApiGetComponentsWithWaiversRequest {
	return ApiGetComponentsWithWaiversRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiComponentWaiversDTO
func (a *ReportsAPIService) GetComponentsWithWaiversExecute(r ApiGetComponentsWithWaiversRequest) (*ApiComponentWaiversDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiComponentWaiversDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetComponentsWithWaivers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/reports/components/waivers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
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

type ApiGetMetricsRequest struct {
	ctx context.Context
	ApiService *ReportsAPIService
	apiMetricsReportingQueryDTOV2 *ApiMetricsReportingQueryDTOV2
}

func (r ApiGetMetricsRequest) ApiMetricsReportingQueryDTOV2(apiMetricsReportingQueryDTOV2 ApiMetricsReportingQueryDTOV2) ApiGetMetricsRequest {
	r.apiMetricsReportingQueryDTOV2 = &apiMetricsReportingQueryDTOV2
	return r
}

func (r ApiGetMetricsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetMetricsExecute(r)
}

/*
GetMetrics Method for GetMetrics

Use this method to retrieve metrics data such as policy evaluation metrics, violation and remediation metrics aggregated monthly or weekly.

Permissions required: View IQ Elements

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMetricsRequest
*/
func (a *ReportsAPIService) GetMetrics(ctx context.Context) ApiGetMetricsRequest {
	return ApiGetMetricsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ReportsAPIService) GetMetricsExecute(r ApiGetMetricsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetMetrics")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/reports/metrics"

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
	localVarHTTPHeaderAccepts := []string{"application/json", "text/csv"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiMetricsReportingQueryDTOV2
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

type ApiGetReportHistoryForApplicationRequest struct {
	ctx context.Context
	ApiService *ReportsAPIService
	applicationId string
	stage *string
	limit *int32
}

// Enter the specific stage, for which you want retrieve the scan history, e.g. &#39;build&#39; 
func (r ApiGetReportHistoryForApplicationRequest) Stage(stage string) ApiGetReportHistoryForApplicationRequest {
	r.stage = &stage
	return r
}

// Enter the exact no. of most recent reports to retrieve.
func (r ApiGetReportHistoryForApplicationRequest) Limit(limit int32) ApiGetReportHistoryForApplicationRequest {
	r.limit = &limit
	return r
}

func (r ApiGetReportHistoryForApplicationRequest) Execute() (*ApiReportHistoryDTO, *http.Response, error) {
	return r.ApiService.GetReportHistoryForApplicationExecute(r)
}

/*
GetReportHistoryForApplication Method for GetReportHistoryForApplication

Use this method to retrieve previous application scan reports (100 max.) for the specified application. You can view application reports only for applications to which you have access.  

Permissions required: View IQ Elements 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Enter the internal application Id. You can use the Applications REST API to get the internal application Id. 
 @return ApiGetReportHistoryForApplicationRequest
*/
func (a *ReportsAPIService) GetReportHistoryForApplication(ctx context.Context, applicationId string) ApiGetReportHistoryForApplicationRequest {
	return ApiGetReportHistoryForApplicationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApiReportHistoryDTO
func (a *ReportsAPIService) GetReportHistoryForApplicationExecute(r ApiGetReportHistoryForApplicationRequest) (*ApiReportHistoryDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiReportHistoryDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetReportHistoryForApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/reports/applications/{applicationId}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.stage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stage", r.stage, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
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

type ApiGetStaleWaiversRequest struct {
	ctx context.Context
	ApiService *ReportsAPIService
}

func (r ApiGetStaleWaiversRequest) Execute() (*ApiStaleWaiversResponseDTO, *http.Response, error) {
	return r.ApiService.GetStaleWaiversExecute(r)
}

/*
GetStaleWaivers Method for GetStaleWaivers

Stale waivers pose a risk because they could be applied unintentionally. Use this method to retrieve stale waivers to eliminate this risk for future application evaluations. 

Permissions required: View IQ Elements. You can view stale waivers only for applications/repositories  to which you have access. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetStaleWaiversRequest
*/
func (a *ReportsAPIService) GetStaleWaivers(ctx context.Context) ApiGetStaleWaiversRequest {
	return ApiGetStaleWaiversRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiStaleWaiversResponseDTO
func (a *ReportsAPIService) GetStaleWaiversExecute(r ApiGetStaleWaiversRequest) (*ApiStaleWaiversResponseDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStaleWaiversResponseDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetStaleWaivers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/reports/waivers/stale"

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

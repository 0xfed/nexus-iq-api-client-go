/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 165
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


// EvaluationAPIService EvaluationAPI service
type EvaluationAPIService service

type ApiDeprecatedManifestEvaluationRequest struct {
	ctx context.Context
	ApiService *EvaluationAPIService
	applicationId string
	apiSourceControlEvaluationRequestDTO *ApiSourceControlEvaluationRequestDTO
}

func (r ApiDeprecatedManifestEvaluationRequest) ApiSourceControlEvaluationRequestDTO(apiSourceControlEvaluationRequestDTO ApiSourceControlEvaluationRequestDTO) ApiDeprecatedManifestEvaluationRequest {
	r.apiSourceControlEvaluationRequestDTO = &apiSourceControlEvaluationRequestDTO
	return r
}

func (r ApiDeprecatedManifestEvaluationRequest) Execute() (*ApiApplicationEvaluationStatusDTOV2, *http.Response, error) {
	return r.ApiService.DeprecatedManifestEvaluationExecute(r)
}

/*
DeprecatedManifestEvaluation Method for DeprecatedManifestEvaluation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @return ApiDeprecatedManifestEvaluationRequest

Deprecated
*/
func (a *EvaluationAPIService) DeprecatedManifestEvaluation(ctx context.Context, applicationId string) ApiDeprecatedManifestEvaluationRequest {
	return ApiDeprecatedManifestEvaluationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApiApplicationEvaluationStatusDTOV2
// Deprecated
func (a *EvaluationAPIService) DeprecatedManifestEvaluationExecute(r ApiDeprecatedManifestEvaluationRequest) (*ApiApplicationEvaluationStatusDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationEvaluationStatusDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EvaluationAPIService.DeprecatedManifestEvaluation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/evaluation/applications/{applicationId}/manifestEvaluation"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiSourceControlEvaluationRequestDTO
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
			var v ApiApplicationEvaluationStatusDTOV2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiEvaluateComponentsRequest struct {
	ctx context.Context
	ApiService *EvaluationAPIService
	applicationId string
	apiComponentEvaluationRequestDTOV2 *ApiComponentEvaluationRequestDTOV2
}

func (r ApiEvaluateComponentsRequest) ApiComponentEvaluationRequestDTOV2(apiComponentEvaluationRequestDTOV2 ApiComponentEvaluationRequestDTOV2) ApiEvaluateComponentsRequest {
	r.apiComponentEvaluationRequestDTOV2 = &apiComponentEvaluationRequestDTOV2
	return r
}

func (r ApiEvaluateComponentsRequest) Execute() (*ApiComponentEvaluationTicketDTOV2, *http.Response, error) {
	return r.ApiService.EvaluateComponentsExecute(r)
}

/*
EvaluateComponents Method for EvaluateComponents

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @return ApiEvaluateComponentsRequest
*/
func (a *EvaluationAPIService) EvaluateComponents(ctx context.Context, applicationId string) ApiEvaluateComponentsRequest {
	return ApiEvaluateComponentsRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApiComponentEvaluationTicketDTOV2
func (a *EvaluationAPIService) EvaluateComponentsExecute(r ApiEvaluateComponentsRequest) (*ApiComponentEvaluationTicketDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiComponentEvaluationTicketDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EvaluationAPIService.EvaluateComponents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/evaluation/applications/{applicationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiComponentEvaluationRequestDTOV2
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
			var v ApiComponentEvaluationTicketDTOV2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiEvaluateSourceControlRequest struct {
	ctx context.Context
	ApiService *EvaluationAPIService
	applicationId string
	apiSourceControlEvaluationRequestDTO *ApiSourceControlEvaluationRequestDTO
}

func (r ApiEvaluateSourceControlRequest) ApiSourceControlEvaluationRequestDTO(apiSourceControlEvaluationRequestDTO ApiSourceControlEvaluationRequestDTO) ApiEvaluateSourceControlRequest {
	r.apiSourceControlEvaluationRequestDTO = &apiSourceControlEvaluationRequestDTO
	return r
}

func (r ApiEvaluateSourceControlRequest) Execute() (*ApiApplicationEvaluationStatusDTOV2, *http.Response, error) {
	return r.ApiService.EvaluateSourceControlExecute(r)
}

/*
EvaluateSourceControl Method for EvaluateSourceControl

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @return ApiEvaluateSourceControlRequest
*/
func (a *EvaluationAPIService) EvaluateSourceControl(ctx context.Context, applicationId string) ApiEvaluateSourceControlRequest {
	return ApiEvaluateSourceControlRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApiApplicationEvaluationStatusDTOV2
func (a *EvaluationAPIService) EvaluateSourceControlExecute(r ApiEvaluateSourceControlRequest) (*ApiApplicationEvaluationStatusDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationEvaluationStatusDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EvaluationAPIService.EvaluateSourceControl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/evaluation/applications/{applicationId}/sourceControlEvaluation"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiSourceControlEvaluationRequestDTO
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
			var v ApiApplicationEvaluationStatusDTOV2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetApplicationEvaluationStatusRequest struct {
	ctx context.Context
	ApiService *EvaluationAPIService
	applicationId string
	statusId string
}

func (r ApiGetApplicationEvaluationStatusRequest) Execute() (*ApiApplicationEvaluationResultDTOV2, *http.Response, error) {
	return r.ApiService.GetApplicationEvaluationStatusExecute(r)
}

/*
GetApplicationEvaluationStatus Method for GetApplicationEvaluationStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @param statusId
 @return ApiGetApplicationEvaluationStatusRequest
*/
func (a *EvaluationAPIService) GetApplicationEvaluationStatus(ctx context.Context, applicationId string, statusId string) ApiGetApplicationEvaluationStatusRequest {
	return ApiGetApplicationEvaluationStatusRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		statusId: statusId,
	}
}

// Execute executes the request
//  @return ApiApplicationEvaluationResultDTOV2
func (a *EvaluationAPIService) GetApplicationEvaluationStatusExecute(r ApiGetApplicationEvaluationStatusRequest) (*ApiApplicationEvaluationResultDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationEvaluationResultDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EvaluationAPIService.GetApplicationEvaluationStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/evaluation/applications/{applicationId}/status/{statusId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"statusId"+"}", url.PathEscape(parameterValueToString(r.statusId, "statusId")), -1)

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
			var v ApiApplicationEvaluationResultDTOV2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetComponentEvaluationRequest struct {
	ctx context.Context
	ApiService *EvaluationAPIService
	applicationId string
	resultId string
}

func (r ApiGetComponentEvaluationRequest) Execute() (*ApiComponentEvaluationResultDTOV2, *http.Response, error) {
	return r.ApiService.GetComponentEvaluationExecute(r)
}

/*
GetComponentEvaluation Method for GetComponentEvaluation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @param resultId
 @return ApiGetComponentEvaluationRequest
*/
func (a *EvaluationAPIService) GetComponentEvaluation(ctx context.Context, applicationId string, resultId string) ApiGetComponentEvaluationRequest {
	return ApiGetComponentEvaluationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		resultId: resultId,
	}
}

// Execute executes the request
//  @return ApiComponentEvaluationResultDTOV2
func (a *EvaluationAPIService) GetComponentEvaluationExecute(r ApiGetComponentEvaluationRequest) (*ApiComponentEvaluationResultDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiComponentEvaluationResultDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EvaluationAPIService.GetComponentEvaluation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/evaluation/applications/{applicationId}/results/{resultId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resultId"+"}", url.PathEscape(parameterValueToString(r.resultId, "resultId")), -1)

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
			var v ApiComponentEvaluationResultDTOV2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPromoteScanRequest struct {
	ctx context.Context
	ApiService *EvaluationAPIService
	applicationId string
	apiPromoteScanRequestDTOV2 *ApiPromoteScanRequestDTOV2
}

func (r ApiPromoteScanRequest) ApiPromoteScanRequestDTOV2(apiPromoteScanRequestDTOV2 ApiPromoteScanRequestDTOV2) ApiPromoteScanRequest {
	r.apiPromoteScanRequestDTOV2 = &apiPromoteScanRequestDTOV2
	return r
}

func (r ApiPromoteScanRequest) Execute() (*ApiApplicationEvaluationStatusDTOV2, *http.Response, error) {
	return r.ApiService.PromoteScanExecute(r)
}

/*
PromoteScan Method for PromoteScan

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @return ApiPromoteScanRequest
*/
func (a *EvaluationAPIService) PromoteScan(ctx context.Context, applicationId string) ApiPromoteScanRequest {
	return ApiPromoteScanRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApiApplicationEvaluationStatusDTOV2
func (a *EvaluationAPIService) PromoteScanExecute(r ApiPromoteScanRequest) (*ApiApplicationEvaluationStatusDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationEvaluationStatusDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EvaluationAPIService.PromoteScan")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/evaluation/applications/{applicationId}/promoteScan"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiPromoteScanRequestDTOV2
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
			var v ApiApplicationEvaluationStatusDTOV2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

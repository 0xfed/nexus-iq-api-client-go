/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.174.0-01
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


// SourceControlAPIService SourceControlAPI service
type SourceControlAPIService service

type ApiAddSourceControlRequest struct {
	ctx context.Context
	ApiService *SourceControlAPIService
	ownerType string
	internalOwnerId string
	apiSourceControlDTO *ApiSourceControlDTO
}

func (r ApiAddSourceControlRequest) ApiSourceControlDTO(apiSourceControlDTO ApiSourceControlDTO) ApiAddSourceControlRequest {
	r.apiSourceControlDTO = &apiSourceControlDTO
	return r
}

func (r ApiAddSourceControlRequest) Execute() (*ApiSourceControlDTO, *http.Response, error) {
	return r.ApiService.AddSourceControlExecute(r)
}

/*
AddSourceControl Method for AddSourceControl

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiAddSourceControlRequest
*/
func (a *SourceControlAPIService) AddSourceControl(ctx context.Context, ownerType string, internalOwnerId string) ApiAddSourceControlRequest {
	return ApiAddSourceControlRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
//  @return ApiSourceControlDTO
func (a *SourceControlAPIService) AddSourceControlExecute(r ApiAddSourceControlRequest) (*ApiSourceControlDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiSourceControlDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourceControlAPIService.AddSourceControl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sourceControl/{ownerType}/{internalOwnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)

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
	localVarPostBody = r.apiSourceControlDTO
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
			var v ApiSourceControlDTO
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

type ApiDeleteSourceControlRequest struct {
	ctx context.Context
	ApiService *SourceControlAPIService
	ownerType string
	internalOwnerId string
}

func (r ApiDeleteSourceControlRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSourceControlExecute(r)
}

/*
DeleteSourceControl Method for DeleteSourceControl

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiDeleteSourceControlRequest
*/
func (a *SourceControlAPIService) DeleteSourceControl(ctx context.Context, ownerType string, internalOwnerId string) ApiDeleteSourceControlRequest {
	return ApiDeleteSourceControlRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
func (a *SourceControlAPIService) DeleteSourceControlExecute(r ApiDeleteSourceControlRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourceControlAPIService.DeleteSourceControl")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sourceControl/{ownerType}/{internalOwnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"*/*"}

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

type ApiGetSourceControl1Request struct {
	ctx context.Context
	ApiService *SourceControlAPIService
	ownerType string
	internalOwnerId string
}

func (r ApiGetSourceControl1Request) Execute() (*ApiSourceControlDTO, *http.Response, error) {
	return r.ApiService.GetSourceControl1Execute(r)
}

/*
GetSourceControl1 Method for GetSourceControl1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiGetSourceControl1Request
*/
func (a *SourceControlAPIService) GetSourceControl1(ctx context.Context, ownerType string, internalOwnerId string) ApiGetSourceControl1Request {
	return ApiGetSourceControl1Request{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
//  @return ApiSourceControlDTO
func (a *SourceControlAPIService) GetSourceControl1Execute(r ApiGetSourceControl1Request) (*ApiSourceControlDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiSourceControlDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourceControlAPIService.GetSourceControl1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sourceControl/{ownerType}/{internalOwnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)

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
			var v ApiSourceControlDTO
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

type ApiUpdateSourceControlRequest struct {
	ctx context.Context
	ApiService *SourceControlAPIService
	ownerType string
	internalOwnerId string
	apiSourceControlDTO *ApiSourceControlDTO
}

func (r ApiUpdateSourceControlRequest) ApiSourceControlDTO(apiSourceControlDTO ApiSourceControlDTO) ApiUpdateSourceControlRequest {
	r.apiSourceControlDTO = &apiSourceControlDTO
	return r
}

func (r ApiUpdateSourceControlRequest) Execute() (*ApiSourceControlDTO, *http.Response, error) {
	return r.ApiService.UpdateSourceControlExecute(r)
}

/*
UpdateSourceControl Method for UpdateSourceControl

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiUpdateSourceControlRequest
*/
func (a *SourceControlAPIService) UpdateSourceControl(ctx context.Context, ownerType string, internalOwnerId string) ApiUpdateSourceControlRequest {
	return ApiUpdateSourceControlRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
//  @return ApiSourceControlDTO
func (a *SourceControlAPIService) UpdateSourceControlExecute(r ApiUpdateSourceControlRequest) (*ApiSourceControlDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiSourceControlDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SourceControlAPIService.UpdateSourceControl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sourceControl/{ownerType}/{internalOwnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)

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
	localVarPostBody = r.apiSourceControlDTO
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
			var v ApiSourceControlDTO
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

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
	"strings"
)


// ConfigArtifactoryConnectionAPIService ConfigArtifactoryConnectionAPI service
type ConfigArtifactoryConnectionAPIService service

type ApiAddArtifactoryConnectionRequest struct {
	ctx context.Context
	ApiService *ConfigArtifactoryConnectionAPIService
	ownerType string
	internalOwnerId string
	apiArtifactoryConnectionDTO *ApiArtifactoryConnectionDTO
}

func (r ApiAddArtifactoryConnectionRequest) ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO ApiArtifactoryConnectionDTO) ApiAddArtifactoryConnectionRequest {
	r.apiArtifactoryConnectionDTO = &apiArtifactoryConnectionDTO
	return r
}

func (r ApiAddArtifactoryConnectionRequest) Execute() (*ApiArtifactoryConnectionDTO, *http.Response, error) {
	return r.ApiService.AddArtifactoryConnectionExecute(r)
}

/*
AddArtifactoryConnection Method for AddArtifactoryConnection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiAddArtifactoryConnectionRequest
*/
func (a *ConfigArtifactoryConnectionAPIService) AddArtifactoryConnection(ctx context.Context, ownerType string, internalOwnerId string) ApiAddArtifactoryConnectionRequest {
	return ApiAddArtifactoryConnectionRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
//  @return ApiArtifactoryConnectionDTO
func (a *ConfigArtifactoryConnectionAPIService) AddArtifactoryConnectionExecute(r ApiAddArtifactoryConnectionRequest) (*ApiArtifactoryConnectionDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiArtifactoryConnectionDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArtifactoryConnectionAPIService.AddArtifactoryConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}"
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
	localVarPostBody = r.apiArtifactoryConnectionDTO
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
			var v ApiArtifactoryConnectionDTO
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

type ApiDeleteArtifactoryConnectionRequest struct {
	ctx context.Context
	ApiService *ConfigArtifactoryConnectionAPIService
	ownerType string
	internalOwnerId string
	artifactoryConnectionId string
}

func (r ApiDeleteArtifactoryConnectionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteArtifactoryConnectionExecute(r)
}

/*
DeleteArtifactoryConnection Method for DeleteArtifactoryConnection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @param artifactoryConnectionId
 @return ApiDeleteArtifactoryConnectionRequest
*/
func (a *ConfigArtifactoryConnectionAPIService) DeleteArtifactoryConnection(ctx context.Context, ownerType string, internalOwnerId string, artifactoryConnectionId string) ApiDeleteArtifactoryConnectionRequest {
	return ApiDeleteArtifactoryConnectionRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
		artifactoryConnectionId: artifactoryConnectionId,
	}
}

// Execute executes the request
func (a *ConfigArtifactoryConnectionAPIService) DeleteArtifactoryConnectionExecute(r ApiDeleteArtifactoryConnectionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArtifactoryConnectionAPIService.DeleteArtifactoryConnection")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactoryConnectionId"+"}", url.PathEscape(parameterValueToString(r.artifactoryConnectionId, "artifactoryConnectionId")), -1)

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

type ApiGetArtifactoryConnectionRequest struct {
	ctx context.Context
	ApiService *ConfigArtifactoryConnectionAPIService
	ownerType string
	internalOwnerId string
	artifactoryConnectionId string
}

func (r ApiGetArtifactoryConnectionRequest) Execute() (*ApiArtifactoryConnectionDTO, *http.Response, error) {
	return r.ApiService.GetArtifactoryConnectionExecute(r)
}

/*
GetArtifactoryConnection Method for GetArtifactoryConnection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @param artifactoryConnectionId
 @return ApiGetArtifactoryConnectionRequest
*/
func (a *ConfigArtifactoryConnectionAPIService) GetArtifactoryConnection(ctx context.Context, ownerType string, internalOwnerId string, artifactoryConnectionId string) ApiGetArtifactoryConnectionRequest {
	return ApiGetArtifactoryConnectionRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
		artifactoryConnectionId: artifactoryConnectionId,
	}
}

// Execute executes the request
//  @return ApiArtifactoryConnectionDTO
func (a *ConfigArtifactoryConnectionAPIService) GetArtifactoryConnectionExecute(r ApiGetArtifactoryConnectionRequest) (*ApiArtifactoryConnectionDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiArtifactoryConnectionDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArtifactoryConnectionAPIService.GetArtifactoryConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactoryConnectionId"+"}", url.PathEscape(parameterValueToString(r.artifactoryConnectionId, "artifactoryConnectionId")), -1)

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
			var v ApiArtifactoryConnectionDTO
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

type ApiGetOwnerArtifactoryConnectionRequest struct {
	ctx context.Context
	ApiService *ConfigArtifactoryConnectionAPIService
	ownerType string
	internalOwnerId string
	inherit *bool
}

func (r ApiGetOwnerArtifactoryConnectionRequest) Inherit(inherit bool) ApiGetOwnerArtifactoryConnectionRequest {
	r.inherit = &inherit
	return r
}

func (r ApiGetOwnerArtifactoryConnectionRequest) Execute() (*ApiOwnerArtifactoryConnectionDTO, *http.Response, error) {
	return r.ApiService.GetOwnerArtifactoryConnectionExecute(r)
}

/*
GetOwnerArtifactoryConnection Method for GetOwnerArtifactoryConnection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiGetOwnerArtifactoryConnectionRequest
*/
func (a *ConfigArtifactoryConnectionAPIService) GetOwnerArtifactoryConnection(ctx context.Context, ownerType string, internalOwnerId string) ApiGetOwnerArtifactoryConnectionRequest {
	return ApiGetOwnerArtifactoryConnectionRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
//  @return ApiOwnerArtifactoryConnectionDTO
func (a *ConfigArtifactoryConnectionAPIService) GetOwnerArtifactoryConnectionExecute(r ApiGetOwnerArtifactoryConnectionRequest) (*ApiOwnerArtifactoryConnectionDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiOwnerArtifactoryConnectionDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArtifactoryConnectionAPIService.GetOwnerArtifactoryConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.inherit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inherit", r.inherit, "")
	} else {
		var defaultValue bool = false
		r.inherit = &defaultValue
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
			var v ApiOwnerArtifactoryConnectionDTO
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

type ApiTestArtifactoryConnectionRequest struct {
	ctx context.Context
	ApiService *ConfigArtifactoryConnectionAPIService
	ownerType string
	internalOwnerId string
	apiArtifactoryConnectionDTO *ApiArtifactoryConnectionDTO
}

func (r ApiTestArtifactoryConnectionRequest) ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO ApiArtifactoryConnectionDTO) ApiTestArtifactoryConnectionRequest {
	r.apiArtifactoryConnectionDTO = &apiArtifactoryConnectionDTO
	return r
}

func (r ApiTestArtifactoryConnectionRequest) Execute() (*ApiStatusDTO, *http.Response, error) {
	return r.ApiService.TestArtifactoryConnectionExecute(r)
}

/*
TestArtifactoryConnection Method for TestArtifactoryConnection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiTestArtifactoryConnectionRequest
*/
func (a *ConfigArtifactoryConnectionAPIService) TestArtifactoryConnection(ctx context.Context, ownerType string, internalOwnerId string) ApiTestArtifactoryConnectionRequest {
	return ApiTestArtifactoryConnectionRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
//  @return ApiStatusDTO
func (a *ConfigArtifactoryConnectionAPIService) TestArtifactoryConnectionExecute(r ApiTestArtifactoryConnectionRequest) (*ApiStatusDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStatusDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArtifactoryConnectionAPIService.TestArtifactoryConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/test"
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
	// body params
	localVarPostBody = r.apiArtifactoryConnectionDTO
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
			var v ApiStatusDTO
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

type ApiTestArtifactoryConnection1Request struct {
	ctx context.Context
	ApiService *ConfigArtifactoryConnectionAPIService
	ownerType string
	internalOwnerId string
	artifactoryConnectionId string
}

func (r ApiTestArtifactoryConnection1Request) Execute() (*ApiStatusDTO, *http.Response, error) {
	return r.ApiService.TestArtifactoryConnection1Execute(r)
}

/*
TestArtifactoryConnection1 Method for TestArtifactoryConnection1

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @param artifactoryConnectionId
 @return ApiTestArtifactoryConnection1Request
*/
func (a *ConfigArtifactoryConnectionAPIService) TestArtifactoryConnection1(ctx context.Context, ownerType string, internalOwnerId string, artifactoryConnectionId string) ApiTestArtifactoryConnection1Request {
	return ApiTestArtifactoryConnection1Request{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
		artifactoryConnectionId: artifactoryConnectionId,
	}
}

// Execute executes the request
//  @return ApiStatusDTO
func (a *ConfigArtifactoryConnectionAPIService) TestArtifactoryConnection1Execute(r ApiTestArtifactoryConnection1Request) (*ApiStatusDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStatusDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArtifactoryConnectionAPIService.TestArtifactoryConnection1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId}/test"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactoryConnectionId"+"}", url.PathEscape(parameterValueToString(r.artifactoryConnectionId, "artifactoryConnectionId")), -1)

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
			var v ApiStatusDTO
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

type ApiUpdateArtifactoryConnectionRequest struct {
	ctx context.Context
	ApiService *ConfigArtifactoryConnectionAPIService
	ownerType string
	internalOwnerId string
	artifactoryConnectionId string
	apiArtifactoryConnectionDTO *ApiArtifactoryConnectionDTO
}

func (r ApiUpdateArtifactoryConnectionRequest) ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO ApiArtifactoryConnectionDTO) ApiUpdateArtifactoryConnectionRequest {
	r.apiArtifactoryConnectionDTO = &apiArtifactoryConnectionDTO
	return r
}

func (r ApiUpdateArtifactoryConnectionRequest) Execute() (*ApiArtifactoryConnectionDTO, *http.Response, error) {
	return r.ApiService.UpdateArtifactoryConnectionExecute(r)
}

/*
UpdateArtifactoryConnection Method for UpdateArtifactoryConnection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @param artifactoryConnectionId
 @return ApiUpdateArtifactoryConnectionRequest
*/
func (a *ConfigArtifactoryConnectionAPIService) UpdateArtifactoryConnection(ctx context.Context, ownerType string, internalOwnerId string, artifactoryConnectionId string) ApiUpdateArtifactoryConnectionRequest {
	return ApiUpdateArtifactoryConnectionRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
		artifactoryConnectionId: artifactoryConnectionId,
	}
}

// Execute executes the request
//  @return ApiArtifactoryConnectionDTO
func (a *ConfigArtifactoryConnectionAPIService) UpdateArtifactoryConnectionExecute(r ApiUpdateArtifactoryConnectionRequest) (*ApiArtifactoryConnectionDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiArtifactoryConnectionDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArtifactoryConnectionAPIService.UpdateArtifactoryConnection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactoryConnectionId"+"}", url.PathEscape(parameterValueToString(r.artifactoryConnectionId, "artifactoryConnectionId")), -1)

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
	localVarPostBody = r.apiArtifactoryConnectionDTO
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
			var v ApiArtifactoryConnectionDTO
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

type ApiUpdateOwnerArtifactoryConnectionStatusRequest struct {
	ctx context.Context
	ApiService *ConfigArtifactoryConnectionAPIService
	ownerType string
	internalOwnerId string
	apiArtifactoryConnectionStatusRequestDTO *ApiArtifactoryConnectionStatusRequestDTO
}

func (r ApiUpdateOwnerArtifactoryConnectionStatusRequest) ApiArtifactoryConnectionStatusRequestDTO(apiArtifactoryConnectionStatusRequestDTO ApiArtifactoryConnectionStatusRequestDTO) ApiUpdateOwnerArtifactoryConnectionStatusRequest {
	r.apiArtifactoryConnectionStatusRequestDTO = &apiArtifactoryConnectionStatusRequestDTO
	return r
}

func (r ApiUpdateOwnerArtifactoryConnectionStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateOwnerArtifactoryConnectionStatusExecute(r)
}

/*
UpdateOwnerArtifactoryConnectionStatus Method for UpdateOwnerArtifactoryConnectionStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiUpdateOwnerArtifactoryConnectionStatusRequest
*/
func (a *ConfigArtifactoryConnectionAPIService) UpdateOwnerArtifactoryConnectionStatus(ctx context.Context, ownerType string, internalOwnerId string) ApiUpdateOwnerArtifactoryConnectionStatusRequest {
	return ApiUpdateOwnerArtifactoryConnectionStatusRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
func (a *ConfigArtifactoryConnectionAPIService) UpdateOwnerArtifactoryConnectionStatusExecute(r ApiUpdateOwnerArtifactoryConnectionStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArtifactoryConnectionAPIService.UpdateOwnerArtifactoryConnectionStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}"
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
	localVarPostBody = r.apiArtifactoryConnectionStatusRequestDTO
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

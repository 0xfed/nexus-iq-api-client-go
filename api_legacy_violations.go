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


// LegacyViolationsAPIService LegacyViolationsAPI service
type LegacyViolationsAPIService service

type ApiRestLegacyViolationsOrganizationOrganizationIdGetRequest struct {
	ctx context.Context
	ApiService *LegacyViolationsAPIService
	organizationId string
}

func (r ApiRestLegacyViolationsOrganizationOrganizationIdGetRequest) Execute() (*LegacyViolationsResponse, *http.Response, error) {
	return r.ApiService.RestLegacyViolationsOrganizationOrganizationIdGetExecute(r)
}

/*
RestLegacyViolationsOrganizationOrganizationIdGet Get legacy violations for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The ID of the organization
 @return ApiRestLegacyViolationsOrganizationOrganizationIdGetRequest
*/
func (a *LegacyViolationsAPIService) RestLegacyViolationsOrganizationOrganizationIdGet(ctx context.Context, organizationId string) ApiRestLegacyViolationsOrganizationOrganizationIdGetRequest {
	return ApiRestLegacyViolationsOrganizationOrganizationIdGetRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return LegacyViolationsResponse
func (a *LegacyViolationsAPIService) RestLegacyViolationsOrganizationOrganizationIdGetExecute(r ApiRestLegacyViolationsOrganizationOrganizationIdGetRequest) (*LegacyViolationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LegacyViolationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyViolationsAPIService.RestLegacyViolationsOrganizationOrganizationIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/legacyViolations/organization/{organizationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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

type ApiRestLegacyViolationsOrganizationOrganizationIdPutRequest struct {
	ctx context.Context
	ApiService *LegacyViolationsAPIService
	organizationId string
	restLegacyViolationsOrganizationOrganizationIdPutRequest *RestLegacyViolationsOrganizationOrganizationIdPutRequest
}

func (r ApiRestLegacyViolationsOrganizationOrganizationIdPutRequest) RestLegacyViolationsOrganizationOrganizationIdPutRequest(restLegacyViolationsOrganizationOrganizationIdPutRequest RestLegacyViolationsOrganizationOrganizationIdPutRequest) ApiRestLegacyViolationsOrganizationOrganizationIdPutRequest {
	r.restLegacyViolationsOrganizationOrganizationIdPutRequest = &restLegacyViolationsOrganizationOrganizationIdPutRequest
	return r
}

func (r ApiRestLegacyViolationsOrganizationOrganizationIdPutRequest) Execute() (*LegacyViolationsResponse, *http.Response, error) {
	return r.ApiService.RestLegacyViolationsOrganizationOrganizationIdPutExecute(r)
}

/*
RestLegacyViolationsOrganizationOrganizationIdPut Update legacy violations for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The ID of cthe organization
 @return ApiRestLegacyViolationsOrganizationOrganizationIdPutRequest
*/
func (a *LegacyViolationsAPIService) RestLegacyViolationsOrganizationOrganizationIdPut(ctx context.Context, organizationId string) ApiRestLegacyViolationsOrganizationOrganizationIdPutRequest {
	return ApiRestLegacyViolationsOrganizationOrganizationIdPutRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return LegacyViolationsResponse
func (a *LegacyViolationsAPIService) RestLegacyViolationsOrganizationOrganizationIdPutExecute(r ApiRestLegacyViolationsOrganizationOrganizationIdPutRequest) (*LegacyViolationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LegacyViolationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LegacyViolationsAPIService.RestLegacyViolationsOrganizationOrganizationIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/legacyViolations/organization/{organizationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.restLegacyViolationsOrganizationOrganizationIdPutRequest == nil {
		return localVarReturnValue, nil, reportError("restLegacyViolationsOrganizationOrganizationIdPutRequest is required and must be specified")
	}

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
	localVarPostBody = r.restLegacyViolationsOrganizationOrganizationIdPutRequest
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

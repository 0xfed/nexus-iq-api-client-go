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


// ConfigReverseProxyAuthenticationAPIService ConfigReverseProxyAuthenticationAPI service
type ConfigReverseProxyAuthenticationAPIService service

type ApiDeleteConfiguration4Request struct {
	ctx context.Context
	ApiService *ConfigReverseProxyAuthenticationAPIService
}

func (r ApiDeleteConfiguration4Request) Execute() (*http.Response, error) {
	return r.ApiService.DeleteConfiguration4Execute(r)
}

/*
DeleteConfiguration4 Method for DeleteConfiguration4

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteConfiguration4Request
*/
func (a *ConfigReverseProxyAuthenticationAPIService) DeleteConfiguration4(ctx context.Context) ApiDeleteConfiguration4Request {
	return ApiDeleteConfiguration4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigReverseProxyAuthenticationAPIService) DeleteConfiguration4Execute(r ApiDeleteConfiguration4Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigReverseProxyAuthenticationAPIService.DeleteConfiguration4")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/reverseProxyAuthentication"

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

type ApiGetConfiguration4Request struct {
	ctx context.Context
	ApiService *ConfigReverseProxyAuthenticationAPIService
}

func (r ApiGetConfiguration4Request) Execute() (*ApiReverseProxyAuthenticationConfigurationDTO, *http.Response, error) {
	return r.ApiService.GetConfiguration4Execute(r)
}

/*
GetConfiguration4 Method for GetConfiguration4

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetConfiguration4Request
*/
func (a *ConfigReverseProxyAuthenticationAPIService) GetConfiguration4(ctx context.Context) ApiGetConfiguration4Request {
	return ApiGetConfiguration4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiReverseProxyAuthenticationConfigurationDTO
func (a *ConfigReverseProxyAuthenticationAPIService) GetConfiguration4Execute(r ApiGetConfiguration4Request) (*ApiReverseProxyAuthenticationConfigurationDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiReverseProxyAuthenticationConfigurationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigReverseProxyAuthenticationAPIService.GetConfiguration4")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/reverseProxyAuthentication"

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
			var v ApiReverseProxyAuthenticationConfigurationDTO
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

type ApiSetConfiguration4Request struct {
	ctx context.Context
	ApiService *ConfigReverseProxyAuthenticationAPIService
	apiReverseProxyAuthenticationConfigurationDTO *ApiReverseProxyAuthenticationConfigurationDTO
}

func (r ApiSetConfiguration4Request) ApiReverseProxyAuthenticationConfigurationDTO(apiReverseProxyAuthenticationConfigurationDTO ApiReverseProxyAuthenticationConfigurationDTO) ApiSetConfiguration4Request {
	r.apiReverseProxyAuthenticationConfigurationDTO = &apiReverseProxyAuthenticationConfigurationDTO
	return r
}

func (r ApiSetConfiguration4Request) Execute() (*http.Response, error) {
	return r.ApiService.SetConfiguration4Execute(r)
}

/*
SetConfiguration4 Method for SetConfiguration4

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSetConfiguration4Request
*/
func (a *ConfigReverseProxyAuthenticationAPIService) SetConfiguration4(ctx context.Context) ApiSetConfiguration4Request {
	return ApiSetConfiguration4Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigReverseProxyAuthenticationAPIService) SetConfiguration4Execute(r ApiSetConfiguration4Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigReverseProxyAuthenticationAPIService.SetConfiguration4")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/reverseProxyAuthentication"

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
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiReverseProxyAuthenticationConfigurationDTO
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

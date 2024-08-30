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
)


// SecurityOverridesAPIService SecurityOverridesAPI service
type SecurityOverridesAPIService service

type ApiGetSecurityVulnerabilityOverridesRequest struct {
	ctx context.Context
	ApiService *SecurityOverridesAPIService
	refId *string
	componentPurl *string
	ownerId *string
}

func (r ApiGetSecurityVulnerabilityOverridesRequest) RefId(refId string) ApiGetSecurityVulnerabilityOverridesRequest {
	r.refId = &refId
	return r
}

func (r ApiGetSecurityVulnerabilityOverridesRequest) ComponentPurl(componentPurl string) ApiGetSecurityVulnerabilityOverridesRequest {
	r.componentPurl = &componentPurl
	return r
}

func (r ApiGetSecurityVulnerabilityOverridesRequest) OwnerId(ownerId string) ApiGetSecurityVulnerabilityOverridesRequest {
	r.ownerId = &ownerId
	return r
}

func (r ApiGetSecurityVulnerabilityOverridesRequest) Execute() (*ApiSecurityVulnerabilityOverrideResponseDTOV2, *http.Response, error) {
	return r.ApiService.GetSecurityVulnerabilityOverridesExecute(r)
}

/*
GetSecurityVulnerabilityOverrides Method for GetSecurityVulnerabilityOverrides

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSecurityVulnerabilityOverridesRequest
*/
func (a *SecurityOverridesAPIService) GetSecurityVulnerabilityOverrides(ctx context.Context) ApiGetSecurityVulnerabilityOverridesRequest {
	return ApiGetSecurityVulnerabilityOverridesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiSecurityVulnerabilityOverrideResponseDTOV2
func (a *SecurityOverridesAPIService) GetSecurityVulnerabilityOverridesExecute(r ApiGetSecurityVulnerabilityOverridesRequest) (*ApiSecurityVulnerabilityOverrideResponseDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiSecurityVulnerabilityOverrideResponseDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityOverridesAPIService.GetSecurityVulnerabilityOverrides")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/securityOverrides"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.refId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refId", r.refId, "form", "")
	}
	if r.componentPurl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "componentPurl", r.componentPurl, "form", "")
	}
	if r.ownerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ownerId", r.ownerId, "form", "")
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
			var v ApiSecurityVulnerabilityOverrideResponseDTOV2
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

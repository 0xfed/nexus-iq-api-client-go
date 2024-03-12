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


// CompositeSourceControlAPIService CompositeSourceControlAPI service
type CompositeSourceControlAPIService service

type ApiGetCompositeSourceControlByOwnerRequest struct {
	ctx context.Context
	ApiService *CompositeSourceControlAPIService
	ownerType string
	internalOwnerId string
}

func (r ApiGetCompositeSourceControlByOwnerRequest) Execute() (*ApiCompositeSourceControlDTO, *http.Response, error) {
	return r.ApiService.GetCompositeSourceControlByOwnerExecute(r)
}

/*
GetCompositeSourceControlByOwner Method for GetCompositeSourceControlByOwner

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiGetCompositeSourceControlByOwnerRequest
*/
func (a *CompositeSourceControlAPIService) GetCompositeSourceControlByOwner(ctx context.Context, ownerType string, internalOwnerId string) ApiGetCompositeSourceControlByOwnerRequest {
	return ApiGetCompositeSourceControlByOwnerRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
//  @return ApiCompositeSourceControlDTO
func (a *CompositeSourceControlAPIService) GetCompositeSourceControlByOwnerExecute(r ApiGetCompositeSourceControlByOwnerRequest) (*ApiCompositeSourceControlDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiCompositeSourceControlDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CompositeSourceControlAPIService.GetCompositeSourceControlByOwner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/compositeSourceControl/{ownerType}/{internalOwnerId}"
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
			var v ApiCompositeSourceControlDTO
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

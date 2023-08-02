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


// RepositoriesAPIService RepositoriesAPI service
type RepositoriesAPIService service

type ApiGetQuarantinedByPathRequest struct {
	ctx context.Context
	ApiService *RepositoriesAPIService
	repositoryManagerInstanceId string
	repositoryPublicId string
	requestBody *[]string
}

func (r ApiGetQuarantinedByPathRequest) RequestBody(requestBody []string) ApiGetQuarantinedByPathRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiGetQuarantinedByPathRequest) Execute() (*ApiRepositoryPathResponseDTO, *http.Response, error) {
	return r.ApiService.GetQuarantinedByPathExecute(r)
}

/*
GetQuarantinedByPath Method for GetQuarantinedByPath

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param repositoryManagerInstanceId
 @param repositoryPublicId
 @return ApiGetQuarantinedByPathRequest
*/
func (a *RepositoriesAPIService) GetQuarantinedByPath(ctx context.Context, repositoryManagerInstanceId string, repositoryPublicId string) ApiGetQuarantinedByPathRequest {
	return ApiGetQuarantinedByPathRequest{
		ApiService: a,
		ctx: ctx,
		repositoryManagerInstanceId: repositoryManagerInstanceId,
		repositoryPublicId: repositoryPublicId,
	}
}

// Execute executes the request
//  @return ApiRepositoryPathResponseDTO
func (a *RepositoriesAPIService) GetQuarantinedByPathExecute(r ApiGetQuarantinedByPathRequest) (*ApiRepositoryPathResponseDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiRepositoryPathResponseDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAPIService.GetQuarantinedByPath")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/repositories/{repositoryManagerInstanceId}/{repositoryPublicId}/components/quarantined/pathnames"
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryManagerInstanceId"+"}", url.PathEscape(parameterValueToString(r.repositoryManagerInstanceId, "repositoryManagerInstanceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repositoryPublicId"+"}", url.PathEscape(parameterValueToString(r.repositoryPublicId, "repositoryPublicId")), -1)

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
	localVarPostBody = r.requestBody
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
			var v ApiRepositoryPathResponseDTO
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

type ApiReleaseQuarantineWithoutReEvalRequest struct {
	ctx context.Context
	ApiService *RepositoriesAPIService
	quarantineId string
	body *string
}

func (r ApiReleaseQuarantineWithoutReEvalRequest) Body(body string) ApiReleaseQuarantineWithoutReEvalRequest {
	r.body = &body
	return r
}

func (r ApiReleaseQuarantineWithoutReEvalRequest) Execute() (*ApiComponentReleasedFromQuarantineDTO, *http.Response, error) {
	return r.ApiService.ReleaseQuarantineWithoutReEvalExecute(r)
}

/*
ReleaseQuarantineWithoutReEval Method for ReleaseQuarantineWithoutReEval

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param quarantineId
 @return ApiReleaseQuarantineWithoutReEvalRequest
*/
func (a *RepositoriesAPIService) ReleaseQuarantineWithoutReEval(ctx context.Context, quarantineId string) ApiReleaseQuarantineWithoutReEvalRequest {
	return ApiReleaseQuarantineWithoutReEvalRequest{
		ApiService: a,
		ctx: ctx,
		quarantineId: quarantineId,
	}
}

// Execute executes the request
//  @return ApiComponentReleasedFromQuarantineDTO
func (a *RepositoriesAPIService) ReleaseQuarantineWithoutReEvalExecute(r ApiReleaseQuarantineWithoutReEvalRequest) (*ApiComponentReleasedFromQuarantineDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiComponentReleasedFromQuarantineDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAPIService.ReleaseQuarantineWithoutReEval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/repositories/quarantine/{quarantineId}/release"
	localVarPath = strings.Replace(localVarPath, "{"+"quarantineId"+"}", url.PathEscape(parameterValueToString(r.quarantineId, "quarantineId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	localVarPostBody = r.body
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
			var v ApiComponentReleasedFromQuarantineDTO
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

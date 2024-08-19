/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.180.0-01
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


// RoleMembershipsAPIService RoleMembershipsAPI service
type RoleMembershipsAPIService service

type ApiGetRoleMembershipsApplicationOrOrganizationRequest struct {
	ctx context.Context
	ApiService *RoleMembershipsAPIService
	ownerType string
	internalOwnerId string
}

func (r ApiGetRoleMembershipsApplicationOrOrganizationRequest) Execute() (*ApiRoleMemberMappingListDTO, *http.Response, error) {
	return r.ApiService.GetRoleMembershipsApplicationOrOrganizationExecute(r)
}

/*
GetRoleMembershipsApplicationOrOrganization Method for GetRoleMembershipsApplicationOrOrganization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @return ApiGetRoleMembershipsApplicationOrOrganizationRequest
*/
func (a *RoleMembershipsAPIService) GetRoleMembershipsApplicationOrOrganization(ctx context.Context, ownerType string, internalOwnerId string) ApiGetRoleMembershipsApplicationOrOrganizationRequest {
	return ApiGetRoleMembershipsApplicationOrOrganizationRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
	}
}

// Execute executes the request
//  @return ApiRoleMemberMappingListDTO
func (a *RoleMembershipsAPIService) GetRoleMembershipsApplicationOrOrganizationExecute(r ApiGetRoleMembershipsApplicationOrOrganizationRequest) (*ApiRoleMemberMappingListDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiRoleMemberMappingListDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleMembershipsAPIService.GetRoleMembershipsApplicationOrOrganization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/roleMemberships/{ownerType}/{internalOwnerId}"
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
			var v ApiRoleMemberMappingListDTO
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

type ApiGetRoleMembershipsGlobalOrRepositoryContainerRequest struct {
	ctx context.Context
	ApiService *RoleMembershipsAPIService
	ownerType string
}

func (r ApiGetRoleMembershipsGlobalOrRepositoryContainerRequest) Execute() (*ApiRoleMemberMappingListDTO, *http.Response, error) {
	return r.ApiService.GetRoleMembershipsGlobalOrRepositoryContainerExecute(r)
}

/*
GetRoleMembershipsGlobalOrRepositoryContainer Method for GetRoleMembershipsGlobalOrRepositoryContainer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @return ApiGetRoleMembershipsGlobalOrRepositoryContainerRequest
*/
func (a *RoleMembershipsAPIService) GetRoleMembershipsGlobalOrRepositoryContainer(ctx context.Context, ownerType string) ApiGetRoleMembershipsGlobalOrRepositoryContainerRequest {
	return ApiGetRoleMembershipsGlobalOrRepositoryContainerRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
	}
}

// Execute executes the request
//  @return ApiRoleMemberMappingListDTO
func (a *RoleMembershipsAPIService) GetRoleMembershipsGlobalOrRepositoryContainerExecute(r ApiGetRoleMembershipsGlobalOrRepositoryContainerRequest) (*ApiRoleMemberMappingListDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiRoleMemberMappingListDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleMembershipsAPIService.GetRoleMembershipsGlobalOrRepositoryContainer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/roleMemberships/{ownerType}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)

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
			var v ApiRoleMemberMappingListDTO
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

type ApiGrantRoleMembershipApplicationOrOrganizationRequest struct {
	ctx context.Context
	ApiService *RoleMembershipsAPIService
	ownerType string
	internalOwnerId string
	roleId string
	memberType string
	memberName string
}

func (r ApiGrantRoleMembershipApplicationOrOrganizationRequest) Execute() (*http.Response, error) {
	return r.ApiService.GrantRoleMembershipApplicationOrOrganizationExecute(r)
}

/*
GrantRoleMembershipApplicationOrOrganization Method for GrantRoleMembershipApplicationOrOrganization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @param roleId
 @param memberType
 @param memberName
 @return ApiGrantRoleMembershipApplicationOrOrganizationRequest
*/
func (a *RoleMembershipsAPIService) GrantRoleMembershipApplicationOrOrganization(ctx context.Context, ownerType string, internalOwnerId string, roleId string, memberType string, memberName string) ApiGrantRoleMembershipApplicationOrOrganizationRequest {
	return ApiGrantRoleMembershipApplicationOrOrganizationRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
		roleId: roleId,
		memberType: memberType,
		memberName: memberName,
	}
}

// Execute executes the request
func (a *RoleMembershipsAPIService) GrantRoleMembershipApplicationOrOrganizationExecute(r ApiGrantRoleMembershipApplicationOrOrganizationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleMembershipsAPIService.GrantRoleMembershipApplicationOrOrganization")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/roleMemberships/{ownerType}/{internalOwnerId}/role/{roleId}/{memberType}/{memberName}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberType"+"}", url.PathEscape(parameterValueToString(r.memberType, "memberType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberName"+"}", url.PathEscape(parameterValueToString(r.memberName, "memberName")), -1)

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

type ApiGrantRoleMembershipGlobalOrRepositoryContainerRequest struct {
	ctx context.Context
	ApiService *RoleMembershipsAPIService
	ownerType string
	roleId string
	memberType string
	memberName string
}

func (r ApiGrantRoleMembershipGlobalOrRepositoryContainerRequest) Execute() (*http.Response, error) {
	return r.ApiService.GrantRoleMembershipGlobalOrRepositoryContainerExecute(r)
}

/*
GrantRoleMembershipGlobalOrRepositoryContainer Method for GrantRoleMembershipGlobalOrRepositoryContainer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param roleId
 @param memberType
 @param memberName
 @return ApiGrantRoleMembershipGlobalOrRepositoryContainerRequest
*/
func (a *RoleMembershipsAPIService) GrantRoleMembershipGlobalOrRepositoryContainer(ctx context.Context, ownerType string, roleId string, memberType string, memberName string) ApiGrantRoleMembershipGlobalOrRepositoryContainerRequest {
	return ApiGrantRoleMembershipGlobalOrRepositoryContainerRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		roleId: roleId,
		memberType: memberType,
		memberName: memberName,
	}
}

// Execute executes the request
func (a *RoleMembershipsAPIService) GrantRoleMembershipGlobalOrRepositoryContainerExecute(r ApiGrantRoleMembershipGlobalOrRepositoryContainerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleMembershipsAPIService.GrantRoleMembershipGlobalOrRepositoryContainer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/roleMemberships/{ownerType}/role/{roleId}/{memberType}/{memberName}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberType"+"}", url.PathEscape(parameterValueToString(r.memberType, "memberType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberName"+"}", url.PathEscape(parameterValueToString(r.memberName, "memberName")), -1)

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

type ApiRevokeRoleMembershipApplicationOrOrganizationRequest struct {
	ctx context.Context
	ApiService *RoleMembershipsAPIService
	ownerType string
	internalOwnerId string
	roleId string
	memberType string
	memberName string
}

func (r ApiRevokeRoleMembershipApplicationOrOrganizationRequest) Execute() (*http.Response, error) {
	return r.ApiService.RevokeRoleMembershipApplicationOrOrganizationExecute(r)
}

/*
RevokeRoleMembershipApplicationOrOrganization Method for RevokeRoleMembershipApplicationOrOrganization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param internalOwnerId
 @param roleId
 @param memberType
 @param memberName
 @return ApiRevokeRoleMembershipApplicationOrOrganizationRequest
*/
func (a *RoleMembershipsAPIService) RevokeRoleMembershipApplicationOrOrganization(ctx context.Context, ownerType string, internalOwnerId string, roleId string, memberType string, memberName string) ApiRevokeRoleMembershipApplicationOrOrganizationRequest {
	return ApiRevokeRoleMembershipApplicationOrOrganizationRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
		roleId: roleId,
		memberType: memberType,
		memberName: memberName,
	}
}

// Execute executes the request
func (a *RoleMembershipsAPIService) RevokeRoleMembershipApplicationOrOrganizationExecute(r ApiRevokeRoleMembershipApplicationOrOrganizationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleMembershipsAPIService.RevokeRoleMembershipApplicationOrOrganization")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/roleMemberships/{ownerType}/{internalOwnerId}/role/{roleId}/{memberType}/{memberName}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberType"+"}", url.PathEscape(parameterValueToString(r.memberType, "memberType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberName"+"}", url.PathEscape(parameterValueToString(r.memberName, "memberName")), -1)

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

type ApiRevokeRoleMembershipGlobalOrRepositoryContainerRequest struct {
	ctx context.Context
	ApiService *RoleMembershipsAPIService
	ownerType string
	roleId string
	memberType string
	memberName string
}

func (r ApiRevokeRoleMembershipGlobalOrRepositoryContainerRequest) Execute() (*http.Response, error) {
	return r.ApiService.RevokeRoleMembershipGlobalOrRepositoryContainerExecute(r)
}

/*
RevokeRoleMembershipGlobalOrRepositoryContainer Method for RevokeRoleMembershipGlobalOrRepositoryContainer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param roleId
 @param memberType
 @param memberName
 @return ApiRevokeRoleMembershipGlobalOrRepositoryContainerRequest
*/
func (a *RoleMembershipsAPIService) RevokeRoleMembershipGlobalOrRepositoryContainer(ctx context.Context, ownerType string, roleId string, memberType string, memberName string) ApiRevokeRoleMembershipGlobalOrRepositoryContainerRequest {
	return ApiRevokeRoleMembershipGlobalOrRepositoryContainerRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		roleId: roleId,
		memberType: memberType,
		memberName: memberName,
	}
}

// Execute executes the request
func (a *RoleMembershipsAPIService) RevokeRoleMembershipGlobalOrRepositoryContainerExecute(r ApiRevokeRoleMembershipGlobalOrRepositoryContainerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoleMembershipsAPIService.RevokeRoleMembershipGlobalOrRepositoryContainer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/roleMemberships/{ownerType}/role/{roleId}/{memberType}/{memberName}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"roleId"+"}", url.PathEscape(parameterValueToString(r.roleId, "roleId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberType"+"}", url.PathEscape(parameterValueToString(r.memberType, "memberType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"memberName"+"}", url.PathEscape(parameterValueToString(r.memberName, "memberName")), -1)

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

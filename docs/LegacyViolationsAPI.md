# \LegacyViolationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestLegacyViolationsOrganizationOrganizationIdGet**](LegacyViolationsAPI.md#RestLegacyViolationsOrganizationOrganizationIdGet) | **Get** /rest/legacyViolations/organization/{organizationId} | Get legacy violations for an organization
[**RestLegacyViolationsOrganizationOrganizationIdPut**](LegacyViolationsAPI.md#RestLegacyViolationsOrganizationOrganizationIdPut) | **Put** /rest/legacyViolations/organization/{organizationId} | Update legacy violations for an organization



## RestLegacyViolationsOrganizationOrganizationIdGet

> LegacyViolationsResponse RestLegacyViolationsOrganizationOrganizationIdGet(ctx, organizationId).Execute()

Get legacy violations for an organization

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatypeiq "github.com/0xfed/nexus-iq-api-client-go"
)

func main() {
	organizationId := "organizationId_example" // string | The ID of the organization

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyViolationsAPI.RestLegacyViolationsOrganizationOrganizationIdGet(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyViolationsAPI.RestLegacyViolationsOrganizationOrganizationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestLegacyViolationsOrganizationOrganizationIdGet`: LegacyViolationsResponse
	fmt.Fprintf(os.Stdout, "Response from `LegacyViolationsAPI.RestLegacyViolationsOrganizationOrganizationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The ID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestLegacyViolationsOrganizationOrganizationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LegacyViolationsResponse**](LegacyViolationsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestLegacyViolationsOrganizationOrganizationIdPut

> LegacyViolationsResponse RestLegacyViolationsOrganizationOrganizationIdPut(ctx, organizationId).RestLegacyViolationsOrganizationOrganizationIdPutRequest(restLegacyViolationsOrganizationOrganizationIdPutRequest).Execute()

Update legacy violations for an organization

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatypeiq "github.com/0xfed/nexus-iq-api-client-go"
)

func main() {
	organizationId := "organizationId_example" // string | The ID of cthe organization
	restLegacyViolationsOrganizationOrganizationIdPutRequest := *sonatypeiq.NewRestLegacyViolationsOrganizationOrganizationIdPutRequest(false, false) // RestLegacyViolationsOrganizationOrganizationIdPutRequest | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyViolationsAPI.RestLegacyViolationsOrganizationOrganizationIdPut(context.Background(), organizationId).RestLegacyViolationsOrganizationOrganizationIdPutRequest(restLegacyViolationsOrganizationOrganizationIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyViolationsAPI.RestLegacyViolationsOrganizationOrganizationIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestLegacyViolationsOrganizationOrganizationIdPut`: LegacyViolationsResponse
	fmt.Fprintf(os.Stdout, "Response from `LegacyViolationsAPI.RestLegacyViolationsOrganizationOrganizationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The ID of cthe organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestLegacyViolationsOrganizationOrganizationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restLegacyViolationsOrganizationOrganizationIdPutRequest** | [**RestLegacyViolationsOrganizationOrganizationIdPutRequest**](RestLegacyViolationsOrganizationOrganizationIdPutRequest.md) |  | 

### Return type

[**LegacyViolationsResponse**](LegacyViolationsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


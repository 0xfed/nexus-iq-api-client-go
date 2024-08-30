# \CompositeSourceControlConfigValidatorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateSourceControlConfig**](CompositeSourceControlConfigValidatorAPI.md#ValidateSourceControlConfig) | **Get** /api/v2/compositeSourceControlConfigValidator/application/{applicationId} | 



## ValidateSourceControlConfig

> ConfigurationValidationResult ValidateSourceControlConfig(ctx, applicationId).Execute()



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
	applicationId := "applicationId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CompositeSourceControlConfigValidatorAPI.ValidateSourceControlConfig(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompositeSourceControlConfigValidatorAPI.ValidateSourceControlConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSourceControlConfig`: ConfigurationValidationResult
	fmt.Fprintf(os.Stdout, "Response from `CompositeSourceControlConfigValidatorAPI.ValidateSourceControlConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateSourceControlConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationValidationResult**](ConfigurationValidationResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


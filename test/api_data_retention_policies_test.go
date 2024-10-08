/*
Sonatype Lifecycle Public REST API

Testing DataRetentionPoliciesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sonatypeiq

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	sonatypeiq "github.com/0xfed/nexus-iq-api-client-go"
)

func Test_sonatypeiq_DataRetentionPoliciesAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test DataRetentionPoliciesAPIService GetDataRetentionPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DataRetentionPoliciesAPI.GetDataRetentionPolicies(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataRetentionPoliciesAPIService GetParentDataRetentionPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.DataRetentionPoliciesAPI.GetParentDataRetentionPolicies(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DataRetentionPoliciesAPIService SetDataRetentionPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		httpRes, err := apiClient.DataRetentionPoliciesAPI.SetDataRetentionPolicies(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

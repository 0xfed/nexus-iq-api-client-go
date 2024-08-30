/*
Sonatype Lifecycle Public REST API

Testing PolicyViolationDetailsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sonatypeiq

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	sonatypeiq "github.com/4077/nexus-iq-api-client-go"
)

func Test_sonatypeiq_PolicyViolationDetailsAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test PolicyViolationDetailsAPIService GetApplicableWaivers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var violationId string

		resp, httpRes, err := apiClient.PolicyViolationDetailsAPI.GetApplicableWaivers(context.Background(), violationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyViolationDetailsAPIService GetCrossStagePolicyViolationByConstituentId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PolicyViolationDetailsAPI.GetCrossStagePolicyViolationByConstituentId(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyViolationDetailsAPIService GetCrossStagePolicyViolationById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var violationId string

		resp, httpRes, err := apiClient.PolicyViolationDetailsAPI.GetCrossStagePolicyViolationById(context.Background(), violationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyViolationDetailsAPIService GetPolicyViolations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PolicyViolationDetailsAPI.GetPolicyViolations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyViolationDetailsAPIService GetSimilarWaivers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var violationId string

		resp, httpRes, err := apiClient.PolicyViolationDetailsAPI.GetSimilarWaivers(context.Background(), violationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyViolationDetailsAPIService GetTransitivePolicyViolationsByAppScanComponent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var scanId string

		resp, httpRes, err := apiClient.PolicyViolationDetailsAPI.GetTransitivePolicyViolationsByAppScanComponent(context.Background(), ownerType, ownerId, scanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PolicyViolationDetailsAPIService GetTransitivePolicyViolationsByOwnerStageComponent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string
		var stageId string

		resp, httpRes, err := apiClient.PolicyViolationDetailsAPI.GetTransitivePolicyViolationsByOwnerStageComponent(context.Background(), ownerType, ownerId, stageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

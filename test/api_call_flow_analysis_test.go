/*
Sonatype Lifecycle Public REST API

Testing CallFlowAnalysisAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sonatypeiq

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	sonatypeiq "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func Test_sonatypeiq_CallFlowAnalysisAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test CallFlowAnalysisAPIService DeleteCallFlowAnalysisConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string

		httpRes, err := apiClient.CallFlowAnalysisAPI.DeleteCallFlowAnalysisConfig(context.Background(), ownerType, ownerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallFlowAnalysisAPIService GetCallFlowAnalysisConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string

		resp, httpRes, err := apiClient.CallFlowAnalysisAPI.GetCallFlowAnalysisConfig(context.Background(), ownerType, ownerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallFlowAnalysisAPIService UpsertCallFlowAnalysisConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerType string
		var ownerId string

		resp, httpRes, err := apiClient.CallFlowAnalysisAPI.UpsertCallFlowAnalysisConfig(context.Background(), ownerType, ownerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

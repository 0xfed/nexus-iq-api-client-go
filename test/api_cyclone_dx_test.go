/*
Sonatype Lifecycle Public REST API

Testing CycloneDXAPIService

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

func Test_sonatypeiq_CycloneDXAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test CycloneDXAPIService GetByReportId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var reportId string
		var cdxVersion string

		httpRes, err := apiClient.CycloneDXAPI.GetByReportId(context.Background(), applicationId, reportId, cdxVersion).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CycloneDXAPIService GetLatest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var stageId string
		var cdxVersion string

		httpRes, err := apiClient.CycloneDXAPI.GetLatest(context.Background(), applicationId, stageId, cdxVersion).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

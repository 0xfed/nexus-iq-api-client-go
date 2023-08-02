/*
Sonatype IQ Server

Testing CompositeSourceControlConfigValidatorAPIService

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

func Test_sonatypeiq_CompositeSourceControlConfigValidatorAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test CompositeSourceControlConfigValidatorAPIService ValidateSourceControlConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.CompositeSourceControlConfigValidatorAPI.ValidateSourceControlConfig(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

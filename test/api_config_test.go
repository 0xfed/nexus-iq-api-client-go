/*
Sonatype Lifecycle Public REST API

Testing ConfigAPIService

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

func Test_sonatypeiq_ConfigAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test ConfigAPIService DeleteConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConfigAPI.DeleteConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigAPIService DisableFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var feature string

		httpRes, err := apiClient.ConfigAPI.DisableFeature(context.Background(), feature).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigAPIService EnabledFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var feature string

		httpRes, err := apiClient.ConfigAPI.EnabledFeature(context.Background(), feature).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigAPIService GetConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigAPI.GetConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigAPIService SetConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConfigAPI.SetConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

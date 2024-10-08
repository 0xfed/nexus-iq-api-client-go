/*
Sonatype Lifecycle Public REST API

Testing ConfigJIRAAPIService

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

func Test_sonatypeiq_ConfigJIRAAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test ConfigJIRAAPIService DeleteConfiguration1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConfigJIRAAPI.DeleteConfiguration1(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigJIRAAPIService GetConfiguration1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigJIRAAPI.GetConfiguration1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigJIRAAPIService SetConfiguration1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConfigJIRAAPI.SetConfiguration1(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

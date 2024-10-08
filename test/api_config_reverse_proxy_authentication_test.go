/*
Sonatype Lifecycle Public REST API

Testing ConfigReverseProxyAuthenticationAPIService

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

func Test_sonatypeiq_ConfigReverseProxyAuthenticationAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test ConfigReverseProxyAuthenticationAPIService DeleteConfiguration4", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConfigReverseProxyAuthenticationAPI.DeleteConfiguration4(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigReverseProxyAuthenticationAPIService GetConfiguration4", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigReverseProxyAuthenticationAPI.GetConfiguration4(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigReverseProxyAuthenticationAPIService SetConfiguration4", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConfigReverseProxyAuthenticationAPI.SetConfiguration4(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

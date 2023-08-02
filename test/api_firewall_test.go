/*
Sonatype IQ Server

Testing FirewallAPIService

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

func Test_sonatypeiq_FirewallAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test FirewallAPIService GetFirewallAutoUnquarantineConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FirewallAPI.GetFirewallAutoUnquarantineConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService GetFirewallUnquarantineSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FirewallAPI.GetFirewallUnquarantineSummary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService GetQuarantineList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.FirewallAPI.GetQuarantineList(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService GetQuarantineSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FirewallAPI.GetQuarantineSummary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService GetQuarantinedComponentViewAnonymousAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.FirewallAPI.GetQuarantinedComponentViewAnonymousAccess(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService GetUnquarantineList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.FirewallAPI.GetUnquarantineList(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService SetFirewallAutoUnquarantineConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FirewallAPI.SetFirewallAutoUnquarantineConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService SetQuarantinedComponentViewAnonymousAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var enabled bool

		httpRes, err := apiClient.FirewallAPI.SetQuarantinedComponentViewAnonymousAccess(context.Background(), enabled).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

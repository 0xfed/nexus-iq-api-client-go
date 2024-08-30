/*
Sonatype Lifecycle Public REST API

Testing ConfigMailAPIService

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

func Test_sonatypeiq_ConfigMailAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test ConfigMailAPIService DeleteConfiguration2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConfigMailAPI.DeleteConfiguration2(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigMailAPIService GetConfiguration2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigMailAPI.GetConfiguration2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigMailAPIService SetConfiguration2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConfigMailAPI.SetConfiguration2(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigMailAPIService TestConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var recipientEmail string

		httpRes, err := apiClient.ConfigMailAPI.TestConfiguration(context.Background(), recipientEmail).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

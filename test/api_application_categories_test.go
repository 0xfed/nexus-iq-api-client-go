/*
Sonatype Lifecycle Public REST API

Testing ApplicationCategoriesAPIService

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

func Test_sonatypeiq_ApplicationCategoriesAPIService(t *testing.T) {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)

	t.Run("Test ApplicationCategoriesAPIService AddTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ApplicationCategoriesAPI.AddTag(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCategoriesAPIService DeleteTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string
		var tagId string

		httpRes, err := apiClient.ApplicationCategoriesAPI.DeleteTag(context.Background(), organizationId, tagId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCategoriesAPIService GetApplicableTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ApplicationCategoriesAPI.GetApplicableTags(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCategoriesAPIService GetApplicableTagsByApplicationPublicId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationPublicId string

		resp, httpRes, err := apiClient.ApplicationCategoriesAPI.GetApplicableTagsByApplicationPublicId(context.Background(), applicationPublicId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCategoriesAPIService GetApplicationApplicableTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationPublicId string

		resp, httpRes, err := apiClient.ApplicationCategoriesAPI.GetApplicationApplicableTags(context.Background(), applicationPublicId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCategoriesAPIService GetAppliedPolicyTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ApplicationCategoriesAPI.GetAppliedPolicyTags(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCategoriesAPIService GetAppliedTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ApplicationCategoriesAPI.GetAppliedTags(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCategoriesAPIService GetTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ApplicationCategoriesAPI.GetTags(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCategoriesAPIService GetTagsUsedByApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ApplicationCategoriesAPI.GetTagsUsedByApplications(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationCategoriesAPIService UpdateTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationId string

		resp, httpRes, err := apiClient.ApplicationCategoriesAPI.UpdateTag(context.Background(), organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

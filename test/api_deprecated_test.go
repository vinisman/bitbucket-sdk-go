/*
Bitbucket Data Center

Testing DeprecatedAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/vinisman/bitbucket-sdk-go/openapi"
)

func Test_openapi_DeprecatedAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeprecatedAPIService AddBuildStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var commitId string

		httpRes, err := apiClient.DeprecatedAPI.AddBuildStatus(context.Background(), commitId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService AddGroupToUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DeprecatedAPI.AddGroupToUser(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService AddUserToGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DeprecatedAPI.AddUserToGroup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService Approve", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectKey string
		var pullRequestId string
		var repositorySlug string

		resp, httpRes, err := apiClient.DeprecatedAPI.Approve(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService GetBuildStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var commitId string

		resp, httpRes, err := apiClient.DeprecatedAPI.GetBuildStatus(context.Background(), commitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService GetBuildStatusStats", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var commitId string

		resp, httpRes, err := apiClient.DeprecatedAPI.GetBuildStatusStats(context.Background(), commitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService GetDefaultBranch1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectKey string
		var repositorySlug string

		resp, httpRes, err := apiClient.DeprecatedAPI.GetDefaultBranch1(context.Background(), projectKey, repositorySlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService GetMultipleBuildStatusStats", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DeprecatedAPI.GetMultipleBuildStatusStats(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService RemoveUserFromGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DeprecatedAPI.RemoveUserFromGroup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService SetDefaultBranch1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectKey string
		var repositorySlug string

		httpRes, err := apiClient.DeprecatedAPI.SetDefaultBranch1(context.Background(), projectKey, repositorySlug).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService UnassignParticipantRole1", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectKey string
		var pullRequestId string
		var repositorySlug string

		httpRes, err := apiClient.DeprecatedAPI.UnassignParticipantRole1(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeprecatedAPIService WithdrawApproval", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectKey string
		var pullRequestId string
		var repositorySlug string

		resp, httpRes, err := apiClient.DeprecatedAPI.WithdrawApproval(context.Background(), projectKey, pullRequestId, repositorySlug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

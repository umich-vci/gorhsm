/*
RHSM-API

Testing ImagesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gorhsm

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/umich-vci/gorhsm"
)

func Test_gorhsm_ImagesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImagesApiService DownloadImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var checksum string

		httpRes, err := apiClient.ImagesApi.DownloadImage(context.Background(), checksum).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ListImageDownloadsByVersionArch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var version string
		var arch string

		resp, httpRes, err := apiClient.ImagesApi.ListImageDownloadsByVersionArch(context.Background(), version, arch).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesApiService ListImagesByContentSet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var contentSet string

		resp, httpRes, err := apiClient.ImagesApi.ListImagesByContentSet(context.Background(), contentSet).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

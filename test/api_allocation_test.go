/*
RHSM-API

Testing AllocationAPIService

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

func Test_gorhsm_AllocationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AllocationAPIService AttachEntitlementAllocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.AllocationAPI.AttachEntitlementAllocation(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService CreateSatellite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AllocationAPI.CreateSatellite(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService ExportAllocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.AllocationAPI.ExportAllocation(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService ExportJobAllocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string
		var exportJobID string

		resp, httpRes, err := apiClient.AllocationAPI.ExportJobAllocation(context.Background(), uuid, exportJobID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService GetExportAllocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string
		var exportID string

		resp, httpRes, err := apiClient.AllocationAPI.GetExportAllocation(context.Background(), uuid, exportID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService ListAllocationPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.AllocationAPI.ListAllocationPools(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService ListAllocations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AllocationAPI.ListAllocations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService ListVersionsAllocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AllocationAPI.ListVersionsAllocation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService RemoveAllocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		httpRes, err := apiClient.AllocationAPI.RemoveAllocation(context.Background(), uuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService RemoveAllocationEntitlement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string
		var entitlementID string

		httpRes, err := apiClient.AllocationAPI.RemoveAllocationEntitlement(context.Background(), uuid, entitlementID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService RemoveAllocationEntitlementDeprecated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string
		var entitlementID string

		httpRes, err := apiClient.AllocationAPI.RemoveAllocationEntitlementDeprecated(context.Background(), uuid, entitlementID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService ShowAllocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.AllocationAPI.ShowAllocation(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService UpdateAllocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		httpRes, err := apiClient.AllocationAPI.UpdateAllocation(context.Background(), uuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AllocationAPIService UpdateEntitlementAllocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string
		var entitlementID string

		resp, httpRes, err := apiClient.AllocationAPI.UpdateEntitlementAllocation(context.Background(), uuid, entitlementID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

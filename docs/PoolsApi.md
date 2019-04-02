# \PoolsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePool**](PoolsApi.md#DeletePool) | **Delete** /pools/{pool_id} | deletePool
[**GetPool**](PoolsApi.md#GetPool) | **Get** /pools/{pool_id} | getPool
[**GetPoolCdn**](PoolsApi.md#GetPoolCdn) | **Get** /pools/{pool_id}/cdn | getPoolCdn
[**GetPoolEntitlements**](PoolsApi.md#GetPoolEntitlements) | **Get** /pools/{pool_id}/entitlements | getPoolEntitlements
[**GetSubCertAsPem**](PoolsApi.md#GetSubCertAsPem) | **Get** /pools/{pool_id}/cert | getSubCertAsPem
[**List**](PoolsApi.md#List) | **Get** /pools | 


# **DeletePool**
> DeletePool(ctx, poolId)
deletePool

Remove a Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPool**
> Pool GetPool(ctx, poolId, optional)
getPool

Retrieves a single Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 
 **optional** | ***GetPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumer** | **optional.String**|  | 
 **activeon** | **optional.String**| Uses ISO 8601 format | 

### Return type

[**Pool**](pool.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoolCdn**
> CdnDto GetPoolCdn(ctx, poolId)
getPoolCdn

Retrieve a CDN for a Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

[**CdnDto**](CdnDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoolEntitlements**
> []EntitlementDto GetPoolEntitlements(ctx, poolId)
getPoolEntitlements

Retrieve a list of Entitlements for a Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

[**[]EntitlementDto**](EntitlementDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubCertAsPem**
> string GetSubCertAsPem(ctx, poolId)
getSubCertAsPem

Retrieves a Subscription Certificate As a PEM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []Pool List(ctx, optional)


Retrieves a list of Pools @deprecated Use the method on /owners

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **optional.String**|  | 
 **consumer** | **optional.String**|  | 
 **product** | **optional.String**|  | 
 **listall** | **optional.Bool**| Use with consumerUuid to list all pools available to the consumer. This will include pools which would otherwise be omitted due to a rules warning. (i.e. not recommended) Pools that trigger an error however will still be omitted. (no entitlements available, consumer type mismatch, etc) | [default to false]
 **activeon** | **optional.String**| Uses ISO 8601 format | 

### Return type

[**[]Pool**](pool.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


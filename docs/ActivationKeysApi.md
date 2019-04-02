# \ActivationKeysApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContentOverrides**](ActivationKeysApi.md#AddContentOverrides) | **Put** /activation_keys/{activation_key_id}/content_overrides | 
[**AddPoolToKey**](ActivationKeysApi.md#AddPoolToKey) | **Post** /activation_keys/{activation_key_id}/pools/{pool_id} | Add Pool to Key
[**AddProductIdToKey**](ActivationKeysApi.md#AddProductIdToKey) | **Post** /activation_keys/{activation_key_id}/product/{product_id} | Add Product ID to key
[**DeleteActivationKey**](ActivationKeysApi.md#DeleteActivationKey) | **Delete** /activation_keys/{activation_key_id} | deleteActivationKey
[**DeleteContentOverrides**](ActivationKeysApi.md#DeleteContentOverrides) | **Delete** /activation_keys/{activation_key_id}/content_overrides | 
[**FindActivationKey**](ActivationKeysApi.md#FindActivationKey) | **Get** /activation_keys | findActivationKey
[**GetActivationKey**](ActivationKeysApi.md#GetActivationKey) | **Get** /activation_keys/{activation_key_id} | Get Activation Key
[**GetActivationKeyPools**](ActivationKeysApi.md#GetActivationKeyPools) | **Get** /activation_keys/{activation_key_id}/pools | Get Activation Key Pools
[**GetContentOverrideList**](ActivationKeysApi.md#GetContentOverrideList) | **Get** /activation_keys/{activation_key_id}/content_overrides | 
[**RemovePoolFromKey**](ActivationKeysApi.md#RemovePoolFromKey) | **Delete** /activation_keys/{activation_key_id}/pools/{pool_id} | Remove Pool From Key
[**RemoveProductIdFromKey**](ActivationKeysApi.md#RemoveProductIdFromKey) | **Delete** /activation_keys/{activation_key_id}/product/{product_id} | Remove Product Id from Key
[**UpdateActivationKey**](ActivationKeysApi.md#UpdateActivationKey) | **Put** /activation_keys/{activation_key_id} | Update Activation Key


# **AddContentOverrides**
> []ContentOverride AddContentOverrides(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddContentOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddContentOverridesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentOverride** | [**optional.Interface of []ContentOverride**](array.md)|  | 

### Return type

[**[]ContentOverride**](ContentOverride.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPoolToKey**
> ActivationKeyDto AddPoolToKey(ctx, activationKeyId, poolId, optional)
Add Pool to Key

Adds a Pool to an Activation Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activationKeyId** | **string**|  | 
  **poolId** | **string**|  | 
 **optional** | ***AddPoolToKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddPoolToKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **quantity** | **optional.Int64**|  | 

### Return type

[**ActivationKeyDto**](ActivationKeyDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddProductIdToKey**
> ActivationKeyDto AddProductIdToKey(ctx, activationKeyId, productId)
Add Product ID to key

Adds an Product ID to an Activation Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activationKeyId** | **string**|  | 
  **productId** | **string**|  | 

### Return type

[**ActivationKeyDto**](ActivationKeyDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteActivationKey**
> DeleteActivationKey(ctx, activationKeyId)
deleteActivationKey

Removes an Activation Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activationKeyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContentOverrides**
> []ContentOverride DeleteContentOverrides(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteContentOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteContentOverridesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentOverride** | [**optional.Interface of []ContentOverride**](array.md)|  | 

### Return type

[**[]ContentOverride**](ContentOverride.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindActivationKey**
> []ActivationKey FindActivationKey(ctx, )
findActivationKey

Retrieves a list of Activation Keys

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ActivationKey**](ActivationKey.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivationKey**
> ActivationKeyDto GetActivationKey(ctx, activationKeyId)
Get Activation Key

Retrieves a single Activation Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activationKeyId** | **string**|  | 

### Return type

[**ActivationKeyDto**](ActivationKeyDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetActivationKeyPools**
> []Pool GetActivationKeyPools(ctx, activationKeyId)
Get Activation Key Pools

Retrieves a list of Pools based on the Activation Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activationKeyId** | **string**|  | 

### Return type

[**[]Pool**](pool.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentOverrideList**
> []ContentOverride GetContentOverrideList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ContentOverride**](ContentOverride.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePoolFromKey**
> ActivationKeyDto RemovePoolFromKey(ctx, activationKeyId, poolId)
Remove Pool From Key

Removes a Pool from an Activation Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activationKeyId** | **string**|  | 
  **poolId** | **string**|  | 

### Return type

[**ActivationKeyDto**](ActivationKeyDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveProductIdFromKey**
> ActivationKeyDto RemoveProductIdFromKey(ctx, activationKeyId, productId)
Remove Product Id from Key

Removes a Product ID from an Activation Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activationKeyId** | **string**|  | 
  **productId** | **string**|  | 

### Return type

[**ActivationKeyDto**](ActivationKeyDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateActivationKey**
> ActivationKeyDto UpdateActivationKey(ctx, activationKeyId, activationKeyDto)
Update Activation Key

Updates an Activation Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activationKeyId** | **string**|  | 
  **activationKeyDto** | [**ActivationKeyDto**](ActivationKeyDto.md)|  | 

### Return type

[**ActivationKeyDto**](ActivationKeyDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


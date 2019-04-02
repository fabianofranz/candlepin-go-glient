# \EntitlementsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntitlement**](EntitlementsApi.md#GetEntitlement) | **Get** /entitlements/{entitlement_id} | getEntitlement
[**GetUpstreamCert**](EntitlementsApi.md#GetUpstreamCert) | **Get** /entitlements/{dbid}/upstream_cert | getUpstreamCert
[**HasEntitlement**](EntitlementsApi.md#HasEntitlement) | **Get** /entitlements/consumer/{consumer_uuid}/product/{product_id} | hasEntitlement
[**ListAllForConsumer**](EntitlementsApi.md#ListAllForConsumer) | **Get** /entitlements | listAllForConsumer
[**MigrateEntitlement**](EntitlementsApi.md#MigrateEntitlement) | **Put** /entitlements/{entitlement_id}/migrate | migrateEntitlement
[**RegenerateEntitlementCertificatesForProduct**](EntitlementsApi.md#RegenerateEntitlementCertificatesForProduct) | **Put** /entitlements/product/{product_id} | regenerateEntitlementCertificatesForProduct
[**Unbind**](EntitlementsApi.md#Unbind) | **Delete** /entitlements/{dbid} | unbind
[**UpdateEntitlement**](EntitlementsApi.md#UpdateEntitlement) | **Put** /entitlements/{entitlement_id} | updateEntitlement


# **GetEntitlement**
> EntitlementDto GetEntitlement(ctx, entitlementId)
getEntitlement

Retrieves a single Entitlement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entitlementId** | **string**|  | 

### Return type

[**EntitlementDto**](EntitlementDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpstreamCert**
> string GetUpstreamCert(ctx, dbid)
getUpstreamCert

Retrieves a Subscription Certificate.  We can't return CdnInfo at this time, but when the time comes this is the implementation we want to start from. It will require changes to thumbslug.  will also @Produces(MediaType.APPLICATION_JSON)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbid** | **string**|  | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HasEntitlement**
> EntitlementDto HasEntitlement(ctx, consumerUuid, productId)
hasEntitlement

Checks Consumer for Product Entitlement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
  **productId** | **string**|  | 

### Return type

[**EntitlementDto**](EntitlementDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllForConsumer**
> []EntitlementDto ListAllForConsumer(ctx, optional)
listAllForConsumer

Retrieves list of Entitlements

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAllForConsumerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAllForConsumerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consumer** | **optional.String**|  | 
 **matches** | **optional.String**|  | 
 **attribute** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]EntitlementDto**](EntitlementDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrateEntitlement**
> MigrateEntitlement(ctx, entitlementId, optional)
migrateEntitlement

Migrate entitlements from one distributor consumer to another. Can specify full or partial quantity. No specified quantity will lead to full migration of the entitlement.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entitlementId** | **string**|  | 
 **optional** | ***MigrateEntitlementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrateEntitlementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toConsumer** | **optional.String**|  | 
 **quantity** | **optional.Int32**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateEntitlementCertificatesForProduct**
> JobDetail RegenerateEntitlementCertificatesForProduct(ctx, productId, optional)
regenerateEntitlementCertificatesForProduct

Regenerates the Entitlement Certificates for a Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | **string**|  | 
 **optional** | ***RegenerateEntitlementCertificatesForProductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegenerateEntitlementCertificatesForProductOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lazyRegen** | **optional.Bool**|  | [default to true]

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unbind**
> Unbind(ctx, dbid)
unbind

Deletes an Entitlement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dbid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEntitlement**
> UpdateEntitlement(ctx, entitlementId, entitlementDto)
updateEntitlement

Updates an Entitlement. This only works for the quantity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entitlementId** | **string**|  | 
  **entitlementDto** | [**EntitlementDto**](EntitlementDto.md)|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


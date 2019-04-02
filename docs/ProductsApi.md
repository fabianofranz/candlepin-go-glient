# \ProductsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBatchContent**](ProductsApi.md#AddBatchContent) | **Post** /products/{product_uuid}/batch_content | addBatchContent
[**AddContent**](ProductsApi.md#AddContent) | **Post** /products/{product_uuid}/content/{content_id} | addContent
[**CreateProduct**](ProductsApi.md#CreateProduct) | **Post** /products | createProduct
[**DeleteProduct**](ProductsApi.md#DeleteProduct) | **Delete** /products/{product_uuid} | deleteProduct
[**GetProduct**](ProductsApi.md#GetProduct) | **Get** /products/{product_uuid} | getProduct
[**GetProductCertificate**](ProductsApi.md#GetProductCertificate) | **Get** /products/{product_uuid}/certificate | getProductCertificate
[**GetProductOwners**](ProductsApi.md#GetProductOwners) | **Get** /products/owners | getProductOwners
[**RefreshPoolsForProduct**](ProductsApi.md#RefreshPoolsForProduct) | **Put** /products/subscriptions | refreshPoolsForProduct
[**RemoveContent**](ProductsApi.md#RemoveContent) | **Delete** /products/{product_uuid}/content/{content_id} | removeContent
[**UpdateProduct**](ProductsApi.md#UpdateProduct) | **Put** /products/{product_uuid} | updateProduct


# **AddBatchContent**
> ProductDto AddBatchContent(ctx, productUuid, optional)
addBatchContent

Adds Content to a Product Batch mode @deprecated Use per-org version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productUuid** | **string**|  | 
 **optional** | ***AddBatchContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddBatchContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**optional.Interface of map[string]bool**](bool.md)|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddContent**
> ProductDto AddContent(ctx, productUuid, contentId, optional)
addContent

Adds Content to a Product. Single mode @deprecated Use per-org version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productUuid** | **string**|  | 
  **contentId** | **string**|  | 
 **optional** | ***AddContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabled** | **optional.Bool**|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProduct**
> ProductDto CreateProduct(ctx, optional)
createProduct

Creates a Product. Returns either the new created Product or the Product that already existed. @deprecated Use per-org version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateProductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateProductOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productDto** | [**optional.Interface of ProductDto**](ProductDto.md)|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProduct**
> DeleteProduct(ctx, productUuid)
deleteProduct

Removes a Product @deprecated Use per-org version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productUuid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProduct**
> ProductDto GetProduct(ctx, productUuid)
getProduct

Retrieves a single Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productUuid** | **string**|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductCertificate**
> ProductCertificateDto GetProductCertificate(ctx, productUuid)
getProductCertificate

Retreives a Certificate for a Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productUuid** | **string**|  | 

### Return type

[**ProductCertificateDto**](ProductCertificateDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductOwners**
> []Owner GetProductOwners(ctx, product)
getProductOwners

Retrieves a list of Owners by Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **product** | [**[]string**](string.md)| Multiple product UUIDs | 

### Return type

[**[]Owner**](Owner.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshPoolsForProduct**
> []JobDetail RefreshPoolsForProduct(ctx, product, optional)
refreshPoolsForProduct

Refreshes Pools by Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **product** | [**[]string**](string.md)| Multiple product UUIDs | 
 **optional** | ***RefreshPoolsForProductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefreshPoolsForProductOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lazyRegen** | **optional.Bool**|  | [default to true]

### Return type

[**[]JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveContent**
> RemoveContent(ctx, productUuid, contentId)
removeContent

Removes Content from a Product @deprecated Use per-org version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productUuid** | **string**|  | 
  **contentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProduct**
> ProductDto UpdateProduct(ctx, productUuid, productDto)
updateProduct

Updates a Product @deprecated Use per-org version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productUuid** | **string**|  | 
  **productDto** | [**ProductDto**](ProductDto.md)|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


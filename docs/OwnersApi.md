# \OwnersApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBatchContent**](OwnersApi.md#AddBatchContent) | **Post** /owners/{owner_key}/products/{product_id}/batch_content | addBatchContent
[**AddContent**](OwnersApi.md#AddContent) | **Post** /owners/{owner_key}/products/{product_id}/content/{content_id} | addContent
[**CountConsumers**](OwnersApi.md#CountConsumers) | **Get** /owners/{owner_key}/consumers/count | consumers count
[**CreateActivationKey**](OwnersApi.md#CreateActivationKey) | **Post** /owners/{owner_key}/activation_keys | Create Activation Key
[**CreateBatchContent**](OwnersApi.md#CreateBatchContent) | **Post** /owners/{owner_key}/content/batch | createBatchContent
[**CreateContent**](OwnersApi.md#CreateContent) | **Post** /owners/{owner_key}/content | createContent
[**CreateEnv**](OwnersApi.md#CreateEnv) | **Post** /owners/{owner_key}/environments | Create environment
[**CreateOwner**](OwnersApi.md#CreateOwner) | **Post** /owners | Create Owner
[**CreatePool**](OwnersApi.md#CreatePool) | **Post** /owners/{owner_key}/pools | Create Pool
[**CreateProduct**](OwnersApi.md#CreateProduct) | **Post** /owners/{owner_key}/products | createProduct
[**CreateSubscription**](OwnersApi.md#CreateSubscription) | **Post** /owners/{owner_key}/subscriptions | Create Subscription
[**CreateUeberCertificate**](OwnersApi.md#CreateUeberCertificate) | **Post** /owners/{owner_key}/uebercert | Create Ueber Entitlement Certificate
[**DeleteLogLevel**](OwnersApi.md#DeleteLogLevel) | **Delete** /owners/{owner_key}/log | Remove Log Level
[**DeleteOwner**](OwnersApi.md#DeleteOwner) | **Delete** /owners/{owner_key} | Delete Owner
[**DeleteProduct**](OwnersApi.md#DeleteProduct) | **Delete** /owners/{owner_key}/products/{product_id} | deleteProduct
[**GetContent**](OwnersApi.md#GetContent) | **Get** /owners/{owner_key}/content/{content_id} | getContent
[**GetEvents**](OwnersApi.md#GetEvents) | **Get** /owners/{owner_key}/events | Get Events
[**GetHypervisors**](OwnersApi.md#GetHypervisors) | **Get** /owners/{owner_key}/hypervisors | Get Hypervisors
[**GetImports**](OwnersApi.md#GetImports) | **Get** /owners/{owner_key}/imports | Get Imports
[**GetOwner**](OwnersApi.md#GetOwner) | **Get** /owners/{owner_key} | Get Owner
[**GetOwnerAtomFeed**](OwnersApi.md#GetOwnerAtomFeed) | **Get** /owners/{owner_key}/atom | Get Atom Feed
[**GetOwnerInfo**](OwnersApi.md#GetOwnerInfo) | **Get** /owners/{owner_key}/info | Get Owner Info
[**GetProduct**](OwnersApi.md#GetProduct) | **Get** /owners/{owner_key}/products/{product_id} | getProduct
[**GetProductCertificate**](OwnersApi.md#GetProductCertificate) | **Get** /owners/{owner_key}/products/{product_id}/certificate | getProductCertificate
[**GetSubscriptions**](OwnersApi.md#GetSubscriptions) | **Get** /owners/{owner_key}/subscriptions | List Subscriptions
[**GetSyspurpose**](OwnersApi.md#GetSyspurpose) | **Get** /owners/{owner_key}/system_purpose | getSyspurpose
[**GetUeberCertificate**](OwnersApi.md#GetUeberCertificate) | **Get** /owners/{owner_key}/uebercert | Get Ueber Entitlement Certificate
[**GetUpstreamConsumers**](OwnersApi.md#GetUpstreamConsumers) | **Get** /owners/{owner_key}/upstream_consumers | Get Upstream Consumers
[**HealEntire**](OwnersApi.md#HealEntire) | **Post** /owners/{owner_key}/entitlements | Heal owner
[**ImportManifest**](OwnersApi.md#ImportManifest) | **Post** /owners/{owner_key}/imports | Import Manifest
[**ImportManifestAsync**](OwnersApi.md#ImportManifestAsync) | **Post** /owners/{owner_key}/imports/async | Import Manifest Asynchronously
[**List**](OwnersApi.md#List) | **Get** /owners | List Owners
[**ListConsumers**](OwnersApi.md#ListConsumers) | **Get** /owners/{owner_key}/consumers | List Consumers
[**ListContent**](OwnersApi.md#ListContent) | **Get** /owners/{owner_key}/content | list
[**ListEnvironments**](OwnersApi.md#ListEnvironments) | **Get** /owners/{owner_key}/environments | List environments
[**ListPools**](OwnersApi.md#ListPools) | **Get** /owners/{owner_key}/pools | List Pools
[**ListProducts**](OwnersApi.md#ListProducts) | **Get** /owners/{owner_key}/products | List Products for an Owner
[**OwnerActivationKeys**](OwnersApi.md#OwnerActivationKeys) | **Get** /owners/{owner_key}/activation_keys | Owner Activation Keys
[**OwnerEntitlements**](OwnersApi.md#OwnerEntitlements) | **Get** /owners/{owner_key}/entitlements | List Owner Entitlements
[**OwnerServiceLevels**](OwnersApi.md#OwnerServiceLevels) | **Get** /owners/{owner_key}/servicelevels | Get Service Levels
[**RefreshPools**](OwnersApi.md#RefreshPools) | **Put** /owners/{owner_key}/subscriptions | Update Subscription
[**RefreshPoolsForProduct**](OwnersApi.md#RefreshPoolsForProduct) | **Put** /owners/{owner_key}/products/{product_id}/subscriptions | refreshPoolsForProduct
[**Remove**](OwnersApi.md#Remove) | **Delete** /owners/{owner_key}/content/{content_id} | remove
[**RemoveBatchContent**](OwnersApi.md#RemoveBatchContent) | **Delete** /owners/{owner_key}/products/{product_id}/batch_content | addBatchContent
[**RemoveContent**](OwnersApi.md#RemoveContent) | **Delete** /owners/{owner_key}/products/{product_id}/content/{content_id} | removeContent
[**SetLogLevel**](OwnersApi.md#SetLogLevel) | **Put** /owners/{owner_key}/log | Set Log Level
[**UndoImports**](OwnersApi.md#UndoImports) | **Delete** /owners/{owner_key}/imports | Undo Imports
[**UpdateContent**](OwnersApi.md#UpdateContent) | **Put** /owners/{owner_key}/content/{content_id} | updateContent
[**UpdateOwner**](OwnersApi.md#UpdateOwner) | **Put** /owners/{owner_key} | Update Owner
[**UpdatePool**](OwnersApi.md#UpdatePool) | **Put** /owners/{owner_key}/pools | Update Pool
[**UpdateProduct**](OwnersApi.md#UpdateProduct) | **Put** /owners/{owner_key}/products/{product_id} | updateProduct
[**UpdateSubscription**](OwnersApi.md#UpdateSubscription) | **Put** /owners/subscriptions | Update Subscription


# **AddBatchContent**
> ProductDto AddBatchContent(ctx, ownerKey, productId, requestBody)
addBatchContent

Adds one or more Content entities to a Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **productId** | **string**|  | 
  **requestBody** | [**map[string]bool**](bool.md)|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddContent**
> ProductDto AddContent(ctx, ownerKey, productId, contentId, optional)
addContent

Adds a single Content to a Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **productId** | **string**|  | 
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

# **CountConsumers**
> int32 CountConsumers(ctx, ownerKey, optional)
consumers count

Retrieve a count of Consumers for the Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***CountConsumersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CountConsumersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of []string**](string.md)|  | 
 **sku** | [**optional.Interface of []string**](string.md)|  | 
 **subscriptionId** | [**optional.Interface of []string**](string.md)|  | 
 **contract** | [**optional.Interface of []string**](string.md)|  | 

### Return type

**int32**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateActivationKey**
> ActivationKeyDto CreateActivationKey(ctx, ownerKey, activationKeyDto)
Create Activation Key

Creates an Activation Key for the Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **activationKeyDto** | [**ActivationKeyDto**](ActivationKeyDto.md)|  | 

### Return type

[**ActivationKeyDto**](ActivationKeyDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBatchContent**
> []ContentDto CreateBatchContent(ctx, ownerKey, contentDTO)
createBatchContent

Creates Contents in bulk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **contentDTO** | [**[]ContentDto**](array.md)|  | 

### Return type

[**[]ContentDto**](ContentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContent**
> ContentDto CreateContent(ctx, ownerKey, contentDto)
createContent

Creates a Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **contentDto** | [**ContentDto**](ContentDto.md)|  | 

### Return type

[**ContentDto**](ContentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEnv**
> EnvironmentDto CreateEnv(ctx, ownerKey, environmentDto)
Create environment

Creates an Environment for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **environmentDto** | [**EnvironmentDto**](EnvironmentDto.md)|  | 

### Return type

[**EnvironmentDto**](EnvironmentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOwner**
> OwnerDto CreateOwner(ctx, ownerDto)
Create Owner

Creates an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerDto** | [**OwnerDto**](OwnerDto.md)|  | 

### Return type

[**OwnerDto**](OwnerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePool**
> Pool CreatePool(ctx, ownerKey, pool)
Create Pool

Creates a custom pool for an Owner. Floating pools are not tied to any upstream subscription, and are most commonly used for custom content delivery in Satellite. Also helps in on-site deployment testing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **pool** | [**Pool**](Pool.md)|  | 

### Return type

[**Pool**](pool.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProduct**
> ProductDto CreateProduct(ctx, ownerKey, optional)
createProduct

Creates a Product.  Returns either the new created Product or the Product that already existed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
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

# **CreateSubscription**
> Subscription CreateSubscription(ctx, ownerKey, optional)
Create Subscription

Creates a Subscription for an Owner DEPRECATED: Please create pools directly with POST /pools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***CreateSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateSubscriptionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscription** | [**optional.Interface of Subscription**](Subscription.md)|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUeberCertificate**
> Cert CreateUeberCertificate(ctx, ownerKey)
Create Ueber Entitlement Certificate

Creates an Ueber Entitlement Certificate. If a certificate already exists, it will be regenerated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**Cert**](cert.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogLevel**
> DeleteLogLevel(ctx, ownerKey)
Remove Log Level

Remove the Log Level of an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOwner**
> DeleteOwner(ctx, ownerKey, optional)
Delete Owner

Removes an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***DeleteOwnerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteOwnerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revoke** | **optional.Bool**|  | [default to true]
 **force** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProduct**
> DeleteProduct(ctx, ownerKey, productId)
deleteProduct

Removes a Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **productId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContent**
> ContentDto GetContent(ctx, ownerKey, contentId)
getContent

Retrieves a single Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **contentId** | **string**|  | 

### Return type

[**ContentDto**](ContentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvents**
> []EventDto GetEvents(ctx, ownerKey)
Get Events

Retrieves a list of Events for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**[]EventDto**](EventDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHypervisors**
> []Consumer GetHypervisors(ctx, ownerKey, optional)
Get Hypervisors

Retrieves a list of Hypervisors for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***GetHypervisorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetHypervisorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hypervisorId** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]Consumer**](Consumer.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImports**
> []ImportRecord GetImports(ctx, ownerKey)
Get Imports

 Retrieves a list of Import Records for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**[]ImportRecord**](ImportRecord.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOwner**
> OwnerDto GetOwner(ctx, ownerKey)
Get Owner

Retrieves a single Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**OwnerDto**](OwnerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOwnerAtomFeed**
> Feed GetOwnerAtomFeed(ctx, ownerKey)
Get Atom Feed

Retrieves an Event Atom Feed for an owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**Feed**](feed.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/atom+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOwnerInfo**
> OwnerInfo GetOwnerInfo(ctx, ownerKey)
Get Owner Info

Retrieves the Owner Info for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**OwnerInfo**](OwnerInfo.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProduct**
> ProductDto GetProduct(ctx, ownerKey, productId)
getProduct

Retrieves a single Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **productId** | **string**|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductCertificate**
> ProductCertificateDto GetProductCertificate(ctx, ownerKey, productId)
getProductCertificate

Retrieves a Certificate for a Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **productId** | **string**| Numeric product identifier | 

### Return type

[**ProductCertificateDto**](ProductCertificateDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptions**
> []Subscription GetSubscriptions(ctx, ownerKey)
List Subscriptions

Retrieves a list of Subscriptions for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSyspurpose**
> SystemPurposeAttributesDto GetSyspurpose(ctx, ownerKey)
getSyspurpose

Retrieves the system purpose settings available to an owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**SystemPurposeAttributesDto**](SystemPurposeAttributesDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUeberCertificate**
> Cert GetUeberCertificate(ctx, ownerKey)
Get Ueber Entitlement Certificate

Retrieves the Ueber Entitlement Certificate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**Cert**](cert.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpstreamConsumers**
> []UpstreamConsumerDto GetUpstreamConsumers(ctx, ownerKey)
Get Upstream Consumers

 Retrieves a list of Upstream Consumers for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**[]UpstreamConsumerDto**](UpstreamConsumerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HealEntire**
> JobDetail HealEntire(ctx, ownerKey)
Heal owner

Starts an asynchronous healing for the given Owner. At the end of the process the idea is that all of the consumers in the owned by the Owner will be up to date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**| ownerKey id of the owner to be healed. | 

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportManifest**
> ImportRecord ImportManifest(ctx, ownerKey, optional)
Import Manifest

Imports a manifest zip file for the given organization. This will bring in any products, content, and subscriptions that were assigned to the distributor who generated the manifest.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***ImportManifestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportManifestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | [**optional.Interface of []string**](string.md)|  | 
 **preamble** | **optional.String**|  | 
 **parts** | [**optional.Interface of []InputPart**](InputPart.md)|  | 

### Return type

[**ImportRecord**](ImportRecord.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportManifestAsync**
> JobDetail ImportManifestAsync(ctx, ownerKey, optional)
Import Manifest Asynchronously

Initiates an asynchronous manifest import for the given organization. This will bring in any products, content, and subscriptions that were assigned to the distributor who generated the manifest.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***ImportManifestAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportManifestAsyncOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | [**optional.Interface of []string**](string.md)|  | 
 **preamble** | **optional.String**|  | 
 **parts** | [**optional.Interface of []InputPart**](InputPart.md)|  | 

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []OwnerDto List(ctx, optional)
List Owners

Retrieves a list of Owners

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **optional.String**|  | 

### Return type

[**[]OwnerDto**](OwnerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConsumers**
> []Consumer ListConsumers(ctx, ownerKey, optional)
List Consumers

Retrieve a list of Consumers for the Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***ListConsumersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListConsumersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**|  | 
 **type_** | [**optional.Interface of []string**](string.md)|  | 
 **uuid** | [**optional.Interface of []string**](string.md)|  | 
 **hypervisorId** | [**optional.Interface of []string**](string.md)|  | 
 **fact** | [**optional.Interface of []string**](string.md)|  | 
 **sku** | [**optional.Interface of []string**](string.md)|  | 
 **subscriptionId** | [**optional.Interface of []string**](string.md)|  | 
 **contract** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]Consumer**](Consumer.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContent**
> []Content ListContent(ctx, ownerKey)
list

Retrieves list of Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**[]Content**](Content.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnvironments**
> CandlepinQueryEnvironmentDto ListEnvironments(ctx, ownerKey, optional)
List environments

Retrieves a list of Environments for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***ListEnvironmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListEnvironmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Environment name filter to search for. | 

### Return type

[**CandlepinQueryEnvironmentDto**](CandlepinQueryEnvironmentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPools**
> []Pool ListPools(ctx, ownerKey, optional)
List Pools

Retrieves a list of Pools for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***ListPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListPoolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consumer** | **optional.String**|  | 
 **activationKey** | **optional.String**|  | 
 **product** | **optional.String**|  | 
 **subscription** | **optional.String**|  | 
 **listall** | **optional.Bool**| Include pools that are not suited to the unit&#39;s facts. | [default to false]
 **activeon** | **optional.Time**| Date to use as current time for lookup criteria. Defaults to current date if not specified. | [default to now]
 **matches** | [**optional.Interface of []string**](string.md)| Find pools matching the given pattern in a variety of fields; * and ? wildcards are supported; may be specified multiple times | 
 **attribute** | [**optional.Interface of []string**](string.md)| The attributes to return based on the specified types. | 
 **addFuture** | **optional.Bool**| When set to true, it will add future dated pools to the result, based on the activeon date. | [default to false]
 **onlyFuture** | **optional.Bool**| When set to true, it will return only future dated pools to the result, based on the activeon date. | [default to false]
 **after** | **optional.Time**| Will only return pools with a start date after the supplied date. Overrides the activeOn date. | 
 **poolid** | [**optional.Interface of []string**](string.md)| One or more pool IDs to use to filter the output; only pools with IDs matching those provided will be returned; may be specified multiple times | 

### Return type

[**[]Pool**](pool.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProducts**
> []Product ListProducts(ctx, ownerKey, optional)
List Products for an Owner

Retrieves a list of Products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***ListProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListProductsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]Product**](Product.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OwnerActivationKeys**
> CandlepinQueryActivationKeyDto OwnerActivationKeys(ctx, ownerKey, optional)
Owner Activation Keys

Retrieves a list of Activation Keys for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***OwnerActivationKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OwnerActivationKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 

### Return type

[**CandlepinQueryActivationKeyDto**](CandlepinQueryActivationKeyDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OwnerEntitlements**
> []EntitlementDto OwnerEntitlements(ctx, ownerKey, optional)
List Owner Entitlements

Retrieves the list of Entitlements for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***OwnerEntitlementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OwnerEntitlementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | **optional.String**|  | 
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

# **OwnerServiceLevels**
> []string OwnerServiceLevels(ctx, ownerKey, optional)
Get Service Levels

Retrieves a list of Support Levels for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**| ownerKey id of the owner whose support levels are sought. | 
 **optional** | ***OwnerServiceLevelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OwnerServiceLevelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exempt** | **optional.String**|  | [default to false]

### Return type

**[]string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshPools**
> JobDetail RefreshPools(ctx, ownerKey, optional)
Update Subscription

Tickle an owner to have all of their entitlement pools synced with their subscriptions. This method (and the one below may not be entirely RESTful, as the updated data is not supplied as an argument. This API call is only relevant in a top level hosted deployment where subscriptions and products are sourced from adapters. Calling this in an on-site deployment is just a no-op.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***RefreshPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefreshPoolsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoCreateOwner** | **optional.Bool**|  | [default to false]
 **lazyRegen** | **optional.Bool**|  | [default to true]

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshPoolsForProduct**
> JobDetail RefreshPoolsForProduct(ctx, ownerKey, productId, optional)
refreshPoolsForProduct

Refreshes Pools by Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **productId** | **string**|  | 
 **optional** | ***RefreshPoolsForProductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefreshPoolsForProductOpts struct

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

# **Remove**
> Remove(ctx, ownerKey, contentId)
remove

Deletes a Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **contentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveBatchContent**
> ProductDto RemoveBatchContent(ctx, ownerKey, productId, requestBody)
addBatchContent

Adds one or more Content entities to a Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **productId** | **string**|  | 
  **requestBody** | [**[]string**](array.md)|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveContent**
> ProductDto RemoveContent(ctx, ownerKey, productId, contentId)
removeContent

Removes a single Content from a Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **productId** | **string**|  | 
  **contentId** | **string**|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetLogLevel**
> OwnerDto SetLogLevel(ctx, ownerKey, optional)
Set Log Level

Sets the Log Level for an Owner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
 **optional** | ***SetLogLevelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetLogLevelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **level** | **optional.String**|  | [default to DEBUG]

### Return type

[**OwnerDto**](OwnerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UndoImports**
> JobDetail UndoImports(ctx, ownerKey)
Undo Imports

Removes Imports for an Owner. Cleans out all imported subscriptions and triggers a background refresh pools. Link to an upstream distributor is removed for the owner, so they can import from another distributor. Other owners can also now import the manifests originally used in this owner. This  call does not differentiate between any specific import, it just destroys all subscriptions with an upstream pool ID, essentially anything from an import. Custom subscriptions will be left alone. Imports do carry rules and product information which is global to the candlepin server. This import data is *not* undone, we assume that updates to this data can be safely kept. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContent**
> ContentDto UpdateContent(ctx, ownerKey, contentId, contentDto)
updateContent

Updates a Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **contentId** | **string**|  | 
  **contentDto** | [**ContentDto**](ContentDto.md)|  | 

### Return type

[**ContentDto**](ContentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOwner**
> OwnerDto UpdateOwner(ctx, ownerKey, ownerDto)
Update Owner

To un-set the defaultServiceLevel for an owner, submit an empty string.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **ownerDto** | [**OwnerDto**](OwnerDto.md)|  | 

### Return type

[**OwnerDto**](OwnerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePool**
> UpdatePool(ctx, ownerKey, pool)
Update Pool

Updates a pool for an Owner. assumes this is a normal pool, and errors out otherwise cause we cannot create master pools from bonus pools 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **pool** | [**Pool**](Pool.md)|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProduct**
> ProductDto UpdateProduct(ctx, ownerKey, productId, productDto)
updateProduct

Updates a Product

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerKey** | **string**|  | 
  **productId** | **string**|  | 
  **productDto** | [**ProductDto**](ProductDto.md)|  | 

### Return type

[**ProductDto**](ProductDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscription**
> UpdateSubscription(ctx, subscription)
Update Subscription

Updates a Subscription for an Owner.  Please update pools directly with POST /pools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscription** | [**Subscription**](Subscription.md)|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


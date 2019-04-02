# \EnvironmentsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](EnvironmentsApi.md#Create) | **Post** /environments/{env_id}/consumers | create
[**DeleteEnv**](EnvironmentsApi.md#DeleteEnv) | **Delete** /environments/{env_id} | deleteEnv
[**DemoteContent**](EnvironmentsApi.md#DemoteContent) | **Delete** /environments/{env_id}/content | demoteContent
[**GetEnv**](EnvironmentsApi.md#GetEnv) | **Get** /environments/{env_id} | getEnv
[**GetEnvironments**](EnvironmentsApi.md#GetEnvironments) | **Get** /environments | getEnvironments
[**PromoteContent**](EnvironmentsApi.md#PromoteContent) | **Post** /environments/{env_id}/content | promoteContent


# **Create**
> ConsumerDto Create(ctx, envId, consumerDto, optional)
create

Creates an Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envId** | **string**|  | 
  **consumerDto** | [**ConsumerDto**](ConsumerDto.md)|  | 
 **optional** | ***CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **username** | **optional.String**|  | 
 **owner** | **optional.String**|  | 
 **activationKeys** | **optional.String**|  | 

### Return type

[**ConsumerDto**](ConsumerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnv**
> DeleteEnv(ctx, envId)
deleteEnv

Deletes an environment. WARNING: this will delete all consumers in the environment and revoke their entitlement certificates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DemoteContent**
> JobDetail DemoteContent(ctx, envId, optional)
demoteContent

Demotes a Content from an Environment. Consumer's registered to this environment will no see this content in their entitlement certificates. (after they are regenerated and synced to clients) This call accepts multiple content IDs to demote at once, allowing us to mass demote, then trigger a cert regeneration. NOTE: This call expects the actual content IDs, *not* the ID created for each EnvironmentContent object created after a promotion. This is to help integrate with other management apps which should not have to track/lookup a specific ID for the content to demote.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envId** | **string**|  | 
 **optional** | ***DemoteContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DemoteContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | [**optional.Interface of []string**](string.md)|  | 
 **lazyRegen** | **optional.Bool**|  | [default to true]

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnv**
> EnvironmentDto GetEnv(ctx, envId)
getEnv

Retrieves a single Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envId** | **string**|  | 

### Return type

[**EnvironmentDto**](EnvironmentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironments**
> []Environment GetEnvironments(ctx, )
getEnvironments

Lists the Environments.  Only available to super admins.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Environment**](Environment.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PromoteContent**
> JobDetail PromoteContent(ctx, envId, environmentContent, optional)
promoteContent

Promotes a Content into an Environment. This call accepts multiple content sets to promote at once, after which all affected certificates for consumers in the environment will be regenerated. Consumers registered to this environment will now receive this content in their entitlement certificates. Because the certificate regeneraiton can be quite time consuming, this is done as an asynchronous job. The content will be promoted and immediately available for new entitlements, but existing entitlements could take some time to be regenerated and sent down to clients as they check in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envId** | **string**|  | 
  **environmentContent** | [**[]EnvironmentContent**](array.md)|  | 
 **optional** | ***PromoteContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PromoteContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lazyRegen** | **optional.Bool**|  | [default to true]

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


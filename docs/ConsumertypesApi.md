# \ConsumertypesApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ConsumertypesApi.md#Create) | **Post** /consumertypes | create
[**DeleteConsumerType**](ConsumertypesApi.md#DeleteConsumerType) | **Delete** /consumertypes/{id} | deleteConsumerType
[**GetConsumerType**](ConsumertypesApi.md#GetConsumerType) | **Get** /consumertypes/{id} | getConsumerType
[**List**](ConsumertypesApi.md#List) | **Get** /consumertypes | list
[**Update**](ConsumertypesApi.md#Update) | **Put** /consumertypes/{id} | update


# **Create**
> ConsumerTypeDto Create(ctx, consumerTypeDto)
create

Creates a Consumer Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerTypeDto** | [**ConsumerTypeDto**](ConsumerTypeDto.md)|  | 

### Return type

[**ConsumerTypeDto**](ConsumerTypeDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConsumerType**
> DeleteConsumerType(ctx, id)
deleteConsumerType

Removes a Consumer Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumerType**
> ConsumerTypeDto GetConsumerType(ctx, id)
getConsumerType

Retrieves a single Consumer Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**ConsumerTypeDto**](ConsumerTypeDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []ConsumerTypeDto List(ctx, )
list

Retrieves a list of Consumer Types

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConsumerTypeDto**](ConsumerTypeDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> ConsumerTypeDto Update(ctx, consumerTypeDto)
update

Updates a Consumer Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerTypeDto** | [**ConsumerTypeDto**](ConsumerTypeDto.md)|  | 

### Return type

[**ConsumerTypeDto**](ConsumerTypeDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


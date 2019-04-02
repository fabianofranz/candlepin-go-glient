# \CdnApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CdnApi.md#Create) | **Post** /cdn | create
[**Delete**](CdnApi.md#Delete) | **Delete** /cdn/{label} | delete
[**GetContentDeliveryNetworks**](CdnApi.md#GetContentDeliveryNetworks) | **Get** /cdn | getContentDeliveryNetworks
[**Update**](CdnApi.md#Update) | **Put** /cdn/{label} | update


# **Create**
> CdnDto Create(ctx, cdnDto)
create

Creates a CDN @return a Cdn object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cdnDto** | [**CdnDto**](CdnDto.md)|  | 

### Return type

[**CdnDto**](CdnDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, label)
delete

Removes a CDN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **label** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentDeliveryNetworks**
> []CdnDto GetContentDeliveryNetworks(ctx, )
getContentDeliveryNetworks

Retrieves a list of CDN's

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CdnDto**](CdnDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> CdnDto Update(ctx, label, cdnDto)
update

Updates a CDN @return a Cdn object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **label** | **string**|  | 
  **cdnDto** | [**CdnDto**](CdnDto.md)|  | 

### Return type

[**CdnDto**](CdnDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


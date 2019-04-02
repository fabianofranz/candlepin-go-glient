# \CrlApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentCrl**](CrlApi.md#GetCurrentCrl) | **Get** /crl | getCurrentCrl
[**Unrevoke**](CrlApi.md#Unrevoke) | **Delete** /crl | unrevoke


# **GetCurrentCrl**
> string GetCurrentCrl(ctx, )
getCurrentCrl

Retrieves the Certificate Revocation List

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unrevoke**
> Unrevoke(ctx, optional)
unrevoke

Deletes a Certificate from the Revocation List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UnrevokeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UnrevokeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serial** | [**optional.Interface of []string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


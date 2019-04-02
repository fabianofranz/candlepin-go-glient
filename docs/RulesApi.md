# \RulesApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](RulesApi.md#Delete) | **Delete** /rules | delete
[**Get**](RulesApi.md#Get) | **Get** /rules | get
[**Upload**](RulesApi.md#Upload) | **Post** /rules | upload


# **Delete**
> Delete(ctx, )
delete

Removes the Rules  Deletes any uploaded rules, uses bundled rules instead

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> string Get(ctx, )
get

Retrieves the Rules

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Upload**
> string Upload(ctx, optional)
upload

Uploads the Rules Returns a copy of the uploaded rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UploadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **optional.String**|  | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


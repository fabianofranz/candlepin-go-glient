# \ContentApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBatchContent**](ContentApi.md#CreateBatchContent) | **Post** /content/batch | createBatchContent
[**CreateContent**](ContentApi.md#CreateContent) | **Post** /content | createContent
[**GetContent**](ContentApi.md#GetContent) | **Get** /content/{content_uuid} | getContent
[**List**](ContentApi.md#List) | **Get** /content | list
[**Remove**](ContentApi.md#Remove) | **Delete** /content/{content_uuid} | remove
[**UpdateContent**](ContentApi.md#UpdateContent) | **Put** /content/{content_uuid} | updateContent


# **CreateBatchContent**
> IterableContentDto CreateBatchContent(ctx, optional)
createBatchContent

Creates Contents in bulk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateBatchContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateBatchContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentDTO** | [**optional.Interface of []ContentDto**](array.md)|  | 

### Return type

[**IterableContentDto**](IterableContentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContent**
> ContentDto CreateContent(ctx, optional)
createContent

Creates a Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentDto** | [**optional.Interface of ContentDto**](ContentDto.md)|  | 

### Return type

[**ContentDto**](ContentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContent**
> ContentDto GetContent(ctx, contentUuid)
getContent

Retrieves a single Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentUuid** | **string**|  | 

### Return type

[**ContentDto**](ContentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []ContentDto List(ctx, )
list

Retrieves list of Content

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ContentDto**](ContentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Remove**
> Remove(ctx, contentUuid)
remove

Deletes a Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentUuid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContent**
> ContentDto UpdateContent(ctx, contentUuid, optional)
updateContent

Updates a Content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentUuid** | **string**|  | 
 **optional** | ***UpdateContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentDto** | [**optional.Interface of ContentDto**](ContentDto.md)|  | 

### Return type

[**ContentDto**](ContentDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


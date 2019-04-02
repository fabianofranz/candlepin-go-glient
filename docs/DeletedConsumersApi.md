# \DeletedConsumersApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListByDate**](DeletedConsumersApi.md#ListByDate) | **Get** /deleted_consumers | listByDate


# **ListByDate**
> []DeletedConsumer ListByDate(ctx, optional)
listByDate

Retrieves a list of Deleted Consumers By deletion date or all. List returned is the deleted Consumers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListByDateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListByDateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **optional.String**|  | 

### Return type

[**[]DeletedConsumer**](DeletedConsumer.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


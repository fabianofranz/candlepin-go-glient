# \EventsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /events/{uuid} | getEvent
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /events | listEvents


# **GetEvent**
> EventDto GetEvent(ctx, uuid)
getEvent

Retrieves a single Event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 

### Return type

[**EventDto**](EventDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEvents**
> []EventDto ListEvents(ctx, )
listEvents

Retrieves a list of Events

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EventDto**](EventDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


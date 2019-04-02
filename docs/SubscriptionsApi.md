# \SubscriptionsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSubscription**](SubscriptionsApi.md#ActivateSubscription) | **Post** /subscriptions | activateSubscription
[**DeleteSubscription**](SubscriptionsApi.md#DeleteSubscription) | **Delete** /subscriptions/{subscription_id} | deleteSubscription
[**GetSubCertAsPem**](SubscriptionsApi.md#GetSubCertAsPem) | **Get** /subscriptions/{subscription_id}/cert | getSubCertAsPem
[**GetSubscription**](SubscriptionsApi.md#GetSubscription) | **Get** /subscriptions/{subscription_id} | getSubscription
[**GetSubscriptions**](SubscriptionsApi.md#GetSubscriptions) | **Get** /subscriptions | getSubscriptions


# **ActivateSubscription**
> ActivateSubscription(ctx, consumerUuid, email, emailLocale)
activateSubscription

Activates a Subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
  **email** | **string**|  | 
  **emailLocale** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscription**
> DeleteSubscription(ctx, subscriptionId)
deleteSubscription

Removes a Subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubCertAsPem**
> string GetSubCertAsPem(ctx, subscriptionId)
getSubCertAsPem

Retrieves a Subscription Certificate As a PEM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**|  | 

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscription**
> Subscription GetSubscription(ctx, subscriptionId)
getSubscription

Retrieves a single Subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriptions**
> []Subscription GetSubscriptions(ctx, )
getSubscriptions

Retrieves a list of Subscriptions

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


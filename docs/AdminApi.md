# \AdminApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearProductCache**](AdminApi.md#ClearProductCache) | **Delete** /admin/cache/product | Clear product cache
[**GetQueueStats**](AdminApi.md#GetQueueStats) | **Get** /admin/queues | Get Queue Stats
[**Initialize**](AdminApi.md#Initialize) | **Get** /admin/init | initialize


# **ClearProductCache**
> ClearProductCache(ctx, )
Clear product cache

Clears the product cache

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueueStats**
> []QueueStatus GetQueueStats(ctx, )
Get Queue Stats

Basic information on the ActiveMQ queues and how many messages are pending in each.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]QueueStatus**](QueueStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Initialize**
> string Initialize(ctx, )
initialize

Initializes the Candlepin database. Currently this just creates the admin user for standalone deployments using the default user service adapter. It must be called once after candlepin is installed, repeat calls are not required, but will be harmless. The String returned is the description if the db was or already is initialized.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


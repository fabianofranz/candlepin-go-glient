# \JobsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Cancel**](JobsApi.md#Cancel) | **Delete** /jobs/{job_id} | cancel
[**GetSchedulerStatus**](JobsApi.md#GetSchedulerStatus) | **Get** /jobs/scheduler | getSchedulerStatus
[**GetStatus**](JobsApi.md#GetStatus) | **Get** /jobs/{job_id} | getStatus
[**GetStatusAndDeleteIfFinished**](JobsApi.md#GetStatusAndDeleteIfFinished) | **Post** /jobs/{job_id} | getStatusAndDeleteIfFinished
[**GetStatuses**](JobsApi.md#GetStatuses) | **Get** /jobs | getStatuses
[**Retrigger**](JobsApi.md#Retrigger) | **Post** /jobs/retrigger/{task} | retrigger
[**Schedule**](JobsApi.md#Schedule) | **Post** /jobs/schedule/{task} | schedule
[**SetSchedulerStatus**](JobsApi.md#SetSchedulerStatus) | **Post** /jobs/scheduler | setSchedulerStatus


# **Cancel**
> JobStatusDto Cancel(ctx, jobId)
cancel

Cancels a Job Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

[**JobStatusDto**](JobStatusDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchedulerStatus**
> SchedulerStatus GetSchedulerStatus(ctx, )
getSchedulerStatus

Retrieves the Scheduler Status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SchedulerStatus**](SchedulerStatus.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatus**
> JobStatusDto GetStatus(ctx, jobId, optional)
getStatus

Retrieves a single Job Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 
 **optional** | ***GetStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resultData** | **optional.Bool**|  | [default to false]

### Return type

[**JobStatusDto**](JobStatusDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatusAndDeleteIfFinished**
> JobStatusDto GetStatusAndDeleteIfFinished(ctx, jobId)
getStatusAndDeleteIfFinished

Retrieves a Job Status and Removes if finished

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

[**JobStatusDto**](JobStatusDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatuses**
> []JobStatus GetStatuses(ctx, optional)
getStatuses

Retrieves a list of Job Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetStatusesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **optional.String**|  | 
 **consumer** | **optional.String**|  | 
 **principal** | **optional.String**|  | 

### Return type

[**[]JobStatus**](JobStatus.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Retrigger**
> Retrigger(ctx, task)
retrigger

Re-trigger cron jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **task** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Schedule**
> JobStatusDto Schedule(ctx, task)
schedule

Fires cron jobs asynchronously and immediately

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **task** | **string**|  | 

### Return type

[**JobStatusDto**](JobStatusDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSchedulerStatus**
> SchedulerStatus SetSchedulerStatus(ctx, optional)
setSchedulerStatus

Updates the Scheduler Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetSchedulerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SetSchedulerStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **optional.Bool**|  | 

### Return type

[**SchedulerStatus**](SchedulerStatus.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


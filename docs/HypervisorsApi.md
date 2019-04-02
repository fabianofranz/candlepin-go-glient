# \HypervisorsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HypervisorUpdate**](HypervisorsApi.md#HypervisorUpdate) | **Post** /hypervisors | hypervisorUpdate
[**HypervisorUpdateAsync**](HypervisorsApi.md#HypervisorUpdateAsync) | **Post** /hypervisors/{owner} | hypervisorUpdateAsync


# **HypervisorUpdate**
> HypervisorCheckInResult HypervisorUpdate(ctx, optional)
hypervisorUpdate

Updates the list of Hypervisor Guests Allows agents such as virt-who to update its host list and associate the guests for each host. This is typically used when a host is unable to register to candlepin via subscription manager.  In situations where consumers already exist it is probably best not to allow creation of new hypervisor consumers.  Most consumers do not have a hypervisorId attribute, so that should be added manually when necessary by the management environment. @deprecated Use the asynchronous method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HypervisorUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HypervisorUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **optional.String**|  | 
 **createMissing** | **optional.Bool**| specify whether or not to create missing hypervisors.Default is true.  If false is specified, hypervisorIds that are not foundwill result in failed entries in the resulting HypervisorCheckInResult | [default to true]
 **requestBody** | [**optional.Interface of map[string][]GuestIdDto**](array.md)|  | 

### Return type

[**HypervisorCheckInResult**](HypervisorCheckInResult.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HypervisorUpdateAsync**
> JobDetail HypervisorUpdateAsync(ctx, owner, optional)
hypervisorUpdateAsync

Creates or Updates the list of Hypervisor hosts Allows agents such as virt-who to update hosts' information . This is typically used when a host is unable to register to candlepin via subscription manager. In situations where consumers already exist it is probably best not to allow creation of new hypervisor consumers.  Most consumers do not have a hypervisorId attribute, so that should be added manually when necessary by the management environment. Default is true.  If false is specified, hypervisorIds that are not found will result in a failed state of the job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **owner** | **string**|  | 
 **optional** | ***HypervisorUpdateAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HypervisorUpdateAsyncOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMissing** | **optional.Bool**| specify whether or not to create missing hypervisors.Default is true.  If false is specified, hypervisorIds that are not foundwill result in failed entries in the resulting HypervisorCheckInResult | [default to true]
 **reporterId** | **optional.String**|  | 
 **body** | **optional.String**|  | 

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


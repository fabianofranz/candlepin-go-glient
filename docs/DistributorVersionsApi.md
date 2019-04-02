# \DistributorVersionsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DistributorVersionsApi.md#Create) | **Post** /distributor_versions | create
[**Delete**](DistributorVersionsApi.md#Delete) | **Delete** /distributor_versions/{id} | delete
[**GetVersions**](DistributorVersionsApi.md#GetVersions) | **Get** /distributor_versions | getVersions
[**Update**](DistributorVersionsApi.md#Update) | **Put** /distributor_versions/{id} | update


# **Create**
> Distributorversion Create(ctx, distributorversion)
create

Creates a Distributor Version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **distributorversion** | [**Distributorversion**](Distributorversion.md)|  | 

### Return type

[**Distributorversion**](distributorversion.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, id)
delete

Deletes a Distributor Version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersions**
> []Distributorversion GetVersions(ctx, optional)
getVersions

Retrieves list of Distributor Versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetVersionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameSearch** | **optional.String**|  | 
 **capability** | **optional.String**|  | 

### Return type

[**[]Distributorversion**](distributorversion.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> Distributorversion Update(ctx, id, distributorversion)
update

Updates a Distributor Version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **distributorversion** | [**Distributorversion**](Distributorversion.md)|  | 

### Return type

[**Distributorversion**](distributorversion.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


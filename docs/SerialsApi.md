# \SerialsApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCertificateSerial**](SerialsApi.md#GetCertificateSerial) | **Get** /serials/{serial_id} | getCertificateSerial
[**GetCertificateSerials**](SerialsApi.md#GetCertificateSerials) | **Get** /serials | getCertificateSerials


# **GetCertificateSerial**
> CertificateSerialDto GetCertificateSerial(ctx, serialId)
getCertificateSerial

Retrieves single Certificate Serial

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serialId** | **int64**|  | 

### Return type

[**CertificateSerialDto**](CertificateSerialDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificateSerials**
> []CertificateSerial GetCertificateSerials(ctx, )
getCertificateSerials

Retrieves a list of Certificate Serials

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]CertificateSerial**](CertificateSerial.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


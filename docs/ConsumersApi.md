# \ConsumersApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContentOverrides**](ConsumersApi.md#AddContentOverrides) | **Put** /consumers/{consumer_uuid}/content_overrides | 
[**Bind**](ConsumersApi.md#Bind) | **Post** /consumers/{consumer_uuid}/entitlements | Bind Entitlements
[**ConsumerExists**](ConsumersApi.md#ConsumerExists) | **Head** /consumers/{consumer_uuid}/exists | 
[**Create**](ConsumersApi.md#Create) | **Post** /consumers | create
[**DeleteConsumer**](ConsumersApi.md#DeleteConsumer) | **Delete** /consumers/{consumer_uuid} | deleteConsumer
[**DeleteContentOverrides**](ConsumersApi.md#DeleteContentOverrides) | **Delete** /consumers/{consumer_uuid}/content_overrides | 
[**DeleteGuest**](ConsumersApi.md#DeleteGuest) | **Delete** /consumers/{consumer_uuid}/guestids/{guest_id} | deleteGuest
[**DownloadExistingExport**](ConsumersApi.md#DownloadExistingExport) | **Get** /consumers/{consumer_uuid}/export/{export_id} | Async Consumer Export (manifest) Download
[**DryBind**](ConsumersApi.md#DryBind) | **Get** /consumers/{consumer_uuid}/entitlements/dry-run | dryBind
[**ExportData**](ConsumersApi.md#ExportData) | **Get** /consumers/{consumer_uuid}/export | Consumer Export (manifest)
[**ExportDataAsync**](ConsumersApi.md#ExportDataAsync) | **Get** /consumers/{consumer_uuid}/export/async | Async Consumer Export (manifest)
[**GetComplianceStatus**](ConsumersApi.md#GetComplianceStatus) | **Get** /consumers/{consumer_uuid}/compliance | getComplianceStatus
[**GetComplianceStatusList**](ConsumersApi.md#GetComplianceStatusList) | **Get** /consumers/compliance | getComplianceStatusList
[**GetConsumer**](ConsumersApi.md#GetConsumer) | **Get** /consumers/{consumer_uuid} | getConsumer
[**GetConsumerAtomFeed**](ConsumersApi.md#GetConsumerAtomFeed) | **Get** /consumers/{consumer_uuid}/atom | getConsumerAtomFeed
[**GetConsumerEvents**](ConsumersApi.md#GetConsumerEvents) | **Get** /consumers/{consumer_uuid}/events | getConsumerEvents
[**GetContentAccessBody**](ConsumersApi.md#GetContentAccessBody) | **Get** /consumers/{consumer_uuid}/accessible_content | getContentAccessBody
[**GetContentOverrideList**](ConsumersApi.md#GetContentOverrideList) | **Get** /consumers/{consumer_uuid}/content_overrides | 
[**GetEntitlementCertificateSerials**](ConsumersApi.md#GetEntitlementCertificateSerials) | **Get** /consumers/{consumer_uuid}/certificates/serials | getEntitlementCertificateSerials
[**GetEntitlementCertificates**](ConsumersApi.md#GetEntitlementCertificates) | **Get** /consumers/{consumer_uuid}/certificates | getEntitlementCertificates
[**GetGuestId**](ConsumersApi.md#GetGuestId) | **Get** /consumers/{consumer_uuid}/guestids/{guest_id} | getGuestId
[**GetGuestIds**](ConsumersApi.md#GetGuestIds) | **Get** /consumers/{consumer_uuid}/guestids | getGuestIds
[**GetGuests**](ConsumersApi.md#GetGuests) | **Get** /consumers/{consumer_uuid}/guests | getGuests
[**GetHost**](ConsumersApi.md#GetHost) | **Get** /consumers/{consumer_uuid}/host | getHost
[**GetOwner**](ConsumersApi.md#GetOwner) | **Get** /consumers/{consumer_uuid}/owner | getOwner
[**GetRelease**](ConsumersApi.md#GetRelease) | **Get** /consumers/{consumer_uuid}/release | getRelease
[**GetSystemPurposeComplianceStatus**](ConsumersApi.md#GetSystemPurposeComplianceStatus) | **Get** /consumers/{consumer_uuid}/purpose_compliance | getSystemPurposeComplianceStatus
[**List**](ConsumersApi.md#List) | **Get** /consumers | list
[**ListEntitlements**](ConsumersApi.md#ListEntitlements) | **Get** /consumers/{consumer_uuid}/entitlements | listEntitlements
[**RegenerateEntitlementCertificates**](ConsumersApi.md#RegenerateEntitlementCertificates) | **Put** /consumers/{consumer_uuid}/certificates | regenerateEntitlementCertificates
[**RegenerateIdentityCertificates**](ConsumersApi.md#RegenerateIdentityCertificates) | **Post** /consumers/{consumer_uuid} | regenerateIdentityCertificates
[**RemoveDeletionRecord**](ConsumersApi.md#RemoveDeletionRecord) | **Delete** /consumers/{consumer_uuid}/deletionrecord | removeDeletionRecord
[**Unbind**](ConsumersApi.md#Unbind) | **Delete** /consumers/{consumer_uuid}/entitlements/{dbid} | unbind
[**UnbindAll**](ConsumersApi.md#UnbindAll) | **Delete** /consumers/{consumer_uuid}/entitlements | unbindAll
[**UnbindByPool**](ConsumersApi.md#UnbindByPool) | **Delete** /consumers/{consumer_uuid}/entitlements/pool/{pool_id} | unbindByPool
[**UnbindBySerial**](ConsumersApi.md#UnbindBySerial) | **Delete** /consumers/{consumer_uuid}/certificates/{serial} | unbindBySerial
[**UpdateConsumer**](ConsumersApi.md#UpdateConsumer) | **Put** /consumers/{consumer_uuid} | updateConsumer
[**UpdateGuest**](ConsumersApi.md#UpdateGuest) | **Put** /consumers/{consumer_uuid}/guestids/{guest_id} | updateGuest
[**UpdateGuests**](ConsumersApi.md#UpdateGuests) | **Put** /consumers/{consumer_uuid}/guestids | updateGuests


# **AddContentOverrides**
> []ContentOverride AddContentOverrides(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddContentOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddContentOverridesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentOverride** | [**optional.Interface of []ContentOverride**](array.md)|  | 

### Return type

[**[]ContentOverride**](ContentOverride.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Bind**
> Bind(ctx, consumerUuid, optional)
Bind Entitlements

If a pool ID is specified, we know we're binding to that exact pool. Specifying an entitle date in this case makes no sense and will throw an error. If a list of product IDs are specified, we attempt to auto-bind to subscriptions which will provide those products. An optional date can be specified allowing the consumer to get compliant for some date in the future. If no date is specified we assume the current date. If neither a pool nor an ID is specified, this is a healing request. The path is similar to the bind by products, but in this case we use the installed products on the consumer, and their current compliant status, to determine which product IDs should be requested. The entitle date is used the same as with bind by products. The response will contain a list of Entitlement objects if async is false, or a JobDetail object if async is true.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
 **optional** | ***BindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pool** | **optional.String**|  | 
 **product** | [**optional.Interface of []string**](string.md)|  | 
 **quantity** | **optional.Int32**|  | 
 **email** | **optional.String**|  | 
 **emailLocale** | **optional.String**|  | 
 **async** | **optional.Bool**|  | [default to false]
 **entitleDate** | **optional.String**|  | 
 **fromPool** | [**optional.Interface of []string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConsumerExists**
> ConsumerExists(ctx, consumerUuid)


Checks for the existence of a Consumer. This method is used to check if a consumer is available on a particular shard.  There is no need to do a full GET for the consumer for this check.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Create**
> ConsumerDto Create(ctx, consumerDto, optional)
create

Creates a Consumer. NOTE: Opening this method up to everyone, as we have nothing we can reliably verify in the method signature. Instead we have to figure out what owner this consumer is destined for (due to backward compatability with existing clients which do not specify an owner during registration), and then check the access to the specified owner in the method itself.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerDto** | [**ConsumerDto**](ConsumerDto.md)|  | 
 **optional** | ***CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**|  | 
 **owner** | **optional.String**|  | 
 **activationKeys** | **optional.String**|  | 
 **identityCertCreation** | **optional.Bool**|  | [default to true]

### Return type

[**ConsumerDto**](ConsumerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConsumer**
> DeleteConsumer(ctx, consumerUuid)
deleteConsumer

Removes a Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContentOverrides**
> []ContentOverride DeleteContentOverrides(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteContentOverridesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteContentOverridesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentOverride** | [**optional.Interface of []ContentOverride**](array.md)|  | 

### Return type

[**[]ContentOverride**](ContentOverride.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGuest**
> DeleteGuest(ctx, consumerUuid, guestId, optional)
deleteGuest

Removes the Guest from the Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**| consumer who owns or hosts the guest in question | 
  **guestId** | **string**|  | 
 **optional** | ***DeleteGuestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteGuestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unregister** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadExistingExport**
> File DownloadExistingExport(ctx, consumerUuid, exportId)
Async Consumer Export (manifest) Download

Downloads an asynchronously generated consumer export file (manifest).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
  **exportId** | **string**|  | 

### Return type

[**File**](File.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DryBind**
> []PoolQuantityDto DryBind(ctx, consumerUuid, optional)
dryBind

Retrieves a list of Pools and quantities that would be the result of an auto-bind. This is a dry run of an autobind. It allows the client to see what would be the result of an autobind without executing it. It can only do this for the prevously established list of installed products for the consumer If a service level is included in the request, then that level will override the one stored on the consumer. If no service level is included then the existing one will be used. The Response has a list of PoolQuantity objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
 **optional** | ***DryBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DryBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceLevel** | **optional.String**|  | 

### Return type

[**[]PoolQuantityDto**](PoolQuantityDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportData**
> File ExportData(ctx, consumerUuid, optional)
Consumer Export (manifest)

Retrieves a Compressed File representation of a Consumer (manifest).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
 **optional** | ***ExportDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportDataOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnLabel** | **optional.String**|  | 
 **webappPrefix** | **optional.String**|  | 
 **apiUrl** | **optional.String**|  | 
 **ext** | [**optional.Interface of []string**](string.md)| Key/Value pairs to be passed to the extension adapter when generating a manifest | 

### Return type

[**File**](File.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportDataAsync**
> JobDetail ExportDataAsync(ctx, consumerUuid, optional)
Async Consumer Export (manifest)

Initiates an async generation of a Compressed File representation of a Consumer (manifest). The response will contain the id of the job from which its result data  will contain the href to download the generated file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**| The UUID of the target consumer | 
 **optional** | ***ExportDataAsyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportDataAsyncOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnLabel** | **optional.String**| The lable of the target CDN | 
 **webappPrefix** | **optional.String**| the URL pointing to the manifest&#39;s originating web application | 
 **apiUrl** | **optional.String**| the URL pointing to the manifest&#39;s originating candlepin API | 
 **ext** | [**optional.Interface of []string**](string.md)| Key/Value pairs to be passed to the extension adapter when generating a manifest | 

### Return type

[**JobDetail**](JobDetail.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComplianceStatus**
> ComplianceStatus GetComplianceStatus(ctx, consumerUuid, optional)
getComplianceStatus

Retireves the Compliance Status of a Consumer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
 **optional** | ***GetComplianceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetComplianceStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onDate** | **optional.String**| Date to get compliance information for, default is now. | 

### Return type

[**ComplianceStatus**](ComplianceStatus.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComplianceStatusList**
> map[string]ComplianceStatus GetComplianceStatusList(ctx, optional)
getComplianceStatusList

Retrieves a Compliance Status list for a list of Consumers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetComplianceStatusListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetComplianceStatusListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**map[string]ComplianceStatus**](ComplianceStatus.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumer**
> ConsumerDto GetConsumer(ctx, consumerUuid)
getConsumer

Retrieves a single Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**ConsumerDto**](ConsumerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumerAtomFeed**
> Feed GetConsumerAtomFeed(ctx, consumerUuid)
getConsumerAtomFeed

Retrieves and Event Atom Feed for a Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**Feed**](feed.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/atom+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumerEvents**
> []EventDto GetConsumerEvents(ctx, consumerUuid)
getConsumerEvents

Retrieves a list of Consumer Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**[]EventDto**](EventDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentAccessBody**
> string GetContentAccessBody(ctx, consumerUuid, optional)
getContentAccessBody

Retrieves the body of the Content Access Certificate for the Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
 **optional** | ***GetContentAccessBodyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContentAccessBodyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**|  | [default to Thu, 01 Jan 1970 00:00:00 GMT]

### Return type

**string**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentOverrideList**
> []ContentOverride GetContentOverrideList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ContentOverride**](ContentOverride.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementCertificateSerials**
> []Serials GetEntitlementCertificateSerials(ctx, consumerUuid)
getEntitlementCertificateSerials

Retrieves a list of Certiticate Serials Return the client certificate metadata a for the given consumer. This is a small subset of data clients can use to determine which certificates they need to update/fetch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**[]Serials**](serials.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntitlementCertificates**
> []CertificateDto GetEntitlementCertificates(ctx, consumerUuid, optional)
getEntitlementCertificates

Retrieves a list of Entitlement Certificates for the Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
 **optional** | ***GetEntitlementCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEntitlementCertificatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **optional.String**|  | 

### Return type

[**[]CertificateDto**](CertificateDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGuestId**
> GuestIdDto GetGuestId(ctx, consumerUuid, guestId)
getGuestId

Retrieves a single Guest By its consumer and the guest UUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
  **guestId** | **string**|  | 

### Return type

[**GuestIdDto**](GuestIdDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGuestIds**
> CandlepinQueryGuestIdDto GetGuestIds(ctx, consumerUuid)
getGuestIds

Retrieves the List of a Consumer's Guests

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**CandlepinQueryGuestIdDto**](CandlepinQueryGuestIdDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGuests**
> []ConsumerDto GetGuests(ctx, consumerUuid)
getGuests

Retrieves a list of Guest Consumers of a Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**[]ConsumerDto**](ConsumerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHost**
> ConsumerDto GetHost(ctx, consumerUuid)
getHost

Retrieves a Host Consumer of a Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**ConsumerDto**](ConsumerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOwner**
> OwnerDto GetOwner(ctx, consumerUuid)
getOwner

Retrieves the Owner associated to a Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**OwnerDto**](OwnerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelease**
> Release GetRelease(ctx, consumerUuid)
getRelease

Retrieves the Release of a Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**Release**](Release.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemPurposeComplianceStatus**
> SystemPurposeComplianceStatusDto GetSystemPurposeComplianceStatus(ctx, consumerUuid, optional)
getSystemPurposeComplianceStatus

Retrieves the System Purpose Compliance Status of a Consumer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
 **optional** | ***GetSystemPurposeComplianceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSystemPurposeComplianceStatusOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onDate** | **optional.String**| Date to get compliance information for, default is now. | 

### Return type

[**SystemPurposeComplianceStatusDto**](SystemPurposeComplianceStatusDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []Consumer List(ctx, optional)
list

Retrieves a list of the Consumers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**|  | 
 **type_** | [**optional.Interface of []string**](string.md)|  | 
 **owner** | **optional.String**|  | 
 **uuid** | [**optional.Interface of []string**](string.md)|  | 
 **hypervisorId** | [**optional.Interface of []string**](string.md)|  | 
 **fact** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]Consumer**](Consumer.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEntitlements**
> []EntitlementDto ListEntitlements(ctx, consumerUuid, optional)
listEntitlements

Retrives a list of Entitlements

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
 **optional** | ***ListEntitlementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListEntitlementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | **optional.String**|  | 
 **regen** | **optional.Bool**|  | [default to true]
 **matches** | **optional.String**|  | 
 **attribute** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]EntitlementDto**](EntitlementDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateEntitlementCertificates**
> RegenerateEntitlementCertificates(ctx, consumerUuid, optional)
regenerateEntitlementCertificates

Regenerates the Entitlement Certificates for a Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
 **optional** | ***RegenerateEntitlementCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegenerateEntitlementCertificatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlement** | **optional.String**|  | 
 **lazyRegen** | **optional.Bool**|  | [default to true]

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateIdentityCertificates**
> ConsumerDto RegenerateIdentityCertificates(ctx, consumerUuid)
regenerateIdentityCertificates

Retrieves a single Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**ConsumerDto**](ConsumerDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveDeletionRecord**
> RemoveDeletionRecord(ctx, consumerUuid)
removeDeletionRecord

Removes the Deletion Record for a Consumer Allowed for a superadmin. The main use case for this would be if a user accidently deleted a non-RHEL hypervisor, causing it to no longer be auto-detected via virt-who.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unbind**
> Unbind(ctx, consumerUuid, dbid)
unbind

Removes an Entitlement from a Consumer By the Entitlement ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
  **dbid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnbindAll**
> DeleteResult UnbindAll(ctx, consumerUuid)
unbindAll

Unbinds all Entitlements for a Consumer Result contains the total number of entitlements unbound.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 

### Return type

[**DeleteResult**](DeleteResult.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnbindByPool**
> UnbindByPool(ctx, consumerUuid, poolId)
unbindByPool

Removes all Entitlements from a Consumer. By Pool Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnbindBySerial**
> UnbindBySerial(ctx, consumerUuid, serial)
unbindBySerial

Removes an Entitlement from a Consumer By the Certificate Serial

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
  **serial** | **int64**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConsumer**
> UpdateConsumer(ctx, consumerUuid, consumerDto)
updateConsumer

Updates a Consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
  **consumerDto** | [**ConsumerDto**](ConsumerDto.md)|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGuest**
> UpdateGuest(ctx, consumerUuid, guestId, guestIdDto)
updateGuest

Updates a single Guest on a Consumer. Allows virt-who to avoid uploading an entire list of guests

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**| consumer who owns or hosts the guest in question | 
  **guestId** | **string**| guest virtual uuid | 
  **guestIdDto** | [**GuestIdDto**](GuestIdDto.md)| updated guest data to use | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGuests**
> UpdateGuests(ctx, consumerUuid, guestIdDTO)
updateGuests

Updates the List of Guests on a Consumer This method should work just like updating the consumer, except that it only updates GuestIds.  Eventually we should move All the logic here, and depricate updating guests through the consumer update.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerUuid** | **string**|  | 
  **guestIdDTO** | [**[]GuestIdDto**](array.md)|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


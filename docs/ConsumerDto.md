# ConsumerDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**time.Time**](time.Time.md) |  | [optional] 
**Updated** | [**time.Time**](time.Time.md) |  | [optional] 
**Id** | **string** |  | [optional] 
**Uuid** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Username** | **string** |  | [optional] 
**EntitlementStatus** | **string** |  | [optional] 
**ServiceLevel** | **string** |  | [optional] 
**Role** | **string** |  | [optional] 
**Usage** | **string** |  | [optional] 
**AddOns** | **[]string** |  | [optional] 
**SystemPurposeStatus** | **string** |  | [optional] 
**Owner** | [**NestedOwnerDto**](NestedOwnerDTO.md) |  | [optional] 
**Environment** | [**EnvironmentDto**](EnvironmentDTO.md) |  | [optional] 
**EntitlementCount** | **int64** |  | [optional] 
**Facts** | **map[string]string** |  | [optional] 
**LastCheckin** | [**time.Time**](time.Time.md) |  | [optional] 
**InstalledProducts** | [**[]ConsumerInstalledProductDto**](ConsumerInstalledProductDTO.md) |  | [optional] 
**CanActivate** | **bool** |  | [optional] [default to false]
**Capabilities** | [**[]CapabilityDto**](CapabilityDTO.md) |  | [optional] 
**HypervisorId** | [**HypervisorIdDto**](HypervisorIdDTO.md) |  | [optional] 
**ContentTags** | **[]string** |  | [optional] 
**Autoheal** | **bool** |  | [optional] [default to false]
**RecipientOwnerKey** | **string** |  | [optional] 
**Annotations** | **string** |  | [optional] 
**ContentAccessMode** | **string** |  | [optional] 
**Type** | [**ConsumerTypeDto**](ConsumerTypeDTO.md) |  | [optional] 
**IdCert** | [**CertificateDto**](CertificateDTO.md) |  | [optional] 
**GuestIds** | [**[]GuestIdDto**](GuestIdDTO.md) |  | [optional] 
**Href** | **string** |  | [optional] 
**ReleaseVer** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



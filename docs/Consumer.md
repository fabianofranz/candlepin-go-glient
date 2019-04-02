# Consumer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**time.Time**](time.Time.md) |  | [optional] 
**Updated** | [**time.Time**](time.Time.md) |  | [optional] 
**Id** | **string** |  | 
**Uuid** | **string** |  | 
**Name** | **string** |  | 
**Username** | **string** |  | [optional] 
**EntitlementStatus** | **string** |  | [optional] 
**ServiceLevel** | **string** |  | [optional] 
**Role** | **string** |  | [optional] 
**Usage** | **string** |  | [optional] 
**AddOns** | **[]string** |  | [optional] 
**SystemPurposeStatus** | **string** |  | [optional] 
**ReleaseVer** | [**Release**](Release.md) |  | [optional] 
**IdCert** | [**IdentityCertificate**](IdentityCertificate.md) |  | [optional] 
**TypeId** | **string** |  | 
**OwnerId** | **string** |  | [optional] 
**EnvironmentId** | **string** |  | [optional] 
**EntitlementCount** | **int64** |  | 
**Facts** | **map[string]string** |  | [optional] 
**LastCheckin** | [**time.Time**](time.Time.md) |  | [optional] 
**InstalledProducts** | [**[]ConsumerInstalledProduct**](ConsumerInstalledProduct.md) |  | [optional] 
**CanActivate** | **bool** |  | [optional] [default to false]
**Capabilities** | [**[]Consumercapability**](consumercapability.md) |  | [optional] 
**HypervisorId** | [**HypervisorId**](HypervisorId.md) |  | [optional] 
**ContentTags** | **[]string** |  | [optional] 
**Autoheal** | **bool** |  | [optional] [default to false]
**ContentAccessMode** | **string** |  | [optional] 
**RecipientOwnerKey** | **string** |  | [optional] 
**Annotations** | **string** |  | [optional] 
**Dev** | **bool** |  | [optional] [default to false]
**Href** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



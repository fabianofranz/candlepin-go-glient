# Pool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**time.Time**](time.Time.md) |  | [optional] 
**Updated** | [**time.Time**](time.Time.md) |  | [optional] 
**Id** | **string** |  | [optional] 
**Type** | **string** |  | [optional] 
**Owner** | [**NestedOwnerDto**](NestedOwnerDTO.md) |  | [optional] 
**ActiveSubscription** | **bool** |  | [optional] [default to false]
**CreatedByShare** | **bool** |  | [optional] [default to false]
**HasSharedAncestor** | **bool** |  | [optional] [default to false]
**SourceEntitlement** | [**NestedEntitlementDto**](NestedEntitlementDTO.md) |  | [optional] 
**Quantity** | **int64** |  | [optional] 
**StartDate** | [**time.Time**](time.Time.md) |  | [optional] 
**EndDate** | [**time.Time**](time.Time.md) |  | [optional] 
**Attributes** | **map[string]string** |  | [optional] 
**RestrictedToUsername** | **string** |  | [optional] 
**ContractNumber** | **string** |  | [optional] 
**AccountNumber** | **string** |  | [optional] 
**OrderNumber** | **string** |  | [optional] 
**Consumed** | **int64** |  | [optional] 
**Exported** | **int64** |  | [optional] 
**Shared** | **int64** |  | [optional] 
**Branding** | [**[]BrandingDto**](BrandingDTO.md) |  | [optional] 
**CalculatedAttributes** | **map[string]string** |  | [optional] 
**UpstreamPoolId** | **string** |  | [optional] 
**UpstreamEntitlementId** | **string** |  | [optional] 
**UpstreamConsumerId** | **string** |  | [optional] 
**ProductName** | **string** |  | [optional] 
**ProductId** | **string** |  | [optional] 
**ProductAttributes** | **map[string]string** |  | [optional] 
**StackId** | **string** |  | [optional] 
**Stacked** | **bool** |  | [optional] [default to false]
**SourceStackId** | **string** |  | [optional] 
**DevelopmentPool** | **bool** |  | [optional] [default to false]
**DerivedProductAttributes** | **map[string]string** |  | [optional] 
**DerivedProductId** | **string** |  | [optional] 
**DerivedProductName** | **string** |  | [optional] 
**ProvidedProducts** | [**[]ProvidedProductDto**](ProvidedProductDTO.md) |  | [optional] 
**DerivedProvidedProducts** | [**[]ProvidedProductDto**](ProvidedProductDTO.md) |  | [optional] 
**SubscriptionSubKey** | **string** |  | [optional] 
**SubscriptionId** | **string** |  | [optional] 
**Href** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SystemPurposeComplianceStatusDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | [optional] 
**Compliant** | **bool** |  | [optional] [default to false]
**Date** | [**time.Time**](time.Time.md) |  | [optional] 
**NonCompliantRole** | **string** |  | [optional] 
**NonCompliantAddOns** | **[]string** |  | [optional] 
**NonCompliantSLA** | **string** |  | [optional] 
**NonCompliantUsage** | **string** |  | [optional] 
**CompliantRole** | [**map[string][]EntitlementDto**](array.md) |  | [optional] 
**CompliantAddOns** | [**map[string][]EntitlementDto**](array.md) |  | [optional] 
**CompliantSLA** | [**map[string][]EntitlementDto**](array.md) |  | [optional] 
**CompliantUsage** | [**map[string][]EntitlementDto**](array.md) |  | [optional] 
**NonPreferredSLA** | [**map[string][]EntitlementDto**](array.md) |  | [optional] 
**NonPreferredUsage** | [**map[string][]EntitlementDto**](array.md) |  | [optional] 
**Reasons** | **[]string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



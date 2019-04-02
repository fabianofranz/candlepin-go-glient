# ComplianceStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**time.Time**](time.Time.md) |  | [optional] 
**CompliantUntil** | [**time.Time**](time.Time.md) |  | [optional] 
**NonCompliantProducts** | **[]string** |  | [optional] 
**CompliantProducts** | [**map[string][]Entitlement**](array.md) |  | [optional] 
**PartiallyCompliantProducts** | [**map[string][]Entitlement**](array.md) |  | [optional] 
**PartialStacks** | [**map[string][]Entitlement**](array.md) |  | [optional] 
**ProductComplianceDateRanges** | [**map[string]DateRange**](DateRange.md) |  | [optional] 
**Reasons** | [**[]ComplianceReason**](ComplianceReason.md) |  | [optional] 
**Status** | **string** |  | [optional] 
**Compliant** | **bool** |  | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



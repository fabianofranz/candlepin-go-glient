# ActivationKeyDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**time.Time**](time.Time.md) |  | [optional] 
**Updated** | [**time.Time**](time.Time.md) |  | [optional] 
**Id** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Description** | **string** |  | [optional] 
**Owner** | [**NestedOwnerDto**](NestedOwnerDTO.md) |  | [optional] 
**ServiceLevel** | **string** |  | [optional] 
**AutoAttach** | **bool** |  | [optional] [default to false]
**Pools** | [**[]ActivationKeyPoolDto**](ActivationKeyPoolDTO.md) |  | [optional] 
**ContentOverrides** | [**[]ActivationKeyContentOverrideDto**](ActivationKeyContentOverrideDTO.md) |  | [optional] 
**Products** | **[]string** |  | [optional] 
**ReleaseVer** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# \RolesApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRolePermission**](RolesApi.md#AddRolePermission) | **Post** /roles/{role_id}/permissions | addRolePermission
[**AddUser**](RolesApi.md#AddUser) | **Post** /roles/{role_id}/users/{username} | addUser
[**CreateRole**](RolesApi.md#CreateRole) | **Post** /roles | createRole
[**DeleteRole**](RolesApi.md#DeleteRole) | **Delete** /roles/{role_id} | deleteRole
[**DeleteUser**](RolesApi.md#DeleteUser) | **Delete** /roles/{role_id}/users/{username} | deleteUser
[**GetRole**](RolesApi.md#GetRole) | **Get** /roles/{role_id} | getRole
[**GetRoles**](RolesApi.md#GetRoles) | **Get** /roles | getRoles
[**RemoveRolePermission**](RolesApi.md#RemoveRolePermission) | **Delete** /roles/{role_id}/permissions/{perm_id} | removeRolePermission
[**UpdateRole**](RolesApi.md#UpdateRole) | **Put** /roles/{role_id} | updateRole


# **AddRolePermission**
> Role AddRolePermission(ctx, roleId, permissionBlueprint)
addRolePermission

Adds a Permission to a Role. Returns the updated Role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 
  **permissionBlueprint** | [**PermissionBlueprint**](PermissionBlueprint.md)|  | 

### Return type

[**Role**](Role.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUser**
> Role AddUser(ctx, roleId, username)
addUser

Adds a User to a Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 
  **username** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRole**
> Role CreateRole(ctx, role)
createRole

Creates a Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **role** | [**Role**](Role.md)|  | 

### Return type

[**Role**](Role.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> DeleteRole(ctx, roleId)
deleteRole

Removes a Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> Role DeleteUser(ctx, roleId, username)
deleteUser

Removes a User from a Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 
  **username** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> Role GetRole(ctx, roleId)
getRole

Retrieves a single Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoles**
> []Role GetRoles(ctx, )
getRoles

Retrieves a list of Roles

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Role**](Role.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRolePermission**
> Role RemoveRolePermission(ctx, roleId, permId)
removeRolePermission

Removes a Permission from a Role. Returns the updated Role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 
  **permId** | **string**|  | 

### Return type

[**Role**](Role.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRole**
> Role UpdateRole(ctx, roleId, role)
updateRole

Updates a Role.  To avoid race conditions, we do not support updating the user or permission collections. Currently this call will only update the role name. See the specific nested POST/DELETE calls for modifying users and permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 
  **role** | [**Role**](Role.md)|  | 

### Return type

[**Role**](Role.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


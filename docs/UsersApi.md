# \UsersApi

All URIs are relative to *http://http:/candlepin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /users | createUser
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /users/{username} | deleteUser
[**GetUserInfo**](UsersApi.md#GetUserInfo) | **Get** /users/{username} | getUserInfo
[**GetUserRoles**](UsersApi.md#GetUserRoles) | **Get** /users/{username}/roles | getUserRoles
[**List**](UsersApi.md#List) | **Get** /users | list
[**ListUsersOwners**](UsersApi.md#ListUsersOwners) | **Get** /users/{username}/owners | listUsersOwners
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /users/{username} | updateUser


# **CreateUser**
> User CreateUser(ctx, userCreationRequest)
createUser

Creates a User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userCreationRequest** | [**UserCreationRequest**](UserCreationRequest.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, username)
deleteUser

Removes a User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserInfo**
> User GetUserInfo(ctx, username)
getUserInfo

Retrieves a single User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserRoles**
> []Role GetUserRoles(ctx, username)
getUserRoles

Retrieves a list of Roles by User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**|  | 

### Return type

[**[]Role**](Role.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []User List(ctx, )
list

Retrieves a list of Users

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsersOwners**
> IterableOwner ListUsersOwners(ctx, username)
listUsersOwners

Retrieve a list of owners the user can register systems to. Previously this represented owners the user was an admin for. Because the client uses this API call to list the owners a user can register to, when we introduced 'my systems' administrator, we have to change its meaning to listing the owners that can be registered to by default to maintain compatability with released clients.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**|  | 

### Return type

[**IterableOwner**](IterableOwner.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> User UpdateUser(ctx, username, user)
updateUser

Updates a User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**|  | 
  **user** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


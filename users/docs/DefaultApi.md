# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/users*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GroupsBasicGet**](DefaultApi.md#V1GroupsBasicGet) | **Get** /v1/groups/basic/ | List groups info without details
[**V1GroupsGet**](DefaultApi.md#V1GroupsGet) | **Get** /v1/groups/ | List groups with details
[**V1GroupsGroupIdDelete**](DefaultApi.md#V1GroupsGroupIdDelete) | **Delete** /v1/groups/{group_id}/ | Delete a particular group by id
[**V1GroupsGroupIdGet**](DefaultApi.md#V1GroupsGroupIdGet) | **Get** /v1/groups/{group_id}/ | Returns a particular group by id
[**V1GroupsGroupIdLogoDelete**](DefaultApi.md#V1GroupsGroupIdLogoDelete) | **Delete** /v1/groups/{group_id}/logo/ | Delete group logo image.
[**V1GroupsGroupIdLogoPost**](DefaultApi.md#V1GroupsGroupIdLogoPost) | **Post** /v1/groups/{group_id}/logo/ | Upload group logo image.
[**V1GroupsGroupIdPatch**](DefaultApi.md#V1GroupsGroupIdPatch) | **Patch** /v1/groups/{group_id}/ | Update group
[**V1GroupsGroupIdPut**](DefaultApi.md#V1GroupsGroupIdPut) | **Put** /v1/groups/{group_id}/ | Update group
[**V1GroupsGroupIdReindexPost**](DefaultApi.md#V1GroupsGroupIdReindexPost) | **Post** /v1/groups/{group_id}/reindex/ | Reindex a particular group by id
[**V1GroupsGroupIdUsersUserIdDelete**](DefaultApi.md#V1GroupsGroupIdUsersUserIdDelete) | **Delete** /v1/groups/{group_id}/users/{user_id}/ | Delete a user from group
[**V1GroupsGroupIdUsersUserIdPost**](DefaultApi.md#V1GroupsGroupIdUsersUserIdPost) | **Post** /v1/groups/{group_id}/users/{user_id}/ | Add user into a group
[**V1GroupsMappingsGet**](DefaultApi.md#V1GroupsMappingsGet) | **Get** /v1/groups/mappings/ | Get all group mappings
[**V1GroupsMappingsNameDelete**](DefaultApi.md#V1GroupsMappingsNameDelete) | **Delete** /v1/groups/mappings/{name}/ | Delete group mapping by name
[**V1GroupsMappingsNameGet**](DefaultApi.md#V1GroupsMappingsNameGet) | **Get** /v1/groups/mappings/{name}/ | Get a group mapping
[**V1GroupsMappingsPost**](DefaultApi.md#V1GroupsMappingsPost) | **Post** /v1/groups/mappings/ | Create a new group mapping
[**V1GroupsPost**](DefaultApi.md#V1GroupsPost) | **Post** /v1/groups/ | Create a new group
[**V1UsersBasicGet**](DefaultApi.md#V1UsersBasicGet) | **Get** /v1/users/basic/ | List of users without details
[**V1UsersCurrentGet**](DefaultApi.md#V1UsersCurrentGet) | **Get** /v1/users/current/ | Returns current user
[**V1UsersCurrentPatch**](DefaultApi.md#V1UsersCurrentPatch) | **Patch** /v1/users/current/ | Update user
[**V1UsersCurrentPhotoDelete**](DefaultApi.md#V1UsersCurrentPhotoDelete) | **Delete** /v1/users/current/photo/ | Delete current user photo image.
[**V1UsersCurrentPhotoPost**](DefaultApi.md#V1UsersCurrentPhotoPost) | **Post** /v1/users/current/photo/ | Upload current user photo image.
[**V1UsersCurrentPut**](DefaultApi.md#V1UsersCurrentPut) | **Put** /v1/users/current/ | Update user
[**V1UsersCurrentRolesGet**](DefaultApi.md#V1UsersCurrentRolesGet) | **Get** /v1/users/current/roles/ | Returns current user roles
[**V1UsersGet**](DefaultApi.md#V1UsersGet) | **Get** /v1/users/ | List of users with details
[**V1UsersLoginPost**](DefaultApi.md#V1UsersLoginPost) | **Post** /v1/users/login/ | Login a user
[**V1UsersPost**](DefaultApi.md#V1UsersPost) | **Post** /v1/users/ | Create a new user
[**V1UsersUserIdDelete**](DefaultApi.md#V1UsersUserIdDelete) | **Delete** /v1/users/{user_id}/ | Delete a particular user by id
[**V1UsersUserIdGet**](DefaultApi.md#V1UsersUserIdGet) | **Get** /v1/users/{user_id}/ | Returns a particular user by id
[**V1UsersUserIdPatch**](DefaultApi.md#V1UsersUserIdPatch) | **Patch** /v1/users/{user_id}/ | Update user
[**V1UsersUserIdPhotoDelete**](DefaultApi.md#V1UsersUserIdPhotoDelete) | **Delete** /v1/users/{user_id}/photo/ | Delete a photo image of a specified user.
[**V1UsersUserIdPhotoPost**](DefaultApi.md#V1UsersUserIdPhotoPost) | **Post** /v1/users/{user_id}/photo/ | Upload user photo image.
[**V1UsersUserIdPut**](DefaultApi.md#V1UsersUserIdPut) | **Put** /v1/users/{user_id}/ | Update user
[**V1UsersUserIdReindexPost**](DefaultApi.md#V1UsersUserIdReindexPost) | **Post** /v1/users/{user_id}/reindex/ | Reindex a particular user by id
[**V1UsersUserIdRolesGet**](DefaultApi.md#V1UsersUserIdRolesGet) | **Get** /v1/users/{user_id}/roles/ | Returns user roles by user_id
[**V1UsersUserIdRolesRoleGet**](DefaultApi.md#V1UsersUserIdRolesRoleGet) | **Get** /v1/users/{user_id}/roles/{role}/ | Returns user roles by user_id
[**V1UsersUserIdSamlDelete**](DefaultApi.md#V1UsersUserIdSamlDelete) | **Delete** /v1/users/{user_id}/saml/ | Remove a user&#39;s SAML IdP setting
[**V1UsersUserIdSamlPut**](DefaultApi.md#V1UsersUserIdSamlPut) | **Put** /v1/users/{user_id}/saml/ | Update a user&#39;s SAML IdP settings


# **V1GroupsBasicGet**
> GroupsSchema V1GroupsBasicGet(ctx, appID, authToken, optional)
List groups info without details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1GroupsBasicGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1GroupsBasicGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **alias** | **optional.String**| Filter by alias | 
 **description** | **optional.String**| Filter by descripton | 
 **name** | **optional.String**| Filter by name | 
 **roles** | **optional.String**| Filter by roles | 
 **dateCreated** | **optional.String**| Filter by date_created | 
 **dateModified** | **optional.String**| Filter by date_modified | 
 **query** | **optional.String**| Filter by any of first_name, last_name and email with wildcard support | 
 **ids** | **optional.String**| Filter list of id:s (comma separated) | 

### Return type

[**GroupsSchema**](GroupsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGet**
> GroupsSchema V1GroupsGet(ctx, appID, authToken, optional)
List groups with details

 Required roles:  - can_read_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1GroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1GroupsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **alias** | **optional.String**| Filter by alias | 
 **description** | **optional.String**| Filter by descripton | 
 **name** | **optional.String**| Filter by name | 
 **roles** | **optional.String**| Filter by roles | 
 **dateCreated** | **optional.String**| Filter by date_created | 
 **dateModified** | **optional.String**| Filter by date_modified | 
 **query** | **optional.String**| Filter by any of field with wildcard support | 
 **ids** | **optional.String**| Filter list of id:s (comma separated) | 

### Return type

[**GroupsSchema**](GroupsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdDelete**
> V1GroupsGroupIdDelete(ctx, appID, authToken, groupId)
Delete a particular group by id

 Required roles:  - can_delete_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdGet**
> GroupSchema V1GroupsGroupIdGet(ctx, appID, authToken, groupId)
Returns a particular group by id

 Required roles:  - can_read_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdLogoDelete**
> V1GroupsGroupIdLogoDelete(ctx, appID, authToken, groupId)
Delete group logo image.

 Required roles:  - can_write_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdLogoPost**
> InlineResponse200 V1GroupsGroupIdLogoPost(ctx, appID, authToken, groupId, logo)
Upload group logo image.

 Required roles:  - can_write_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **logo** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdPatch**
> GroupSchema V1GroupsGroupIdPatch(ctx, appID, authToken, groupId, body)
Update group

 Required roles:  - can_write_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **body** | [**GroupSchema**](GroupSchema.md)|  | 

### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdPut**
> GroupSchema V1GroupsGroupIdPut(ctx, appID, authToken, groupId, body)
Update group

 Required roles:  - can_write_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **body** | [**GroupSchema**](GroupSchema.md)|  | 

### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdReindexPost**
> UserSchema V1GroupsGroupIdReindexPost(ctx, appID, authToken, groupId)
Reindex a particular group by id

 Required roles:  - can_reindex_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdUsersUserIdDelete**
> UserSchema V1GroupsGroupIdUsersUserIdDelete(ctx, appID, authToken, groupId, userId)
Delete a user from group

 Required roles:  - can_delete_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdUsersUserIdPost**
> UserSchema V1GroupsGroupIdUsersUserIdPost(ctx, appID, authToken, groupId, userId)
Add user into a group

 Required roles:  - can_write_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsMappingsGet**
> GroupMappingsSchema V1GroupsMappingsGet(ctx, appID, authToken, optional)
Get all group mappings

 Required roles:  - can_read_group_mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1GroupsMappingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1GroupsMappingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **lastId** | **optional.String**|  | 

### Return type

[**GroupMappingsSchema**](GroupMappingsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsMappingsNameDelete**
> V1GroupsMappingsNameDelete(ctx, appID, authToken, name)
Delete group mapping by name

 Required roles:  - can_delete_group_mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsMappingsNameGet**
> GroupMappingSchema V1GroupsMappingsNameGet(ctx, appID, authToken, name)
Get a group mapping

 Required roles:  - can_read_group_mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**GroupMappingSchema**](GroupMappingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsMappingsPost**
> GroupMappingSchema V1GroupsMappingsPost(ctx, appID, authToken, body)
Create a new group mapping

 Required roles:  - can_write_group_mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**GroupMappingSchema**](GroupMappingSchema.md)|  | 

### Return type

[**GroupMappingSchema**](GroupMappingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsPost**
> GroupSchema V1GroupsPost(ctx, appID, authToken, body)
Create a new group

 Required roles:  - can_write_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**GroupSchema**](GroupSchema.md)|  | 

### Return type

[**GroupSchema**](GroupSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersBasicGet**
> UsersSchema V1UsersBasicGet(ctx, appID, authToken, optional)
List of users without details



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1UsersBasicGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1UsersBasicGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **dateCreated** | **optional.String**| Filter by date_created | 
 **dateModified** | **optional.String**| Filter by date_modified | 
 **email** | **optional.String**| Filter by email | 
 **firstName** | **optional.String**| Filter by first_name | 
 **lastName** | **optional.String**| Filter by last_name | 
 **groups** | **optional.String**| Filter by groups | 
 **hideEmail** | **optional.String**| Filter by hide_email | 
 **hidePhone** | **optional.String**| Filter by hide_phone | 
 **isAdmin** | **optional.String**| Filter by is_admin | 
 **passwordChanged** | **optional.String**| Filter by password_changed | 
 **phone** | **optional.String**| Filter by phone | 
 **photo** | **optional.String**| Filter by photo | 
 **status** | **optional.String**| Filter by status | 
 **query** | **optional.String**| Filter by any of first_name, last_name and email with wildcard support | 
 **ids** | **optional.String**| Filter list of id:s (comma separated) | 

### Return type

[**UsersSchema**](UsersSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersCurrentGet**
> UserSchema V1UsersCurrentGet(ctx, appID, authToken)
Returns current user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersCurrentPatch**
> UserSchema V1UsersCurrentPatch(ctx, appID, authToken, body)
Update user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**UserSchema**](UserSchema.md)|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersCurrentPhotoDelete**
> V1UsersCurrentPhotoDelete(ctx, appID, authToken)
Delete current user photo image.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersCurrentPhotoPost**
> InlineResponse2001 V1UsersCurrentPhotoPost(ctx, appID, authToken, photo)
Upload current user photo image.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **photo** | **string**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersCurrentPut**
> UserSchema V1UsersCurrentPut(ctx, appID, authToken, body)
Update user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**UserSchema**](UserSchema.md)|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersCurrentRolesGet**
> UserRolesSchema V1UsersCurrentRolesGet(ctx, appID, authToken)
Returns current user roles



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**UserRolesSchema**](UserRolesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersGet**
> UsersSchema V1UsersGet(ctx, appID, authToken, optional)
List of users with details

 Required roles:  - can_read_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1UsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1UsersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **dateCreated** | **optional.String**| Filter by date_created | 
 **dateModified** | **optional.String**| Filter by date_modified | 
 **email** | **optional.String**| Filter by email | 
 **firstName** | **optional.String**| Filter by first_name | 
 **lastName** | **optional.String**| Filter by last_name | 
 **groups** | **optional.String**| Filter by groups | 
 **hideEmail** | **optional.String**| Filter by hide_email | 
 **hidePhone** | **optional.String**| Filter by hide_phone | 
 **isAdmin** | **optional.String**| Filter by is_admin | 
 **passwordChanged** | **optional.String**| Filter by password_changed | 
 **phone** | **optional.String**| Filter by phone | 
 **photo** | **optional.String**| Filter by photo | 
 **status** | **optional.String**| Filter by status | 
 **query** | **optional.String**| Filter by any of first_name, last_name and email with wildcard support | 
 **ids** | **optional.String**| Filter list of id:s (comma separated) | 

### Return type

[**UsersSchema**](UsersSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersLoginPost**
> UserSchema V1UsersLoginPost(ctx, appID, authToken, body)
Login a user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**UserLoginSchema**](UserLoginSchema.md)|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersPost**
> UserSchema V1UsersPost(ctx, appID, authToken, body)
Create a new user

 Required roles:  - can_create_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**UserSchema**](UserSchema.md)|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdDelete**
> V1UsersUserIdDelete(ctx, appID, authToken, userId)
Delete a particular user by id

 Required roles:  - can_delete_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdGet**
> UserSchema V1UsersUserIdGet(ctx, appID, authToken, userId)
Returns a particular user by id

 Required roles:  - can_read_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdPatch**
> UserSchema V1UsersUserIdPatch(ctx, appID, authToken, userId, body)
Update user

 Required roles:  - can_write_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **body** | [**UserSchema**](UserSchema.md)|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdPhotoDelete**
> V1UsersUserIdPhotoDelete(ctx, appID, authToken, userId)
Delete a photo image of a specified user.

 Required roles:  - can_write_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdPhotoPost**
> InlineResponse2001 V1UsersUserIdPhotoPost(ctx, appID, authToken, userId, photo)
Upload user photo image.

 Required roles:  - can_write_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **photo** | **string**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdPut**
> UserSchema V1UsersUserIdPut(ctx, appID, authToken, userId, body)
Update user

 Required roles:  - can_write_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **body** | [**UserSchema**](UserSchema.md)|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdReindexPost**
> UserSchema V1UsersUserIdReindexPost(ctx, appID, authToken, userId)
Reindex a particular user by id

 Required roles:  - can_reindex_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdRolesGet**
> UserRolesSchema V1UsersUserIdRolesGet(ctx, appID, authToken, userId)
Returns user roles by user_id

 Required roles:  - can_read_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**UserRolesSchema**](UserRolesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdRolesRoleGet**
> V1UsersUserIdRolesRoleGet(ctx, appID, authToken, userId, role)
Returns user roles by user_id

 Required roles:  - can_read_users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **role** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdSamlDelete**
> UserSchema V1UsersUserIdSamlDelete(ctx, appID, authToken, userId)
Remove a user's SAML IdP setting



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdSamlPut**
> UserSchema V1UsersUserIdSamlPut(ctx, appID, authToken, userId, body)
Update a user's SAML IdP settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **body** | [**UserSamlIdpUpdateSchema**](UserSamlIdpUpdateSchema.md)|  | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


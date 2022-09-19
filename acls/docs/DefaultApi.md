# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/acls*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AclObjectTypeContentDelete**](DefaultApi.md#V1AclObjectTypeContentDelete) | **Delete** /v1/acl/{object_type}/content/ | Delete acls for content of multiple objects
[**V1AclObjectTypeContentPut**](DefaultApi.md#V1AclObjectTypeContentPut) | **Put** /v1/acl/{object_type}/content/ | Create a new acl for content of multiple objects
[**V1AclObjectTypeDelete**](DefaultApi.md#V1AclObjectTypeDelete) | **Delete** /v1/acl/{object_type}/ | Delete acls for multiple objects
[**V1AclObjectTypeObjectKeyGet**](DefaultApi.md#V1AclObjectTypeObjectKeyGet) | **Get** /v1/acl/{object_type}/{object_key}/ | List of object permissions
[**V1AclObjectTypeObjectKeyPermissionGet**](DefaultApi.md#V1AclObjectTypeObjectKeyPermissionGet) | **Get** /v1/acl/{object_type}/{object_key}/{permission}/ | Check if particular object has required permission
[**V1AclObjectTypeObjectKeyPermissionsGet**](DefaultApi.md#V1AclObjectTypeObjectKeyPermissionsGet) | **Get** /v1/acl/{object_type}/{object_key}/permissions/ | List of permissions for the user
[**V1AclObjectTypePermissionPost**](DefaultApi.md#V1AclObjectTypePermissionPost) | **Post** /v1/acl/{object_type}/{permission}/ | Check if objects have required permission
[**V1AclObjectTypePut**](DefaultApi.md#V1AclObjectTypePut) | **Put** /v1/acl/{object_type}/ | Create a new acl for multiple objects
[**V1AclPost**](DefaultApi.md#V1AclPost) | **Post** /v1/acl/ | Check if objects have required permission
[**V1AclTemplatesGet**](DefaultApi.md#V1AclTemplatesGet) | **Get** /v1/acl/templates/ | Retreive all acl templates
[**V1AclTemplatesPost**](DefaultApi.md#V1AclTemplatesPost) | **Post** /v1/acl/templates/ | Create an acl template
[**V1AclTemplatesTemplateIdDelete**](DefaultApi.md#V1AclTemplatesTemplateIdDelete) | **Delete** /v1/acl/templates/{template_id}/ | Remove an acl template
[**V1AclTemplatesTemplateIdGet**](DefaultApi.md#V1AclTemplatesTemplateIdGet) | **Get** /v1/acl/templates/{template_id}/ | Retreive an acl template
[**V1AclTemplatesTemplateIdObjectTypeObjectKeyPost**](DefaultApi.md#V1AclTemplatesTemplateIdObjectTypeObjectKeyPost) | **Post** /v1/acl/templates/{template_id}/{object_type}/{object_key}/ | Apply template permissions to an object
[**V1AclTemplatesTemplateIdPatch**](DefaultApi.md#V1AclTemplatesTemplateIdPatch) | **Patch** /v1/acl/templates/{template_id}/ | Update an acl template
[**V1AclTemplatesTemplateIdPut**](DefaultApi.md#V1AclTemplatesTemplateIdPut) | **Put** /v1/acl/templates/{template_id}/ | Update an acl template
[**V1GroupsGroupIdAclObjectTypeObjectKeyDelete**](DefaultApi.md#V1GroupsGroupIdAclObjectTypeObjectKeyDelete) | **Delete** /v1/groups/{group_id}/acl/{object_type}/{object_key}/ | Delete a particular acl by id for an object
[**V1GroupsGroupIdAclObjectTypeObjectKeyGet**](DefaultApi.md#V1GroupsGroupIdAclObjectTypeObjectKeyGet) | **Get** /v1/groups/{group_id}/acl/{object_type}/{object_key}/ | List of groups permissions for an object
[**V1GroupsGroupIdAclObjectTypeObjectKeyPermissionGet**](DefaultApi.md#V1GroupsGroupIdAclObjectTypeObjectKeyPermissionGet) | **Get** /v1/groups/{group_id}/acl/{object_type}/{object_key}/{permission}/ | Check if group has particular permission for an object
[**V1GroupsGroupIdAclObjectTypeObjectKeyPut**](DefaultApi.md#V1GroupsGroupIdAclObjectTypeObjectKeyPut) | **Put** /v1/groups/{group_id}/acl/{object_type}/{object_key}/ | Update or create group acl for an object
[**V1SharesObjectTypeObjectKeyGet**](DefaultApi.md#V1SharesObjectTypeObjectKeyGet) | **Get** /v1/shares/{object_type}/{object_key}/ | List of share acls
[**V1SharesShareIdAclObjectTypeObjectKeyDelete**](DefaultApi.md#V1SharesShareIdAclObjectTypeObjectKeyDelete) | **Delete** /v1/shares/{share_id}/acl/{object_type}/{object_key}/ | Delete a share acl for an object
[**V1SharesShareIdAclObjectTypeObjectKeyGet**](DefaultApi.md#V1SharesShareIdAclObjectTypeObjectKeyGet) | **Get** /v1/shares/{share_id}/acl/{object_type}/{object_key}/ | List of share permissions for an object
[**V1SharesShareIdAclObjectTypeObjectKeyPermissionGet**](DefaultApi.md#V1SharesShareIdAclObjectTypeObjectKeyPermissionGet) | **Get** /v1/shares/{share_id}/acl/{object_type}/{object_key}/{permission}/ | Returns a share acl for an object
[**V1SharesShareIdAclObjectTypeObjectKeyPost**](DefaultApi.md#V1SharesShareIdAclObjectTypeObjectKeyPost) | **Post** /v1/shares/{share_id}/acl/{object_type}/{object_key}/ | Create a new share acl for an object
[**V1SharesShareIdAclObjectTypeObjectKeyPut**](DefaultApi.md#V1SharesShareIdAclObjectTypeObjectKeyPut) | **Put** /v1/shares/{share_id}/acl/{object_type}/{object_key}/ | Update share acl for an object
[**V1SharesShareIdAclObjectTypePut**](DefaultApi.md#V1SharesShareIdAclObjectTypePut) | **Put** /v1/shares/{share_id}/acl/{object_type}/ | Create a new acl for multiple share objects
[**V1UsersUserIdAclObjectTypeObjectKeyDelete**](DefaultApi.md#V1UsersUserIdAclObjectTypeObjectKeyDelete) | **Delete** /v1/users/{user_id}/acl/{object_type}/{object_key}/ | Delete a user acl for an object
[**V1UsersUserIdAclObjectTypeObjectKeyGet**](DefaultApi.md#V1UsersUserIdAclObjectTypeObjectKeyGet) | **Get** /v1/users/{user_id}/acl/{object_type}/{object_key}/ | List of user permissions for an object
[**V1UsersUserIdAclObjectTypeObjectKeyPermissionGet**](DefaultApi.md#V1UsersUserIdAclObjectTypeObjectKeyPermissionGet) | **Get** /v1/users/{user_id}/acl/{object_type}/{object_key}/{permission}/ | Returns a user acl for an object
[**V1UsersUserIdAclObjectTypeObjectKeyPut**](DefaultApi.md#V1UsersUserIdAclObjectTypeObjectKeyPut) | **Put** /v1/users/{user_id}/acl/{object_type}/{object_key}/ | Update or create user acl for an object


# **V1AclObjectTypeContentDelete**
> V1AclObjectTypeContentDelete(ctx, appID, authToken, objectType, body)
Delete acls for content of multiple objects

 Required roles:  - can_delete_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**DeleteBulkAcLsSchema**](DeleteBulkAcLsSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclObjectTypeContentPut**
> V1AclObjectTypeContentPut(ctx, appID, authToken, objectType, body)
Create a new acl for content of multiple objects

 Required roles:  - can_write_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**CreateBulkAcLsSchema**](CreateBulkAcLsSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclObjectTypeDelete**
> V1AclObjectTypeDelete(ctx, appID, authToken, objectType, body)
Delete acls for multiple objects

 Required roles:  - can_delete_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**DeleteAcLsSchema**](DeleteAcLsSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclObjectTypeObjectKeyGet**
> AclSchema V1AclObjectTypeObjectKeyGet(ctx, appID, authToken, objectType, objectKey)
List of object permissions

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

[**AclSchema**](ACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclObjectTypeObjectKeyPermissionGet**
> V1AclObjectTypeObjectKeyPermissionGet(ctx, appID, authToken, objectType, objectKey, permission)
Check if particular object has required permission

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 
  **permission** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclObjectTypeObjectKeyPermissionsGet**
> CombinedPermissionsSchema V1AclObjectTypeObjectKeyPermissionsGet(ctx, appID, authToken, objectType, objectKey)
List of permissions for the user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

[**CombinedPermissionsSchema**](CombinedPermissionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclObjectTypePermissionPost**
> BulkAclSchema V1AclObjectTypePermissionPost(ctx, appID, authToken, objectType, permission)
Check if objects have required permission

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **permission** | **string**|  | 

### Return type

[**BulkAclSchema**](BulkACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclObjectTypePut**
> CreateAcLsResultSchema V1AclObjectTypePut(ctx, appID, authToken, objectType, body)
Create a new acl for multiple objects

 Required roles:  - can_write_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**CreateAcLsSchema**](CreateAcLsSchema.md)|  | 

### Return type

[**CreateAcLsResultSchema**](CreateACLsResultSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclPost**
> BulkAclSchema V1AclPost(ctx, appID, authToken)
Check if objects have required permission

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**BulkAclSchema**](BulkACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclTemplatesGet**
> AclTemplatesSchema V1AclTemplatesGet(ctx, appID, authToken)
Retreive all acl templates

 Required roles:  - can_read_acl_templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**AclTemplatesSchema**](ACLTemplatesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclTemplatesPost**
> AclTemplateSchema V1AclTemplatesPost(ctx, appID, authToken, body)
Create an acl template

 Required roles:  - can_write_acl_templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**AclTemplateSchema**](AclTemplateSchema.md)|  | 

### Return type

[**AclTemplateSchema**](ACLTemplateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclTemplatesTemplateIdDelete**
> V1AclTemplatesTemplateIdDelete(ctx, appID, authToken, templateId)
Remove an acl template

 Required roles:  - can_delete_acl_templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **templateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclTemplatesTemplateIdGet**
> AclTemplateSchema V1AclTemplatesTemplateIdGet(ctx, appID, authToken, templateId)
Retreive an acl template

 Required roles:  - can_read_acl_templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **templateId** | **string**|  | 

### Return type

[**AclTemplateSchema**](ACLTemplateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclTemplatesTemplateIdObjectTypeObjectKeyPost**
> V1AclTemplatesTemplateIdObjectTypeObjectKeyPost(ctx, appID, authToken, templateId, objectType, objectKey)
Apply template permissions to an object



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **templateId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclTemplatesTemplateIdPatch**
> AclTemplateSchema V1AclTemplatesTemplateIdPatch(ctx, appID, authToken, templateId, body)
Update an acl template

 Required roles:  - can_write_acl_templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **templateId** | **string**|  | 
  **body** | [**AclTemplateSchema**](AclTemplateSchema.md)|  | 

### Return type

[**AclTemplateSchema**](ACLTemplateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AclTemplatesTemplateIdPut**
> AclTemplateSchema V1AclTemplatesTemplateIdPut(ctx, appID, authToken, templateId, body)
Update an acl template

 Required roles:  - can_write_acl_templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **templateId** | **string**|  | 
  **body** | [**AclTemplateSchema**](AclTemplateSchema.md)|  | 

### Return type

[**AclTemplateSchema**](ACLTemplateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdAclObjectTypeObjectKeyDelete**
> V1GroupsGroupIdAclObjectTypeObjectKeyDelete(ctx, appID, authToken, groupId, objectType, objectKey)
Delete a particular acl by id for an object

 Required roles:  - can_delete_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdAclObjectTypeObjectKeyGet**
> GroupAclSchema V1GroupsGroupIdAclObjectTypeObjectKeyGet(ctx, appID, authToken, groupId, objectType, objectKey)
List of groups permissions for an object

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

[**GroupAclSchema**](GroupACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdAclObjectTypeObjectKeyPermissionGet**
> V1GroupsGroupIdAclObjectTypeObjectKeyPermissionGet(ctx, appID, authToken, groupId, objectType, objectKey, permission)
Check if group has particular permission for an object

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 
  **permission** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupsGroupIdAclObjectTypeObjectKeyPut**
> GroupAclSchema V1GroupsGroupIdAclObjectTypeObjectKeyPut(ctx, appID, authToken, groupId, objectType, objectKey, body)
Update or create group acl for an object

 Required roles:  - can_write_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 
  **body** | [**GroupAclSchema**](GroupAclSchema.md)|  | 

### Return type

[**GroupAclSchema**](GroupACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesObjectTypeObjectKeyGet**
> SharesAclSchema V1SharesObjectTypeObjectKeyGet(ctx, appID, authToken, objectType, objectKey)
List of share acls

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

[**SharesAclSchema**](SharesACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesShareIdAclObjectTypeObjectKeyDelete**
> V1SharesShareIdAclObjectTypeObjectKeyDelete(ctx, appID, authToken, shareId, objectType, objectKey)
Delete a share acl for an object

 Required roles:  - can_delete_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **shareId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesShareIdAclObjectTypeObjectKeyGet**
> ShareAclSchema V1SharesShareIdAclObjectTypeObjectKeyGet(ctx, appID, authToken, shareId, objectType, objectKey)
List of share permissions for an object

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **shareId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

[**ShareAclSchema**](ShareACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesShareIdAclObjectTypeObjectKeyPermissionGet**
> V1SharesShareIdAclObjectTypeObjectKeyPermissionGet(ctx, appID, authToken, shareId, objectType, objectKey, permission)
Returns a share acl for an object

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **shareId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 
  **permission** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesShareIdAclObjectTypeObjectKeyPost**
> ShareAclSchema V1SharesShareIdAclObjectTypeObjectKeyPost(ctx, appID, authToken, shareId, objectType, objectKey, body)
Create a new share acl for an object

 Required roles:  - can_write_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **shareId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 
  **body** | [**ShareAclSchema**](ShareAclSchema.md)|  | 

### Return type

[**ShareAclSchema**](ShareACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesShareIdAclObjectTypeObjectKeyPut**
> ShareAclSchema V1SharesShareIdAclObjectTypeObjectKeyPut(ctx, appID, authToken, shareId, objectType, objectKey, body)
Update share acl for an object

 Required roles:  - can_write_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **shareId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 
  **body** | [**ShareAclSchema**](ShareAclSchema.md)|  | 

### Return type

[**ShareAclSchema**](ShareACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesShareIdAclObjectTypePut**
> CreateAcLsResultSchema V1SharesShareIdAclObjectTypePut(ctx, appID, authToken, shareId, objectType, body)
Create a new acl for multiple share objects

 Required roles:  - can_write_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **shareId** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**CreateShareAcLsSchema**](CreateShareAcLsSchema.md)|  | 

### Return type

[**CreateAcLsResultSchema**](CreateACLsResultSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdAclObjectTypeObjectKeyDelete**
> V1UsersUserIdAclObjectTypeObjectKeyDelete(ctx, appID, authToken, userId, objectType, objectKey)
Delete a user acl for an object

 Required roles:  - can_delete_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdAclObjectTypeObjectKeyGet**
> UserAclSchema V1UsersUserIdAclObjectTypeObjectKeyGet(ctx, appID, authToken, userId, objectType, objectKey)
List of user permissions for an object

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 

### Return type

[**UserAclSchema**](UserACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdAclObjectTypeObjectKeyPermissionGet**
> UserAclSchema V1UsersUserIdAclObjectTypeObjectKeyPermissionGet(ctx, appID, authToken, userId, objectType, objectKey, permission)
Returns a user acl for an object

 Required roles:  - can_read_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 
  **permission** | **string**|  | 

### Return type

[**UserAclSchema**](UserACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdAclObjectTypeObjectKeyPut**
> UserAclSchema V1UsersUserIdAclObjectTypeObjectKeyPut(ctx, appID, authToken, userId, objectType, objectKey, body)
Update or create user acl for an object

 Required roles:  - can_write_acls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectKey** | **string**|  | 
  **body** | [**UserAclSchema**](UserAclSchema.md)|  | 

### Return type

[**UserAclSchema**](UserACLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


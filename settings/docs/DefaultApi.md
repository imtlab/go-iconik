# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/settings*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CorsHostsCorsHostIdDelete**](DefaultApi.md#V1CorsHostsCorsHostIdDelete) | **Delete** /v1/cors_hosts/{cors_host_id}/ | Delete a particular CORS host by id
[**V1CorsHostsCorsHostIdGet**](DefaultApi.md#V1CorsHostsCorsHostIdGet) | **Get** /v1/cors_hosts/{cors_host_id}/ | Returns a particular CORS host by id
[**V1CorsHostsGet**](DefaultApi.md#V1CorsHostsGet) | **Get** /v1/cors_hosts/ | List of CORS hosts
[**V1CorsHostsPost**](DefaultApi.md#V1CorsHostsPost) | **Post** /v1/cors_hosts/ | Create a new CORS host
[**V1GroupGroupIdDelete**](DefaultApi.md#V1GroupGroupIdDelete) | **Delete** /v1/group/{group_id}/ | Delete group settings
[**V1GroupGroupIdGet**](DefaultApi.md#V1GroupGroupIdGet) | **Get** /v1/group/{group_id}/ | Group settings
[**V1GroupGroupIdPatch**](DefaultApi.md#V1GroupGroupIdPatch) | **Patch** /v1/group/{group_id}/ | Change group settings
[**V1GroupGroupIdPut**](DefaultApi.md#V1GroupGroupIdPut) | **Put** /v1/group/{group_id}/ | Change group settings
[**V1MergedCurrentGet**](DefaultApi.md#V1MergedCurrentGet) | **Get** /v1/merged/current/ | Get merged settings for current user
[**V1MergedUserIdGet**](DefaultApi.md#V1MergedUserIdGet) | **Get** /v1/merged/{user_id}/ | Get merged settings for a specific user
[**V1SystemCurrentGet**](DefaultApi.md#V1SystemCurrentGet) | **Get** /v1/system/current/ | System settings
[**V1SystemCurrentPatch**](DefaultApi.md#V1SystemCurrentPatch) | **Patch** /v1/system/current/ | Change system settings
[**V1SystemCurrentPut**](DefaultApi.md#V1SystemCurrentPut) | **Put** /v1/system/current/ | Change system settings
[**V1SystemSystemDomainIdGet**](DefaultApi.md#V1SystemSystemDomainIdGet) | **Get** /v1/system/{system_domain_id}/ | System settings
[**V1SystemSystemDomainIdPatch**](DefaultApi.md#V1SystemSystemDomainIdPatch) | **Patch** /v1/system/{system_domain_id}/ | Change system settings
[**V1SystemSystemDomainIdPut**](DefaultApi.md#V1SystemSystemDomainIdPut) | **Put** /v1/system/{system_domain_id}/ | Change system settings
[**V1UserUserIdDelete**](DefaultApi.md#V1UserUserIdDelete) | **Delete** /v1/user/{user_id}/ | Delete user settings
[**V1UserUserIdGet**](DefaultApi.md#V1UserUserIdGet) | **Get** /v1/user/{user_id}/ | User settings
[**V1UserUserIdPatch**](DefaultApi.md#V1UserUserIdPatch) | **Patch** /v1/user/{user_id}/ | Change user settings
[**V1UserUserIdPut**](DefaultApi.md#V1UserUserIdPut) | **Put** /v1/user/{user_id}/ | Change user settings


# **V1CorsHostsCorsHostIdDelete**
> V1CorsHostsCorsHostIdDelete(ctx, appID, authToken, corsHostId)
Delete a particular CORS host by id

 Required roles:  - can_delete_cors_hosts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **corsHostId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CorsHostsCorsHostIdGet**
> CorsHostSchema V1CorsHostsCorsHostIdGet(ctx, appID, authToken, corsHostId)
Returns a particular CORS host by id

 Required roles:  - can_read_cors_hosts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **corsHostId** | **string**|  | 

### Return type

[**CorsHostSchema**](CORSHostSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CorsHostsGet**
> CorsHostsSchema V1CorsHostsGet(ctx, appID, authToken, systemDomainId)
List of CORS hosts

 Required roles:  - can_read_cors_hosts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 

### Return type

[**CorsHostsSchema**](CORSHostsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CorsHostsPost**
> CorsHostSchema V1CorsHostsPost(ctx, appID, authToken, body)
Create a new CORS host

 Required roles:  - can_write_cors_hosts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**CorsHostSchema**](CorsHostSchema.md)|  | 

### Return type

[**CorsHostSchema**](CORSHostSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupGroupIdDelete**
> V1GroupGroupIdDelete(ctx, appID, authToken, groupId)
Delete group settings



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

# **V1GroupGroupIdGet**
> GroupSettingPublicSchema V1GroupGroupIdGet(ctx, appID, authToken, groupId)
Group settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

[**GroupSettingPublicSchema**](GroupSettingPublicSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupGroupIdPatch**
> GroupSettingPublicSchema V1GroupGroupIdPatch(ctx, appID, authToken, groupId, body)
Change group settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **body** | [**GroupSettingPublicSchema**](GroupSettingPublicSchema.md)|  | 

### Return type

[**GroupSettingPublicSchema**](GroupSettingPublicSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1GroupGroupIdPut**
> GroupSettingPublicSchema V1GroupGroupIdPut(ctx, appID, authToken, groupId, body)
Change group settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **body** | [**GroupSettingPublicSchema**](GroupSettingPublicSchema.md)|  | 

### Return type

[**GroupSettingPublicSchema**](GroupSettingPublicSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MergedCurrentGet**
> MergedSettingsSchema V1MergedCurrentGet(ctx, appID, authToken)
Get merged settings for current user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**MergedSettingsSchema**](MergedSettingsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MergedUserIdGet**
> MergedSettingsSchema V1MergedUserIdGet(ctx, appID, authToken, userId)
Get merged settings for a specific user



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**MergedSettingsSchema**](MergedSettingsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemCurrentGet**
> SystemSettingPublicSchema V1SystemCurrentGet(ctx, appID, authToken)
System settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemCurrentPatch**
> SystemSettingPublicSchema V1SystemCurrentPatch(ctx, appID, authToken, body)
Change system settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)|  | 

### Return type

[**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemCurrentPut**
> SystemSettingPublicSchema V1SystemCurrentPut(ctx, appID, authToken, body)
Change system settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)|  | 

### Return type

[**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemSystemDomainIdGet**
> SystemSettingPublicSchema V1SystemSystemDomainIdGet(ctx, appID, authToken, systemDomainId)
System settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 

### Return type

[**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemSystemDomainIdPatch**
> SystemSettingPublicSchema V1SystemSystemDomainIdPatch(ctx, appID, authToken, systemDomainId, body)
Change system settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 
  **body** | [**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)|  | 

### Return type

[**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemSystemDomainIdPut**
> SystemSettingPublicSchema V1SystemSystemDomainIdPut(ctx, appID, authToken, systemDomainId, body)
Change system settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 
  **body** | [**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)|  | 

### Return type

[**SystemSettingPublicSchema**](SystemSettingPublicSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserUserIdDelete**
> V1UserUserIdDelete(ctx, appID, authToken, userId)
Delete user settings



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

# **V1UserUserIdGet**
> UserSettingSchema V1UserUserIdGet(ctx, appID, authToken, userId)
User settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 

### Return type

[**UserSettingSchema**](UserSettingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserUserIdPatch**
> UserSettingSchema V1UserUserIdPatch(ctx, appID, authToken, userId, body)
Change user settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **body** | [**UserSettingSchema**](UserSettingSchema.md)|  | 

### Return type

[**UserSettingSchema**](UserSettingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserUserIdPut**
> UserSettingSchema V1UserUserIdPut(ctx, appID, authToken, userId, body)
Change user settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **userId** | **string**|  | 
  **body** | [**UserSettingSchema**](UserSettingSchema.md)|  | 

### Return type

[**UserSettingSchema**](UserSettingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


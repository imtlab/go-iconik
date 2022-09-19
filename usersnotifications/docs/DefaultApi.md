# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/users-notifications*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1NotificationSettingsGet**](DefaultApi.md#V1NotificationSettingsGet) | **Get** /v1/notification_settings/ | Returns a particular notification_setting by id
[**V1NotificationSettingsObjectTypeSubObjectTypeEventTypeProtocolGet**](DefaultApi.md#V1NotificationSettingsObjectTypeSubObjectTypeEventTypeProtocolGet) | **Get** /v1/notification_settings/{object_type}/{sub_object_type}/{event_type}/{protocol}/ | Returns a particular notification_setting by id
[**V1NotificationSettingsObjectTypeSubObjectTypeEventTypeProtocolPut**](DefaultApi.md#V1NotificationSettingsObjectTypeSubObjectTypeEventTypeProtocolPut) | **Put** /v1/notification_settings/{object_type}/{sub_object_type}/{event_type}/{protocol}/ | Create a new notification_setting
[**V1NotificationsAllReadPut**](DefaultApi.md#V1NotificationsAllReadPut) | **Put** /v1/notifications/all/read/ | Update notification
[**V1NotificationsGet**](DefaultApi.md#V1NotificationsGet) | **Get** /v1/notifications/ | Returns a list of notifications
[**V1NotificationsNotificationIdDelete**](DefaultApi.md#V1NotificationsNotificationIdDelete) | **Delete** /v1/notifications/{notification_id}/ | Delete a particular notification by id
[**V1NotificationsNotificationIdGet**](DefaultApi.md#V1NotificationsNotificationIdGet) | **Get** /v1/notifications/{notification_id}/ | Returns a particular notification by id
[**V1NotificationsPost**](DefaultApi.md#V1NotificationsPost) | **Post** /v1/notifications/ | Create a new notification
[**V1NotificationsSystemPost**](DefaultApi.md#V1NotificationsSystemPost) | **Post** /v1/notifications/system/ | Create a new system notification
[**V1ObjectTypeObjectIdSubscriptionsAllDelete**](DefaultApi.md#V1ObjectTypeObjectIdSubscriptionsAllDelete) | **Delete** /v1/{object_type}/{object_id}/subscriptions/all/ | Delete all user subscriptions for a specific object_type and object_id
[**V1ObjectTypeObjectIdSubscriptionsGet**](DefaultApi.md#V1ObjectTypeObjectIdSubscriptionsGet) | **Get** /v1/{object_type}/{object_id}/subscriptions/ | Returns user subscriptions for a specific object_type and object_id
[**V1SubscriptionsGet**](DefaultApi.md#V1SubscriptionsGet) | **Get** /v1/subscriptions/ | Returns all user subscriptions
[**V1SubscriptionsPost**](DefaultApi.md#V1SubscriptionsPost) | **Post** /v1/subscriptions/ | Create a new subscription
[**V1SubscriptionsSubscriptionIdDelete**](DefaultApi.md#V1SubscriptionsSubscriptionIdDelete) | **Delete** /v1/subscriptions/{subscription_id}/ | Delete a particular subscription by id
[**V1SubscriptionsSubscriptionIdGet**](DefaultApi.md#V1SubscriptionsSubscriptionIdGet) | **Get** /v1/subscriptions/{subscription_id}/ | Returns a particular subscription by id


# **V1NotificationSettingsGet**
> NotificationSettingsSchema V1NotificationSettingsGet(ctx, appID, authToken, optional)
Returns a particular notification_setting by id

 Required roles:  - can_read_notification_settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1NotificationSettingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1NotificationSettingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file set on previous page | 

### Return type

[**NotificationSettingsSchema**](NotificationSettingsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationSettingsObjectTypeSubObjectTypeEventTypeProtocolGet**
> NotificationSettingSchema V1NotificationSettingsObjectTypeSubObjectTypeEventTypeProtocolGet(ctx, appID, authToken, objectType, subObjectType, eventType, protocol)
Returns a particular notification_setting by id

 Required roles:  - can_read_notification_settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **subObjectType** | **string**|  | 
  **eventType** | **string**|  | 
  **protocol** | **string**|  | 

### Return type

[**NotificationSettingSchema**](NotificationSettingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationSettingsObjectTypeSubObjectTypeEventTypeProtocolPut**
> NotificationSettingSchema V1NotificationSettingsObjectTypeSubObjectTypeEventTypeProtocolPut(ctx, appID, authToken, objectType, subObjectType, eventType, protocol, body)
Create a new notification_setting



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **subObjectType** | **string**|  | 
  **eventType** | **string**|  | 
  **protocol** | **string**|  | 
  **body** | [**NotificationSettingSchema**](NotificationSettingSchema.md)|  | 

### Return type

[**NotificationSettingSchema**](NotificationSettingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationsAllReadPut**
> V1NotificationsAllReadPut(ctx, appID, authToken)
Update notification

 Required roles:  - can_read_notifications

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

# **V1NotificationsGet**
> NotificationsSchema V1NotificationsGet(ctx, appID, authToken, optional)
Returns a list of notifications

 Required roles:  - can_read_notifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1NotificationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1NotificationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file set on previous page | 

### Return type

[**NotificationsSchema**](NotificationsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationsNotificationIdDelete**
> V1NotificationsNotificationIdDelete(ctx, appID, authToken, notificationId)
Delete a particular notification by id

 Required roles:  - can_delete_notifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **notificationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationsNotificationIdGet**
> NotificationSchema V1NotificationsNotificationIdGet(ctx, appID, authToken, notificationId)
Returns a particular notification by id

 Required roles:  - can_read_notifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **notificationId** | **string**|  | 

### Return type

[**NotificationSchema**](NotificationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationsPost**
> NotificationSchema V1NotificationsPost(ctx, appID, authToken, body)
Create a new notification



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**NotificationSchema**](NotificationSchema.md)|  | 

### Return type

[**NotificationSchema**](NotificationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1NotificationsSystemPost**
> NotificationSchema V1NotificationsSystemPost(ctx, appID, authToken, body)
Create a new system notification



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**SystemNotificationSchema**](SystemNotificationSchema.md)|  | 

### Return type

[**NotificationSchema**](NotificationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSubscriptionsAllDelete**
> SubscriptionSchema V1ObjectTypeObjectIdSubscriptionsAllDelete(ctx, appID, authToken, objectType, objectId)
Delete all user subscriptions for a specific object_type and object_id

 Required roles:  - can_read_subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 

### Return type

[**SubscriptionSchema**](SubscriptionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSubscriptionsGet**
> SubscriptionsSchema V1ObjectTypeObjectIdSubscriptionsGet(ctx, appID, authToken, objectType, objectId)
Returns user subscriptions for a specific object_type and object_id

 Required roles:  - can_read_subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 

### Return type

[**SubscriptionsSchema**](SubscriptionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SubscriptionsGet**
> SubscriptionSchema V1SubscriptionsGet(ctx, appID, authToken)
Returns all user subscriptions

 Required roles:  - can_read_subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**SubscriptionSchema**](SubscriptionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SubscriptionsPost**
> SubscriptionSchema V1SubscriptionsPost(ctx, appID, authToken, body)
Create a new subscription



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**SubscriptionSchema**](SubscriptionSchema.md)|  | 

### Return type

[**SubscriptionSchema**](SubscriptionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SubscriptionsSubscriptionIdDelete**
> V1SubscriptionsSubscriptionIdDelete(ctx, appID, authToken, subscriptionId)
Delete a particular subscription by id

 Required roles:  - can_write_subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **subscriptionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SubscriptionsSubscriptionIdGet**
> SubscriptionSchema V1SubscriptionsSubscriptionIdGet(ctx, appID, authToken, subscriptionId)
Returns a particular subscription by id

 Required roles:  - can_read_subscriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **subscriptionId** | **string**|  | 

### Return type

[**SubscriptionSchema**](SubscriptionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


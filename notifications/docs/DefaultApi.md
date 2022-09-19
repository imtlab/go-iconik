# \DefaultApi

All URIs are relative to *http://app.iconik.io/API/webhooks*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1WebhooksGet**](DefaultApi.md#V1WebhooksGet) | **Get** /v1/webhooks/ | 
[**V1WebhooksPost**](DefaultApi.md#V1WebhooksPost) | **Post** /v1/webhooks/ | 
[**V1WebhooksWebhookIdDelete**](DefaultApi.md#V1WebhooksWebhookIdDelete) | **Delete** /v1/webhooks/{webhook_id}/ | 
[**V1WebhooksWebhookIdGet**](DefaultApi.md#V1WebhooksWebhookIdGet) | **Get** /v1/webhooks/{webhook_id}/ | 
[**V1WebhooksWebhookIdPut**](DefaultApi.md#V1WebhooksWebhookIdPut) | **Put** /v1/webhooks/{webhook_id}/ | 


# **V1WebhooksGet**
> interface{} V1WebhooksGet(ctx, authToken, appID)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WebhooksPost**
> interface{} V1WebhooksPost(ctx, authToken, appID, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **body** | [**Body**](Body.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WebhooksWebhookIdDelete**
> V1WebhooksWebhookIdDelete(ctx, authToken, appID, webhookId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **webhookId** | [**interface{}**](.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WebhooksWebhookIdGet**
> interface{} V1WebhooksWebhookIdGet(ctx, authToken, appID, webhookId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **webhookId** | [**interface{}**](.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WebhooksWebhookIdPut**
> interface{} V1WebhooksWebhookIdPut(ctx, authToken, appID, webhookId, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **webhookId** | [**interface{}**](.md)|  | 
  **body** | [**Body1**](Body1.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


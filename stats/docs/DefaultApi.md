# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/stats*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AssetsByPeriodGet**](DefaultApi.md#V1AssetsByPeriodGet) | **Get** /v1/assets/by/{period}/ | Returns all asset usage
[**V1AssetsPost**](DefaultApi.md#V1AssetsPost) | **Post** /v1/assets/ | Sets asset usage.
[**V1BillingChargesChargeIdReceiptUrlGet**](DefaultApi.md#V1BillingChargesChargeIdReceiptUrlGet) | **Get** /v1/billing/charges/{charge_id}/receipt_url/ | Returns billing receipt
[**V1BillingCreditsPost**](DefaultApi.md#V1BillingCreditsPost) | **Post** /v1/billing/credits/ | Add credits to an account
[**V1BillingCreditsPriceGet**](DefaultApi.md#V1BillingCreditsPriceGet) | **Get** /v1/billing/credits/price/ | Checks the total price that needs to be paid including VAT if it&#39;s needed
[**V1BillingCreditsVerifyPost**](DefaultApi.md#V1BillingCreditsVerifyPost) | **Post** /v1/billing/credits/verify/ | Verify status of add credits to an account
[**V1BillingCustomerCardDelete**](DefaultApi.md#V1BillingCustomerCardDelete) | **Delete** /v1/billing/customer/card/ | Creates billing customer card
[**V1BillingCustomerCardPost**](DefaultApi.md#V1BillingCustomerCardPost) | **Post** /v1/billing/customer/card/ | Creates billing customer card
[**V1BillingCustomerGet**](DefaultApi.md#V1BillingCustomerGet) | **Get** /v1/billing/customer/ | Returns billing customer
[**V1BillingCustomerPost**](DefaultApi.md#V1BillingCustomerPost) | **Post** /v1/billing/customer/ | Updates billing customer
[**V1BillingGet**](DefaultApi.md#V1BillingGet) | **Get** /v1/billing/ | Returns billing info
[**V1BillingInvoicesGet**](DefaultApi.md#V1BillingInvoicesGet) | **Get** /v1/billing/invoices/ | Returns billing invoices
[**V1BillingPost**](DefaultApi.md#V1BillingPost) | **Post** /v1/billing/ | Updates Billing (Requires super admin access).
[**V1BillingRecipientsGet**](DefaultApi.md#V1BillingRecipientsGet) | **Get** /v1/billing/recipients/ | Updates Billing Recipients
[**V1BillingRecipientsPut**](DefaultApi.md#V1BillingRecipientsPut) | **Put** /v1/billing/recipients/ | Updates Billing Recipients
[**V1BillingSettingsGet**](DefaultApi.md#V1BillingSettingsGet) | **Get** /v1/billing/settings/ | Updates Billing Settings
[**V1BillingSettingsPut**](DefaultApi.md#V1BillingSettingsPut) | **Put** /v1/billing/settings/ | Updates Billing Settings
[**V1BillingStatusGet**](DefaultApi.md#V1BillingStatusGet) | **Get** /v1/billing/status/ | Returns billing status
[**V1BillingSystemDomainIdBillingIdDelete**](DefaultApi.md#V1BillingSystemDomainIdBillingIdDelete) | **Delete** /v1/billing/{system_domain_id}/{billing_id}/ | Delete billing record (Requires super admin access).
[**V1CollectionsByPeriodGet**](DefaultApi.md#V1CollectionsByPeriodGet) | **Get** /v1/collections/by/{period}/ | Returns all collection usage
[**V1IdObjectIdInfoGet**](DefaultApi.md#V1IdObjectIdInfoGet) | **Get** /v1/id/{object_id}/info/ | Internal endpoint to convert ID to system domain
[**V1StorageAccessByPeriodGet**](DefaultApi.md#V1StorageAccessByPeriodGet) | **Get** /v1/storage/access/by/{period}/ | Returns storage_access for all storages
[**V1StorageUsageByPeriodGet**](DefaultApi.md#V1StorageUsageByPeriodGet) | **Get** /v1/storage/usage/by/{period}/ | Returns storage_usage for all storages
[**V1TranscoderUsageByPeriodGet**](DefaultApi.md#V1TranscoderUsageByPeriodGet) | **Get** /v1/transcoder/usage/by/{period}/ | Returns transcoder_usage for all transcoders
[**V1UserAuditByPeriodGet**](DefaultApi.md#V1UserAuditByPeriodGet) | **Get** /v1/user/audit/by/{period}/ | Returns all audit


# **V1AssetsByPeriodGet**
> AssetUsageSchema V1AssetsByPeriodGet(ctx, appID, authToken, period, optional)
Returns all asset usage

 Required roles:  - can_read_stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **period** | **string**| Period of stats (month or day) | 
 **optional** | ***DefaultApiV1AssetsByPeriodGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsByPeriodGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromDate** | **optional.String**| Filter by from_date | 
 **toDate** | **optional.String**| Filter by to_date | 

### Return type

[**AssetUsageSchema**](AssetUsageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsPost**
> AssetUsageSchema V1AssetsPost(ctx, appID, authToken, body)
Sets asset usage.

<br/>system_domain_id will be automatically added when <br/>posting to this end point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**AssetUsageSchema**](AssetUsageSchema.md)|  | 

### Return type

[**AssetUsageSchema**](AssetUsageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingChargesChargeIdReceiptUrlGet**
> BillingReceiptSchema V1BillingChargesChargeIdReceiptUrlGet(ctx, appID, authToken, chargeId)
Returns billing receipt

 Required roles:  - can_read_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **chargeId** | **string**|  | 

### Return type

[**BillingReceiptSchema**](BillingReceiptSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingCreditsPost**
> BillingCreditsSchema V1BillingCreditsPost(ctx, appID, authToken, body)
Add credits to an account

 Required roles:  - can_write_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BillingCreditsSchema**](BillingCreditsSchema.md)|  | 

### Return type

[**BillingCreditsSchema**](BillingCreditsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingCreditsPriceGet**
> CreditsSchema V1BillingCreditsPriceGet(ctx, appID, authToken, credits)
Checks the total price that needs to be paid including VAT if it's needed

 Required roles:  - can_write_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **credits** | **int32**|  | 

### Return type

[**CreditsSchema**](CreditsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingCreditsVerifyPost**
> BillingCreditsVerifySchema V1BillingCreditsVerifyPost(ctx, appID, authToken, body)
Verify status of add credits to an account

 Required roles:  - can_write_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BillingCreditsVerifySchema**](BillingCreditsVerifySchema.md)|  | 

### Return type

[**BillingCreditsVerifySchema**](BillingCreditsVerifySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingCustomerCardDelete**
> V1BillingCustomerCardDelete(ctx, appID, authToken)
Creates billing customer card

 Required roles:  - can_write_billing

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

# **V1BillingCustomerCardPost**
> BillingCustomerCardSchema V1BillingCustomerCardPost(ctx, appID, authToken, body)
Creates billing customer card

 Required roles:  - can_write_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BillingCustomerCardSchema**](BillingCustomerCardSchema.md)|  | 

### Return type

[**BillingCustomerCardSchema**](BillingCustomerCardSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingCustomerGet**
> BillingSchema V1BillingCustomerGet(ctx, appID, authToken, body)
Returns billing customer

 Required roles:  - can_read_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BillingSchema**](BillingSchema.md)|  | 

### Return type

[**BillingSchema**](BillingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingCustomerPost**
> BillingCustomerSchema V1BillingCustomerPost(ctx, appID, authToken, body)
Updates billing customer

 Required roles:  - can_write_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BillingCustomerSchema**](BillingCustomerSchema.md)|  | 

### Return type

[**BillingCustomerSchema**](BillingCustomerSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingGet**
> BillingsSchema V1BillingGet(ctx, appID, authToken, optional)
Returns billing info

 Required roles:  - can_read_stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1BillingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1BillingGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDate** | **optional.String**| Filter by from_date | 
 **toDate** | **optional.String**| Filter by from_date | 
 **perPage** | **optional.Int32**| The number of items for each page | [default to 100]
 **lastId** | **optional.String**| ID of a last file on previous page | 

### Return type

[**BillingsSchema**](BillingsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingInvoicesGet**
> V1BillingInvoicesGet(ctx, appID, authToken, optional)
Returns billing invoices

 Required roles:  - can_read_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1BillingInvoicesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1BillingInvoicesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startingAfter** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingPost**
> BillingSchema V1BillingPost(ctx, appID, authToken, body)
Updates Billing (Requires super admin access).



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BillingSchema**](BillingSchema.md)|  | 

### Return type

[**BillingSchema**](BillingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingRecipientsGet**
> BillingRecipientsSchema V1BillingRecipientsGet(ctx, appID, authToken)
Updates Billing Recipients

 Required roles:  - can_read_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**BillingRecipientsSchema**](BillingRecipientsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingRecipientsPut**
> BillingRecipientsSchema V1BillingRecipientsPut(ctx, appID, authToken, body)
Updates Billing Recipients

 Required roles:  - can_write_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BillingRecipientsSchema**](BillingRecipientsSchema.md)|  | 

### Return type

[**BillingRecipientsSchema**](BillingRecipientsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingSettingsGet**
> BillingSettingsSchema V1BillingSettingsGet(ctx, appID, authToken)
Updates Billing Settings

 Required roles:  - can_read_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**BillingSettingsSchema**](BillingSettingsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingSettingsPut**
> BillingSettingsSchema V1BillingSettingsPut(ctx, appID, authToken, body)
Updates Billing Settings

 Required roles:  - can_write_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BillingSettingsSchema**](BillingSettingsSchema.md)|  | 

### Return type

[**BillingSettingsSchema**](BillingSettingsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingStatusGet**
> BillingStatsSchema V1BillingStatusGet(ctx, appID, authToken)
Returns billing status

 Required roles:  - can_read_billing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**BillingStatsSchema**](BillingStatsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1BillingSystemDomainIdBillingIdDelete**
> V1BillingSystemDomainIdBillingIdDelete(ctx, appID, authToken, systemDomainId, billingId)
Delete billing record (Requires super admin access).



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 
  **billingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsByPeriodGet**
> CollectionUsageSchema V1CollectionsByPeriodGet(ctx, appID, authToken, period, optional)
Returns all collection usage

 Required roles:  - can_read_stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **period** | **string**| Period of stats (month or day) | 
 **optional** | ***DefaultApiV1CollectionsByPeriodGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectionsByPeriodGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromDate** | **optional.String**| Filter by from_date | 
 **toDate** | **optional.String**| Filter by to_date | 

### Return type

[**CollectionUsageSchema**](CollectionUsageSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1IdObjectIdInfoGet**
> V1IdObjectIdInfoGet(ctx, appID, authToken, objectId)
Internal endpoint to convert ID to system domain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**| Object ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StorageAccessByPeriodGet**
> StorageAccessesSchema V1StorageAccessByPeriodGet(ctx, appID, authToken, period, optional)
Returns storage_access for all storages

 Required roles:  - can_read_stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **period** | **string**| Period of stats (month or day) | 
 **optional** | ***DefaultApiV1StorageAccessByPeriodGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StorageAccessByPeriodGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromDate** | **optional.String**| Filter by from_date | 
 **toDate** | **optional.String**| Filter by to_date | 

### Return type

[**StorageAccessesSchema**](StorageAccessesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StorageUsageByPeriodGet**
> StorageUsagesSchema V1StorageUsageByPeriodGet(ctx, appID, authToken, period, optional)
Returns storage_usage for all storages

 Required roles:  - can_read_stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **period** | **string**| Period of stats (month or day) | 
 **optional** | ***DefaultApiV1StorageUsageByPeriodGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StorageUsageByPeriodGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromDate** | **optional.String**| Filter by from_date | 
 **toDate** | **optional.String**| Filter by to_date | 

### Return type

[**StorageUsagesSchema**](StorageUsagesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscoderUsageByPeriodGet**
> TranscoderUsagesSchema V1TranscoderUsageByPeriodGet(ctx, appID, authToken, period, optional)
Returns transcoder_usage for all transcoders

 Required roles:  - can_read_stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **period** | **string**|  | 
 **optional** | ***DefaultApiV1TranscoderUsageByPeriodGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1TranscoderUsageByPeriodGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromDate** | **optional.String**| Filter by from_date | 
 **toDate** | **optional.String**| Filter by to_date | 

### Return type

[**TranscoderUsagesSchema**](TranscoderUsagesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserAuditByPeriodGet**
> UserUsagesSchema V1UserAuditByPeriodGet(ctx, appID, authToken, period, optional)
Returns all audit

 Required roles:  - can_read_stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **period** | **string**| Period of stats (month, day or day_detailed) | 
 **optional** | ***DefaultApiV1UserAuditByPeriodGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1UserAuditByPeriodGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fromDate** | **optional.String**| Filter by from_date | 
 **toDate** | **optional.String**| Filter by to_date | 
 **systemDomainId** | **optional.String**| Filter by system_domain_id (Only for super admins) | 

### Return type

[**UserUsagesSchema**](UserUsagesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# Go API client for stats

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./stats"
```

## Documentation for API Endpoints

All URIs are relative to *https://app.iconik.io/API/stats*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**V1AssetsByPeriodGet**](docs/DefaultApi.md#v1assetsbyperiodget) | **Get** /v1/assets/by/{period}/ | Returns all asset usage
*DefaultApi* | [**V1AssetsPost**](docs/DefaultApi.md#v1assetspost) | **Post** /v1/assets/ | Sets asset usage.
*DefaultApi* | [**V1BillingChargesChargeIdReceiptUrlGet**](docs/DefaultApi.md#v1billingchargeschargeidreceipturlget) | **Get** /v1/billing/charges/{charge_id}/receipt_url/ | Returns billing receipt
*DefaultApi* | [**V1BillingCreditsPost**](docs/DefaultApi.md#v1billingcreditspost) | **Post** /v1/billing/credits/ | Add credits to an account
*DefaultApi* | [**V1BillingCreditsPriceGet**](docs/DefaultApi.md#v1billingcreditspriceget) | **Get** /v1/billing/credits/price/ | Checks the total price that needs to be paid including VAT if it&#39;s needed
*DefaultApi* | [**V1BillingCreditsVerifyPost**](docs/DefaultApi.md#v1billingcreditsverifypost) | **Post** /v1/billing/credits/verify/ | Verify status of add credits to an account
*DefaultApi* | [**V1BillingCustomerCardDelete**](docs/DefaultApi.md#v1billingcustomercarddelete) | **Delete** /v1/billing/customer/card/ | Creates billing customer card
*DefaultApi* | [**V1BillingCustomerCardPost**](docs/DefaultApi.md#v1billingcustomercardpost) | **Post** /v1/billing/customer/card/ | Creates billing customer card
*DefaultApi* | [**V1BillingCustomerGet**](docs/DefaultApi.md#v1billingcustomerget) | **Get** /v1/billing/customer/ | Returns billing customer
*DefaultApi* | [**V1BillingCustomerPost**](docs/DefaultApi.md#v1billingcustomerpost) | **Post** /v1/billing/customer/ | Updates billing customer
*DefaultApi* | [**V1BillingGet**](docs/DefaultApi.md#v1billingget) | **Get** /v1/billing/ | Returns billing info
*DefaultApi* | [**V1BillingInvoicesGet**](docs/DefaultApi.md#v1billinginvoicesget) | **Get** /v1/billing/invoices/ | Returns billing invoices
*DefaultApi* | [**V1BillingPost**](docs/DefaultApi.md#v1billingpost) | **Post** /v1/billing/ | Updates Billing (Requires super admin access).
*DefaultApi* | [**V1BillingRecipientsGet**](docs/DefaultApi.md#v1billingrecipientsget) | **Get** /v1/billing/recipients/ | Updates Billing Recipients
*DefaultApi* | [**V1BillingRecipientsPut**](docs/DefaultApi.md#v1billingrecipientsput) | **Put** /v1/billing/recipients/ | Updates Billing Recipients
*DefaultApi* | [**V1BillingSettingsGet**](docs/DefaultApi.md#v1billingsettingsget) | **Get** /v1/billing/settings/ | Updates Billing Settings
*DefaultApi* | [**V1BillingSettingsPut**](docs/DefaultApi.md#v1billingsettingsput) | **Put** /v1/billing/settings/ | Updates Billing Settings
*DefaultApi* | [**V1BillingStatusGet**](docs/DefaultApi.md#v1billingstatusget) | **Get** /v1/billing/status/ | Returns billing status
*DefaultApi* | [**V1BillingSystemDomainIdBillingIdDelete**](docs/DefaultApi.md#v1billingsystemdomainidbillingiddelete) | **Delete** /v1/billing/{system_domain_id}/{billing_id}/ | Delete billing record (Requires super admin access).
*DefaultApi* | [**V1CollectionsByPeriodGet**](docs/DefaultApi.md#v1collectionsbyperiodget) | **Get** /v1/collections/by/{period}/ | Returns all collection usage
*DefaultApi* | [**V1IdObjectIdInfoGet**](docs/DefaultApi.md#v1idobjectidinfoget) | **Get** /v1/id/{object_id}/info/ | Internal endpoint to convert ID to system domain
*DefaultApi* | [**V1StorageAccessByPeriodGet**](docs/DefaultApi.md#v1storageaccessbyperiodget) | **Get** /v1/storage/access/by/{period}/ | Returns storage_access for all storages
*DefaultApi* | [**V1StorageUsageByPeriodGet**](docs/DefaultApi.md#v1storageusagebyperiodget) | **Get** /v1/storage/usage/by/{period}/ | Returns storage_usage for all storages
*DefaultApi* | [**V1TranscoderUsageByPeriodGet**](docs/DefaultApi.md#v1transcoderusagebyperiodget) | **Get** /v1/transcoder/usage/by/{period}/ | Returns transcoder_usage for all transcoders
*DefaultApi* | [**V1UserAuditByPeriodGet**](docs/DefaultApi.md#v1userauditbyperiodget) | **Get** /v1/user/audit/by/{period}/ | Returns all audit


## Documentation For Models

 - [AssetUsageSchema](docs/AssetUsageSchema.md)
 - [AssetUsagesElasticSchema](docs/AssetUsagesElasticSchema.md)
 - [AssetUsagesSchema](docs/AssetUsagesSchema.md)
 - [BillingCreditsSchema](docs/BillingCreditsSchema.md)
 - [BillingCreditsVerifySchema](docs/BillingCreditsVerifySchema.md)
 - [BillingCustomerCardSchema](docs/BillingCustomerCardSchema.md)
 - [BillingCustomerSchema](docs/BillingCustomerSchema.md)
 - [BillingCustomerShippingAddressSchema](docs/BillingCustomerShippingAddressSchema.md)
 - [BillingCustomerShippingSchema](docs/BillingCustomerShippingSchema.md)
 - [BillingReceiptSchema](docs/BillingReceiptSchema.md)
 - [BillingRecipientsSchema](docs/BillingRecipientsSchema.md)
 - [BillingSchema](docs/BillingSchema.md)
 - [BillingSettingsSchema](docs/BillingSettingsSchema.md)
 - [BillingStatsSchema](docs/BillingStatsSchema.md)
 - [BillingsSchema](docs/BillingsSchema.md)
 - [CollectionUsageSchema](docs/CollectionUsageSchema.md)
 - [CollectionUsagesElasticSchema](docs/CollectionUsagesElasticSchema.md)
 - [CollectionUsagesSchema](docs/CollectionUsagesSchema.md)
 - [CreditsSchema](docs/CreditsSchema.md)
 - [DateFilterSchema](docs/DateFilterSchema.md)
 - [ListObjectsSchema](docs/ListObjectsSchema.md)
 - [StorageAccessElasticSchema](docs/StorageAccessElasticSchema.md)
 - [StorageAccessSchema](docs/StorageAccessSchema.md)
 - [StorageAccessesSchema](docs/StorageAccessesSchema.md)
 - [StorageUsageSchema](docs/StorageUsageSchema.md)
 - [StorageUsagesElasticSchema](docs/StorageUsagesElasticSchema.md)
 - [StorageUsagesSchema](docs/StorageUsagesSchema.md)
 - [TranscoderUsageSchema](docs/TranscoderUsageSchema.md)
 - [TranscoderUsagesElasticSchema](docs/TranscoderUsagesElasticSchema.md)
 - [TranscoderUsagesSchema](docs/TranscoderUsagesSchema.md)
 - [UserAuditSchema](docs/UserAuditSchema.md)
 - [UserUsagesDetailedSchema](docs/UserUsagesDetailedSchema.md)
 - [UserUsagesElasticSchema](docs/UserUsagesElasticSchema.md)
 - [UserUsagesSchema](docs/UserUsagesSchema.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



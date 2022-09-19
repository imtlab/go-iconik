# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/auth*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AppsAppIdDelete**](DefaultApi.md#V1AppsAppIdDelete) | **Delete** /v1/apps/{app_id}/ | Delete a particular app by id
[**V1AppsAppIdGet**](DefaultApi.md#V1AppsAppIdGet) | **Get** /v1/apps/{app_id}/ | Returns a particular app by id
[**V1AppsAppIdPatch**](DefaultApi.md#V1AppsAppIdPatch) | **Patch** /v1/apps/{app_id}/ | Update app
[**V1AppsAppIdPut**](DefaultApi.md#V1AppsAppIdPut) | **Put** /v1/apps/{app_id}/ | Update app
[**V1AppsAppIdTokenPost**](DefaultApi.md#V1AppsAppIdTokenPost) | **Post** /v1/apps/{app_id}/token/ | Creates app token by id and returns it&#39;s data
[**V1AppsExternalAuthPost**](DefaultApi.md#V1AppsExternalAuthPost) | **Post** /v1/apps/external/auth/ | Create a new token for the logged in user and store it for an external app
[**V1AppsExternalAuthSecretGet**](DefaultApi.md#V1AppsExternalAuthSecretGet) | **Get** /v1/apps/external/auth/{secret}/ | Gets a token requested by an external app
[**V1AppsGet**](DefaultApi.md#V1AppsGet) | **Get** /v1/apps/ | List of apps
[**V1AppsInstanceApprovedInstanceIdDelete**](DefaultApi.md#V1AppsInstanceApprovedInstanceIdDelete) | **Delete** /v1/apps/instance/{approved_instance_id}/ | Delete an approved instance of an app
[**V1AppsInstanceApprovedInstanceIdGet**](DefaultApi.md#V1AppsInstanceApprovedInstanceIdGet) | **Get** /v1/apps/instance/{approved_instance_id}/ | Gets an approved instance of an app
[**V1AppsInstancePost**](DefaultApi.md#V1AppsInstancePost) | **Post** /v1/apps/instance/ | Create a new app instance
[**V1AppsPost**](DefaultApi.md#V1AppsPost) | **Post** /v1/apps/ | Create a new app
[**V1AuthAdLoginPost**](DefaultApi.md#V1AuthAdLoginPost) | **Post** /v1/auth/ad/login/ | Login by ActiveDirectory
[**V1AuthOauthLoginPost**](DefaultApi.md#V1AuthOauthLoginPost) | **Post** /v1/auth/oauth/login/ | Login by OAuth
[**V1AuthSamlAcsPublicIdPost**](DefaultApi.md#V1AuthSamlAcsPublicIdPost) | **Post** /v1/auth/saml/acs/{public_id}/ | SAML Assertion Consumer Service
[**V1AuthSamlAcsSystemDomainIdIdentityProviderIdPost**](DefaultApi.md#V1AuthSamlAcsSystemDomainIdIdentityProviderIdPost) | **Post** /v1/auth/saml/acs/{system_domain_id}/{identity_provider_id}/ | SAML Assertion Consumer Service
[**V1AuthSamlDomainsDomainDelete**](DefaultApi.md#V1AuthSamlDomainsDomainDelete) | **Delete** /v1/auth/saml/domains/{domain}/ | Unbind domain from identity provider
[**V1AuthSamlDomainsPost**](DefaultApi.md#V1AuthSamlDomainsPost) | **Post** /v1/auth/saml/domains/ | Bind domain to identity provider
[**V1AuthSamlIdpConvertPost**](DefaultApi.md#V1AuthSamlIdpConvertPost) | **Post** /v1/auth/saml/idp/convert/ | Convert an IdP EntityDescriptor XML into json suitable as a settings configuration.
[**V1AuthSamlIdpGet**](DefaultApi.md#V1AuthSamlIdpGet) | **Get** /v1/auth/saml/idp/ | Get list of identity providers
[**V1AuthSamlIdpIdentityProviderIdDelete**](DefaultApi.md#V1AuthSamlIdpIdentityProviderIdDelete) | **Delete** /v1/auth/saml/idp/{identity_provider_id}/ | Delete a particular identity provider by id
[**V1AuthSamlIdpIdentityProviderIdGet**](DefaultApi.md#V1AuthSamlIdpIdentityProviderIdGet) | **Get** /v1/auth/saml/idp/{identity_provider_id}/ | Get a particular identity provider by id
[**V1AuthSamlIdpIdentityProviderIdPatch**](DefaultApi.md#V1AuthSamlIdpIdentityProviderIdPatch) | **Patch** /v1/auth/saml/idp/{identity_provider_id}/ | Update a particular identity provider by id
[**V1AuthSamlIdpIdentityProviderIdPut**](DefaultApi.md#V1AuthSamlIdpIdentityProviderIdPut) | **Put** /v1/auth/saml/idp/{identity_provider_id}/ | Update a particular identity provider by id
[**V1AuthSamlIdpPost**](DefaultApi.md#V1AuthSamlIdpPost) | **Post** /v1/auth/saml/idp/ | Create a new identity provider.
[**V1AuthSamlLoginEmailGet**](DefaultApi.md#V1AuthSamlLoginEmailGet) | **Get** /v1/auth/saml/login/{email}/ | SAML Single sign-on url by domain
[**V1AuthSamlLoginPost**](DefaultApi.md#V1AuthSamlLoginPost) | **Post** /v1/auth/saml/login/ | SAML Single sign-on url by domain
[**V1AuthSamlLogoutPublicIdPost**](DefaultApi.md#V1AuthSamlLogoutPublicIdPost) | **Post** /v1/auth/saml/logout/{public_id}/ | Initiate SAML Single logout
[**V1AuthSamlMetadataPublicIdGet**](DefaultApi.md#V1AuthSamlMetadataPublicIdGet) | **Get** /v1/auth/saml/metadata/{public_id}/ | SAML Single Logout Service
[**V1AuthSamlMetadataSystemDomainIdIdentityProviderIdGet**](DefaultApi.md#V1AuthSamlMetadataSystemDomainIdIdentityProviderIdGet) | **Get** /v1/auth/saml/metadata/{system_domain_id}/{identity_provider_id}/ | SAML Single Logout Service
[**V1AuthSamlSloPublicIdGet**](DefaultApi.md#V1AuthSamlSloPublicIdGet) | **Get** /v1/auth/saml/slo/{public_id}/ | SAML Single Logout Service
[**V1AuthSamlSloSystemDomainIdIdentityProviderIdGet**](DefaultApi.md#V1AuthSamlSloSystemDomainIdIdentityProviderIdGet) | **Get** /v1/auth/saml/slo/{system_domain_id}/{identity_provider_id}/ | SAML Single Logout Service
[**V1AuthSamlSsoPublicIdGet**](DefaultApi.md#V1AuthSamlSsoPublicIdGet) | **Get** /v1/auth/saml/sso/{public_id}/ | SAML Single sign-on Service
[**V1AuthSamlSsoSystemDomainIdIdentityProviderIdGet**](DefaultApi.md#V1AuthSamlSsoSystemDomainIdIdentityProviderIdGet) | **Get** /v1/auth/saml/sso/{system_domain_id}/{identity_provider_id}/ | SAML Single sign-on Service
[**V1AuthSimpleLoginPost**](DefaultApi.md#V1AuthSimpleLoginPost) | **Post** /v1/auth/simple/login/ | Login by using email and password
[**V1AuthTokenDelete**](DefaultApi.md#V1AuthTokenDelete) | **Delete** /v1/auth/token/ | Revoke token
[**V1AuthTokenGet**](DefaultApi.md#V1AuthTokenGet) | **Get** /v1/auth/token/ | Check if auth token valid
[**V1AuthTokenPut**](DefaultApi.md#V1AuthTokenPut) | **Put** /v1/auth/token/ | Refresh token
[**V1PasswordForgotPost**](DefaultApi.md#V1PasswordForgotPost) | **Post** /v1/password/forgot/ | Receives email address and sends email to this address with link for resetting password
[**V1PasswordResetResetHashPut**](DefaultApi.md#V1PasswordResetResetHashPut) | **Put** /v1/password/reset/{reset_hash}/ | Changes password to a new one
[**V1ReferralCodesCodeDelete**](DefaultApi.md#V1ReferralCodesCodeDelete) | **Delete** /v1/referral_codes/{code}/ | Delete a referral_code
[**V1ReferralCodesCodeGet**](DefaultApi.md#V1ReferralCodesCodeGet) | **Get** /v1/referral_codes/{code}/ | Get a referral_code
[**V1ReferralCodesGet**](DefaultApi.md#V1ReferralCodesGet) | **Get** /v1/referral_codes/ | Get all referral_codes
[**V1ReferralCodesPost**](DefaultApi.md#V1ReferralCodesPost) | **Post** /v1/referral_codes/ | Create a new referral_code
[**V1RegistrationsCountriesGet**](DefaultApi.md#V1RegistrationsCountriesGet) | **Get** /v1/registrations/countries/ | Returns list of countries
[**V1RegistrationsPost**](DefaultApi.md#V1RegistrationsPost) | **Post** /v1/registrations/ | Create a new registration
[**V1RegistrationsVerifyEmailHashPost**](DefaultApi.md#V1RegistrationsVerifyEmailHashPost) | **Post** /v1/registrations/verify/{email_hash}/ | Verify email address and create system domain from template is email address valid
[**V1SystemDomainsGet**](DefaultApi.md#V1SystemDomainsGet) | **Get** /v1/system_domains/ | List of system domains
[**V1SystemDomainsPost**](DefaultApi.md#V1SystemDomainsPost) | **Post** /v1/system_domains/ | Create a new system domain
[**V1SystemDomainsReferralCodeReferralCodePost**](DefaultApi.md#V1SystemDomainsReferralCodeReferralCodePost) | **Post** /v1/system_domains/referral_code/{referral_code}/ | Create a new system domain from a referral code (That is associated to your domain)
[**V1SystemDomainsSystemDomainIdDelete**](DefaultApi.md#V1SystemDomainsSystemDomainIdDelete) | **Delete** /v1/system_domains/{system_domain_id}/ | Delete a particular system_domain by id
[**V1SystemDomainsSystemDomainIdGet**](DefaultApi.md#V1SystemDomainsSystemDomainIdGet) | **Get** /v1/system_domains/{system_domain_id}/ | Returns a particular system domain by id
[**V1SystemDomainsSystemDomainIdLogoDelete**](DefaultApi.md#V1SystemDomainsSystemDomainIdLogoDelete) | **Delete** /v1/system_domains/{system_domain_id}/logo/ | Delete system domain logo image.
[**V1SystemDomainsSystemDomainIdLogoPost**](DefaultApi.md#V1SystemDomainsSystemDomainIdLogoPost) | **Post** /v1/system_domains/{system_domain_id}/logo/ | Upload system domain logo image.
[**V1SystemDomainsSystemDomainIdPatch**](DefaultApi.md#V1SystemDomainsSystemDomainIdPatch) | **Patch** /v1/system_domains/{system_domain_id}/ | Update system domain
[**V1SystemDomainsSystemDomainIdPut**](DefaultApi.md#V1SystemDomainsSystemDomainIdPut) | **Put** /v1/system_domains/{system_domain_id}/ | Update system domain
[**V1SystemDomainsTemplatesGet**](DefaultApi.md#V1SystemDomainsTemplatesGet) | **Get** /v1/system_domains/templates/ | List of system domain templates


# **V1AppsAppIdDelete**
> V1AppsAppIdDelete(ctx, appID, authToken, appId)
Delete a particular app by id

 Required roles:  - can_delete_apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **appId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsAppIdGet**
> AppSchema V1AppsAppIdGet(ctx, authToken, appId)
Returns a particular app by id

 Required roles:  - can_read_apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appId** | **string**|  | 

### Return type

[**AppSchema**](AppSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsAppIdPatch**
> AppSchema V1AppsAppIdPatch(ctx, appID, authToken, appId, body)
Update app

 Required roles:  - can_write_apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **appId** | **string**|  | 
  **body** | [**AppSchema**](AppSchema.md)|  | 

### Return type

[**AppSchema**](AppSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsAppIdPut**
> AppSchema V1AppsAppIdPut(ctx, appID, authToken, appId, body)
Update app

 Required roles:  - can_write_apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **appId** | **string**|  | 
  **body** | [**AppSchema**](AppSchema.md)|  | 

### Return type

[**AppSchema**](AppSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsAppIdTokenPost**
> TokenSchema V1AppsAppIdTokenPost(ctx, appID, authToken, appId)
Creates app token by id and returns it's data

 Required roles:  - can_read_apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **appId** | **string**|  | 

### Return type

[**TokenSchema**](TokenSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsExternalAuthPost**
> ExternalAuthRequestResponseSchema V1AppsExternalAuthPost(ctx, appID, authToken, body)
Create a new token for the logged in user and store it for an external app



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**ExternalAuthRequestSchema**](ExternalAuthRequestSchema.md)|  | 

### Return type

[**ExternalAuthRequestResponseSchema**](ExternalAuthRequestResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsExternalAuthSecretGet**
> ExternalAuthSchema V1AppsExternalAuthSecretGet(ctx, secret)
Gets a token requested by an external app



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **secret** | **string**|  | 

### Return type

[**ExternalAuthSchema**](ExternalAuthSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsGet**
> AppsSchema V1AppsGet(ctx, appID, authToken, optional)
List of apps

 Required roles:  - can_read_apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1AppsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AppsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last file set on previous page | 

### Return type

[**AppsSchema**](AppsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsInstanceApprovedInstanceIdDelete**
> V1AppsInstanceApprovedInstanceIdDelete(ctx, appID, authToken, approvedInstanceId)
Delete an approved instance of an app



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **approvedInstanceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsInstanceApprovedInstanceIdGet**
> ExternalAuthSchema V1AppsInstanceApprovedInstanceIdGet(ctx, appID, authToken, approvedInstanceId)
Gets an approved instance of an app



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **approvedInstanceId** | **string**|  | 

### Return type

[**ExternalAuthSchema**](ExternalAuthSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsInstancePost**
> ApprovedAppInstanceSchema V1AppsInstancePost(ctx, appID, authToken, body)
Create a new app instance



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**ApprovedAppInstanceSchema**](ApprovedAppInstanceSchema.md)|  | 

### Return type

[**ApprovedAppInstanceSchema**](ApprovedAppInstanceSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AppsPost**
> AppSchema V1AppsPost(ctx, appID, authToken, body)
Create a new app

 Required roles:  - can_write_apps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**AppSchema**](AppSchema.md)|  | 

### Return type

[**AppSchema**](AppSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthAdLoginPost**
> TokenSchema V1AuthAdLoginPost(ctx, body)
Login by ActiveDirectory

<br/>This function is not yet implemented.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | **interface{}**|  | 

### Return type

[**TokenSchema**](TokenSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthOauthLoginPost**
> TokenSchema V1AuthOauthLoginPost(ctx, body)
Login by OAuth

<br/>This function is not yet implemented.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | **interface{}**|  | 

### Return type

[**TokenSchema**](TokenSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlAcsPublicIdPost**
> V1AuthSamlAcsPublicIdPost(ctx, publicId)
SAML Assertion Consumer Service



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlAcsSystemDomainIdIdentityProviderIdPost**
> V1AuthSamlAcsSystemDomainIdIdentityProviderIdPost(ctx, systemDomainId, identityProviderId)
SAML Assertion Consumer Service



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemDomainId** | **string**|  | 
  **identityProviderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlDomainsDomainDelete**
> V1AuthSamlDomainsDomainDelete(ctx, appID, authToken, domain)
Unbind domain from identity provider



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **domain** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlDomainsPost**
> DomainIdentityProviderMapSchema V1AuthSamlDomainsPost(ctx, appID, authToken, body)
Bind domain to identity provider



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DomainIdentityProviderMapSchema**](DomainIdentityProviderMapSchema.md)|  | 

### Return type

[**DomainIdentityProviderMapSchema**](DomainIdentityProviderMapSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlIdpConvertPost**
> IdentityProviderSchema V1AuthSamlIdpConvertPost(ctx, appID, authToken, body)
Convert an IdP EntityDescriptor XML into json suitable as a settings configuration.

<br/>Input should be a SAML EntityDescriptor XML.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**IdentityProviderSchema**](IdentityProviderSchema.md)|  | 

### Return type

[**IdentityProviderSchema**](IdentityProviderSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlIdpGet**
> IdentityProvidersSchema V1AuthSamlIdpGet(ctx, appID, authToken, optional)
Get list of identity providers

 Required roles:  - can_read_identity_providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1AuthSamlIdpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AuthSamlIdpGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **lastId** | **optional.String**|  | 

### Return type

[**IdentityProvidersSchema**](IdentityProvidersSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlIdpIdentityProviderIdDelete**
> V1AuthSamlIdpIdentityProviderIdDelete(ctx, appID, authToken, identityProviderId)
Delete a particular identity provider by id

 Required roles:  - can_delete_identity_providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **identityProviderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlIdpIdentityProviderIdGet**
> IdentityProviderSchema V1AuthSamlIdpIdentityProviderIdGet(ctx, appID, authToken, identityProviderId)
Get a particular identity provider by id

 Required roles:  - can_read_identity_providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **identityProviderId** | **string**|  | 

### Return type

[**IdentityProviderSchema**](IdentityProviderSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlIdpIdentityProviderIdPatch**
> IdentityProviderSchema V1AuthSamlIdpIdentityProviderIdPatch(ctx, appID, authToken, identityProviderId, body)
Update a particular identity provider by id

 Required roles:  - can_write_identity_providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **identityProviderId** | **string**|  | 
  **body** | [**IdentityProviderSchema**](IdentityProviderSchema.md)|  | 

### Return type

[**IdentityProviderSchema**](IdentityProviderSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlIdpIdentityProviderIdPut**
> IdentityProviderSchema V1AuthSamlIdpIdentityProviderIdPut(ctx, appID, authToken, identityProviderId, body)
Update a particular identity provider by id

 Required roles:  - can_write_identity_providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **identityProviderId** | **string**|  | 
  **body** | [**IdentityProviderSchema**](IdentityProviderSchema.md)|  | 

### Return type

[**IdentityProviderSchema**](IdentityProviderSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlIdpPost**
> IdentityProviderSchema V1AuthSamlIdpPost(ctx, appID, authToken, body)
Create a new identity provider.

<br/>Input can either be an IdentityProviderSchema as json or a SAML<br/>EntityDescriptor XML. Required roles:  - can_write_identity_providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**IdentityProviderSchema**](IdentityProviderSchema.md)|  | 

### Return type

[**IdentityProviderSchema**](IdentityProviderSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlLoginEmailGet**
> InlineResponse200 V1AuthSamlLoginEmailGet(ctx, email)
SAML Single sign-on url by domain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlLoginPost**
> InlineResponse200 V1AuthSamlLoginPost(ctx, body)
SAML Single sign-on url by domain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SamlLoginSchema**](SamlLoginSchema.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlLogoutPublicIdPost**
> InlineResponse2001 V1AuthSamlLogoutPublicIdPost(ctx, publicId)
Initiate SAML Single logout



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicId** | **string**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlMetadataPublicIdGet**
> V1AuthSamlMetadataPublicIdGet(ctx, publicId)
SAML Single Logout Service



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlMetadataSystemDomainIdIdentityProviderIdGet**
> V1AuthSamlMetadataSystemDomainIdIdentityProviderIdGet(ctx, systemDomainId, identityProviderId)
SAML Single Logout Service



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemDomainId** | **string**|  | 
  **identityProviderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlSloPublicIdGet**
> V1AuthSamlSloPublicIdGet(ctx, publicId)
SAML Single Logout Service



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlSloSystemDomainIdIdentityProviderIdGet**
> V1AuthSamlSloSystemDomainIdIdentityProviderIdGet(ctx, systemDomainId, identityProviderId)
SAML Single Logout Service



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemDomainId** | **string**|  | 
  **identityProviderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlSsoPublicIdGet**
> V1AuthSamlSsoPublicIdGet(ctx, publicId)
SAML Single sign-on Service



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **publicId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSamlSsoSystemDomainIdIdentityProviderIdGet**
> V1AuthSamlSsoSystemDomainIdIdentityProviderIdGet(ctx, systemDomainId, identityProviderId)
SAML Single sign-on Service



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemDomainId** | **string**|  | 
  **identityProviderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthSimpleLoginPost**
> TokenSchema V1AuthSimpleLoginPost(ctx, body)
Login by using email and password



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SimpleLoginSchema**](SimpleLoginSchema.md)|  | 

### Return type

[**TokenSchema**](TokenSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthTokenDelete**
> V1AuthTokenDelete(ctx, appID, authToken)
Revoke token



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

# **V1AuthTokenGet**
> V1AuthTokenGet(ctx, appID, authToken)
Check if auth token valid



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

# **V1AuthTokenPut**
> TokenSchema V1AuthTokenPut(ctx, appID, authToken)
Refresh token



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**TokenSchema**](TokenSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PasswordForgotPost**
> V1PasswordForgotPost(ctx, body)
Receives email address and sends email to this address with link for resetting password



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ForgotPasswordSchema**](ForgotPasswordSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PasswordResetResetHashPut**
> V1PasswordResetResetHashPut(ctx, resetHash, body)
Changes password to a new one



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resetHash** | **string**|  | 
  **body** | [**ResetPasswordSchema**](ResetPasswordSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ReferralCodesCodeDelete**
> V1ReferralCodesCodeDelete(ctx, appID, authToken, code)
Delete a referral_code



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **code** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ReferralCodesCodeGet**
> ReferralCodeSchema V1ReferralCodesCodeGet(ctx, appID, authToken, code)
Get a referral_code



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **code** | **string**|  | 

### Return type

[**ReferralCodeSchema**](ReferralCodeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ReferralCodesGet**
> ReferralCodesSchema V1ReferralCodesGet(ctx, appID, authToken)
Get all referral_codes



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**ReferralCodesSchema**](ReferralCodesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ReferralCodesPost**
> ReferralCodeSchema V1ReferralCodesPost(ctx, appID, authToken, body)
Create a new referral_code



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**ReferralCodeSchema**](ReferralCodeSchema.md)|  | 

### Return type

[**ReferralCodeSchema**](ReferralCodeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1RegistrationsCountriesGet**
> CountriesSchema V1RegistrationsCountriesGet(ctx, )
Returns list of countries



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CountriesSchema**](CountriesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1RegistrationsPost**
> RegistrationSchema V1RegistrationsPost(ctx, body)
Create a new registration



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegistrationSchema**](RegistrationSchema.md)|  | 

### Return type

[**RegistrationSchema**](RegistrationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1RegistrationsVerifyEmailHashPost**
> SystemDomainFromTemplateSchema V1RegistrationsVerifyEmailHashPost(ctx, emailHash)
Verify email address and create system domain from template is email address valid



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailHash** | **string**|  | 

### Return type

[**SystemDomainFromTemplateSchema**](SystemDomainFromTemplateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsGet**
> SystemDomainsSchema V1SystemDomainsGet(ctx, appID, authToken, optional)
List of system domains



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1SystemDomainsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1SystemDomainsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **query** | **optional.String**| Query the name | 
 **statuses** | **optional.String**| Comma separated list of statuses to show | 

### Return type

[**SystemDomainsSchema**](SystemDomainsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsPost**
> SystemDomainSchema V1SystemDomainsPost(ctx, appID, authToken, body)
Create a new system domain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**SystemDomainSchema**](SystemDomainSchema.md)|  | 

### Return type

[**SystemDomainSchema**](SystemDomainSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsReferralCodeReferralCodePost**
> SystemDomainFromTemplateSchema V1SystemDomainsReferralCodeReferralCodePost(ctx, appID, authToken, referralCode, body)
Create a new system domain from a referral code (That is associated to your domain)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **referralCode** | **string**|  | 
  **body** | [**SystemDomainFromReferralCodeSchema**](SystemDomainFromReferralCodeSchema.md)|  | 

### Return type

[**SystemDomainFromTemplateSchema**](SystemDomainFromTemplateSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsSystemDomainIdDelete**
> V1SystemDomainsSystemDomainIdDelete(ctx, appID, authToken, systemDomainId)
Delete a particular system_domain by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsSystemDomainIdGet**
> SystemDomainSchema V1SystemDomainsSystemDomainIdGet(ctx, appID, authToken, systemDomainId)
Returns a particular system domain by id

 Required roles:  - can_read_system_domains

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 

### Return type

[**SystemDomainSchema**](SystemDomainSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsSystemDomainIdLogoDelete**
> V1SystemDomainsSystemDomainIdLogoDelete(ctx, appID, authToken, systemDomainId)
Delete system domain logo image.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsSystemDomainIdLogoPost**
> InlineResponse201 V1SystemDomainsSystemDomainIdLogoPost(ctx, appID, authToken, systemDomainId, logo)
Upload system domain logo image.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 
  **logo** | **string**|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsSystemDomainIdPatch**
> SystemDomainSchema V1SystemDomainsSystemDomainIdPatch(ctx, appID, authToken, systemDomainId, body)
Update system domain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 
  **body** | [**SystemDomainSchema**](SystemDomainSchema.md)|  | 

### Return type

[**SystemDomainSchema**](SystemDomainSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsSystemDomainIdPut**
> SystemDomainSchema V1SystemDomainsSystemDomainIdPut(ctx, appID, authToken, systemDomainId, body)
Update system domain



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **systemDomainId** | **string**|  | 
  **body** | [**SystemDomainSchema**](SystemDomainSchema.md)|  | 

### Return type

[**SystemDomainSchema**](SystemDomainSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SystemDomainsTemplatesGet**
> SystemDomainsSchema V1SystemDomainsTemplatesGet(ctx, appID, authToken)
List of system domain templates



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**SystemDomainsSchema**](SystemDomainsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


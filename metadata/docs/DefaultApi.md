# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/metadata*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AssetsAssetIdObjectTypeObjectIdVersionsVersionIdViewsViewIdGet**](DefaultApi.md#V1AssetsAssetIdObjectTypeObjectIdVersionsVersionIdViewsViewIdGet) | **Get** /v1/assets/{asset_id}/{object_type}/{object_id}/versions/{version_id}/views/{view_id}/ | Get asset metadata by object type, object ID, version ID and view ID
[**V1AssetsAssetIdObjectTypeObjectIdViewsViewIdGet**](DefaultApi.md#V1AssetsAssetIdObjectTypeObjectIdViewsViewIdGet) | **Get** /v1/assets/{asset_id}/{object_type}/{object_id}/views/{view_id}/ | Get asset metadata by object type, object ID and view ID
[**V1AssetsAssetIdObjectTypeObjectIdViewsViewIdPut**](DefaultApi.md#V1AssetsAssetIdObjectTypeObjectIdViewsViewIdPut) | **Put** /v1/assets/{asset_id}/{object_type}/{object_id}/views/{view_id}/ | Edit view metadata values for sub-objects of an asset (Such as segments)
[**V1AssetsAssetIdVersionsVersionIdViewsViewIdGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdViewsViewIdGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/views/{view_id}/ | Get object metadata by object type, object ID, version ID and view ID
[**V1FieldsFieldNameDelete**](DefaultApi.md#V1FieldsFieldNameDelete) | **Delete** /v1/fields/{field_name}/ | Delete a particular field by name
[**V1FieldsFieldNameGet**](DefaultApi.md#V1FieldsFieldNameGet) | **Get** /v1/fields/{field_name}/ | Returns a particular field by name
[**V1FieldsFieldNamePatch**](DefaultApi.md#V1FieldsFieldNamePatch) | **Patch** /v1/fields/{field_name}/ | Update field by name
[**V1FieldsFieldNamePut**](DefaultApi.md#V1FieldsFieldNamePut) | **Put** /v1/fields/{field_name}/ | Update field by name
[**V1FieldsGet**](DefaultApi.md#V1FieldsGet) | **Get** /v1/fields/ | List the fields defined in the system
[**V1FieldsPost**](DefaultApi.md#V1FieldsPost) | **Post** /v1/fields/ | Create a new field
[**V1MappingFieldsFieldNameGet**](DefaultApi.md#V1MappingFieldsFieldNameGet) | **Get** /v1/mapping/fields/{field_name}/ | Get the metadata field mapping
[**V1MappingFieldsPost**](DefaultApi.md#V1MappingFieldsPost) | **Post** /v1/mapping/fields/ | Create a new metadata field mapping
[**V1MappingOptionsGet**](DefaultApi.md#V1MappingOptionsGet) | **Get** /v1/mapping/options/ | List the metadata field mapping options
[**V1ObjectTypeCategoriesGet**](DefaultApi.md#V1ObjectTypeCategoriesGet) | **Get** /v1/{object_type}/categories/ | Get metadata categories
[**V1ObjectTypeCategoriesNameDelete**](DefaultApi.md#V1ObjectTypeCategoriesNameDelete) | **Delete** /v1/{object_type}/categories/{name}/ | Delete metadata category by object type and category name
[**V1ObjectTypeCategoriesNameGet**](DefaultApi.md#V1ObjectTypeCategoriesNameGet) | **Get** /v1/{object_type}/categories/{name}/ | Get metadata category by object type and category name
[**V1ObjectTypeCategoriesNamePut**](DefaultApi.md#V1ObjectTypeCategoriesNamePut) | **Put** /v1/{object_type}/categories/{name}/ | Edit metadata category for an object type
[**V1ObjectTypeCategoriesNameViewsGet**](DefaultApi.md#V1ObjectTypeCategoriesNameViewsGet) | **Get** /v1/{object_type}/categories/{name}/views/ | Get metadata views with field for object type and category
[**V1ObjectTypeCategoriesPost**](DefaultApi.md#V1ObjectTypeCategoriesPost) | **Post** /v1/{object_type}/categories/ | Add a metadata category for an object type
[**V1ObjectTypeContentViewsViewIdPut**](DefaultApi.md#V1ObjectTypeContentViewsViewIdPut) | **Put** /v1/{object_type}/content/views/{view_id}/ | Edit view metadata values for collection or saved search content.
[**V1ObjectTypeObjectIdGet**](DefaultApi.md#V1ObjectTypeObjectIdGet) | **Get** /v1/{object_type}/{object_id}/ | Get object metadata by object type and object ID
[**V1ObjectTypeObjectIdPut**](DefaultApi.md#V1ObjectTypeObjectIdPut) | **Put** /v1/{object_type}/{object_id}/ | Edit metadata values directly without a view. Admin access required.
[**V1ObjectTypeObjectIdViewsViewIdGet**](DefaultApi.md#V1ObjectTypeObjectIdViewsViewIdGet) | **Get** /v1/{object_type}/{object_id}/views/{view_id}/ | Get object metadata by object type, object ID and view ID
[**V1ObjectTypeObjectIdViewsViewIdPut**](DefaultApi.md#V1ObjectTypeObjectIdViewsViewIdPut) | **Put** /v1/{object_type}/{object_id}/views/{view_id}/ | Edit view metadata values for a single object
[**V1ObjectTypeViewsViewIdPut**](DefaultApi.md#V1ObjectTypeViewsViewIdPut) | **Put** /v1/{object_type}/views/{view_id}/ | Edit view metadata values for multiple objects (Assets, Collections or Segments)
[**V1UserFieldsGet**](DefaultApi.md#V1UserFieldsGet) | **Get** /v1/user/fields/ | List the fields that can be accessed by a user
[**V1ViewsGet**](DefaultApi.md#V1ViewsGet) | **Get** /v1/views/ | List the views defined in the system
[**V1ViewsPost**](DefaultApi.md#V1ViewsPost) | **Post** /v1/views/ | Create a new view
[**V1ViewsViewIdDelete**](DefaultApi.md#V1ViewsViewIdDelete) | **Delete** /v1/views/{view_id}/ | Delete a particular view by id
[**V1ViewsViewIdGet**](DefaultApi.md#V1ViewsViewIdGet) | **Get** /v1/views/{view_id}/ | Returns a particular view by id
[**V1ViewsViewIdPatch**](DefaultApi.md#V1ViewsViewIdPatch) | **Patch** /v1/views/{view_id}/ | Update view
[**V1ViewsViewIdPut**](DefaultApi.md#V1ViewsViewIdPut) | **Put** /v1/views/{view_id}/ | Update view


# **V1AssetsAssetIdObjectTypeObjectIdVersionsVersionIdViewsViewIdGet**
> MetadataValuesSchema V1AssetsAssetIdObjectTypeObjectIdVersionsVersionIdViewsViewIdGet(ctx, appID, authToken, objectType, assetId, objectId, versionId, viewId)
Get asset metadata by object type, object ID, version ID and view ID

 Required roles:  - can_read_metadata_values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **assetId** | **string**|  | 
  **objectId** | **string**|  | 
  **versionId** | **string**|  | 
  **viewId** | **string**|  | 

### Return type

[**MetadataValuesSchema**](MetadataValuesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdObjectTypeObjectIdViewsViewIdGet**
> MetadataValuesSchema V1AssetsAssetIdObjectTypeObjectIdViewsViewIdGet(ctx, appID, authToken, objectType, assetId, objectId, viewId)
Get asset metadata by object type, object ID and view ID

 Required roles:  - can_read_metadata_values Required roles:  - can_read_metadata_values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **assetId** | **string**|  | 
  **objectId** | **string**|  | 
  **viewId** | **string**|  | 

### Return type

[**MetadataValuesSchema**](MetadataValuesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdObjectTypeObjectIdViewsViewIdPut**
> MetadataValuesSchema V1AssetsAssetIdObjectTypeObjectIdViewsViewIdPut(ctx, appID, authToken, assetId, objectType, objectId, viewId, body)
Edit view metadata values for sub-objects of an asset (Such as segments)

 Required roles:  - can_write_metadata_values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **viewId** | **string**|  | 
  **body** | [**MetadataValuesInputSchema**](MetadataValuesInputSchema.md)|  | 

### Return type

[**MetadataValuesSchema**](MetadataValuesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdViewsViewIdGet**
> MetadataValuesSchema V1AssetsAssetIdVersionsVersionIdViewsViewIdGet(ctx, appID, authToken, assetId, versionId, viewId)
Get object metadata by object type, object ID, version ID and view ID

 Required roles:  - can_read_metadata_values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **viewId** | **string**|  | 

### Return type

[**MetadataValuesSchema**](MetadataValuesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FieldsFieldNameDelete**
> V1FieldsFieldNameDelete(ctx, appID, authToken, fieldName)
Delete a particular field by name

 Required roles:  - can_delete_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **fieldName** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FieldsFieldNameGet**
> MetadataFieldSchema V1FieldsFieldNameGet(ctx, appID, authToken, fieldName)
Returns a particular field by name

 Required roles:  - can_read_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **fieldName** | **string**|  | 

### Return type

[**MetadataFieldSchema**](MetadataFieldSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FieldsFieldNamePatch**
> MetadataFieldSchema V1FieldsFieldNamePatch(ctx, appID, authToken, fieldName, body)
Update field by name

 Required roles:  - can_write_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **fieldName** | **string**|  | 
  **body** | [**MetadataFieldSchema**](MetadataFieldSchema.md)|  | 

### Return type

[**MetadataFieldSchema**](MetadataFieldSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FieldsFieldNamePut**
> MetadataFieldSchema V1FieldsFieldNamePut(ctx, appID, authToken, fieldName, body)
Update field by name

 Required roles:  - can_write_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **fieldName** | **string**|  | 
  **body** | [**MetadataFieldSchema**](MetadataFieldSchema.md)|  | 

### Return type

[**MetadataFieldSchema**](MetadataFieldSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FieldsGet**
> MetadataFieldsSchema V1FieldsGet(ctx, appID, authToken, optional)
List the fields defined in the system

 Required roles:  - can_read_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1FieldsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1FieldsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page (Default 500) | 
 **lastFieldName** | **optional.String**| If your request returns per_page entries, send the last value of \&quot;name\&quot; to fetch next page | 
 **filter** | **optional.String**| A comma separated list of fieldnames For example - first_name,last_name,salary | 

### Return type

[**MetadataFieldsSchema**](MetadataFieldsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1FieldsPost**
> MetadataFieldSchema V1FieldsPost(ctx, appID, authToken, body)
Create a new field

 Required roles:  - can_write_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**MetadataFieldCreateSchema**](MetadataFieldCreateSchema.md)|  | 

### Return type

[**MetadataFieldSchema**](MetadataFieldSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MappingFieldsFieldNameGet**
> MetadataFieldMappingSchema V1MappingFieldsFieldNameGet(ctx, appID, authToken, fieldName)
Get the metadata field mapping

 Required roles:  - can_read_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **fieldName** | **string**|  | 

### Return type

[**MetadataFieldMappingSchema**](MetadataFieldMappingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MappingFieldsPost**
> MetadataFieldMappingSchema V1MappingFieldsPost(ctx, appID, authToken, body)
Create a new metadata field mapping

 Required roles:  - can_read_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**MetadataFieldMappingSchema**](MetadataFieldMappingSchema.md)|  | 

### Return type

[**MetadataFieldMappingSchema**](MetadataFieldMappingSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1MappingOptionsGet**
> MetadataFieldMappingOptionsSchema V1MappingOptionsGet(ctx, appID, authToken)
List the metadata field mapping options

 Required roles:  - can_read_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**MetadataFieldMappingOptionsSchema**](MetadataFieldMappingOptionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeCategoriesGet**
> MetadataCategoriesSchema V1ObjectTypeCategoriesGet(ctx, appID, authToken, objectType)
Get metadata categories

 Required roles:  - can_read_metadata_categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 

### Return type

[**MetadataCategoriesSchema**](MetadataCategoriesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeCategoriesNameDelete**
> V1ObjectTypeCategoriesNameDelete(ctx, appID, authToken, objectType, name)
Delete metadata category by object type and category name

 Required roles:  - can_delete_metadata_categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeCategoriesNameGet**
> MetadataCategorySchema V1ObjectTypeCategoriesNameGet(ctx, appID, authToken, objectType, name)
Get metadata category by object type and category name

 Required roles:  - can_read_metadata_categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**MetadataCategorySchema**](MetadataCategorySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeCategoriesNamePut**
> MetadataCategorySchema V1ObjectTypeCategoriesNamePut(ctx, appID, authToken, objectType, name, body)
Edit metadata category for an object type

 Required roles:  - can_write_metadata_categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **name** | **string**|  | 
  **body** | [**MetadataCategorySchema**](MetadataCategorySchema.md)|  | 

### Return type

[**MetadataCategorySchema**](MetadataCategorySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeCategoriesNameViewsGet**
> MetadataCategorySchema V1ObjectTypeCategoriesNameViewsGet(ctx, appID, authToken, objectType, name, optional)
Get metadata views with field for object type and category

 Required roles:  - can_read_metadata_categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **name** | **string**|  | 
 **optional** | ***DefaultApiV1ObjectTypeCategoriesNameViewsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ObjectTypeCategoriesNameViewsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **extOptions** | **optional.Bool**|  | 
 **writableOnly** | **optional.Bool**|  | 

### Return type

[**MetadataCategorySchema**](MetadataCategorySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeCategoriesPost**
> MetadataCategorySchema V1ObjectTypeCategoriesPost(ctx, appID, authToken, objectType, body)
Add a metadata category for an object type

 Required roles:  - can_write_metadata_categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**MetadataCategorySchema**](MetadataCategorySchema.md)|  | 

### Return type

[**MetadataCategorySchema**](MetadataCategorySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeContentViewsViewIdPut**
> V1ObjectTypeContentViewsViewIdPut(ctx, appID, authToken, objectType, viewId, body)
Edit view metadata values for collection or saved search content.

 Required roles:  - can_write_metadata_values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **viewId** | **string**|  | 
  **body** | [**CollectionMetadataValuesBatchSchema**](CollectionMetadataValuesBatchSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdGet**
> MetadataValuesSchema V1ObjectTypeObjectIdGet(ctx, appID, authToken, objectType, objectId)
Get object metadata by object type and object ID

This view should be available only internally

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 

### Return type

[**MetadataValuesSchema**](MetadataValuesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdPut**
> MetadataValuesSchema V1ObjectTypeObjectIdPut(ctx, appID, authToken, objectType, objectId, body)
Edit metadata values directly without a view. Admin access required.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **body** | [**MetadataValuesInputSchema**](MetadataValuesInputSchema.md)|  | 

### Return type

[**MetadataValuesSchema**](MetadataValuesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdViewsViewIdGet**
> MetadataValuesSchema V1ObjectTypeObjectIdViewsViewIdGet(ctx, appID, authToken, objectType, objectId, viewId)
Get object metadata by object type, object ID and view ID

 Required roles:  - can_read_metadata_values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **viewId** | **string**|  | 

### Return type

[**MetadataValuesSchema**](MetadataValuesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdViewsViewIdPut**
> MetadataValuesSchema V1ObjectTypeObjectIdViewsViewIdPut(ctx, appID, authToken, objectType, objectId, viewId, body)
Edit view metadata values for a single object

 Required roles:  - can_write_metadata_values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **viewId** | **string**|  | 
  **body** | [**MetadataValuesInputSchema**](MetadataValuesInputSchema.md)|  | 

### Return type

[**MetadataValuesSchema**](MetadataValuesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeViewsViewIdPut**
> V1ObjectTypeViewsViewIdPut(ctx, appID, authToken, objectType, viewId, body)
Edit view metadata values for multiple objects (Assets, Collections or Segments)

 Required roles:  - can_write_metadata_values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **viewId** | **string**|  | 
  **body** | [**MetadataValuesBatchSchema**](MetadataValuesBatchSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserFieldsGet**
> MetadataFieldsSchema V1UserFieldsGet(ctx, appID, authToken)
List the fields that can be accessed by a user

 Required roles:  - can_read_metadata_fields

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**MetadataFieldsSchema**](MetadataFieldsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ViewsGet**
> MetadataViewsSchema V1ViewsGet(ctx, appID, authToken, optional)
List the views defined in the system

 Required roles:  - can_read_metadata_views

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1ViewsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ViewsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **filter** | **optional.String**| A comma separated list of fieldnames with order For example - first_name,eq,Vlad;last_name,eq,Gudkov | 

### Return type

[**MetadataViewsSchema**](MetadataViewsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ViewsPost**
> MetadataViewSchema V1ViewsPost(ctx, appID, authToken, body)
Create a new view

 Required roles:  - can_write_metadata_views

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**MetadataViewInputSchema**](MetadataViewInputSchema.md)|  | 

### Return type

[**MetadataViewSchema**](MetadataViewSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ViewsViewIdDelete**
> V1ViewsViewIdDelete(ctx, appID, authToken, viewId)
Delete a particular view by id

 Required roles:  - can_delete_metadata_views

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **viewId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ViewsViewIdGet**
> MetadataViewSchema V1ViewsViewIdGet(ctx, appID, authToken, viewId, optional)
Returns a particular view by id

 Required roles:  - can_read_metadata_views

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **viewId** | **string**|  | 
 **optional** | ***DefaultApiV1ViewsViewIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ViewsViewIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mergeFields** | **optional.Bool**|  | [default to true]

### Return type

[**MetadataViewSchema**](MetadataViewSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ViewsViewIdPatch**
> MetadataViewSchema V1ViewsViewIdPatch(ctx, appID, authToken, viewId, body)
Update view

 Required roles:  - can_write_metadata_views

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **viewId** | **string**|  | 
  **body** | [**MetadataViewInputSchema**](MetadataViewInputSchema.md)|  | 

### Return type

[**MetadataViewSchema**](MetadataViewSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ViewsViewIdPut**
> MetadataViewSchema V1ViewsViewIdPut(ctx, appID, authToken, viewId, body)
Update view

 Required roles:  - can_write_metadata_views

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **viewId** | **string**|  | 
  **body** | [**MetadataViewInputSchema**](MetadataViewInputSchema.md)|  | 

### Return type

[**MetadataViewSchema**](MetadataViewSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


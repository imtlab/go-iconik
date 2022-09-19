# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/search*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1DiscoveryDefaultEntitiesAdminGet**](DefaultApi.md#V1DiscoveryDefaultEntitiesAdminGet) | **Get** /v1/discovery/default/entities/admin/ | Returns the discovery entities that are used to build the discovery view.
[**V1DiscoveryDefaultEntitiesEntityIdDelete**](DefaultApi.md#V1DiscoveryDefaultEntitiesEntityIdDelete) | **Delete** /v1/discovery/default/entities/{entity_id}/ | Delete a discovery entity by id
[**V1DiscoveryDefaultEntitiesEntityIdGet**](DefaultApi.md#V1DiscoveryDefaultEntitiesEntityIdGet) | **Get** /v1/discovery/default/entities/{entity_id}/ | Returns discovery entity
[**V1DiscoveryDefaultEntitiesEntityIdPatch**](DefaultApi.md#V1DiscoveryDefaultEntitiesEntityIdPatch) | **Patch** /v1/discovery/default/entities/{entity_id}/ | Update a discovery entity by id
[**V1DiscoveryDefaultEntitiesEntityIdPut**](DefaultApi.md#V1DiscoveryDefaultEntitiesEntityIdPut) | **Put** /v1/discovery/default/entities/{entity_id}/ | Update a discovery entity by id
[**V1DiscoveryDefaultEntitiesGet**](DefaultApi.md#V1DiscoveryDefaultEntitiesGet) | **Get** /v1/discovery/default/entities/ | Returns the discovery entities that are used to build the discovery view.
[**V1DiscoveryDefaultEntitiesPost**](DefaultApi.md#V1DiscoveryDefaultEntitiesPost) | **Post** /v1/discovery/default/entities/ | Adds a new discovery entity.
[**V1DiscoveryDefaultPut**](DefaultApi.md#V1DiscoveryDefaultPut) | **Put** /v1/discovery/default/ | Update default discovery view
[**V1SearchHistoryGet**](DefaultApi.md#V1SearchHistoryGet) | **Get** /v1/search/history/ | Returns the current search history
[**V1SearchHistorySearchHistoryIdDelete**](DefaultApi.md#V1SearchHistorySearchHistoryIdDelete) | **Delete** /v1/search/history/{search_history_id}/ | Delete a search from history by its id
[**V1SearchHistorySearchHistoryIdGet**](DefaultApi.md#V1SearchHistorySearchHistoryIdGet) | **Get** /v1/search/history/{search_history_id}/ | Returns results of search history
[**V1SearchPost**](DefaultApi.md#V1SearchPost) | **Post** /v1/search/ | Search
[**V1SearchSavedGet**](DefaultApi.md#V1SearchSavedGet) | **Get** /v1/search/saved/ | Returns list of saved searches
[**V1SearchSavedGroupGroupIdDelete**](DefaultApi.md#V1SearchSavedGroupGroupIdDelete) | **Delete** /v1/search/saved/group/{group_id}/ | Delete a saved search group by it&#39;s id
[**V1SearchSavedGroupGroupIdGet**](DefaultApi.md#V1SearchSavedGroupGroupIdGet) | **Get** /v1/search/saved/group/{group_id}/ | Returns saved search group data
[**V1SearchSavedGroupGroupIdPatch**](DefaultApi.md#V1SearchSavedGroupGroupIdPatch) | **Patch** /v1/search/saved/group/{group_id}/ | Update and return saved search group data
[**V1SearchSavedGroupGroupIdPut**](DefaultApi.md#V1SearchSavedGroupGroupIdPut) | **Put** /v1/search/saved/group/{group_id}/ | Update and return saved search group data
[**V1SearchSavedGroupGroupIdSearchSearchIdDelete**](DefaultApi.md#V1SearchSavedGroupGroupIdSearchSearchIdDelete) | **Delete** /v1/search/saved/group/{group_id}/search/{search_id}/ | Delete saved search from search group
[**V1SearchSavedGroupGroupIdSearchSearchIdPost**](DefaultApi.md#V1SearchSavedGroupGroupIdSearchSearchIdPost) | **Post** /v1/search/saved/group/{group_id}/search/{search_id}/ | Adds saved search to group
[**V1SearchSavedGroupPost**](DefaultApi.md#V1SearchSavedGroupPost) | **Post** /v1/search/saved/group/ | Create and return saved search group data
[**V1SearchSavedGroupsGet**](DefaultApi.md#V1SearchSavedGroupsGet) | **Get** /v1/search/saved/groups/ | Returns paginated list of search groups
[**V1SearchSavedGroupsGroupIdReindexPost**](DefaultApi.md#V1SearchSavedGroupsGroupIdReindexPost) | **Post** /v1/search/saved/groups/{group_id}/reindex/ | Reindex a particular saved search group by id
[**V1SearchSavedPost**](DefaultApi.md#V1SearchSavedPost) | **Post** /v1/search/saved/ | Search, save and return result of this search
[**V1SearchSavedSearchIdDelete**](DefaultApi.md#V1SearchSavedSearchIdDelete) | **Delete** /v1/search/saved/{search_id}/ | Delete a saved search by its id
[**V1SearchSavedSearchIdGet**](DefaultApi.md#V1SearchSavedSearchIdGet) | **Get** /v1/search/saved/{search_id}/ | Returns results of saved search
[**V1SearchSavedSearchIdPut**](DefaultApi.md#V1SearchSavedSearchIdPut) | **Put** /v1/search/saved/{search_id}/ | Search and save this search
[**V1SearchSavedSearchIdReindexPost**](DefaultApi.md#V1SearchSavedSearchIdReindexPost) | **Post** /v1/search/saved/{search_id}/reindex/ | Reindex a particular saved search by id
[**V1SearchSuggestPost**](DefaultApi.md#V1SearchSuggestPost) | **Post** /v1/search/suggest/ | Returns search suggestions for a particular query.


# **V1DiscoveryDefaultEntitiesAdminGet**
> DiscoveryEntitiesSchema V1DiscoveryDefaultEntitiesAdminGet(ctx, appID, authToken)
Returns the discovery entities that are used to build the discovery view.

 Required roles:  - can_write_discovery_entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**DiscoveryEntitiesSchema**](DiscoveryEntitiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DiscoveryDefaultEntitiesEntityIdDelete**
> V1DiscoveryDefaultEntitiesEntityIdDelete(ctx, appID, authToken, entityId)
Delete a discovery entity by id

 Required roles:  - can_delete_discovery_entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **entityId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DiscoveryDefaultEntitiesEntityIdGet**
> DiscoveryEntitySchema V1DiscoveryDefaultEntitiesEntityIdGet(ctx, appID, authToken, entityId)
Returns discovery entity

 Required roles:  - can_read_discovery_entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **entityId** | **string**|  | 

### Return type

[**DiscoveryEntitySchema**](DiscoveryEntitySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DiscoveryDefaultEntitiesEntityIdPatch**
> V1DiscoveryDefaultEntitiesEntityIdPatch(ctx, appID, authToken, entityId, body)
Update a discovery entity by id

 Required roles:  - can_write_discovery_entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **entityId** | **string**|  | 
  **body** | [**DiscoveryEntitySchema**](DiscoveryEntitySchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DiscoveryDefaultEntitiesEntityIdPut**
> V1DiscoveryDefaultEntitiesEntityIdPut(ctx, appID, authToken, entityId, body)
Update a discovery entity by id

 Required roles:  - can_write_discovery_entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **entityId** | **string**|  | 
  **body** | [**DiscoveryEntitySchema**](DiscoveryEntitySchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DiscoveryDefaultEntitiesGet**
> DiscoveryEntitiesSchema V1DiscoveryDefaultEntitiesGet(ctx, appID, authToken)
Returns the discovery entities that are used to build the discovery view.

 Required roles:  - can_read_discovery_entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**DiscoveryEntitiesSchema**](DiscoveryEntitiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DiscoveryDefaultEntitiesPost**
> DiscoveryEntitySchema V1DiscoveryDefaultEntitiesPost(ctx, appID, authToken, body)
Adds a new discovery entity.

<br/>This creates an entry for the discovery view to show collections and saved searches.<br/>Object Type should be one of COLLECTION, SAVED_SEARCH, ASSET, RECOMMENDATION or TRENDING<br/>Object ID is only needed in the case of COLLECTION, SAVED_SEARCH or ASSET<br/>metadata is for user defined extra data.<br/> Required roles:  - can_write_discovery_entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DiscoveryEntitySchema**](DiscoveryEntitySchema.md)|  | 

### Return type

[**DiscoveryEntitySchema**](DiscoveryEntitySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DiscoveryDefaultPut**
> DiscoveryViewSettingsSchema V1DiscoveryDefaultPut(ctx, appID, authToken, body)
Update default discovery view

 Required roles:  - can_write_discovery_entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DiscoveryViewSettingsSchema**](DiscoveryViewSettingsSchema.md)|  | 

### Return type

[**DiscoveryViewSettingsSchema**](DiscoveryViewSettingsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchHistoryGet**
> SearchHistoryListSchema V1SearchHistoryGet(ctx, appID, authToken)
Returns the current search history

 Required roles:  - can_read_search_history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**SearchHistoryListSchema**](SearchHistoryListSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchHistorySearchHistoryIdDelete**
> V1SearchHistorySearchHistoryIdDelete(ctx, appID, authToken, searchHistoryId)
Delete a search from history by its id

 Required roles:  - can_delete_search_history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **searchHistoryId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchHistorySearchHistoryIdGet**
> SearchDocumentsSchema V1SearchHistorySearchHistoryIdGet(ctx, appID, authToken, searchHistoryId)
Returns results of search history

 Required roles:  - can_read_search_history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **searchHistoryId** | **string**|  | 

### Return type

[**SearchDocumentsSchema**](SearchDocumentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchPost**
> SearchDocumentsSchema V1SearchPost(ctx, appID, authToken, body, optional)
Search

 Required roles:  - can_search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**SearchCriteriaSchema**](SearchCriteriaSchema.md)|  | 
 **optional** | ***DefaultApiV1SearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1SearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **optional.Int32**| The number of documents for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **scroll** | **optional.Bool**| If true passed then uses scroll pagination instead of default one | 
 **scrollId** | **optional.String**| In order to get next batch of results using scroll pagination the scroll_id is required | 
 **generateSignedUrl** | **optional.Bool**| Set to false if you don&#39;t need a URL, will speed things up | [default to true]
 **generateSignedDownloadUrl** | **optional.Bool**| Set to true if you also want the file download URLs generated | [default to false]
 **saveSearchHistory** | **optional.Bool**| Set to false if you don&#39;t want to save the search to the history | [default to true]

### Return type

[**SearchDocumentsSchema**](SearchDocumentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedGet**
> SavedSearchesSchema V1SearchSavedGet(ctx, appID, authToken, optional)
Returns list of saved searches

 Required roles:  - can_read_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1SearchSavedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1SearchSavedGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - name,asc;group_id,desc | 
 **groupId** | **optional.String**| Group ID | 
 **ids** | **optional.String**| Filter list of id:s (comma separated) | 
 **query** | **optional.String**| Search using query | 

### Return type

[**SavedSearchesSchema**](SavedSearchesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedGroupGroupIdDelete**
> V1SearchSavedGroupGroupIdDelete(ctx, appID, authToken, groupId)
Delete a saved search group by it's id

 Required roles:  - can_delete_saved_search_groups

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

# **V1SearchSavedGroupGroupIdGet**
> SavedSearchGroupSchema V1SearchSavedGroupGroupIdGet(ctx, appID, authToken, groupId)
Returns saved search group data

 Required roles:  - can_read_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

[**SavedSearchGroupSchema**](SavedSearchGroupSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedGroupGroupIdPatch**
> SavedSearchGroupSchema V1SearchSavedGroupGroupIdPatch(ctx, appID, authToken, groupId, body)
Update and return saved search group data

 Required roles:  - can_write_saved_search_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **body** | [**SavedSearchGroupSchema**](SavedSearchGroupSchema.md)|  | 

### Return type

[**SavedSearchGroupSchema**](SavedSearchGroupSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedGroupGroupIdPut**
> SavedSearchGroupSchema V1SearchSavedGroupGroupIdPut(ctx, appID, authToken, groupId, body)
Update and return saved search group data

 Required roles:  - can_write_saved_search_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **body** | [**SavedSearchGroupSchema**](SavedSearchGroupSchema.md)|  | 

### Return type

[**SavedSearchGroupSchema**](SavedSearchGroupSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedGroupGroupIdSearchSearchIdDelete**
> V1SearchSavedGroupGroupIdSearchSearchIdDelete(ctx, appID, authToken, groupId, searchId)
Delete saved search from search group

 Required roles:  - can_write_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **searchId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedGroupGroupIdSearchSearchIdPost**
> V1SearchSavedGroupGroupIdSearchSearchIdPost(ctx, appID, authToken, groupId, searchId)
Adds saved search to group

 Required roles:  - can_write_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **groupId** | **string**|  | 
  **searchId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedGroupPost**
> SavedSearchGroupSchema V1SearchSavedGroupPost(ctx, appID, authToken, body)
Create and return saved search group data

 Required roles:  - can_write_saved_search_groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**SavedSearchGroupSchema**](SavedSearchGroupSchema.md)|  | 

### Return type

[**SavedSearchGroupSchema**](SavedSearchGroupSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedGroupsGet**
> SavedSearchGroupsSchema V1SearchSavedGroupsGet(ctx, appID, authToken, optional)
Returns paginated list of search groups

 Required roles:  - can_read_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1SearchSavedGroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1SearchSavedGroupsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **ids** | **optional.String**| A comma separated list of IDs. for example - 3c2db7d8-1f39-46e3-8c77-992101e5e616,683a2c63-63a0-42b0-aed8-5b62dedba840 | 
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 

### Return type

[**SavedSearchGroupsSchema**](SavedSearchGroupsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedGroupsGroupIdReindexPost**
> V1SearchSavedGroupsGroupIdReindexPost(ctx, appID, authToken, groupId)
Reindex a particular saved search group by id

 Required roles:  - can_reindex_saved_searches

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

# **V1SearchSavedPost**
> SavedSearchSchema V1SearchSavedPost(ctx, appID, authToken, body)
Search, save and return result of this search

 Required roles:  - can_write_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**SavedSearchSchema**](SavedSearchSchema.md)|  | 

### Return type

[**SavedSearchSchema**](SavedSearchSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedSearchIdDelete**
> V1SearchSavedSearchIdDelete(ctx, appID, authToken, searchId)
Delete a saved search by its id

 Required roles:  - can_delete_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **searchId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedSearchIdGet**
> SavedSearchResultsSchema V1SearchSavedSearchIdGet(ctx, appID, authToken, searchId, optional)
Returns results of saved search

 Required roles:  - can_read_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **searchId** | **string**|  | 
 **optional** | ***DefaultApiV1SearchSavedSearchIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1SearchSavedSearchIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **includeResults** | **optional.Bool**| Set to false if you only want the search definition | [default to true]

### Return type

[**SavedSearchResultsSchema**](SavedSearchResultsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedSearchIdPut**
> SearchDocumentsSchema V1SearchSavedSearchIdPut(ctx, appID, authToken, searchId, body)
Search and save this search

 Required roles:  - can_write_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **searchId** | **string**|  | 
  **body** | [**SavedSearchSchema**](SavedSearchSchema.md)|  | 

### Return type

[**SearchDocumentsSchema**](SearchDocumentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSavedSearchIdReindexPost**
> V1SearchSavedSearchIdReindexPost(ctx, appID, authToken, searchId, body)
Reindex a particular saved search by id

 Required roles:  - can_reindex_saved_searches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **searchId** | **string**|  | 
  **body** | [**ReindexSavedSearchSchema**](ReindexSavedSearchSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SearchSuggestPost**
> SearchSuggestsResponseSchema V1SearchSuggestPost(ctx, appID, authToken, optional)
Returns search suggestions for a particular query.

 Required roles:  - can_search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1SearchSuggestPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1SearchSuggestPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of documents for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **filter** | **optional.String**| A comma separated list of fieldnames with order For example - first_name,eq,Vlad;last_name,eq,Gudkov | 
 **body** | [**optional.Interface of SearchSuggestSchema**](SearchSuggestSchema.md)|  | 

### Return type

[**SearchSuggestsResponseSchema**](SearchSuggestsResponseSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


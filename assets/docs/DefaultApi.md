# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/assets*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AssetsAssetIdDelete**](DefaultApi.md#V1AssetsAssetIdDelete) | **Delete** /v1/assets/{asset_id}/ | Delete a particular asset by id
[**V1AssetsAssetIdGet**](DefaultApi.md#V1AssetsAssetIdGet) | **Get** /v1/assets/{asset_id}/ | Returns a particular asset by id
[**V1AssetsAssetIdHistoryGet**](DefaultApi.md#V1AssetsAssetIdHistoryGet) | **Get** /v1/assets/{asset_id}/history/ | Get list of assets
[**V1AssetsAssetIdHistoryHistoryEntityIdDelete**](DefaultApi.md#V1AssetsAssetIdHistoryHistoryEntityIdDelete) | **Delete** /v1/assets/{asset_id}/history/{history_entity_id}/ | Deletes an asset history entity
[**V1AssetsAssetIdHistoryHistoryEntityIdGet**](DefaultApi.md#V1AssetsAssetIdHistoryHistoryEntityIdGet) | **Get** /v1/assets/{asset_id}/history/{history_entity_id}/ | Get an asset history entity
[**V1AssetsAssetIdHistoryHistoryEntityIdReindexPost**](DefaultApi.md#V1AssetsAssetIdHistoryHistoryEntityIdReindexPost) | **Post** /v1/assets/{asset_id}/history/{history_entity_id}/reindex/ | Reindex asset history entity
[**V1AssetsAssetIdHistoryPost**](DefaultApi.md#V1AssetsAssetIdHistoryPost) | **Post** /v1/assets/{asset_id}/history/ | Create an asset history entity
[**V1AssetsAssetIdPatch**](DefaultApi.md#V1AssetsAssetIdPatch) | **Patch** /v1/assets/{asset_id}/ | Update asset
[**V1AssetsAssetIdPurgeDelete**](DefaultApi.md#V1AssetsAssetIdPurgeDelete) | **Delete** /v1/assets/{asset_id}/purge/ | Purges a particular asset by id immediately
[**V1AssetsAssetIdPut**](DefaultApi.md#V1AssetsAssetIdPut) | **Put** /v1/assets/{asset_id}/ | Update asset
[**V1AssetsAssetIdReindexPost**](DefaultApi.md#V1AssetsAssetIdReindexPost) | **Post** /v1/assets/{asset_id}/reindex/ | Reindex asset
[**V1AssetsAssetIdRelationsGet**](DefaultApi.md#V1AssetsAssetIdRelationsGet) | **Get** /v1/assets/{asset_id}/relations/ | Returns an assets relations
[**V1AssetsAssetIdRelationsPost**](DefaultApi.md#V1AssetsAssetIdRelationsPost) | **Post** /v1/assets/{asset_id}/relations/ | Create a new asset relation
[**V1AssetsAssetIdRelationsRelationTypeGet**](DefaultApi.md#V1AssetsAssetIdRelationsRelationTypeGet) | **Get** /v1/assets/{asset_id}/relations/{relation_type}/ | Returns assets that has a relation to this asset
[**V1AssetsAssetIdRelationsRelationTypeRelatedToAssetIdDelete**](DefaultApi.md#V1AssetsAssetIdRelationsRelationTypeRelatedToAssetIdDelete) | **Delete** /v1/assets/{asset_id}/relations/{relation_type}/{related_to_asset_id}/ | Delete a particular asset by id
[**V1AssetsAssetIdRelationsRelationTypeRelatedToAssetIdPost**](DefaultApi.md#V1AssetsAssetIdRelationsRelationTypeRelatedToAssetIdPost) | **Post** /v1/assets/{asset_id}/relations/{relation_type}/{related_to_asset_id}/ | Create a new asset relation
[**V1AssetsAssetIdRestorePut**](DefaultApi.md#V1AssetsAssetIdRestorePut) | **Put** /v1/assets/{asset_id}/restore/ | Restore deleted asset by id
[**V1AssetsAssetIdSearchDocumentPut**](DefaultApi.md#V1AssetsAssetIdSearchDocumentPut) | **Put** /v1/assets/{asset_id}/search_document/ | Update metadata for asset
[**V1AssetsAssetIdSegmentsBulkPost**](DefaultApi.md#V1AssetsAssetIdSegmentsBulkPost) | **Post** /v1/assets/{asset_id}/segments/bulk/ | Create multiple new segments for a single asset
[**V1AssetsAssetIdSegmentsGet**](DefaultApi.md#V1AssetsAssetIdSegmentsGet) | **Get** /v1/assets/{asset_id}/segments/ | List of segments
[**V1AssetsAssetIdSegmentsPost**](DefaultApi.md#V1AssetsAssetIdSegmentsPost) | **Post** /v1/assets/{asset_id}/segments/ | Create a new segment
[**V1AssetsAssetIdSegmentsReindexPost**](DefaultApi.md#V1AssetsAssetIdSegmentsReindexPost) | **Post** /v1/assets/{asset_id}/segments/reindex/ | Reindex assets segments
[**V1AssetsAssetIdSegmentsSegmentIdDelete**](DefaultApi.md#V1AssetsAssetIdSegmentsSegmentIdDelete) | **Delete** /v1/assets/{asset_id}/segments/{segment_id}/ | Delete a particular segment from an asset by id
[**V1AssetsAssetIdSegmentsSegmentIdGet**](DefaultApi.md#V1AssetsAssetIdSegmentsSegmentIdGet) | **Get** /v1/assets/{asset_id}/segments/{segment_id}/ | Get a segment by ID
[**V1AssetsAssetIdSegmentsSegmentIdPatch**](DefaultApi.md#V1AssetsAssetIdSegmentsSegmentIdPatch) | **Patch** /v1/assets/{asset_id}/segments/{segment_id}/ | Update segment
[**V1AssetsAssetIdSegmentsSegmentIdReindexPost**](DefaultApi.md#V1AssetsAssetIdSegmentsSegmentIdReindexPost) | **Post** /v1/assets/{asset_id}/segments/{segment_id}/reindex/ | Reindex assets segment
[**V1AssetsAssetIdSegmentsSegmentTypeGet**](DefaultApi.md#V1AssetsAssetIdSegmentsSegmentTypeGet) | **Get** /v1/assets/{asset_id}/segments/{segment_type}/ | List of segments
[**V1AssetsAssetIdSegmentsSrtGet**](DefaultApi.md#V1AssetsAssetIdSegmentsSrtGet) | **Get** /v1/assets/{asset_id}/segments/srt/ | List of segments as SRT file
[**V1AssetsAssetIdSegmentsVttGet**](DefaultApi.md#V1AssetsAssetIdSegmentsVttGet) | **Get** /v1/assets/{asset_id}/segments/vtt/ | List of segments as WebVTT file
[**V1AssetsAssetIdUploadsDelete**](DefaultApi.md#V1AssetsAssetIdUploadsDelete) | **Delete** /v1/assets/{asset_id}/uploads/ | Delete a particular asset by id on failed uplaod
[**V1AssetsAssetIdVersionsFromAssetsSourceAssetIdPost**](DefaultApi.md#V1AssetsAssetIdVersionsFromAssetsSourceAssetIdPost) | **Post** /v1/assets/{asset_id}/versions/from/assets/{source_asset_id}/ | Create a new asset&#39;s version from another asset
[**V1AssetsAssetIdVersionsOldDelete**](DefaultApi.md#V1AssetsAssetIdVersionsOldDelete) | **Delete** /v1/assets/{asset_id}/versions/old/ | Delete all asset versions except the latest one
[**V1AssetsAssetIdVersionsPost**](DefaultApi.md#V1AssetsAssetIdVersionsPost) | **Post** /v1/assets/{asset_id}/versions/ | Add asset version
[**V1AssetsAssetIdVersionsVersionIdDelete**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdDelete) | **Delete** /v1/assets/{asset_id}/versions/{version_id}/ | Delete a particular asset version by id
[**V1AssetsAssetIdVersionsVersionIdPatch**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdPatch) | **Patch** /v1/assets/{asset_id}/versions/{version_id}/ | Edit asset version
[**V1AssetsAssetIdVersionsVersionIdPromotePut**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdPromotePut) | **Put** /v1/assets/{asset_id}/versions/{version_id}/promote/ | Promote a particular asset version to a latest version
[**V1AssetsAssetIdVersionsVersionIdPut**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdPut) | **Put** /v1/assets/{asset_id}/versions/{version_id}/ | Edit asset version
[**V1AssetsAssetIdVersionsVersionIdTranscriptionsPropertiesGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdTranscriptionsPropertiesGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/transcriptions/properties/ | Get a list of transcription properties
[**V1AssetsAssetIdVersionsVersionIdTranscriptionsPropertiesPost**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdTranscriptionsPropertiesPost) | **Post** /v1/assets/{asset_id}/versions/{version_id}/transcriptions/properties/ | Add a new transcription properties
[**V1AssetsAssetIdVersionsVersionIdTranscriptionsSubtitlesPost**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdTranscriptionsSubtitlesPost) | **Post** /v1/assets/{asset_id}/versions/{version_id}/transcriptions/subtitles/ | Add a new transcription properties
[**V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesDelete**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesDelete) | **Delete** /v1/assets/{asset_id}/versions/{version_id}/transcriptions/{transcription_id}/properties/ | Delete transcription properties by ID
[**V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesGet**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesGet) | **Get** /v1/assets/{asset_id}/versions/{version_id}/transcriptions/{transcription_id}/properties/ | Get a transcription properties by ID
[**V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesPatch**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesPatch) | **Patch** /v1/assets/{asset_id}/versions/{version_id}/transcriptions/{transcription_id}/properties/ | Update transcription properties by ID
[**V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesPut**](DefaultApi.md#V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesPut) | **Put** /v1/assets/{asset_id}/versions/{version_id}/transcriptions/{transcription_id}/properties/ | Update transcription properties by ID
[**V1AssetsGet**](DefaultApi.md#V1AssetsGet) | **Get** /v1/assets/ | Get list of assets
[**V1AssetsPatch**](DefaultApi.md#V1AssetsPatch) | **Patch** /v1/assets/ | Bulk update assets
[**V1AssetsPost**](DefaultApi.md#V1AssetsPost) | **Post** /v1/assets/ | Create a new asset
[**V1AssetsPut**](DefaultApi.md#V1AssetsPut) | **Put** /v1/assets/ | Bulk update assets
[**V1AssetsReindexPost**](DefaultApi.md#V1AssetsReindexPost) | **Post** /v1/assets/reindex/ | Trigger reindexing of all assets
[**V1AssetsRelationTypesGet**](DefaultApi.md#V1AssetsRelationTypesGet) | **Get** /v1/assets/relation_types/ | Create a new asset relation type
[**V1AssetsRelationTypesPost**](DefaultApi.md#V1AssetsRelationTypesPost) | **Post** /v1/assets/relation_types/ | Create a new asset relation type
[**V1AssetsRelationTypesRelationTypeDelete**](DefaultApi.md#V1AssetsRelationTypesRelationTypeDelete) | **Delete** /v1/assets/relation_types/{relation_type}/ | Delete an asset relation type
[**V1AssetsRelationTypesRelationTypeGet**](DefaultApi.md#V1AssetsRelationTypesRelationTypeGet) | **Get** /v1/assets/relation_types/{relation_type}/ | Get a relation type
[**V1AssetsRelationTypesRelationTypePatch**](DefaultApi.md#V1AssetsRelationTypesRelationTypePatch) | **Patch** /v1/assets/relation_types/{relation_type}/ | Update an asset relation type
[**V1AssetsRelationTypesRelationTypePut**](DefaultApi.md#V1AssetsRelationTypesRelationTypePut) | **Put** /v1/assets/relation_types/{relation_type}/ | Update an asset relation type
[**V1AssetsSegmentsReindexPost**](DefaultApi.md#V1AssetsSegmentsReindexPost) | **Post** /v1/assets/segments/reindex/ | Trigger reindexing of all segments
[**V1CollectionsCollectionIdAncestorsGet**](DefaultApi.md#V1CollectionsCollectionIdAncestorsGet) | **Get** /v1/collections/{collection_id}/ancestors/ | Returns list of ancestors of a collection
[**V1CollectionsCollectionIdContentInfoGet**](DefaultApi.md#V1CollectionsCollectionIdContentInfoGet) | **Get** /v1/collections/{collection_id}/content/info/ | Returns a sub-collections and assets number for a specific collection
[**V1CollectionsCollectionIdContentsGet**](DefaultApi.md#V1CollectionsCollectionIdContentsGet) | **Get** /v1/collections/{collection_id}/contents/ | Returns contents of a collection by id
[**V1CollectionsCollectionIdContentsObjectTypeObjectIdDelete**](DefaultApi.md#V1CollectionsCollectionIdContentsObjectTypeObjectIdDelete) | **Delete** /v1/collections/{collection_id}/contents/{object_type}/{object_id}/ | Delete a particular content object in a collection by id
[**V1CollectionsCollectionIdContentsObjectTypeObjectIdPut**](DefaultApi.md#V1CollectionsCollectionIdContentsObjectTypeObjectIdPut) | **Put** /v1/collections/{collection_id}/contents/{object_type}/{object_id}/ | Update an order of a particular content object in a collection
[**V1CollectionsCollectionIdContentsObjectTypeObjectIdReindexPost**](DefaultApi.md#V1CollectionsCollectionIdContentsObjectTypeObjectIdReindexPost) | **Post** /v1/collections/{collection_id}/contents/{object_type}/{object_id}/reindex/ | Reindex collection content
[**V1CollectionsCollectionIdContentsOrderingCustomDelete**](DefaultApi.md#V1CollectionsCollectionIdContentsOrderingCustomDelete) | **Delete** /v1/collections/{collection_id}/contents/ordering/custom/ | Disable custom ordering for a collection&#39;s content
[**V1CollectionsCollectionIdContentsOrderingCustomPost**](DefaultApi.md#V1CollectionsCollectionIdContentsOrderingCustomPost) | **Post** /v1/collections/{collection_id}/contents/ordering/custom/ | Enable custom ordering for a collection&#39;s content
[**V1CollectionsCollectionIdContentsPost**](DefaultApi.md#V1CollectionsCollectionIdContentsPost) | **Post** /v1/collections/{collection_id}/contents/ | Add an object to a collection
[**V1CollectionsCollectionIdDelete**](DefaultApi.md#V1CollectionsCollectionIdDelete) | **Delete** /v1/collections/{collection_id}/ | Delete a particular collection by id
[**V1CollectionsCollectionIdGet**](DefaultApi.md#V1CollectionsCollectionIdGet) | **Get** /v1/collections/{collection_id}/ | Returns a particular collection by id
[**V1CollectionsCollectionIdKeyframesPost**](DefaultApi.md#V1CollectionsCollectionIdKeyframesPost) | **Post** /v1/collections/{collection_id}/keyframes/ | Pick up to three asset_ids for collection keyframes
[**V1CollectionsCollectionIdPatch**](DefaultApi.md#V1CollectionsCollectionIdPatch) | **Patch** /v1/collections/{collection_id}/ | Update collection
[**V1CollectionsCollectionIdPurgeDelete**](DefaultApi.md#V1CollectionsCollectionIdPurgeDelete) | **Delete** /v1/collections/{collection_id}/purge/ | Purges deleted collection by id immediately
[**V1CollectionsCollectionIdPut**](DefaultApi.md#V1CollectionsCollectionIdPut) | **Put** /v1/collections/{collection_id}/ | Update collection
[**V1CollectionsCollectionIdReindexContentsPost**](DefaultApi.md#V1CollectionsCollectionIdReindexContentsPost) | **Post** /v1/collections/{collection_id}/reindex/contents/ | Reindex collection and its content
[**V1CollectionsCollectionIdReindexPost**](DefaultApi.md#V1CollectionsCollectionIdReindexPost) | **Post** /v1/collections/{collection_id}/reindex/ | Reindex collection
[**V1CollectionsCollectionIdRestorePut**](DefaultApi.md#V1CollectionsCollectionIdRestorePut) | **Put** /v1/collections/{collection_id}/restore/ | Restore deleted collection by id
[**V1CollectionsCollectionIdSearchDocumentPut**](DefaultApi.md#V1CollectionsCollectionIdSearchDocumentPut) | **Put** /v1/collections/{collection_id}/search_document/ | Update metadata for collection
[**V1CollectionsCollectionIdSubcollectionsPost**](DefaultApi.md#V1CollectionsCollectionIdSubcollectionsPost) | **Post** /v1/collections/{collection_id}/subcollections/ | Copy a collection (recursively) in to another collection
[**V1CollectionsGet**](DefaultApi.md#V1CollectionsGet) | **Get** /v1/collections/ | Get list of collections
[**V1CollectionsPost**](DefaultApi.md#V1CollectionsPost) | **Post** /v1/collections/ | Create a new collection
[**V1CollectionsReindexPost**](DefaultApi.md#V1CollectionsReindexPost) | **Post** /v1/collections/reindex/ | Trigger reindexing of all collections
[**V1CustomActionsContextActionIdCallbackPost**](DefaultApi.md#V1CustomActionsContextActionIdCallbackPost) | **Post** /v1/custom_actions/{context}/{action_id}/callback/ | Schedules a celery task that will call custom action
[**V1CustomActionsContextActionIdDelete**](DefaultApi.md#V1CustomActionsContextActionIdDelete) | **Delete** /v1/custom_actions/{context}/{action_id}/ | Deletes an custom action
[**V1CustomActionsContextActionIdGet**](DefaultApi.md#V1CustomActionsContextActionIdGet) | **Get** /v1/custom_actions/{context}/{action_id}/ | Get an asset custom action
[**V1CustomActionsContextActionIdPatch**](DefaultApi.md#V1CustomActionsContextActionIdPatch) | **Patch** /v1/custom_actions/{context}/{action_id}/ | Update an custom action
[**V1CustomActionsContextActionIdPut**](DefaultApi.md#V1CustomActionsContextActionIdPut) | **Put** /v1/custom_actions/{context}/{action_id}/ | Update an custom action
[**V1CustomActionsContextGet**](DefaultApi.md#V1CustomActionsContextGet) | **Get** /v1/custom_actions/{context}/ | Get list of custom actions by context
[**V1CustomActionsContextPost**](DefaultApi.md#V1CustomActionsContextPost) | **Post** /v1/custom_actions/{context}/ | Create an custom action
[**V1CustomActionsGet**](DefaultApi.md#V1CustomActionsGet) | **Get** /v1/custom_actions/ | Get list of custom actions
[**V1CustomActionsSharedContextActionIdCallbackPost**](DefaultApi.md#V1CustomActionsSharedContextActionIdCallbackPost) | **Post** /v1/custom_actions/shared/{context}/{action_id}/callback/ | Schedules a celery task that will call custom action on shares
[**V1DeleteQueueAssetsDelete**](DefaultApi.md#V1DeleteQueueAssetsDelete) | **Delete** /v1/delete_queue/assets/ | Delete assets from delete queue (Mark assets as active again)
[**V1DeleteQueueAssetsGet**](DefaultApi.md#V1DeleteQueueAssetsGet) | **Get** /v1/delete_queue/assets/ | Get deleted objects
[**V1DeleteQueueAssetsPost**](DefaultApi.md#V1DeleteQueueAssetsPost) | **Post** /v1/delete_queue/assets/ | Add assets to a delete queue (Mark assets as deleted)
[**V1DeleteQueueAssetsPurgeAllPost**](DefaultApi.md#V1DeleteQueueAssetsPurgeAllPost) | **Post** /v1/delete_queue/assets/purge/all/ | Purge all assets from delete queue (Permanently delete)
[**V1DeleteQueueAssetsPurgePost**](DefaultApi.md#V1DeleteQueueAssetsPurgePost) | **Post** /v1/delete_queue/assets/purge/ | Purge assets from delete queue (Permanently delete)
[**V1DeleteQueueBulkPost**](DefaultApi.md#V1DeleteQueueBulkPost) | **Post** /v1/delete_queue/bulk/ | Bulk delete objects
[**V1DeleteQueueCollectionsDelete**](DefaultApi.md#V1DeleteQueueCollectionsDelete) | **Delete** /v1/delete_queue/collections/ | Delete collections from delete queue (Mark collections as active again)
[**V1DeleteQueueCollectionsGet**](DefaultApi.md#V1DeleteQueueCollectionsGet) | **Get** /v1/delete_queue/collections/ | Get list of collections
[**V1DeleteQueueCollectionsPost**](DefaultApi.md#V1DeleteQueueCollectionsPost) | **Post** /v1/delete_queue/collections/ | Add collections to a delete queue (Mark collections as deleted)
[**V1DeleteQueueCollectionsPurgeAllPost**](DefaultApi.md#V1DeleteQueueCollectionsPurgeAllPost) | **Post** /v1/delete_queue/collections/purge/all/ | Purge all collections from delete queue (Permanently delete)
[**V1DeleteQueueCollectionsPurgePost**](DefaultApi.md#V1DeleteQueueCollectionsPurgePost) | **Post** /v1/delete_queue/collections/purge/ | Purge collections from delete queue (Permanently delete)
[**V1DeleteQueuePurgeAllPost**](DefaultApi.md#V1DeleteQueuePurgeAllPost) | **Post** /v1/delete_queue/purge/all/ | Purge all assets and collections from delete queue (Permanently delete)
[**V1ObjectTypeObjectIdApprovalsDelete**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsDelete) | **Delete** /v1/{object_type}/{object_id}/approvals/ | Deletes an objects approval status
[**V1ObjectTypeObjectIdApprovalsExternalEmailDelete**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsExternalEmailDelete) | **Delete** /v1/{object_type}/{object_id}/approvals/external/{email}/ | Deletes an objects approval status by user_id
[**V1ObjectTypeObjectIdApprovalsGet**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsGet) | **Get** /v1/{object_type}/{object_id}/approvals/ | Returns an objects approval request
[**V1ObjectTypeObjectIdApprovalsPut**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsPut) | **Put** /v1/{object_type}/{object_id}/approvals/ | Returns an objects approval status
[**V1ObjectTypeObjectIdApprovalsRequestDelete**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsRequestDelete) | **Delete** /v1/{object_type}/{object_id}/approvals/request/ | Deletes an objects approval request
[**V1ObjectTypeObjectIdApprovalsRequestGet**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsRequestGet) | **Get** /v1/{object_type}/{object_id}/approvals/request/ | Returns an objects approval request
[**V1ObjectTypeObjectIdApprovalsRequestPatch**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsRequestPatch) | **Patch** /v1/{object_type}/{object_id}/approvals/request/ | Edits an approval request
[**V1ObjectTypeObjectIdApprovalsRequestPost**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsRequestPost) | **Post** /v1/{object_type}/{object_id}/approvals/request/ | Creates an objects approval request
[**V1ObjectTypeObjectIdApprovalsRequestPut**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsRequestPut) | **Put** /v1/{object_type}/{object_id}/approvals/request/ | Edits an approval request
[**V1ObjectTypeObjectIdApprovalsUserUserIdDelete**](DefaultApi.md#V1ObjectTypeObjectIdApprovalsUserUserIdDelete) | **Delete** /v1/{object_type}/{object_id}/approvals/user/{user_id}/ | Deletes an objects approval status by user_id
[**V1ObjectTypeObjectIdSharesGet**](DefaultApi.md#V1ObjectTypeObjectIdSharesGet) | **Get** /v1/{object_type}/{object_id}/shares/ | Get list of shares
[**V1ObjectTypeObjectIdSharesPost**](DefaultApi.md#V1ObjectTypeObjectIdSharesPost) | **Post** /v1/{object_type}/{object_id}/shares/ | Create a new share
[**V1ObjectTypeObjectIdSharesShareIdDelete**](DefaultApi.md#V1ObjectTypeObjectIdSharesShareIdDelete) | **Delete** /v1/{object_type}/{object_id}/shares/{share_id}/ | Delete a particular share by id
[**V1ObjectTypeObjectIdSharesShareIdGet**](DefaultApi.md#V1ObjectTypeObjectIdSharesShareIdGet) | **Get** /v1/{object_type}/{object_id}/shares/{share_id}/ | Returns a particular share by id
[**V1ObjectTypeObjectIdSharesShareIdPut**](DefaultApi.md#V1ObjectTypeObjectIdSharesShareIdPut) | **Put** /v1/{object_type}/{object_id}/shares/{share_id}/ | Update share
[**V1ObjectTypeObjectIdSharesShareIdUsersGet**](DefaultApi.md#V1ObjectTypeObjectIdSharesShareIdUsersGet) | **Get** /v1/{object_type}/{object_id}/shares/{share_id}/users/ | Get list of share users
[**V1ObjectTypeObjectIdSharesShareIdUsersPost**](DefaultApi.md#V1ObjectTypeObjectIdSharesShareIdUsersPost) | **Post** /v1/{object_type}/{object_id}/shares/{share_id}/users/ | Add a new share_user to a share
[**V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdDelete**](DefaultApi.md#V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdDelete) | **Delete** /v1/{object_type}/{object_id}/shares/{share_id}/users/{share_user_id}/ | Delete a particular share_user user by id
[**V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdGet**](DefaultApi.md#V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdGet) | **Get** /v1/{object_type}/{object_id}/shares/{share_id}/users/{share_user_id}/ | Returns a particular share user by id
[**V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdPatch**](DefaultApi.md#V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdPatch) | **Patch** /v1/{object_type}/{object_id}/shares/{share_id}/users/{share_user_id}/ | Update share user
[**V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdPut**](DefaultApi.md#V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdPut) | **Put** /v1/{object_type}/{object_id}/shares/{share_id}/users/{share_user_id}/ | Update share user
[**V1ObjectTypeObjectIdSharesUrlPost**](DefaultApi.md#V1ObjectTypeObjectIdSharesUrlPost) | **Post** /v1/{object_type}/{object_id}/shares/url/ | Generates a URL for the shared object
[**V1ObjectTypeObjectIdVersionsVersionIdApprovalsGet**](DefaultApi.md#V1ObjectTypeObjectIdVersionsVersionIdApprovalsGet) | **Get** /v1/{object_type}/{object_id}/versions/{version_id}/approvals/ | Returns an objects approval request by version
[**V1ObjectTypeObjectIdVersionsVersionIdApprovalsRequestGet**](DefaultApi.md#V1ObjectTypeObjectIdVersionsVersionIdApprovalsRequestGet) | **Get** /v1/{object_type}/{object_id}/versions/{version_id}/approvals/request/ | Returns an objects approval request by version
[**V1ShareObjectTypePost**](DefaultApi.md#V1ShareObjectTypePost) | **Post** /v1/share/{object_type}/ | Create a new share of multiple objects (currently only assets are supported)
[**V1SharesAuthLoginPost**](DefaultApi.md#V1SharesAuthLoginPost) | **Post** /v1/shares/auth/login/ | Login for share
[**V1SharesAuthTokenGet**](DefaultApi.md#V1SharesAuthTokenGet) | **Get** /v1/shares/auth/token/ | Check if a token is valid


# **V1AssetsAssetIdDelete**
> V1AssetsAssetIdDelete(ctx, appID, authToken, assetId)
Delete a particular asset by id

 Required roles:  - can_delete_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdGet**
> AssetSchema V1AssetsAssetIdGet(ctx, appID, authToken, assetId, optional)
Returns a particular asset by id

 Required roles:  - can_read_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeCollections** | **optional.Bool**| include collection membership | [default to false]
 **includeUsers** | **optional.Bool**| include info about the users who have interacted with this asset | [default to false]

### Return type

[**AssetSchema**](AssetSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdHistoryGet**
> AssetHistoryEntitiesSchema V1AssetsAssetIdHistoryGet(ctx, appID, authToken, assetId, optional)
Get list of assets

 Required roles:  - can_read_assets_history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdHistoryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdHistoryGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **filter** | **optional.String**| A comma separated list of fieldnames with order For example - first_name,eq,Vlad;last_name,eq,Gudkov | 

### Return type

[**AssetHistoryEntitiesSchema**](AssetHistoryEntitiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdHistoryHistoryEntityIdDelete**
> AssetHistorySchema V1AssetsAssetIdHistoryHistoryEntityIdDelete(ctx, appID, authToken, assetId, historyEntityId)
Deletes an asset history entity

 Required roles:  - can_delete_assets_history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **historyEntityId** | **string**|  | 

### Return type

[**AssetHistorySchema**](AssetHistorySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdHistoryHistoryEntityIdGet**
> AssetHistorySchema V1AssetsAssetIdHistoryHistoryEntityIdGet(ctx, appID, authToken, assetId, historyEntityId)
Get an asset history entity

 Required roles:  - can_read_assets_history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **historyEntityId** | **string**|  | 

### Return type

[**AssetHistorySchema**](AssetHistorySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdHistoryHistoryEntityIdReindexPost**
> V1AssetsAssetIdHistoryHistoryEntityIdReindexPost(ctx, appID, authToken, assetId, historyEntityId)
Reindex asset history entity

 Required roles:  - can_reindex_assets_history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **historyEntityId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdHistoryPost**
> AssetHistorySchema V1AssetsAssetIdHistoryPost(ctx, appID, authToken, assetId, body)
Create an asset history entity

 Required roles:  - can_write_assets_history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**AssetHistorySchema**](AssetHistorySchema.md)|  | 

### Return type

[**AssetHistorySchema**](AssetHistorySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdPatch**
> AssetSchema V1AssetsAssetIdPatch(ctx, appID, authToken, assetId, body)
Update asset

 Required roles:  - can_write_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**AssetSchema**](AssetSchema.md)|  | 

### Return type

[**AssetSchema**](AssetSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdPurgeDelete**
> V1AssetsAssetIdPurgeDelete(ctx, appID, authToken, assetId)
Purges a particular asset by id immediately

 Required roles:  - can_purge_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdPut**
> AssetSchema V1AssetsAssetIdPut(ctx, appID, authToken, assetId, body)
Update asset

 Required roles:  - can_write_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**AssetSchema**](AssetSchema.md)|  | 

### Return type

[**AssetSchema**](AssetSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdReindexPost**
> V1AssetsAssetIdReindexPost(ctx, appID, authToken, assetId, body)
Reindex asset

 Required roles:  - can_reindex_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**ReindexAssetSchema**](ReindexAssetSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdRelationsGet**
> AssetsSchema V1AssetsAssetIdRelationsGet(ctx, appID, authToken, assetId, optional)
Returns an assets relations

 Required roles:  - can_read_asset_relations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdRelationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdRelationsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeDeleted** | **optional.Bool**| Also show assets from recycle bin in relations | [default to false]
 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 

### Return type

[**AssetsSchema**](AssetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdRelationsPost**
> RelationSchema V1AssetsAssetIdRelationsPost(ctx, appID, authToken, assetId, body)
Create a new asset relation

 Required roles:  - can_create_asset_relations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**RelationSchema**](RelationSchema.md)|  | 

### Return type

[**RelationSchema**](RelationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdRelationsRelationTypeGet**
> AssetsSchema V1AssetsAssetIdRelationsRelationTypeGet(ctx, appID, authToken, assetId, relationType, optional)
Returns assets that has a relation to this asset

 Required roles:  - can_read_asset_relations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **relationType** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdRelationsRelationTypeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdRelationsRelationTypeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeDeleted** | **optional.Bool**| Also show assets from recycle bin in relations | [default to false]

### Return type

[**AssetsSchema**](AssetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdRelationsRelationTypeRelatedToAssetIdDelete**
> V1AssetsAssetIdRelationsRelationTypeRelatedToAssetIdDelete(ctx, appID, authToken, assetId, relationType, relatedToAssetId)
Delete a particular asset by id

 Required roles:  - can_delete_asset_relations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **relationType** | **string**|  | 
  **relatedToAssetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdRelationsRelationTypeRelatedToAssetIdPost**
> RelationSchema V1AssetsAssetIdRelationsRelationTypeRelatedToAssetIdPost(ctx, appID, authToken, assetId, relationType, relatedToAssetId, body)
Create a new asset relation

 Required roles:  - can_create_asset_relations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **relationType** | **string**|  | 
  **relatedToAssetId** | **string**|  | 
  **body** | [**RelationSchema**](RelationSchema.md)|  | 

### Return type

[**RelationSchema**](RelationSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdRestorePut**
> V1AssetsAssetIdRestorePut(ctx, appID, authToken, assetId)
Restore deleted asset by id

 Required roles:  - can_write_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSearchDocumentPut**
> V1AssetsAssetIdSearchDocumentPut(ctx, appID, authToken, assetId, body)
Update metadata for asset

 Required roles:  - can_reindex_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**AssetElasticSchema**](AssetElasticSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsBulkPost**
> V1AssetsAssetIdSegmentsBulkPost(ctx, appID, authToken, assetId, body)
Create multiple new segments for a single asset

 Required roles:  - can_create_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**BulkCreateSegmentsSchema**](BulkCreateSegmentsSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsGet**
> SegmentsSchema V1AssetsAssetIdSegmentsGet(ctx, appID, authToken, assetId, optional)
List of segments

 Required roles:  - can_read_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdSegmentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdSegmentsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sort** | **optional.String**| Sort on field (Comma separated) | 
 **ids** | **optional.String**| Only include these segments (Comma separated) | 
 **query** | **optional.String**| Search using query | 
 **includes** | **optional.String**| Only include fields (Comma separated) | 
 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **scroll** | **optional.Bool**| If true passed then uses scroll pagination instead of default one | 
 **scrollId** | **optional.String**| In order to get next batch of results using scroll pagination the scroll_id is required | 
 **transcriptionId** | **optional.String**| Filter segments by transcription_id | 
 **versionId** | **optional.String**| Filter segments by version_id | 
 **segmentType** | **optional.String**| Filter segments by segment_type | 
 **segmentColor** | **optional.String**| Filter segments by segment_color | 
 **status** | **optional.String**| Filter segments by status | 
 **includeUsers** | **optional.Bool**| Include segment&#39;s authors info | 

### Return type

[**SegmentsSchema**](SegmentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsPost**
> SegmentSchema V1AssetsAssetIdSegmentsPost(ctx, appID, authToken, assetId, body, optional)
Create a new segment

 Required roles:  - can_create_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**SegmentSchema**](SegmentSchema.md)|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdSegmentsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdSegmentsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **shareUserEmail** | **optional.String**| This header is used for shares by URL to identify user. Only valid emails are allowed. | 

### Return type

[**SegmentSchema**](SegmentSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsReindexPost**
> V1AssetsAssetIdSegmentsReindexPost(ctx, appID, authToken, assetId, body)
Reindex assets segments

 Required roles:  - can_reindex_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**ReindexSegmentSchema**](ReindexSegmentSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsSegmentIdDelete**
> V1AssetsAssetIdSegmentsSegmentIdDelete(ctx, appID, authToken, assetId, segmentId, optional)
Delete a particular segment from an asset by id

 Required roles:  - can_delete_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **segmentId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdSegmentsSegmentIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdSegmentsSegmentIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **softDelete** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsSegmentIdGet**
> SegmentSchema V1AssetsAssetIdSegmentsSegmentIdGet(ctx, appID, authToken, assetId, segmentId, optional)
Get a segment by ID

 Required roles:  - can_read_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **segmentId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdSegmentsSegmentIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdSegmentsSegmentIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **scroll** | **optional.Bool**| If true passed then uses scroll pagination instead of default one | 
 **scrollId** | **optional.String**| In order to get next batch of results using scroll pagination the scroll_id is required | 

### Return type

[**SegmentSchema**](SegmentSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsSegmentIdPatch**
> SegmentSchema V1AssetsAssetIdSegmentsSegmentIdPatch(ctx, appID, authToken, assetId, segmentId, body)
Update segment

 Required roles:  - can_write_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **segmentId** | **string**|  | 
  **body** | [**EditSegmentSchema**](EditSegmentSchema.md)|  | 

### Return type

[**SegmentSchema**](SegmentSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsSegmentIdReindexPost**
> V1AssetsAssetIdSegmentsSegmentIdReindexPost(ctx, appID, authToken, assetId, segmentId, body)
Reindex assets segment

 Required roles:  - can_reindex_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **segmentId** | **string**|  | 
  **body** | [**ReindexSegmentSchema**](ReindexSegmentSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsSegmentTypeGet**
> SegmentsSchema V1AssetsAssetIdSegmentsSegmentTypeGet(ctx, appID, authToken, assetId, segmentType, optional)
List of segments

 Required roles:  - can_read_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **segmentType** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdSegmentsSegmentTypeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdSegmentsSegmentTypeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sort** | **optional.String**| Sort on field (Comma separated) | 
 **ids** | **optional.String**| Only include these segments (Comma separated) | 
 **query** | **optional.String**| Search using query | 
 **includes** | **optional.String**| Only include fields (Comma separated) | 
 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **scroll** | **optional.Bool**| If true passed then uses scroll pagination instead of default one | 
 **scrollId** | **optional.String**| In order to get next batch of results using scroll pagination the scroll_id is required | 

### Return type

[**SegmentsSchema**](SegmentsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsSrtGet**
> V1AssetsAssetIdSegmentsSrtGet(ctx, appID, authToken, assetId, optional)
List of segments as SRT file

 Required roles:  - can_read_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdSegmentsSrtGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdSegmentsSrtGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ids** | **optional.String**| Only include these segments (Comma separated) | 
 **query** | **optional.String**| Search using query | 
 **scroll** | **optional.Bool**| If true passed then uses scroll pagination instead of default one | 
 **scrollId** | **optional.String**| In order to get next batch of results using scroll pagination the scroll_id is required | 
 **transcriptionId** | **optional.String**| Filter segments by transcription_id | 
 **versionId** | **optional.String**| Filter segments by version_id | 
 **segmentType** | **optional.String**| Filter segments by segment_type | 
 **segmentColor** | **optional.String**| Filter segments by segment_color | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdSegmentsVttGet**
> V1AssetsAssetIdSegmentsVttGet(ctx, appID, authToken, assetId, optional)
List of segments as WebVTT file

 Required roles:  - can_read_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsAssetIdSegmentsVttGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsAssetIdSegmentsVttGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ids** | **optional.String**| Only include these segments (Comma separated) | 
 **query** | **optional.String**| Search using query | 
 **scroll** | **optional.Bool**| If true passed then uses scroll pagination instead of default one | 
 **scrollId** | **optional.String**| In order to get next batch of results using scroll pagination the scroll_id is required | 
 **transcriptionId** | **optional.String**| Filter segments by transcription_id | 
 **versionId** | **optional.String**| Filter segments by version_id | 
 **segmentType** | **optional.String**| Filter segments by segment_type | 
 **segmentColor** | **optional.String**| Filter segments by segment_color | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdUploadsDelete**
> V1AssetsAssetIdUploadsDelete(ctx, appID, authToken, assetId)
Delete a particular asset by id on failed uplaod

 Required roles:  - can_create_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsFromAssetsSourceAssetIdPost**
> V1AssetsAssetIdVersionsFromAssetsSourceAssetIdPost(ctx, appID, authToken, assetId, sourceAssetId, body)
Create a new asset's version from another asset

 Required roles:  - can_write_versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **sourceAssetId** | **string**|  | 
  **body** | [**CreateAssetVersionFromAssetSchema**](CreateAssetVersionFromAssetSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsOldDelete**
> V1AssetsAssetIdVersionsOldDelete(ctx, appID, authToken, assetId)
Delete all asset versions except the latest one

 Required roles:  - can_delete_versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsPost**
> AssetVersionsSchema V1AssetsAssetIdVersionsPost(ctx, appID, authToken, assetId, body)
Add asset version

 Required roles:  - can_write_versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **body** | [**CreateAssetVersionSchema**](CreateAssetVersionSchema.md)|  | 

### Return type

[**AssetVersionsSchema**](AssetVersionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdDelete**
> V1AssetsAssetIdVersionsVersionIdDelete(ctx, appID, authToken, assetId, versionId)
Delete a particular asset version by id

 Required roles:  - can_delete_versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdPatch**
> AssetVersionSchema V1AssetsAssetIdVersionsVersionIdPatch(ctx, appID, authToken, assetId, versionId, body)
Edit asset version

 Required roles:  - can_write_versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **body** | [**AssetVersionSchema**](AssetVersionSchema.md)|  | 

### Return type

[**AssetVersionSchema**](AssetVersionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdPromotePut**
> V1AssetsAssetIdVersionsVersionIdPromotePut(ctx, appID, authToken, assetId, versionId)
Promote a particular asset version to a latest version

 Required roles:  - can_write_versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdPut**
> AssetVersionSchema V1AssetsAssetIdVersionsVersionIdPut(ctx, appID, authToken, assetId, versionId, body)
Edit asset version

 Required roles:  - can_write_versions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **body** | [**AssetVersionSchema**](AssetVersionSchema.md)|  | 

### Return type

[**AssetVersionSchema**](AssetVersionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdTranscriptionsPropertiesGet**
> AssetTranscriptionsPropertiesSchema V1AssetsAssetIdVersionsVersionIdTranscriptionsPropertiesGet(ctx, appID, authToken, assetId, versionId)
Get a list of transcription properties

 Required roles:  - can_read_transcriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 

### Return type

[**AssetTranscriptionsPropertiesSchema**](AssetTranscriptionsPropertiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdTranscriptionsPropertiesPost**
> AssetTranscriptionPropertiesSchema V1AssetsAssetIdVersionsVersionIdTranscriptionsPropertiesPost(ctx, appID, authToken, assetId, versionId, body)
Add a new transcription properties

 Required roles:  - can_write_transcriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **body** | [**AssetTranscriptionPropertiesSchema**](AssetTranscriptionPropertiesSchema.md)|  | 

### Return type

[**AssetTranscriptionPropertiesSchema**](AssetTranscriptionPropertiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdTranscriptionsSubtitlesPost**
> AssetTranscriptionPropertiesSchema V1AssetsAssetIdVersionsVersionIdTranscriptionsSubtitlesPost(ctx, appID, authToken, assetId, versionId, body)
Add a new transcription properties

 Required roles:  - can_write_transcriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **body** | [**AssetTranscriptionFromSubtitleSchema**](AssetTranscriptionFromSubtitleSchema.md)|  | 

### Return type

[**AssetTranscriptionPropertiesSchema**](AssetTranscriptionPropertiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesDelete**
> V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesDelete(ctx, appID, authToken, assetId, versionId, transcriptionId)
Delete transcription properties by ID

 Required roles:  - can_delete_transcriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **transcriptionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesGet**
> AssetTranscriptionPropertiesSchema V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesGet(ctx, appID, authToken, assetId, versionId, transcriptionId)
Get a transcription properties by ID

 Required roles:  - can_read_transcriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **transcriptionId** | **string**|  | 

### Return type

[**AssetTranscriptionPropertiesSchema**](AssetTranscriptionPropertiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesPatch**
> AssetTranscriptionPropertiesSchema V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesPatch(ctx, appID, authToken, assetId, versionId, transcriptionId, body)
Update transcription properties by ID

 Required roles:  - can_write_transcriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **transcriptionId** | **string**|  | 
  **body** | [**AssetTranscriptionPropertiesSchema**](AssetTranscriptionPropertiesSchema.md)|  | 

### Return type

[**AssetTranscriptionPropertiesSchema**](AssetTranscriptionPropertiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesPut**
> AssetTranscriptionPropertiesSchema V1AssetsAssetIdVersionsVersionIdTranscriptionsTranscriptionIdPropertiesPut(ctx, appID, authToken, assetId, versionId, transcriptionId, body)
Update transcription properties by ID

 Required roles:  - can_write_transcriptions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **versionId** | **string**|  | 
  **transcriptionId** | **string**|  | 
  **body** | [**AssetTranscriptionPropertiesSchema**](AssetTranscriptionPropertiesSchema.md)|  | 

### Return type

[**AssetTranscriptionPropertiesSchema**](AssetTranscriptionPropertiesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsGet**
> AssetsSchema V1AssetsGet(ctx, appID, authToken, optional)
Get list of assets

 Required roles:  - can_read_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **fieldName** | **optional.String**| filter by field_name | 

### Return type

[**AssetsSchema**](AssetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsPatch**
> V1AssetsPatch(ctx, appID, authToken, body)
Bulk update assets

 Required roles:  - can_write_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BulkAssetEditSchema**](BulkAssetEditSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsPost**
> AssetSchema V1AssetsPost(ctx, appID, authToken, body, optional)
Create a new asset

 Required roles:  - can_create_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**AssetCreateSchema**](AssetCreateSchema.md)|  | 
 **optional** | ***DefaultApiV1AssetsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applyDefaultAcls** | **optional.Bool**| Adds default ACLs to an asset | [default to true]
 **applyCollectionAcls** | **optional.Bool**| Adds containing collection&#39;s ACLs to an asset | [default to false]
 **assignToCollection** | **optional.Bool**| Adds the asset to the collection specified in the body | [default to false]

### Return type

[**AssetSchema**](AssetSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsPut**
> V1AssetsPut(ctx, appID, authToken, body)
Bulk update assets

 Required roles:  - can_write_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BulkAssetEditSchema**](BulkAssetEditSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsReindexPost**
> V1AssetsReindexPost(ctx, appID, authToken, optional)
Trigger reindexing of all assets

 Required roles:  - can_reindex_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1AssetsReindexPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AssetsReindexPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ReindexAllAssetsSchema**](ReindexAllAssetsSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsRelationTypesGet**
> RelationTypesSchema V1AssetsRelationTypesGet(ctx, appID, authToken)
Create a new asset relation type

 Required roles:  - can_read_asset_relations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**RelationTypesSchema**](RelationTypesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsRelationTypesPost**
> RelationTypeSchema V1AssetsRelationTypesPost(ctx, appID, authToken, body)
Create a new asset relation type

 Required roles:  - can_write_asset_relation_types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**RelationTypeSchema**](RelationTypeSchema.md)|  | 

### Return type

[**RelationTypeSchema**](RelationTypeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsRelationTypesRelationTypeDelete**
> V1AssetsRelationTypesRelationTypeDelete(ctx, appID, authToken, relationType)
Delete an asset relation type

 Required roles:  - can_delete_asset_relation_types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **relationType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsRelationTypesRelationTypeGet**
> RelationTypeSchema V1AssetsRelationTypesRelationTypeGet(ctx, appID, authToken, relationType)
Get a relation type

 Required roles:  - can_read_asset_relations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **relationType** | **string**|  | 

### Return type

[**RelationTypeSchema**](RelationTypeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsRelationTypesRelationTypePatch**
> RelationTypeSchema V1AssetsRelationTypesRelationTypePatch(ctx, appID, authToken, relationType, body)
Update an asset relation type

 Required roles:  - can_write_asset_relation_types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **relationType** | **string**|  | 
  **body** | [**RelationTypeSchema**](RelationTypeSchema.md)|  | 

### Return type

[**RelationTypeSchema**](RelationTypeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsRelationTypesRelationTypePut**
> RelationTypeSchema V1AssetsRelationTypesRelationTypePut(ctx, appID, authToken, relationType, body)
Update an asset relation type

 Required roles:  - can_write_asset_relation_types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **relationType** | **string**|  | 
  **body** | [**RelationTypeSchema**](RelationTypeSchema.md)|  | 

### Return type

[**RelationTypeSchema**](RelationTypeSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AssetsSegmentsReindexPost**
> V1AssetsSegmentsReindexPost(ctx, appID, authToken, body)
Trigger reindexing of all segments

 Required roles:  - can_reindex_segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**ReindexAllSegmentsSchema**](ReindexAllSegmentsSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdAncestorsGet**
> CollectionsSchema V1CollectionsCollectionIdAncestorsGet(ctx, appID, authToken, collectionId)
Returns list of ancestors of a collection

 Required roles:  - can_read_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 

### Return type

[**CollectionsSchema**](CollectionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdContentInfoGet**
> CollectionContentInfoSchema V1CollectionsCollectionIdContentInfoGet(ctx, appID, authToken, collectionId)
Returns a sub-collections and assets number for a specific collection

 Required roles:  - can_read_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 

### Return type

[**CollectionContentInfoSchema**](CollectionContentInfoSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdContentsGet**
> AssetsSchema V1CollectionsCollectionIdContentsGet(ctx, appID, authToken, collectionId, optional)
Returns contents of a collection by id

 Required roles:  - can_read_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
 **optional** | ***DefaultApiV1CollectionsCollectionIdContentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectionsCollectionIdContentsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **objectTypes** | **optional.String**| Comma separated list of content types. Example - assets,collections | 
 **objectIds** | **optional.String**| Comma separated list of content ids. | 
 **externalId** | **optional.String**|  | 
 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **filter** | **optional.String**| A comma separated list of fieldnames with order For example - first_name,eq,Vlad;last_name,eq,Gudkov | 

### Return type

[**AssetsSchema**](AssetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdContentsObjectTypeObjectIdDelete**
> V1CollectionsCollectionIdContentsObjectTypeObjectIdDelete(ctx, appID, authToken, collectionId, objectType, objectId)
Delete a particular content object in a collection by id

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdContentsObjectTypeObjectIdPut**
> V1CollectionsCollectionIdContentsObjectTypeObjectIdPut(ctx, appID, authToken, collectionId, objectType, objectId, body)
Update an order of a particular content object in a collection

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **body** | [**CollectionContentOrderingSchema**](CollectionContentOrderingSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdContentsObjectTypeObjectIdReindexPost**
> V1CollectionsCollectionIdContentsObjectTypeObjectIdReindexPost(ctx, appID, authToken, collectionId, objectType, objectId, body)
Reindex collection content

 Required roles:  - can_reindex_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **body** | [**ReindexCollectionContentSchema**](ReindexCollectionContentSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdContentsOrderingCustomDelete**
> V1CollectionsCollectionIdContentsOrderingCustomDelete(ctx, appID, authToken, collectionId)
Disable custom ordering for a collection's content

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdContentsOrderingCustomPost**
> V1CollectionsCollectionIdContentsOrderingCustomPost(ctx, appID, authToken, collectionId, body)
Enable custom ordering for a collection's content

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **body** | [**CreateCollectionContentOrderingSchema**](CreateCollectionContentOrderingSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdContentsPost**
> V1CollectionsCollectionIdContentsPost(ctx, appID, authToken, collectionId, body)
Add an object to a collection

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **body** | [**CollectionContentSchema**](CollectionContentSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdDelete**
> V1CollectionsCollectionIdDelete(ctx, appID, authToken, collectionId)
Delete a particular collection by id

 Required roles:  - can_delete_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdGet**
> CollectionSchema V1CollectionsCollectionIdGet(ctx, appID, authToken, collectionId)
Returns a particular collection by id

 Required roles:  - can_read_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 

### Return type

[**CollectionSchema**](CollectionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdKeyframesPost**
> V1CollectionsCollectionIdKeyframesPost(ctx, appID, authToken, collectionId, body)
Pick up to three asset_ids for collection keyframes

 Required roles:  - can_reindex_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **body** | [**SynchronizeCollectionKeyframesSchema**](SynchronizeCollectionKeyframesSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdPatch**
> CollectionSchema V1CollectionsCollectionIdPatch(ctx, appID, authToken, collectionId, body, optional)
Update collection

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **body** | [**CollectionInputSchema**](CollectionInputSchema.md)|  | 
 **optional** | ***DefaultApiV1CollectionsCollectionIdPatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectionsCollectionIdPatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **changeParentMode** | **optional.String**| Set to &#39;move&#39; or &#39;copy&#39;. Ignored if &#39;parent_id&#39; hasn&#39;t changed | [default to move]

### Return type

[**CollectionSchema**](CollectionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdPurgeDelete**
> V1CollectionsCollectionIdPurgeDelete(ctx, appID, authToken, collectionId)
Purges deleted collection by id immediately

 Required roles:  - can_purge_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdPut**
> CollectionSchema V1CollectionsCollectionIdPut(ctx, appID, authToken, collectionId, body, optional)
Update collection

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **body** | [**CollectionInputSchema**](CollectionInputSchema.md)|  | 
 **optional** | ***DefaultApiV1CollectionsCollectionIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectionsCollectionIdPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **changeParentMode** | **optional.String**| Set to &#39;move&#39; or &#39;copy&#39;. Ignored if &#39;parent_id&#39; hasn&#39;t changed | [default to move]

### Return type

[**CollectionSchema**](CollectionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdReindexContentsPost**
> V1CollectionsCollectionIdReindexContentsPost(ctx, appID, authToken, collectionId)
Reindex collection and its content

 Required roles:  - can_reindex_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdReindexPost**
> V1CollectionsCollectionIdReindexPost(ctx, appID, authToken, collectionId, body)
Reindex collection

 Required roles:  - can_reindex_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **body** | [**ReindexCollectionSchema**](ReindexCollectionSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdRestorePut**
> V1CollectionsCollectionIdRestorePut(ctx, appID, authToken, collectionId)
Restore deleted collection by id

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdSearchDocumentPut**
> V1CollectionsCollectionIdSearchDocumentPut(ctx, appID, authToken, collectionId, body)
Update metadata for collection

 Required roles:  - can_reindex_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **body** | [**CollectionSchema**](CollectionSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsCollectionIdSubcollectionsPost**
> V1CollectionsCollectionIdSubcollectionsPost(ctx, appID, authToken, collectionId, body)
Copy a collection (recursively) in to another collection

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
  **body** | [**CollectionContentSchema**](CollectionContentSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsGet**
> CollectionsSchema V1CollectionsGet(ctx, appID, authToken, optional)
Get list of collections

 Required roles:  - can_read_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1CollectionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - title,asc;is_root,desc | 
 **isRoot** | **optional.String**| Filter by is_root | 
 **status** | **optional.String**| Filter by status | 

### Return type

[**CollectionsSchema**](CollectionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsPost**
> CollectionSchema V1CollectionsPost(ctx, appID, authToken, body, optional)
Create a new collection

 Required roles:  - can_create_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**CollectionInputSchema**](CollectionInputSchema.md)|  | 
 **optional** | ***DefaultApiV1CollectionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectionsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applyDefaultAcls** | **optional.Bool**| Adds default ACLs to a collection | [default to true]
 **applyCollectionAcls** | **optional.Bool**| Adds containing collection&#39;s ACLs to a collection | [default to false]

### Return type

[**CollectionSchema**](CollectionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CollectionsReindexPost**
> V1CollectionsReindexPost(ctx, appID, authToken, optional)
Trigger reindexing of all collections

 Required roles:  - can_reindex_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1CollectionsReindexPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1CollectionsReindexPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ReindexAllCollectionsSchema**](ReindexAllCollectionsSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CustomActionsContextActionIdCallbackPost**
> CustomActionCallbackReplySchema V1CustomActionsContextActionIdCallbackPost(ctx, appID, authToken, context, actionId, body)
Schedules a celery task that will call custom action

 Required roles:  - can_read_custom_actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **context** | **string**|  | 
  **actionId** | **string**|  | 
  **body** | [**CustomActionCallbackSchema**](CustomActionCallbackSchema.md)|  | 

### Return type

[**CustomActionCallbackReplySchema**](CustomActionCallbackReplySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CustomActionsContextActionIdDelete**
> CustomActionSchema V1CustomActionsContextActionIdDelete(ctx, appID, authToken, context, actionId)
Deletes an custom action



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **context** | **string**|  | 
  **actionId** | **string**|  | 

### Return type

[**CustomActionSchema**](CustomActionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CustomActionsContextActionIdGet**
> CustomActionSchema V1CustomActionsContextActionIdGet(ctx, appID, authToken, context, actionId)
Get an asset custom action

 Required roles:  - can_read_custom_actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **context** | **string**|  | 
  **actionId** | **string**|  | 

### Return type

[**CustomActionSchema**](CustomActionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CustomActionsContextActionIdPatch**
> CustomActionSchema V1CustomActionsContextActionIdPatch(ctx, appID, authToken, context, actionId, body)
Update an custom action

 Required roles:  - can_write_custom_actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **context** | **string**|  | 
  **actionId** | **string**|  | 
  **body** | [**CustomActionSchema**](CustomActionSchema.md)|  | 

### Return type

[**CustomActionSchema**](CustomActionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CustomActionsContextActionIdPut**
> CustomActionSchema V1CustomActionsContextActionIdPut(ctx, appID, authToken, context, actionId, body)
Update an custom action

 Required roles:  - can_write_custom_actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **context** | **string**|  | 
  **actionId** | **string**|  | 
  **body** | [**CustomActionSchema**](CustomActionSchema.md)|  | 

### Return type

[**CustomActionSchema**](CustomActionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CustomActionsContextGet**
> CustomActionsSchema V1CustomActionsContextGet(ctx, appID, authToken, context)
Get list of custom actions by context

 Required roles:  - can_read_custom_actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **context** | **string**|  | 

### Return type

[**CustomActionsSchema**](CustomActionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CustomActionsContextPost**
> CustomActionSchema V1CustomActionsContextPost(ctx, appID, authToken, context, body)
Create an custom action

 Required roles:  - can_write_custom_actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **context** | **string**|  | 
  **body** | [**CustomActionSchema**](CustomActionSchema.md)|  | 

### Return type

[**CustomActionSchema**](CustomActionSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CustomActionsGet**
> CustomActionsSchema V1CustomActionsGet(ctx, appID, authToken)
Get list of custom actions

 Required roles:  - can_read_custom_actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**CustomActionsSchema**](CustomActionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1CustomActionsSharedContextActionIdCallbackPost**
> CustomActionCallbackReplySchema V1CustomActionsSharedContextActionIdCallbackPost(ctx, appID, authToken, context, actionId, body)
Schedules a celery task that will call custom action on shares

 Required roles:  - can_read_custom_actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **context** | **string**|  | 
  **actionId** | **string**|  | 
  **body** | [**CustomActionCallbackSchema**](CustomActionCallbackSchema.md)|  | 

### Return type

[**CustomActionCallbackReplySchema**](CustomActionCallbackReplySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueAssetsDelete**
> V1DeleteQueueAssetsDelete(ctx, appID, authToken, body)
Delete assets from delete queue (Mark assets as active again)

 Required roles:  - can_write_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueAssetsGet**
> AssetsSchema V1DeleteQueueAssetsGet(ctx, appID, authToken, optional)
Get deleted objects

 Required roles:  - can_read_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1DeleteQueueAssetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1DeleteQueueAssetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **filter** | **optional.String**| A comma separated list of fieldnames with order For example - first_name,eq,Vlad;last_name,eq,Gudkov | 

### Return type

[**AssetsSchema**](AssetsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueAssetsPost**
> V1DeleteQueueAssetsPost(ctx, appID, authToken, body)
Add assets to a delete queue (Mark assets as deleted)

 Required roles:  - can_delete_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueAssetsPurgeAllPost**
> V1DeleteQueueAssetsPurgeAllPost(ctx, appID, authToken)
Purge all assets from delete queue (Permanently delete)

 Required roles:  - can_purge_assets

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

# **V1DeleteQueueAssetsPurgePost**
> V1DeleteQueueAssetsPurgePost(ctx, appID, authToken, body)
Purge assets from delete queue (Permanently delete)

 Required roles:  - can_purge_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueBulkPost**
> V1DeleteQueueBulkPost(ctx, appID, authToken, body)
Bulk delete objects

 Required roles:  - can_delete_assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BulkDeleteSchema**](BulkDeleteSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueCollectionsDelete**
> V1DeleteQueueCollectionsDelete(ctx, appID, authToken, body)
Delete collections from delete queue (Mark collections as active again)

 Required roles:  - can_write_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueCollectionsGet**
> CollectionsSchema V1DeleteQueueCollectionsGet(ctx, appID, authToken, optional)
Get list of collections

 Required roles:  - can_read_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1DeleteQueueCollectionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1DeleteQueueCollectionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**| The number of items for each page | 
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **filter** | **optional.String**| A comma separated list of fieldnames with order For example - first_name,eq,Vlad;last_name,eq,Gudkov | 

### Return type

[**CollectionsSchema**](CollectionsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueCollectionsPost**
> V1DeleteQueueCollectionsPost(ctx, appID, authToken, body)
Add collections to a delete queue (Mark collections as deleted)

 Required roles:  - can_delete_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueueCollectionsPurgeAllPost**
> V1DeleteQueueCollectionsPurgeAllPost(ctx, appID, authToken)
Purge all collections from delete queue (Permanently delete)

 Required roles:  - can_purge_collections

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

# **V1DeleteQueueCollectionsPurgePost**
> V1DeleteQueueCollectionsPurgePost(ctx, appID, authToken, body)
Purge collections from delete queue (Permanently delete)

 Required roles:  - can_purge_collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**DeleteQueueSchema**](DeleteQueueSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DeleteQueuePurgeAllPost**
> V1DeleteQueuePurgeAllPost(ctx, appID, authToken)
Purge all assets and collections from delete queue (Permanently delete)

 Required roles:  - can_purge_assets - can_purge_collections

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

# **V1ObjectTypeObjectIdApprovalsDelete**
> V1ObjectTypeObjectIdApprovalsDelete(ctx, appID, authToken, objectId, objectType)
Deletes an objects approval status



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdApprovalsExternalEmailDelete**
> V1ObjectTypeObjectIdApprovalsExternalEmailDelete(ctx, appID, authToken, objectId, objectType, email)
Deletes an objects approval status by user_id

 Required roles:  - can_delete_approval_status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 
  **email** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdApprovalsGet**
> ApprovalsBySchema V1ObjectTypeObjectIdApprovalsGet(ctx, appID, authToken, objectId, objectType)
Returns an objects approval request

 Required roles:  - can_read_approval_request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 

### Return type

[**ApprovalsBySchema**](ApprovalsBySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdApprovalsPut**
> ApprovalBySchema V1ObjectTypeObjectIdApprovalsPut(ctx, appID, authToken, objectId, objectType, body)
Returns an objects approval status

 Required roles:  - can_write_approval_status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**ApprovalBySchema**](ApprovalBySchema.md)|  | 

### Return type

[**ApprovalBySchema**](ApprovalBySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdApprovalsRequestDelete**
> ApprovalSchema V1ObjectTypeObjectIdApprovalsRequestDelete(ctx, appID, authToken, objectId, objectType)
Deletes an objects approval request

 Required roles:  - can_delete_approval_request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 

### Return type

[**ApprovalSchema**](ApprovalSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdApprovalsRequestGet**
> ApprovalSchema V1ObjectTypeObjectIdApprovalsRequestGet(ctx, appID, authToken, objectId, objectType)
Returns an objects approval request

 Required roles:  - can_read_approval_request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 

### Return type

[**ApprovalSchema**](ApprovalSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdApprovalsRequestPatch**
> ApprovalSchema V1ObjectTypeObjectIdApprovalsRequestPatch(ctx, appID, authToken, objectId, objectType, body)
Edits an approval request

 Required roles:  - can_write_approval_request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**ApprovalSchema**](ApprovalSchema.md)|  | 

### Return type

[**ApprovalSchema**](ApprovalSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdApprovalsRequestPost**
> ApprovalSchema V1ObjectTypeObjectIdApprovalsRequestPost(ctx, appID, authToken, objectId, objectType, body)
Creates an objects approval request

 Required roles:  - can_write_approval_request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**ApprovalSchema**](ApprovalSchema.md)|  | 

### Return type

[**ApprovalSchema**](ApprovalSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdApprovalsRequestPut**
> ApprovalSchema V1ObjectTypeObjectIdApprovalsRequestPut(ctx, appID, authToken, objectId, objectType, body)
Edits an approval request

 Required roles:  - can_write_approval_request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**ApprovalSchema**](ApprovalSchema.md)|  | 

### Return type

[**ApprovalSchema**](ApprovalSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdApprovalsUserUserIdDelete**
> V1ObjectTypeObjectIdApprovalsUserUserIdDelete(ctx, appID, authToken, objectId, objectType, userId)
Deletes an objects approval status by user_id

 Required roles:  - can_delete_approval_status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectId** | **string**|  | 
  **objectType** | **string**|  | 
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesGet**
> SharesSchema V1ObjectTypeObjectIdSharesGet(ctx, appID, authToken, objectType, objectId, optional)
Get list of shares

 Required roles:  - can_read_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
 **optional** | ***DefaultApiV1ObjectTypeObjectIdSharesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ObjectTypeObjectIdSharesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **perPage** | **optional.Int32**| The number of items for each page | 
 **lastId** | **optional.String**|  | 

### Return type

[**SharesSchema**](SharesSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesPost**
> ShareSchema V1ObjectTypeObjectIdSharesPost(ctx, appID, authToken, objectType, objectId, body)
Create a new share

 Required roles:  - can_write_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **body** | [**ShareCreateSchema**](ShareCreateSchema.md)|  | 

### Return type

[**ShareSchema**](ShareSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesShareIdDelete**
> V1ObjectTypeObjectIdSharesShareIdDelete(ctx, appID, authToken, objectType, objectId, shareId)
Delete a particular share by id

 Required roles:  - can_delete_object_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **shareId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesShareIdGet**
> ShareSchema V1ObjectTypeObjectIdSharesShareIdGet(ctx, appID, authToken, objectType, objectId, shareId)
Returns a particular share by id



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **shareId** | **string**|  | 

### Return type

[**ShareSchema**](ShareSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesShareIdPut**
> ShareSchema V1ObjectTypeObjectIdSharesShareIdPut(ctx, appID, authToken, objectType, objectId, shareId, body)
Update share

 Required roles:  - can_write_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **shareId** | **string**|  | 
  **body** | [**ShareSchema**](ShareSchema.md)|  | 

### Return type

[**ShareSchema**](ShareSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesShareIdUsersGet**
> ShareUsersSchema V1ObjectTypeObjectIdSharesShareIdUsersGet(ctx, appID, authToken, objectType, objectId, shareId, optional)
Get list of share users

 Required roles:  - can_read_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **shareId** | **string**|  | 
 **optional** | ***DefaultApiV1ObjectTypeObjectIdSharesShareIdUsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1ObjectTypeObjectIdSharesShareIdUsersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **perPage** | **optional.Int32**| The number of items for each page | 
 **lastId** | **optional.String**|  | 

### Return type

[**ShareUsersSchema**](ShareUsersSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesShareIdUsersPost**
> ShareUserSchema V1ObjectTypeObjectIdSharesShareIdUsersPost(ctx, appID, authToken, objectType, objectId, shareId, body)
Add a new share_user to a share

 Required roles:  - can_write_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **shareId** | **string**|  | 
  **body** | [**ShareUserSchema**](ShareUserSchema.md)|  | 

### Return type

[**ShareUserSchema**](ShareUserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdDelete**
> V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdDelete(ctx, appID, authToken, objectType, objectId, shareId, shareUserId)
Delete a particular share_user user by id

 Required roles:  - can_write_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **shareId** | **string**|  | 
  **shareUserId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdGet**
> ShareUserSchema V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdGet(ctx, appID, authToken, objectType, objectId, shareId, shareUserId)
Returns a particular share user by id

 Required roles:  - can_read_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **shareId** | **string**|  | 
  **shareUserId** | **string**|  | 

### Return type

[**ShareUserSchema**](ShareUserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdPatch**
> ShareUserSchema V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdPatch(ctx, appID, authToken, objectType, objectId, shareId, shareUserId, body)
Update share user

 Required roles:  - can_write_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **shareId** | **string**|  | 
  **shareUserId** | **string**|  | 
  **body** | [**ShareUserSchema**](ShareUserSchema.md)|  | 

### Return type

[**ShareUserSchema**](ShareUserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdPut**
> ShareUserSchema V1ObjectTypeObjectIdSharesShareIdUsersShareUserIdPut(ctx, appID, authToken, objectType, objectId, shareId, shareUserId, body)
Update share user

 Required roles:  - can_write_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **shareId** | **string**|  | 
  **shareUserId** | **string**|  | 
  **body** | [**ShareUserSchema**](ShareUserSchema.md)|  | 

### Return type

[**ShareUserSchema**](ShareUserSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdSharesUrlPost**
> ShareUrlSchema V1ObjectTypeObjectIdSharesUrlPost(ctx, appID, authToken, objectType, objectId, body)
Generates a URL for the shared object

 Required roles:  - can_write_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **body** | [**ShareUrlCreateSchema**](ShareUrlCreateSchema.md)|  | 

### Return type

[**ShareUrlSchema**](ShareURLSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdVersionsVersionIdApprovalsGet**
> ApprovalsBySchema V1ObjectTypeObjectIdVersionsVersionIdApprovalsGet(ctx, appID, authToken, objectType, objectId, versionId)
Returns an objects approval request by version

 Required roles:  - can_read_approval_request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **versionId** | **string**|  | 

### Return type

[**ApprovalsBySchema**](ApprovalsBySchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ObjectTypeObjectIdVersionsVersionIdApprovalsRequestGet**
> ApprovalSchema V1ObjectTypeObjectIdVersionsVersionIdApprovalsRequestGet(ctx, appID, authToken, objectType, objectId, versionId)
Returns an objects approval request by version

 Required roles:  - can_read_approval_request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **objectId** | **string**|  | 
  **versionId** | **string**|  | 

### Return type

[**ApprovalSchema**](ApprovalSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ShareObjectTypePost**
> ShareSchema V1ShareObjectTypePost(ctx, appID, authToken, objectType, body)
Create a new share of multiple objects (currently only assets are supported)

 Required roles:  - can_write_shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **objectType** | **string**|  | 
  **body** | [**BulkShareCreateSchema**](BulkShareCreateSchema.md)|  | 

### Return type

[**ShareSchema**](ShareSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesAuthLoginPost**
> ShareTokenSchema V1SharesAuthLoginPost(ctx, appID, body)
Login for share



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **body** | [**ShareLoginSchema**](ShareLoginSchema.md)|  | 

### Return type

[**ShareTokenSchema**](ShareTokenSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1SharesAuthTokenGet**
> ShareTokenSchema V1SharesAuthTokenGet(ctx, appID, shareAuthToken)
Check if a token is valid



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **shareAuthToken** | **string**|  | 

### Return type

[**ShareTokenSchema**](ShareTokenSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


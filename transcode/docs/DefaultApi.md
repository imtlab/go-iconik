# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/transcode*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AnalyzeAssetsAssetIdPost**](DefaultApi.md#V1AnalyzeAssetsAssetIdPost) | **Post** /v1/analyze/assets/{asset_id}/ | Start a job that sends an asset for analysis
[**V1AnalyzeAssetsAssetIdProfilesDefaultMediaTypePost**](DefaultApi.md#V1AnalyzeAssetsAssetIdProfilesDefaultMediaTypePost) | **Post** /v1/analyze/assets/{asset_id}/profiles/default/{media_type}/ | Start a job that sends an asset for analysis
[**V1AnalyzeAssetsAssetIdProfilesDefaultPost**](DefaultApi.md#V1AnalyzeAssetsAssetIdProfilesDefaultPost) | **Post** /v1/analyze/assets/{asset_id}/profiles/default/ | Start a job that sends an asset for analysis with a default analysis profile
[**V1AnalyzeAssetsAssetIdProfilesProfileIdPost**](DefaultApi.md#V1AnalyzeAssetsAssetIdProfilesProfileIdPost) | **Post** /v1/analyze/assets/{asset_id}/profiles/{profile_id}/ | Start a job that sends an asset for analysis with a custom analysis profile
[**V1AnalyzeBulkPost**](DefaultApi.md#V1AnalyzeBulkPost) | **Post** /v1/analyze/bulk/ | Start a job that sends objects for analysis using a custom analysis profile
[**V1EdgeTranscodeJobsJobIdAcknowledgePost**](DefaultApi.md#V1EdgeTranscodeJobsJobIdAcknowledgePost) | **Post** /v1/edge_transcode/jobs/{job_id}/acknowledge/ | Acknowledge an edge transcode job
[**V1EdgeTranscodeWorkersGet**](DefaultApi.md#V1EdgeTranscodeWorkersGet) | **Get** /v1/edge_transcode/workers/ | Get edge transcode workers
[**V1EdgeTranscodeWorkersPost**](DefaultApi.md#V1EdgeTranscodeWorkersPost) | **Post** /v1/edge_transcode/workers/ | Create a new edge transcode worker
[**V1EdgeTranscodeWorkersWorkerIdDelete**](DefaultApi.md#V1EdgeTranscodeWorkersWorkerIdDelete) | **Delete** /v1/edge_transcode/workers/{worker_id}/ | Delete a edge transcode worker
[**V1EdgeTranscodeWorkersWorkerIdGet**](DefaultApi.md#V1EdgeTranscodeWorkersWorkerIdGet) | **Get** /v1/edge_transcode/workers/{worker_id}/ | Get a edge transcode worker
[**V1EdgeTranscodeWorkersWorkerIdPatch**](DefaultApi.md#V1EdgeTranscodeWorkersWorkerIdPatch) | **Patch** /v1/edge_transcode/workers/{worker_id}/ | Update a edge transcode worker
[**V1EdgeTranscodeWorkersWorkerIdPut**](DefaultApi.md#V1EdgeTranscodeWorkersWorkerIdPut) | **Put** /v1/edge_transcode/workers/{worker_id}/ | Update a edge transcode worker
[**V1KeyframesCollectionsCollectionIdPost**](DefaultApi.md#V1KeyframesCollectionsCollectionIdPost) | **Post** /v1/keyframes/collections/{collection_id}/ | Start a job that creates a collection keyframe
[**V1StoragesStorageIdDelete**](DefaultApi.md#V1StoragesStorageIdDelete) | **Delete** /v1/storages/{storage_id}/ | Cancel all transcode jobs linked to the storage
[**V1StoragesStorageIdEdgeTranscodeJobsGet**](DefaultApi.md#V1StoragesStorageIdEdgeTranscodeJobsGet) | **Get** /v1/storages/{storage_id}/edge_transcode/jobs/ | Get a edge transcode jobs from the job queue
[**V1StoragesStorageIdFilesFileIdTranscodeDelete**](DefaultApi.md#V1StoragesStorageIdFilesFileIdTranscodeDelete) | **Delete** /v1/storages/{storage_id}/files/{file_id}/transcode/ | Delete local storage transcode job.
[**V1StoragesStorageIdTranscodeGet**](DefaultApi.md#V1StoragesStorageIdTranscodeGet) | **Get** /v1/storages/{storage_id}/transcode/ | Get pending local storage transcode jobs.
[**V1StoragesStorageIdTranscodeRecordIdDelete**](DefaultApi.md#V1StoragesStorageIdTranscodeRecordIdDelete) | **Delete** /v1/storages/{storage_id}/transcode/{record_id}/ | Delete local storage transcode job.
[**V1StoragesStorageIdTranscodeRecordIdGet**](DefaultApi.md#V1StoragesStorageIdTranscodeRecordIdGet) | **Get** /v1/storages/{storage_id}/transcode/{record_id}/ | Get local storage transcode job.
[**V1TranscodePost**](DefaultApi.md#V1TranscodePost) | **Post** /v1/transcode/ | Starts a new transcode.
[**V1TranscodeQueueGet**](DefaultApi.md#V1TranscodeQueueGet) | **Get** /v1/transcode/queue/ | Get all the statuses of the queued transcode jobs
[**V1TranscodeQueueSystemGet**](DefaultApi.md#V1TranscodeQueueSystemGet) | **Get** /v1/transcode/queue/system/ | Get the status of the transcode job queues
[**V1TranscodeTranscodeJobIdDelete**](DefaultApi.md#V1TranscodeTranscodeJobIdDelete) | **Delete** /v1/transcode/{transcode_job_id}/ | Cancel a particular transcode job by id
[**V1TranscodeTranscodeJobIdGet**](DefaultApi.md#V1TranscodeTranscodeJobIdGet) | **Get** /v1/transcode/{transcode_job_id}/ | Get transcode job
[**V1TranscodeTranscodeJobIdPositionPositionPost**](DefaultApi.md#V1TranscodeTranscodeJobIdPositionPositionPost) | **Post** /v1/transcode/{transcode_job_id}/position/{position}/ | Move transcode job to top or bottom of the queue
[**V1TranscodeTranscodeJobIdPriorityPriorityPut**](DefaultApi.md#V1TranscodeTranscodeJobIdPriorityPriorityPut) | **Put** /v1/transcode/{transcode_job_id}/priority/{priority}/ | Change transcode job priority
[**V1TranscribeAssetsAssetIdProfilesDefaultPost**](DefaultApi.md#V1TranscribeAssetsAssetIdProfilesDefaultPost) | **Post** /v1/transcribe/assets/{asset_id}/profiles/default/ | Start a job that sends an asset to default transcription service
[**V1TranscribeBulkPost**](DefaultApi.md#V1TranscribeBulkPost) | **Post** /v1/transcribe/bulk/ | Start a job that sends multiple objects to transcription service


# **V1AnalyzeAssetsAssetIdPost**
> V1AnalyzeAssetsAssetIdPost(ctx, appID, authToken, assetId, optional)
Start a job that sends an asset for analysis

 Required roles:  - can_analyze_content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AnalyzeAssetsAssetIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AnalyzeAssetsAssetIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of AnalyzeSchema**](AnalyzeSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalyzeAssetsAssetIdProfilesDefaultMediaTypePost**
> V1AnalyzeAssetsAssetIdProfilesDefaultMediaTypePost(ctx, appID, authToken, assetId, mediaType, optional)
Start a job that sends an asset for analysis

with a default analysis profile of specified media type Required roles:  - can_analyze_content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **mediaType** | **string**|  | 
 **optional** | ***DefaultApiV1AnalyzeAssetsAssetIdProfilesDefaultMediaTypePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AnalyzeAssetsAssetIdProfilesDefaultMediaTypePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of Body1**](Body1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalyzeAssetsAssetIdProfilesDefaultPost**
> V1AnalyzeAssetsAssetIdProfilesDefaultPost(ctx, appID, authToken, assetId, optional)
Start a job that sends an asset for analysis with a default analysis profile

 Required roles:  - can_analyze_content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1AnalyzeAssetsAssetIdProfilesDefaultPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AnalyzeAssetsAssetIdProfilesDefaultPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of Body**](Body.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalyzeAssetsAssetIdProfilesProfileIdPost**
> V1AnalyzeAssetsAssetIdProfilesProfileIdPost(ctx, appID, authToken, assetId, profileId, optional)
Start a job that sends an asset for analysis with a custom analysis profile

 Required roles:  - can_analyze_content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
  **profileId** | **string**|  | 
 **optional** | ***DefaultApiV1AnalyzeAssetsAssetIdProfilesProfileIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AnalyzeAssetsAssetIdProfilesProfileIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of Body2**](Body2.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AnalyzeBulkPost**
> V1AnalyzeBulkPost(ctx, appID, authToken, optional)
Start a job that sends objects for analysis using a custom analysis profile

 Required roles:  - can_analyze_content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1AnalyzeBulkPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1AnalyzeBulkPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BulkAnalyzeSchema**](BulkAnalyzeSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1EdgeTranscodeJobsJobIdAcknowledgePost**
> V1EdgeTranscodeJobsJobIdAcknowledgePost(ctx, appID, authToken, jobId)
Acknowledge an edge transcode job

 Required roles:  - can_read_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **jobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1EdgeTranscodeWorkersGet**
> EdgeTranscodeWorkersSchema V1EdgeTranscodeWorkersGet(ctx, appID, authToken)
Get edge transcode workers

 Required roles:  - is_storage_worker - can_read_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 

### Return type

[**EdgeTranscodeWorkersSchema**](EdgeTranscodeWorkersSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1EdgeTranscodeWorkersPost**
> EdgeTranscodeWorkerSchema V1EdgeTranscodeWorkersPost(ctx, appID, authToken, body)
Create a new edge transcode worker

 Required roles:  - is_storage_worker - can_write_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**EdgeTranscodeWorkerSchema**](EdgeTranscodeWorkerSchema.md)|  | 

### Return type

[**EdgeTranscodeWorkerSchema**](EdgeTranscodeWorkerSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1EdgeTranscodeWorkersWorkerIdDelete**
> V1EdgeTranscodeWorkersWorkerIdDelete(ctx, appID, authToken, workerId)
Delete a edge transcode worker

 Required roles:  - is_storage_worker - can_write_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **workerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1EdgeTranscodeWorkersWorkerIdGet**
> EdgeTranscodeWorkerSchema V1EdgeTranscodeWorkersWorkerIdGet(ctx, appID, authToken, workerId)
Get a edge transcode worker

 Required roles:  - is_storage_worker - can_read_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **workerId** | **string**|  | 

### Return type

[**EdgeTranscodeWorkerSchema**](EdgeTranscodeWorkerSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1EdgeTranscodeWorkersWorkerIdPatch**
> EdgeTranscodeWorkerSchema V1EdgeTranscodeWorkersWorkerIdPatch(ctx, appID, authToken, workerId, body)
Update a edge transcode worker

 Required roles:  - is_storage_worker - can_write_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **workerId** | **string**|  | 
  **body** | [**EdgeTranscodeWorkerSchema**](EdgeTranscodeWorkerSchema.md)|  | 

### Return type

[**EdgeTranscodeWorkerSchema**](EdgeTranscodeWorkerSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1EdgeTranscodeWorkersWorkerIdPut**
> EdgeTranscodeWorkerSchema V1EdgeTranscodeWorkersWorkerIdPut(ctx, appID, authToken, workerId, body)
Update a edge transcode worker

 Required roles:  - is_storage_worker - can_write_transcoders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **workerId** | **string**|  | 
  **body** | [**EdgeTranscodeWorkerSchema**](EdgeTranscodeWorkerSchema.md)|  | 

### Return type

[**EdgeTranscodeWorkerSchema**](EdgeTranscodeWorkerSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1KeyframesCollectionsCollectionIdPost**
> V1KeyframesCollectionsCollectionIdPost(ctx, appID, authToken, collectionId, optional)
Start a job that creates a collection keyframe

 Required roles:  - can_write_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **collectionId** | **string**|  | 
 **optional** | ***DefaultApiV1KeyframesCollectionsCollectionIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1KeyframesCollectionsCollectionIdPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of GenerateCollectionKeyframeSchema**](GenerateCollectionKeyframeSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdDelete**
> V1StoragesStorageIdDelete(ctx, appID, authToken, storageId, optional)
Cancel all transcode jobs linked to the storage

 Required roles:  - can_delete_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of AbortStorageTranscodeJobsSchema**](AbortStorageTranscodeJobsSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdEdgeTranscodeJobsGet**
> EdgeTranscodeJobsSchema V1StoragesStorageIdEdgeTranscodeJobsGet(ctx, appID, authToken, storageId, optional)
Get a edge transcode jobs from the job queue

 Required roles:  - can_read_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdEdgeTranscodeJobsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdEdgeTranscodeJobsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Int32**| The max number of items to fetch | [default to 10]

### Return type

[**EdgeTranscodeJobsSchema**](EdgeTranscodeJobsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdFilesFileIdTranscodeDelete**
> V1StoragesStorageIdFilesFileIdTranscodeDelete(ctx, authToken, appID, storageId, fileId)
Delete local storage transcode job.

 Required roles:  - can_read_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**|  | 
  **fileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTranscodeGet**
> LocalStorageFileTranscodeJobsSchema V1StoragesStorageIdTranscodeGet(ctx, authToken, appID, storageId, optional)
Get pending local storage transcode jobs.

 Required roles:  - can_read_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**|  | 
 **optional** | ***DefaultApiV1StoragesStorageIdTranscodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1StoragesStorageIdTranscodeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **lastId** | **optional.String**| ID of a last transcode job entity on previous page | 

### Return type

[**LocalStorageFileTranscodeJobsSchema**](LocalStorageFileTranscodeJobsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTranscodeRecordIdDelete**
> V1StoragesStorageIdTranscodeRecordIdDelete(ctx, authToken, appID, storageId, recordId)
Delete local storage transcode job.

 Required roles:  - can_read_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**|  | 
  **recordId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1StoragesStorageIdTranscodeRecordIdGet**
> LocalStorageFileTranscodeJobSchema V1StoragesStorageIdTranscodeRecordIdGet(ctx, authToken, appID, storageId, recordId)
Get local storage transcode job.

 Required roles:  - can_read_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authToken** | **string**|  | 
  **appID** | **string**|  | 
  **storageId** | **string**|  | 
  **recordId** | **string**|  | 

### Return type

[**LocalStorageFileTranscodeJobSchema**](LocalStorageFileTranscodeJobSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodePost**
> JobSchema V1TranscodePost(ctx, appID, authToken, body)
Starts a new transcode.

Use /API/files/v1/assets/ID/files/ID/keyframes instead Required roles:  - can_write_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**JobSchema**](JobSchema.md)|  | 

### Return type

[**JobSchema**](JobSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodeQueueGet**
> TranscodeQueueSchema V1TranscodeQueueGet(ctx, appID, authToken, optional)
Get all the statuses of the queued transcode jobs

 Required roles:  - can_read_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1TranscodeQueueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1TranscodeQueueGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **optional.Int32**|  | 
 **page** | **optional.Int32**|  | 
 **sort** | **optional.String**| A comma separated list of fieldnames without spaces. object_type,user_id,retry_count,priority,type,status | 

### Return type

[**TranscodeQueueSchema**](TranscodeQueueSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodeQueueSystemGet**
> TranscodeQueueSchema V1TranscodeQueueSystemGet(ctx, appID, authToken, optional)
Get the status of the transcode job queues



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1TranscodeQueueSystemGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1TranscodeQueueSystemGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perDomainId** | **optional.Bool**|  | 
 **perPage** | **optional.Int32**|  | 
 **page** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 

### Return type

[**TranscodeQueueSchema**](TranscodeQueueSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodeTranscodeJobIdDelete**
> V1TranscodeTranscodeJobIdDelete(ctx, appID, authToken, transcodeJobId)
Cancel a particular transcode job by id

 Required roles:  - can_delete_transcode_jobs Required roles:  - can_delete_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcodeJobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodeTranscodeJobIdGet**
> JobSchema V1TranscodeTranscodeJobIdGet(ctx, appID, authToken, transcodeJobId)
Get transcode job

 Required roles:  - can_read_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcodeJobId** | **string**|  | 

### Return type

[**JobSchema**](JobSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodeTranscodeJobIdPositionPositionPost**
> V1TranscodeTranscodeJobIdPositionPositionPost(ctx, appID, authToken, transcodeJobId, position)
Move transcode job to top or bottom of the queue

 Required roles:  - can_write_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcodeJobId** | **string**|  | 
  **position** | **string**| move transcode job to \&quot;top\&quot; or \&quot;bottom\&quot; position | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscodeTranscodeJobIdPriorityPriorityPut**
> V1TranscodeTranscodeJobIdPriorityPriorityPut(ctx, appID, authToken, transcodeJobId, priority)
Change transcode job priority

 Required roles:  - can_write_transcode_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **transcodeJobId** | **string**|  | 
  **priority** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscribeAssetsAssetIdProfilesDefaultPost**
> V1TranscribeAssetsAssetIdProfilesDefaultPost(ctx, appID, authToken, assetId, optional)
Start a job that sends an asset to default transcription service

 Required roles:  - can_transcribe_content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **assetId** | **string**|  | 
 **optional** | ***DefaultApiV1TranscribeAssetsAssetIdProfilesDefaultPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1TranscribeAssetsAssetIdProfilesDefaultPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of TranscribeSchema**](TranscribeSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TranscribeBulkPost**
> V1TranscribeBulkPost(ctx, appID, authToken, body)
Start a job that sends multiple objects to transcription service

 Required roles:  - can_transcribe_content

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**BulkTranscribeSchema**](BulkTranscribeSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# Go API client for transcode

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./transcode"
```

## Documentation for API Endpoints

All URIs are relative to *https://app.iconik.io/API/transcode*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**V1AnalyzeAssetsAssetIdPost**](docs/DefaultApi.md#v1analyzeassetsassetidpost) | **Post** /v1/analyze/assets/{asset_id}/ | Start a job that sends an asset for analysis
*DefaultApi* | [**V1AnalyzeAssetsAssetIdProfilesDefaultMediaTypePost**](docs/DefaultApi.md#v1analyzeassetsassetidprofilesdefaultmediatypepost) | **Post** /v1/analyze/assets/{asset_id}/profiles/default/{media_type}/ | Start a job that sends an asset for analysis
*DefaultApi* | [**V1AnalyzeAssetsAssetIdProfilesDefaultPost**](docs/DefaultApi.md#v1analyzeassetsassetidprofilesdefaultpost) | **Post** /v1/analyze/assets/{asset_id}/profiles/default/ | Start a job that sends an asset for analysis with a default analysis profile
*DefaultApi* | [**V1AnalyzeAssetsAssetIdProfilesProfileIdPost**](docs/DefaultApi.md#v1analyzeassetsassetidprofilesprofileidpost) | **Post** /v1/analyze/assets/{asset_id}/profiles/{profile_id}/ | Start a job that sends an asset for analysis with a custom analysis profile
*DefaultApi* | [**V1AnalyzeBulkPost**](docs/DefaultApi.md#v1analyzebulkpost) | **Post** /v1/analyze/bulk/ | Start a job that sends objects for analysis using a custom analysis profile
*DefaultApi* | [**V1EdgeTranscodeJobsJobIdAcknowledgePost**](docs/DefaultApi.md#v1edgetranscodejobsjobidacknowledgepost) | **Post** /v1/edge_transcode/jobs/{job_id}/acknowledge/ | Acknowledge an edge transcode job
*DefaultApi* | [**V1EdgeTranscodeWorkersGet**](docs/DefaultApi.md#v1edgetranscodeworkersget) | **Get** /v1/edge_transcode/workers/ | Get edge transcode workers
*DefaultApi* | [**V1EdgeTranscodeWorkersPost**](docs/DefaultApi.md#v1edgetranscodeworkerspost) | **Post** /v1/edge_transcode/workers/ | Create a new edge transcode worker
*DefaultApi* | [**V1EdgeTranscodeWorkersWorkerIdDelete**](docs/DefaultApi.md#v1edgetranscodeworkersworkeriddelete) | **Delete** /v1/edge_transcode/workers/{worker_id}/ | Delete a edge transcode worker
*DefaultApi* | [**V1EdgeTranscodeWorkersWorkerIdGet**](docs/DefaultApi.md#v1edgetranscodeworkersworkeridget) | **Get** /v1/edge_transcode/workers/{worker_id}/ | Get a edge transcode worker
*DefaultApi* | [**V1EdgeTranscodeWorkersWorkerIdPatch**](docs/DefaultApi.md#v1edgetranscodeworkersworkeridpatch) | **Patch** /v1/edge_transcode/workers/{worker_id}/ | Update a edge transcode worker
*DefaultApi* | [**V1EdgeTranscodeWorkersWorkerIdPut**](docs/DefaultApi.md#v1edgetranscodeworkersworkeridput) | **Put** /v1/edge_transcode/workers/{worker_id}/ | Update a edge transcode worker
*DefaultApi* | [**V1KeyframesCollectionsCollectionIdPost**](docs/DefaultApi.md#v1keyframescollectionscollectionidpost) | **Post** /v1/keyframes/collections/{collection_id}/ | Start a job that creates a collection keyframe
*DefaultApi* | [**V1StoragesStorageIdDelete**](docs/DefaultApi.md#v1storagesstorageiddelete) | **Delete** /v1/storages/{storage_id}/ | Cancel all transcode jobs linked to the storage
*DefaultApi* | [**V1StoragesStorageIdEdgeTranscodeJobsGet**](docs/DefaultApi.md#v1storagesstorageidedgetranscodejobsget) | **Get** /v1/storages/{storage_id}/edge_transcode/jobs/ | Get a edge transcode jobs from the job queue
*DefaultApi* | [**V1StoragesStorageIdFilesFileIdTranscodeDelete**](docs/DefaultApi.md#v1storagesstorageidfilesfileidtranscodedelete) | **Delete** /v1/storages/{storage_id}/files/{file_id}/transcode/ | Delete local storage transcode job.
*DefaultApi* | [**V1StoragesStorageIdTranscodeGet**](docs/DefaultApi.md#v1storagesstorageidtranscodeget) | **Get** /v1/storages/{storage_id}/transcode/ | Get pending local storage transcode jobs.
*DefaultApi* | [**V1StoragesStorageIdTranscodeRecordIdDelete**](docs/DefaultApi.md#v1storagesstorageidtranscoderecordiddelete) | **Delete** /v1/storages/{storage_id}/transcode/{record_id}/ | Delete local storage transcode job.
*DefaultApi* | [**V1StoragesStorageIdTranscodeRecordIdGet**](docs/DefaultApi.md#v1storagesstorageidtranscoderecordidget) | **Get** /v1/storages/{storage_id}/transcode/{record_id}/ | Get local storage transcode job.
*DefaultApi* | [**V1TranscodePost**](docs/DefaultApi.md#v1transcodepost) | **Post** /v1/transcode/ | Starts a new transcode.
*DefaultApi* | [**V1TranscodeQueueGet**](docs/DefaultApi.md#v1transcodequeueget) | **Get** /v1/transcode/queue/ | Get all the statuses of the queued transcode jobs
*DefaultApi* | [**V1TranscodeQueueSystemGet**](docs/DefaultApi.md#v1transcodequeuesystemget) | **Get** /v1/transcode/queue/system/ | Get the status of the transcode job queues
*DefaultApi* | [**V1TranscodeTranscodeJobIdDelete**](docs/DefaultApi.md#v1transcodetranscodejobiddelete) | **Delete** /v1/transcode/{transcode_job_id}/ | Cancel a particular transcode job by id
*DefaultApi* | [**V1TranscodeTranscodeJobIdGet**](docs/DefaultApi.md#v1transcodetranscodejobidget) | **Get** /v1/transcode/{transcode_job_id}/ | Get transcode job
*DefaultApi* | [**V1TranscodeTranscodeJobIdPositionPositionPost**](docs/DefaultApi.md#v1transcodetranscodejobidpositionpositionpost) | **Post** /v1/transcode/{transcode_job_id}/position/{position}/ | Move transcode job to top or bottom of the queue
*DefaultApi* | [**V1TranscodeTranscodeJobIdPriorityPriorityPut**](docs/DefaultApi.md#v1transcodetranscodejobidprioritypriorityput) | **Put** /v1/transcode/{transcode_job_id}/priority/{priority}/ | Change transcode job priority
*DefaultApi* | [**V1TranscribeAssetsAssetIdProfilesDefaultPost**](docs/DefaultApi.md#v1transcribeassetsassetidprofilesdefaultpost) | **Post** /v1/transcribe/assets/{asset_id}/profiles/default/ | Start a job that sends an asset to default transcription service
*DefaultApi* | [**V1TranscribeBulkPost**](docs/DefaultApi.md#v1transcribebulkpost) | **Post** /v1/transcribe/bulk/ | Start a job that sends multiple objects to transcription service


## Documentation For Models

 - [AbortStorageTranscodeJobsSchema](docs/AbortStorageTranscodeJobsSchema.md)
 - [AnalyzeSchema](docs/AnalyzeSchema.md)
 - [Body](docs/Body.md)
 - [Body1](docs/Body1.md)
 - [Body2](docs/Body2.md)
 - [BulkActionSchema](docs/BulkActionSchema.md)
 - [BulkAnalyzeSchema](docs/BulkAnalyzeSchema.md)
 - [BulkTranscribeSchema](docs/BulkTranscribeSchema.md)
 - [EdgeThumbnailJobFieldSchema](docs/EdgeThumbnailJobFieldSchema.md)
 - [EdgeTranscodeEndpointSchema](docs/EdgeTranscodeEndpointSchema.md)
 - [EdgeTranscodeInputSchema](docs/EdgeTranscodeInputSchema.md)
 - [EdgeTranscodeJobFieldSchema](docs/EdgeTranscodeJobFieldSchema.md)
 - [EdgeTranscodeJobSchema](docs/EdgeTranscodeJobSchema.md)
 - [EdgeTranscodeJobsSchema](docs/EdgeTranscodeJobsSchema.md)
 - [EdgeTranscodeWorkerSchema](docs/EdgeTranscodeWorkerSchema.md)
 - [EdgeTranscodeWorkersSchema](docs/EdgeTranscodeWorkersSchema.md)
 - [EndpointSchema](docs/EndpointSchema.md)
 - [FacetBucketSchema](docs/FacetBucketSchema.md)
 - [FacetSchema](docs/FacetSchema.md)
 - [GenerateCollectionKeyframeSchema](docs/GenerateCollectionKeyframeSchema.md)
 - [InputSchema](docs/InputSchema.md)
 - [JobBaseSchema](docs/JobBaseSchema.md)
 - [JobSchema](docs/JobSchema.md)
 - [JobStepSchema](docs/JobStepSchema.md)
 - [JobsPrioritySchema](docs/JobsPrioritySchema.md)
 - [JobsStateSchema](docs/JobsStateSchema.md)
 - [ListObjectsSchema](docs/ListObjectsSchema.md)
 - [LocalStorageFileTranscodeJobSchema](docs/LocalStorageFileTranscodeJobSchema.md)
 - [LocalStorageFileTranscodeJobsSchema](docs/LocalStorageFileTranscodeJobsSchema.md)
 - [LocalTranscodeInputSchema](docs/LocalTranscodeInputSchema.md)
 - [LocalTranscodeJobSchema](docs/LocalTranscodeJobSchema.md)
 - [OutputEndpointSchema](docs/OutputEndpointSchema.md)
 - [ReindexQueueRecordSchema](docs/ReindexQueueRecordSchema.md)
 - [SpecifiedKeyframesSchema](docs/SpecifiedKeyframesSchema.md)
 - [ThumbnailJobSchema](docs/ThumbnailJobSchema.md)
 - [TranscodeJobSchema](docs/TranscodeJobSchema.md)
 - [TranscodeQueueObjectSchema](docs/TranscodeQueueObjectSchema.md)
 - [TranscodeQueueSchema](docs/TranscodeQueueSchema.md)
 - [TranscodeValidateMediaInfoSchema](docs/TranscodeValidateMediaInfoSchema.md)
 - [TranscodersSchema](docs/TranscodersSchema.md)
 - [TranscribeSchema](docs/TranscribeSchema.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author




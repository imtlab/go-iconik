# \DefaultApi

All URIs are relative to *https://app.iconik.io/API/jobs*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1JobsDelete**](DefaultApi.md#V1JobsDelete) | **Delete** /v1/jobs/ | Delete multiple jobs by ids list
[**V1JobsGet**](DefaultApi.md#V1JobsGet) | **Get** /v1/jobs/ | Get list of jobs
[**V1JobsJobIdDelete**](DefaultApi.md#V1JobsJobIdDelete) | **Delete** /v1/jobs/{job_id}/ | Delete a particular job by id
[**V1JobsJobIdGet**](DefaultApi.md#V1JobsJobIdGet) | **Get** /v1/jobs/{job_id}/ | Returns a particular job by id
[**V1JobsJobIdPatch**](DefaultApi.md#V1JobsJobIdPatch) | **Patch** /v1/jobs/{job_id}/ | Update job
[**V1JobsJobIdPut**](DefaultApi.md#V1JobsJobIdPut) | **Put** /v1/jobs/{job_id}/ | Update job
[**V1JobsJobIdReindexPost**](DefaultApi.md#V1JobsJobIdReindexPost) | **Post** /v1/jobs/{job_id}/reindex/ | Reindex job
[**V1JobsJobIdStepsJobStepIdPatch**](DefaultApi.md#V1JobsJobIdStepsJobStepIdPatch) | **Patch** /v1/jobs/{job_id}/steps/{job_step_id}/ | Update job step
[**V1JobsJobIdStepsJobStepIdPut**](DefaultApi.md#V1JobsJobIdStepsJobStepIdPut) | **Put** /v1/jobs/{job_id}/steps/{job_step_id}/ | Update job step
[**V1JobsPost**](DefaultApi.md#V1JobsPost) | **Post** /v1/jobs/ | Create a new job
[**V1JobsPriorityPut**](DefaultApi.md#V1JobsPriorityPut) | **Put** /v1/jobs/priority/ | Change jobs priority
[**V1JobsStatePut**](DefaultApi.md#V1JobsStatePut) | **Put** /v1/jobs/state/ | Change jobs state


# **V1JobsDelete**
> V1JobsDelete(ctx, appID, authToken, body)
Delete multiple jobs by ids list

 Required roles:  - can_delete_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**JobsBulkDeleteSchema**](JobsBulkDeleteSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1JobsGet**
> JobsSchema V1JobsGet(ctx, appID, authToken, optional)
Get list of jobs

 Required roles:  - can_read_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
 **optional** | ***DefaultApiV1JobsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiV1JobsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **facets** | **optional.Bool**| If facets should be returned | [default to true]
 **page** | **optional.Int32**| Which page number to fetch | [default to 1]
 **perPage** | **optional.Int32**| The number of items for each page | [default to 10]
 **sort** | **optional.String**| A comma separated list of fieldnames with order. For example - first_name,asc;last_name,desc | 
 **type_** | **optional.String**| Filter by type | 
 **objectType** | **optional.String**| Filter by object_type | 
 **parentId** | **optional.String**| Filter by parent_id | 
 **objectId** | **optional.String**| Filter by object_id | 
 **status** | **optional.String**| Filter by status | 
 **createdBy** | **optional.String**| Filter by created_by | 
 **dateCreated** | **optional.String**| Filter by date_created. Can either be a single ISO8601 timestamp or two timestamps separated by a semicolon &#x60;;&#x60;. The timestamp can also be expressed as number of milliseconds since Jan 1 1970 (epoch). Either timestamp can also be replaced with an asterisk &#x60;*&#x60; to make the query open ended. For example: 2018-01-01T10:00:00Z;2018-01-01T15:00:00Z | 
 **dateModified** | **optional.String**| Filter by date_modified Can either be a single ISO8601 timestamp or two timestamps separated by a semicolon &#x60;;&#x60;. The timestamp can also be expressed as number of milliseconds since Jan 1 1970 (epoch). Either timestamp can also be replaced with an asterisk &#x60;*&#x60; to make the query open ended. For example: *;1544450400 | 
 **query** | **optional.String**| Filter by any of the above with wildcard support | 
 **ids** | **optional.String**| Filter list of id:s (comma separated) | 

### Return type

[**JobsSchema**](JobsSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1JobsJobIdDelete**
> V1JobsJobIdDelete(ctx, appID, authToken, jobId)
Delete a particular job by id

 Required roles:  - can_delete_jobs

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

# **V1JobsJobIdGet**
> JobSchema V1JobsJobIdGet(ctx, appID, authToken, jobId)
Returns a particular job by id

 Required roles:  - can_read_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **jobId** | **string**|  | 

### Return type

[**JobSchema**](JobSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1JobsJobIdPatch**
> JobSchema V1JobsJobIdPatch(ctx, appID, authToken, jobId, body)
Update job

 Required roles:  - can_write_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **jobId** | **string**|  | 
  **body** | [**JobSchema**](JobSchema.md)|  | 

### Return type

[**JobSchema**](JobSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1JobsJobIdPut**
> JobSchema V1JobsJobIdPut(ctx, appID, authToken, jobId, body)
Update job

 Required roles:  - can_write_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **jobId** | **string**|  | 
  **body** | [**JobSchema**](JobSchema.md)|  | 

### Return type

[**JobSchema**](JobSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1JobsJobIdReindexPost**
> V1JobsJobIdReindexPost(ctx, appID, authToken, jobId, body)
Reindex job

 Required roles:  - can_write_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **jobId** | **string**|  | 
  **body** | [**ReindexJobSchema**](ReindexJobSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1JobsJobIdStepsJobStepIdPatch**
> JobSchema V1JobsJobIdStepsJobStepIdPatch(ctx, systemDomainID, appID, authToken, jobId, jobStepId, body)
Update job step

 Required roles:  - can_write_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemDomainID** | **string**|  | 
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **jobId** | **string**|  | 
  **jobStepId** | **string**|  | 
  **body** | [**JobStepSchema**](JobStepSchema.md)|  | 

### Return type

[**JobSchema**](JobSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1JobsJobIdStepsJobStepIdPut**
> JobSchema V1JobsJobIdStepsJobStepIdPut(ctx, systemDomainID, appID, authToken, jobId, jobStepId, body)
Update job step

 Required roles:  - can_write_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemDomainID** | **string**|  | 
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **jobId** | **string**|  | 
  **jobStepId** | **string**|  | 
  **body** | [**JobStepSchema**](JobStepSchema.md)|  | 

### Return type

[**JobSchema**](JobSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1JobsPost**
> JobSchema V1JobsPost(ctx, appID, authToken, body)
Create a new job



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

# **V1JobsPriorityPut**
> V1JobsPriorityPut(ctx, appID, authToken, body)
Change jobs priority

 Required roles:  - can_write_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**JobsPrioritySchema**](JobsPrioritySchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1JobsStatePut**
> V1JobsStatePut(ctx, appID, authToken, body)
Change jobs state

 Required roles:  - can_write_jobs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appID** | **string**|  | 
  **authToken** | **string**|  | 
  **body** | [**JobsStateSchema**](JobsStateSchema.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


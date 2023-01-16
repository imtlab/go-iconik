/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

import (
	"github.com/imtlab/go-iconik/shared"
)

type TransferCloudSchema struct {
	AssetId string `json:"asset_id,omitempty"`
	CeleryTaskId string `json:"celery_task_id,omitempty"`
	ComponentIds []string `json:"component_ids,omitempty"`
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	DeleteOriginal bool `json:"delete_original,omitempty"`
	DestinationDirectoryPath string `json:"destination_directory_path,omitempty"`
	DestinationFilename string `json:"destination_filename,omitempty"`
	DestinationStorageId string `json:"destination_storage_id,omitempty"`
	Id string `json:"id,omitempty"`
	JobId string `json:"job_id,omitempty"`
	JobSteps interface{} `json:"job_steps,omitempty"`
	OriginalFileSetId string `json:"original_file_set_id,omitempty"`
	OriginalStorageId string `json:"original_storage_id,omitempty"`
	OriginalUrl string `json:"original_url,omitempty"`
	ParentJobId string `json:"parent_job_id,omitempty"`
	Priority int32 `json:"priority,omitempty"`
	Status string `json:"status,omitempty"`
	Success string `json:"success,omitempty"`
	TransferType string `json:"transfer_type,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

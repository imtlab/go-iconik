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
	"github.com/imtlab/iconik/shared"
)

type GatewayReportSchema struct {
	AwaitChecksumFilesNumber int32 `json:"await_checksum_files_number,omitempty"`
	DateReported shared.Time `json:"date_reported,omitempty"`
	EmptyFilesNumber int32 `json:"empty_files_number,omitempty"`
	ErrorLogLines []string `json:"error_log_lines,omitempty"`
	ExportedFilesNumber int32 `json:"exported_files_number,omitempty"`
	FaultyFilesNumber int32 `json:"faulty_files_number,omitempty"`
	HostInfo string `json:"host_info,omitempty"`
	HostName string `json:"host_name,omitempty"`
	Id string `json:"id,omitempty"`
	IngestedFilesNumber int32 `json:"ingested_files_number,omitempty"`
	IngestedFilesUploadsNumber int32 `json:"ingested_files_uploads_number,omitempty"`
	IngestingFilesNumber int32 `json:"ingesting_files_number,omitempty"`
	LastScanTime int32 `json:"last_scan_time,omitempty"`
	LogLines []string `json:"log_lines,omitempty"`
	MissingFilesNumber int32 `json:"missing_files_number,omitempty"`
	ScannedFilesNumber int32 `json:"scanned_files_number,omitempty"`
	SkippedFilesNumber int32 `json:"skipped_files_number,omitempty"`
	StartLastDate shared.Time `json:"start_last_date,omitempty"`
	StartStatus string `json:"start_status,omitempty"`
	StartStatusMessage string `json:"start_status_message,omitempty"`
	Status string `json:"status,omitempty"`
	StorageId string `json:"storage_id,omitempty"`
	TotalFilesNumber int32 `json:"total_files_number,omitempty"`
	TotalFoldersNumber int32 `json:"total_folders_number,omitempty"`
	Version string `json:"version,omitempty"`
	WaitingPreviewTranscodeJobsNumber int32 `json:"waiting_preview_transcode_jobs_number,omitempty"`
	WaitingTranscodeJobsNumber int32 `json:"waiting_transcode_jobs_number,omitempty"`
}
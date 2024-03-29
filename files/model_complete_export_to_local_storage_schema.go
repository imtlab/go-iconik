/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type CompleteExportToLocalStorageSchema struct {
	AddFileSet bool `json:"add_file_set,omitempty"`
	AssetId string `json:"asset_id,omitempty"`
	ComponentIds []string `json:"component_ids,omitempty"`
	DeleteRemoteFileSetAfterDownload bool `json:"delete_remote_file_set_after_download,omitempty"`
	DeleteSourceFileSetAfterDownload bool `json:"delete_source_file_set_after_download,omitempty"`
	DestinationDirectoryPath string `json:"destination_directory_path"`
	DestinationFileSetName string `json:"destination_file_set_name"`
	DestinationFilename string `json:"destination_filename,omitempty"`
	ExportMetadataFormat string `json:"export_metadata_format,omitempty"`
	ExportMetadataView string `json:"export_metadata_view,omitempty"`
	ExportPosters bool `json:"export_posters,omitempty"`
	ExportTranscriptionFormat string `json:"export_transcription_format,omitempty"`
	FileSetId string `json:"file_set_id,omitempty"`
	FormatId string `json:"format_id,omitempty"`
	Id string `json:"id,omitempty"`
	JobId string `json:"job_id,omitempty"`
	JobSteps interface{} `json:"job_steps,omitempty"`
	LocalStorageId string `json:"local_storage_id,omitempty"`
	OriginalStorageId string `json:"original_storage_id,omitempty"`
	OriginalUrl string `json:"original_url,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	ParentJobId string `json:"parent_job_id,omitempty"`
	TemporaryFileSetSource bool `json:"temporary_file_set_source,omitempty"`
	TransferType string `json:"transfer_type,omitempty"`
}

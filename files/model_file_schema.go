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

type FileSchema struct {
	AssetId string `json:"asset_id,omitempty"`
	Checksum string `json:"checksum,omitempty"`
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	DirectoryPath string `json:"directory_path"`
	FileDateCreated shared.Time `json:"file_date_created,omitempty"`
	FileDateModified shared.Time `json:"file_date_modified,omitempty"`
	FileSetId string `json:"file_set_id"`
	FileSetStatus string `json:"file_set_status,omitempty"`
	FormatId string `json:"format_id,omitempty"`
	FormatStatus string `json:"format_status,omitempty"`
	Id string `json:"id,omitempty"`
	MultipartUploadUrl string `json:"multipart_upload_url,omitempty"`
	// If not specified, name will be autogenerated based on `id` and `original_name`
	Name string `json:"name,omitempty"`
	OriginalName string `json:"original_name"`
	ParentId string `json:"parent_id,omitempty"`
	Size int64 `json:"size,omitempty"`
	Status string `json:"status,omitempty"`
	StorageId string `json:"storage_id,omitempty"`
	StorageMethod string `json:"storage_method,omitempty"`
	Type_ string `json:"type"`
	UploadCredentials interface{} `json:"upload_credentials,omitempty"`
	UploadFilename string `json:"upload_filename,omitempty"`
	UploadMethod string `json:"upload_method,omitempty"`
	// On a file creation for Backblaze B2 storage in case when request to Backblaze failed, upload_url field will be empty. You can try getting upload_url again by requesting created file with `generate_signed_post_url` set to true.
	UploadUrl string `json:"upload_url,omitempty"`
	Url string `json:"url,omitempty"`
	UserId string `json:"user_id,omitempty"`
	VersionId string `json:"version_id,omitempty"`
}
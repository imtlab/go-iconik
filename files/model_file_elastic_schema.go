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

type FileElasticSchema struct {
	AssetId string `json:"asset_id,omitempty"`
	Checksum string `json:"checksum,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	DateModified string `json:"date_modified,omitempty"`
	DirectoryPath string `json:"directory_path"`
	FileDateCreated shared.Time `json:"file_date_created,omitempty"`
	FileDateModified shared.Time `json:"file_date_modified,omitempty"`
	FileSetId string `json:"file_set_id,omitempty"`
	FileSetStatus string `json:"file_set_status,omitempty"`
	FormatId string `json:"format_id,omitempty"`
	FormatStatus string `json:"format_status,omitempty"`
	Id string `json:"id,omitempty"`
	// If not specified, name will be autogenerated based on `id` and `original_name`
	Name string `json:"name,omitempty"`
	OriginalName string `json:"original_name"`
	ParentId string `json:"parent_id,omitempty"`
	Size int64 `json:"size,omitempty"`
	Status string `json:"status,omitempty"`
	StorageId string `json:"storage_id,omitempty"`
	Type_ string `json:"type"`
	UserId string `json:"user_id,omitempty"`
	VersionId string `json:"version_id,omitempty"`
}

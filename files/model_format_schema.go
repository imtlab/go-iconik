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

type FormatSchema struct {
	ArchiveStatus string `json:"archive_status,omitempty"`
	AssetId string `json:"asset_id,omitempty"`
	Components []ComponentSchema `json:"components,omitempty"`
	DateDeleted shared.Time `json:"date_deleted,omitempty"`
	DeletedByUser string `json:"deleted_by_user,omitempty"`
	Id string `json:"id,omitempty"`
	IsOnline bool `json:"is_online,omitempty"`
	Metadata []interface{} `json:"metadata,omitempty"`
	Name string `json:"name"`
	Status string `json:"status,omitempty"`
	StorageMethods string `json:"storage_methods,omitempty"`
	UserId string `json:"user_id,omitempty"`
	VersionId string `json:"version_id,omitempty"`
}

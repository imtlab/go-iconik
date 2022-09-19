/*
 * iconik_assets
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package assets

import (
	"github.com/imtlab/iconik/shared"
)

type AssetElasticSchema struct {
	AnalyzeStatus string `json:"analyze_status,omitempty"`
	ArchiveStatus string `json:"archive_status,omitempty"`
	Category string `json:"category,omitempty"`
	CreatedByUser string `json:"created_by_user,omitempty"`
	CreatedByUserInfo *UserSchema `json:"created_by_user_info,omitempty"`
	CustomKeyframe string `json:"custom_keyframe,omitempty"`
	CustomPoster string `json:"custom_poster,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	DateDeleted string `json:"date_deleted,omitempty"`
	DateImported shared.Time `json:"date_imported,omitempty"`
	DateModified string `json:"date_modified,omitempty"`
	DeletedByUser string `json:"deleted_by_user,omitempty"`
	DeletedByUserInfo *UserSchema `json:"deleted_by_user_info,omitempty"`
	DurationMilliseconds int64 `json:"duration_milliseconds,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
	Files []interface{} `json:"files,omitempty"`
	Id string `json:"id,omitempty"`
	InCollections []string `json:"in_collections,omitempty"`
	IsBlocked bool `json:"is_blocked,omitempty"`
	IsOnline bool `json:"is_online,omitempty"`
	Keyframes []interface{} `json:"keyframes,omitempty"`
	MediaType string `json:"media_type,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
	Position int32 `json:"position,omitempty"`
	Proxies []interface{} `json:"proxies,omitempty"`
	Relations []RelationElasticSchema `json:"relations,omitempty"`
	Status string `json:"status,omitempty"`
	Title string `json:"title"`
	Type_ string `json:"type,omitempty"`
	UpdatedByUser string `json:"updated_by_user,omitempty"`
	UpdatedByUserInfo *UserSchema `json:"updated_by_user_info,omitempty"`
	Versions []AssetVersionSchema `json:"versions,omitempty"`
	VersionsNumber int32 `json:"versions_number,omitempty"`
	Warning string `json:"warning,omitempty"`
}
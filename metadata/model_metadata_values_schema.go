/*
 * iconik_metadata
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package metadata

import (
	"github.com/imtlab/iconik/shared"
)

type MetadataValuesSchema struct {
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	MetadataValues string `json:"metadata_values"`
	ObjectId string `json:"object_id,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
	VersionId string `json:"version_id,omitempty"`
}

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

type MetadataViewSchema struct {
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	ViewFields []MetadataFieldSchema `json:"view_fields"`
}

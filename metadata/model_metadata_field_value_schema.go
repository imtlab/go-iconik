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

type MetadataFieldValueSchema struct {
	DateCreated shared.Time `json:"date_created,omitempty"`
	FieldValues []interface{} `json:"field_values,omitempty"`
	Mode string `json:"mode,omitempty"`
}
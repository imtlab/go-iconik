/*
 * iconik_transcode
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package transcode

import (
	"github.com/imtlab/iconik/shared"
)

type JobStepSchema struct {
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateUpdated shared.Time `json:"date_updated,omitempty"`
	Id string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
}
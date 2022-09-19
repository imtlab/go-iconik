/*
 * iconik_acls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package acls

import (
	"github.com/imtlab/iconik/shared"
)

type UserAclSchema struct {
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	ObjectKey string `json:"object_key,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
	Permissions []string `json:"permissions"`
	UserId string `json:"user_id,omitempty"`
}

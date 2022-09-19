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

type ApprovalSchema struct {
	ChangeDate shared.Time `json:"change_date,omitempty"`
	Externals string `json:"externals,omitempty"`
	Groups string `json:"groups,omitempty"`
	MinNumber int32 `json:"min_number,omitempty"`
	ObjectId string `json:"object_id,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
	RequestDate shared.Time `json:"request_date,omitempty"`
	RequestedBy string `json:"requested_by,omitempty"`
	ShareId string `json:"share_id,omitempty"`
	Status string `json:"status,omitempty"`
	UserStatus string `json:"user_status,omitempty"`
	Users string `json:"users,omitempty"`
	UsersInfo []UserSchema `json:"users_info,omitempty"`
	VersionId string `json:"version_id,omitempty"`
}
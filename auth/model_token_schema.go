/*
 * iconik_auth
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package auth

import (
	"github.com/imtlab/go-iconik/shared"
)

type TokenSchema struct {
	AppId string `json:"app_id,omitempty"`
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	Expires shared.Time `json:"expires,omitempty"`
	Id string `json:"id,omitempty"`
	IsAdmin bool `json:"is_admin,omitempty"`
	IsSuperAdmin bool `json:"is_super_admin,omitempty"`
	SystemDomainId string `json:"system_domain_id,omitempty"`
	SystemDomainStatus string `json:"system_domain_status,omitempty"`
	SystemDomainWarningMessage string `json:"system_domain_warning_message,omitempty"`
	SystemDomains []string `json:"system_domains,omitempty"`
	Token string `json:"token"`
	UserId string `json:"user_id,omitempty"`
}

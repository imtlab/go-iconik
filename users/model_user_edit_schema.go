/*
 * iconik_users
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package users

import (
	"github.com/imtlab/go-iconik/shared"
)

type UserEditSchema struct {
	CurrentPassword string `json:"current_password,omitempty"`
	DateTermsAccepted shared.Time `json:"date_terms_accepted,omitempty"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	Groups []string `json:"groups,omitempty"`
	HideEmail bool `json:"hide_email,omitempty"`
	HidePhone bool `json:"hide_phone,omitempty"`
	Id string `json:"id,omitempty"`
	IdentityProviderId string `json:"identity_provider_id,omitempty"`
	IsAdmin bool `json:"is_admin,omitempty"`
	IsSuperAdmin bool `json:"is_super_admin,omitempty"`
	LastName string `json:"last_name,omitempty"`
	LastSuccessfulAuth shared.Time `json:"last_successful_auth,omitempty"`
	LastUnsuccessfulAuth shared.Time `json:"last_unsuccessful_auth,omitempty"`
	LastWebLogin shared.Time `json:"last_web_login,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	Password string `json:"password"`
	Phone string `json:"phone,omitempty"`
	Photo string `json:"photo,omitempty"`
	PhotoBig string `json:"photo_big,omitempty"`
	PhotoSmall string `json:"photo_small,omitempty"`
	PrimaryGroup string `json:"primary_group,omitempty"`
	Status string `json:"status,omitempty"`
	SystemDomains []string `json:"system_domains,omitempty"`
	SystemMetadata *UserBaseSchemaSystemMetadata `json:"system_metadata,omitempty"`
	SystemName string `json:"system_name,omitempty"`
	Type_ string `json:"type"`
}

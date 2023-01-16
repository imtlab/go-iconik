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

type SystemDomainFromReferralCodeSchema struct {
	AdminEmail string `json:"admin_email"`
	AdminFirstName string `json:"admin_first_name"`
	AdminLastName string `json:"admin_last_name,omitempty"`
	AdminPassword string `json:"admin_password"`
	CountryCode string `json:"country_code"`
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
}

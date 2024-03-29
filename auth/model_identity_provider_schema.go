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

type IdentityProviderSchema struct {
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	Id string `json:"id,omitempty"`
	PublicId string `json:"public_id,omitempty"`
	SamlSettings string `json:"saml_settings,omitempty"`
	Settings string `json:"settings"`
	Type_ string `json:"type"`
	VerboseLogging bool `json:"verbose_logging,omitempty"`
}

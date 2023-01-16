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

type RegistrationSchema struct {
	BaseUrl string `json:"base_url,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	Country string `json:"country"`
	DateCreated shared.Time `json:"date_created,omitempty"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	Id string `json:"id,omitempty"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	ReferralCode string `json:"referral_code,omitempty"`
	StripeId string `json:"stripe_id,omitempty"`
}

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

type SystemDomainSchema struct {
	BaseUrl string `json:"base_url"`
	CustomTerms bool `json:"custom_terms,omitempty"`
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	DeactivateDate shared.Time `json:"deactivate_date,omitempty"`
	Description string `json:"description,omitempty"`
	DiscountPercent float32 `json:"discount_percent,omitempty"`
	DoNotChargeEdgeTranscoder bool `json:"do_not_charge_edge_transcoder,omitempty"`
	DoNotChargeRemoteProxies bool `json:"do_not_charge_remote_proxies,omitempty"`
	FreezeDate shared.Time `json:"freeze_date,omitempty"`
	Id string `json:"id,omitempty"`
	InvoiceEndOfMonth bool `json:"invoice_end_of_month,omitempty"`
	IsTemplate bool `json:"is_template,omitempty"`
	Name string `json:"name"`
	ReferralCode string `json:"referral_code,omitempty"`
	Secret string `json:"secret,omitempty"`
	Status string `json:"status,omitempty"`
	StripeId string `json:"stripe_id,omitempty"`
	Type_ string `json:"type,omitempty"`
	WarningMessage string `json:"warning_message,omitempty"`
}

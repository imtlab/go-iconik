/*
 * iconik_auth
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package auth

type OktaSettingsSchema struct {
	CertFingerprint string `json:"cert_fingerprint,omitempty"`
	CertFingerprintAlgorithm string `json:"cert_fingerprint_algorithm,omitempty"`
	DigestAlgorithm string `json:"digest_algorithm,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
	IdpX509cert string `json:"idp_x509cert,omitempty"`
	OktaAppId string `json:"okta_app_id,omitempty"`
	OktaName string `json:"okta_name"`
	OktaPreview bool `json:"okta_preview,omitempty"`
	OktaSso string `json:"okta_sso,omitempty"`
	SignatureAlgorithm string `json:"signature_algorithm,omitempty"`
}

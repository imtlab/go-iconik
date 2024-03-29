/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type AnalysisProfileSchema struct {
	AnalysisServiceAccountId string `json:"analysis_service_account_id"`
	Enabled bool `json:"enabled,omitempty"`
	Id string `json:"id,omitempty"`
	IsDefault bool `json:"is_default,omitempty"`
	MediaType string `json:"media_type"`
	Name string `json:"name"`
	ServiceType string `json:"service_type,omitempty"`
	Settings string `json:"settings,omitempty"`
}

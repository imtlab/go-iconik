/*
 * iconik_auth
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package auth

type ExternalAuthRequestSchema struct {
	AppId string `json:"app_id,omitempty"`
	AppName string `json:"app_name,omitempty"`
	RedirectInfo *ExternalAuthRequestSchemaRedirectInfo `json:"redirect_info,omitempty"`
	Secret string `json:"secret"`
}

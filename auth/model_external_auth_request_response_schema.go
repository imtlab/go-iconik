/*
 * iconik_auth
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package auth

type ExternalAuthRequestResponseSchema struct {
	AppId string `json:"app_id,omitempty"`
	RedirectInfo *RedirectInfoTypeSchema `json:"redirect_info,omitempty"`
}
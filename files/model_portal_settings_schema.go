/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type PortalSettingsSchema struct {
	Delete bool `json:"delete,omitempty"`
	IsSystem bool `json:"is_system,omitempty"`
	Read bool `json:"read,omitempty"`
	Scan bool `json:"scan,omitempty"`
	Write bool `json:"write,omitempty"`
}

/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type SftpSettingsSchema struct {
	Address string `json:"address"`
	Delete bool `json:"delete,omitempty"`
	IsSystem bool `json:"is_system,omitempty"`
	Password string `json:"password"`
	Read bool `json:"read,omitempty"`
	Scan bool `json:"scan,omitempty"`
	User string `json:"user"`
	Write bool `json:"write,omitempty"`
}
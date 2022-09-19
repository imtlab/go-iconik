/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type VantageSettingsSchema struct {
	Host string `json:"host"`
	Local bool `json:"local,omitempty"`
	Port int32 `json:"port"`
	ShareName string `json:"share_name"`
	TempProxyFolder string `json:"temp_proxy_folder"`
	WorkflowId string `json:"workflow_id"`
}

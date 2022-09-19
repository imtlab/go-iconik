/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type TelestreamCloudSchema struct {
	AccessKey string `json:"access_key"`
	ApiHost string `json:"api_host,omitempty"`
	ApiPort int32 `json:"api_port,omitempty"`
	FactoryId string `json:"factory_id"`
	KeyframesProfileId string `json:"keyframes_profile_id"`
	Local bool `json:"local,omitempty"`
	ProxyProfileId string `json:"proxy_profile_id"`
	SecretKey string `json:"secret_key"`
	StorageId string `json:"storage_id"`
}

/*
 * iconik_transcode
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package transcode

type EndpointSchema struct {
	Data interface{} `json:"data,omitempty"`
	Headers interface{} `json:"headers,omitempty"`
	Method string `json:"method,omitempty"`
	StorageMethod string `json:"storage_method,omitempty"`
	Type_ string `json:"type,omitempty"`
	Url string `json:"url"`
}

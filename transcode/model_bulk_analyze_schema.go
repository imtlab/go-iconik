/*
 * iconik_transcode
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package transcode

type BulkAnalyzeSchema struct {
	Force bool `json:"force,omitempty"`
	ObjectIds []string `json:"object_ids"`
	ObjectType string `json:"object_type"`
	ProfileId string `json:"profile_id,omitempty"`
}

/*
 * iconik_transcode
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package transcode

type JobBaseSchema struct {
	AssetId string `json:"asset_id,omitempty"`
	CollectionId string `json:"collection_id,omitempty"`
	Input *InputSchema `json:"input"`
	JobId string `json:"job_id,omitempty"`
	JobSteps []JobStepSchema `json:"job_steps,omitempty"`
	MediaInfo string `json:"media_info,omitempty"`
	Thumbnail []ThumbnailJobSchema `json:"thumbnail,omitempty"`
	Transcode []TranscodeJobSchema `json:"transcode,omitempty"`
}
/*
 * iconik_transcode
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package transcode

type GenerateCollectionKeyframeSchema struct {
	DeletedAssetId string `json:"deleted_asset_id,omitempty"`
	Force bool `json:"force,omitempty"`
	SpecifiedAssetIds []string `json:"specified_asset_ids,omitempty"`
	SpecifiedKeyframes []SpecifiedKeyframesSchema `json:"specified_keyframes,omitempty"`
}
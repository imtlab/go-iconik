/*
 * iconik_assets
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package assets

type RelationElasticSchema struct {
	DateCreated string `json:"date_created,omitempty"`
	DateModified string `json:"date_modified,omitempty"`
	Description string `json:"description,omitempty"`
	RelatedFromAssetId string `json:"related_from_asset_id,omitempty"`
	RelatedToAssetId string `json:"related_to_asset_id,omitempty"`
	RelationType string `json:"relation_type,omitempty"`
}
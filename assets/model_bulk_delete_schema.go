/*
 * iconik_assets
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package assets

type BulkDeleteSchema struct {
	// If set to `False`, will also delete entities of type `object_type` specified in `object_ids`.
	ContentOnly bool `json:"content_only,omitempty"`
	ObjectIds []string `json:"object_ids"`
	ObjectType string `json:"object_type"`
}

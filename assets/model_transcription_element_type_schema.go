/*
 * iconik_assets
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package assets

type TranscriptionElementTypeSchema struct {
	EndMs int64 `json:"end_ms"`
	Score float32 `json:"score,omitempty"`
	StartMs int64 `json:"start_ms"`
	Value string `json:"value"`
}

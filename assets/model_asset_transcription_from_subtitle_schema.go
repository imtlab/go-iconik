/*
 * iconik_assets
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package assets

type AssetTranscriptionFromSubtitleSchema struct {
	Content string `json:"content,omitempty"`
	DeleteOldTranscriptions bool `json:"delete_old_transcriptions,omitempty"`
	Format string `json:"format,omitempty"`
	Language string `json:"language,omitempty"`
	// Set to source subtitle_id or do not set and use the content fields instead
	SourceSubtitleId string `json:"source_subtitle_id,omitempty"`
}

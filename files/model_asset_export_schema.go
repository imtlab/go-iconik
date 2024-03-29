/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type AssetExportSchema struct {
	ExportMetadata bool `json:"export_metadata,omitempty"`
	ExportToAssetFolder bool `json:"export_to_asset_folder,omitempty"`
	FileName string `json:"file_name,omitempty"`
	FormatId string `json:"format_id,omitempty"`
	MetadataFormat string `json:"metadata_format,omitempty"`
	MetadataView string `json:"metadata_view,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	TranscodeProfileIds []string `json:"transcode_profile_ids,omitempty"`
}

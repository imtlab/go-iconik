/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type BulkFilesetExportSchema struct {
	AllowDuplicateTransfers bool `json:"allow_duplicate_transfers,omitempty"`
	DeleteOriginal bool `json:"delete_original,omitempty"`
	ExportMetadata bool `json:"export_metadata,omitempty"`
	ExportToAssetFolder bool `json:"export_to_asset_folder,omitempty"`
	KeepCollectionStructure bool `json:"keep_collection_structure,omitempty"`
	MetadataFormat string `json:"metadata_format,omitempty"`
	MetadataView string `json:"metadata_view,omitempty"`
	ObjectIds []string `json:"object_ids"`
	ObjectType string `json:"object_type"`
	Overwrite bool `json:"overwrite,omitempty"`
}

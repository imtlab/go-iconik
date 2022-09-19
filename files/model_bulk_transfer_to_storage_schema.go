/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type BulkTransferToStorageSchema struct {
	AllowDuplicateTransfers bool `json:"allow_duplicate_transfers,omitempty"`
	DeleteOriginal bool `json:"delete_original,omitempty"`
	FilePath string `json:"file_path,omitempty"`
	FormatName string `json:"format_name,omitempty"`
	KeepCollectionStructure bool `json:"keep_collection_structure,omitempty"`
	ObjectIds []string `json:"object_ids"`
	ObjectType string `json:"object_type"`
	Overwrite bool `json:"overwrite,omitempty"`
}

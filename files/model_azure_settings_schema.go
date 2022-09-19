/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type AzureSettingsSchema struct {
	AccessGroupId string `json:"access_group_id,omitempty"`
	AclTemplateId string `json:"acl_template_id,omitempty"`
	AddUuidToFilenames bool `json:"add_uuid_to_filenames,omitempty"`
	AggregateIdenticalFiles bool `json:"aggregate_identical_files,omitempty"`
	AggregateIgnore string `json:"aggregate_ignore,omitempty"`
	AggregateOnlyOnSameStorage bool `json:"aggregate_only_on_same_storage,omitempty"`
	ConnectionString string `json:"connection_string"`
	Container string `json:"container"`
	Delete bool `json:"delete,omitempty"`
	EnableCollectionDirectoryMapping bool `json:"enable_collection_directory_mapping,omitempty"`
	FilenameIsExternalId bool `json:"filename_is_external_id,omitempty"`
	FolderNameTagsMetadataFieldName string `json:"folder_name_tags_metadata_field_name,omitempty"`
	FolderNameTagsMetadataViewId string `json:"folder_name_tags_metadata_view_id,omitempty"`
	IsSystem bool `json:"is_system,omitempty"`
	MetadataViewId string `json:"metadata_view_id,omitempty"`
	Path string `json:"path"`
	PreloadEdgeJobs int32 `json:"preload_edge_jobs,omitempty"`
	Read bool `json:"read,omitempty"`
	RootCollectionId string `json:"root_collection_id,omitempty"`
	Scan bool `json:"scan,omitempty"`
	ScanDirectories string `json:"scan_directories,omitempty"`
	ScanIgnore string `json:"scan_ignore,omitempty"`
	SidecarMetadataRequired bool `json:"sidecar_metadata_required,omitempty"`
	TitleIncludesExtension bool `json:"title_includes_extension,omitempty"`
	TranscodeIgnore string `json:"transcode_ignore,omitempty"`
	Write bool `json:"write,omitempty"`
}

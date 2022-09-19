/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type S3SettingsSchema struct {
	AccessGroupId string `json:"access_group_id,omitempty"`
	AccessKey string `json:"access_key"`
	AclTemplateId string `json:"acl_template_id,omitempty"`
	AddUuidToFilenames bool `json:"add_uuid_to_filenames,omitempty"`
	AggregateIdenticalFiles bool `json:"aggregate_identical_files,omitempty"`
	AggregateIgnore string `json:"aggregate_ignore,omitempty"`
	AggregateOnlyOnSameStorage bool `json:"aggregate_only_on_same_storage,omitempty"`
	Bucket string `json:"bucket"`
	Delete bool `json:"delete,omitempty"`
	EnableCollectionDirectoryMapping bool `json:"enable_collection_directory_mapping,omitempty"`
	Endpoint string `json:"endpoint"`
	FilenameIsExternalId bool `json:"filename_is_external_id,omitempty"`
	FolderNameTagsMetadataFieldName string `json:"folder_name_tags_metadata_field_name,omitempty"`
	FolderNameTagsMetadataViewId string `json:"folder_name_tags_metadata_view_id,omitempty"`
	// Keep restored assets online for this many days to allow them to be copied before going back to Glacier
	GlacierRestoreTimeout int32 `json:"glacier_restore_timeout,omitempty"`
	IsSystem bool `json:"is_system,omitempty"`
	MetadataViewId string `json:"metadata_view_id,omitempty"`
	NotificationSqsUrl string `json:"notification_sqs_url,omitempty"`
	Path string `json:"path"`
	PreloadEdgeJobs int32 `json:"preload_edge_jobs,omitempty"`
	Read bool `json:"read,omitempty"`
	Region string `json:"region"`
	RootCollectionId string `json:"root_collection_id,omitempty"`
	Scan bool `json:"scan,omitempty"`
	ScanDirectories string `json:"scan_directories,omitempty"`
	ScanIgnore string `json:"scan_ignore,omitempty"`
	SecretKey string `json:"secret_key"`
	SidecarMetadataRequired bool `json:"sidecar_metadata_required,omitempty"`
	TitleIncludesExtension bool `json:"title_includes_extension,omitempty"`
	TranscodeIgnore string `json:"transcode_ignore,omitempty"`
	UseAcceleration bool `json:"use_acceleration,omitempty"`
	Write bool `json:"write,omitempty"`
}
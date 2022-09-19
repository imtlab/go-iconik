/*
 * iconik_settings
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package settings

type MergedSettingsSchema struct {
	AclTemplateId string `json:"acl_template_id,omitempty"`
	AssetDefaultSections []string `json:"asset_default_sections,omitempty"`
	CorsHosts []string `json:"cors_hosts,omitempty"`
	CustomTerms bool `json:"custom_terms,omitempty"`
	DateFormat string `json:"date_format,omitempty"`
	DatetimeFormat string `json:"datetime_format,omitempty"`
	DefaultUploadStorageId string `json:"default_upload_storage_id,omitempty"`
	// Grace period that indicate how long objects will live in recycle bin. Unit: hours
	DeleteGracePeriod int32 `json:"delete_grace_period,omitempty"`
	ExternalShare bool `json:"external_share,omitempty"`
	FacetFields []FacetFieldSchema `json:"facet_fields,omitempty"`
	FiltersDefaultMetadataViewId string `json:"filters_default_metadata_view_id,omitempty"`
	GroupId string `json:"group_id,omitempty"`
	HomePage string `json:"home_page,omitempty"`
	ImagePropertiesMetadataField string `json:"image_properties_metadata_field,omitempty"`
	JobsDashboard *GroupSettingPublicSchemaJobsDashboard `json:"jobs_dashboard,omitempty"`
	LocationsMetadataField string `json:"locations_metadata_field,omitempty"`
	LogoStorageId string `json:"logo_storage_id,omitempty"`
	LogoUrl string `json:"logo_url,omitempty"`
	LogosMetadataField string `json:"logos_metadata_field,omitempty"`
	MaxBrowseUsers int32 `json:"max_browse_users,omitempty"`
	MaxPowerUsers int32 `json:"max_power_users,omitempty"`
	MaxStandardUsers int32 `json:"max_standard_users,omitempty"`
	MaxStorageBytes int64 `json:"max_storage_bytes,omitempty"`
	MaxTrafficBytes int64 `json:"max_traffic_bytes,omitempty"`
	PasswordChecks *GroupSettingPublicSchemaJobsDashboard `json:"password_checks,omitempty"`
	RequiredMetadataViews []string `json:"required_metadata_views,omitempty"`
	SafeSearchesMetadataField string `json:"safe_searches_metadata_field,omitempty"`
	SamlRequireGroups bool `json:"saml_require_groups,omitempty"`
	SearchDefaultSections []string `json:"search_default_sections,omitempty"`
	SearchDisplayFields []SearchDisplayFieldSchema `json:"search_display_fields,omitempty"`
	SearchInTranscriptions bool `json:"search_in_transcriptions,omitempty"`
	// Default share expiration time that indicate how long share will be valid. Unit: days
	ShareExpirationTime int32 `json:"share_expiration_time,omitempty"`
	SupportAccess bool `json:"support_access,omitempty"`
	SystemDomainId string `json:"system_domain_id,omitempty"`
	TagsMetadataField string `json:"tags_metadata_field,omitempty"`
	TextsMetadataField string `json:"texts_metadata_field,omitempty"`
}

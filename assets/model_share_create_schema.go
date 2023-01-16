/*
 * iconik_assets
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package assets

import (
	"github.com/imtlab/go-iconik/shared"
)

type ShareCreateSchema struct {
	AllowApprovingComments bool `json:"allow_approving_comments"`
	AllowComments bool `json:"allow_comments"`
	AllowCustomActions bool `json:"allow_custom_actions,omitempty"`
	AllowDownload bool `json:"allow_download"`
	AllowDownloadProxy bool `json:"allow_download_proxy,omitempty"`
	AllowSettingApproveStatus bool `json:"allow_setting_approve_status"`
	AllowUpload bool `json:"allow_upload,omitempty"`
	AllowViewTranscriptions bool `json:"allow_view_transcriptions,omitempty"`
	AllowViewVersions bool `json:"allow_view_versions,omitempty"`
	AppId string `json:"app_id,omitempty"`
	Emails []string `json:"emails"`
	Expires shared.Time `json:"expires,omitempty"`
	Id string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	MetadataViews []string `json:"metadata_views,omitempty"`
	ObjectId string `json:"object_id,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
	OwnerId string `json:"owner_id,omitempty"`
	Password string `json:"password,omitempty"`
	SystemDomainId string `json:"system_domain_id,omitempty"`
	UploadStorageId string `json:"upload_storage_id,omitempty"`
}

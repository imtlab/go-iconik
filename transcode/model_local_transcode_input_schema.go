/*
 * iconik_transcode
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package transcode

type LocalTranscodeInputSchema struct {
	AssetId string `json:"asset_id"`
	Checksum string `json:"checksum,omitempty"`
	ClosedCaptions bool `json:"closed_captions,omitempty"`
	CollectionId string `json:"collection_id,omitempty"`
	ComponentIds []string `json:"component_ids,omitempty"`
	DirectoryPath string `json:"directory_path"`
	Endpoint *EndpointSchema `json:"endpoint"`
	FileId string `json:"file_id"`
	FileSetId string `json:"file_set_id"`
	Filename string `json:"filename"`
	FormatId string `json:"format_id"`
	Language string `json:"language,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
	ProxyId string `json:"proxy_id,omitempty"`
	Size int64 `json:"size"`
	StorageId string `json:"storage_id"`
	UpdateProxyMediainfo bool `json:"update_proxy_mediainfo,omitempty"`
	VersionId string `json:"version_id"`
}

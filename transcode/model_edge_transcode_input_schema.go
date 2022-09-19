/*
 * iconik_transcode
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package transcode

type EdgeTranscodeInputSchema struct {
	AssetId string `json:"asset_id,omitempty"`
	ClosedCaptions bool `json:"closed_captions,omitempty"`
	Endpoint *EdgeTranscodeEndpointSchema `json:"endpoint"`
	FileId string `json:"file_id,omitempty"`
	FileSetId string `json:"file_set_id,omitempty"`
	FormatId string `json:"format_id,omitempty"`
	Language string `json:"language,omitempty"`
	OriginalName string `json:"original_name,omitempty"`
	ProxyId string `json:"proxy_id,omitempty"`
	StorageId string `json:"storage_id,omitempty"`
}

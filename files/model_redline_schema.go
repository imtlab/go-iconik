/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type RedlineSchema struct {
	EditProxyUploadStorageId string `json:"edit_proxy_upload_storage_id,omitempty"`
	EditProxyUploadStoragePath string `json:"edit_proxy_upload_storage_path,omitempty"`
	Format string `json:"format,omitempty"`
	KeepRedlineProxy bool `json:"keep_redline_proxy,omitempty"`
	Local bool `json:"local,omitempty"`
	OpenclDeviceIndexes string `json:"opencl_device_indexes,omitempty"`
	QtCodec string `json:"qt_codec,omitempty"`
}
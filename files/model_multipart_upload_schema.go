/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type MultipartUploadSchema struct {
	PartsGroup int64 `json:"parts_group,omitempty"`
	PartsRange []string `json:"parts_range"`
}
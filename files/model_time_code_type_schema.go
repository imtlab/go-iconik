/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

type TimeCodeTypeSchema struct {
	FramesNumber int64 `json:"frames_number,omitempty"`
	IsDropFrame bool `json:"is_drop_frame,omitempty"`
	TimeBase *TimeBaseTypeSchema `json:"time_base,omitempty"`
}
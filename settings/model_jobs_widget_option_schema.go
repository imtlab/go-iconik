/*
 * iconik_settings
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package settings

type JobsWidgetOptionSchema struct {
	Columns []string `json:"columns,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	Sort []SortSchema `json:"sort,omitempty"`
}
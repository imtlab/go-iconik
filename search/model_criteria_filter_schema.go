/*
 * iconik_search
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package search

type CriteriaFilterSchema struct {
	Filters []CriteriaFilterSchemaFilters1 `json:"filters,omitempty"`
	Operator string `json:"operator"`
	Terms []CriteriaTermSchema `json:"terms,omitempty"`
}
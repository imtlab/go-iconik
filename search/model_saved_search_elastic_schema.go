/*
 * iconik_search
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package search

type SavedSearchElasticSchema struct {
	Criteria string `json:"criteria,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
	DateModified string `json:"date_modified,omitempty"`
	GroupId string `json:"group_id,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Permissions []string `json:"permissions,omitempty"`
}

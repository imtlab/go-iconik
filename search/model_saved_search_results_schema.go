/*
 * iconik_search
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package search

type SavedSearchResultsSchema struct {
	Facets interface{} `json:"facets,omitempty"`
	FirstUrl string `json:"first_url,omitempty"`
	Id string `json:"id,omitempty"`
	LastUrl string `json:"last_url,omitempty"`
	Name string `json:"name,omitempty"`
	NextUrl string `json:"next_url,omitempty"`
	Objects []SearchDocumentSchema `json:"objects,omitempty"`
	Page int32 `json:"page,omitempty"`
	Pages int32 `json:"pages,omitempty"`
	PerPage int32 `json:"per_page,omitempty"`
	PrevUrl string `json:"prev_url,omitempty"`
	SearchCriteriaDocument *SavedSearchSchema `json:"search_criteria_document,omitempty"`
	Total int32 `json:"total,omitempty"`
}

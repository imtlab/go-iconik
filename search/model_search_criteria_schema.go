/*
 * iconik_search
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package search

type SearchCriteriaSchema struct {
	DocTypes []string `json:"doc_types,omitempty"`
	ExcludeFields []string `json:"exclude_fields,omitempty"`
	Facets []string `json:"facets,omitempty"`
	FacetsFilters []FacetFilterSchema `json:"facets_filters,omitempty"`
	Filter *CriteriaFilterSchema `json:"filter,omitempty"`
	IncludeFields []string `json:"include_fields,omitempty"`
	MetadataViewId string `json:"metadata_view_id,omitempty"`
	Query string `json:"query,omitempty"`
	SearchFields []string `json:"search_fields,omitempty"`
	Sort []CriteriaSortSchema `json:"sort,omitempty"`
}

/*
 * iconik_stats
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stats

type ListObjectsSchema struct {
	FirstUrl string `json:"first_url,omitempty"`
	LastUrl string `json:"last_url,omitempty"`
	NextUrl string `json:"next_url,omitempty"`
	Page int32 `json:"page,omitempty"`
	Pages int32 `json:"pages,omitempty"`
	PerPage int32 `json:"per_page,omitempty"`
	PrevUrl string `json:"prev_url,omitempty"`
	Total int64 `json:"total,omitempty"`
}

/*
 * iconik_transcode
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package transcode

type FacetSchema struct {
	Buckets []FacetBucketSchema `json:"buckets,omitempty"`
	DocCountErrorUpperBound int64 `json:"doc_count_error_upper_bound,omitempty"`
	SumOtherDocCount int64 `json:"sum_other_doc_count,omitempty"`
}

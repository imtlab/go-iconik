/*
 * iconik_acls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package acls

type BulkAcLsObjectSchema struct {
	ObjectKeys []string `json:"object_keys"`
	ObjectType string `json:"object_type"`
	Permissions []string `json:"permissions"`
}

/*
 * iconik_acls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package acls

type CreateAcLsSchema struct {
	GroupIds []string `json:"group_ids,omitempty"`
	Mode string `json:"mode,omitempty"`
	// The number of object_keys in the list is limited to a minimum of 0 and a maximum of 500
	ObjectKeys []string `json:"object_keys"`
	ObjectType string `json:"object_type,omitempty"`
	Permissions []string `json:"permissions"`
	UserIds []string `json:"user_ids,omitempty"`
}
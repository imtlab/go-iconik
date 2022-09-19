/*
 * iconik_users
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package users

type GroupBaseSchema struct {
	Alias string `json:"alias,omitempty"`
	DefaultUserType string `json:"default_user_type,omitempty"`
	Description string `json:"description,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
	Id string `json:"id,omitempty"`
	IsSamlGroup bool `json:"is_saml_group,omitempty"`
	Logo string `json:"logo,omitempty"`
	Name string `json:"name"`
	RoleCategories *GroupBaseSchemaRoleCategories `json:"role_categories,omitempty"`
	Roles []string `json:"roles,omitempty"`
	SamlPrimaryGroupPriority int32 `json:"saml_primary_group_priority,omitempty"`
}

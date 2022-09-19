/*
 * iconik_acls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package acls

type AclSchema struct {
	GroupsAcl []GroupAclSchema `json:"groups_acl,omitempty"`
	InheritedGroupsAcl []GroupAclSchema `json:"inherited_groups_acl,omitempty"`
	InheritedUsersAcl []UserAclSchema `json:"inherited_users_acl,omitempty"`
	PropagatingGroupsAcl []GroupAclSchema `json:"propagating_groups_acl,omitempty"`
	PropagatingUsersAcl []UserAclSchema `json:"propagating_users_acl,omitempty"`
	UsersAcl []UserAclSchema `json:"users_acl,omitempty"`
}

/*
 * iconik_users_notifications
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package usersnotifications

import (
	"github.com/imtlab/go-iconik/shared"
)

type NotificationSchema struct {
	Context interface{} `json:"context,omitempty"`
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	EventType string `json:"event_type"`
	ExcludeUsers []string `json:"exclude_users,omitempty"`
	Id string `json:"id,omitempty"`
	MessageLong string `json:"message_long"`
	MessageShort string `json:"message_short"`
	ObjectId string `json:"object_id"`
	ObjectType string `json:"object_type"`
	RecipientId string `json:"recipient_id"`
	SenderId string `json:"sender_id"`
	ShareId string `json:"share_id,omitempty"`
	ShareUserId string `json:"share_user_id,omitempty"`
	Status string `json:"status,omitempty"`
	SubObjectId string `json:"sub_object_id,omitempty"`
	SubObjectType string `json:"sub_object_type,omitempty"`
	SystemDomainId string `json:"system_domain_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

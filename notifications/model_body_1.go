/*
 * iconik_webhooks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package notifications

type Body1 struct {
	// Type of events. Example assets, collections
	EventType string `json:"event_type"`
	// ID of a particular object you want to subscribe on
	ObjectId string `json:"object_id,omitempty"`
	// Realm of event. Example entity, contents, acls, metadata
	Realm string `json:"realm,omitempty"`
	// Operation of event. Example create, update, delete
	Operation string `json:"operation,omitempty"`
	// URL you want to be called when notification is appeared
	Url string `json:"url"`
	// Define the key-value pairs your third party provider requires here
	Headers interface{} `json:"headers,omitempty"`
}

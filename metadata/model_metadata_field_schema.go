/*
 * iconik_metadata
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package metadata

import (
	"github.com/imtlab/go-iconik/shared"
)

type MetadataFieldSchema struct {
	AutoSet bool `json:"auto_set,omitempty"`
	DateCreated shared.Time `json:"date_created,omitempty"`
	DateModified shared.Time `json:"date_modified,omitempty"`
	Description string `json:"description,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
	FieldType string `json:"field_type"`
	HideIfNotSet bool `json:"hide_if_not_set,omitempty"`
	IsBlockField bool `json:"is_block_field,omitempty"`
	IsWarningField bool `json:"is_warning_field,omitempty"`
	Label string `json:"label,omitempty"`
	MappedFieldName string `json:"mapped_field_name,omitempty"`
	MaxValue float32 `json:"max_value,omitempty"`
	MinValue float32 `json:"min_value,omitempty"`
	Multi bool `json:"multi,omitempty"`
	Name string `json:"name,omitempty"`
	Options []FieldOptionsSchema `json:"options,omitempty"`
	ReadOnly bool `json:"read_only,omitempty"`
	Representative bool `json:"representative,omitempty"`
	Required bool `json:"required,omitempty"`
	Sortable bool `json:"sortable,omitempty"`
	// Will be used to upload MetadataField's `options`. Cannot be set or used as long as `options` are set.  **Example**: The value is `https://external-url.com/foo/`. In that case `GET` request will be sent to `https://external-url.com/foo/?user_id=uuid1&view_id=uuid1&field_name=bar&view_name=user_view&system_domain_id=uuid1`. Please note that some query parameters were added by *iconik* to get values that were predefined in your system for each user [user_id] and view [view_id]. Metadata field name [field_name], view's name [view_name] and system domain identifier [system_domain_id] will be also passed in each request. *iconik* will successfully parse the response from that URL if it will be sent in JSON formatted string: `{\"bar\": [{\"value\": \"1\", \"label\": \"1st\"}, {\"value\": \"2\", \"label\": \"2nd\"}]}`
	SourceUrl string `json:"source_url,omitempty"`
	UseAsFacet bool `json:"use_as_facet,omitempty"`
}

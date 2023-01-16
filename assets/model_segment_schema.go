/*
 * iconik_assets
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package assets

import (
	"github.com/imtlab/go-iconik/shared"
)

type SegmentSchema struct {
	AssetId string `json:"asset_id,omitempty"`
	DateCreated shared.Time `json:"date_created,omitempty"`
	Drawing *DrawingSchema `json:"drawing,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
	HasDrawing bool `json:"has_drawing,omitempty"`
	Id string `json:"id,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	MetadataViewId string `json:"metadata_view_id,omitempty"`
	ParentId string `json:"parent_id,omitempty"`
	Path string `json:"path,omitempty"`
	SegmentChecked bool `json:"segment_checked,omitempty"`
	SegmentColor string `json:"segment_color,omitempty"`
	SegmentText string `json:"segment_text,omitempty"`
	SegmentTrack string `json:"segment_track,omitempty"`
	SegmentType string `json:"segment_type"`
	ShareUserEmail string `json:"share_user_email,omitempty"`
	Status string `json:"status,omitempty"`
	TimeEndMilliseconds int64 `json:"time_end_milliseconds,omitempty"`
	TimeStartMilliseconds int64 `json:"time_start_milliseconds,omitempty"`
	TopLevel bool `json:"top_level,omitempty"`
	Transcription *TranscriptionTypeSchema `json:"transcription,omitempty"`
	TranscriptionId string `json:"transcription_id,omitempty"`
	UserFirstName string `json:"user_first_name,omitempty"`
	UserId string `json:"user_id,omitempty"`
	UserInfo *UserSchema `json:"user_info,omitempty"`
	UserLastName string `json:"user_last_name,omitempty"`
	UserPhoto string `json:"user_photo,omitempty"`
	VersionId string `json:"version_id,omitempty"`
}

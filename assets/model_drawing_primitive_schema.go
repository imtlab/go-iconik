/*
 * iconik_assets
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package assets

type DrawingPrimitiveSchema struct {
	Color string `json:"color,omitempty"`
	Points []DrawingPointSchema `json:"points,omitempty"`
	Text string `json:"text,omitempty"`
	Type_ string `json:"type"`
}
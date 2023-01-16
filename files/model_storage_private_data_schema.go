/*
 * iconik_files
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package files

import (
	"github.com/imtlab/go-iconik/shared"
)

type StoragePrivateDataSchema struct {
	BucketLocation string `json:"bucket_location,omitempty"`
	Default_ bool `json:"default,omitempty"`
	Id string `json:"id,omitempty"`
	LastScanned shared.Time `json:"last_scanned,omitempty"`
	Method string `json:"method"`
	Name string `json:"name"`
	Purpose string `json:"purpose"`
	ScannerStatus string `json:"scanner_status,omitempty"`
	Settings string `json:"settings"`
	Status string `json:"status,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	Version string `json:"version,omitempty"`
}

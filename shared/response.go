/*
 * iconik_shared
 *
 * iconik client code shared by all API packages
 *
 * API version: 1.0
 */

package shared

import (
	"net/http"
)

type APIResponse struct {
	*http.Response		`json:"-"`
	Message		string	`json:"message,omitempty"`
	//	Operation is the name of the swagger operation.
	Operation	string	`json:"operation,omitempty"`
	//	RequestURL is the request URL.
	//	This value is always available, even if the embedded *http.Response is nil.
	RequestURL	string	`json:"url,omitempty"`
	//	Method is the HTTP method used for the request.
	//	This value is always available, even if the embedded *http.Response is nil.
	Method		string	`json:"method,omitempty"`
	//	Payload holds the contents of the response body (which may be nil or empty).
	//	This is provided here as the raw response.Body() reader will have already been drained.
	Payload		[]byte	`json:"-"`
}

func NewAPIResponse(pResponse *http.Response) *APIResponse {
	return &APIResponse{Response: pResponse}
}

func NewAPIResponseWithError(errorMessage string) *APIResponse {
	return &APIResponse{Message: errorMessage}
}

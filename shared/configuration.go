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

//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//	BEGIN COMMENT OUT

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//\\//	END COMMENT OUT

type Configuration struct {
	BasePath		string				`json:"basePath,omitempty"`
	Host			string				`json:"host,omitempty"`
	Scheme			string				`json:"scheme,omitempty"`
	DefaultHeader	map[string]string	`json:"defaultHeader,omitempty"`
	UserAgent		string				`json:"userAgent,omitempty"`
	HTTPClient		*http.Client
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

/*
NewConfiguration Create a new Configuration object with appID and authToken in its DefaultHeader map
 * @param appID
 * @param authToken
*/
func NewConfiguration(appID, authToken string) *Configuration {
	pConfiguration := &Configuration {
		BasePath:		"https://app.iconik.io/API/",
		DefaultHeader:	make(map[string]string),
		UserAgent:		"Swagger-Codegen/1.0.0/go",
	}

	pConfiguration.AddDefaultHeader("App-ID", appID)
	pConfiguration.AddDefaultHeader("Auth-Token", authToken)

	return pConfiguration
}


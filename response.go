/*
 * Cloudmailer
 *
 * Reference for the public Cloudmailer API calls.  Generated packages use semantic versioning since 2.0.0. The endpoints (v1) itself should always be backwards compatible.  The API uses rate limiting based on requests per minute. Every API response contains the following two headers: 'X-RateLimit-Limit', 'X-RateLimit-Remaining'. `X-RateLimit-Limit` indicates the request limit per minute. `X-RateLimit-Remaining` indicates the amount of available requests.
 *
 * API version: 3.0.0
 * Contact: support@slash2.nl
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudmailer

import (
	"net/http"
)

// APIResponse stores the API response returned by the server.
type APIResponse struct {
	*http.Response `json:"-"`
	Message        string `json:"message,omitempty"`
	// Operation is the name of the OpenAPI operation.
	Operation string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

// NewAPIResponse returns a new APIResonse object.
func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

// NewAPIResponseWithError returns a new APIResponse object with the provided error message.
func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}

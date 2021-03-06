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
// GroupMemberIndexResponse struct for GroupMemberIndexResponse
type GroupMemberIndexResponse struct {
	CurrentPage int32 `json:"current_page,omitempty"`
	Data []GroupMemberIndexResponseData `json:"data,omitempty"`
	FirstPageUrl string `json:"first_page_url,omitempty"`
	From int32 `json:"from,omitempty"`
	LastPage int32 `json:"last_page,omitempty"`
	LastPageUrl string `json:"last_page_url,omitempty"`
	NextPageUrl string `json:"next_page_url,omitempty"`
	Path string `json:"path,omitempty"`
	PerPage int32 `json:"per_page,omitempty"`
	PrevPageUrl string `json:"prev_page_url,omitempty"`
	To int32 `json:"to,omitempty"`
	Total int32 `json:"total,omitempty"`
	Sortable []string `json:"sortable,omitempty"`
}

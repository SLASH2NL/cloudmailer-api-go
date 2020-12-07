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
// MemberResponseData struct for MemberResponseData
type MemberResponseData struct {
	Id int32 `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	IsBlacklisted bool `json:"is_blacklisted,omitempty"`
	ValidatedAt string `json:"validated_at,omitempty"`
	RequestedValidationAt string `json:"requested_validation_at,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	MetaFields []GroupMemberIndexResponseMetaFields `json:"meta_fields,omitempty"`
	Groups []GroupIndexResponseData `json:"groups,omitempty"`
}

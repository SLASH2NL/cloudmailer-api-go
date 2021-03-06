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
// AutomationShowNameResponseData struct for AutomationShowNameResponseData
type AutomationShowNameResponseData struct {
	Id int32 `json:"id,omitempty"`
	InternName string `json:"intern_name,omitempty"`
	Name string `json:"name,omitempty"`
	Subject string `json:"subject,omitempty"`
	PreviewText string `json:"preview_text,omitempty"`
	TotalMailsSent int32 `json:"total_mails_sent,omitempty"`
	WizardStepName string `json:"wizard_step_name,omitempty"`
	ClientId int32 `json:"client_id,omitempty"`
	Locale string `json:"locale,omitempty"`
	TestingEmail string `json:"testing_email,omitempty"`
	HasOnlineView bool `json:"has_online_view,omitempty"`
	HasUnsubscribe bool `json:"has_unsubscribe,omitempty"`
	HasTracking bool `json:"has_tracking,omitempty"`
	HasLabel bool `json:"has_label,omitempty"`
	HasUtm bool `json:"has_utm,omitempty"`
	IsTesting bool `json:"is_testing,omitempty"`
	IsPublished bool `json:"is_published,omitempty"`
	IsSystemDefault bool `json:"is_system_default,omitempty"`
	Replacements []ResponseAutomationReplacements `json:"replacements,omitempty"`
	ViewUrl string `json:"view_url,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

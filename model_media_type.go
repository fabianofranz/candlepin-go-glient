/*
 * Candlepin
 *
 * Candlepin is subscription management server written in Java. It helps with management of software subscriptions.
 *
 * API version: 2.3.10
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MediaType struct {
	Type string `json:"type,omitempty"`
	Subtype string `json:"subtype,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
	WildcardType bool `json:"wildcardType,omitempty"`
	WildcardSubtype bool `json:"wildcardSubtype,omitempty"`
}

/*
 * Candlepin
 *
 * Candlepin is subscription management server written in Java. It helps with management of software subscriptions.
 *
 * API version: 2.3.10
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ComplianceReason struct {
	Key string `json:"key,omitempty"`
	Message string `json:"message,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}
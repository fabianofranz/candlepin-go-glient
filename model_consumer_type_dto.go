/*
 * Candlepin
 *
 * Candlepin is subscription management server written in Java. It helps with management of software subscriptions.
 *
 * API version: 2.3.10
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DTO representing a consumer type
type ConsumerTypeDto struct {
	Id string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Manifest bool `json:"manifest,omitempty"`
}

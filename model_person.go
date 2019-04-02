/*
 * Candlepin
 *
 * Candlepin is subscription management server written in Java. It helps with management of software subscriptions.
 *
 * API version: 2.3.10
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Person struct {
	ExtensionAttributes map[string]map[string]interface{} `json:"extensionAttributes,omitempty"`
	Lang string `json:"lang,omitempty"`
	Base string `json:"base,omitempty"`
	Name string `json:"name,omitempty"`
	Uri string `json:"uri,omitempty"`
	Email string `json:"email,omitempty"`
}
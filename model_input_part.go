/*
 * Candlepin
 *
 * Candlepin is subscription management server written in Java. It helps with management of software subscriptions.
 *
 * API version: 2.3.10
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InputPart struct {
	BodyAsString string `json:"bodyAsString,omitempty"`
	ContentTypeFromMessage bool `json:"contentTypeFromMessage,omitempty"`
	MediaType MediaType `json:"mediaType,omitempty"`
	Headers map[string][]string `json:"headers,omitempty"`
}

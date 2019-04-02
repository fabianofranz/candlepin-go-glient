/*
 * Candlepin
 *
 * Candlepin is subscription management server written in Java. It helps with management of software subscriptions.
 *
 * API version: 2.3.10
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

// DTO representing a GuestIdDTO
type GuestIdDto struct {
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Id string `json:"id,omitempty"`
	GuestId string `json:"guestId,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

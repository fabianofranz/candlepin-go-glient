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

type PermissionBlueprint struct {
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Id string `json:"id"`
	Owner NestedOwner `json:"owner"`
	Access string `json:"access,omitempty"`
	Type string `json:"type"`
}

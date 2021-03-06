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

type Role struct {
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Id string `json:"id"`
	Users []User `json:"users,omitempty"`
	Permissions []PermissionBlueprint `json:"permissions,omitempty"`
	Name string `json:"name"`
	Href string `json:"href,omitempty"`
}

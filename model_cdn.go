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

type Cdn struct {
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Id string `json:"id"`
	Label string `json:"label"`
	Name string `json:"name"`
	Url string `json:"url"`
	Certificate CdnCertificate `json:"certificate,omitempty"`
}

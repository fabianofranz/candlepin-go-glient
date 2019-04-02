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

// DTO representing an upstream consumer
type UpstreamConsumerDto struct {
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Id string `json:"id,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	ApiUrl string `json:"apiUrl,omitempty"`
	WebUrl string `json:"webUrl,omitempty"`
	OwnerId string `json:"ownerId,omitempty"`
	ContentAccessMode string `json:"contentAccessMode,omitempty"`
	Type ConsumerTypeDto `json:"type,omitempty"`
	IdCert CertificateDto `json:"idCert,omitempty"`
}

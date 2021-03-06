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

type ConsumerInstalledProduct struct {
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Id string `json:"id"`
	ProductId string `json:"productId"`
	ProductName string `json:"productName,omitempty"`
	Version string `json:"version,omitempty"`
	Arch string `json:"arch,omitempty"`
	Status string `json:"status,omitempty"`
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate time.Time `json:"endDate,omitempty"`
}

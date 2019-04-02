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

type DateRange struct {
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate time.Time `json:"endDate,omitempty"`
}
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

type ComplianceStatus struct {
	Date time.Time `json:"date,omitempty"`
	CompliantUntil time.Time `json:"compliantUntil,omitempty"`
	NonCompliantProducts []string `json:"nonCompliantProducts,omitempty"`
	CompliantProducts map[string][]Entitlement `json:"compliantProducts,omitempty"`
	PartiallyCompliantProducts map[string][]Entitlement `json:"partiallyCompliantProducts,omitempty"`
	PartialStacks map[string][]Entitlement `json:"partialStacks,omitempty"`
	ProductComplianceDateRanges map[string]DateRange `json:"productComplianceDateRanges,omitempty"`
	Reasons []ComplianceReason `json:"reasons,omitempty"`
	Status string `json:"status,omitempty"`
	Compliant bool `json:"compliant,omitempty"`
}

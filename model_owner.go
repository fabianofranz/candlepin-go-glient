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

type Owner struct {
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	ParentOwner NestedOwner `json:"parentOwner,omitempty"`
	Id string `json:"id"`
	Key string `json:"key"`
	DisplayName string `json:"displayName"`
	ContentPrefix string `json:"contentPrefix,omitempty"`
	LastRefreshed time.Time `json:"lastRefreshed,omitempty"`
	DefaultServiceLevel string `json:"defaultServiceLevel,omitempty"`
	UpstreamConsumer UpstreamConsumer `json:"upstreamConsumer,omitempty"`
	LogLevel string `json:"logLevel,omitempty"`
	ContentAccessMode string `json:"contentAccessMode,omitempty"`
	ContentAccessModeList string `json:"contentAccessModeList,omitempty"`
	Href string `json:"href,omitempty"`
	AutobindDisabled bool `json:"autobindDisabled,omitempty"`
}

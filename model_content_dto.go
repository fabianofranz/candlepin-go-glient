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

// DTO representing a certificate
type ContentDto struct {
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Id string `json:"id"`
	Type string `json:"type,omitempty"`
	Label string `json:"label,omitempty"`
	Name string `json:"name,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	ContentUrl string `json:"contentUrl,omitempty"`
	RequiredTags string `json:"requiredTags,omitempty"`
	GpgUrl string `json:"gpgUrl,omitempty"`
	ModifiedProductIds []string `json:"modifiedProductIds,omitempty"`
	Arches string `json:"arches,omitempty"`
	MetadataExpire int64 `json:"metadataExpire,omitempty"`
	ReleaseVer string `json:"releaseVer,omitempty"`
}

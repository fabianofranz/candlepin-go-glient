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

type Source struct {
	Title string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Categories []Category `json:"categories,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Id string `json:"id,omitempty"`
	Links []Link `json:"links,omitempty"`
	Author []Person `json:"author,omitempty"`
	Contributor []Person `json:"contributor,omitempty"`
	Rights string `json:"rights,omitempty"`
	Icon string `json:"icon,omitempty"`
	Logo string `json:"logo,omitempty"`
	Generator Generator `json:"generator,omitempty"`
	ExtensionAttributes map[string]map[string]interface{} `json:"extensionAttributes,omitempty"`
	Lang string `json:"lang,omitempty"`
	Base string `json:"base,omitempty"`
}

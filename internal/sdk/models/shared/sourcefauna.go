// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceFaunaSchemasCollectionDeletionMode string

const (
	SourceFaunaSchemasCollectionDeletionModeDeletedField SourceFaunaSchemasCollectionDeletionMode = "deleted_field"
)

func (e SourceFaunaSchemasCollectionDeletionMode) ToPointer() *SourceFaunaSchemasCollectionDeletionMode {
	return &e
}
func (e *SourceFaunaSchemasCollectionDeletionMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "deleted_field":
		*e = SourceFaunaSchemasCollectionDeletionMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFaunaSchemasCollectionDeletionMode: %v", v)
	}
}

type SourceFaunaEnabled struct {
	// Name of the "deleted at" column.
	Column       *string                                  `default:"deleted_at" json:"column"`
	deletionMode SourceFaunaSchemasCollectionDeletionMode `const:"deleted_field" json:"deletion_mode"`
}

func (s SourceFaunaEnabled) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFaunaEnabled) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFaunaEnabled) GetColumn() *string {
	if o == nil {
		return nil
	}
	return o.Column
}

func (o *SourceFaunaEnabled) GetDeletionMode() SourceFaunaSchemasCollectionDeletionMode {
	return SourceFaunaSchemasCollectionDeletionModeDeletedField
}

type SourceFaunaSchemasDeletionMode string

const (
	SourceFaunaSchemasDeletionModeIgnore SourceFaunaSchemasDeletionMode = "ignore"
)

func (e SourceFaunaSchemasDeletionMode) ToPointer() *SourceFaunaSchemasDeletionMode {
	return &e
}
func (e *SourceFaunaSchemasDeletionMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ignore":
		*e = SourceFaunaSchemasDeletionMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFaunaSchemasDeletionMode: %v", v)
	}
}

type SourceFaunaDisabled struct {
	deletionMode SourceFaunaSchemasDeletionMode `const:"ignore" json:"deletion_mode"`
}

func (s SourceFaunaDisabled) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFaunaDisabled) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFaunaDisabled) GetDeletionMode() SourceFaunaSchemasDeletionMode {
	return SourceFaunaSchemasDeletionModeIgnore
}

type SourceFaunaDeletionModeType string

const (
	SourceFaunaDeletionModeTypeSourceFaunaDisabled SourceFaunaDeletionModeType = "source-fauna_Disabled"
	SourceFaunaDeletionModeTypeSourceFaunaEnabled  SourceFaunaDeletionModeType = "source-fauna_Enabled"
)

// SourceFaunaDeletionMode - <b>This only applies to incremental syncs.</b> <br>
// Enabling deletion mode informs your destination of deleted documents.<br>
// Disabled - Leave this feature disabled, and ignore deleted documents.<br>
// Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.
type SourceFaunaDeletionMode struct {
	SourceFaunaDisabled *SourceFaunaDisabled
	SourceFaunaEnabled  *SourceFaunaEnabled

	Type SourceFaunaDeletionModeType
}

func CreateSourceFaunaDeletionModeSourceFaunaDisabled(sourceFaunaDisabled SourceFaunaDisabled) SourceFaunaDeletionMode {
	typ := SourceFaunaDeletionModeTypeSourceFaunaDisabled

	return SourceFaunaDeletionMode{
		SourceFaunaDisabled: &sourceFaunaDisabled,
		Type:                typ,
	}
}

func CreateSourceFaunaDeletionModeSourceFaunaEnabled(sourceFaunaEnabled SourceFaunaEnabled) SourceFaunaDeletionMode {
	typ := SourceFaunaDeletionModeTypeSourceFaunaEnabled

	return SourceFaunaDeletionMode{
		SourceFaunaEnabled: &sourceFaunaEnabled,
		Type:               typ,
	}
}

func (u *SourceFaunaDeletionMode) UnmarshalJSON(data []byte) error {

	var sourceFaunaDisabled SourceFaunaDisabled = SourceFaunaDisabled{}
	if err := utils.UnmarshalJSON(data, &sourceFaunaDisabled, "", true, true); err == nil {
		u.SourceFaunaDisabled = &sourceFaunaDisabled
		u.Type = SourceFaunaDeletionModeTypeSourceFaunaDisabled
		return nil
	}

	var sourceFaunaEnabled SourceFaunaEnabled = SourceFaunaEnabled{}
	if err := utils.UnmarshalJSON(data, &sourceFaunaEnabled, "", true, true); err == nil {
		u.SourceFaunaEnabled = &sourceFaunaEnabled
		u.Type = SourceFaunaDeletionModeTypeSourceFaunaEnabled
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceFaunaDeletionMode) MarshalJSON() ([]byte, error) {
	if u.SourceFaunaDisabled != nil {
		return utils.MarshalJSON(u.SourceFaunaDisabled, "", true)
	}

	if u.SourceFaunaEnabled != nil {
		return utils.MarshalJSON(u.SourceFaunaEnabled, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SourceFaunaCollection - Settings for the Fauna Collection.
type SourceFaunaCollection struct {
	// <b>This only applies to incremental syncs.</b> <br>
	// Enabling deletion mode informs your destination of deleted documents.<br>
	// Disabled - Leave this feature disabled, and ignore deleted documents.<br>
	// Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.
	Deletions SourceFaunaDeletionMode `json:"deletions"`
	// The page size used when reading documents from the database. The larger the page size, the faster the connector processes documents. However, if a page is too large, the connector may fail. <br>
	// Choose your page size based on how large the documents are. <br>
	// See <a href="https://docs.fauna.com/fauna/current/learn/understanding/types#page">the docs</a>.
	PageSize *int64 `default:"64" json:"page_size"`
}

func (s SourceFaunaCollection) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFaunaCollection) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceFaunaCollection) GetDeletions() SourceFaunaDeletionMode {
	if o == nil {
		return SourceFaunaDeletionMode{}
	}
	return o.Deletions
}

func (o *SourceFaunaCollection) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type Fauna string

const (
	FaunaFauna Fauna = "fauna"
)

func (e Fauna) ToPointer() *Fauna {
	return &e
}
func (e *Fauna) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fauna":
		*e = Fauna(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Fauna: %v", v)
	}
}

type SourceFauna struct {
	// Settings for the Fauna Collection.
	Collection *SourceFaunaCollection `json:"collection,omitempty"`
	// Domain of Fauna to query. Defaults db.fauna.com. See <a href=https://docs.fauna.com/fauna/current/learn/understanding/region_groups#how-to-use-region-groups>the docs</a>.
	Domain *string `default:"db.fauna.com" json:"domain"`
	// Endpoint port.
	Port *int64 `default:"443" json:"port"`
	// URL scheme.
	Scheme *string `default:"https" json:"scheme"`
	// Fauna secret, used when authenticating with the database.
	Secret     string `json:"secret"`
	sourceType Fauna  `const:"fauna" json:"sourceType"`
}

func (s SourceFauna) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFauna) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceFauna) GetCollection() *SourceFaunaCollection {
	if o == nil {
		return nil
	}
	return o.Collection
}

func (o *SourceFauna) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *SourceFauna) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceFauna) GetScheme() *string {
	if o == nil {
		return nil
	}
	return o.Scheme
}

func (o *SourceFauna) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}

func (o *SourceFauna) GetSourceType() Fauna {
	return FaunaFauna
}

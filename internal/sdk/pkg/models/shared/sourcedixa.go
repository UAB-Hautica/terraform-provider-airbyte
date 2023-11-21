// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type Dixa string

const (
	DixaDixa Dixa = "dixa"
)

func (e Dixa) ToPointer() *Dixa {
	return &e
}

func (e *Dixa) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dixa":
		*e = Dixa(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Dixa: %v", v)
	}
}

type SourceDixa struct {
	// Dixa API token
	APIToken string `json:"api_token"`
	// Number of days to batch into one request. Max 31.
	BatchSize  *int64 `default:"31" json:"batch_size"`
	sourceType Dixa   `const:"dixa" json:"sourceType"`
	// The connector pulls records updated from this date onwards.
	StartDate time.Time `json:"start_date"`
}

func (s SourceDixa) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceDixa) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceDixa) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceDixa) GetBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.BatchSize
}

func (o *SourceDixa) GetSourceType() Dixa {
	return DixaDixa
}

func (o *SourceDixa) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

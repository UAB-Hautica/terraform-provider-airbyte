// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type KustomerSinger string

const (
	KustomerSingerKustomerSinger KustomerSinger = "kustomer-singer"
)

func (e KustomerSinger) ToPointer() *KustomerSinger {
	return &e
}

func (e *KustomerSinger) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kustomer-singer":
		*e = KustomerSinger(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KustomerSinger: %v", v)
	}
}

type SourceKustomerSinger struct {
	// Kustomer API Token. See the <a href="https://developer.kustomer.com/kustomer-api-docs/reference/authentication">docs</a> on how to obtain this
	APIToken   string         `json:"api_token"`
	sourceType KustomerSinger `const:"kustomer-singer" json:"sourceType"`
	// The date from which you'd like to replicate the data
	StartDate string `json:"start_date"`
}

func (s SourceKustomerSinger) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceKustomerSinger) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceKustomerSinger) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceKustomerSinger) GetSourceType() KustomerSinger {
	return KustomerSingerKustomerSinger
}

func (o *SourceKustomerSinger) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}

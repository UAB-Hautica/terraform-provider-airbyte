// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Zoom string

const (
	ZoomZoom Zoom = "zoom"
)

func (e Zoom) ToPointer() *Zoom {
	return &e
}

func (e *Zoom) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zoom":
		*e = Zoom(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Zoom: %v", v)
	}
}

type SourceZoom struct {
	// JWT Token
	JwtToken   string `json:"jwt_token"`
	sourceType Zoom   `const:"zoom" json:"sourceType"`
}

func (s SourceZoom) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZoom) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZoom) GetJwtToken() string {
	if o == nil {
		return ""
	}
	return o.JwtToken
}

func (o *SourceZoom) GetSourceType() Zoom {
	return ZoomZoom
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Metabase string

const (
	MetabaseMetabase Metabase = "metabase"
)

func (e Metabase) ToPointer() *Metabase {
	return &e
}
func (e *Metabase) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "metabase":
		*e = Metabase(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Metabase: %v", v)
	}
}

type SourceMetabase struct {
	// URL to your metabase instance API
	InstanceAPIURL string  `json:"instance_api_url"`
	Password       *string `json:"password,omitempty"`
	// To generate your session token, you need to run the following command: ``` curl -X POST \
	//   -H "Content-Type: application/json" \
	//   -d '{"username": "person@metabase.com", "password": "fakepassword"}' \
	//   http://localhost:3000/api/session
	// ``` Then copy the value of the `id` field returned by a successful call to that API.
	// Note that by default, sessions are good for 14 days and needs to be regenerated.
	SessionToken *string  `json:"session_token,omitempty"`
	sourceType   Metabase `const:"metabase" json:"sourceType"`
	Username     *string  `json:"username,omitempty"`
}

func (s SourceMetabase) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMetabase) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMetabase) GetInstanceAPIURL() string {
	if o == nil {
		return ""
	}
	return o.InstanceAPIURL
}

func (o *SourceMetabase) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceMetabase) GetSessionToken() *string {
	if o == nil {
		return nil
	}
	return o.SessionToken
}

func (o *SourceMetabase) GetSourceType() Metabase {
	return MetabaseMetabase
}

func (o *SourceMetabase) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

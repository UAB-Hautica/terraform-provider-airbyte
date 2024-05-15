// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type PaypalTransaction string

const (
	PaypalTransactionPaypalTransaction PaypalTransaction = "paypal-transaction"
)

func (e PaypalTransaction) ToPointer() *PaypalTransaction {
	return &e
}
func (e *PaypalTransaction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "paypal-transaction":
		*e = PaypalTransaction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaypalTransaction: %v", v)
	}
}

type SourcePaypalTransaction struct {
	// The Client ID of your Paypal developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Paypal developer application.
	ClientSecret string `json:"client_secret"`
	// Start Date parameter for the list dispute endpoint in <a href=\"https://datatracker.ietf.org/doc/html/rfc3339#section-5.6\">ISO format</a>. This Start Date must be in range within 180 days before present time, and requires ONLY 3 miliseconds(mandatory). If you don't use this option, it defaults to a start date set 180 days in the past.
	DisputeStartDate *time.Time `json:"dispute_start_date,omitempty"`
	// End Date for data extraction in <a href=\"https://datatracker.ietf.org/doc/html/rfc3339#section-5.6\">ISO format</a>. This can be help you select specific range of time, mainly for test purposes  or data integrity tests. When this is not used, now_utc() is used by the streams. This does not apply to Disputes and Product streams.
	EndDate *time.Time `json:"end_date,omitempty"`
	// Determines whether to use the sandbox or production environment.
	IsSandbox *bool `default:"false" json:"is_sandbox"`
	// The key to refresh the expired access token.
	RefreshToken *string           `json:"refresh_token,omitempty"`
	sourceType   PaypalTransaction `const:"paypal-transaction" json:"sourceType"`
	// Start Date for data extraction in <a href=\"https://datatracker.ietf.org/doc/html/rfc3339#section-5.6\">ISO format</a>. Date must be in range from 3 years till 12 hrs before present time.
	StartDate time.Time `json:"start_date"`
	// The number of days per request. Must be a number between 1 and 31.
	TimeWindow *int64 `default:"7" json:"time_window"`
}

func (s SourcePaypalTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePaypalTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePaypalTransaction) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourcePaypalTransaction) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourcePaypalTransaction) GetDisputeStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.DisputeStartDate
}

func (o *SourcePaypalTransaction) GetEndDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourcePaypalTransaction) GetIsSandbox() *bool {
	if o == nil {
		return nil
	}
	return o.IsSandbox
}

func (o *SourcePaypalTransaction) GetRefreshToken() *string {
	if o == nil {
		return nil
	}
	return o.RefreshToken
}

func (o *SourcePaypalTransaction) GetSourceType() PaypalTransaction {
	return PaypalTransactionPaypalTransaction
}

func (o *SourcePaypalTransaction) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourcePaypalTransaction) GetTimeWindow() *int64 {
	if o == nil {
		return nil
	}
	return o.TimeWindow
}

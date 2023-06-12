// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package setplanmodifier

import (
	"airbyte/internal/planmodifiers/utils"
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// TrustRead returns a plan modifier that propagates a state value into the planned value, when it is Known, and the Plan Value is Unknown
func TrustRead() planmodifier.Set {
	return trustRead{}
}

// trustRead implements the plan modifier.
type trustRead struct{}

// Description returns a human-readable description of the plan modifier.
func (m trustRead) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m trustRead) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifySet implements the plan modification logic.
func (m trustRead) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// Do nothing if there is a known planned value.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value
	if req.ConfigValue.IsUnknown() {
		return
	}

	if utils.IsAllStateUnknown(ctx, req.State) {
		return
	}

	resp.PlanValue = req.StateValue
}

package utils

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func IsAllStateUnknown(ctx context.Context, state tfsdk.State) bool {
	attrs := state.Schema.GetAttributes()
	anyFound := false
	for k, _ := range attrs {
		attrValue := new(attr.Value)
		state.GetAttribute(ctx, path.Root(k), attrValue)
		if attrValue != nil && !(*attrValue).IsUnknown() && !(*attrValue).IsNull() {
			anyFound = true
			break
		}
	}

	return !anyFound
}

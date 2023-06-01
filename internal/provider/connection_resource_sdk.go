// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ConnectionResourceModel) ToCreateSDKType() *shared.ConnectionCreateRequest {
	var configurations *shared.StreamConfigurations
	if r.Configurations != nil {
		streams := make([]shared.StreamConfiguration, 0)
		for _, streamsItem := range r.Configurations.Streams {
			cursorField := make([]string, 0)
			for _, cursorFieldItem := range streamsItem.CursorField {
				cursorField = append(cursorField, cursorFieldItem.ValueString())
			}
			name := streamsItem.Name.ValueString()
			primaryKey := make([][]string, 0)
			for _, primaryKeyItem := range streamsItem.PrimaryKey {
				primaryKeyTmp := make([]string, 0)
				for _, item := range primaryKeyItem {
					primaryKeyTmp = append(primaryKeyTmp, item.ValueString())
				}
				primaryKey = append(primaryKey, primaryKeyTmp)
			}
			syncMode := new(shared.ConnectionSyncModeEnum)
			if !streamsItem.SyncMode.IsUnknown() && !streamsItem.SyncMode.IsNull() {
				*syncMode = shared.ConnectionSyncModeEnum(streamsItem.SyncMode.ValueString())
			} else {
				syncMode = nil
			}
			streams = append(streams, shared.StreamConfiguration{
				CursorField: cursorField,
				Name:        name,
				PrimaryKey:  primaryKey,
				SyncMode:    syncMode,
			})
		}
		configurations = &shared.StreamConfigurations{
			Streams: streams,
		}
	}
	dataResidency := new(shared.GeographyEnum)
	if !r.DataResidency.IsUnknown() && !r.DataResidency.IsNull() {
		*dataResidency = shared.GeographyEnum(r.DataResidency.ValueString())
	} else {
		dataResidency = nil
	}
	destinationID := r.DestinationID.ValueString()
	name1 := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name1 = r.Name.ValueString()
	} else {
		name1 = nil
	}
	namespaceDefinition := new(shared.NamespaceDefinitionEnum)
	if !r.NamespaceDefinition.IsUnknown() && !r.NamespaceDefinition.IsNull() {
		*namespaceDefinition = shared.NamespaceDefinitionEnum(r.NamespaceDefinition.ValueString())
	} else {
		namespaceDefinition = nil
	}
	namespaceFormat := new(string)
	if !r.NamespaceFormat.IsUnknown() && !r.NamespaceFormat.IsNull() {
		*namespaceFormat = r.NamespaceFormat.ValueString()
	} else {
		namespaceFormat = nil
	}
	nonBreakingSchemaUpdatesBehavior := new(shared.NonBreakingSchemaUpdatesBehaviorEnum)
	if !r.NonBreakingSchemaUpdatesBehavior.IsUnknown() && !r.NonBreakingSchemaUpdatesBehavior.IsNull() {
		*nonBreakingSchemaUpdatesBehavior = shared.NonBreakingSchemaUpdatesBehaviorEnum(r.NonBreakingSchemaUpdatesBehavior.ValueString())
	} else {
		nonBreakingSchemaUpdatesBehavior = nil
	}
	prefix := new(string)
	if !r.Prefix.IsUnknown() && !r.Prefix.IsNull() {
		*prefix = r.Prefix.ValueString()
	} else {
		prefix = nil
	}
	var schedule *shared.ConnectionSchedule
	if r.Schedule != nil {
		cronExpression := new(string)
		if !r.Schedule.CronExpression.IsUnknown() && !r.Schedule.CronExpression.IsNull() {
			*cronExpression = r.Schedule.CronExpression.ValueString()
		} else {
			cronExpression = nil
		}
		scheduleType := shared.ScheduleTypeEnum(r.Schedule.ScheduleType.ValueString())
		schedule = &shared.ConnectionSchedule{
			CronExpression: cronExpression,
			ScheduleType:   scheduleType,
		}
	}
	sourceID := r.SourceID.ValueString()
	status := new(shared.ConnectionStatusEnum)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = shared.ConnectionStatusEnum(r.Status.ValueString())
	} else {
		status = nil
	}
	out := shared.ConnectionCreateRequest{
		Configurations:                   configurations,
		DataResidency:                    dataResidency,
		DestinationID:                    destinationID,
		Name:                             name1,
		NamespaceDefinition:              namespaceDefinition,
		NamespaceFormat:                  namespaceFormat,
		NonBreakingSchemaUpdatesBehavior: nonBreakingSchemaUpdatesBehavior,
		Prefix:                           prefix,
		Schedule:                         schedule,
		SourceID:                         sourceID,
		Status:                           status,
	}
	return &out
}

func (r *ConnectionResourceModel) ToDeleteSDKType() *shared.ConnectionCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *ConnectionResourceModel) RefreshFromCreateResponse(resp *shared.ConnectionResponse) {
	r.Configurations.Streams = nil
	for _, streamsItem := range resp.Configurations.Streams {
		var streams1 StreamConfiguration
		streams1.CursorField = nil
		for _, v := range streamsItem.CursorField {
			streams1.CursorField = append(streams1.CursorField, types.StringValue(v))
		}
		streams1.Name = types.StringValue(streamsItem.Name)
		// Not Implemented streamsItem.PrimaryKey, {"Format":"","Discriminator":null,"BaseName":"","Fields":[],"Enum":null,"RefType":"","Input":false,"AdditionalProperties":null,"Type":"array","Scope":"","Truncated":false,"Extensions":{},"Example":null,"AssociatedTypes":[],"ItemType":{"Extensions":{"Symbol":"PrimaryKey"},"Input":false,"Type":"array","Fields":[],"AssociatedTypes":[],"RefType":"","Comments":null,"Output":false,"Example":null,"Format":"","Name":"","ItemType":{"Scope":"","Comments":null,"Extensions":{"Symbol":"CursorField"},"Example":null,"RefType":"","Truncated":false,"Name":"","AssociatedTypes":[],"Enum":null,"BaseName":"","Output":false,"AdditionalProperties":null,"Discriminator":null,"Type":"string","ItemType":null,"Fields":[],"Input":false,"Format":""},"Scope":"","BaseName":"","Enum":null,"Truncated":false,"AdditionalProperties":null,"Discriminator":null},"Comments":null,"Output":false,"Name":""}, true, , , streams1.PrimaryKey
		if streamsItem.SyncMode != nil {
			streams1.SyncMode = types.StringValue(string(*streamsItem.SyncMode))
		} else {
			streams1.SyncMode = types.StringNull()
		}
		r.Configurations.Streams = append(r.Configurations.Streams, streams1)
	}
	r.ConnectionID = types.StringValue(resp.ConnectionID)
	r.DataResidency = types.StringValue(string(resp.DataResidency))
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.Name = types.StringValue(resp.Name)
	if resp.NamespaceDefinition != nil {
		r.NamespaceDefinition = types.StringValue(string(*resp.NamespaceDefinition))
	} else {
		r.NamespaceDefinition = types.StringNull()
	}
	if resp.NamespaceFormat != nil {
		r.NamespaceFormat = types.StringValue(*resp.NamespaceFormat)
	} else {
		r.NamespaceFormat = types.StringNull()
	}
	if resp.NonBreakingSchemaUpdatesBehavior != nil {
		r.NonBreakingSchemaUpdatesBehavior = types.StringValue(string(*resp.NonBreakingSchemaUpdatesBehavior))
	} else {
		r.NonBreakingSchemaUpdatesBehavior = types.StringNull()
	}
	if resp.Prefix != nil {
		r.Prefix = types.StringValue(*resp.Prefix)
	} else {
		r.Prefix = types.StringNull()
	}
	r.Schedule.ScheduleType = types.StringValue(string(resp.Schedule.ScheduleType))
	r.SourceID = types.StringValue(resp.SourceID)
	r.Status = types.StringValue(string(resp.Status))
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

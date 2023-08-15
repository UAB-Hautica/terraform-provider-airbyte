// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceClickupAPIDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceClickupAPIDataSource{}

func NewSourceClickupAPIDataSource() datasource.DataSource {
	return &SourceClickupAPIDataSource{}
}

// SourceClickupAPIDataSource is the data source implementation.
type SourceClickupAPIDataSource struct {
	client *sdk.SDK
}

// SourceClickupAPIDataSourceModel describes the data model.
type SourceClickupAPIDataSourceModel struct {
	Configuration SourceClickupAPI `tfsdk:"configuration"`
	Name          types.String     `tfsdk:"name"`
	SecretID      types.String     `tfsdk:"secret_id"`
	SourceID      types.String     `tfsdk:"source_id"`
	WorkspaceID   types.String     `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceClickupAPIDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_clickup_api"
}

// Schema defines the schema for the data source.
func (r *SourceClickupAPIDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceClickupAPI DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"api_token": schema.StringAttribute{
						Computed:    true,
						Description: `Every ClickUp API call required authentication. This field is your personal API token. See <a href="https://clickup.com/api/developer-portal/authentication/#personal-token">here</a>.`,
					},
					"folder_id": schema.StringAttribute{
						Computed:    true,
						Description: `The ID of your folder in your space. Retrieve it from the ` + "`" + `/space/{space_id}/folder` + "`" + ` of the ClickUp API. See <a href="https://clickup.com/api/clickupreference/operation/GetFolders/">here</a>.`,
					},
					"include_closed_tasks": schema.BoolAttribute{
						Computed:    true,
						Description: `Include or exclude closed tasks. By default, they are excluded. See <a https://clickup.com/api/clickupreference/operation/GetTasks/#!in=query&path=include_closed&t=request">here</a>.`,
					},
					"list_id": schema.StringAttribute{
						Computed:    true,
						Description: `The ID of your list in your folder. Retrieve it from the ` + "`" + `/folder/{folder_id}/list` + "`" + ` of the ClickUp API. See <a href="https://clickup.com/api/clickupreference/operation/GetLists/">here</a>.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"clickup-api",
							),
						},
						Description: `must be one of ["clickup-api"]`,
					},
					"space_id": schema.StringAttribute{
						Computed:    true,
						Description: `The ID of your space in your workspace. Retrieve it from the ` + "`" + `/team/{team_id}/space` + "`" + ` of the ClickUp API. See <a href="https://clickup.com/api/clickupreference/operation/GetSpaces/">here</a>.`,
					},
					"team_id": schema.StringAttribute{
						Computed:    true,
						Description: `The ID of your team in ClickUp. Retrieve it from the ` + "`" + `/team` + "`" + ` of the ClickUp API. See <a href="https://clickup.com/api/clickupreference/operation/GetAuthorizedTeams/">here</a>.`,
					},
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret_id": schema.StringAttribute{
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SourceClickupAPIDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceClickupAPIDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceClickupAPIDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceClickupAPIRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceClickupAPI(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

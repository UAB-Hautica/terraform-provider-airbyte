// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceZohoCrmDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceZohoCrmDataSource{}

func NewSourceZohoCrmDataSource() datasource.DataSource {
	return &SourceZohoCrmDataSource{}
}

// SourceZohoCrmDataSource is the data source implementation.
type SourceZohoCrmDataSource struct {
	client *sdk.SDK
}

// SourceZohoCrmDataSourceModel describes the data model.
type SourceZohoCrmDataSourceModel struct {
	Configuration SourceZohoCrm `tfsdk:"configuration"`
	Name          types.String  `tfsdk:"name"`
	SecretID      types.String  `tfsdk:"secret_id"`
	SourceID      types.String  `tfsdk:"source_id"`
	WorkspaceID   types.String  `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceZohoCrmDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_zoho_crm"
}

// Schema defines the schema for the data source.
func (r *SourceZohoCrmDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceZohoCrm DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Computed:    true,
						Description: `OAuth2.0 Client ID`,
					},
					"client_secret": schema.StringAttribute{
						Computed:    true,
						Description: `OAuth2.0 Client Secret`,
					},
					"dc_region": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"US",
								"AU",
								"EU",
								"IN",
								"CN",
								"JP",
							),
						},
						MarkdownDescription: `must be one of ["US", "AU", "EU", "IN", "CN", "JP"]` + "\n" +
							`Please choose the region of your Data Center location. More info by this <a href="https://www.zoho.com/crm/developer/docs/api/v2/multi-dc.html">Link</a>`,
					},
					"edition": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Free",
								"Standard",
								"Professional",
								"Enterprise",
								"Ultimate",
							),
						},
						MarkdownDescription: `must be one of ["Free", "Standard", "Professional", "Enterprise", "Ultimate"]` + "\n" +
							`Choose your Edition of Zoho CRM to determine API Concurrency Limits`,
					},
					"environment": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Production",
								"Developer",
								"Sandbox",
							),
						},
						MarkdownDescription: `must be one of ["Production", "Developer", "Sandbox"]` + "\n" +
							`Please choose the environment`,
					},
					"refresh_token": schema.StringAttribute{
						Computed:    true,
						Description: `OAuth2.0 Refresh Token`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"zoho-crm",
							),
						},
						Description: `must be one of ["zoho-crm"]`,
					},
					"start_datetime": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `ISO 8601, for instance: ` + "`" + `YYYY-MM-DD` + "`" + `, ` + "`" + `YYYY-MM-DD HH:MM:SS+HH:MM` + "`" + ``,
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

func (r *SourceZohoCrmDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceZohoCrmDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceZohoCrmDataSourceModel
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
	request := operations.GetSourceZohoCrmRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceZohoCrm(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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

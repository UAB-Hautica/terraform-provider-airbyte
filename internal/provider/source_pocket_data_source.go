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
var _ datasource.DataSource = &SourcePocketDataSource{}
var _ datasource.DataSourceWithConfigure = &SourcePocketDataSource{}

func NewSourcePocketDataSource() datasource.DataSource {
	return &SourcePocketDataSource{}
}

// SourcePocketDataSource is the data source implementation.
type SourcePocketDataSource struct {
	client *sdk.SDK
}

// SourcePocketDataSourceModel describes the data model.
type SourcePocketDataSourceModel struct {
	Configuration SourcePocket `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourcePocketDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_pocket"
}

// Schema defines the schema for the data source.
func (r *SourcePocketDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourcePocket DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"access_token": schema.StringAttribute{
						Computed:    true,
						Description: `The user's Pocket access token.`,
					},
					"consumer_key": schema.StringAttribute{
						Computed:    true,
						Description: `Your application's Consumer Key.`,
					},
					"content_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"article",
								"video",
								"image",
							),
						},
						MarkdownDescription: `must be one of ["article", "video", "image"]` + "\n" +
							`Select the content type of the items to retrieve.`,
					},
					"detail_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"simple",
								"complete",
							),
						},
						MarkdownDescription: `must be one of ["simple", "complete"]` + "\n" +
							`Select the granularity of the information about each item.`,
					},
					"domain": schema.StringAttribute{
						Computed:    true,
						Description: `Only return items from a particular ` + "`" + `domain` + "`" + `.`,
					},
					"favorite": schema.BoolAttribute{
						Computed:    true,
						Description: `Retrieve only favorited items.`,
					},
					"search": schema.StringAttribute{
						Computed:    true,
						Description: `Only return items whose title or url contain the ` + "`" + `search` + "`" + ` string.`,
					},
					"since": schema.StringAttribute{
						Computed:    true,
						Description: `Only return items modified since the given timestamp.`,
					},
					"sort": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"newest",
								"oldest",
								"title",
								"site",
							),
						},
						MarkdownDescription: `must be one of ["newest", "oldest", "title", "site"]` + "\n" +
							`Sort retrieved items by the given criteria.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"pocket",
							),
						},
						Description: `must be one of ["pocket"]`,
					},
					"state": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"unread",
								"archive",
								"all",
							),
						},
						MarkdownDescription: `must be one of ["unread", "archive", "all"]` + "\n" +
							`Select the state of the items to retrieve.`,
					},
					"tag": schema.StringAttribute{
						Computed:    true,
						Description: `Return only items tagged with this tag name. Use _untagged_ for retrieving only untagged items.`,
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

func (r *SourcePocketDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourcePocketDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourcePocketDataSourceModel
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
	request := operations.GetSourcePocketRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourcePocket(ctx, request)
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

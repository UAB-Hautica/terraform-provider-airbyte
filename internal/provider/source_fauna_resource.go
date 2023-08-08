// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	speakeasy_stringplanmodifier "airbyte/internal/planmodifiers/stringplanmodifier"
	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceFaunaResource{}
var _ resource.ResourceWithImportState = &SourceFaunaResource{}

func NewSourceFaunaResource() resource.Resource {
	return &SourceFaunaResource{}
}

// SourceFaunaResource defines the resource implementation.
type SourceFaunaResource struct {
	client *sdk.SDK
}

// SourceFaunaResourceModel describes the resource data model.
type SourceFaunaResourceModel struct {
	Configuration SourceFauna  `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	SourceType    types.String `tfsdk:"source_type"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

func (r *SourceFaunaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_fauna"
}

func (r *SourceFaunaResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceFauna Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"collection": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"deletions": schema.SingleNestedAttribute{
								Required: true,
								Attributes: map[string]schema.Attribute{
									"source_fauna_collection_deletion_mode_disabled": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"deletion_mode": schema.StringAttribute{
												Required: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"ignore",
													),
												},
												Description: `must be one of ["ignore"]`,
											},
										},
										MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
											`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
											`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
											`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
									},
									"source_fauna_collection_deletion_mode_enabled": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"column": schema.StringAttribute{
												Required:    true,
												Description: `Name of the "deleted at" column.`,
											},
											"deletion_mode": schema.StringAttribute{
												Required: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"deleted_field",
													),
												},
												Description: `must be one of ["deleted_field"]`,
											},
										},
										MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
											`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
											`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
											`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
									},
									"source_fauna_update_collection_deletion_mode_disabled": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"deletion_mode": schema.StringAttribute{
												Required: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"ignore",
													),
												},
												Description: `must be one of ["ignore"]`,
											},
										},
										MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
											`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
											`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
											`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
									},
									"source_fauna_update_collection_deletion_mode_enabled": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"column": schema.StringAttribute{
												Required:    true,
												Description: `Name of the "deleted at" column.`,
											},
											"deletion_mode": schema.StringAttribute{
												Required: true,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"deleted_field",
													),
												},
												Description: `must be one of ["deleted_field"]`,
											},
										},
										MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
											`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
											`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
											`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
									},
								},
								Validators: []validator.Object{
									validators.ExactlyOneChild(),
								},
								MarkdownDescription: `<b>This only applies to incremental syncs.</b> <br>` + "\n" +
									`Enabling deletion mode informs your destination of deleted documents.<br>` + "\n" +
									`Disabled - Leave this feature disabled, and ignore deleted documents.<br>` + "\n" +
									`Enabled - Enables this feature. When a document is deleted, the connector exports a record with a "deleted at" column containing the time that the document was deleted.`,
							},
							"page_size": schema.Int64Attribute{
								Required: true,
								MarkdownDescription: `The page size used when reading documents from the database. The larger the page size, the faster the connector processes documents. However, if a page is too large, the connector may fail. <br>` + "\n" +
									`Choose your page size based on how large the documents are. <br>` + "\n" +
									`See <a href="https://docs.fauna.com/fauna/current/learn/understanding/types#page">the docs</a>.`,
							},
						},
						Description: `Settings for the Fauna Collection.`,
					},
					"domain": schema.StringAttribute{
						Required:    true,
						Description: `Domain of Fauna to query. Defaults db.fauna.com. See <a href=https://docs.fauna.com/fauna/current/learn/understanding/region_groups#how-to-use-region-groups>the docs</a>.`,
					},
					"port": schema.Int64Attribute{
						Required:    true,
						Description: `Endpoint port.`,
					},
					"scheme": schema.StringAttribute{
						Required:    true,
						Description: `URL scheme.`,
					},
					"secret": schema.StringAttribute{
						Required:    true,
						Description: `Fauna secret, used when authenticating with the database.`,
					},
					"source_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"fauna",
							),
						},
						Description: `must be one of ["fauna"]`,
					},
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required: true,
			},
			"secret_id": schema.StringAttribute{
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"source_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required: true,
			},
		},
	}
}

func (r *SourceFaunaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceFaunaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceFaunaResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	request := *data.ToCreateSDKType()
	res, err := r.client.Sources.CreateSourceFauna(ctx, request)
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
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceFaunaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceFaunaResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.GetSourceFaunaRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceFauna(ctx, request)
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

func (r *SourceFaunaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceFaunaResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceFaunaPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceFaunaRequest{
		SourceFaunaPutRequest: sourceFaunaPutRequest,
		SourceID:              sourceID,
	}
	res, err := r.client.Sources.PutSourceFauna(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	sourceId1 := data.SourceID.ValueString()
	getRequest := operations.GetSourceFaunaRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceFauna(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if getResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", getResponse))
		return
	}
	if getResponse.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", getResponse.StatusCode), debugResponse(getResponse.RawResponse))
		return
	}
	if getResponse.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceFaunaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceFaunaResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.DeleteSourceFaunaRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceFauna(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceFaunaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}

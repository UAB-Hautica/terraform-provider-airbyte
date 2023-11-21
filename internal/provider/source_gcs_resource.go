// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"

	speakeasy_stringplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/stringplanmodifier"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/operations"
	"github.com/airbytehq/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceGcsResource{}
var _ resource.ResourceWithImportState = &SourceGcsResource{}

func NewSourceGcsResource() resource.Resource {
	return &SourceGcsResource{}
}

// SourceGcsResource defines the resource implementation.
type SourceGcsResource struct {
	client *sdk.SDK
}

// SourceGcsResourceModel describes the resource data model.
type SourceGcsResourceModel struct {
	Configuration SourceGcs    `tfsdk:"configuration"`
	DefinitionID  types.String `tfsdk:"definition_id"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	SourceType    types.String `tfsdk:"source_type"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

func (r *SourceGcsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_gcs"
}

func (r *SourceGcsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceGcs Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"bucket": schema.StringAttribute{
						Required:    true,
						Description: `Name of the GCS bucket where the file(s) exist.`,
					},
					"service_account": schema.StringAttribute{
						Required:    true,
						Description: `Enter your Google Cloud <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">service account key</a> in JSON format`,
					},
					"start_date": schema.StringAttribute{
						Optional:    true,
						Description: `UTC date and time in the format 2017-01-25T00:00:00.000000Z. Any file modified before this date will not be replicated.`,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
					},
					"streams": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"days_to_sync_if_history_is_full": schema.Int64Attribute{
									Optional: true,
									MarkdownDescription: `Default: 3` + "\n" +
										`When the state history of the file store is full, syncs will only read files that were last modified in the provided day range.`,
								},
								"format": schema.SingleNestedAttribute{
									Required: true,
									Attributes: map[string]schema.Attribute{
										"csv_format": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"delimiter": schema.StringAttribute{
													Optional: true,
													MarkdownDescription: `Default: ","` + "\n" +
														`The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'.`,
												},
												"double_quote": schema.BoolAttribute{
													Optional: true,
													MarkdownDescription: `Default: true` + "\n" +
														`Whether two quotes in a quoted CSV value denote a single quote in the data.`,
												},
												"encoding": schema.StringAttribute{
													Optional: true,
													MarkdownDescription: `Default: "utf8"` + "\n" +
														`The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options.`,
												},
												"escape_char": schema.StringAttribute{
													Optional:    true,
													Description: `The character used for escaping special characters. To disallow escaping, leave this field blank.`,
												},
												"false_values": schema.ListAttribute{
													Optional:    true,
													ElementType: types.StringType,
													Description: `A set of case-sensitive strings that should be interpreted as false values.`,
												},
												"header_definition": schema.SingleNestedAttribute{
													Optional: true,
													Attributes: map[string]schema.Attribute{
														"autogenerated": schema.SingleNestedAttribute{
															Optional:    true,
															Attributes:  map[string]schema.Attribute{},
															Description: `How headers will be defined. ` + "`" + `User Provided` + "`" + ` assumes the CSV does not have a header row and uses the headers provided and ` + "`" + `Autogenerated` + "`" + ` assumes the CSV does not have a header row and the CDK will generate headers using for ` + "`" + `f{i}` + "`" + ` where ` + "`" + `i` + "`" + ` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows.`,
														},
														"from_csv": schema.SingleNestedAttribute{
															Optional:    true,
															Attributes:  map[string]schema.Attribute{},
															Description: `How headers will be defined. ` + "`" + `User Provided` + "`" + ` assumes the CSV does not have a header row and uses the headers provided and ` + "`" + `Autogenerated` + "`" + ` assumes the CSV does not have a header row and the CDK will generate headers using for ` + "`" + `f{i}` + "`" + ` where ` + "`" + `i` + "`" + ` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows.`,
														},
														"user_provided": schema.SingleNestedAttribute{
															Optional: true,
															Attributes: map[string]schema.Attribute{
																"column_names": schema.ListAttribute{
																	Required:    true,
																	ElementType: types.StringType,
																	Description: `The column names that will be used while emitting the CSV records`,
																},
															},
															Description: `How headers will be defined. ` + "`" + `User Provided` + "`" + ` assumes the CSV does not have a header row and uses the headers provided and ` + "`" + `Autogenerated` + "`" + ` assumes the CSV does not have a header row and the CDK will generate headers using for ` + "`" + `f{i}` + "`" + ` where ` + "`" + `i` + "`" + ` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows.`,
														},
													},
													Description: `How headers will be defined. ` + "`" + `User Provided` + "`" + ` assumes the CSV does not have a header row and uses the headers provided and ` + "`" + `Autogenerated` + "`" + ` assumes the CSV does not have a header row and the CDK will generate headers using for ` + "`" + `f{i}` + "`" + ` where ` + "`" + `i` + "`" + ` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows.`,
													Validators: []validator.Object{
														validators.ExactlyOneChild(),
													},
												},
												"inference_type": schema.StringAttribute{
													Optional: true,
													MarkdownDescription: `must be one of ["None", "Primitive Types Only"]; Default: "None"` + "\n" +
														`How to infer the types of the columns. If none, inference default to strings.`,
													Validators: []validator.String{
														stringvalidator.OneOf(
															"None",
															"Primitive Types Only",
														),
													},
												},
												"null_values": schema.ListAttribute{
													Optional:    true,
													ElementType: types.StringType,
													Description: `A set of case-sensitive strings that should be interpreted as null values. For example, if the value 'NA' should be interpreted as null, enter 'NA' in this field.`,
												},
												"quote_char": schema.StringAttribute{
													Optional: true,
													MarkdownDescription: `Default: "\""` + "\n" +
														`The character used for quoting CSV values. To disallow quoting, make this field blank.`,
												},
												"skip_rows_after_header": schema.Int64Attribute{
													Optional: true,
													MarkdownDescription: `Default: 0` + "\n" +
														`The number of rows to skip after the header row.`,
												},
												"skip_rows_before_header": schema.Int64Attribute{
													Optional: true,
													MarkdownDescription: `Default: 0` + "\n" +
														`The number of rows to skip before the header row. For example, if the header row is on the 3rd row, enter 2 in this field.`,
												},
												"strings_can_be_null": schema.BoolAttribute{
													Optional: true,
													MarkdownDescription: `Default: true` + "\n" +
														`Whether strings can be interpreted as null values. If true, strings that match the null_values set will be interpreted as null. If false, strings that match the null_values set will be interpreted as the string itself.`,
												},
												"true_values": schema.ListAttribute{
													Optional:    true,
													ElementType: types.StringType,
													Description: `A set of case-sensitive strings that should be interpreted as true values.`,
												},
											},
											Description: `The configuration options that are used to alter how to read incoming files that deviate from the standard formatting.`,
										},
									},
									Description: `The configuration options that are used to alter how to read incoming files that deviate from the standard formatting.`,
									Validators: []validator.Object{
										validators.ExactlyOneChild(),
									},
								},
								"globs": schema.ListAttribute{
									Optional:    true,
									ElementType: types.StringType,
									Description: `The pattern used to specify which files should be selected from the file system. For more information on glob pattern matching look <a href="https://en.wikipedia.org/wiki/Glob_(programming)">here</a>.`,
								},
								"input_schema": schema.StringAttribute{
									Optional:    true,
									Description: `The schema that will be used to validate records extracted from the file. This will override the stream schema that is auto-detected from incoming files.`,
								},
								"legacy_prefix": schema.StringAttribute{
									Optional:    true,
									Description: `The path prefix configured in previous versions of the GCS connector. This option is deprecated in favor of a single glob.`,
								},
								"name": schema.StringAttribute{
									Required:    true,
									Description: `The name of the stream.`,
								},
								"primary_key": schema.StringAttribute{
									Optional:    true,
									Sensitive:   true,
									Description: `The column or columns (for a composite key) that serves as the unique identifier of a record.`,
								},
								"schemaless": schema.BoolAttribute{
									Optional: true,
									MarkdownDescription: `Default: false` + "\n" +
										`When enabled, syncs will not validate or structure records against the stream's schema.`,
								},
								"validation_policy": schema.StringAttribute{
									Optional: true,
									MarkdownDescription: `must be one of ["Emit Record", "Skip Record", "Wait for Discover"]; Default: "Emit Record"` + "\n" +
										`The name of the validation policy that dictates sync behavior when a record does not adhere to the stream schema.`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"Emit Record",
											"Skip Record",
											"Wait for Discover",
										),
									},
								},
							},
						},
						Description: `Each instance of this configuration defines a <a href=https://docs.airbyte.com/cloud/core-concepts#stream>stream</a>. Use this to define which files belong in the stream, their format, and how they should be parsed and validated. When sending data to warehouse destination such as Snowflake or BigQuery, each stream is a separate table.`,
					},
				},
				MarkdownDescription: `NOTE: When this Spec is changed, legacy_config_transformer.py must also be` + "\n" +
					`modified to uptake the changes because it is responsible for converting` + "\n" +
					`legacy GCS configs into file based configs using the File-Based CDK.`,
			},
			"definition_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.`,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required:    true,
				Description: `Name of the source e.g. dev-mysql-instance.`,
			},
			"secret_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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

func (r *SourceGcsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceGcsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceGcsResourceModel
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

	request := data.ToCreateSDKType()
	res, err := r.client.Sources.CreateSourceGcs(ctx, request)
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
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGcsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceGcsResourceModel
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
	request := operations.GetSourceGcsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceGcs(ctx, request)
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

func (r *SourceGcsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceGcsResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceGcsPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceGcsRequest{
		SourceGcsPutRequest: sourceGcsPutRequest,
		SourceID:            sourceID,
	}
	res, err := r.client.Sources.PutSourceGcs(ctx, request)
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
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	sourceId1 := data.SourceID.ValueString()
	getRequest := operations.GetSourceGcsRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceGcs(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
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
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceGcsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceGcsResourceModel
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
	request := operations.DeleteSourceGcsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceGcs(ctx, request)
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
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceGcsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_id"), req.ID)...)
}

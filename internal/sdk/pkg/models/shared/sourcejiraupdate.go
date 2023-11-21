// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type IssuesStreamExpandWith string

const (
	IssuesStreamExpandWithRenderedFields IssuesStreamExpandWith = "renderedFields"
	IssuesStreamExpandWithTransitions    IssuesStreamExpandWith = "transitions"
	IssuesStreamExpandWithChangelog      IssuesStreamExpandWith = "changelog"
)

func (e IssuesStreamExpandWith) ToPointer() *IssuesStreamExpandWith {
	return &e
}

func (e *IssuesStreamExpandWith) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "renderedFields":
		fallthrough
	case "transitions":
		fallthrough
	case "changelog":
		*e = IssuesStreamExpandWith(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IssuesStreamExpandWith: %v", v)
	}
}

type SourceJiraUpdate struct {
	// Jira API Token. See the <a href="https://docs.airbyte.com/integrations/sources/jira">docs</a> for more information on how to generate this key. API Token is used for Authorization to your account by BasicAuth.
	APIToken string `json:"api_token"`
	// The Domain for your Jira account, e.g. airbyteio.atlassian.net, airbyteio.jira.com, jira.your-domain.com
	Domain string `json:"domain"`
	// The user email for your Jira account which you used to generate the API token. This field is used for Authorization to your account by BasicAuth.
	Email string `json:"email"`
	// Allow the use of experimental streams which rely on undocumented Jira API endpoints. See https://docs.airbyte.com/integrations/sources/jira#experimental-tables for more info.
	EnableExperimentalStreams *bool `default:"false" json:"enable_experimental_streams"`
	// (DEPRECATED) Expand the changelog when replicating issues.
	ExpandIssueChangelog *bool `default:"false" json:"expand_issue_changelog"`
	// (DEPRECATED) Expand the transitions when replicating issues.
	ExpandIssueTransition *bool `default:"false" json:"expand_issue_transition"`
	// Select fields to Expand the `Issues` stream when replicating with:
	IssuesStreamExpandWith []IssuesStreamExpandWith `json:"issues_stream_expand_with,omitempty"`
	// List of Jira project keys to replicate data for, or leave it empty if you want to replicate data for all projects.
	Projects []string `json:"projects,omitempty"`
	// (DEPRECATED) Render issue fields in HTML format in addition to Jira JSON-like format.
	RenderFields *bool `default:"false" json:"render_fields"`
	// The date from which you want to replicate data from Jira, use the format YYYY-MM-DDT00:00:00Z. Note that this field only applies to certain streams, and only data generated on or after the start date will be replicated. Or leave it empty if you want to replicate all data. For more information, refer to the <a href="https://docs.airbyte.com/integrations/sources/jira/">documentation</a>.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceJiraUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceJiraUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceJiraUpdate) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceJiraUpdate) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *SourceJiraUpdate) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SourceJiraUpdate) GetEnableExperimentalStreams() *bool {
	if o == nil {
		return nil
	}
	return o.EnableExperimentalStreams
}

func (o *SourceJiraUpdate) GetExpandIssueChangelog() *bool {
	if o == nil {
		return nil
	}
	return o.ExpandIssueChangelog
}

func (o *SourceJiraUpdate) GetExpandIssueTransition() *bool {
	if o == nil {
		return nil
	}
	return o.ExpandIssueTransition
}

func (o *SourceJiraUpdate) GetIssuesStreamExpandWith() []IssuesStreamExpandWith {
	if o == nil {
		return nil
	}
	return o.IssuesStreamExpandWith
}

func (o *SourceJiraUpdate) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *SourceJiraUpdate) GetRenderFields() *bool {
	if o == nil {
		return nil
	}
	return o.RenderFields
}

func (o *SourceJiraUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceRailzPutRequest struct {
	Configuration SourceRailzUpdate `json:"configuration"`
	Name          string            `json:"name"`
	WorkspaceID   string            `json:"workspaceId"`
}

func (o *SourceRailzPutRequest) GetConfiguration() SourceRailzUpdate {
	if o == nil {
		return SourceRailzUpdate{}
	}
	return o.Configuration
}

func (o *SourceRailzPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceRailzPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

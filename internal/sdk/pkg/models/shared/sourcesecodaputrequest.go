// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSecodaPutRequest struct {
	Configuration SourceSecodaUpdate `json:"configuration"`
	Name          string             `json:"name"`
	WorkspaceID   string             `json:"workspaceId"`
}

func (o *SourceSecodaPutRequest) GetConfiguration() SourceSecodaUpdate {
	if o == nil {
		return SourceSecodaUpdate{}
	}
	return o.Configuration
}

func (o *SourceSecodaPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceSecodaPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

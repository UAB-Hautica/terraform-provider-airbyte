// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceLokalisePutRequest struct {
	Configuration SourceLokaliseUpdate `json:"configuration"`
	Name          string               `json:"name"`
	WorkspaceID   string               `json:"workspaceId"`
}

func (o *SourceLokalisePutRequest) GetConfiguration() SourceLokaliseUpdate {
	if o == nil {
		return SourceLokaliseUpdate{}
	}
	return o.Configuration
}

func (o *SourceLokalisePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceLokalisePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

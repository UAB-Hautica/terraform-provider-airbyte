// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceAirtablePutRequest struct {
	Configuration SourceAirtableUpdate `json:"configuration"`
	Name          string               `json:"name"`
	WorkspaceID   string               `json:"workspaceId"`
}

func (o *SourceAirtablePutRequest) GetConfiguration() SourceAirtableUpdate {
	if o == nil {
		return SourceAirtableUpdate{}
	}
	return o.Configuration
}

func (o *SourceAirtablePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceAirtablePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

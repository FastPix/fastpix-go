package components

type SigningKeyWorkspaceDTO struct {
	// FastPix generates a unique identifier for each workspace.
	ID *string `json:"id,omitempty"`
	// Designated title for the workspace.
	Name *string `json:"name,omitempty"`
	// Describes the type of a workspace.  Possible value: QA, staging, production, or development.
	WorkspaceType *string `json:"workspaceType,omitempty"`
}

func (s *SigningKeyWorkspaceDTO) GetID() *string {
	if s == nil {
		return nil
	}
	return s.ID
}

func (s *SigningKeyWorkspaceDTO) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SigningKeyWorkspaceDTO) GetWorkspaceType() *string {
	if s == nil {
		return nil
	}
	return s.WorkspaceType
}

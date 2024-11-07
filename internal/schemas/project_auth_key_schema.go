package schemas

import "github.com/google/uuid"

type ProjectAuthKeyInput struct {
	ProjectId string `validate:"required|uuid4" json:"project_id"`
}

func (p ProjectAuthKeyInput) GetProjectID() uuid.UUID {
	return uuid.MustParse(p.ProjectId)
}

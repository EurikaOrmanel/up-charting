package schemas

type ProjectCreateInput struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required|string" json:"description"`
}

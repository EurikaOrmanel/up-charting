package schemas

type WebhookCreateInput struct {
	ProjectId   string `validate:"required|uuid" json:"project_id"`
	Url         string `validate:"required|url" json:"url"`
	PayHeaderID string `validate:"required|string" json:"pay_header_id"`
}

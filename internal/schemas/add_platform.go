package schemas

type AddPlatformInput struct {
	Name string `validate:"string|required" json:"name"`
	Url  string `validate:"required|fullUrl" json:"url"`
}

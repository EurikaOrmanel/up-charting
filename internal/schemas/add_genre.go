package schemas

type AddGenreInput struct {
	Name string `validate:"string|required" json:"name"`
}

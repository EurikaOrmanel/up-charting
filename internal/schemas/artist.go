package schemas

type AddArtistInput struct {
	Stagename string                `validate:"required|string" json:"stagename"`
	Firstname string                `validate:"required|string" json:"firstname"`
	Lastname  string                `validate:"required|string" json:"lastname"`
	Platforms ArtistPlatformInputs `json:"platforms"`
}

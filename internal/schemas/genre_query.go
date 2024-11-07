package schemas

type GenreQuery struct {
	PaginationQuery
	Q string `query:"q" validate:"string"`
}

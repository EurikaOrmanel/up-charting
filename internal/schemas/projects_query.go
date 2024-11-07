package schemas


type ProjectsQuery struct{
	PaginationQuery
	Q string `query:"q" validate:"string"`
}
package schemas

type PlatformQuery struct {
	PaginationQuery
	Q string `query:"q" validate:"string"`
}

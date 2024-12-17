package schemas

type SongChartQuery struct {
	PaginationQuery
	Genre string `query:"genre"`
}

package schemas

type PaginationQuery struct {
	Count int `validate:"integer" query:"count"`
	Page  int `validate:"integer" query:"page"`
}

func (nQ *PaginationQuery) GetOffset() int {
	return (nQ.GetPage() - 1) * nQ.GetCount()
}

func (nQ *PaginationQuery) GetCount() int {
	if nQ.Count == 0 {
		nQ.Count = 10
	}
	return nQ.Count
}

func (nQ *PaginationQuery) GetPage() int {
	if nQ.Page == 0 {
		nQ.Page = 1
	}
	return nQ.Page
}

func (nQ *PaginationQuery) Prepare() {
	if nQ.Count == 0 {
		nQ.Count = 10
	}
	if nQ.Page == 0 {

		nQ.Page = 1
	}
}

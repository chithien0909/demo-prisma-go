package utils

func BuildPagination(page, limit int) (skip, take int) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}
	skip = (page - 1) * limit
	take = limit
	return
}

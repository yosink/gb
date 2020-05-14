package utils

func GetPageOffset(page, perPage int) int {
	if page < 1 {
		return 0
	}
	return (page - 1) * perPage
}

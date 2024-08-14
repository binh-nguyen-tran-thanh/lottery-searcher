package util

const (
	DEFAULT_LIMIT  = 10
	MAX_LIMIT      = 20
	DEFAULT_OFFSET = 0
)

func GeneratePagingParamsWithDefaultValue(limit uint, offset uint) (validLimit uint, validOffset uint) {
	validLimit = limit
	validOffset = offset
	if limit <= 0 {
		validLimit = DEFAULT_LIMIT
	}

	if limit > MAX_LIMIT {
		validLimit = MAX_LIMIT
	}

	return
}

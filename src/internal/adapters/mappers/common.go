package mappers

import "github.com/RenZorRUS/todo-backend/src/internal/utils"

func GetOrDefaultRowsConstraint(value *uint64, defaultValue int64) (int64, error) {
	if value == nil {
		return defaultValue, nil
	}

	return utils.UintToInt[uint64, int64](*value)
}

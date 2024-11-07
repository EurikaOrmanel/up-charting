package repositories

import (
	"EurikaOrmanel/up-charter/internal/schemas"

	"gorm.io/gorm"
)

func paginate(pagination schemas.PaginationQuery, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetCount())
	}

}

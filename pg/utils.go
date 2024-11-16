package pg

import (
	"errors"

	"git.sriss.uz/shared/shared_service/sharedutil"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Paginate = sharedutil.Paginate
	Filter   = func(tx *gorm.DB) *gorm.DB
)

type pageEntity[T any] struct {
	Total int64
	Data  T `gorm:"embedded"`
}

type Model interface {
	TableName() string
}

var ErrRowsAffected = errors.New("not rows affected")

func NewReturning(columns ...string) clause.Returning {
	var clauseReturning clause.Returning
	{
		for _, column := range columns {
			clauseReturning.Columns = append(
				clauseReturning.Columns, clause.Column{
					Name: column,
				},
			)
		}
	}
	return clauseReturning
}

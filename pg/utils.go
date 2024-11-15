package repository

import (
	"errors"

	"gitea.avtomig.uz/mydream/shared_service/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Paginate = common.Paginate
	Filter   = func(tx *gorm.DB) *gorm.DB
	// PageResult[T any] response.PageData[T]
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

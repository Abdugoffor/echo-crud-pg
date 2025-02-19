package pg

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Filter = func(tx *gorm.DB) *gorm.DB
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

func IsTx(db *gorm.DB) bool {
	return db.Statement != nil && db.Statement.ConnPool != db.ConnPool
}

func Transaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {

	tx := db
	{
		if !IsTx(tx) {
			tx = tx.Begin()
		}
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

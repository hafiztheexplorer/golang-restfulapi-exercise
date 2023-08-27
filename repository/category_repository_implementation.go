package repository

import (
	"context"
	"database/sql"
	"golang-restfulapi-exercise/helper"
	"golang-restfulapi-exercise/model/domain"
)

type CategoryRepositoryImplem struct{}

// funtion untuk post
func (repository *CategoryRepositoryImplem) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sqlscript := "insert into category (namakategori) values (?)"
	result, error := tx.ExecContext(ctx, sqlscript, category.Namakategori)

	/*
		agar tidak perlu melakaukan ngetik handling error berkali2
		maka buat saja function PanicIfError di folder helper/error.go
		isinya:
		if err != nil {
			panic(err)
		}
	*/
	helper.PanicIfError(error)

	id, error := result.LastInsertId()
	helper.PanicIfError(error)

	category.Id = id
	return category

}

func (repository *CategoryRepositoryImplem) Update(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category {
	panic("")
}

func (repository *CategoryRepositoryImplem) Delete(ctx context.Context, tx *sql.Tx, Category domain.Category) {
	panic("")
}

func (repository *CategoryRepositoryImplem) FindById(ctx context.Context, tx *sql.Tx, idKategori int64) domain.Category {
	panic("")
}

func (repository *CategoryRepositoryImplem) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	panic("")
}

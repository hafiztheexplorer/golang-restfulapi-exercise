package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-restfulapi-exercise/helper"
	"golang-restfulapi-exercise/model/domain"
)

type CategoryRepositoryImplem struct{}

// funtion interaksi dari go ke database untuk post nama kategori
func (repository *CategoryRepositoryImplem) Create(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category {
	sqlscript := "insert into category (namakategori) values (?)"
	result, error := tx.ExecContext(ctx, sqlscript, Category.Namakategori)

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

	Category.Id = id
	return Category

}

// function interaksi dari go ke database untuk put/update/replace data di database, by menentukan namakategorinya apa dan di id berapa
func (repository *CategoryRepositoryImplem) Update(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category {
	sqlscript := "update category set namakategori = ? where id = ?"
	result, error := tx.ExecContext(ctx, sqlscript, Category.Namakategori, Category.Id)
	fmt.Print(result)
	helper.PanicIfError(error)

	return Category
}

// function interaksi dari go ke database untuk delete data di database by id
func (repository *CategoryRepositoryImplem) Delete(ctx context.Context, tx *sql.Tx, Category domain.Category) {
	sqlscript := "delete from where id = ?"
	result, error := tx.ExecContext(ctx, sqlscript, Category.Id)
	fmt.Print(result)
	helper.PanicIfError(error)

}

// function interaksi dari go ke database untuk menemukan(select) row data(1 row isinya id dan namakategori), dari databse by Id
func (repository *CategoryRepositoryImplem) FindById(ctx context.Context, tx *sql.Tx, idKategori int64) (domain.Category, error) {
	sqlscript := "select id, namakategori from category where id = ?"
	rows, error := tx.QueryContext(ctx, sqlscript, idKategori)
	helper.PanicIfError(error)

	category := domain.Category{}

	if rows.Next() {
		error := rows.Scan(&category.Id, &category.Namakategori)
		helper.PanicIfError(error)
		return category, nil

	} else {
		return category, errors.New("kategori tidak ditemukan")
	}

}

// function interaksi dari go ke database untuk mendisplaykan (select) row semua data(1 row berupa id, namakategori), diuang dengan menggunakan perulangan for dan rows.Next
func (repository *CategoryRepositoryImplem) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sqlscript := "select id,namakategori from category"
	rows, error := tx.QueryContext(ctx, sqlscript)
	helper.PanicIfError(error)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		error := rows.Scan(&category.Id, &category.Namakategori)
		helper.PanicIfError(error)
		categories = append(categories, domain.Category{})
	}

	return categories
}

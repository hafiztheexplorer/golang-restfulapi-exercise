package repository

import (
	"context"
	"database/sql"
	"golang-restfulapi-exercise/model/domain"
)

// kita buat interfacenya dulu jangan langsung struct
type CategoryRepository interface {
	// interface untuk post
	Create(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, Category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, idKategori int64) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}

package repository

import (
	"context"
	"database/sql"
	"hafiztheexplorer/golang-restfulapi-exercise/model/domain"
)

// kita buat interfacenya dulu jangan langsung struct
type CategoryRepository interface {
	Create(Category context.Context, sql.Tx, Category domain.Category)
	Update(ctx context.Context, sql.Tx)
	Delete(ctx context.Context, sql.Tx)
	FindById(ctx context.Context, sql.Tx)
	FindAll(ctx context.Context, sql.Tx)
}

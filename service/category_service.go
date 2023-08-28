package service

import (
	"context"
	"golang-restfulapi-exercise/model/web"
)

type CategoryService interface {
	//Findbyid category spesifik ke id korelasi dengan API GetById
	FindByIdGet(ctx context.Context, idKategori int64) web.CategoryResponse
	//FindAll category tidak spesifik, korelasi dengan API Get
	FindAllGet(ctx context.Context) []web.CategoryResponse
	//Create category korelasi dengan API Post
	CreatePost(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	//Update/replace category korelasi dengan API Put
	UpdatePut(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	//Delete category korelasi dengan API Get
	DeleteDelete(ctx context.Context, idKategori int64)
}

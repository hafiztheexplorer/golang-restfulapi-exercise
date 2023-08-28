package web

import (
	"context"
	"golang-restfulapi-exercise/model/web"
)

type CategoryCreateRequest interface {
	//Findbyid category spesifik ke id korelasi dengan API GetById
	FindByIdGet(ctx context.Context)
	//FindAll category tidak spesifik, korelasi dengan API Get
	FindAllGet(ctx context.Context)
	//Create category korelasi dengan API Post
	CreatePost(ctx context.Context, request web.CategoryCreateRequest)
	//Update/replace category korelasi dengan API Put
	UpdatePut(ctx context.Context)
	//Delete category korelasi dengan API Get
	DeleteDelete(ctx context.Context)
}

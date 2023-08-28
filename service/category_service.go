package service

import "context"

type CategoryService interface {
	//Findbyid category spesifik ke id korelasi dengan API GetById
	FindByIdGet(ctx context.Context)
	//FindAll category tidak spesifik, korelasi dengan API Get
	FindAllGet(ctx context.Context)
	//Create category korelasi dengan API Post
	CreatePost(ctx context.Context)
	//Update/replace category korelasi dengan API Put
	UpdatePut(ctx context.Context)
	//Delete category korelasi dengan API Get
	DeleteDelete(ctx context.Context)
}

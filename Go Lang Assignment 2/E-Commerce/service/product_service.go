package service

import (
	"ecommerce/model"
	"ecommerce/repository"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: productRepo}
}

func (service *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
	return service.ProductRepo.CreateProduct(product)
}

func (service *ProductService) GetProduct(id int) (*model.Product, error) {
	return service.ProductRepo.GetProduct(id)
}

func (service *ProductService) GetAllProducts() ([]model.Product, error) {
	return service.ProductRepo.GetAllProducts()
}

func (service *ProductService) UpdateProduct(product *model.Product) (*model.Product, error) {
	return service.ProductRepo.UpdateProduct(product)
}

func (service *ProductService) DeleteProduct(id int) error {
	return service.ProductRepo.DeleteProduct(id)
}

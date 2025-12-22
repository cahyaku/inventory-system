package controller

import (
	"fmt"
	"inventory-system/internal/service"
	"inventory-system/internal/utils"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service}
}

func (c *CategoryController) ShowCategories() {
	categories, _ := c.service.GetAll()

	if len(categories) == 0 {
		println("No categories available.")
		return
	}

	for _, category := range categories {
		println(category.Name)
	}
}

func (c *CategoryController) CreateCategory() {
	name := utils.ReadLine("Category name: ")

	err := c.service.Create(name)
	if err != nil {
		println("Error:", err.Error())
		return
	}

	fmt.Println("Category created successfully")
}

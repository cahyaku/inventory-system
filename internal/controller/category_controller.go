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

func (c *CategoryController) UpdateCategory() {
	categories, err := c.service.GetAll()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	if len(categories) == 0 {
		fmt.Println("No categories to edit.")
		return
	}

	// tampilkan list
	for i, cat := range categories {
		fmt.Printf("%d. %s\n", i+1, cat.Name)
	}

	index, err := utils.ReadInt("Select category number: ")
	if err != nil || index < 1 || index > len(categories) {
		fmt.Println("Invalid category number.")
		return
	}

	newName := utils.ReadLine("New category name: ")

	err = c.service.Update(categories[index-1].ID, newName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Category updated successfully ✅")
}

func (c *CategoryController) DeleteCategory() {
	categories, err := c.service.GetAll()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	if len(categories) == 0 {
		fmt.Println("No categories to delete.")
		return
	}

	for i, cat := range categories {
		fmt.Printf("%d. %s\n", i+1, cat.Name)
	}

	index, err := utils.ReadInt("Select category number: ")
	if err != nil || index < 1 || index > len(categories) {
		fmt.Println("Invalid category number.")
		return
	}

	err = c.service.Delete(categories[index-1].ID)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Category deleted successfully 🗑️")
}

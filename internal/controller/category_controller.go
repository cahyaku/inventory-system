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
	categories, err := c.service.GetAll()
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	if len(categories) == 0 {
		fmt.Println("No categories available!")
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("--------Available Categories--------")
	for i, category := range categories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}

	utils.PressEnterToContinue()
}

func (c *CategoryController) CreateCategory() {
	fmt.Println("--------Create Category--------")
	name := utils.ReadLine("Category name: ")

	err := c.service.Create(name)
	if err != nil {
		println("Error:", err.Error())
		return
	}

	fmt.Println("Category (", name, ") created successfully")
	utils.PressEnterToContinue()
}

func (c *CategoryController) UpdateCategory() {
	categories, err := c.service.GetAll()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	if len(categories) == 0 {
		fmt.Println("No categories to edit!")
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("--------Update Categories--------")
	// tampilkan list
	for i, cat := range categories {
		fmt.Printf("%d. %s\n", i+1, cat.Name)
	}

	index, err := utils.ReadInt("Select category number: ")
	if err != nil || index < 1 || index > len(categories) {
		fmt.Println("Invalid category number.")
		utils.PressEnterToContinue()
		return
	}

	newName := utils.ReadLine("New category name: ")

	err = c.service.Update(categories[index-1].ID, newName)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Category (", categories[index-1].Name, ") updated successfully ✅")
	utils.PressEnterToContinue()
}

func (c *CategoryController) DeleteCategory() {
	categories, err := c.service.GetAll()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	if len(categories) == 0 {
		fmt.Println("No categories to delete!")
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("--------Delete Categories--------")
	for i, cat := range categories {
		fmt.Printf("%d. %s\n", i+1, cat.Name)
	}

	index, err := utils.ReadInt("Select category number: ")
	if err != nil || index < 1 || index > len(categories) {
		fmt.Println("Invalid category number.")
		utils.PressEnterToContinue()
		return
	}

	err = c.service.Delete(categories[index-1].ID)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Category deleted (", categories[index-1].Name, ") successfully!")
	utils.PressEnterToContinue()
}

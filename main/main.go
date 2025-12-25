package main

import (
	"fmt"
	"inventory-system/internal/controller"
	"inventory-system/internal/repository"
	"inventory-system/internal/service"
	"inventory-system/internal/utils"
)

func main() {
	// Repository
	itemRepository := repository.NewInMemoryItemRepository()
	categoryRepository := repository.NewInMemoryCategoryRepository()

	// Service
	itemService := service.NewItemService(itemRepository, categoryRepository)
	categoryService := service.NewCategoryService(categoryRepository, itemRepository)

	// Controller
	itemController := controller.NewItemController(itemService, categoryService)
	categoryController := controller.NewCategoryController(categoryService)

	for {
		showMainMenu()
		input, err := utils.ReadInt("Enter your choice: ")
		if err != nil {
			fmt.Println("Please input number between 1 and 6 👺")
			utils.PressEnterToContinue()
			continue
		}

		switch input {
		case 1:
			itemController.ShowItems()
			fmt.Println()
		case 2:
			itemController.CreateItem()
			fmt.Println()
		case 3:
			itemController.UpdateItem()
			fmt.Println()
		case 4:
			itemController.DeleteItem()
			fmt.Println()
		case 5:
			categoryMenu(categoryController)
		case 6:
			fmt.Println("Thank you for using this app bye bye....👋😊")
			return
		default:
			fmt.Println("Please input number between 1 and 6 👺👺👺👺👺")
			utils.PressEnterToContinue()
			continue
		}
	}
}

/**
 * Function to display the main menu
 */
func showMainMenu() {
	fmt.Println("══════════════════════════════════════════════════════════")
	fmt.Println("════════════════════ INVENTORY SYSTEM ════════════════════")
	fmt.Println("══════════════════════════════════════════════════════════")

	fmt.Println("Menu:")
	fmt.Println("1. Show items")
	fmt.Println("2. Add items")
	fmt.Println("3. Update item stock")
	fmt.Println("4. Remove item")
	fmt.Println("5. Manage categories")
	fmt.Println("6. Exit")
}

func showCategoryMenu() {
	fmt.Println("══════════════ CATEGORY MANAGEMENT ══════════════")
	fmt.Println("1. Show categories")
	fmt.Println("2. Add category")
	fmt.Println("3. Edit category")
	fmt.Println("4. Remove category")
	fmt.Println("5. Back to main menu")
}

func categoryMenu(categoryController *controller.CategoryController) {
	for {
		fmt.Println()
		showCategoryMenu()

		input, err := utils.ReadInt("Enter your choice: ")
		if err != nil {
			fmt.Println("Invalid choice (please input number between 1 and 5 👺👺👺👺👺)")
			utils.PressEnterToContinue()
			continue
		}

		switch input {
		case 1:
			categoryController.ShowCategories()
		case 2:
			categoryController.CreateCategory()
		case 3:
			categoryController.UpdateCategory()
		case 4:
			categoryController.DeleteCategory()
		case 5:
			return
		default:
			fmt.Println("Invalid choice (please input number between 1 and 5 👺👺👺👺👺)")
			utils.PressEnterToContinue()
		}
	}
}

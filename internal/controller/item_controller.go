package controller

import (
	"fmt"
	"inventory-system/internal/service"
	"inventory-system/internal/utils"
	"strconv"
	"strings"
)

type ItemController struct {
	itemService     *service.ItemService
	categoryService *service.CategoryService
}

func NewItemController(itemService *service.ItemService, categoryService *service.CategoryService) *ItemController {
	return &ItemController{
		itemService:     itemService,
		categoryService: categoryService,
	}
}

func (c *ItemController) ShowItems() {
	items, _ := c.itemService.GetAllItems()

	if len(items) == 0 {
		fmt.Println("No items available.")
		utils.PressEnterToContinue()
		return
	}

	for _, item := range items {
		fmt.Printf("%s - Stock: %d - Categories: ", item.Name, item.Stock)
		for i, cat := range item.Categories {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(cat.Name)
		}
		fmt.Println()
	}
}

func (c *ItemController) CreateItem() {
	// 1. Cek apakah kategori ada
	categories, err := c.categoryService.GetAll()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	if len(categories) == 0 {
		fmt.Println("Cannot create item.")
		fmt.Println("No categories available (please create at least one category first).")
		utils.PressEnterToContinue()
		return
	}

	// 2. Tampilkan kategori agar user tahu IDnya
	fmt.Println("Available categories:")
	for _, cat := range categories {
		fmt.Printf("%d. %s\n", cat.ID, cat.Name)
	}

	// 3. Input item
	name := utils.ReadLine("Item name: ")

	// 4. Input stok awal
	stock, err := utils.ReadInt("Initial stock: ")
	if err != nil || stock == 0 {
		fmt.Println("Stock must be zero or greater.")
		utils.PressEnterToContinue()
		return
	}

	// 5. Input kategori
	input := utils.ReadLine("Select category IDs (comma separated, e.g 1,2):")

	parts := strings.Split(input, ",")
	var ids []int

	for _, part := range parts {
		id, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			fmt.Println("Invalid category ID:", part)
			utils.PressEnterToContinue()
			return
		}
		ids = append(ids, id)
	}

	if len(ids) == 0 {
		fmt.Println("You must select at least one category")
		utils.PressEnterToContinue()
		return
	}

	// 6. Panggil service untuk CreateItem
	err = c.itemService.CreateItem(name, stock, ids)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Item created successfully ✅")
	utils.PressEnterToContinue()
}

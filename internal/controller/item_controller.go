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

	fmt.Println("====== Show Items =====")
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

/* =========================
   CREATE ITEM
========================= */

func (c *ItemController) CreateItem() {
	categories, err := c.categoryService.GetAll()
	if err != nil || len(categories) == 0 {
		fmt.Println("Cannot create item. No categories available.")
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("====== Create Item =====")

	name := utils.ReadLine("Item name: ")

	stock, err := utils.ReadInt("Initial stock: ")
	if err != nil || stock < 0 {
		fmt.Println("Stock must be zero or greater.")
		utils.PressEnterToContinue()
		return
	}

	ids, err := c.readCategoryIDs()
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	err = c.itemService.Create(name, stock, ids)
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("Item (", name, ") created successfully ✅")
	utils.PressEnterToContinue()
}

/* =========================
   UPDATE ITEM
========================= */

func (c *ItemController) UpdateItem() {
	items, _ := c.itemService.GetAllItems()
	if len(items) == 0 {
		fmt.Println("No items available.")
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("====== Update Item =====")
	for i, item := range items {
		fmt.Printf("%d. %s (Stock: %d)\n", i+1, item.Name, item.Stock)
	}

	index, err := utils.ReadInt("Select item number: ")
	if err != nil || index < 1 || index > len(items) {
		fmt.Println("Invalid item number.")
		utils.PressEnterToContinue()
		return
	}

	item := items[index-1]

	newName := utils.ReadLine("New item name: ")
	newStock, err := utils.ReadInt("New stock: ")
	if err != nil || newStock < 0 {
		fmt.Println("Stock must be zero or greater.")
		utils.PressEnterToContinue()
		return
	}

	ids, err := c.readCategoryIDs()
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	err = c.itemService.Update(item.ID, newName, newStock, ids)
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("Item updated successfully ✅")
	utils.PressEnterToContinue()
}

func (c *ItemController) DeleteItem() {
	items, err := c.itemService.GetAllItems()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	if len(items) == 0 {
		fmt.Println("No items to delete.")
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("====== Delete Item =====")
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item.Name)
	}

	index, err := utils.ReadInt("Select item number: ")
	if err != nil || index < 1 || index > len(items) {
		fmt.Println("Invalid item number.")
		utils.PressEnterToContinue()
		return
	}

	err = c.itemService.Delete(items[index-1].ID)
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("Item deleted successfully!")
	utils.PressEnterToContinue()
}

/* =========================
   HELPER: READ CATEGORY IDS
========================= */

func (c *ItemController) readCategoryIDs() ([]int, error) {
	categories, err := c.categoryService.GetAll()
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, fmt.Errorf("no categories available")
	}

	fmt.Println("Available categories:")
	for i, cat := range categories {
		fmt.Printf("%d. %s\n", i+1, cat.Name)
	}

	input := utils.ReadLine("Select category numbers (comma separated, e.g 1,2,3): ")
	parts := strings.Split(input, ",")

	if len(parts) == 0 {
		return nil, fmt.Errorf("you must select at least one category")
	}

	var ids []int
	selected := make(map[int]bool)

	for _, part := range parts {
		index, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid category number: %s", part)
		}

		// VALIDASI INDEX (1-based → 0-based)
		if index < 1 || index > len(categories) {
			return nil, fmt.Errorf("category number %d is out of range", index)
		}

		realID := categories[index-1].ID

		// Cegah duplikat
		if selected[realID] {
			return nil, fmt.Errorf("duplicate category selection: %d", index)
		}

		selected[realID] = true
		ids = append(ids, realID)
	}

	return ids, nil
}

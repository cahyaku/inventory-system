package controller

import (
	"fmt"
	"inventory-system/internal/service"
	"inventory-system/internal/utils"
	"strconv"
	"strings"
)

type ItemController struct {
	service *service.ItemService
}

func NewItemController(service *service.ItemService) *ItemController {
	return &ItemController{service: service}
}

func (c *ItemController) ShowItems() {
	items, _ := c.service.GetAllItems()

	if len(items) == 0 {
		fmt.Println("No items available.")
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
	name := utils.ReadLine("Item name: ")
	stock, err := utils.ReadInt("Initial stock: ")
	if err != nil {
		fmt.Println("Stock must be a number")
		return
	}

	fmt.Println("Select category IDs (comma separated, e.g 1,2):")
	input := utils.ReadLine(">> ")

	parts := strings.Split(input, ",")
	var ids []int

	for _, part := range parts {
		id, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			fmt.Println("Invalid category ID:", part)
			return
		}
		ids = append(ids, id)
	}

	err = c.service.CreateItem(name, stock, ids)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("Item created successfully ✅")
}

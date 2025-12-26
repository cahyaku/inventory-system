package controller

import (
	"fmt"
	"inventory-system/internal/entity"
	"inventory-system/internal/service"
	"inventory-system/internal/utils"
	"strconv"
	"strings"
)

// ItemController membuat controller untuk item
type ItemController struct {
	itemService     *service.ItemService
	categoryService *service.CategoryService
}

// NewItemController menginisialisasi controller item (constructor)
func NewItemController(itemService *service.ItemService, categoryService *service.CategoryService) *ItemController {
	return &ItemController{
		itemService:     itemService,
		categoryService: categoryService,
	}
}

// ShowItems menampilkan semua item dan kategori masing-masing.
func (c *ItemController) ShowItems() {
	items, _ := c.itemService.GetAllItems()

	if len(items) == 0 {
		fmt.Println("No items available.")
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("====== Show Items =====")
	for i, item := range items {
		fmt.Printf("%d. %s - Stock: %d - Categories: ", i+1, item.Name, item.Stock)
		for i, cat := range item.Categories {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(cat.Name)
		}
		fmt.Println()
	}
	utils.PressEnterToContinue()
}

// CreateItem membuat item baru dengan nama, stock, dan kategori yang dipilih.
func (c *ItemController) CreateItem() {
	categories, err := c.categoryService.GetAll()
	if err != nil || len(categories) == 0 {
		fmt.Println("Cannot create item. No categories available.")
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("====== Create Item =====")

	// Baca input user dan parsing
	name, stock, ids, err := c.readItemInput()
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	if err := c.itemService.Create(name, stock, ids); err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("Item (", name, ") created successfully ✅")
	utils.PressEnterToContinue()
}

// UpdateItem memperbarui item yang dipilih dengan nama, stock, dan kategori yang dipilih.
func (c *ItemController) UpdateItem() {
	item, err := c.selectItem("====== Update Item =====")
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	name, stock, ids, err := c.readItemInput()
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	if !utils.Confirm("Are you sure want to update item?") {
		fmt.Println("Edit item cancelled")
		return
	}

	if err := c.itemService.Update(item.ID, name, stock, ids); err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("Item updated successfully ✅")
	utils.PressEnterToContinue()
}

// DeleteItem menghapus item yang dipilih.
func (c *ItemController) DeleteItem() {
	item, err := c.selectItem("====== Delete Item =====")
	if err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	if !utils.Confirm("Are you sure want to delete item?") {
		fmt.Println("Delete item cancelled")
		return
	}

	if err := c.itemService.Delete(item.ID); err != nil {
		fmt.Println("Error:", err.Error())
		utils.PressEnterToContinue()
		return
	}

	fmt.Println("Item deleted successfully ✅")
	utils.PressEnterToContinue()
}

// readCategoryIDs membaca input user untuk memilih kategori.
func (c *ItemController) readCategoryIDs() ([]int, error) {
	categories, err := c.categoryService.GetAll()
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, fmt.Errorf("no categories available")
	}

	fmt.Println("---------------------")
	fmt.Println("Available categories:")
	// loop tampilkan semua kategori
	for i, cat := range categories {
		fmt.Printf("%d. %s\n", i+1, cat.Name)
	}

	// user memilih kategori nomor
	input := utils.ReadLine("Select category numbers (comma separated, e.g 1,2,3...): ")
	parts := strings.Split(input, ",")

	if len(parts) == 0 {
		return nil, fmt.Errorf("you must select at least one category")
	}

	// menyimpan semua id kategori yang dipilih (id aslinya bukan no urut yang dipilih)
	var ids []int

	selected := make(map[int]bool) // agar input tidak duplikat
	// Jadi pada selected key: realID kategori(ex. 5, 8, 12) dengan value true.

	for _, part := range parts {
		index, err := strconv.Atoi(strings.TrimSpace(part)) // konversi string ke int
		if err != nil {
			return nil, fmt.Errorf("invalid category number: %s", part)
		}

		// VALIDASI INDEX range (supaya user tidak memasukkan index yang salah => memilih kategori yang tidak ada)
		if index < 1 || index > len(categories) {
			return nil, fmt.Errorf("category number %d is out of range", index)
		}

		realID := categories[index-1].ID // karena user lihat dari 1, sedangkan slice dimulai dari 0

		// Cegah duplikat
		if selected[realID] {
			return nil, fmt.Errorf("duplicate category selection: %d", index)
		}

		selected[realID] = true   // tandai sudah dipilih
		ids = append(ids, realID) // simpan id kategori
	}
	return ids, nil
}

// selectItem membaca input user untuk memilih item.
func (c *ItemController) selectItem(prompt string) (*entity.Item, error) {
	items, err := c.itemService.GetAllItems()
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("no items available")
	}

	fmt.Println(prompt)
	for i, item := range items {
		var categories []string
		for _, cat := range item.Categories {
			categories = append(categories, cat.Name)
		}
		fmt.Printf("%d. %s (Stock: %d) Categories: %s\n", i+1, item.Name, item.Stock, strings.Join(categories, ", "))
	}

	index, err := utils.ReadInt("Select item number: ")
	if err != nil || index < 1 || index > len(items) {
		return nil, fmt.Errorf("item number %d is out of range", index)
	}

	//fmt.Println("DEBUG ITEM ID:", items[index-1].ID)

	return &items[index-1], nil
}

// readItemInput membaca input user untuk membuat item baru.
func (c *ItemController) readItemInput() (string, int, []int, error) {
	name := utils.ReadLine("Item name: ")
	stock := readValidStock()
	categoryIDs := c.readCategoryInput()

	return name, stock, categoryIDs, nil
}

// readValidStock membaca input user untuk stock, hanya jika stock >= 0
func readValidStock() int {
	for {
		stock, err := utils.ReadInt("Stock: ")
		if err == nil && stock >= 0 {
			return stock
		}
		fmt.Println("Stock must be zero or greater. Please try again.")
	}
}

// readCategoryInput membaca input user untuk memilih kategori.
func (c *ItemController) readCategoryInput() []int {
	for {
		ids, err := c.readCategoryIDs()
		if err == nil {
			return ids
		}
		fmt.Println("Invalid category selection. Please try again.", err.Error())
	}
}

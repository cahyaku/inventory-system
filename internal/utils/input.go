package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * Ini untuk mendapatkan jumlah inputan lebih dari 1 kata.
 */
var reader = bufio.NewReader(os.Stdin)

func ReadLine(prompt string) string {
	// Perulangan terjadi selama iput belum valid
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n') // membaca jika tekan enter
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input cannot be empty, please try again.")
			continue
		}

		return input
	}
}

func ReadInt(prompt string) (int, error) {
	input := ReadLine(prompt)
	return strconv.Atoi(input)
}

func PressEnterToContinue() {
	fmt.Println()
	fmt.Print("Press Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

// Confirm menampilkan pertanyaan konfirmasi (y/n)
// return true jika user setuju, false jika batal
func Confirm(prompt string) bool {
	for {
		fmt.Printf("%s (y/n): ", prompt)
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		switch input {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Input tidak valid. Masukkan y atau n.")
		}
	}
}

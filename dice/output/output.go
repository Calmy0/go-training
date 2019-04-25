package output

import "fmt"

/* prints 12-numbers row, with fixed column spacing */
func PrintRow(arr []int) {
	for _, value := range arr {
		fmt.Printf("%-6d", value)
	}
	print("\n")
}

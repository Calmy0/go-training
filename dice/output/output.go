package output

import "fmt"

/* prints 12-numbers row, with fixed column spacing */
func PrintRow(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%-6d", arr[i])
	}
	print("\n")
}

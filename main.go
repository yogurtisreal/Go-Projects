// Create package
package main

// Our Imports
import(
	"fmt"
	"unsafe"
)

// Main Code Execution

func main() {
	var x int8
	fmt.Println(unsafe.Sizeof(x), "Byte(s)")
}

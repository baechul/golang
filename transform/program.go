package transform

// When a package inside the module is dependent on another package of the same module,
// you can simply provide the full URL just like what client app does.
// Go is smart enough not to download it but simply refer the package from the same module source code
import (
	"fmt"
	"github.com/baechul/golang/calc"
)

// SomeFunc xxx
func SomeFunc() {
	fmt.Println(calc.Add(1, 2))
}
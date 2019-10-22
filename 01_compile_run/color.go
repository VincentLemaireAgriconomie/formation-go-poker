package main





import (
	color "github.com/logrusorgru/aurora"
	"fmt"
)

func main() {
	   fmt.Println(    "Hello,",     color.Magenta("Aurora")   )
	fmt.Println(color.Bold(    color.Cyan(     "Cya!")))
}

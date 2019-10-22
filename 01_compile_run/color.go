package main

import (
	"fmt"
	color "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Println("Hello,", color.Magenta("Aurora"))
	fmt.Println(color.Bold(color.Cyan("Cya!")))
}

package main

import (
	"fmt"

	"github.com/KznRkjp/go-keeper.git/internal/buildinfo"
)

func main() {
	buildinfo.PrintBuildVersionDate()
	fmt.Println("pass")

}

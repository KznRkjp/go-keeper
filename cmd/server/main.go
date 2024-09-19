package main

import (
	"github.com/KznRkjp/go-keeper.git/internal/buildinfo"
	"github.com/KznRkjp/go-keeper.git/internal/flags"
)

func main() {
	flags.ParseFlags()
	buildinfo.PrintBuildVersionDate()

}

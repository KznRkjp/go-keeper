package buildinfo

import (
	"fmt"
	"time"

	"github.com/KznRkjp/go-keeper.git/internal/flags"
)

var BuildVersion string
var BuildDate string

func PrintBuildVersionDate() {

	if flags.FlagBuildVersion != "" {
		BuildVersion = flags.FlagBuildVersion
	} else {
		BuildVersion = "0.0.0-a.1"
	}
	BuildDate = time.Now().Format("2006-01-02")
	fmt.Println("Build date:", BuildDate, "\nVersion:", BuildVersion)

}

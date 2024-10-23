package buildinfo

import (
	"fmt"
	"time"

	"github.com/KznRkjp/go-keeper.git/internal/flags"
)

var BuildVersion string
var BuildDate string

//  Печатается билд и дата, при этом данные для надо конено передавать через билд пайплайном
// примерно так: https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
func PrintBuildVersionDate() {

	if flags.FlagBuildVersion != "" {
		BuildVersion = flags.FlagBuildVersion
	} else {
		BuildVersion = "0.0.0-a.1"
	}
	BuildDate = time.Now().Format("2006-01-02")
	fmt.Println("Build date:", BuildDate, "\nVersion:", BuildVersion)

}

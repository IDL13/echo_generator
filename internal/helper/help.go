package helper

import (
	"codegenerator/pkg/logger"
	"fmt"
	"os"
)

func Help() {
	logg := logger.NewLogger("../../log.txt")
	help, err := os.ReadFile("../helper.txt")
	if err != nil {
		logg.Write("Error from helper file", err)
	}
	fmt.Println(string(help))
}

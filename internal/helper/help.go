package helper

import (
	"fmt"
	"log"
	"os"
)

func Help() {
	help, err := os.ReadFile("../helper.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(help))
}

package make_executor

import (
	_const "codegenerator/internal/makefile/const"
	"fmt"
	"os"
)

func CreateNewMakefile() {
	f, err := os.OpenFile("./Makefile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(fmt.Sprintf(_const.MakefileBase))
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}

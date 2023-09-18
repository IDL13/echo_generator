package make_executor

import (
	_const "codegenerator/internal/makefile/const"
	"codegenerator/pkg/logger"
	"fmt"
	"os"
)

func CreateNewMakefile() {
	logg := logger.NewLogger("../../log.txt")
	f, err := os.OpenFile("./Makefile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		logg.Write("Error from echo_executor file", err)
	}
	_, err = f.WriteString(fmt.Sprintf(_const.MakefileBase))
	if err != nil {
		logg.Write("Error from echo_executor file", err)
	}
	defer f.Close().Error()
}

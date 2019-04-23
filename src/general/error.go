package general

import (
	"fmt"
)

//CheckError - funcao generia para apresentar mensagem em caso de erro
func CheckError(err error, msg string) bool {
	hasError := (err != nil)

	if hasError {
		if msg == "" {
			fmt.Println(err)
		} else {
			fmt.Println(msg)
		}
	}
	return hasError
}

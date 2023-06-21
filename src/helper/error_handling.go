package helper

import "fmt"

func ErrorNotNil(err error, message string) {
	if err != nil {
		fmt.Println(err, "<< error")
		Logger.Warn().Err(err).Msg(message)
		return
	}
}

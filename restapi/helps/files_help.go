package helps

import (
	"log"
	"os"
)

func GetCurrentPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err.Error(), "err.Error() helps/files_help.go:11")
		return ""
	}

	return pwd
}

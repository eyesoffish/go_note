package util

import (
	"log"
	"os"
)

var  (
	INFO *log.Logger
	ERROR *log.Logger
	WARN *log.Logger
)

func init() {
	
	file, err := os.OpenFile(".log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	INFO = log.New(file, "INFO: ", log.LstdFlags | log.Llongfile);
	ERROR = log.New(file, "ERROR: ", log.LstdFlags | log.Llongfile);
	WARN = log.New(file, "WARN: ", log.LstdFlags | log.Llongfile);
}
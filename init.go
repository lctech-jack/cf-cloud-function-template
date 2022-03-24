package cloudfunc

import (
	"log"
	"os"
)

// custom logger
var logger *log.Logger

// read shared env (optional)
// var envVar = os.Getenv("SOME_VAR")

// write shared initialization here
func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

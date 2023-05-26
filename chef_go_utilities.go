package main //chef_go_utilities

import (
	// local references
	"chef_go_server_utilities/pkg/logging"

	// GoLang standard libraries
	"fmt"
)

func main () {
	myLog := logging.StdOutLogger{Val: 7}

	a := myLog.Log()
	fmt.Println(a)
}

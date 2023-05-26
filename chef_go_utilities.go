package main //chef_go_utilities

import (
	"chef_go_utilities/pkg/logging"
	"fmt"
)

func main () {
	myLog := logging.StdOutLogger{Val: 7}

	a := myLog.Log()
	fmt.Println(a)
}

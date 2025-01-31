package check

import (
	"development/util"
	"fmt"
)

func CheckPackages() {
	path := util.GetPath()
	fmt.Println("Checking everything in ", path)
}

package cli_helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DigiConvent/testd9t/core/db"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func ListPackages(sysService sys_service.SysServiceInterface) {
	dbPackages, _ := sysService.GetPackages()
	packages := db.ListPackages()

	cols := []int{9, 13, 9, 8}
	fmt.Printf("┏%s┳%s┳%s┳%s┓\n", strings.Repeat("━", cols[0]), strings.Repeat("━", cols[1]), strings.Repeat("━", cols[2]), strings.Repeat("━", cols[3]))
	fmt.Println("┃ Package ┃ Initialised ┃ Version ┃ Max V. ┃")

	for _, pkg := range packages {
		fmt.Printf("┣%s╋%s╋%s╋%s┫\n", strings.Repeat("━", cols[0]), strings.Repeat("━", cols[1]), strings.Repeat("━", cols[2]), strings.Repeat("━", cols[3]))
		format := "┃ %" + strconv.Itoa(cols[0]-2) + "s ┃ %" + strconv.Itoa(cols[1]-2) + "s ┃ %" + strconv.Itoa(cols[2]-2) + "s ┃ %" + strconv.Itoa(cols[3]-2) + "s ┃\n"
		val, ok := dbPackages[pkg]
		initialised := "✓"
		var version, maxVersion string
		versions, _ := sysService.GetPackageVersions(pkg)
		sys_domain.Sort(versions, true)
		if len(versions) > 0 {
			maxVersion = versions[len(versions)-1].String()
		}
		if !ok {
			initialised = "✗"
		} else {
			version = val.Version.String()
		}
		fmt.Printf(format, pkg, initialised, version, maxVersion)
	}

	fmt.Printf("┗%s┻%s┻%s┻%s┛\n", strings.Repeat("━", cols[0]), strings.Repeat("━", cols[1]), strings.Repeat("━", cols[2]), strings.Repeat("━", cols[3]))
}

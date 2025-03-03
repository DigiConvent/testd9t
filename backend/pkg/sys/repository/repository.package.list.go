package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) GetPackages() ([]sys_domain.Package, core.Status) {
	res, err := r.db.Query("select name, major, minor, patch from packages")
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	var packages []sys_domain.Package
	for res.Next() {
		var p sys_domain.Package
		err = res.Scan(&p.Name, &p.Version.Major, &p.Version.Minor, &p.Version.Patch)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		packages = append(packages, p)
	}

	return packages, *core.StatusSuccess()
}

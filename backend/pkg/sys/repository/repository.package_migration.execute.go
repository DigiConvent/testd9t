package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/log"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) MigratePackage(pkgName string, toVersion *sys_domain.Version) core.Status {
	pkgs, _ := r.GetPackages()
	fromVersion := sys_domain.Version{Major: -1, Minor: -1, Patch: -1}
	for i := range pkgs {
		pkg := pkgs[i]
		if pkg.Name == pkgName {
			fromVersion = pkg.Version
		}
	}

	versions, status := r.GetPackageVersions(pkgName)
	if status.Err() {
		return status
	}

	sys_domain.Sort(versions, true)

	versionsToInstall := make([]sys_domain.Version, 0)

	for i := range versions {
		version := versions[i]
		if version.SmallerThan(&fromVersion) {
			continue
		}
		if version.Equals(&fromVersion) {
			continue
		}
		if toVersion.SmallerThan(&version) {
			continue
		}
		versionsToInstall = append(versionsToInstall, version)
	}

	for i := range versionsToInstall {
		version := versionsToInstall[i]
		script, status := r.GetPackageMigrationScript(pkgName, &version)
		if status.Err() {
			return status
		}

		conn := db.NewSqliteDB(pkgName)

		_, err := conn.Exec(script)
		if err != nil {
			return *core.InternalError("Could not execute migration script for package " + pkgName + " version " + version.String() + ": " + err.Error())
		} else {
			log.Success("Migrated package " + pkgName + " from " + fromVersion.String() + " to " + version.String())
		}
	}

	if fromVersion.Major == -1 {
		_, err := r.db.Exec("INSERT INTO packages (name, major, minor, patch) VALUES (?, ?, ?, ?)", pkgName, toVersion.Major, toVersion.Minor, toVersion.Patch)
		if err != nil {
			return *core.InternalError("Could not register package " + pkgName + " version " + toVersion.String())
		}
	} else {
		_, err := r.db.Exec("UPDATE packages SET major = ?, minor = ?, patch = ? WHERE name = ?", toVersion.Major, toVersion.Minor, toVersion.Patch, pkgName)
		if err != nil {
			return *core.InternalError("Could not update package " + pkgName + " version to " + toVersion.String())
		} else {
			log.Success("Updated package " + pkgName + " to version " + toVersion.String())
		}
	}

	return *core.StatusSuccess()
}

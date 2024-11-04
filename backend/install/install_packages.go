package install

import (
	"fmt"

	installation_utils "github.com/DigiConvent/testd9t/install/utils"
)

func PostgresInstalled() bool {
	return installation_utils.Execute("which psql", true) != ""
}

func InstallPostgres(force bool) {
	if PostgresInstalled() && !force {
		fmt.Println("Postgres is already installed, ignoring\nUse --force to reinstall.")
		return
	}
	installation_utils.Execute("sudo apt update -y", false)
	installation_utils.Execute("sudo apt install -y ca-certificates gpg certbot", false)
	installation_utils.Execute("sudo sh -c `echo \"deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main\" > /etc/apt/sources.list.d/pgdg.list`", false)
	installation_utils.Execute("wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc", false)
	installation_utils.Execute("sudo apt-key add -", false)
	installation_utils.Execute("sudo apt update -y", false)
	installation_utils.Execute("sudo apt install -y postgresql postgresql-contrib", false)
	StartPostgres()
}

func StartPostgres() {
	installation_utils.Execute("sudo systemctl enable postgresql", false)
	installation_utils.Execute("sudo systemctl start postgresql", false)
	fmt.Println("Postgres enabled and started")
}

func StopPostgres() {
	installation_utils.Execute("sudo systemctl stop postgresql", false)
	installation_utils.Execute("sudo systemctl disable postgresql", false)
	fmt.Println("Postgres disabled and stopped")
}

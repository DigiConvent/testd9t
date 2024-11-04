package install

func Install(forceFlag bool) {
	CreateUser()
	InstallPostgres(forceFlag)
}

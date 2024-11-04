package version

type SemVerMigrations struct {
	Version    SemVer            `json:"version"`    // 0.0.0
	Migrations map[string]string `json:"migrations"` // 01_init.sql -> create table if not exists versions ....
}

type Migrator interface {
	ListMigrationVersions() []SemVer
	ListMigrationFiles(version *SemVer) map[string]string
}

type remoteMigrationsMigrator struct{}
type localMigrationsMigrator struct{}

func NewRemoteMigrationsMigrator() Migrator {
	return &remoteMigrationsMigrator{}
}

func NewLocalMigrationsMigrator() Migrator {
	return &localMigrationsMigrator{}
}

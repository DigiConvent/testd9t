package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) SetMaintenance(maintenance bool) core.Status {
	_, err := r.db.Exec("update configurations set maintenance = ?", maintenance)

	if err != nil {
		return *core.InternalError("Failed to set maintenance flag")
	}
	sys_domain.SetMaintenance(maintenance)

	return *core.StatusSuccess()
}

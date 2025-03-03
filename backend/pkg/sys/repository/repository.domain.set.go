package sys_repository

import "github.com/DigiConvent/testd9t/core"

// SetDomain implements SysRepositoryInterface.
func (r *SysRepository) SetDomain(domain string) core.Status {
	result, err := r.db.Exec("update configurations set domain = ?", domain)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return *core.NotFoundError("configuration not found")
	}

	return *core.StatusNoContent()
}

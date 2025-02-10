package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/jackc/pgtype"
)

func (r *SysRepository) ClaimAdmin(tgId string) core.Status {
	res := r.db.QueryRow(`SELECT id FROM users WHERE super = true`)

	var superId string
	err := res.Scan(&superId)
	if err == nil {
		return *core.BadRequestError("Super user already exists")
	}

	var exists int
	res = r.db.QueryRow(`SELECT count(*) FROM users WHERE telegram_id = ?`, tgId)
	err = res.Scan(&exists)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	var userId string
	if exists > 0 {
		update, err := r.db.Exec(`UPDATE users SET super = true WHERE telegram_id = ?`, tgId)
		if err != nil {
			return *core.InternalError(err.Error())
		}
		rowsAffected, _ := update.RowsAffected()
		if rowsAffected == 0 {
			return *core.NotFoundError("User not found")
		}
		row := r.db.QueryRow(`SELECT id FROM users WHERE telegram_id = ?`, tgId)
		err = row.Scan(&userId)
		if err != nil {
			return *core.InternalError(err.Error())
		}
	} else {
		row := r.db.QueryRow(`INSERT INTO users (telegram_id, super) VALUES (?, true) returning id`, tgId)
		err = row.Scan(&userId)
		if err != nil {
			return *core.InternalError(err.Error())
		}
	}

	var superPermissionGroupId pgtype.UUID

	row := r.db.QueryRow(`SELECT id FROM permission_groups WHERE name = 'super'`)

	err = row.Scan(&superPermissionGroupId)

	if err != nil {
		return *core.NotFoundError("Super permission group does not exist")
	}

	_, err = r.db.Exec(`INSERT INTO permission_group_has_user (permission_group, "user") VALUES (?, ?)`, superPermissionGroupId, userId)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusSuccess()
}

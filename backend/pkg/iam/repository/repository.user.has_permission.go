package iam_repository

import (
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) UserHasPermission(userId *uuid.UUID, permission string) bool {
	var hasPermission bool
	row := r.db.QueryRow(`select 1 from user_has_permissions where user = ? and permission = ?`, userId.String(), permission)

	err := row.Scan(&hasPermission)

	if err != nil {
		return false
	}
	return hasPermission
}

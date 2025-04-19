package iam_repository

import (
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) UserHasPermission(userId *uuid.UUID, permission string) bool {
	var hasPermission bool
	permissionLike := permission + ".%"
	row := r.db.QueryRow(`select 1 from user_has_permissions where user = ? and (permission = ? or permission like ?)`, userId.String(), permission, permissionLike)

	err := row.Scan(&hasPermission)

	if err != nil {
		return false
	}
	return hasPermission
}

package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) UserHasPermission(userId *uuid.UUID, permission string) (bool, core.Status) {
	var hasPermission bool
	err := r.DB.QueryRow(`
with recursive relevant_groups as (
  select permission_group as id from permission_group_has_permission 
	join permissions on permission_group_has_permission.permission = permissions.id
	where permissions.name = ?
  union all
  select child.id as permission_group from permission_groups child
    inner join relevant_groups s on s.id = child.parent
)
select true as result from permission_group_has_user pghu
join relevant_groups on pghu.permission_group = relevant_groups.id
where pghu."user" = ?;`, userId.String(), permission).Scan(&hasPermission)

	if err != nil {
		return false, *core.InternalError(err.Error())
	}
	return hasPermission, *core.StatusSuccess()
}

package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListStatusesFromUser(id *uuid.UUID) ([]*iam_domain.UserBecameStatusRead, core.Status) {
	var userBecameStatusses []*iam_domain.UserBecameStatusRead

	// 	r.db.QueryDebug(`select s.id, uf.id,  uf.first_name, uf.last_name, ubs.description, ubs.start, ubs."end"
	//  from permission_group_has_user ubs
	//  join user_facades uf on ubs."user" = uf.id
	//  join permission_groups s on s.meta = 'status' and ubs.status = s.id
	//  where ubs."user" = ?`, id.String())
	rows, err := r.db.Query(`select s.id, uf.id,  uf.first_name, uf.last_name, ubs.comment, ubs.start, ubs."end" 
 from permission_group_has_user ubs 
 join user_facades uf on ubs."user" = uf.id 
 join permission_groups s on s.meta = 'status' and ubs.permission_group = s.id 
 where ubs."user" = ?`, id.String())
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		ubs := &iam_domain.UserBecameStatusRead{
			User: iam_domain.UserFacade{},
		}

		err := rows.Scan(&ubs.UserStatus, &ubs.User.Id, &ubs.User.FirstName, &ubs.User.LastName, &ubs.Comment, &ubs.Start, &ubs.End)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		userBecameStatusses = append(userBecameStatusses, ubs)
	}

	return userBecameStatusses, *core.StatusSuccess()
}

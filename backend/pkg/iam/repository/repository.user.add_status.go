package iam_repository

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) AddUserStatusToUser(d *iam_domain.AddUserStatusToUser) core.Status {
	fmt.Println(d.When)
	_, err := r.db.Exec(`insert into user_became_status ("user", status, start) values (?, ?, ?)`, d.UserID, d.StatusID, d.When)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusSuccess()
}

package iam_domain

type PermissionWrite struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Meta        string `json:"meta"`
}

type PermissionRead struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Meta        string `json:"meta"`
	Archived    bool   `json:"archived"`
}

type PermissionProfile struct {
	Permission       *PermissionRead          `json:"permission"`
	Descendants      []*PermissionFacade      `json:"descendants"`
	PermissionGroups []*PermissionGroupFacade `json:"permission_groups"`
	Users            []*UserFacade            `json:"users"`
}

type PermissionFacade struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Implied     bool   `json:"implied"`
	Meta        string `json:"meta"`
}

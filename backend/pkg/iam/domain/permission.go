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
	Generated   bool   `json:"generated"`
	Archived    bool   `json:"archived"`
}

type PermissionProfile struct {
	Permission       *PermissionRead          `json:"permission"`
	PermissionGroups []*PermissionGroupFacade `json:"permission_groups"`
	Users            []*UserFacade            `json:"users"`
}

type PermissionFacade struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Implied     bool   `json:"implied"`
}

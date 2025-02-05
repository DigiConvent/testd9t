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
}

type PermissionFacade struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Implied     bool   `json:"implied"`
}

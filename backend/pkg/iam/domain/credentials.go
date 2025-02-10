package iam_domain

type PasswordReset struct {
	Email string
	Code  string
}

type SetPassword struct {
	Email    string
	Code     string
	Password string
}

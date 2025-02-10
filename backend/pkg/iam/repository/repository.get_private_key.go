package iam_repository

import "crypto/rsa"

func (r *IAMRepository) GetPrivateKey() *rsa.PrivateKey {
	return &r.privateKey
}

package core

import (
	"context"
)

type (
	// Cert  struct
	cert struct {
		Issuer    string `json:"issuer"`
		CreateAt  string `json:"create_at"`
		EndAt     string `json:"end_at"`
		ExpiredAt int    `json:"expired_at"`
	}
	// CertInfo .
	CertInfo struct {
		Info map[string]*cert `json:"host"`
	}

	// HostCert host cert info check
	HostCert interface {
		Create(context.Context, string) error
		Check(context.Context) error
		Info(context.Context) *CertInfo
		Run(context.Context, string) error
	}
)

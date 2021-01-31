package core

import (
	"context"
)

type (
	// Cert  struct
	cert struct {
		Issuer    string  `json:"Issuer"`
		CreateAt  int64   `json:"create_at"`
		EndAt     int64   `json:"end_at"`
		ExpiredAt float64 `json:"expired_at"`
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

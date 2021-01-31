package core

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"sync"
	"time"
)

var errNotFound = errors.New("domain not founded in memory")

type host struct {
	sync.Mutex
	hosts map[string]*cert
}

// New return a new host
func New() HostCert {
	return &host{
		hosts: make(map[string]*cert),
	}
}

func (h *host) Create(ctx context.Context, domain string) error {
	h.Lock()
	h.hosts[domain] = &cert{}
	h.Unlock()
	return nil
}

func (h *host) Check(ctx context.Context) error {
	h.Lock()
	defer h.Unlock()
	if len(h.hosts) == 0 {
		return errNotFound
	}

	// cert info
	for host := range h.hosts {
		conn, err := tls.Dial("tcp", host+":443", nil)
		if err != nil {
			log.Printf("host: %s cannot create a connect, error: %s\n", host, err)
			continue
		}

		defer conn.Close()

		err = conn.VerifyHostname(host)
		if err != nil {
			log.Printf("Hostname doesn't match with certificate: %s\n", err.Error())
		}

		cert := conn.ConnectionState().PeerCertificates[0]
		timeNow := time.Now()
		h.hosts[host].CreateAt = cert.NotBefore.Local().Format(time.RFC850)
		h.hosts[host].EndAt = cert.NotAfter.Local().Format(time.RFC850)
		h.hosts[host].Issuer = cert.Issuer.String()
		h.hosts[host].ExpiredAt = int(cert.NotAfter.Sub(timeNow).Hours())
	}

	return nil
}

func (h *host) Info(ctx context.Context) *CertInfo {
	h.Lock()
	defer h.Unlock()

	info := &CertInfo{
		Info: make(map[string]*cert),
	}

	for host, certinfo := range h.hosts {
		info.Info[host] = certinfo
	}

	return info
}

func (h *host) Run(ctx context.Context, location string) error {
	if location == "" {
		return errNotFound
	}
	// load config
	for _, domain := range read(location) {
		h.Create(ctx, domain)
	}

	h.Check(ctx)
	// add cron job
	return nil
}

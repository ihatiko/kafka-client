package kafka_client

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/ihatiko/log"
)

func (c *Config) NewTlsConfig() *tls.Config {
	if c.UseSSL {
		caCertPool := x509.NewCertPool()
		ok := caCertPool.AppendCertsFromPEM([]byte(c.SslCaPem))
		if !ok {
			log.FatalF("SSL pem error %v", ok)
		}
		return &tls.Config{
			RootCAs:            caCertPool,
			InsecureSkipVerify: true,
			MinVersion:         tls.VersionTLS12,
		}
	}
	return &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
}

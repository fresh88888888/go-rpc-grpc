package help

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"google.golang.org/grpc/credentials"
)

func GetClientCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	cert_pool := x509.NewCertPool()
	ca, _ := os.ReadFile("cert/ca.crt")
	cert_pool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "www.umbrella.com",
		RootCAs:      cert_pool,
	})

	return creds
}

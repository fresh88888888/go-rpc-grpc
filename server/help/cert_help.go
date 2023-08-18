package help

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"

	"google.golang.org/grpc/credentials"
)

func GetServerCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	cert_pool := x509.NewCertPool()
	ca, _ := os.ReadFile("cert/server.crt")
	cert_pool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    cert_pool,
	})

	return creds
}

func GetClientCreds() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("../server/cert/client.pem", "../server/cert/client.key")
	if err != nil {
		log.Fatal(err)
	}
	cert_pool := x509.NewCertPool()
	ca, err := os.ReadFile("../server/cert/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	cert_pool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "www.umbrella.com",
		RootCAs:      cert_pool,
	})

	return creds
}

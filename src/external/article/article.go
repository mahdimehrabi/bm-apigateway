package article

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
)

const articleServer = "localhost:6001"

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pmServerCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pmServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	//load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientCert},
	}

	return credentials.NewTLS(config), nil
}

type ArticleGRPC struct {
	CC *grpc.ClientConn
}

func (a *ArticleGRPC) Connect() (err error) {
	tlsCredential, err := loadTLSCredentials()
	a.CC, err = grpc.Dial(articleServer, grpc.WithTransportCredentials(tlsCredential))
	if err != nil {
		return err
	}
	return nil
}

func (a *ArticleGRPC) Close() error {
	return a.CC.Close()
}

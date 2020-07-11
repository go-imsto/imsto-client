package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"
	"sync"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials"
)

var (
	address = defaultAddr
	conn    *grpc.ClientConn
	onceCC  sync.Once

	retryOptMax = grpc_retry.WithMax(maxRetries)
)

func init() {
	addr := os.Getenv(envAddr)
	if addr != "" {
		address = addr
	}
}

// defaultConn ...
func defaultConn() *grpc.ClientConn {
	onceCC.Do(func() {
		var err error
		conn, err = newClientConn()
		if err != nil {
			logger().Fatalw("get client conn fail", "err", err)
		}
		logger().Infow("connection", "state", conn.GetState())
	})
	return conn
}

func loadTrCr(caCrt, clientCrt, clientKey, serverName string) credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair(clientCrt, clientKey)
	if err != nil {
		logger().Fatalw("tls.LoadX509KeyPair fail", "err", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caCrt)
	if err != nil {
		logger().Fatalw("ioutil.ReadFile fail", "err", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		logger().Fatalw("certPool.AppendCertsFromPEM fail")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   serverName,
		RootCAs:      certPool,
	})
	return c
}

// SetAddr set connection address
func SetAddr(addr string) {
	address = addr
}

func newClientConn() (*grpc.ClientConn, error) {

	caCrt := os.Getenv(envCaCrt)
	clientCrt := os.Getenv(envClientCrt)
	clientKey := os.Getenv(envClientKey)
	serverName := os.Getenv(envServerName)

	var trcr credentials.TransportCredentials
	if len(caCrt) > 0 && len(clientCrt) > 0 && len(clientKey) > 0 && len(serverName) > 0 {
		logger().Infow("grpc: using TLS", "cert", clientCrt, "key", clientKey)
		trcr = loadTrCr(caCrt, clientCrt, clientKey, serverName)
	}

	hasTLS := trcr != nil
	logger().Infow("grpc: dial", "addr", address, "tls", hasTLS, "ver", grpc.Version)
	retryOpt := grpc_retry.WithBackoff(grpc_retry.BackoffExponential(backoffRetryBase))
	var retryOpts = []grpc_retry.CallOption{retryOpt, retryOptMax}

	bc := backoff.DefaultConfig
	bc.MaxDelay = backoffMaxDelay
	opts := []grpc.DialOption{
		grpc.WithConnectParams(grpc.ConnectParams{Backoff: bc}),
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(retryOpts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)),
	}
	if hasTLS {
		opts = append(opts, grpc.WithTransportCredentials(trcr))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

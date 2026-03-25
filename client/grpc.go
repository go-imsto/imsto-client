package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"

	pb "github.com/go-imsto/imsto-client/impb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = defaultAddr
	conn    *grpc.ClientConn
	client  pb.ImageSvcClient
	onceCC  sync.Once
)

// retryPolicy is the gRPC built-in retry policy configured via service config.
// MaxAttempts: maximum number of retry attempts before failing.
// InitialBackoff/MaxBackoff/BackoffMultiplier: exponential backoff configuration.
// RetryableStatusCodes: status codes that trigger retry (UNAVAILABLE in this case).
var retryPolicy = fmt.Sprintf(`{
"methodConfig":[{
  "name": [{}],
  "retryPolicy": {
    "MaxAttempts": %d,
    "InitialBackoff": "%s",
    "MaxBackoff": "%s",
    "BackoffMultiplier": 1.0,
    "RetryableStatusCodes": ["UNAVAILABLE"]
  }
}]
}`, maxRetries, backoffRetryBase, backoffMaxDelay)

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
		client = pb.NewImageSvcClient(conn)
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
	ca, err := os.ReadFile(caCrt)
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

	bc := backoff.DefaultConfig
	bc.MaxDelay = backoffMaxDelay
	opts := []grpc.DialOption{
		grpc.WithConnectParams(grpc.ConnectParams{Backoff: bc}),
		grpc.WithDefaultServiceConfig(retryPolicy),
	}
	if hasTLS {
		opts = append(opts, grpc.WithTransportCredentials(trcr))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

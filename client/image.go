//
// Package client imsto 客户端实现
//
package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/go-imsto/imsto-client/impb"
)

// 环境变量
const (
	envGRPCAddr = "IMSTO_GRPC_ADDR"

	defaultTimeout time.Duration = time.Second * 5
)

var (
	address = "127.0.0.1:12010"
	client  pb.ImageSvcClient
	conn    *grpc.ClientConn
	once    sync.Once
	trcr    credentials.TransportCredentials

	maxRetries  uint = 5
	retryOptMax      = grpc_retry.WithMax(maxRetries)
)

func init() {
	addr := os.Getenv(envGRPCAddr)
	if addr != "" {
		address = addr
	}
	caCrt := os.Getenv("IMSTO_GRPC_CA")
	clientCrt := os.Getenv("IMSTO_GRPC_CLIENT_CRT")
	clientKey := os.Getenv("IMSTO_GRPC_CLIENT_KEY")
	serverName := os.Getenv("IMSTO_GRPC_SERVER_NAME")
	if len(caCrt) > 0 && len(clientCrt) > 0 && len(clientKey) > 0 && len(serverName) > 0 {
		log.Printf("sso: using TLS %s,%s", clientCrt, clientKey)
		trcr = loadTrCr(caCrt, clientCrt, clientKey, serverName)
	}
}

func loadTrCr(caCrt, clientCrt, clientKey, serverName string) credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair(clientCrt, clientKey)
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caCrt)
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   serverName,
		RootCAs:      certPool,
	})
	return c
}

// SetAddr 设置连接地址
func SetAddr(addr string) {
	address = addr
}

func GetClient() pb.ImageSvcClient {
	once.Do(func() {
		hasTLS := trcr != nil
		log.Printf("dial to %q, TLS %v, gRPC ver %s", address, hasTLS, grpc.Version)
		retryOpt := grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond))
		retryTimesOpt := grpc_retry.WithMax(maxRetries)
		var retryOpts = []grpc_retry.CallOption{retryOpt, retryTimesOpt}

		opts := []grpc.DialOption{
			grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(retryOpts...)),
			grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)),
		}
		var err error
		if hasTLS {
			opts = append(opts, grpc.WithTransportCredentials(trcr))
		} else {
			opts = append(opts, grpc.WithInsecure())
		}

		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer cancel()
		conn, err = grpc.DialContext(ctx, address, opts...)
		if err != nil {
			panic(err)
		}
		log.Printf("connection state %s", conn.GetState())
		client = pb.NewImageSvcClient(conn)
	})

	return client
}

// Close 关闭所有连接，不是必须
func Close() error {
	if conn != nil {
		defer func() { client = nil }()
		return conn.Close()
	}
	return nil
}

// IImage ...
type IImage interface {
	GetPath() string
	GetHost() string
	GetUri() string
	GetID() uint64
}

// FetchInput ...
type FetchInput = pb.FetchInput

// ImageInput ...
type ImageInput = pb.ImageInput

// ImageOutput ...
type ImageOutput = pb.ImageOutput

// Fetch ...
func Fetch(ctx context.Context, apiKey, uri string) (IImage, error) {
	r, err := GetClient().Fetch(ctx, &pb.FetchInput{
		ApiKey: apiKey,
		Uri:    uri,
	})
	if err != nil {
		log.Printf("call Fetch(%s, %s) ERR %s", apiKey, uri, err)
		return nil, err
	}

	return r, nil
}

// Store ...
func Store(ctx context.Context, apiKey string, rd io.Reader) (IImage, error) {
	data, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, err
	}
	r, err := GetClient().Store(ctx, &pb.ImageInput{
		ApiKey: apiKey,
		Image:  data,
	})
	if err != nil {
		log.Printf("call Store(%s, %d bytes) ERR %s", apiKey, len(data), err)
		return nil, err
	}

	return r, nil
}

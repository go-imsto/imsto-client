package client

import (
	"time"
)

// consts of grpc connect
const (
	envAddr       = "IMSTO_GRPC_ADDR"
	envCaCrt      = "IMSTO_GRPC_CA"
	envClientCrt  = "IMSTO_GRPC_CLIENT_CRT"
	envClientKey  = "IMSTO_GRPC_CLIENT_KEY"
	envServerName = "IMSTO_GRPC_SERVER_NAME"
)

// consts of default
const (
	defaultAddr           = "127.0.0.1:8969"
	defaultTimeout        = time.Second * 5
	backoffMaxDelay       = time.Second * 20
	backoffRetryBase      = time.Millisecond * 100
	maxRetries       uint = 3
)

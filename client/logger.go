package client

import (
	zlog "github.com/go-imsto/imsto-client/log"
)

func logger() zlog.Logger {
	return zlog.Get()
}

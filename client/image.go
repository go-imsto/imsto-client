//
// Package client imsto 客户端实现
//
package client

import (
	"context"
	"io"
	"io/ioutil"
	"sync"

	pb "github.com/go-imsto/imsto-client/impb"
)

// vars
var (
	firstAPIKey string
	StageHost   string // Host of image uri, like CDN

	client pb.ImageSvcClient
	onceIC sync.Once
)

// GetClient ...
func GetClient() pb.ImageSvcClient {
	onceIC.Do(func() {
		client = pb.NewImageSvcClient(defaultConn())
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
	GetID() string
}

// FetchInput ...
type FetchInput = pb.FetchInput

// ImageInput ...
type ImageInput = pb.ImageInput

// ImageOutput ...
type ImageOutput = pb.ImageOutput

// MakeURL ...
func MakeURL(path, sizeOp string) string {
	if len(sizeOp) == 0 {
		sizeOp = "orig"
	}
	if len(StageHost) == 0 {
		return "/show/" + sizeOp + "/" + path
	}
	return "//" + StageHost + "/show/" + sizeOp + "/" + path
}

// Fetch ...
func Fetch(ctx context.Context, in FetchInput) (IImage, error) {
	if in.ApiKey == "" {
		in.ApiKey = firstAPIKey
	}
	if uid, ok := UIDFromContext(ctx); ok {
		in.UserID = uid
	}
	r, err := GetClient().Fetch(ctx, &in)
	if err != nil {
		logger().Infow("call Fetch() fail", "apiKey", in.ApiKey, "roof", in.Roof, "uri", in.Uri, "err", err)
		return nil, err
	}

	return r, nil
}

// Store ...
func Store(ctx context.Context, in *ImageInput) (IImage, error) {
	if in.ApiKey == "" {
		in.ApiKey = firstAPIKey
	}
	if uid, ok := UIDFromContext(ctx); ok {
		in.UserID = uid
	}
	r, err := GetClient().Store(ctx, in)
	if err != nil {
		logger().Infow("call Store() fail", "apiKey", in.ApiKey, "roof", in.Roof, "image bytes", len(in.Image), "err", err)
		return nil, err
	}
	return r, nil
}

// StoreReader ...
func StoreReader(ctx context.Context, roof, name string, rd io.Reader) (IImage, error) {
	data, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, err
	}
	return Store(ctx, &pb.ImageInput{
		Image: data,
		Roof:  roof,
		Name:  name,
	})

}

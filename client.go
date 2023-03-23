package tiktokengrpc

import (
	"context"
	"fmt"
	"time"

	"buf.build/gen/go/yorinasub17/tiktoken-grpc/grpc/go/tiktoken/v1/tiktokenv1grpc"
	tiktokenv1 "buf.build/gen/go/yorinasub17/tiktoken-grpc/protocolbuffers/go/tiktoken/v1"
	"google.golang.org/grpc"
)

// Client manages the connection to the Tiktoken gRPC service and provides the main interface for sending requests and
// receiving responses to/from the service.
type Client struct {
	c       tiktokenv1grpc.TiktokenServiceClient
	conn    *grpc.ClientConn
	timeout time.Duration
}

// NewClient initializes a new connection to the grpc service using the provided options.
func NewClient(host string, port int, opts ...grpc.DialOption) (Client, error) {
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return Client{}, err
	}
	c := tiktokenv1grpc.NewTiktokenServiceClient(conn)
	return Client{c: c, conn: conn}, nil
}

// Close tears down the client connection to the service.
func (c Client) Close() error {
	return c.conn.Close()
}

// NumTokens calls the NumTokens API of the tiktoken gRPC service for the given text, with the encoding identified by
// name.
func (c Client) NumTokens(ctx context.Context, encodingName TiktokenEncodingName, text string) (int32, error) {
	if !encodingName.IsValid() {
		return -1, InvalidEncodingNameErr(encodingName)
	}

	request := tiktokenv1.NumTokensRequest{
		Text:     text,
		Encoding: &tiktokenv1.NumTokensRequest_ByName{ByName: string(encodingName)},
	}
	resp, err := c.c.NumTokens(ctx, &request)
	if err != nil {
		return -1, err
	}
	return resp.Count, nil
}

// NumTokensByModel calls the NumTokens API of the tiktoken gRPC service for the given text, with the encoding
// identified by model name.
func (c Client) NumTokensByModel(ctx context.Context, modelName string, text string) (int32, error) {
	request := tiktokenv1.NumTokensRequest{
		Text:     text,
		Encoding: &tiktokenv1.NumTokensRequest_ByModelName{ByModelName: modelName},
	}
	resp, err := c.c.NumTokens(ctx, &request)
	if err != nil {
		return -1, err
	}
	return resp.Count, nil
}

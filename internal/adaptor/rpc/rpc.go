package rpc

import (
	"context"
	"time"
)

const (
	rpcTimeout = time.Second * 2
)

type Client struct{}

// Connect implements Rpc
func (c *Client) Connect(ctx context.Context) error {
	return nil
}

// Close implements Rpc
func (c *Client) Close() {
}

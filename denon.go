package denon

import (
	"context"
	"denon/packager"
	"denon/transport"
	"log"
	"sync"
	"time"
)

const (
	// Default timeout
	defaultCommandDelay = 250 * time.Millisecond
)

type Client struct {
	// context
	ctx         context.Context
	transporter transport.Transport
	packager    packager.Packager

	// Command delay
	Delay        time.Duration
	lastActivity time.Time

	// TCP connection
	mu sync.Mutex
}

func New(ctx context.Context, address string, opts *transport.Options) (*Client, error) {
	delay := defaultCommandDelay

	if opts != nil && opts.Delay != nil {
		delay = *opts.Delay
	}

	var logger *log.Logger
	if opts != nil && opts.Logger != nil {
		logger = opts.Logger
	}

	// Transport
	t, _ := transport.NewTcpTransport(ctx, address, opts)

	// Packager
	p := packager.NewCommand(ctx, logger)

	return &Client{
		ctx:         ctx,
		Delay:       delay,
		transporter: t,
		packager:    p,
	}, nil
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.transporter.Close()
}

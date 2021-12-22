package transport

import (
	"context"
	"log"
	"net"
	"sync"
	"time"
)

const (
	connectionTimeout     = 5 * time.Second
	connectionIdleTimeout = 60 * time.Second
	sendingDelay          = 500 * time.Millisecond
)

type tcp struct {
	ctx context.Context

	// Connect string = ip_address:port
	address string
	// Connect & Read timeout
	timeout time.Duration
	// Idle timeout to close the connection
	idleTimeout time.Duration
	// command delay timeout
	delay time.Duration

	// Logger
	logger *log.Logger

	// TCP connection
	mu           sync.Mutex
	conn         net.Conn
	lastActivity time.Time
	closeTimer   *time.Timer
}

func NewTcpTransport(ctx context.Context, address string, opts *Options) (*tcp, error) {
	timeout := connectionTimeout
	idleTimeout := connectionIdleTimeout
	delay := sendingDelay

	var logger *log.Logger

	if opts != nil {
		if opts.Delay != nil {
			delay = *opts.Delay
		}

		if opts.Timeout != nil {
			timeout = *opts.Timeout
		}

		if opts.IdleTimeout != nil {
			idleTimeout = *opts.IdleTimeout
		}

		if opts.Logger != nil {
			logger = opts.Logger
		}
	}

	return &tcp{
		ctx:         ctx,
		address:     address,
		delay:       delay,
		timeout:     timeout,
		idleTimeout: idleTimeout,
		logger:      logger,
	}, nil
}

func (t *tcp) Send(command []byte) (response []byte, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.logf("Sending command: %s", command)

	// TODO is needed?
	if t.delay > 0 {
		if t.lastActivity.Add(t.delay).After(time.Now()) {
			time.Sleep(t.delay)
		}
	}

	// Establish a new connection if not connected
	if err = t.Connect(); err != nil {
		return
	}

	// Set timer to close when idle
	t.lastActivity = time.Now()
	t.startCloseTimer()

	// Set write and read timeout
	var timeout time.Time
	if t.timeout > 0 {
		timeout = t.lastActivity.Add(t.timeout)
	}
	if err = t.conn.SetDeadline(timeout); err != nil {
		return
	}

	// Send data
	if _, err = t.conn.Write(command); err != nil {
		return
	}
	t.logf("Command sent")

	buffer := make([]byte, 135)
	size, rerr := t.conn.Read(buffer)

	t.logf("Read response size: %d\n", size)

	if size > 2 {
		response = buffer[:size-1]
	}

	err = rerr
	return
}

func (t *tcp) Connect() error {
	if t.conn == nil {
		dialer := net.Dialer{Timeout: t.timeout}
		conn, err := dialer.Dial("tcp", t.address)
		if err != nil {
			return err
		}
		t.conn = conn
	}

	return nil
}

func (t *tcp) startCloseTimer() {
	if t.idleTimeout <= 0 {
		return
	}
	if t.closeTimer == nil {
		t.closeTimer = time.AfterFunc(t.idleTimeout, t.closeIdle)
	} else {
		t.closeTimer.Reset(t.idleTimeout)
	}
}

// closeIdle closes the connection if last activity is passed behind IdleTimeout.
func (t *tcp) closeIdle() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.idleTimeout <= 0 {
		return
	}
	idle := time.Now().Sub(t.lastActivity)
	if idle >= t.idleTimeout {
		t.logf("modbus: closing connection due to idle timeout: %v", idle)
		t.Close()
	}
}

func (t *tcp) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.logf("Close connection")
	if t.conn != nil {
		return t.conn.Close()
	}

	return nil
}

func (t *tcp) logf(format string, v ...interface{}) {
	if t.logger != nil {
		t.logger.Printf(format, v...)
	}
}

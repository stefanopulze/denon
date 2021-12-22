package transport

import (
	"log"
	"time"
)

type Options struct {
	// Command Delay
	Delay *time.Duration
	// Connect & Read timeout
	Timeout *time.Duration
	// Idle timeout to close the connection
	IdleTimeout *time.Duration

	// Transmission logger
	Logger *log.Logger
}

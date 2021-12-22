package packager

import (
	"context"
	"errors"
	"log"
)

type command struct {
	ctx    context.Context
	logger *log.Logger
}

func NewCommand(ctx context.Context, logger *log.Logger) *command {
	return &command{
		ctx:    ctx,
		logger: logger,
	}
}

func (c *command) Encode(command string) ([]byte, error) {
	return []byte(command), nil
}

func (c *command) Verify(request []byte, response []byte) error {
	if len(request) < 2 {
		return errors.New("invalid request length")
	}

	if len(response) < 2 {
		return errors.New("invalid response length")
	}

	// If the Denon accept the command it return the same first two bytes
	if request[0] != response[0] || request[1] != response[1] {
		return errors.New("response is invalid for the request")
	}

	return nil
}

func (c *command) Decode(response []byte) (string, error) {
	if len(response) < 2 {
		return "", errors.New("cannot decode response, invalid length")
	}
	rc := string(response[2:])

	if c.logger != nil {
		c.logger.Printf("Decode response: %s\n", rc)
	}
	return rc, nil
}

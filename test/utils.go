package test

import (
	"context"
	"denon"
)

func DefaultDenonClient() (*denon.Client, error) {
	return denon.New(context.Background(), "192.168.1.22:23", nil)
}

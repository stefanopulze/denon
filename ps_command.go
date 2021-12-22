package denon

import (
	"fmt"
)

type DynamicVolume string

const (
	DynamicVolHeavy  DynamicVolume = "HEV"
	DynamicVolMedium               = "MED"
	DynamicVolLight                = "LIT"
	DynamicVolOff                  = "OFF"
)

// Tone Control

func (c *Client) ToneControlStatus() error {
	_, err := c.sendCommand("PSTONE CTRL ?")
	return err
}

func (c *Client) SetToneControl(on bool) error {
	_, err := c.sendCommand(fmt.Sprintf("PSTONE CTRL %s", onOff(on)))
	return err
}

// Cinema EQ

func (c *Client) CinemaEQStatus() (bool, error) {
	rs, err := c.sendCommand("PSCINEMA EQ. ?")
	return rs[10:] == "ON", err
}

func (c *Client) SetCinemaEQ(on bool) error {
	_, err := c.sendCommand(fmt.Sprintf("PSCINEMA EQ.%s", onOff(on)))
	return err
}

// Loudness Management

func (c *Client) LoudnessManagementStatus() error {
	_, err := c.sendCommand("PSLOM ?")
	return err
}

func (c *Client) SetLoudnessManagement(on bool) error {
	_, err := c.sendCommand(fmt.Sprintf("PSLOM %s", onOff(on)))
	return err
}

func (c *Client) DynamicEqStatus() (bool, error) {
	rs, err := c.sendCommand("PSDYNEQ ?")
	return rs[6:] == "ON", err
}

func (c *Client) SetDynamicEq(on bool) error {
	_, err := c.sendCommand(fmt.Sprintf("PSDYNEQ %s", onOff(on)))
	return err
}

func (c *Client) DynamicVolumeStatus() (bool, error) {
	rs, err := c.sendCommand("PSDYNVOL ?")
	return rs[6:] == "ON", err
}

func (c *Client) SetDynamicVolume(mode DynamicVolume) error {
	_, err := c.sendCommand(fmt.Sprintf("PSDYNVOL %s", mode))
	return err
}

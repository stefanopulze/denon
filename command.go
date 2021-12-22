package denon

import (
	"fmt"
	"strconv"
)

func (c *Client) IsPoweredOn() (bool, error) {
	if r, err := c.sendCommand("ZM?"); err != nil || r != "ON" {
		return false, nil
	}

	return true, nil
}

func (c *Client) SetPower(on bool) error {
	_, err := c.sendCommand(fmt.Sprintf("ZM%s", onOff(on)))
	return err
}

func (c *Client) SetAllZonePower(on bool) error {
	cmd := "PWSTANDBY"
	if on {
		cmd = "PWON"
	}

	_, err := c.sendCommand(cmd)
	return err
}

func (c *Client) VolumeStatus() (int, error) {
	r, err := c.sendCommand("MV?")
	v, _ := strconv.Atoi(r)
	return v, err
}

func (c *Client) SetVolume(vol int) error {
	v := vol

	if v < 0 {
		v = 0
	} else if v > 98 {
		v = 98
	}

	_, err := c.sendCommand(fmt.Sprintf("MV%d", v))
	return err
}

// MuteStatus indicate if Denon state is mute
func (c *Client) MuteStatus() (bool, error) {
	status, err := c.sendCommand("MU?")
	if err != nil {
		return false, err
	}

	return status == "ON", nil
}

// SetMute can mute the Denon audio
func (c *Client) SetMute(mute bool) error {
	_, err := c.sendCommand(fmt.Sprintf("MU%s", onOff(mute)))
	return err
}

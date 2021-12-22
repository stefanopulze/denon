package denon

import "time"

// sendCommand internal function to send data to denon
func (c *Client) sendCommand(command string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Wait denon process the request
	if c.lastActivity.Add(c.Delay).After(time.Now()) {
		time.Sleep(c.Delay)
	}
	c.lastActivity = time.Now()

	// Encode command
	request, err := c.packager.Encode(command)
	if err != nil {
		return "", err
	}

	// Send command
	response, err := c.transporter.Send(request)
	if err != nil {
		return "", err
	}

	// Verify command response
	if err := c.packager.Verify(request, response); err != nil {
		return "", err
	}

	// Decode command
	result, err := c.packager.Decode(response)
	if err != nil {
		return "", err
	}

	return result, nil
}

// SendCommand public
func (c *Client) SendCommand(command string) (string, error) {
	return c.sendCommand(command)
}

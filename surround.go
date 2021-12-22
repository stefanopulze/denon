package denon

type Surround string

const (
	SurroundMovie        Source = "MSMOVIE"
	SurroundMusic               = "MSMUSIC"
	SurroundGame                = "MSGAME"
	SurroundDirect              = "MSDIRECT"
	SurroundStereo              = "MSSTEREO"
	SurroundStandard            = "MSSTANDARD"
	SurroundDolbyDigital        = "MSDOLBY DIGITAL"
	SurroundDts                 = "MSDTS SURROUND"
)

func (c *Client) SetSurround(mode Surround) error {
	_, err := c.sendCommand(string(mode))
	return err
}

func (c *Client) SurroundStatus() error {
	_, err := c.sendCommand("MS?")
	return err
}

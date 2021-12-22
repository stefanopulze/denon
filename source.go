package denon

type Source string

const (
	SourceTuner       Source = "SITUNER"
	SourceDvd                = "SIDVD"
	SourceBluRay             = "SIBD"
	SourceTv                 = "SITV"
	SourceSat                = "SISAT"
	SourceMediaPlayer        = "SIMPLAY"
	SourceGame               = "SIGAME"
	SourceAux                = "SIAUX1"
	SourceSpotify            = "SPOTIFY"
)

func (c *Client) Source(source Source) error {
	_, err := c.sendCommand(string(source))
	return err
}

package test

import (
	"log"
	"testing"
)

func Test_powerOn_command(t *testing.T) {
	d, _ := DefaultDenonClient()
	defer d.Close()

	b, err := d.IsPoweredOn()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	log.Printf("Power status: %v", b)
}

func Test_powerOff_command(t *testing.T) {
	d, _ := DefaultDenonClient()
	defer d.Close()
	d.SetPower(false)
}

func Test_mute_command(t *testing.T) {
	d, _ := DefaultDenonClient()
	defer d.Close()

	mute, err := d.MuteStatus()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	log.Printf("Mute status: %v", mute)

	d.SetMute(true)
	mute, _ = d.MuteStatus()

	if !mute {
		t.Errorf("cannot set mute mode on")
		t.Fail()
	}

	d.SetMute(false)
}

func Test_volume_command(t *testing.T) {
	d, _ := DefaultDenonClient()
	defer d.Close()

	vol, err := d.VolumeStatus()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	log.Printf("Volume status: %d", vol)

	d.SetVolume(20)
	vol, _ = d.VolumeStatus()

	if vol != 20 {
		t.Errorf("cannot set volume to 20")
		t.Fail()
	}
}

func Test_cinemaeq_command(t *testing.T) {
	d, _ := DefaultDenonClient()
	defer d.Close()

	status, err := d.CinemaEQStatus()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	log.Printf("Cinema EQ status: %v", status)

	d.SetCinemaEQ(true)
	ns, _ := d.CinemaEQStatus()
	if !ns {
		t.Errorf("cannot set cinema eq to on")
		t.Fail()
	}

	d.SetCinemaEQ(status)
}

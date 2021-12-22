package test

import (
	"denon"
	"log"
	"testing"
)

func Test_dynamiceq_command(t *testing.T) {
	d, _ := DefaultDenonClient()
	defer d.Close()

	s, _ := d.DynamicEqStatus()
	log.Printf("Dynamic eq status: %v", s)
	d.SetDynamicEq(false)
}

func Test_dynamic_volume_command(t *testing.T) {
	d, _ := DefaultDenonClient()
	defer d.Close()

	s, _ := d.DynamicVolumeStatus()
	log.Printf("Dynamic volume status: %v", s)
	d.SetDynamicVolume(denon.DynamicVolMedium)
	d.SetDynamicVolume(denon.DynamicVolOff)
}

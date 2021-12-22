package main

import (
	"context"
	"denon"
	"denon/transport"
	"log"
	"os"
)

func main() {
	address := "192.168.1.22:23"
	logger := log.New(os.Stdout, "DENON ", 0)

	d, _ := denon.New(context.Background(), address, &transport.Options{
		Logger: logger,
	})
	defer d.Close()

	//powerOn, err := d.IsPoweredOn()
	//if err != nil {
	//	logger.Println(err)
	//	return
	//}
	//log.Println("Power on: ", powerOn)

	//v, err := d.VolumeStatus()
	//if err != nil {
	//	logger.Println(err)
	//	return
	//}
	//log.Println("VolumeStatus: ", v)

	//if !powerOn {
	//	if err = d.SetPower(true); err != nil {
	//		panic(err)
	//	}
	//}

	if err := d.SetSurround(denon.SurroundStandard); err != nil {
		panic(err)
	}
}

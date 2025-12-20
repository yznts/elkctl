package main

import (
	"log"

	"github.com/yznts/elkctl/ctlv2"
	"github.com/yznts/zen/v3/mapx"
	"github.com/yznts/zen/v3/slice"
)

type DeviceObj struct {
	Name    string
	Addr    string
	Enabled bool
}

type DeviceService struct{}

func (d *DeviceService) GetDevices() []DeviceObj {
	return slice.Map(mapx.Values(ctlv2.Devices), func(d *ctlv2.Device) DeviceObj {
		return DeviceObj{
			Name:    d.Name,
			Addr:    d.Addr,
			Enabled: d.Enabled,
		}
	})
}

func (d *DeviceService) AddDevice(name, addr string) {
	log.Println("Adding device:", name, addr)
	ctlv2.AddDevice(name, addr)
	// Power on and enable the device by default
	ctlv2.PowerOnDevice(name)
	ctlv2.EnableDevice(name)
}

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
	Powered bool
}

type DeviceService struct{}

func (d *DeviceService) GetDevices() []DeviceObj {
	return slice.Map(mapx.Values(ctlv2.Devices), func(d *ctlv2.Device) DeviceObj {
		return DeviceObj{
			Name:    d.Name,
			Addr:    d.Addr,
			Enabled: d.Enabled,
			Powered: d.Powered,
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

func (d *DeviceService) RemoveDevice(name string) {
	log.Println("Removing device:", name)
	ctlv2.RemoveDevice(name)
}

func (d *DeviceService) SwitchEnableState(name string) {
	if ctlv2.Devices[name].Enabled {
		log.Println("Disabling device:", name)
		ctlv2.DisableDevice(name)
	} else {
		log.Println("Enabling device:", name)
		ctlv2.EnableDevice(name)
	}
}

func (d *DeviceService) SwitchPowerState(name string) {
	if ctlv2.Devices[name].Powered {
		log.Println("Powering off device:", name)
		ctlv2.PowerOffDevice(name)
	} else {
		log.Println("Powering on device:", name)
		ctlv2.PowerOnDevice(name)
	}
}

func (d *DeviceService) SetMode(mode string) {
	log.Println("Setting mode:", mode)
	ctlv2.SetMode(mode)
}

func (d *DeviceService) StaticRgbSetColor(color string) {
	ctlv2.StaticRgbState["color"] = color
}

func (d *DeviceService) StaticRgbSetBrightness(brightness string) {
	ctlv2.StaticRgbState["brightness"] = brightness
}

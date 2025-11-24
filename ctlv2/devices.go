package ctlv2

import "github.com/yznts/elkctl/elkd"

var (
	// Currently connected ELK devices
	Devices = map[string]Device{}
)

type Device struct {
	// Device configuration
	Name    string
	Addr    string
	Enabled bool
	// Device connectivity
	Elk *elkd.Elk
}

func AddDevice(name, addr string) {
	elk := elkd.New(addr, elkd.Options{})
	elk.Start()
	Devices[name] = Device{
		Name:    name,
		Addr:    addr,
		Enabled: true,
		Elk:     elk,
	}
}

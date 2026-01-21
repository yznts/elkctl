package ctlv2

import (
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/yznts/elkctl/elkd"
)

var (
	// Currently connected ELK devices
	Devices = map[string]*Device{}
)

type Device struct {
	// Device configuration
	Name    string
	Addr    string
	Enabled bool
	Powered bool
	// Device connectivity
	Elk *elkd.Elk
}

// Add/Remove

func AddDevice(name, addr string) {
	// Determine elkd path.
	// First, we will check for macOS app bundle.
	path := "elkd"
	if runtime.GOOS == "darwin" {
		exe, err := os.Executable()
		if err != nil {
			path = "elkd"
		} else {
			exedir := filepath.Dir(exe)
			path = filepath.Clean(filepath.Join(exedir, "..", "Resources", "elkd"))
		}
	}
	elk := elkd.New(addr, elkd.Options{
		Path: path,
	})
	elk.Start()
	Devices[name] = &Device{
		Name:    name,
		Addr:    addr,
		Enabled: true,
		Elk:     elk,
	}
}

func RemoveDevice(name string) {
	delete(Devices, name)
}

// Enable/Disable

func EnableDevice(name string) {
	Devices[name].Enabled = true
}

func DisableDevice(name string) {
	Devices[name].Enabled = false
}

// Power On/Off

func PowerOnDevice(name string) {
	Devices[name].Elk.Exec("power_on:", 5*time.Second)
	Devices[name].Powered = true
}

func PowerOffDevice(name string) {
	Devices[name].Elk.Exec("power_off:", 5*time.Second)
	Devices[name].Powered = false
}

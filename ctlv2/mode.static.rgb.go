package ctlv2

import (
	"log"
	"time"
)

var (
	StaticRgbState = map[string]any{}
)

func StaticRgbMode() {
	// State defaults.
	if _, ok := StaticRgbState["color"]; !ok {
		StaticRgbState["color"] = "15,15,15"
	}
	if _, ok := StaticRgbState["brightness"]; !ok {
		StaticRgbState["brightness"] = "100"
	}
	// Current color and brightness to avoid redundant updates.
	var (
		currentColor      = ""
		currentBrightness = ""
	)
	// Mode mainloop.
	for {
		// Check for mode cancel.
		if err := ModeCtx.Err(); err != nil {
			ModeWaitGroup.Done()
			return
		}
		// If color changed, update it.
		if StaticRgbState["color"] != currentColor {
			currentColor = StaticRgbState["color"].(string)
			for _, device := range Devices {
				// Pass, if device is internally disabled or powered off.
				if !device.Enabled || !device.Powered {
					continue
				}
				_, err := device.Elk.Exec("set_color:"+currentColor, 5*time.Second)
				if err != nil {
					log.Println("Error setting color on device", device.Name, ":", err)
				}
			}
		}
		// If brightness changed, update it.
		if StaticRgbState["brightness"] != currentBrightness {
			currentBrightness = StaticRgbState["brightness"].(string)
			for _, device := range Devices {
				// Pass, if device is internally disabled or powered off.
				if !device.Enabled || !device.Powered {
					continue
				}
				_, err := device.Elk.Exec("set_brightness:"+currentBrightness, 5*time.Second)
				if err != nil {
					log.Println("Error setting brightness on device", device.Name, ":", err)
				}
			}
		}
		// Dummy sleep to prevent busy loop.
		time.Sleep(100 * time.Millisecond)
	}
}

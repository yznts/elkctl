package ctlv2

import (
	"context"
	"sync"
)

var (
	Mode          = ""                // Current mode name.
	ModeCtx       context.Context     // Current mode context.
	ModeCancel    context.CancelFunc  // Current mode cancel function.
	ModeWaitGroup = &sync.WaitGroup{} // ModeWG is used to wait for mode goroutines to finish. Use it for graceful shutdowns.
)

// SetMode stops the current mode (if any, or differs)
// and starts the new mode.
func SetMode(mode string) {
	// No change.
	if Mode == mode {
		return
	}
	// Stop current mode.
	if ModeCancel != nil {
		ModeCancel()
		ModeWaitGroup.Wait()
	}
	// Reset mode state.
	Mode = mode
	// Prepare mode context.
	ModeCtx, ModeCancel = context.WithCancel(context.Background())
	// Start new mode, depending on mode name.
	ModeWaitGroup.Add(1)
	switch Mode {
	case "static:rgb":
		go StaticRgbMode()
	default:
		panic("unknown mode: " + Mode)
	}
}

package ctlv2

// init is responsible for initial actions,
// like state loading,
// starting initial mode goroutine, etc.
func init() {
	// State loading
	// ...

	// Start state saving goroutine
	// ...

	// Start mode according to loaded state
	if Mode != "" {
		SetMode(Mode)
	}

}

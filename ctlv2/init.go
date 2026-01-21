package ctlv2

// init is responsible for initial actions,
// like state loading,
// starting initial mode goroutine, etc.
func init() {
	// State loading
	// ...

	// Start state saving goroutine
	// ...

	// Start default mode.
	// Will be replaced with a saved configuration later.
	SetMode("static:rgb")
}

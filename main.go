package main

import (
	"embed"
	_ "embed"
	"log"
	"os"

	// Keep for init
	_ "github.com/yznts/elkctl/ctlv2"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func NewDevelopmentWindow(app *application.App) *application.WebviewWindow {
	return app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "elkctl",
		Width:  325,
		Height: 450,
	})
}

func NewProductionWindow(app *application.App) *application.WebviewWindow {
	return app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:         "elkctl",
		Width:         325,
		Height:        450,
		AlwaysOnTop:   true,
		Hidden:        true,
		DisableResize: true,
		Frameless:     true,
		// Hide buttons
		MinimiseButtonState: application.ButtonHidden,
		MaximiseButtonState: application.ButtonHidden,
		CloseButtonState:    application.ButtonHidden,
		Mac: application.MacWindow{
			Backdrop:    application.MacBackdropTranslucent,
			TitleBar:    application.MacTitleBarHidden,
			WindowLevel: application.MacWindowLevelPopUpMenu,
		},
		URL: "/",
	})
}

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "elkctl",
		Description: "A control panel for ELK LED devices.",
		Services: []application.Service{
			application.NewService(&GreetService{}),
			application.NewService(&DeviceService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ActivationPolicy: application.ActivationPolicyAccessory,
			ApplicationShouldTerminateAfterLastWindowClosed: false,
		},
	})

	// Create a new window depending on the environment.
	var window *application.WebviewWindow
	if !app.Env.Info().Debug || os.Getenv("ENV") == "prod" {
		window = NewProductionWindow(app)
	} else {
		window = NewDevelopmentWindow(app)
	}

	// Create a system tray menu.
	systray := app.SystemTray.New()
	systray.SetLabel("elkctl")

	// Attach window to the system tray menu.
	systray.AttachWindow(window)

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}

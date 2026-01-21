package main

import (
	"embed"
	_ "embed"
	"log"
	"os"

	// Keep for init
	"github.com/yznts/elkctl/ctlv2"
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
			application.NewService(&StatusBarService{}),
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
	systray.SetLabel("💡")

	// Attach window to the system tray menu.
	if !app.Env.Info().Debug || os.Getenv("ENV") == "prod" {
		systray.AttachWindow(window)
	}

	// Add test devices
	if !app.Env.Info().Debug || os.Getenv("ENV") == "prod" {
		// None for production

		// Temporary solution
		ctlv2.AddDevice("Lamp 1", "A7FB96B6-5ED1-5A8D-BAB7-62271EA0B9C7")
		ctlv2.PowerOnDevice("Lamp 1")
		ctlv2.EnableDevice("Lamp 1")
		ctlv2.AddDevice("Lamp 2", "DDDB7C05-8FC0-0F16-3017-1AD8F101FE99")
		ctlv2.PowerOnDevice("Lamp 2")
		ctlv2.EnableDevice("Lamp 2")
	} else {
		ctlv2.AddDevice("Lamp 1", "A7FB96B6-5ED1-5A8D-BAB7-62271EA0B9C7")
		ctlv2.PowerOnDevice("Lamp 1")
		ctlv2.EnableDevice("Lamp 1")
		ctlv2.AddDevice("Lamp 2", "DDDB7C05-8FC0-0F16-3017-1AD8F101FE99")
		ctlv2.PowerOnDevice("Lamp 2")
		ctlv2.EnableDevice("Lamp 2")
	}

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}

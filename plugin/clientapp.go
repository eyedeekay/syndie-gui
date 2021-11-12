package plugin

import (
	"os"
	"os/user"

	"github.com/kpetku/syndie-gui/ui"
	"i2pgit.org/idk/clientapp"
)

// SyndieClientApp is an example of a ClientApp. In java, use this declaration:
// class SyndieClientApp extends go.clientapp.ClientApp.SyndieClientApp implements ClientApp {
// this needs to be done in the code generation phase for apps that import this code
type SyndieClientApp struct {
	DisplayName string
	Name        string
	State       clientapp.ClientAppState
	client      *ui.UI
}

// GetDisplayName() implements ClientApp in Java and in Go
func (e *SyndieClientApp) GetDisplayName() string {
	return e.DisplayName
}

// GetName() implements ClientApp in Java and in Go
func (e *SyndieClientApp) GetName() string {
	return e.Name
}

// GetState() implements ClientApp in Java and in Go
func (e *SyndieClientApp) GetState() clientapp.ClientAppState {
	return e.State
}

func (e *SyndieClientApp) GetClientState() int {
	return int(e.State)
}

// Startup() implements ClientApp in Java and in Go
func (e *SyndieClientApp) StartupClient() {
	e.State = clientapp.STARTING
	e.client = ui.NewUI()
	usr, err := user.Current()
	if err != nil {
		panic("Error obtaining current user")
	}
	pluginpath := usr.HomeDir + "/.i2p/plugins/syndie"
	if usr.Username == "i2psvc" {
		pluginpath = "/var/lib/i2p/i2p-config/plugins/syndie"
	}
	os.MkdirAll(pluginpath, 0755)
	e.client.Start(pluginpath + "/.syndie")
	e.State = clientapp.RUNNING
}

// Shutdown() implements ClientApp in Java and in Go. In go, Shutdown never takes args.
func (e *SyndieClientApp) Shutdown() {
	e.State = clientapp.STOPPING

	e.State = clientapp.STOPPED
}

var ClientAppImpl interface{} = &SyndieClientApp{}

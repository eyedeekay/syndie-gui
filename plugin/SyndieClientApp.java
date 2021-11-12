
package net.i2p.SyndieClientApp;
import net.i2p.app.ClientApp;
import net.i2p.app.ClientAppState;
public class SyndieClientApp implements ClientApp {
	go.clientapp.Clientapp.ExampleClientApp _clientApp;
	@Override
	public String getDisplayName() {
		return _clientApp.getDisplayName();
	}
	@Override
	public String getName() {
		return _clientApp.getName();
	}
	private ClientAppState convertClientState(long cas){
		switch((int)cas){
			case 0:
				return ClientAppState.UNINITIALIZED;
			case 1:
				return ClientAppState.INITIALIZED;
			case 2:
				return ClientAppState.STARTING;
			case 3:
				return ClientAppState.START_FAILED;
			case 4:
				return ClientAppState.RUNNING;
			case 5:
				return ClientAppState.STOPPING;
			case 6:
				return ClientAppState.STOPPED;
			case 7:
				return ClientAppState.CRASHED;
			case 8:
				return ClientAppState.FORKED;
			default:
				return ClientAppState.UNINITIALIZED;
		}
	}	@Override
	public ClientAppState getState() {
		return convertClientState(_clientApp.GetClientState());
	}
	@Override
	public void startup() {
		_clientApp.StartupClient();
	}
	@Override
	public void shutdown(String[] args) {
		_clientApp.Shutdown();
	}
}

import { DeviceService } from "./bindings/github.com/yznts/elkctl";
import { SwitchPowerState } from "./bindings/github.com/yznts/elkctl/deviceservice";

// Devices list module
function device() {
  return {
    state: {
      devices: [],
    },
    bindings: {
      name: "",
      addr: "",
    },

    // Init module
    init() {
      // Peroidically refresh devices list
      setInterval(this.refreshDevices.bind(this), 500);
    },

    addDevice() {
      // Get form values
      let form = document.querySelector("#add-device-form");
      let name = form.querySelector("input[name='name']").value;
      let addr = form.querySelector("input[name='addr']").value;
      // Call the backend service
      DeviceService.AddDevice(name, addr);
    },

    removeDevice(name) {
      DeviceService.RemoveDevice(name);
    },

    switchEnableState(name) {
      DeviceService.SwitchEnableState(name);
    },

    switchPowerState(name) {
      DeviceService.SwitchPowerState(name);
    },

    async refreshDevices() {
      const devices = await DeviceService.GetDevices();
      this.state.devices = devices.sort((a, b) => a.Name.localeCompare(b.Name));
    },
  };
}

// Export
window.device = device;

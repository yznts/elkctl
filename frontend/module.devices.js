import { DeviceService } from "./bindings/github.com/yznts/elkctl";

// Devices module
function devices() {
  const devices = {
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
      setInterval(this.refreshDevices.bind(this), 2000);
    },

    // Handle device add
    addDevice() {
      // Get form values
      let form = document.querySelector("#add-device-form");
      let name = form.querySelector("input[name='name']").value;
      let addr = form.querySelector("input[name='addr']").value;
      // Call the backend service
      DeviceService.AddDevice(name, addr);
    },

    async refreshDevices() {
      this.state.devices = await DeviceService.GetDevices();
    },
  };

  return devices;
}

// Export
window.devices = devices;

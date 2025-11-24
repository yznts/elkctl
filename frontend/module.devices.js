import { DeviceService } from "./bindings/github.com/yznts/elkctl";

// Devices module
const devices = {
  state: {
    devices: [],
  },
  bindings: {
    name: "",
    addr: "",
  },

  // Handle device add
  addDevice: () => {
    // Get form values
    let form = document.querySelector("#add-device-form");
    let name = form.querySelector("input[name='name']").value;
    let addr = form.querySelector("input[name='addr']").value;
    // Call the backend service
    DeviceService.AddDevice({ name: name, addr: addr });
  },
};

// Export
window.devices = devices;

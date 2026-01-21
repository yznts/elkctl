import { DeviceService } from "./bindings/github.com/yznts/elkctl";

// Mode select and configuration module
function mode() {
  return {
    state: {},
    bindings: {
      mode: "",
      staticRgbColor: "",
      staticRgbBrightness: 0,
    },

    init() {},

    // Handle mode change
    changeMode() {
      DeviceService.SetMode(this.bindings.mode);
    },

    staticRgbSetColor() {
      // Convert hex to rgb
      const hex = this.bindings.staticRgbColor.replace("#", "");
      const r = parseInt(hex.substring(0, 2), 16);
      const g = parseInt(hex.substring(2, 4), 16);
      const b = parseInt(hex.substring(4, 6), 16);
      // Send to the backend
      DeviceService.StaticRgbSetColor(`${r},${g},${b}`);
    },

    staticRgbSetBrightness() {
      DeviceService.StaticRgbSetBrightness(this.bindings.staticRgbBrightness);
    },
  };
}

// Export
window.mode = mode;

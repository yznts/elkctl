import { StatusBarService } from "./bindings/github.com/yznts/elkctl";

function statusbar() {
  return {
    init() {},

    quit() {
      StatusBarService.Quit();
    },
  };
}

window.statusbar = statusbar;

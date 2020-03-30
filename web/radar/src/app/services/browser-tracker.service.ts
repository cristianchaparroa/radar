import { Injectable } from "@angular/core";

import { Plugins } from "@capacitor/core";
const { Device } = Plugins;

@Injectable({
  providedIn: "root"
})
export class BrowserTrackerService {
  private WebPlatform = "web";

  constructor() {}

  async isWebPlatform() {
    let info = await Device.getInfo();
    return info.platform == this.WebPlatform;
  }
}

import { Injectable } from "@angular/core";
import { BrowserTrackerService } from "./browser-tracker.service";
import {
  BackgroundGeolocation,
  BackgroundGeolocationConfig,
  BackgroundGeolocationResponse,
  BackgroundGeolocationEvents
} from "@ionic-native/background-geolocation/ngx";

import { LocationApiService } from "../api/location-api.service";
import { LocalProfileService } from "./local-profile.service";
import { LocationModel} from "../models/location.model";

@Injectable({
  providedIn: "root"
})
export class TrackerService {
  constructor(
    private local: LocalProfileService,
    private webTracker: BrowserTrackerService,
    private backgroundGeolocation: BackgroundGeolocation,
    private locationApi: LocationApiService
  ) {}

  async start() {
    /*
    let isWeb = await this.webTracker.isWebPlatform();
    if (isWeb) {
      console.log(
        "Is web application. The web tracking is not implement yet!."
      );
      return;
    }*/

    this.startMobileTracking();
  }

  async startMobileTracking() {

    console.log("--> StartMobileTracking");
    let config = this.getConfiguration();

    if (config == null) {
      return;
    }

    this.setupBackground(config);

    // start recording location
    console.log("--> this.backgroundGeolocation.start()");
    this.backgroundGeolocation.start();
  }

  getConfiguration() {

    try {
      const config: BackgroundGeolocationConfig = {
        interval: 1000,
        desiredAccuracy: 10,
        stationaryRadius: 20,
        distanceFilter: 30,
        debug: true, //  enable this hear sounds for background-geolocation life-cycle.
        stopOnTerminate: false // enable this to clear background location settings when the app terminates
      };
      return config;
    } catch(e) {
        console.log("--> Error trying to getConfiguration: ", e)
    }

    return null;
  }

  async setupBackground(config) {

    this.backgroundGeolocation.configure(config).then(() => {
      this.backgroundGeolocation
        .on(BackgroundGeolocationEvents.location)
        .subscribe((location: BackgroundGeolocationResponse) => {
          console.log("--> Location");
          console.log(location);
          this.sendLocation(location);
          this.backgroundGeolocation.finish();
        });
    });
  }

  async sendLocation(position) {

    let current = await this.local.getProfile();
    let profile = current.profile;
    let location = new LocationModel(position.latitude, position.longitude);
    location.profileID = profile.id;
    await this.locationApi.register(location);
    console.log("--> sendLocation", position.latitude, position.longitude);
  }

  stop() {
    // If you wish to turn OFF background-tracking, call the #stop method.
    this.backgroundGeolocation.stop();
  }
}

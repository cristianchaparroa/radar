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

  startMobileTracking() {

    console.log("--> StartMobileTracking");
    let config = this.getConfiguration();

    this.backgroundGeolocation.configure(config).then(() => {
      this.backgroundGeolocation
        .on(BackgroundGeolocationEvents.location)
        .subscribe((location: BackgroundGeolocationResponse) => {
          console.log("--> Location");
          console.log(location);
          this.sendLocation(location);

          // IMPORTANT:  You must execute the finish method here to inform the native plugin that you're finished,
          // and the background-task may be completed.  You must do this regardless if your operations are successful or not.
          // IF YOU DON'T, ios will CRASH YOUR APP for spending too much time in the background.
          this.backgroundGeolocation.finish(); // FOR IOS ONLY
        });
    });


    // start recording location
    console.log("--> this.backgroundGeolocation.start()");
    this.backgroundGeolocation.start();
  }

  getConfiguration() {
    const config: BackgroundGeolocationConfig = {
      interval: 1000,
      desiredAccuracy: 10,
      stationaryRadius: 20,
      distanceFilter: 30,
      debug: true, //  enable this hear sounds for background-geolocation life-cycle.
      stopOnTerminate: false // enable this to clear background location settings when the app terminates
    };
    return config;
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

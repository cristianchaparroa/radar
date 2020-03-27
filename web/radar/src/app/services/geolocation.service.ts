import { Injectable } from "@angular/core";
import { Geolocation } from "@ionic-native/geolocation/ngx";
import { LocationModel } from "../models/location.model";

@Injectable({
  providedIn: "root"
})
export class GeolocationService {
  constructor(private geolocation: Geolocation) {}

  async getLocation() {
    let self = this;
    let promise = await this.geolocation.getCurrentPosition();
    let location = new LocationModel(
      promise.coords.latitude,
      promise.coords.longitude
    );
    return location;
  }
}

import { Injectable } from "@angular/core";
import { UniqueDeviceID } from "@ionic-native/unique-device-id/ngx";
import { v4 as uuid } from "uuid";

import { ProfileModel } from "../models/profile.model";
import { LocalProfileService } from "./local-profile.service";
import { ProfileApiService } from "../api/profile-api.service";

@Injectable({
  providedIn: "root"
})
export class ProfileService {
  constructor(
    private uniqueDeviceID: UniqueDeviceID,
    private local: LocalProfileService,
    private profileAPI: ProfileApiService
  ) {}

  async create() {
    try {
      let profile = await this.createProfileModel();
      let profileResponse = await this.profileAPI.create(profile);
      this.local.create(profile);
    } catch (e) {
      console.log(e);
    }
  }

  async createProfileModel() {
    let id = uuid();
    let deviceID = await this.getDeviceID();
    return new ProfileModel(id, deviceID);
  }

  async getDeviceID() {
    try {
      let deviceID = await this.uniqueDeviceID.get();
      return deviceID;
    } catch (e) {
      return "";
    }
  }

  async load() {
    let isProfile = await this.exist();
    if (!isProfile) {
      this.create();
    }
  }

  async exist() {
    let isProfile = await this.local.exist();
    return isProfile;
  }
}

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

  /**
   * Init verifies if is saved a profile in the local storage if not it will
   * create a new.
   */
  async init() {
    let isProfile = await this.exist();
    if (!isProfile) {
      this.create();
    }
  }

  /**
  * It verifies if already exist a profile in the locale storage.
  */
  async exist() {
    let isProfile = await this.local.exist();
    return isProfile;
  }
  /**
   * It creates a profile in the backend server and in the local storage.
   */
  async create() {
    try {
      let profile = await this.createProfileModel();
      let profileResponse = await this.profileAPI.create(profile);
      this.local.create(profile);
    } catch (e) {
      console.log(e);
    }
  }

  /**
   * It creates a model of profile.
   */
  async createProfileModel() {
    let id = uuid();
    let deviceID = await this.getDeviceID();
    return new ProfileModel(id, deviceID);
  }

  /**
   * It try to retrieves the devices ID. If is a browser it returns empty in
   * other case it will return the deviceID.
   */
  async getDeviceID() {
    try {
      let deviceID = await this.uniqueDeviceID.get();
      return deviceID;
    } catch (e) {
      return "";
    }
  }
}

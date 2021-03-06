import { Injectable } from "@angular/core";
import { Plugins } from "@capacitor/core";
import { CurrentProfileModel } from "../models/current-profile.model";

const { Storage } = Plugins;

/**
 * LocalProfileService is in charge to save all data related to profile
 */
@Injectable({
  providedIn: "root"
})
export class LocalProfileService {
  constructor() {}

  async create(profile: CurrentProfileModel) {
    await Storage.set({
      key: "profile",
      value: JSON.stringify(profile)
    });
  }

  /**
   * It retrieves the profile from local storage
   */
  async getProfile() {
    const ret = await Storage.get({ key: "profile" });
    const profile = JSON.parse(ret.value);
    return profile;
  }

  /**
   * It verfies if exist a profile saved in the local storage.
   */
  async exist() {
    try {
      let currentProfile = await this.getProfile();
      return !this.isNotValidCurrentProfile(currentProfile);
    } catch (e) {
      console.log(e);
      return false;
    }
    return false;
  }

  isNotValidCurrentProfile(currentProfile) {
    return (
      currentProfile == null ||
      currentProfile.profile == null ||
      currentProfile.profile.id == null
    );
  }
}

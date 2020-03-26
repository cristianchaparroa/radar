import { Injectable } from "@angular/core";
import { Plugins } from "@capacitor/core";
import { ProfileModel } from "../models/profile.model";

const { Storage } = Plugins;

@Injectable({
  providedIn: "root"
})
export class LocalProfileService {
  constructor() {}

  async create(profile: ProfileModel) {
    await Storage.set({
      key: "profile",
      value: JSON.stringify(profile)
    });
  }

  async getProfile() {
    const ret = await Storage.get({ key: "profile" });
    const profile = JSON.parse(ret.value);
    return profile;
  }

  async exist() {
    try {
      let profile = await this.getProfile();
      if (profile == null || profile.id == null) {
        return false;
      }
      return true;
    } catch (e) {
      console.log(e);
      return false;
    }
    return false;
  }
}

import { Injectable } from "@angular/core";
import { Plugins } from "@capacitor/core";
import { StatusModel } from "../models/status.model";

const { Storage } = Plugins;

@Injectable({
  providedIn: "root"
})
export class LocalStatusService {
  constructor() {}

  async updateStatus(status: StatusModel) {
    const ret = await Storage.get({ key: "profile" });
    const profile = JSON.parse(ret.value);
    profile.status = status;

    await Storage.set({
      key: "profile",
      value: JSON.stringify(profile)
    });

    return profile;
  }
}

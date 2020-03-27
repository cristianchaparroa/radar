import { ProfileModel } from "../../models/profile.model";
import { LocationModel } from "../../models/location.model";

export class CreateProfileRequest {
  profile: ProfileModel;
  location: LocationModel;

  constructor(profile: ProfileModel, location: LocationModel) {
    this.profile = profile;
    this.location = location;
    this.location.profileId = this.profile.id;
  }

  JSON() {
    return JSON.stringify(this);
  }
}

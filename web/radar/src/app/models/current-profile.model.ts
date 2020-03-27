import { ProfileModel } from "./profile.model";
import { StatusModel } from "./status.model";
import { LocationModel } from "./location.model";


/**
 * CurrentProfileModel contains the information in the last status update.
 */
export class CurrentProfileModel {
  /**
   * This is the profile active.
   */
  profile: ProfileModel;

  /**
   * This is the current status;
   */
  status: StatusModel;

  /**
   * This is the location when is changed the status;
   */
  location: LocationModel;

  constructor(profile: ProfileModel, status: StatusModel, location:LocationModel) {
    this.profile = profile;
    this.status = status;
  }
}

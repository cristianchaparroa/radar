import { StatusModel } from "../../models/status.model";
import { LocationModel } from "../../models/location.model";

export class CreateStatusRequest {
  status: StatusModel;
  location: LocationModel;

  constructor(status: StatusModel, location: LocationModel) {
    this.status = status;
    this.location = location;
  }

  JSON() {
    return JSON.stringify(this);
  }
}

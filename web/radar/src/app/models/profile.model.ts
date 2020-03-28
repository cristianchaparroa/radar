export class ProfileModel {
  id: string;
  deviceID: string;
  createdAt: any;

  constructor(id: string, deviceID: string) {
    this.id = id;
    this.deviceID = deviceID;
  }

  JSON() {
    return JSON.stringify(this);
  }
}

export function NewProfilefromJSON(json: any) {
  let id = json.profile.id;
  let deviceID = json.profile.deviceID;
  let createdAt = json.profile.createdAt;
  let model = new ProfileModel(id, deviceID);
  model.createdAt = createdAt;
  return model;
}

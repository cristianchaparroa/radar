export class ProfileModel {
  id: string;
  deviceID: string;
  createdAt: any;

  constructor(id: string, deviceID: string) {
    this.id = id;
    this.deviceID = deviceID;
  }

  fromJSON(json: any) {
    let id = json.id;
    let deviceID = json.device_id;
    let createdAt = json.created_at;
    let model = new ProfileModel(id, deviceID);
    model.createdAt = createdAt;
    return model;
  }

  JSON() {
    return JSON.stringify(this);
  }
}

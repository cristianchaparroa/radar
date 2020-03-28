export class StatusModel {
  id: string;
  profileID: string;
  name: string;
  createdAt: any;

  constructor(id?: string, profileID?: string, name?: string, createdAt?: any) {
    this.id = id;
    this.profileID = profileID;
    this.name = name;
    this.createdAt = createdAt;
  }

  JSON() {
    return JSON.stringify(this);
  }
}

export function NewStatusModelFromJSON(json: any) {
  let id = json.status.id;
  let profileId = json.status.profileID;
  let name = json.status.name;
  let createdAt = json.status.createdAt;
  return new StatusModel(id, profileId, name, createdAt);
}

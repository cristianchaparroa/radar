
export class ProfileModel {

  userID:number;
  deviceID:string;


  constructor(userID:string, deviceID:string) {
    this.userID   = userID
    this.deviceID = deviceID;
  }
}

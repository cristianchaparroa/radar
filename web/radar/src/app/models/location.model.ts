
export class LocationModel {

  userId:number;
  latitude:number;
  longitude:number;

  constructor(lat:number, lng:number) {
    this.latitude = lat;
    this.longitude = lng;
  }
}

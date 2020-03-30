import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { LOCATIONS_ENDPOINT } from "../../environments/environment";
import { LocationModel } from "../models/location.model";



@Injectable({
  providedIn: 'root'
})
export class LocationApiService {

  httpOptions = {
    headers: new HttpHeaders({
      "Content-Type": "application/json"
    })
  };

  constructor(protected httpClient: HttpClient) {}

  /**
   * It performs a post request to register a new location.
   */
  async post(location: LocationModel) {
    return this.httpClient
      .post<LocationModel>(LOCATIONS_ENDPOINT, location.JSON(), this.httpOptions)
      .toPromise();
  }

  async register(location:LocationModel) {
    let response = await this.post(location);
    console.log(response);
  }
}

import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders } from "@angular/common/http";

import { ProfileModel } from "../models/profile.model";
import { PROFILES_ENDPOINT } from "../../environments/environment";

@Injectable({
  providedIn: "root"
})
export class ProfileApiService {
  httpOptions = {
    headers: new HttpHeaders({
      "Content-Type": "application/json"
    })
  };

  constructor(protected httpClient: HttpClient) {}

  async post(profile: ProfileModel) {
    return this.httpClient
      .post<ProfileModel>(PROFILES_ENDPOINT, profile.JSON(), this.httpOptions)
      .toPromise();
  }

  async create(profile: ProfileModel) {
    let response = await this.post(profile);
    console.log(response);
    let p = new ProfileModel("","");
    return p.fromJSON(response);
  }
}

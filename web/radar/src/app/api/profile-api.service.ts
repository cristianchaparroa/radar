import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders } from "@angular/common/http";

import { ProfileModel, NewProfilefromJSON } from "../models/profile.model";
import { StatusModel, NewStatusModelFromJSON } from "../models/status.model";
import { CurrentProfileModel } from "../models/current-profile.model";
import {CreateProfileRequest} from "./request/create-profile.request";

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

  /**
   * It performs a post request to create a new profile.
   */
  async post(profile: CreateProfileRequest) {
    return this.httpClient
      .post<ProfileModel>(PROFILES_ENDPOINT, profile.JSON(), this.httpOptions)
      .toPromise();
  }
  /**
   * It creates a profile and parse the response properly.
   */
  async create(profile: CreateProfileRequest) {
    let response = await this.post(profile);
    let profileResponse = NewProfilefromJSON(response);
    let statusResponse = NewStatusModelFromJSON(response);
    return new CurrentProfileModel(profileResponse, statusResponse, profile.location);
  }
}

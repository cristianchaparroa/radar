import { Injectable } from "@angular/core";
import { CreateStatusRequest } from "./request/create-status.request";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { NewStatusModelFromJSON } from "../models/status.model";

import { STATUSES_ENDPOINT } from "../../environments/environment";

@Injectable({
  providedIn: "root"
})
export class StatusesApiService {
  httpOptions = {
    headers: new HttpHeaders({
      "Content-Type": "application/json"
    })
  };

  constructor(protected httpClient: HttpClient) {}
  async post(request: CreateStatusRequest) {
    return this.httpClient
      .post<CreateStatusRequest>(
        STATUSES_ENDPOINT,
        request.JSON(),
        this.httpOptions
      )
      .toPromise();
  }

  async create(request: CreateStatusRequest) {
    let response = await this.post(request);
    let statusResponse = NewStatusModelFromJSON(response);
    return statusResponse;
  }
}

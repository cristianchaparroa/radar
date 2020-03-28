import { Injectable } from "@angular/core";
import { GeolocationService } from "./geolocation.service";
import { LocalStatusService } from "./local-status.service";

import { StatusModel } from "../models/status.model";
import { StatusesApiService } from "../api/statuses-api.service";
import { CreateStatusRequest } from "../api/request/create-status.request";

@Injectable({
  providedIn: "root"
})
export class StatusesService {
  private statuses = [
    {
      index: 0,
      label: "Positivo",
      value: "POSITIVE",
      isChecked: false,
      icon: "add-circle",
      color: "danger"
    },
    {
      index: 1,
      label: "Negativo",
      value: "NEGATIVE",
      isChecked: true,
      icon: "thumbs-up",
      color: "warning"
    },
    {
      index: 2,
      label: "Recuperado",
      value: "RECOVERED",
      isChecked: false,
      icon: "checkmark-circle",
      color: "success"
    }
  ];

  constructor(
    private geolocationService: GeolocationService,
    private local: LocalStatusService,
    private statusesApiService: StatusesApiService
  ) {}

  getStatuses() {
    return this.statuses;
  }

  getIndexSelectedByName(name) {
    for (let s of this.statuses) {
      if (s.value == name) {
        return s.index;
      }
    }
    return -1;
  }

  async create(profileId, statusName) {
    let location = await this.geolocationService.getLocation();
    let status = this.createStatus(profileId, statusName);
    let request = new CreateStatusRequest(status, location);
    let statusResponse = await this.statusesApiService.create(request);
    let currentProfile = await this.local.updateStatus(statusResponse);
    return currentProfile;
  }

  createStatus(profileId, statusName) {
    let status = new StatusModel();
    status.profileID = profileId;
    status.name = statusName;
    return status;
  }
}

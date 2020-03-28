import { Component, OnInit } from "@angular/core";
import { ProfileService } from "../../services/profile.service";
import { StatusesService } from "../../services/statuses.service";
import { CurrentProfileModel } from "../../models/current-profile.model";

@Component({
  selector: "app-profile",
  templateUrl: "./profile.page.html",
  styleUrls: ["./profile.page.scss"]
})
export class ProfilePage implements OnInit {
  /**
   * currentProfile contains the current profile
   */
  private currentProfile: CurrentProfileModel;

  /**
   * This is the index of the status options selected
   */
  private selectedStatusIndex = 1;

  /**
   * This is the status options showed.
   */
  private statuses = [];

  constructor(
    private profileService: ProfileService,
    private statusesService: StatusesService
  ) {}

  async ngOnInit() {
    // Fill the status options
    this.statuses = this.statusesService.getStatuses();

    // Load the current profile
    this.currentProfile = await this.profileService.init();

    // retrieve the current satus name
    let statusName = this.currentProfile.status.name;

    // set the option selected
    this.selectedStatusIndex = this.statusesService.getIndexSelectedByName(
      statusName
    );
  }

  /**
   * function triggered when is changed the status
   */
  async changeStatus(event, statusName) {
    let isChecked = event.detail.checked;
    if (isChecked) {
      console.log(this.currentProfile);
      let profileId = this.currentProfile.profile.id;
      this.currentProfile = await this.statusesService.create(
        profileId,
        statusName
      );
    }
  }
}

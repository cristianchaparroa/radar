import { Component, OnInit } from "@angular/core";
import { ToastController } from "@ionic/angular";
import { BackgroundMode } from '@ionic-native/background-mode/ngx';

import { ProfileService } from "../../services/profile.service";
import { StatusesService } from "../../services/statuses.service";
import { CurrentProfileModel } from "../../models/current-profile.model";
import { TrackerService} from "../../services/tracker.service";

@Component({
  selector: "app-profile",
  templateUrl: "./profile.page.html",
  styleUrls: ["./profile.page.scss"]
})
export class ProfilePage implements OnInit {

  /**
   * Identifies if the tracking is enable;
  */
  private isTracking:boolean;
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
    public toastController: ToastController,
    private profileService: ProfileService,
    private statusesService: StatusesService,
    private backgroundMode: BackgroundMode,
    private tracker:TrackerService
  ) {}

  async ngOnInit() {

    this.backgroundMode.enable();
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

       this.showUpdateMessage();
    }
  }


  changeTracking(event) {
    let isAllowedTracking = event.detail.checked;
    if (isAllowedTracking) {
      console.log("Tracking enabled");
      this.initTracking();
    } else {
      console.log("Tracking disabled");
      this.stopTracking();
    }
  }

  async showUpdateMessage() {
    const toast = await this.toastController.create({
      color: "light",
      message: "El estado ha sido actualizado",
      duration: 2000
    });
    toast.present();
  }

  initTracking() {
      this.tracker.start();
  }

  stopTracking() {
    this.tracker.stop();
  }

}

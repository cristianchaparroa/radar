import { Component, OnInit } from "@angular/core";
import { ProfileService } from "../../services/profile.service";

@Component({
  selector: "app-profile",
  templateUrl: "./profile.page.html",
  styleUrls: ["./profile.page.scss"]
})
export class ProfilePage implements OnInit {
  private statuses = [
    {
      label: "Positivo",
      value: "positive",
      isChecked: false,
      icon: "add-circle",
      color: "danger"
    },
    {
      label: "Negativo",
      value: "negative",
      isChecked: true,
      icon: "thumbs-up",
      color: "warning"
    },
    {
      label: "Recuperado",
      value: "recovered",
      isChecked: false,
      icon: "checkmark-circle",
      color: "success"
    }
  ];

  constructor(private profileService: ProfileService) {}

  async ngOnInit() {
    this.profileService.init();
  }

  isProfile() {}

  changeStatus(status) {
    for (let s of this.statuses) {
      if (s.value !== status) {
        s.isChecked = false;
      }
    }
  }
}

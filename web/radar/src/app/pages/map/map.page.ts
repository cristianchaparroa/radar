import { Component, OnInit } from '@angular/core';
import {LocationService} from './../../services/location.service';


@Component({
  selector: 'app-map',
  templateUrl: './map.page.html',
  styleUrls: ['./map.page.scss'],
})
export class MapPage implements OnInit {

  constructor(private locationService:LocationService) {
      this.loadLocation();
  }

  ngOnInit() {}

  async loadLocation() {
    let location = await this.locationService.getLocation();
    console.log(location);
  }



}

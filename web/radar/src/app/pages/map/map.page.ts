import { Component, OnInit , ElementRef, ViewChild} from '@angular/core';
import {GeolocationService} from './../../services/geolocation.service';
import { Map, Marker } from "mapbox-gl";
import * as mapboxgl from "mapbox-gl";


@Component({
  selector: 'app-map',
  templateUrl: './map.page.html',
  styleUrls: ['./map.page.scss'],
})
export class MapPage implements OnInit {

  @ViewChild('map', {static: false}) mapElement: ElementRef;

  map: mapboxgl.Map;

  constructor(private geolocationService:GeolocationService) {
      Object.getOwnPropertyDescriptor(mapboxgl, "accessToken").set("pk.eyJ1IjoiY3Jpc3RpYW5jaGFwYXJyb2EiLCJhIjoiY2s4NmxpYW16MGZ3MzNnbXJ2M2o4YmkwNCJ9.Mt9Redxy3_5QsHvoTpfGtQ");
  }


  async  loadMap() {

      let location = await this.geolocationService.getLocation();

      this.map = new Map({
          container: 'map', // container id
          style: 'mapbox://styles/mapbox/dark-v10',
          center: [location.longitude, location.latitude],
          zoom: 13 // starting zoom
      });
      this.map.resize();
      this.map.addControl(new mapboxgl.NavigationControl());

      var marker = new Marker();
      marker.setLngLat([location.longitude, location.latitude]);
      marker.addTo(this.map);

      var dos = new Marker();
      dos.setLngLat([4.725111, -74.268288]);
      dos.addTo(this.map);


      var tres = new Marker();
      tres.setLngLat([4.725156, -74.269401]);
      tres.addTo(this.map);


  }

  ngOnInit() {
      this.loadMap();
  }




}

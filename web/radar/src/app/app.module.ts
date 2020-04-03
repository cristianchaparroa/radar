import { NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";
import { RouteReuseStrategy } from "@angular/router";

import { IonicModule, IonicRouteStrategy } from "@ionic/angular";
import { SplashScreen } from "@ionic-native/splash-screen/ngx";
import { StatusBar } from "@ionic-native/status-bar/ngx";

import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";
import { Geolocation } from "@ionic-native/geolocation/ngx";
import { HttpClientModule } from "@angular/common/http";
import { UniqueDeviceID } from "@ionic-native/unique-device-id/ngx";
import { BackgroundGeolocation } from "@ionic-native/background-geolocation/ngx";
import { BackgroundMode } from '@ionic-native/background-mode/ngx';
import {TrackerService} from './services/tracker.service';

@NgModule({
  declarations: [AppComponent],
  entryComponents: [],
  imports: [
    BrowserModule,
    IonicModule.forRoot(),
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [
    StatusBar,
    SplashScreen,
    { provide: RouteReuseStrategy, useClass: IonicRouteStrategy },
    Geolocation,
    UniqueDeviceID,
    BackgroundMode,
    BackgroundGeolocation,
    TrackerService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}

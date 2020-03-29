import { Injectable } from '@angular/core';
import { BrowserTrackerService} from './browser-tracker.service';

@Injectable({
  providedIn: 'root'
})
export class TrackerService {

  constructor(private webTracker:BrowserTrackerService) {
  }
}

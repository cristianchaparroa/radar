import { TestBed } from '@angular/core/testing';

import { BrowserTrackerService } from './browser-tracker.service';

describe('BrowserTrackerService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: BrowserTrackerService = TestBed.get(BrowserTrackerService);
    expect(service).toBeTruthy();
  });
});

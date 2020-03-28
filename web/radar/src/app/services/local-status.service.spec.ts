import { TestBed } from '@angular/core/testing';

import { LocalStatusService } from './local-status.service';

describe('LocalStatusService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: LocalStatusService = TestBed.get(LocalStatusService);
    expect(service).toBeTruthy();
  });
});

import { TestBed } from '@angular/core/testing';

import { LocalProfileService } from './local-profile.service';

describe('LocalProfileService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: LocalProfileService = TestBed.get(LocalProfileService);
    expect(service).toBeTruthy();
  });
});

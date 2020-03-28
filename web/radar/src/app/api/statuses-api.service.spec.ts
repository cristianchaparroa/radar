import { TestBed } from '@angular/core/testing';

import { StatusesApiService } from './statuses-api.service';

describe('StatusesApiService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: StatusesApiService = TestBed.get(StatusesApiService);
    expect(service).toBeTruthy();
  });
});

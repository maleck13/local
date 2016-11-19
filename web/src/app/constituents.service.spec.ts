/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { ConstituentsService } from './constituents.service';

describe('Service: Constituents', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ConstituentsService]
    });
  });

  it('should ...', inject([ConstituentsService], (service: ConstituentsService) => {
    expect(service).toBeTruthy();
  }));
});

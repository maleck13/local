/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { CommunicationsService } from './communications.service';

describe('Service: Communications', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CommunicationsService]
    });
  });

  it('should ...', inject([CommunicationsService], (service: CommunicationsService) => {
    expect(service).toBeTruthy();
  }));
});

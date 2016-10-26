/* tslint:disable:no-unused-variable */

import { addProviders, async, inject } from '@angular/core/testing';
import { CouncillorsService } from './councillors.service';

describe('Service: Councillors', () => {
  beforeEach(() => {
    addProviders([CouncillorsService]);
  });

  it('should ...',
    inject([CouncillorsService],
      (service: CouncillorsService) => {
        expect(service).toBeTruthy();
      }));
});


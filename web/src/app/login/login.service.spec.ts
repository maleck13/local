/* tslint:disable:no-unused-variable */

import { addProviders, async, inject } from '@angular/core/testing';
import { LoginService } from './login.service';

describe('Service: Login', () => {
  beforeEach(() => {
    addProviders([LoginService]);
  });

  it('should ...',
    inject([LoginService],
      (service: LoginService) => {
        expect(service).toBeTruthy();
      }));
});

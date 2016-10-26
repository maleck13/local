/* tslint:disable:no-unused-variable */

import { By }           from '@angular/platform-browser';
import { DebugElement } from '@angular/core';
import { addProviders, async, inject } from '@angular/core/testing';
import { CouncillorsComponent } from './councillors.component';

describe('Component: Councillors', () => {
  it('should create an instance', () => {
    let component = new CouncillorsComponent();
    expect(component).toBeTruthy();
  });
});

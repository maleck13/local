import { bootstrap } from '@angular/platform-browser-dynamic';
import { enableProdMode, provide } from '@angular/core';
import { AppComponent, environment } from './app/';
import { appRouterProviders } from './app/app.routes';
import { LocationStrategy,HashLocationStrategy } from '@angular/common';
import {HTTP_PROVIDERS} from '@angular/http'


if (environment.production) {
  enableProdMode();
}

bootstrap(AppComponent, [
  HTTP_PROVIDERS,
  { provide: LocationStrategy, useClass: HashLocationStrategy },
  appRouterProviders
]).catch(err => console.error(err));




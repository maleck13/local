import { Component, OnInit, ViewChild, ElementRef, Renderer, NgZone} from '@angular/core';
import { ROUTER_DIRECTIVES, Router } from '@angular/router';
import {HomeComponent} from "./home/home.component";
import {SignupComponent} from "./signup/signup.component";
import {LoginComponent} from "./login/login.component";
import {AdminCoucillorComponent} from "./admin/admin.councillor.component";
import {SignupService} from "./signup/signup.service";
import {ProfileComponent} from "./profile/profile.component";
import {ProfileService, UserData} from "./profile.service";
import {IsLoggedIn} from './is-logged-in.directive'
import {Isguest} from './isguest.directive'
import {Isadmin} from './Isadmin.directive'

declare var gapi: any;
@Component({
  moduleId: module.id,
  selector: 'app-root',
  templateUrl: 'app.component.html',
  styleUrls: ['app.component.css'],
  directives: [ROUTER_DIRECTIVES, IsLoggedIn, Isguest, Isadmin],
  precompile: [HomeComponent, SignupComponent, LoginComponent, ProfileComponent, AdminCoucillorComponent],
  providers: [SignupService, ProfileService]
})
export class AppComponent implements OnInit {

  @ViewChild('home') home: ElementRef

  constructor(private profileService: ProfileService, private router: Router, private renderer: Renderer, private _zone: NgZone) {

  }

  userData: UserData = new UserData("", "","");

  signOut() {
    this._zone.run(() => {
      var self = this;
      let event = new MouseEvent('click', { bubbles: true })
      self.profileService.deleteUserData()
        .then(() => {
          gapi.load('auth2', function () {
            gapi.auth2.init();
            let auth2 = gapi.auth2.getAuthInstance();
            if (auth2 && auth2.signOut) {
              auth2.then(() => {
                auth2.signOut().then(function () {
                  self.renderer.invokeElementMethod(self.home.nativeElement, 'click', [event]);
                  console.log('User signed out.');
                });
              });
            }
          });
        })
        .catch((e)=>console.error);
    });
  }

  ngAfterViewInit() {

  }
  ngOnInit() {

    this.profileService.getUserData().
      then((data) => {
        this.userData = data;
      })
      .catch((err) => {
        console.error("error getting user data", err);
      });
  }
}

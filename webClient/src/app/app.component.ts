import { Component, OnInit, OnDestroy, ViewChild, ElementRef, Renderer, NgZone} from '@angular/core';
import { ROUTER_DIRECTIVES, Router } from '@angular/router';
import {HomeComponent} from "./home/home.component";
import {SignupComponent} from "./signup/signup.component";
import {LoginComponent} from "./login/login.component";
import {AdminCoucillorComponent} from "./admin/admin.councillor.component";
import {SignupService} from "./signup/signup.service";
import {LoginService} from "./login/login.service";
import {ProfileComponent} from "./profile/profile.component";
import {CouncillorsComponent} from "./councillors/councillors.component";
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
  precompile: [HomeComponent, SignupComponent, LoginComponent, ProfileComponent, AdminCoucillorComponent, CouncillorsComponent],
  providers: [SignupService, ProfileService, LoginService]
})
export class AppComponent implements OnInit, OnDestroy {

  @ViewChild('home') home: ElementRef

  constructor(
    private profileService: ProfileService,
    private loginService: LoginService,
    private router: Router,
    private renderer: Renderer, private _zone: NgZone) {

  }

  userData: UserData;

  isLoggedIn: Boolean = false;

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
                  self.isLoggedIn = false;
                  self.userData = null;
                  self.renderer.invokeElementMethod(self.home.nativeElement, 'click', [event]);
                });
              });
            }
          });
        })
        .catch((e) => console.error);
    });
  }
  ngOnDestroy() {
    console.log("ngOnDestroy app.component")
  }
  resetUserData(){
    this.profileService.getUserData().then((data)=>{
        console.log("userData AppComponent",data);
        if (data) {
          this.userData = data;
          this.isLoggedIn = true
        }
      })
      .catch((err) => {
        console.error("error getting user data", err);
      }); 
  }

  ngOnInit() {
    this.loginService.loggedIn$.subscribe((value) => { 
      this.isLoggedIn = value;
    });
    this.profileService.userDataObservable.subscribe((value)=>{
      console.log("userdata from observer", value);
      this.userData = value;
    });
  }
}

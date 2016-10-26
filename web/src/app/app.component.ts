import { Component, OnInit, OnDestroy, ViewChild, ElementRef, Renderer, NgZone} from '@angular/core';
import {SignupService} from "./signup/signup.service";
import {LoginService} from "./login/login.service";
import {ProfileService, UserData} from "./profile.service";
import {CommunicationsService} from './communications.service'

declare var gapi: any;
@Component({
  selector: 'app-root',
  templateUrl: 'app.component.html',
  styleUrls: ['app.component.css'],
  providers: [SignupService, ProfileService, LoginService, CommunicationsService]
})
export class AppComponent implements OnInit, OnDestroy {

  @ViewChild('home') home: ElementRef

  constructor(
    private profileService: ProfileService,
    private loginService: LoginService,
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
      this.userData = value;
    });
    this.resetUserData();
  }
}

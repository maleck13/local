import { Component, OnInit, NgZone } from '@angular/core';
import {SignupService} from "./signup.service";
import {Signup} from "./signup.service";
import {ROUTER_DIRECTIVES, Router} from '@angular/router';
import {LoginService, Login} from '../login/login.service'
import {ProfileService, UserData} from '../profile.service'

declare var gapi: any;

@Component({
  moduleId: module.id,
  selector: 'app-signup',
  templateUrl: 'signup.component.html',
  styleUrls: ['signup.component.css']
})
export class SignupComponent implements OnInit {

  private message: string;
  private signUpError: string;

  constructor(private _zone: NgZone, private signUpService: SignupService,
    private router: Router, private loginService: LoginService,
    private profileService: ProfileService) { }

  ngOnInit() {
    this.message = "Signup Message"
  }

  ngAfterViewInit() {
    // Converts the Google login button stub to an actual button.
    var self = this;
    gapi.load('auth2', function () {
      gapi.auth2.init();
      gapi.signin2.render(
        "google-signin",
        {
          "onSuccess": self.onGoogleLoginSuccess,
          "scope": "profile"
        });
    });

  }

  // Triggered after a user successfully logs in using the Google external
  // login provider.
  onGoogleLoginSuccess = (loggedInUser) => {
    this._zone.run(() => {
      this.userAuthToken = loggedInUser.getAuthResponse().id_token;
      this.userDisplayName = loggedInUser.getBasicProfile().getName();
      this.userEmail = loggedInUser.getBasicProfile().getEmail();
      this.givenName = loggedInUser.getBasicProfile().getGivenName();
      this.familyName = loggedInUser.getBasicProfile().getFamilyName();
      let model = new Signup("google", this.userAuthToken, this.userEmail, this.givenName, this.familyName);

      this.signUpService.signUp(model)
        .then(() => {
          this.loginService.login(new Login(model.token, model.email, model.signUpType))
            .then((res) => {
              let ud = new UserData(res.id, res.token, res.type);
              let pStore = this.profileService.storeUserData(ud)
              pStore.then(() => {
                this.router.navigate(["/profile/" + res.id],{"queryParams":{"newSignUp":true}});
              });
              pStore.catch((err) => {
                console.error("signup error", err);
                this.signUpError = err.Message;
              });
            })
            .catch((err)=>{
              this.signUpError = "Failed to log you in after sign up. Please try from the login page.";  
            });
        })
        .catch((err) => {
          console.log("sign up error", err);
          if (err.status && err.status == 409) {
            this.signUpError = "You have already registered. Please sign in";
          } else {
            this.signUpError = "unexpected error occurred. Please try again later"
          }
        });

    });
  }

}

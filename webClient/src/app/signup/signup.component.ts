import { Component, OnInit,NgZone } from '@angular/core';
import {SignupService} from "./signup.service";
import {Signup} from "./signup.service";

declare var gapi:any;

@Component({
  moduleId: module.id,
  selector: 'app-signup',
  templateUrl: 'signup.component.html',
  styleUrls: ['signup.component.css']
})
export class SignupComponent implements OnInit {

  private message :string;

  constructor(private _zone: NgZone, private signUpService: SignupService) { }

  ngOnInit() {
     this.message = "Signup Message"
  }

  ngAfterViewInit() {
    // Converts the Google login button stub to an actual button.
    var self = this;
    gapi.load('auth2', function() {
      gapi.auth2.init();
      gapi.signin2.render(
          "google-signin",
          {
            "onSuccess": self.onGoogleLoginSuccess,
            "scope": "profile",
            "theme": "dark"
          });
    });

  }

  // Triggered after a user successfully logs in using the Google external
  // login provider.
  onGoogleLoginSuccess = (loggedInUser) => {
    console.log("here");
    this._zone.run(() => {
      this.userAuthToken = loggedInUser.getAuthResponse().id_token;
      this.userDisplayName = loggedInUser.getBasicProfile().getName();
      this.userEmail = loggedInUser.getBasicProfile().getEmail();
      this.givenName = loggedInUser.getBasicProfile().getGivenName();
      this.familyName = loggedInUser.getBasicProfile().getFamilyName();

      console.log("here", loggedInUser.getBasicProfile());

      console.log(this.userDisplayName, this.userAuthToken, this.userEmail);
      let model = new Signup("google",this.userAuthToken,this.userEmail,this.givenName, this.familyName);
      this.signUpService.signUp(model)
          .then((res) => {
              console.log(res)
              //redirect to complete sign up
          })
      .catch((err) => {
        console.log("sign up error",err);
          if (err.status && err.status == 409){
              this.message = "You have already registered. Please sign in";
          }else{
              this.message = "unexpected error occurred. Please try again later"
          }
      });

    });
  }

}

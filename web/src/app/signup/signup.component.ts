import { Component, OnInit, NgZone, Renderer, ViewChildren, QueryList, ElementRef} from '@angular/core';
import {SignupService} from "./signup.service";
import {Signup} from "./signup.service";
import {Router} from '@angular/router';
import {LoginService, Login} from '../login/login.service'
import {ProfileService, UserData} from '../profile.service'

declare var gapi: any;

@Component({
  selector: 'app-signup',
  templateUrl: 'signup.component.html',
  styleUrls: ['signup.component.css']
})
export class SignupComponent implements OnInit {

  @ViewChildren('userType') viewChildren: QueryList<ElementRef>;

  private message: string;
  private signUpError: string;

  private userAuthToken: string;
  private userDisplayName: string;
  private userEmail: string;
  private givenName: string;
  private familyName: string;
  private userSignupType: string = "local";
  private signUp: Signup = new Signup(null,null,null,null,null,null,null)
   private counties = ["Waterford"]
  private electorialAreas = { "Waterford": ["Tramore Waterford City West", "Waterford City East", "Waterford City South"] }
  private signUpMessage:string; 

  constructor(private _zone: NgZone, private signUpService: SignupService,
    private router: Router, private loginService: LoginService,
    private profileService: ProfileService, private render: Renderer) { }

  ngOnInit() {
    this.message = "Signup either using the social login or using the form provided."
  }

  ngAfterViewInit() {
    var self = this;
    // Converts the Google login button stub to an actual button.
    gapi.load('auth2', function () {
      gapi.auth2.init();
      gapi.signin2.render(
        "google-signin",
        {
          "onSuccess": self.onGoogleLoginSuccess,
          "scope": "profile",
          "onfailure": (err) => {
            self.signUpError = err;
          }
        });
    });

  }

  loadGoogleAuth() {
    var self = this;
    gapi.signin2.render(
      "google-signin",
      {
        "onSuccess": self.onGoogleLoginSuccess,
        "scope": "profile",
        "onfailure": (err) => {
          self.signUpError = err;
        }
      });
  }

  userTypeSelect(event: any) {
    this.signUpMessage = null;
    event.preventDefault()
    let t = event.srcElement.innerText.toLowerCase();
    this.viewChildren.map((e) => {
      e.nativeElement.className = "inactive"
      if (t.toLowerCase() == e.nativeElement.id.toLowerCase()) {
        e.nativeElement.className = "active"
      }
    })
    this.userSignupType = t;
    setTimeout(() => { //needed as the elment may not be immediately present.
      if (t === "local") {
        this.loadGoogleAuth();
      }
    }, 200);
  }

  manualSignup(){
    this.signUp.signUpType = "local";
    this.signupAndLogin(this.signUp);
  }

  councillorSignUpCheck(){
    console.log(this.signUp.email)
    this.signUpService.councillorUserExists(this.signUp.email)
    .then((res)=>{
        console.log("exists ",res);
        this.signUpMessage = "A verification email has been sent. It should arrive in your mail shortly.";
    })
    .catch((err)=>{
       this.signUpMessage = "No councillor. Do you have the right email address? Please feel free to contact support@locals.ie";
    })
  }

  signupAndLogin(signUp: Signup){
    this.signUpService.signUp(signUp)
        .then(() => {
          this.loginService.login(new Login(signUp.token, signUp.email, signUp.signUpType))
            .then((res) => {
              let ud = new UserData(res.id, res.token, res.type, "", "");
              let pStore = this.profileService.storeUserData(ud)
              pStore.then(() => {
                this.router.navigate([res.type + "/profile/" + res.id], { "queryParams": { "newSignUp": true } });
              });
              pStore.catch((err) => {
                console.error("signup error", err);
                this.signUpError = err.Message;
              });
            })
            .catch((err) => {
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
      this.signUp = new Signup("google", this.userAuthToken, this.userEmail, this.givenName, this.familyName,"local",null);
      this.signupAndLogin(this.signUp);
    });
  }

}
